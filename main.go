package main

import (
	"github.com/kenfdev/opa-api-auth-go/entity"
	"github.com/kenfdev/opa-api-auth-go/gateway"
	"github.com/kenfdev/opa-api-auth-go/handler"
	localmiddleware "github.com/kenfdev/opa-api-auth-go/middleware"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// inits
	policyGateway := gateway.NewPolicyLocalGateway()
	salaryGateway := gateway.NewSalaryInMemoryGateway()

	// middleware
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret"),
		ContextKey: "token",
		Claims:     &entity.TokenClaims{},
	}))
	e.Use(localmiddleware.EnforcePolicy(policyGateway))

	// routing
	e.GET("/finance/salary/:id", handler.SalaryHandler(salaryGateway))

	e.Logger.Fatal(e.Start(":1323"))
}
