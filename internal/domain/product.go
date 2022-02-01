package domain

import (
	"fmt"
	"github.com/divilla/golastore/framework/format"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/shopspring/decimal"
)

type (
	ListProduct struct {
		Id          uuid.UUID
		Code        string
		Name        string
		Slug        string
		OldPrice    pgtype.Numeric
		Price       pgtype.Numeric
		Description string
	}
)

func (p *ListProduct) LinkToProduct() string {
	return fmt.Sprintf("/p/%s", p.Slug)
}

func (p *ListProduct) ImageURL() string {
	return fmt.Sprintf("https://yekupi.blob.core.windows.net/ekupihr/300Wx300H/%s_1.image", p.Code)
}

func (p *ListProduct) OldPriceFormat() string {
	return format.MoneyNumeric(p.OldPrice)
}

func (p *ListProduct) PriceFormat() string {
	return format.MoneyNumeric(p.Price)
}

func (p *ListProduct) Discount() string {
	if p.OldPrice.Status == pgtype.Null || p.Price.Status == pgtype.Null {
		return ""
	}

	var oldPrice, price float64
	if _ = p.OldPrice.AssignTo(&oldPrice); oldPrice == 0 {
		return ""
	}
	if _ = p.Price.AssignTo(&price); price == 0 {
		return ""
	}

	discount := decimal.NewFromFloat((oldPrice - price) / oldPrice * 100).Round(0)
	return discount.String() + "%"
}
