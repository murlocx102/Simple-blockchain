package wallet

import (
	"os"
	"testing"
)

var (
	testWallet   HDWallet
	testAccount  Accounts
	testBusiness *Business
)

func TestMain(m *testing.M) {
	// 创建钱包服务
	testWallet = NewHDWalletService()

	mnemonic := "simple comfort unfold paper season modify inherit chaos pen faith truck abstract" // 助记词已变更

	// 创建hd钱包.
	testWallet.CreateHDWallet(mnemonic)

	testAccount = NewAccountService(testWallet.GetHDWallet())

	// !钱包交易需要eth.
	testBusiness = NewBusiness(testWallet.GetHDWallet(), nil)

	os.Exit(m.Run())
}
