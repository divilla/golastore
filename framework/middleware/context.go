package middleware

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

const (
	LoggerFieldsKey = "loggerFields"
)

type (
	CustomContext struct {
		echo.Context
	}

	CustomContextFunc func(ctx *CustomContext) error

	ContextIdentity struct {
		ID    string
		Email string
		Scope string
		Token string
	}
)

func NewCustomContext(ctx echo.Context) *CustomContext {
	return &CustomContext{
		Context: ctx,
	}
}

func CustomContextMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := NewCustomContext(c)
			return next(cc)
		}
	}
}

func UseCustomContext(ccf CustomContextFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ccf(NewCustomContext(ctx))
	}
}

func (cc *CustomContext) LogString(key, val string) {
	l := cc.Get(LoggerFieldsKey).(*zapLogger)
	l.fields = append(l.fields, zap.String(key, val))
}

func (cc *CustomContext) LogError(key string, val error) {
	l := cc.Get(LoggerFieldsKey).(*zapLogger)
	l.fields = append(l.fields, zap.NamedError(key, val))
}

func (cc *CustomContext) LogAny(key string, val interface{}) {
	l := cc.Get(LoggerFieldsKey).(*zapLogger)
	l.fields = append(l.fields, zap.Any(key, val))
}
