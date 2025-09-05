package unit

import (
	"fmt"
	"testing"
	"time"

	"github.com/sawpanic/CGemini/api"
)

func TestCircuitBreaker(t *testing.T) {
	cb := api.NewCircuitBreaker("test", 1*time.Second, 1*time.Second, 1)

	// The breaker should be closed initially
	if cb.GetState() != api.StateClosed {
		t.Fatalf("expected state to be Closed, got %v", cb.GetState())
	}

	// The first request should fail and trip the breaker
	err := cb.Execute(func() error {
		return fmt.Errorf("failed")
	})
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	// The breaker should now be open
	if cb.GetState() != api.StateOpen {
		t.Fatalf("expected state to be Open, got %v", cb.GetState())
	}

	// The next request should fail immediately
	err = cb.Execute(func() error {
		return nil
	})
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	// Wait for the reset timeout
	time.Sleep(2 * time.Second)

	// The breaker should now be half-open
	if cb.GetState() != api.StateHalfOpen {
		t.Fatalf("expected state to be HalfOpen, got %v", cb.GetState())
	}

	// The next request should succeed and close the breaker
	err = cb.Execute(func() error {
		return nil
	})
	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}

	// The breaker should now be closed
	if cb.GetState() != api.StateClosed {
		t.Fatalf("expected state to be Closed, got %v", cb.GetState())
	}
}
