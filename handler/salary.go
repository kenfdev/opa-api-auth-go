package handler

import (
	"github.com/kenfdev/opa-api-auth-go/gateway"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// SalaryHandler is the handler responsible for salary resources
func SalaryHandler(gw gateway.SalaryGateway) echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		id := c.Param("id")
		logrus.WithFields(logrus.Fields{
			"id": id,
		}).Info("Processing salary request")

		salary := gw.FetchSalary(id)

		return c.String(http.StatusOK, strconv.Itoa(salary))
	}
}
