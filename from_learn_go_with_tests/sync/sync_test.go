package sync

import (
	"sync"
	"testing"
)

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, &counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCount) // Add 1000 times for waiting

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done() // Decrease 1 time from 1000 since it's done 1 time
			}()
		}
		wg.Wait()

		assertCounter(t, &counter, wantedCount)
	})
}
