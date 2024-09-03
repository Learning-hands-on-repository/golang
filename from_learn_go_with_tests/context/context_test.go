package context

import (
	"context"
	"errors"
	"log"
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

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestServer(t *testing.T) {

	t.Run("returns data from store", func(t *testing.T) {
		// Arrange
		data := "hello, world"
		store := &SpyStore{response: data}
		svr := Server(store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder() // How to get 'response' from fake calling http

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
		//response := httptest.NewRecorder() // Need to not use this since httpTest.NewRecorder can not check 'reponse should not have been written'
		response := &SpyResponseWriter{}

		// Act
		svr.ServeHTTP(response, request)

		// Assert
		if response.written {
			t.Error("a response should not have been written since we run 'cancel' after pass 5 ms")
		}
	})
}
