package gateway

import (
	"bytes"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
)

type PolicyGateway interface {
	Ask(echo.Context) bool
}

type opaRequest struct {
	// Input wraps the OPA request (https://www.openpolicyagent.org/docs/latest/rest-api/#get-a-document-with-input)
	Input *opaInput `json:"input"`
}

type opaInput struct {
	// The token of the requester
	Token string `json:"token"`
	// The path to which the request was made split to an array
	Path []string `json:"path"`
	// The HTTP Method
	Method string `json:"method"`
}

type opaResponse struct {
	Result bool `json:"result"`
}

// PolicyOpaGateway makes policy decision requests to OPA
type PolicyOpaGateway struct {
	endpoint string
}

func NewPolicyOpaGateway(endpoint string) *PolicyOpaGateway {
	return &PolicyOpaGateway{
		endpoint: endpoint,
	}
}

// Ask requests to OPA with required inputs and returns the decision made by OPA
func (gw *PolicyOpaGateway) Ask(c echo.Context) bool {
	token := c.Get("token").(*jwt.Token)
	// After splitting, the first element isn't necessary
	// "/finance/salary/alice" -> ["", "finance", "salary", "alice"]
	paths := strings.Split(c.Request().RequestURI, "/")[1:]
	method := c.Request().Method

	// create input to send to OPA
	input := &opaInput{
		Token:  token.Raw,
		Path:   paths,
		Method: method,
	}
	opaRequest := &opaRequest{
		Input: input,
	}

	logrus.WithFields(logrus.Fields{
		"token":  input.Token,
		"path":   input.Path,
		"method": input.Method,
	}).Info("Requesting PDP for decision")

	requestBody, err := json.Marshal(opaRequest)
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("Marshalling request body failed")
		return false
	}

	// request OPA
	resp, err := http.Post(gw.endpoint, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("PDP Request failed")
		return false
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("Reading body of response failed")
		return false
	}

	var opaResponse opaResponse
	err = json.Unmarshal(body, &opaResponse)
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("Unmarshalling response body failed")
		return false
	}

	logrus.WithFields(logrus.Fields{
		"result": opaResponse.Result,
	}).Info("Decision")

	return opaResponse.Result
}
