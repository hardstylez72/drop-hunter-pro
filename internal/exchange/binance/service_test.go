package binance

import (
	"context"
	"testing"
)

func TestService(t *testing.T) {
	//cfg := tests.GetConfig()
	//c := NewService(cfg.BinanceKey, cfg.BinanceSecret, cfg.CliL1)
	//
	//c.GetWithDrawOptions(context.Background())
	//t.Run("assets", func(t *testing.T) {
	//	ctx := context.Background()
	//	c.GetUserAsset(ctx, "ETH")
	//})

	c := NewWsClient()
	c.GetPairRate(context.Background(), "STGUSDT")

	//t.Run("networks", func(t *testing.T) {
	//	ctx := context.Background()
	//	c.GetTokenAvailableNetworks(ctx, []string{"ETH"})
	//})

}
