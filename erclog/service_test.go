package erclog

import (
	"os"
	"testing"
)

var (
	testLog   *LogService
	tokenAddr       = "0x4354c7002ffe5b399b9199b8c32a1f1a5c26f94d" // 币地址
	Addr            = "0xC56bE5E6B20F6cf225A9ff4412d7cCEcd53e3037" // 目标账户 1-0
	fromBlock int64 = 21230135                                     // 起始块
	toBlock   int64 = 21230135                                     // 结束块
)

func TestMain(m *testing.M) {
	// 创建币查询服务
	testLog = NewLogService(nil)

	os.Exit(m.Run())
}

func Test_TransactionLog(t *testing.T) {
	testLog.TransactionLog(tokenAddr, Addr, fromBlock, toBlock)
}

/*
交易记录
Log Block Number: 21230135
Log Index: 0
Log Name: Transfer
Log Block Number: 21230135
Log Index: 0
Log Hash: 0xbc9fc83d39ce43841325b387f51d17cf56172289c3b56ccefa0c5625a33d5fb0
From: 0xC56bE5E6B20F6cf225A9ff4412d7cCEcd53e3037
To: 0x2D72D0e085447bBf86E098f5A0e644a9F25f64FE
Tokens: 100000500000000000000000
*/
