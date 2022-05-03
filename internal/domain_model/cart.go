package domain_model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type (
	Cart struct {
		Items []CartItem `json:"items" query:"items"`
	}

	CartItem struct {
		ProductID uuid.UUID       `json:"product_id" query:"product_id"`
		Price     decimal.Decimal `json:"price"      query:"price"`
		Quantity  decimal.Decimal `json:"quantity"   query:"quantity"`
	}
)

func (c *Cart) Add(cartItem *CartItem) {
	c.Items = append(c.Items, *cartItem)
}

func (c *Cart) RemoveKey(key int) {
	if key == len(c.Items)-1 {
		c.Items = c.Items[0:key]
		return
	}

	c.Items = append(c.Items[0:key], c.Items[key+1:]...)
}

func (c *Cart) Contains(productID uuid.UUID) *CartItem {
	for _, item := range c.Items {
		if item.ProductID == productID {
			return &item
		}
	}

	return nil
}

func (c *Cart) ContainsIndex(productID uuid.UUID) int {
	for key, item := range c.Items {
		if item.ProductID == productID {
			return key
		}
	}

	return -1
}

func (c *Cart) ItemsQuantity() int {
	return len(c.Items)
}

func (c *Cart) Total() decimal.Decimal {
	var total decimal.Decimal
	for _, item := range c.Items {
		total.Add(item.Total())
	}
	return total
}

func (ci *CartItem) ZeroQuantity() bool {
	return ci.Quantity == decimal.NewFromFloat(0)
}

func (ci *CartItem) Total() decimal.Decimal {
	return ci.Price.Mul(ci.Quantity).Round(2)
}
