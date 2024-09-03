package context

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
	Cancel()
}

// NOTE: this Server still failed test since it will cancel everytime request was made
// func Server(store Store) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, store.Fetch())
// 		store.Cancel()
// 	}
// }

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprint(w, "fake response") // this will use 'fake response' as result of httpCall
		data, _ := store.Fetch(r.Context())
		fmt.Fprint(w, data) // this will use 'fake response' as result of httpCall
	}
	// NOTE: context has a method Done() which returns a channel which gets sent a signal when the context is "done" or "cancelled".
	// We want to listen to that signal and call store.Cancel if we get it but we want to ignore it if our Store manages to Fetch before it.
}
