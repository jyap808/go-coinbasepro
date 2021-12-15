// Package Coinbase Pro is an implementation of the Coinbase Pro API in Golang.
package coinbasepro

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	client *Client
)

func setup() func() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = New(BaseURL(server.URL + "/"))

	return func() {
		server.Close()
	}
}

func fixture(path string) string {
	b, err := ioutil.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func TestClient_GetCurrencies(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/currencies", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("currencies.json"))
	})

	type fields struct {
		baseURL    string
		httpClient *http.Client
	}
	tests := []struct {
		name           string
		fields         fields
		wantCurrencies []Currency
		wantErr        bool
	}{
		{
			name: "Test",
			fields: fields{
				baseURL:    client.baseURL,
				httpClient: client.httpClient,
			},
			wantCurrencies: []Currency{
				{Id: "OXT", Name: "Orchid",
					Details: Details{
						NetworkConfirmations: json.Number("35")}},
				{Id: "IOTX", Name: "IoTeX (ERC-20)",
					Details: Details{
						NetworkConfirmations: json.Number("35")}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Client{
				baseURL:    tt.fields.baseURL,
				httpClient: tt.fields.httpClient,
			}
			gotCurrencies, err := b.GetCurrencies()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetCurrencies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCurrencies, tt.wantCurrencies) {
				t.Errorf("Client.GetCurrencies() = %v, want %v", gotCurrencies, tt.wantCurrencies)
			}
		})
	}
}

func TestClient_GetStats(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/products/stats", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("stats.json"))
	})

	type fields struct {
		baseURL    string
		httpClient *http.Client
	}
	tests := []struct {
		name      string
		fields    fields
		wantStats map[string]Stat
		wantErr   bool
	}{
		{
			name: "Test",
			fields: fields{
				baseURL:    client.baseURL,
				httpClient: client.httpClient,
			},
			wantStats: map[string]Stat{
				"BCH-BTC": {
					Stats24Hour{
						Volume: json.Number("3263.97247188"), Last: json.Number("0.00902")}}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Client{
				baseURL:    tt.fields.baseURL,
				httpClient: tt.fields.httpClient,
			}
			gotStats, err := b.GetStats()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetStats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStats, tt.wantStats) {
				t.Errorf("Client.GetStats() = %v, want %v", gotStats, tt.wantStats)
			}
		})
	}
}
