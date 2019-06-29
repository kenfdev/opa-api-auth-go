package handler

import (
	"github.com/kenfdev/opa-api-auth-go/gateway"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

// SalaryHandler is the handler responsible for salary resources
func SalaryHandler(gw gateway.SalaryGateway) echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		id := c.Param("id")
		salary := gw.FetchSalary(id)

		return c.String(http.StatusOK, strconv.Itoa(salary))
	}
}
