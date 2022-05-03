package domain_model

import (
	"fmt"
	"github.com/divilla/golastore/framework/format"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/shopspring/decimal"
)

type (
	ProductListItem struct {
		Id          uuid.UUID
		Code        string
		Name        string
		Slug        string
		OldPrice    pgtype.Numeric
		Price       pgtype.Numeric
		Description string
	}
)

func (p *ProductListItem) LinkToProduct() string {
	return fmt.Sprintf("/p/%s", p.Slug)
}

func (p *ProductListItem) ImageURL() string {
	return fmt.Sprintf("https://yekupi.blob.core.windows.net/ekupihr/300Wx300H/%s_1.image", p.Code)
}

func (p *ProductListItem) OldPriceFormat() string {
	return format.MoneyNumeric(p.OldPrice)
}

func (p *ProductListItem) PriceFormat() string {
	return format.MoneyNumeric(p.Price)
}

func (p *ProductListItem) Discount() string {
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
