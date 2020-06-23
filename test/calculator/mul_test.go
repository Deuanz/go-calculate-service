package test

import (
	"go-calculate-service/router"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestMul(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		router := router.Setup()

		w := httptest.NewRecorder()

		body := []byte(`{
			"a":2,
			"b":3
		}`)
		req, _ := http.NewRequest(http.MethodPost, "/calculator.mul", strings.NewReader(string(body)))

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, `{"result":"6"}`, w.Body.String())
	})

}
