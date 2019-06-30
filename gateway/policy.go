package gateway

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kenfdev/opa-api-auth-go/entity"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"regexp"
)

type PolicyGateway interface {
	Ask(echo.Context) bool
}

type PolicyLocalGateway struct {
	getSalaryPathRegex *regexp.Regexp
}

func NewPolicyLocalGateway() *PolicyLocalGateway {
	return &PolicyLocalGateway{
		getSalaryPathRegex: regexp.MustCompile("GET /finance/salary/.*"),
	}
}

// Ask checks if an action is permitted as it inspects the request context
func (gw *PolicyLocalGateway) Ask(c echo.Context) bool {
	path := c.Request().Method + " " + c.Path()

	raw := c.Get("token").(*jwt.Token)
	claims := raw.Claims.(*entity.TokenClaims)

	var allow = false
	switch {
	case gw.getSalaryPathRegex.MatchString(path):
		allow = gw.checkGETSalary(c, claims)
	}

	return allow
}

func (gw *PolicyLocalGateway) checkGETSalary(c echo.Context, claims *entity.TokenClaims) bool {
	userID := c.Param("id")
	logrus.WithFields(logrus.Fields{
		"userID": userID,
		"claims": claims,
	}).Info("Checking GET salary policies")

	if yes := gw.checkIfOwner(userID, claims); yes {
		logrus.Info("Allowing because requester is the owner")
		return true
	}

	if yes := gw.checkIfSubordinate(userID, claims); yes {
		logrus.Info("Allowing because target is a subordinate of requester")
		return true
	}

	if yes := gw.checkIfHR(claims); yes {
		logrus.Info("Allowing because requester is a member of HR")
		return true
	}

	logrus.Info("Denying request")
	return false
}

func (*PolicyLocalGateway) checkIfOwner(userID string, claims *entity.TokenClaims) bool {
	return userID == claims.User
}

func (*PolicyLocalGateway) checkIfSubordinate(userID string, claims *entity.TokenClaims) bool {
	for _, subordinate := range claims.Subordinates {
		if subordinate == userID {
			return true
		}
	}
	return false
}

func (*PolicyLocalGateway) checkIfHR(claims *entity.TokenClaims) bool {
	return claims.BelongsToHR
}
