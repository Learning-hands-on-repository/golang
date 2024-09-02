package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

// Fetch implements Store.
func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func TestServer(t *testing.T) {

	t.Run("returns data from store", func(t *testing.T) {
		// Arrange
		data := "hello, world"
		store := &SpyStore{response: data}
		svr := Server(store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		// Act
		svr.ServeHTTP(response, request)

		// Assert
		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}
		if store.cancelled {
			t.Error("it should not have cancelled the store")
		}
	})

	t.Run("tell store to cancel work if request is cancelled", func(t *testing.T) {

		// Arrange
		mockData := "hello world"
		store := &SpyStore{response: mockData}
		svr := Server(store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		// NOTE: It's important that you derive your contexts so that cancellations are propagated throughout the call stack for a given request.
		// What we do is derive a new cancellingCtx from our request which returns us a cancel function. We then schedule that function to be called in 5 milliseconds by using time.AfterFunc.
		// Finally we use this new context in our request by calling request.WithContext.
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)
		response := httptest.NewRecorder()

		// Act
		svr.ServeHTTP(response, request)

		// Assert
		if !store.cancelled {
			t.Error("store was not told to cancel")
		}
	})
}
