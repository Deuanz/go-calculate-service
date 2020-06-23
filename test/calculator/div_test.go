package test

import (
	"go-calculate-service/router"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestDiv(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		router := router.Setup()

		w := httptest.NewRecorder()

		body := []byte(`{
			"a":5,
			"b":2
		}`)
		req, _ := http.NewRequest(http.MethodPost, "/calculator.div", strings.NewReader(string(body)))

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, `{"result":"2.5"}`, w.Body.String())
	})

	t.Run("divide by zero", func(t *testing.T) {
		router := router.Setup()

		w := httptest.NewRecorder()

		body := []byte(`{
			"a":5,
			"b":0
		}`)
		req, _ := http.NewRequest(http.MethodPost, "/calculator.div", strings.NewReader(string(body)))

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, `{"error":"denominator can not be zero"}`, w.Body.String())
	})

}
