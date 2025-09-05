package integration

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sawpanic/CGemini/api"
	"github.com/sawpanic/CGemini/config"
)

func TestKrakenAPI(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"error":[],"result":{"XXBTZUSD":{"a":["60000.00000","1","1.000"],"b":["59999.00000","1","1.000"],"c":["60000.00000","0.00010000"],"v":["1000.00000000","2000.00000000"],"p":["60000.00000","60000.00000"],"t":[100,200],"l":["59000.00000","59000.00000"],"h":["61000.00000","61000.00000"],"o":"60500.00000"}}}`))
	}))
	defer server.Close()

	cfg := &config.APIsConfig{}
	client := api.NewKrakenClient(cfg, api.NewCircuitBreaker("test", 0, 0, 0))

	// Override the base URL to point to the test server
	if k, ok := client.(*api.KrakenClientImpl); ok {
		k.SetBaseURL(server.URL)
	} else {
		t.Fatal("Failed to cast client to KrakenClientImpl")
	}

	_, err := client.GetTicker("XBTUSD")
	if err != nil {
		t.Fatalf("GetTicker failed: %v", err)
	}
}
