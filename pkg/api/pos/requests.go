package pos

import (
	"context"
	"encoding/json"
	"github.com/erply/api-go-wrapper/internal/common"
	erro "github.com/erply/api-go-wrapper/internal/errors"
)

// GetPointsOfSale will list points of sale according to specified filters.
func (cli *Client) GetPointsOfSale(ctx context.Context, filters map[string]string) ([]PointOfSale, error) {
	resp, err := cli.SendRequest(ctx, "getPointsOfSale", filters)
	if err != nil {
		return nil, err
	}
	var res GetPointsOfSaleResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("failed to unmarshal GetPointsOfSaleResponse", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewFromResponseStatus(&res.Status)
	}
	return res.PointsOfSale, nil
}

// GetDayClosings will retrieve a log of POS day openings and closings according to specified filters.
func (cli *Client) GetDayClosings(ctx context.Context, filters map[string]string) ([]DayClosing, error) {
	resp, err := cli.SendRequest(ctx, "getDayClosings", filters)
	if err != nil {
		return nil, err
	}
	var res GetDayClosingsResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("failed to unmarshal GetDayClosingsResponse", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewFromResponseStatus(&res.Status)
	}
	return res.DayClosings, nil
}

// GetCashIns will retrieve POS cash drops and cash payouts according to specified filters.
func (cli *Client) GetCashIns(ctx context.Context, filters map[string]string) ([]CashIn, error) {
	resp, err := cli.SendRequest(ctx, "getCashIns", filters)
	if err != nil {
		return nil, err
	}
	var res GetCashInsResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("failed to unmarshal GetCashInsResponse", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewFromResponseStatus(&res.Status)
	}
	return res.CashIns, nil
}

// GetReasonCodes will get a list of reason codes. according to specified filters.
func (cli *Client) GetReasonCodes(ctx context.Context, filters map[string]string) ([]ReasonCode, error) {
	resp, err := cli.SendRequest(ctx, "getReasonCodes", filters)
	if err != nil {
		return nil, err
	}
	var res GetReasonCodesResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("failed to unmarshal getReasonCodes", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewFromResponseStatus(&res.Status)
	}
	return res.ReasonCodes, nil
}
