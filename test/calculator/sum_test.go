package test

import (
	"go-calculate-service/router"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestSum(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		router := router.Setup()

		w := httptest.NewRecorder()

		body := []byte(`{
			"a":1,
			"b":2
		}`)
		req, _ := http.NewRequest(http.MethodPost, "/calculator.sum", strings.NewReader(string(body)))

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, `{"result":"3"}`, w.Body.String())
	})

}
