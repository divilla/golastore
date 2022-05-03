package cart

import (
	"github.com/divilla/golastore/internal/domain_model"
)

type (
	Service struct{}
)

func NewService() *Service {
	return &Service{}
}

func (s *Service) Set(cart *domain_model.Cart, cartItem *domain_model.CartItem) {
	key := cart.ContainsIndex(cartItem.ProductID)
	if key >= 0 && cartItem.ZeroQuantity() {
		cart.RemoveKey(key)
		return
	}
	if key >= 0 {
		cart.Items[key].Price = cartItem.Price
		cart.Items[key].Quantity = cartItem.Quantity
		return
	}

	cart.Add(cartItem)
}
