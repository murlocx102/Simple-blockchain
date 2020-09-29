package wallet

import (
	"log"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

type HDWalletService struct {
	wallet *hdwallet.Wallet
}

// NewHDWalletService 创建hd钱包服务
func NewHDWalletService() HDWallet {
	return &HDWalletService{}
}

// CreateMnemonic 生成助记词  bits 128=12个单词  256=24个单词
func (h *HDWalletService) CreateMnemonic(bits int) (string, error) {
	entropy, err := hdwallet.NewEntropy(bits)
	if err != nil {
		log.Fatal(err)
	}

	mnemonic, err := hdwallet.NewMnemonicFromEntropy(entropy)
	if err != nil {
		log.Fatal(err)
	}

	return mnemonic, nil
}

// CreateHDWallet 创建hd钱包(真实使用应只执行一次)
func (h *HDWalletService) CreateHDWallet(mnemonic string) error {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return err
	}

	// 创建主账户,使用了i:非增强索引码,i‘:表示增强索引码
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	// pin 是否记录账户到钱包中
	_, err = wallet.Derive(path, true)
	if err != nil {
		return err
	}

	h.wallet = wallet

	return nil
}

// GetHDWallet 获取HD钱包
func (h *HDWalletService) GetHDWallet() *hdwallet.Wallet {
	return h.wallet
}
