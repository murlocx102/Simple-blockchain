package erclog

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 创建与ERC-20事件日志签名类型相匹配的结构类型
type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}

type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}

type LogService struct {
	client *ethclient.Client // 客户端
}

// NewLogService 创建日志服务
func NewLogService(client *ethclient.Client) *LogService {
	if client == nil {
		client, _ = ethclient.Dial("https://kovan.infura.io/v3/6dd6106f58e444f3bbc42e0d32ea6079") // 使用kovan测试节点
	}

	return &LogService{
		client: client,
	}
}

// TransactionLog 交易日志
func (l *LogService) TransactionLog(tokenAddr string, addr string, fromBlock, toBlock int64) {
	tokenAddress := common.HexToAddress(tokenAddr)

	// 创建所需的块范围的FilterQuery
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(fromBlock),
		ToBlock:   big.NewInt(toBlock),
		Addresses: []common.Address{
			tokenAddress,
		},
	}

	// 过滤日志,获取所需块日志
	logs, err := l.client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	// 解析JSON abi,用于解压缩原始日志数据
	contractAbi, err := abi.JSON(strings.NewReader(string(ErclogABI)))
	if err != nil {
		log.Fatal(err)
	}

	// 按某种日志类型进行过滤
	logTransferSig := []byte("Transfer(address,address,uint256)")
	LogApprovalSig := []byte("Approval(address,address,uint256)")

	// 计算keccak256哈希
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)

	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)
		// 事件日志函数签名哈希始终是topic [0]
		switch vLog.Topics[0].Hex() {
		case logTransferSigHash.Hex():
			// 解析Transfer事件日志
			var transferEvent LogTransfer

			err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

			// 获取指定账户的交易记录
			if strings.EqualFold(transferEvent.From.Hex(), addr) || strings.EqualFold(transferEvent.To.Hex(), addr) {
				fmt.Printf("Log Name: Transfer\n")
				fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
				fmt.Printf("Log Index: %d\n", vLog.Index)
				fmt.Printf("Log Hash: %s\n", vLog.TxHash.String())
				fmt.Printf("From: %s \n", transferEvent.From.Hex())
				fmt.Printf("To: %s \n", transferEvent.To.Hex())
				fmt.Printf("Tokens: %s\n", transferEvent.Tokens.String())
			}

		case logApprovalSigHash.Hex():
			fmt.Printf("Log Name: Approval\n")

			var approvalEvent LogApproval

			err := contractAbi.UnpackIntoInterface(&approvalEvent, "Approval", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			approvalEvent.TokenOwner = common.HexToAddress(vLog.Topics[1].Hex())
			approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("Token Owner: %s\n", approvalEvent.TokenOwner.Hex())
			fmt.Printf("Spender: %s\n", approvalEvent.Spender.Hex())
			fmt.Printf("Tokens: %s\n", approvalEvent.Tokens.String())
		}

		// indexed事件类型,它们存储在topics,必须单独解析
	}
}
