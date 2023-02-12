package futures

import (
	"context"
	"fmt"
	"net/http"

	"github.com/adshao/go-binance/v2/common"
)

// KlinesService list klines
type KlinesService struct {
	c         *Client
	symbol    string
	interval  string
	limit     *int
	startTime *int64
	endTime   *int64
}

// Symbol set symbol
func (s *KlinesService) Symbol(symbol string) *KlinesService {
	s.symbol = symbol
	return s
}

// Interval set interval
func (s *KlinesService) Interval(interval string) *KlinesService {
	s.interval = interval
	return s
}

// Limit set limit
func (s *KlinesService) Limit(limit int) *KlinesService {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *KlinesService) StartTime(startTime int64) *KlinesService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *KlinesService) EndTime(endTime int64) *KlinesService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *KlinesService) Do(ctx context.Context, opts ...RequestOption) (res []*Kline, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/klines",
	}
	r.setParam("symbol", s.symbol)
	r.setParam("interval", s.interval)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*Kline{}, err
	}
	j, err := newJSON(data)
	if err != nil {
		return []*Kline{}, err
	}
	num := len(j.MustArray())
	res = make([]*Kline, num)
	for i := 0; i < num; i++ {
		item := j.GetIndex(i)
		if len(item.MustArray()) < 11 {
			err = fmt.Errorf("invalid kline response")
			return []*Kline{}, err
		}

		res[i] = &Kline{
			OpenTime:                 item.GetIndex(0).MustInt64(),
			Open:                     common.ParseFloat64Str(item.GetIndex(1).MustString()),
			High:                     common.ParseFloat64Str(item.GetIndex(2).MustString()),
			Low:                      common.ParseFloat64Str(item.GetIndex(3).MustString()),
			Close:                    common.ParseFloat64Str(item.GetIndex(4).MustString()),
			Volume:                   common.ParseFloat64Str(item.GetIndex(5).MustString()),
			CloseTime:                item.GetIndex(6).MustInt64(),
			QuoteAssetVolume:         common.ParseFloat64Str(item.GetIndex(7).MustString()),
			TradeNum:                 item.GetIndex(8).MustInt64(),
			TakerBuyBaseAssetVolume:  common.ParseFloat64Str(item.GetIndex(9).MustString()),
			TakerBuyQuoteAssetVolume: common.ParseFloat64Str(item.GetIndex(10).MustString()),
		}
	}
	return res, nil
}

// Kline define kline info
type Kline struct {
	OpenTime                 int64          `json:"openTime"`
	Open                     common.Float64 `json:"open"`
	High                     common.Float64 `json:"high"`
	Low                      common.Float64 `json:"low"`
	Close                    common.Float64 `json:"close"`
	Volume                   common.Float64 `json:"volume"`
	CloseTime                int64          `json:"closeTime"`
	QuoteAssetVolume         common.Float64 `json:"quoteAssetVolume"`
	TradeNum                 int64          `json:"tradeNum"`
	TakerBuyBaseAssetVolume  common.Float64 `json:"takerBuyBaseAssetVolume"`
	TakerBuyQuoteAssetVolume common.Float64 `json:"takerBuyQuoteAssetVolume"`
}
