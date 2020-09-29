package coin

import (
	"context"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	abiToken "github.com/simple-blockchain/abi"
)

// CoinService 代币相关操作
type CoinService struct {
	client *ethclient.Client // 客户端
}

// NewCoinService 创建代币服务
func NewCoinService(client *ethclient.Client) *CoinService {
	if client == nil {
		client, _ = ethclient.Dial("https://kovan.infura.io/v3/6dd6106f58e444f3bbc42e0d32ea6079") // 使用kovan测试节点
	}

	return &CoinService{
		client: client,
	}
}

// QueryCoinDetail 代币详情查询
func (c *CoinService) QueryCoinDetail(tokenAddr, accountAddr string) (string, string, error) {
	// 将指定token导入并实例化.
	tokenAddress := common.HexToAddress(tokenAddr)
	instance, err := abiToken.NewToken(tokenAddress, c.client)
	if err != nil {
		return "", "", err
	}

	address := common.HexToAddress(accountAddr)
	// 查询用户的代币余额单位(wei)
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		return "", "", err
	}

	// 获取代币公共信息
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		return "", "", err
	}

	// 将代币余额(wei)转换为可读的十进制格式
	tokenBal := ConversionFormat(bal, decimals)

	// 查询用户eth余额
	ethBal, err := c.client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		return "", "", err
	}

	return tokenBal, ConversionFormat(ethBal, 18), nil
}

// ConversionFormat 转换格式化
func ConversionFormat(value *big.Int, decimals uint8) string {
	fBal := new(big.Float)
	fBal.SetString(value.String())
	result := new(big.Float).Quo(fBal, big.NewFloat(math.Pow10(int(decimals))))

	// result.Text('g', 10) 对于“ g”和“ G”，10是数字的总数(更高精度使用更大的数字)
	return result.String() // 输出10位数字
}
