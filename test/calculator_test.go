package calculator

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go-calculate-service/router"

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

func TestSub(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		router := router.Setup()

		w := httptest.NewRecorder()

		body := []byte(`{
			"a":4,
			"b":3
		}`)
		req, _ := http.NewRequest(http.MethodPost, "/calculator.sub", strings.NewReader(string(body)))

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, `{"result":"1"}`, w.Body.String())
	})

}

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
