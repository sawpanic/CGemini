package circuitbreaker

import (
	"sync"
	"time"

	"github.com/sawpanic/CGemini/api"
	"github.com/sawpanic/CGemini/config"
)

// Manager manages a collection of circuit breakers for different API providers.
type Manager struct {
	mu       sync.RWMutex
	breakers map[string]*api.CircuitBreaker
	config   *config.CircuitsConfig
}

// NewManager creates a new CircuitBreaker Manager.
func NewManager(cfg *config.CircuitsConfig) *Manager {
	m := &Manager{
		breakers: make(map[string]*api.CircuitBreaker),
		config:   cfg,
	}
	m.initBreakers()
	return m
}

// initBreakers initializes the circuit breakers based on the configuration.
func (m *Manager) initBreakers() {
	for name, cfg := range m.config.CircuitBreakers {
		resetTimeout, _ := time.ParseDuration(cfg.ResetTimeout)
		probeInterval, _ := time.ParseDuration(cfg.ProbeInterval)
		m.breakers[name] = api.NewCircuitBreaker(name, resetTimeout, probeInterval, cfg.MaxFailures)
	}
}

// GetBreaker returns the circuit breaker for a given provider name.
// If a breaker for the provider does not exist, it creates a default one.
func (m *Manager) GetBreaker(name string) *api.CircuitBreaker {
	m.mu.RLock()
	breaker, ok := m.breakers[name]
	m.mu.RUnlock()

	if !ok {
		// If no specific configuration exists, create a default breaker.
		// This can be adjusted to return an error or use a more sophisticated default.
		breaker = api.NewCircuitBreaker(name, 30*time.Second, 5*time.Second, 5)
		m.mu.Lock()
		m.breakers[name] = breaker
		m.mu.Unlock()
	}

	return breaker
}
