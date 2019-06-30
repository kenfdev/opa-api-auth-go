package middleware

import (
	"github.com/kenfdev/opa-api-auth-go/gateway"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"
)

// EnforcePolicy is a middleware to check all the requests and ask if the requester has permissions
func EnforcePolicy(gateway gateway.PolicyGateway) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			logrus.Info("Enforcing policy with middleware")

			allow := gateway.Ask(c)

			if allow {
				logrus.Info("Action is allowed, continuing process")
				return next(c)
			} else {
				logrus.Info("Action was not allowed, cancelling process")
				return c.String(http.StatusForbidden, "Action not allowed")
			}
		}
	}
}
