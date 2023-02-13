package futures

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/adshao/go-binance/v2/common"
)

// GetPositionRiskService get account balance
type GetPositionRiskService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *GetPositionRiskService) Symbol(symbol string) *GetPositionRiskService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *GetPositionRiskService) Do(ctx context.Context, opts ...RequestOption) (res []*PositionRisk, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v2/positionRisk",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*PositionRisk{}, err
	}
	res = make([]*PositionRisk, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*PositionRisk{}, err
	}
	return res, nil
}

// PositionRisk define position risk info
type PositionRisk struct {
	EntryPrice       common.Float64 `json:"entryPrice"`
	MarginType       string         `json:"marginType"`
	IsAutoAddMargin  string         `json:"isAutoAddMargin"`
	IsolatedMargin   common.Float64 `json:"isolatedMargin"`
	Leverage         common.Float64 `json:"leverage"`
	LiquidationPrice common.Float64 `json:"liquidationPrice"`
	MarkPrice        common.Float64 `json:"markPrice"`
	MaxNotionalValue common.Float64 `json:"maxNotionalValue"`
	PositionAmt      common.Float64 `json:"positionAmt"`
	Symbol           string         `json:"symbol"`
	UnRealizedProfit common.Float64 `json:"unRealizedProfit"`
	PositionSide     string         `json:"positionSide"`
	Notional         string         `json:"notional"`
	IsolatedWallet   string         `json:"isolatedWallet"`
}
