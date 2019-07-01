package main

import (
	"github.com/kenfdev/opa-api-auth-go/entity"
	"github.com/kenfdev/opa-api-auth-go/gateway"
	"github.com/kenfdev/opa-api-auth-go/handler"
	localmiddleware "github.com/kenfdev/opa-api-auth-go/middleware"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

func main() {
	e := echo.New()

	// inits
	endpoint := os.Getenv("PDP_ENDPOINT")
	policyGateway := gateway.NewPolicyOpaGateway(endpoint)
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
