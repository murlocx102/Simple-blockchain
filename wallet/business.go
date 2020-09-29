package wallet

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

type Business struct {
	wallet *hdwallet.Wallet  // 钱包
	client *ethclient.Client // 客户端
}

// NewBusiness 创建交易服务
func NewBusiness(hd *hdwallet.Wallet, client *ethclient.Client) *Business {
	// infura提供公开的Ethereum主网和测试网络节点，到infura.io网站注册后,建立项目即可获取各个网络的节点地址
	if client == nil {
		client, _ = ethclient.Dial("https://kovan.infura.io/v3/6dd6106f58e444f3bbc42e0d32ea6079") // 使用kovan测试节点
	}

	return &Business{
		wallet: hd,
		client: client,
	}
}

// 注意:测试环境的代币或eth交易.均需要 gas.账户需保证有eth.请参考 coin文件中的md.

// TurnOut 代币交易
func (b *Business) TurnOut(account accounts.Account, contractAddress, toAddress string, value float64, nonceNum int64) (string, error) {
	nonce := uint64(nonceNum)                 // 一个序列编号,由构建这个交易的外部账号提供(需跟踪nonce数),用于防止交易的重放攻击。(节点会以任意顺序接收交易;如果nonce是3,那么会直到nonce值为0到2的交易已经被处理,即便它是先接收到的.)
	amount := big.NewInt(0)                   // 如果是代币传输,则不需要传输ETH，因此将交易数量“值”设置为“0”。
	address := common.HexToAddress(toAddress) // 交易接收方以太坊地址
	var data []byte                           // 附在交易中的可变长度的数据。(只包含data的交易是针对合约的调用)

	gasLimit, err := b.client.EstimateGas(context.Background(), ethereum.CallMsg{ // 估算的完成交易所需的估计gas上限(交易发起方愿意为这个交易支付的最大gas数量)
		To:   &address,
		Data: data,
	})
	if err != nil {
		return "", err
	}

	gasLimit = gasLimit + 50000 // 测试网络涨价了...

	gasPrice, err := b.client.SuggestGasPrice(context.Background()) // 根据先前块的x个数来获取平均gas价格(交易发起方愿意支付的gas价格(单位:wei))
	if err != nil {
		return "", err
	}

	// 转账针对的就是账户里的dc代币,contractAddress 是代币地址
	// 具体看这里 https://kovan.etherscan.io/token/0x4354c7002ffe5b399b9199b8c32a1f1a5c26f94d
	tokenAddress := common.HexToAddress(contractAddress)

	// 编写代币合约data
	data = b.CoinOptData(address, value)

	// 构建无符号以太坊交易事务类型
	tx := types.NewTransaction(nonce, tokenAddress, amount, gasLimit, gasPrice, data)

	// 使用发起者账户的私钥对交易进行签名
	// 代币交易需要EIP155签名,还需要从客户端拿到链ID
	chainID, err := b.client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	signedTx, err := b.wallet.SignTxEIP155(account, tx, chainID)
	if err != nil {
		return "", err
	}

	// 通过调用接收已签名交易的客户端上的SendTransaction将交易广播到整个网络
	err = b.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	// 交易tx哈希
	txSent := signedTx.Hash().Hex()

	return txSent, nil
}

// EthTurnOut 以太币交易转出
func (b *Business) EthTurnOut(account accounts.Account, toAddress string, value float64, nonceNum int64) (string, error) {
	nonce := uint64(nonceNum)                 // 一个序列编号,由构建这个交易的外部账号提供(需跟踪nonce数),用于防止交易的重放攻击。(节点会以任意顺序接收交易;如果nonce是3,那么会直到nonce值为0到2的交易已经被处理,即便它是先接收到的.)
	address := common.HexToAddress(toAddress) // 交易接收方以太坊地址
	gasLimit := uint64(21000)                 // 交易发起方愿意为这个交易支付的最大gas数量。应设上限21000
	var data []byte                           // 附在交易中的可变长度的数据。(只包含data的交易是针对合约的调用,value为0)

	gasPrice, err := b.client.SuggestGasPrice(context.Background()) // 根据先前块的x个数来获取平均gas价格(交易发起方愿意支付的gas(单位:wei)价格)
	if err != nil {
		return "", err
	}

	//这是处理位数的代码段 1个ETH为1加18个零
	amountFloat := big.NewFloat(value)
	bigInt := new(big.Int)
	amount, _ := amountFloat.Mul(amountFloat, big.NewFloat(1000000000000000000)).Int(bigInt)

	// 构建无符号以太坊交易事务类型
	tx := types.NewTransaction(nonce, address, amount, gasLimit, gasPrice, data)

	signedTx, err := b.wallet.SignTx(account, tx, nil)
	if err != nil {
		return "", err

	}

	// 通过调用接收已签名交易的客户端上的SendTransaction将交易广播到整个网络
	err = b.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	// 交易tx哈希
	txSent := signedTx.Hash().Hex()

	return txSent, nil
}

// CoinOptData 代币合约转换
func (b *Business) CoinOptData(toAddress common.Address, value float64) []byte {
	// 函数名将是传递函数的名称，即ERC-20规范中的transfer和参数类型。
	transferFnSignature := []byte("transfer(address,uint256)")

	// 生成函数签名的Keccak256哈希
	hashByte := crypto.Keccak256(transferFnSignature)

	// 1.获取前4个字节(8字符)的方法ID
	methodID := hashByte[:4]
	//fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

	// 2.将接收代币(目标)方地址 左填充到32字节
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)

	// 3.发送多少个代币,格式化为wei(18个0) 1000000000000000000 = 1eth
	amountFloat := big.NewFloat(value)
	bigInt := new(big.Int)
	amount, _ := amountFloat.Mul(amountFloat, big.NewFloat(1000000000000000000)).Int(bigInt)
	// 4.代币量同样左填充到32个字节
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	// 5.合并方法ID,填充后的地址和填后的转账量,到将成为我们数据字段的字节片
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	return data
}
