package middleware

import (
	"fmt"
	"github.com/divilla/golastore/pkg/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type (
	zapLogger struct {
		start  time.Time
		fields []zapcore.Field
	}
)

func ZapLoggerMiddleware(log *logger.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now().UTC()
			cc := NewCustomContext(c)
			zl := &zapLogger{
				start: time.Now().UTC(),
			}
			cc.Set(LoggerFieldsKey, zl)

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			req := cc.Request()
			res := cc.Response()

			fields := []zapcore.Field{
				zap.String("remote_ip", c.RealIP()),
				zap.String("latency", time.Since(start).String()),
				zap.String("host", req.Host),
				zap.String("request", fmt.Sprintf("%s %s", req.Method, req.RequestURI)),
				zap.Int("status", res.Status),
				zap.Int64("size", res.Size),
				zap.String("user_agent", req.UserAgent()),
			}

			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
				fields = append(fields, zap.String("request_id", id))
			}

			ctxFields := cc.Get(LoggerFieldsKey).(*zapLogger)
			for _, field := range ctxFields.fields {
				fields = append(fields, field)
			}

			n := res.Status
			if cc.Echo().Debug {
				switch {
				case n >= 500:
					log.Sugar().With("error", err, zap.Stack("stack")).Desugar().Error("Server error", fields...)
				case n >= 400:
					log.Zap().With(zap.Error(err)).Warn("Client error", fields...)
				case n >= 300:
					log.Zap().Info("Redirection", fields...)
				default:
					log.Zap().Info("Success", fields...)
				}
			} else {
				switch {
				case n >= 500:
					log.Zap().With(zap.Error(err)).Error("Server error", fields...)
				case n >= 400:
					log.Zap().With(zap.Error(err)).Warn("Client error", fields...)
				case n >= 300:
					log.Zap().Info("Redirection", fields...)
				default:
					log.Zap().Info("Success", fields...)
				}
			}

			return nil
		}
	}
}
