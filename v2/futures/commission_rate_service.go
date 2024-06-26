package futures

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/adshao/go-binance/v2/common"
)

type CommissionRateService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (service *CommissionRateService) Symbol(symbol string) *CommissionRateService {
	service.symbol = symbol
	return service
}

// Do send request
func (s *CommissionRateService) Do(ctx context.Context, opts ...RequestOption) (*CommissionRate, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/commissionRate",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res CommissionRate
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// Commission Rate
type CommissionRate struct {
	Symbol              string         `json:"symbol"`
	MakerCommissionRate common.Float64 `json:"makerCommissionRate"`
	TakerCommissionRate common.Float64 `json:"takerCommissionRate"`
}
