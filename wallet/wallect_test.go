package wallet

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GetMnemonic(t *testing.T) {
	bits := 128
	mnemonic, err := testWallet.CreateMnemonic(bits)

	require.NoError(t, err, "生成助记词")

	t.Log(mnemonic)
}

func Test_GetWallet(t *testing.T) {
	mnemonic := "simple comfort unfold paper season modify inherit chaos pen faith truck abstract"

	err := testWallet.CreateHDWallet(mnemonic)

	require.NoError(t, err, "创建hd钱包")
}
