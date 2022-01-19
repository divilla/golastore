package middleware

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	SessionCookieName    = "__Host-SID"
	DevSessionCookieName = "SID"
	IdentityKey          = "contextIdentity"
	LoggerFieldsKey      = "loggerFields"
)

type (
	CustomContext struct {
		echo.Context
	}

	CustomContextFunc func(ctx *CustomContext) error

	contextIdentity struct {
		id    string
		email string
		scope string
		token string
	}

	zapLogger struct {
		fields []zapcore.Field
	}
)

func NewCustomContext(ctx echo.Context) *CustomContext {
	return &CustomContext{
		Context: ctx,
	}
}

func SetCustomContext() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := NewCustomContext(c)
			cc.Set(LoggerFieldsKey, &zapLogger{})
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

func (cc *CustomContext) sessionCookieName() string {
	if cc.Echo().Debug {
		return DevSessionCookieName
	}
	return SessionCookieName
}
