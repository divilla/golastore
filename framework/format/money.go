package format

import (
	"github.com/jackc/pgtype"
	"github.com/leekchan/accounting"
	"github.com/shopspring/decimal"
)

var ac = accounting.Accounting{Symbol: "$", Precision: 2}

func MoneyDecimal(val decimal.Decimal) string {
	return ac.FormatMoneyDecimal(val)
}

func MoneyNumeric(val pgtype.Numeric) string {
	if val.Status == pgtype.Null {
		return ""
	}

	var f float64
	err := val.AssignTo(&f)
	if err != nil {
		return ""
	}

	return ac.FormatMoneyFloat64(f)
}
