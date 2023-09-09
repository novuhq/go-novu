package lib_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/novuhq/go-novu/lib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInboundParserService_Get_Success(t *testing.T) {
	expectedResponse := &lib.InboundParserResponse{
		Data: lib.MxRecordConfiguredStatus{MxRecordConfigured: false},
	}

	inboundParserService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Write the expected response as JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(expectedResponse)
		assert.NoError(t, err)

		t.Run("Header must contain ApiKey", func(t *testing.T) {
			authKey := r.Header.Get("Authorization")
			assert.True(t, strings.Contains(authKey, novuApiKey))
			assert.True(t, strings.HasPrefix(authKey, "ApiKey"))
		})

		t.Run("URL and request method is as expected", func(t *testing.T) {
			expectedURL := "/v1/inbound-parse/mx/status"
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, expectedURL, r.RequestURI)
		})

	}))

	defer inboundParserService.Close()

	ctx := context.Background()
	i := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(inboundParserService.URL)})

	resp, err := i.InboundParserApi.Get(ctx)

	require.Nil(t, err)
	require.Equal(t, resp, expectedResponse)
}
