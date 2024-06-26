package futures

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/adshao/go-binance/v2/common"
)

// GetBalanceService get account balance
type GetBalanceService struct {
	c *Client
}

// Do send request
func (s *GetBalanceService) Do(ctx context.Context, opts ...RequestOption) (res []*Balance, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v2/balance",
		secType:  secTypeSigned,
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*Balance{}, err
	}
	res = make([]*Balance, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*Balance{}, err
	}
	return res, nil
}

// Balance define user balance of your account
type Balance struct {
	AccountAlias       string `json:"accountAlias"`
	Asset              string `json:"asset"`
	Balance            string `json:"balance"`
	CrossWalletBalance string `json:"crossWalletBalance"`
	CrossUnPnl         string `json:"crossUnPnl"`
	AvailableBalance   string `json:"availableBalance"`
	MaxWithdrawAmount  string `json:"maxWithdrawAmount"`
}

// GetAccountService get account info
type GetAccountService struct {
	c *Client
}

// Do send request
func (s *GetAccountService) Do(ctx context.Context, opts ...RequestOption) (res *Account, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v2/account",
		secType:  secTypeSigned,
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(Account)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Account define account info
type Account struct {
	Assets                      []*AccountAsset    `json:"assets"`
	FeeTier                     int                `json:"feeTier"`
	CanTrade                    bool               `json:"canTrade"`
	CanDeposit                  bool               `json:"canDeposit"`
	CanWithdraw                 bool               `json:"canWithdraw"`
	UpdateTime                  int64              `json:"updateTime"`
	TotalInitialMargin          common.Float64     `json:"totalInitialMargin"`
	TotalMaintMargin            common.Float64     `json:"totalMaintMargin"`
	TotalWalletBalance          common.Float64     `json:"totalWalletBalance"`
	TotalUnrealizedProfit       common.Float64     `json:"totalUnrealizedProfit"`
	TotalMarginBalance          common.Float64     `json:"totalMarginBalance"`
	TotalPositionInitialMargin  common.Float64     `json:"totalPositionInitialMargin"`
	TotalOpenOrderInitialMargin common.Float64     `json:"totalOpenOrderInitialMargin"`
	TotalCrossWalletBalance     common.Float64     `json:"totalCrossWalletBalance"`
	TotalCrossUnPnl             common.Float64     `json:"totalCrossUnPnl"`
	AvailableBalance            common.Float64     `json:"availableBalance"`
	MaxWithdrawAmount           common.Float64     `json:"maxWithdrawAmount"`
	Positions                   []*AccountPosition `json:"positions"`
}

// AccountAsset define account asset
type AccountAsset struct {
	Asset                  string         `json:"asset"`
	InitialMargin          common.Float64 `json:"initialMargin"`
	MaintMargin            common.Float64 `json:"maintMargin"`
	MarginBalance          common.Float64 `json:"marginBalance"`
	MaxWithdrawAmount      common.Float64 `json:"maxWithdrawAmount"`
	OpenOrderInitialMargin common.Float64 `json:"openOrderInitialMargin"`
	PositionInitialMargin  common.Float64 `json:"positionInitialMargin"`
	UnrealizedProfit       common.Float64 `json:"unrealizedProfit"`
	WalletBalance          common.Float64 `json:"walletBalance"`
	MarginAvailable        bool           `json:"marginAvailable"`
	CrossWalletBalance     common.Float64 `json:"crossWalletBalance"`
	CrossUnPnl             common.Float64 `json:"crossUnPnl"`
	AvailableBalance       common.Float64 `json:"availableBalance"`
	UpdateTime             int64          `json:"updateTime"`
}

// AccountPosition define account position
type AccountPosition struct {
	Isolated               bool             `json:"isolated"`
	Leverage               common.Float64   `json:"leverage"`
	InitialMargin          common.Float64   `json:"initialMargin"`
	MaintMargin            common.Float64   `json:"maintMargin"`
	OpenOrderInitialMargin common.Float64   `json:"openOrderInitialMargin"`
	PositionInitialMargin  common.Float64   `json:"positionInitialMargin"`
	Symbol                 string           `json:"symbol"`
	UnrealizedProfit       common.Float64   `json:"unrealizedProfit"`
	EntryPrice             common.Float64   `json:"entryPrice"`
	MaxNotional            common.Float64   `json:"maxNotional"`
	PositionSide           PositionSideType `json:"positionSide"`
	PositionAmt            common.Float64   `json:"positionAmt"`
	Notional               common.Float64   `json:"notional"`
	IsolatedWallet         common.Float64   `json:"isolatedWallet"`
	UpdateTime             int64            `json:"updateTime"`
}
