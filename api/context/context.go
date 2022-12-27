package context

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type CustomContext struct {
	echo.Context
	StartTime time.Time
	Elapsed   map[string]int64
}

// RegisterElapsedTime generates response status about dependencies
func (cc *CustomContext) RegisterElapsedTime(app string, start time.Time) {
	cc.Elapsed[app] += time.Since(start).Nanoseconds() / int64(time.Millisecond)
}

// MessageIDFromContext get trace id or generate a new
func MessageIDFromContext(c echo.Context) string {

	if value := c.Param("id"); value != "" {
		return value
	}
	return uuid.NewV4().String()
}

// ContextRequestHTTP creates a new request object
func ContextRequestHTTP(c echo.Context) map[string]interface{} {
	var (
		bodyBytes []byte
		token     *jwt.Token
		request   = make(map[string]interface{})
	)

	req := c.Request()

	tokenString := req.Header.Get("Authorization")
	if tokenString != "" {
		token, _ = jwt.ParseWithClaims(strings.Replace(tokenString, "Bearer ", "", -1), jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(nil), nil
		})

		claims := token.Claims.(jwt.MapClaims)
		if len(claims) != 0 {
			request["client"] = token.Claims
		} else {
			request["client"] = tokenString
		}
	}

	if req.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(req.Body)
		if len(bodyBytes) > 0 {
			body := strings.Join(strings.Fields(string(bodyBytes)), "")
			request["body"] = body
			c.Set("request_body", body)
		}
		req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	header, _ := json.Marshal(req.Header)
	request["header"] = string(header)
	request["path"] = c.Path()
	request["query_params"] = req.URL.Query()
	request["method"] = req.Method
	request["uri"] = req.RequestURI

	return request
}

func ContextResponseHTTP(c echo.Context) map[string]interface{} {
	var (
		response = make(map[string]interface{})
		// latencySeconds = make(map[string]interface{})
	)

	if c.Get("response_body") != nil {
		bodyBytes, _ := json.Marshal(c.Get("response_body"))
		body := string(bodyBytes)
		body = strings.Replace(body, "\\\"", "", -1)
		response["body"] = body
		c.Set("resp_body", body)
	}

	// if c.Get("dependencies_latency") != nil {
	// 	elapsed := c.Get("dependencies_latency")
	// 	latencyMs := elapsed.(map[string]int64)
	// 	for service, value := range latencyMs {
	// 		seconds := util.SetPrecisionFloat(float64(value) / 1000)
	// 		latencySeconds[service] = seconds
	// 	}
	// 	response["latency_seconds"] = latencySeconds
	// }
	response["status_code"] = c.Response().Status

	return response
}
