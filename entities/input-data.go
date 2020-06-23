package entities

import "github.com/shopspring/decimal"

type Data struct {
	A decimal.Decimal `json:"a"`
	B decimal.Decimal `json:"b"`
}
