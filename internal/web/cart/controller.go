package cart

import (
	"encoding/json"
	"net/http"

	"github.com/divilla/golastore/framework/middleware"
	"github.com/divilla/golastore/internal/domain_model"
	"github.com/labstack/echo/v4"
)

const (
	cartCookieName = "cart"
)

type (
	controller struct {
		service *Service
	}
)

func NewController(e *echo.Echo, service *Service) {
	c := &controller{
		service: service,
	}

	group := e.Group("/cart")
	group.GET("/set", middleware.UseCustomContext(c.set))
}

func (c *controller) set(ctx *middleware.CustomContext) error {
	var cart domain_model.Cart
	if cartCookie, err := ctx.Cookie(cartCookieName); err == nil {
		if err := json.Unmarshal([]byte(cartCookie.Value), &cart); err != nil {
			return err
		}
	}

	var cartItem domain_model.CartItem
	if err := ctx.Bind(&cartItem); err != nil {
		return err
	}

	c.service.Set(&cart, &cartItem)
	cartBytes, err := json.Marshal(&cart)
	if err != nil {
		return err
	}

	ctx.SetCookie(&http.Cookie{
		Name:     cartCookieName,
		Value:    string(cartBytes),
		MaxAge:   0,
		Secure:   true,
		HttpOnly: true,
	})

	return ctx.Redirect(http.StatusTemporaryRedirect, ctx.QueryParam("return_url"))
}
