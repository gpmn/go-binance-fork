package futures

import (
	"context"
	"encoding/json"
	"net/http"
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
		endpoint: "/fapi/v1/account",
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
	CanDeposit                  bool               `json:"canDeposit"`
	CanTrade                    bool               `json:"canTrade"`
	CanWithdraw                 bool               `json:"canWithdraw"`
	FeeTier                     int                `json:"feeTier"`
	MaxWithdrawAmount           string             `json:"maxWithdrawAmount"`
	Positions                   []*AccountPosition `json:"positions"`
	TotalInitialMargin          string             `json:"totalInitialMargin"`
	TotalMaintMargin            string             `json:"totalMaintMargin"`
	TotalMarginBalance          string             `json:"totalMarginBalance"`
	TotalOpenOrderInitialMargin string             `json:"totalOpenOrderInitialMargin"`
	TotalPositionInitialMargin  string             `json:"totalPositionInitialMargin"`
	TotalUnrealizedProfit       string             `json:"totalUnrealizedProfit"`
	TotalWalletBalance          string             `json:"totalWalletBalance"`
	UpdateTime                  int64              `json:"updateTime"`
}

// AccountAsset define account asset
type AccountAsset struct {
	Asset                  string  `json:"asset"`
	InitialMargin          Float64 `json:"initialMargin"`
	MaintMargin            Float64 `json:"maintMargin"`
	MarginBalance          Float64 `json:"marginBalance"`
	MaxWithdrawAmount      Float64 `json:"maxWithdrawAmount"`
	OpenOrderInitialMargin Float64 `json:"openOrderInitialMargin"`
	PositionInitialMargin  Float64 `json:"positionInitialMargin"`
	UnrealizedProfit       Float64 `json:"unrealizedProfit"`
	WalletBalance          Float64 `json:"walletBalance"`
}

// AccountPosition define account position
type AccountPosition struct {
	Isolated               bool             `json:"isolated"`
	Leverage               string           `json:"leverage"`
	InitialMargin          Float64          `json:"initialMargin"`
	MaintMargin            Float64          `json:"maintMargin"`
	OpenOrderInitialMargin Float64          `json:"openOrderInitialMargin"`
	PositionInitialMargin  Float64          `json:"positionInitialMargin"`
	Symbol                 string           `json:"symbol"`
	UnrealizedProfit       Float64          `json:"unrealizedProfit"`
	EntryPrice             Float64          `json:"entryPrice"`
	MaxNotional            Float64          `json:"maxNotional"`
	PositionSide           PositionSideType `json:"positionSide"`
	PositionAmt            Float64          `json:"positionAmt"`
	Notional               string           `json:"notional"`
	IsolatedWallet         string           `json:"isolatedWallet"`
	UpdateTime             int64            `json:"updateTime"`
}
