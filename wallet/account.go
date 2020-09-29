package wallet

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

const (
	AccountPathFmt = "m/44'/60'/%s'/0/0"
)

type AccountService struct {
	wallet *hdwallet.Wallet
}

// NewAccountService 创建账户服务
func NewAccountService(hd *hdwallet.Wallet) Accounts {
	return &AccountService{
		wallet: hd,
	}
}

// AddAccount 增加子账户
func (a *AccountService) AddAccount(accountID int) (accounts.Account, error) {
	pathFmt := fmt.Sprintf(AccountPathFmt, strconv.Itoa(accountID))
	path := hdwallet.MustParseDerivationPath(pathFmt)

	account, err := a.wallet.Derive(path, true)
	if err != nil {
		return account, err
	}

	return account, nil
}

// GetAllAccounts 获取所有账户
func (a *AccountService) GetAllAccounts() []accounts.Account {
	return a.wallet.Accounts()
}

// Address 获取账户地址
func (a *AccountService) Address(account accounts.Account) string {
	return account.Address.Hex()
}

// FindAccountByAddress 根据地址查询账户
func (a *AccountService) FindAccountByAddress(address string) (accounts.Account, error) {
	panic("not implemented") // TODO: Implement
}

// GetAllBalance 所有账户余额
func (a *AccountService) GetAllBalance() map[string]*big.Float {
	panic("not implemented") // TODO: Implement
}

// PrivateKey 获取指定账户私钥
func (a *AccountService) PrivateKey(account accounts.Account) (string, error) {
	// 返回十六进制格式私钥
	result, err := a.wallet.PrivateKeyHex(account)
	if err != nil {
		return "", err
	}

	// 标明是十六进制
	return fmt.Sprintf("%s%s", "0x", result), nil
}

// CreateReceiveAddress 根据指定账户生成一个新可用地址(本意就是一个账户下属的地址.而账户主地址是0索引号.而这个建立的是0以后的索引号.)
func (a *AccountService) CreateReceiveAddress(account accounts.Account, indexID int, enhance bool) (string, error) {
	accountPath := account.URL.Path
	//  m/44'/60'/1'/0/0
	strArr := strings.Split(accountPath, "/")
	if len(strArr) != 6 {
		return "", errors.New("path error!")
	}

	// 构建是否增强算法索引号
	if enhance {
		strArr[5] = fmt.Sprintf("%s%s", strconv.Itoa(indexID), "'")
	} else {
		strArr[5] = strconv.Itoa(indexID)
	}
	pathStr := strings.Join(strArr, "/")

	path := hdwallet.MustParseDerivationPath(pathStr)
	// 生成指定账户新的可用地址
	account, err := a.wallet.Derive(path, true)
	if err != nil {
		return "", err
	}

	return account.Address.Hex(), nil
}
