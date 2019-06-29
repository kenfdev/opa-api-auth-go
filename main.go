package main

import (
	"github.com/kenfdev/opa-api-auth-go/gateway"
	"github.com/kenfdev/opa-api-auth-go/handler"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// inits
	salaryGateway := gateway.NewSalaryInMemoryGateway()

	// routing
	e.GET("/finance/salary/:id", handler.SalaryHandler(salaryGateway))

	e.Logger.Fatal(e.Start(":1323"))
}
