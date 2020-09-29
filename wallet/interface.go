package wallet

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

type HDWallet interface {
	// 生成助记词
	CreateMnemonic(bits int) (string, error)
	// 生成HD钱包,且已生成主账户
	CreateHDWallet(mnemonic string) error
	// 获取HD钱包实例
	GetHDWallet() *hdwallet.Wallet
}

type Accounts interface {
	// 增加子账户
	AddAccount(accountID int) (accounts.Account, error)
	// 获取所有账户
	GetAllAccounts() []accounts.Account
	// 获取账户地址
	Address(account accounts.Account) string
	// 根据地址查询账户
	FindAccountByAddress(address string) (accounts.Account, error)
	// 所有账户余额
	GetAllBalance() map[string]*big.Float
	// 获取指定账户私钥
	PrivateKey(account accounts.Account) (string, error)
	// 根据指定账户生成一个新可用地址(本意就是一个账户下属的地址.而账户主地址是0索引号.而这个建立的是0以后的索引号.)
	CreateReceiveAddress(account accounts.Account, indexID int, enhance bool) (string, error)
}
