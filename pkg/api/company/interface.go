package company

import (
	"context"
	"net/http"
)

type Manager interface {
	GetCompanyInfo(ctx context.Context) (*Info, error)
	GetConfParameters(ctx context.Context) (*ConfParameter, error)
	GetCustomConfParameters(ctx context.Context) (*http.Response, error)
}
