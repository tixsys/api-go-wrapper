package pos

import (
	common2 "github.com/erply/api-go-wrapper/pkg/api/common"
)

type (
	PointOfSale struct {
		PointOfSaleID uint   `json:"pointOfSaleID"`
		Name          string `json:"name"`
		WarehouseID   int    `json:"warehouseID"`
		WarehouseName string `json:"warehouseName"`
		Added         uint64 `json:"added"`
		LastModified  uint64 `json:"lastModified"`
		StoreHours    string `json:"storeHours"`
	}

	GetPointsOfSaleResponse struct {
		Status       common2.Status `json:"status"`
		PointsOfSale []PointOfSale  `json:"records"`
	}
)
