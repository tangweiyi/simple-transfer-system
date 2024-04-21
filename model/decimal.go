package model

import (
	"encoding/json"

	"github.com/shopspring/decimal"
)

type Decimal5 struct {
	decimal.Decimal
}

func (d Decimal5) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.StringFixed(5))
}

func (d *Decimal5) UnmarshalJSON(decimalBytes []byte) error {
	err := d.Decimal.UnmarshalJSON(decimalBytes)
	if err != nil {
		return err
	}
	d.Decimal = d.Decimal.Truncate(5)
	return nil
}

func (d Decimal5) String() string {
	return d.StringFixed(5)
}
