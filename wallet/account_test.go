package wallet

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
)

func Test_addAccount(t *testing.T) {
	// 创建子账户
	account, err := testAccount.AddAccount(1)

	require.NoError(t, err, "创建子账户")

	// 0xC56bE5E6B20F6cf225A9ff4412d7cCEcd53e3037
	t.Log(testAccount.Address(account))
}

func Test_PrivateKey(t *testing.T) {
	accounts := testAccount.GetAllAccounts()

	for _, v := range accounts {
		priv, err := testAccount.PrivateKey(v)
		require.NoError(t, err, "获取账户私钥")

		t.Log(priv)
	}
}

func Test_getAllAccount(t *testing.T) {
	// 获取所有账户
	accounts := testAccount.GetAllAccounts()

	spew.Dump(accounts)
}

func Test_CreateReceiveAddress(t *testing.T) {
	accounts := testAccount.GetAllAccounts()
	for _, v := range accounts {
		result, err := testAccount.CreateReceiveAddress(v, 1, false)
		require.NoError(t, err, "获取账户可用索引地址私钥")

		// 0xb2c05Ec95B09fe0dDb0E01d412e4658565Af3274
		// 0x47cd8e6dE32278774C36b524B6120b9D467699EE
		t.Log(result)
	}

	Test_PrivateKey(t)
}
