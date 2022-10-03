package custom_source_header_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	plugin "github.com/bonsai-oss/custom-source-header"
)

func dummyHandler(w http.ResponseWriter, r *http.Request) {
	headers := r.Header
	json.NewEncoder(w).Encode(headers)
	return
}

func TestCustomSourceHeader_ServeHTTP(t *testing.T) {
	for _, tt := range []struct {
		name         string
		pluginConfig *plugin.Config
	}{
		{
			name: "test_hdr1",
			pluginConfig: &plugin.Config{
				HeaderName: "X-Forwarded-For",
			},
		},
		{
			name: "test_hdr2",
			pluginConfig: &plugin.Config{
				HeaderName: "X-Source-Address",
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			pluginHandler, pluginHandlerCreateError := plugin.New(context.Background(), http.HandlerFunc(dummyHandler), tt.pluginConfig, tt.name)
			if pluginHandlerCreateError != nil {
				t.Fatal(pluginHandlerCreateError)
			}

			svr := httptest.NewServer(
				pluginHandler,
			)
			defer svr.Close()

			req, err := http.NewRequest("GET", svr.URL, nil)
			if err != nil {
				t.Fatal(err)
			}
			rsp, _ := (&http.Client{}).Do(req)
			defer rsp.Body.Close()

			responseHeaderData := make(map[string][]string)
			json.NewDecoder(rsp.Body).Decode(&responseHeaderData)

			if _, ok := responseHeaderData[tt.pluginConfig.HeaderName]; !ok {
				t.Errorf("expected header %s to be set", tt.pluginConfig.HeaderName)
			}
		})
	}
}

func TestCreateConfig(t *testing.T) {
	config := plugin.CreateConfig()

	if fmt.Sprintf("%T", config) != "*custom_source_header.Config" {
		t.Errorf("expected config to be of type *custom_source_header.Config")
	}
}
