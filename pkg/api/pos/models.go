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

	GetDayClosingsResponse struct {
		Status      common2.Status `json:"status"`
		DayClosings []DayClosing   `json:"records"`
	}

	DayClosing struct {
		DayID                int     `json:"dayID,omitempty"`
		WarehouseID          int     `json:"warehouseID,omitempty"`
		WarehouseName        string  `json:"warehouseName,omitempty"`
		PointOfSaleID        int     `json:"pointOfSaleID,omitempty"`
		PointOfSaleName      string  `json:"pointOfSaleName,omitempty"`
		DrawerID             *int    `json:"drawerID,omitempty"`
		OpenedUnixTime       int     `json:"openedUnixTime,omitempty"`
		OpenedByEmployeeID   int     `json:"openedByEmployeeID,omitempty"`
		OpenedByEmployeeName *string `json:"openedByEmployeeName,omitempty"`
		OpenedSum            string  `json:"openedSum,omitempty"`
		ClosedUnixTime       int     `json:"closedUnixTime,omitempty"`
		ClosedByEmployeeID   int     `json:"closedByEmployeeID,omitempty"`
		ClosedByEmployeeName *string `json:"closedByEmployeeName,omitempty"`
		ClosedSum            string  `json:"closedSum,omitempty"`
		BankedSum            string  `json:"bankedSum,omitempty"`
		Notes                string  `json:"notes,omitempty"`
		ReasonID             int     `json:"reasonID,omitempty"`
		ShiftType            string  `json:"shiftType,omitempty"`
		Employees            []struct {
			EmployeeID int `json:"employeeID,omitempty"`
		} `json:"employees,omitempty"`
	}

	GetCashInsResponse struct {
		Status  common2.Status `json:"status"`
		CashIns []CashIn       `json:"records"`
	}

	CashIn struct {
		TransactionID   int                    `json:"transactionID,omitempty"`
		Sum             float64                `json:"sum,omitempty"`
		CurrencyCode    string                 `json:"currencyCode,omitempty"`
		CurrencyRate    float64                `json:"currencyRate,omitempty"`
		WarehouseID     int                    `json:"warehouseID,omitempty"`
		WarehouseName   string                 `json:"warehouseName,omitempty"`
		PointOfSaleID   int                    `json:"pointOfSaleID,omitempty"`
		PointOfSaleName string                 `json:"pointOfSaleName,omitempty"`
		EmployeeID      int                    `json:"employeeID,omitempty"`
		EmployeeName    string                 `json:"employeeName,omitempty"`
		DateTime        string                 `json:"dateTime,omitempty"`
		Comment         string                 `json:"comment,omitempty"`
		Attributes      []common2.ObjAttribute `json:"attributes,omitempty"`
	}

	GetReasonCodesResponse struct {
		Status      common2.Status `json:"status"`
		ReasonCodes []ReasonCode   `json:"records"`
	}

	ReasonCode struct {
		ReasonID int    `json:"reasonID,omitempty"`
		Name     string `json:"name,omitempty"`
		Purpose  string `json:"purpose,omitempty"`
		Code     string `json:"code,omitempty"`
	}
)
