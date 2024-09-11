package web

import (
	"github.com/stretchr/testify/assert"
	"go-sync/application"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWebHandler(t *testing.T) {
	store := application.NewKeyValueStore()
	h := NewHandlers(store)
	t.Run("should return 400 for empty body", func(t *testing.T) {
		req, _ := http.NewRequest("PUT", "/key/test", strings.NewReader(""))
		w := httptest.NewRecorder()

		h.PutKey(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "Invalid or missing request body\n", w.Body.String())
	})

	t.Run("should return 400 for malformed body", func(t *testing.T) {
		req, _ := http.NewRequest("PUT", "/key/test", nil)
		w := httptest.NewRecorder()

		h.PutKey(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "Invalid or missing request body\n", w.Body.String())
	})

	t.Run("should return 404 for missing key", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/key/test", nil)
		w := httptest.NewRecorder()

		h.GetKey(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, "Key not found\n", w.Body.String())
	})

	t.Run("should add a key", func(t *testing.T) {
		req, _ := http.NewRequest("PUT", "/key/test", strings.NewReader("value"))
		w := httptest.NewRecorder()

		h.PutKey(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "", w.Body.String())
	})

	t.Run("should return value for a valid key", func(t *testing.T) {
		req, _ := http.NewRequest("PUT", "/key/test", strings.NewReader("value"))
		w := httptest.NewRecorder()

		h.PutKey(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "", w.Body.String())

		req, _ = http.NewRequest("GET", "/key/test", nil)
		w = httptest.NewRecorder()

		h.GetKey(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "value", w.Body.String())
	})

	t.Run("should delete a key", func(t *testing.T) {
		req, _ := http.NewRequest("PUT", "/key/test", strings.NewReader("value"))
		w := httptest.NewRecorder()

		h.PutKey(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "", w.Body.String())

		req, _ = http.NewRequest("DELETE", "/key/test", nil)
		w = httptest.NewRecorder()

		h.DeleteKey(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "", w.Body.String())
	})

	t.Run("should update a key", func(t *testing.T) {
		req, _ := http.NewRequest("PUT", "/key/test", strings.NewReader("value"))
		w := httptest.NewRecorder()

		h.PutKey(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "", w.Body.String())

		req, _ = http.NewRequest("PUT", "/key/test", strings.NewReader("new-value"))
		w = httptest.NewRecorder()

		h.PutKey(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "", w.Body.String())

		req, _ = http.NewRequest("GET", "/key/test", nil)
		w = httptest.NewRecorder()

		h.GetKey(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "new-value", w.Body.String())
	})

	t.Run("should list all keys", func(t *testing.T) {
		req, _ := http.NewRequest("PUT", "/key/test1", strings.NewReader("value1"))
		w := httptest.NewRecorder()

		h.PutKey(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "", w.Body.String())

		req, _ = http.NewRequest("PUT", "/key/test2", strings.NewReader("value2"))
		w = httptest.NewRecorder()

		h.PutKey(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "", w.Body.String())

		req, _ = http.NewRequest("PUT", "/key/test3", strings.NewReader("value3"))
		w = httptest.NewRecorder()

		h.PutKey(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "", w.Body.String())

		req, _ = http.NewRequest("GET", "/", nil)
		w = httptest.NewRecorder()

		h.ListKeys(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "test1")
		assert.Contains(t, w.Body.String(), "test2")
		assert.Contains(t, w.Body.String(), "test3")
	})
}
