package health_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/marcelofelixsalgado/financial-period-api/api/controllers/health"

	echo "github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/v1/health", nil)
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)

	handler := NewHealthHandler()

	if assert.NoError(t, handler.Health(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
