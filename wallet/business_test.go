package wallet

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	contractAddress         = "0x4354c7002ffe5b399b9199b8c32a1f1a5c26f94d" // 币地址
	toAddress               = "0x47cd8e6dE32278774C36b524B6120b9D467699EE" // 目标账户 1-1
	toAddress2              = "0xC56bE5E6B20F6cf225A9ff4412d7cCEcd53e3037" // 目标账户 1-0
	value           float64 = 100.1                                        // 代币量
	nonceNum        int64   = 7
)

func Test_TurnOut(t *testing.T) {
	account, err := testAccount.AddAccount(1) // 1-0
	require.NoError(t, err, "子账户")

	result, err := testBusiness.TurnOut(account, contractAddress, toAddress, value, nonceNum)
	require.NoError(t, err, "代币转出")

	t.Log(result) // 0xc88b75ccdbfda4715b3b2e7bbaa30b4527f3d8330156592fa16fa8fce5ac3945
}

func Test_EthTurnOut(t *testing.T) {
	accounts := testAccount.GetAllAccounts() // 0-0

	ethValue := 0.2

	result, err := testBusiness.EthTurnOut(accounts[0], toAddress2, ethValue, nonceNum)
	require.NoError(t, err, "测试以太币转出")

	t.Log(result) // 0xad3c213582ccadcd636dd3780845b87b7df60d2fa9ebe45b7ea39d5994397eba
}
