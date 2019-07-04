package bittrex

import (
	"time"

	"github.com/shopspring/decimal"
)

type Order struct {
	OrderUuid         string          `json:"OrderUuid"`
	Exchange          string          `json:"Exchange"`
	TimeStamp         jTime           `json:"TimeStamp"`
	OrderType         string          `json:"OrderType"`
	Limit             decimal.Decimal `json:"Limit"`
	Quantity          decimal.Decimal `json:"Quantity"`
	QuantityRemaining decimal.Decimal `json:"QuantityRemaining"`
	Commission        decimal.Decimal `json:"Commission"`
	Price             decimal.Decimal `json:"Price"`
	PricePerUnit      decimal.Decimal `json:"PricePerUnit"`
	Opened            string 		  `json:"Opened,omitempty"`
}

// For getorder
type Order2 struct {
	AccountId                  string
	OrderUuid                  string          `json:"OrderUuid"`
	Exchange                   string          `json:"Exchange"`
	Type                       string          `json:"OrderType"`
	Quantity                   decimal.Decimal `json:"Quantity"`
	QuantityRemaining          decimal.Decimal `json:"QuantityRemaining"`
	Limit                      decimal.Decimal `json:"Limit"`
	Reserved                   decimal.Decimal
	ReserveRemaining           decimal.Decimal
	CommissionReserved         decimal.Decimal
	CommissionReserveRemaining decimal.Decimal
	CommissionPaid             decimal.Decimal
	Price                      decimal.Decimal `json:"Price"`
	PricePerUnit               decimal.Decimal `json:"PricePerUnit"`
	Opened                     string
	Closed                     string
	IsOpen                     bool
	Sentinel                   string
	CancelInitiated            bool
	ImmediateOrCancel          bool
	IsConditional              bool
	Condition                  string
	ConditionTarget            decimal.Decimal
}

type NewOrderV3 struct {
	MarketSymbol  string          `json:"marketSymbol"`
	Direction     string          `json:"direction"`
	Type          string          `json:"type"` // LIMIT
	Quantity      decimal.Decimal `json:"quantity"`
	Limit         decimal.Decimal `json:"limit"`
	TimeInForce   string          `json:"timeInForce"` // GOOD_TIL_CANCELLED||IMMEDIATE_OR_CANCEL||FILL_OR_KILL||POST_ONLY_GOOD_TIL_CANCELLED||IMMEDIATE_OR_CANCEL||FILL_OR_KILL||POST_ONLY_GOOD_TIL_CANCELLED
	ClientOrderId *string         `json:"clientOrderId"`
}

type OrderV3 struct {
	ID            string          `json:"id"`
	MarketSymbol  string          `json:"marketSymbol"`
	Direction     string          `json:"direction"`
	Type          string          `json:"type"`
	Quantity      decimal.Decimal `json:"quantity"`
	Limit         decimal.Decimal `json:"limit"`
	Ceiling       decimal.Decimal `json:"ceiling"`
	TimeInForce   string          `json:"timeInForce"`
	ExpiresAt     time.Time       `json:"expiresAt"`
	ClientOrderId string          `json:"clientOrderId"`
	FillQuantity  decimal.Decimal `json:"fillQuantity"`
	Commission    decimal.Decimal `json:"commission"`
	Proceeds      decimal.Decimal `json:"proceeds"`
	Status        string          `json:"status"`
	CreatedAt     time.Time       `json:"createdAt"`
	UpdatedAt     time.Time       `json:"updatedAt"`
	ClosedAt      time.Time       `json:"closedAt"`
}
