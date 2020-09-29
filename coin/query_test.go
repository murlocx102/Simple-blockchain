package coin

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	testCoin  *CoinService
	tokenAddr = "0x4354c7002ffe5b399b9199b8c32a1f1a5c26f94d" // 币地址
	Addr      = "0xC56bE5E6B20F6cf225A9ff4412d7cCEcd53e3037" // 目标账户 1-0
)

func TestMain(m *testing.M) {
	// 创建币查询服务
	testCoin = NewCoinService(nil)

	os.Exit(m.Run())
}

func Test_queryCoin(t *testing.T) {
	tokenNum, ethNum, err := testCoin.QueryCoinDetail(tokenAddr, Addr)
	require.NoError(t, err, "持币数量")

	t.Log(tokenNum, ethNum)
}
