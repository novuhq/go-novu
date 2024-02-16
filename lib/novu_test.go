package lib_test

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/novuhq/go-novu/lib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestError_Retry_With_Custom_Config(t *testing.T) {
	var (
		subscriberBulkPayload lib.SubscriberBulkPayload
		receivedBody          lib.SubscriberBulkPayload
		expectedRequest       lib.SubscriberBulkPayload
	)
	reqCount := 0
	var idempotencyHeader []string
	allElementsSame := func(arr []string) bool {
		if len(arr) == 0 {
			return true // An empty array is considered to have all elements the same.
		}
		firstElement := arr[0]
		for _, element := range arr {
			if element != firstElement {
				return false
			}
		}
		return true
	}
	subscriberService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if err := json.NewDecoder(req.Body).Decode(&receivedBody); err != nil {
			log.Printf("error in unmarshalling %+v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		reqCount++

		t.Run("Header must contain Idempotency-Key", func(t *testing.T) {
			idKey := req.Header.Get("Idempotency-Key")
			idempotencyHeader = append(idempotencyHeader, idKey)
			assert.NotNil(t, idKey)
		})
		t.Run("Header must contain ApiKey", func(t *testing.T) {
			authKey := req.Header.Get("Authorization")
			assert.True(t, strings.Contains(authKey, novuApiKey))
			assert.True(t, strings.HasPrefix(authKey, "ApiKey"))
		})

		t.Run("URL and request method is as expected", func(t *testing.T) {
			expectedURL := "/v1/subscribers/bulk"
			assert.Equal(t, http.MethodPost, req.Method)
			assert.Equal(t, expectedURL, req.RequestURI)
		})

		t.Run("Request is as expected", func(t *testing.T) {
			fileToStruct(filepath.Join("../testdata", "subscriber_bulk.json"), &expectedRequest)
			assert.Equal(t, expectedRequest, receivedBody)
		})

		var resp lib.SubscriberResponse
		fileToStruct(filepath.Join("../testdata", "subscriber_bulk_response.json"), &resp)

		w.WriteHeader(http.StatusInternalServerError)
		bb, _ := json.Marshal(resp)
		w.Write(bb)
	}))

	defer subscriberService.Close()

	ctx := context.Background()
	fileToStruct(filepath.Join("../testdata", "subscriber_bulk.json"), &subscriberBulkPayload)

	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(subscriberService.URL), RetryConfig: &lib.RetryConfigType{RetryMax: 5, InitialDelay: 0 * time.Second}})

	resp, err := c.SubscriberApi.BulkCreate(ctx, subscriberBulkPayload)
	require.NotNil(t, err)
	assert.NotNil(t, resp)

	//idempotency and retry tests
	assert.Equal(t, reqCount, 6)
	assert.Equal(t, len(idempotencyHeader), 6)
	assert.True(t, allElementsSame(idempotencyHeader))
}
func TestError_Retry_With_Default_Config(t *testing.T) {
	var (
		subscriberBulkPayload lib.SubscriberBulkPayload
		receivedBody          lib.SubscriberBulkPayload
		expectedRequest       lib.SubscriberBulkPayload
	)
	reqCount := 0
	var idempotencyHeader []string
	allElementsSame := func(arr []string) bool {
		if len(arr) == 0 {
			return true // An empty array is considered to have all elements the same.
		}
		firstElement := arr[0]
		for _, element := range arr {
			if element != firstElement {
				return false
			}
		}
		return true
	}
	subscriberService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if err := json.NewDecoder(req.Body).Decode(&receivedBody); err != nil {
			log.Printf("error in unmarshalling %+v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		reqCount++

		t.Run("Header must contain Idempotency-Key", func(t *testing.T) {
			idKey := req.Header.Get("Idempotency-Key")
			idempotencyHeader = append(idempotencyHeader, idKey)
			assert.NotNil(t, idKey)
		})
		t.Run("Header must contain ApiKey", func(t *testing.T) {
			authKey := req.Header.Get("Authorization")
			assert.True(t, strings.Contains(authKey, novuApiKey))
			assert.True(t, strings.HasPrefix(authKey, "ApiKey"))
		})

		t.Run("URL and request method is as expected", func(t *testing.T) {
			expectedURL := "/v1/subscribers/bulk"
			assert.Equal(t, http.MethodPost, req.Method)
			assert.Equal(t, expectedURL, req.RequestURI)
		})

		t.Run("Request is as expected", func(t *testing.T) {
			fileToStruct(filepath.Join("../testdata", "subscriber_bulk.json"), &expectedRequest)
			assert.Equal(t, expectedRequest, receivedBody)
		})

		var resp lib.SubscriberResponse
		fileToStruct(filepath.Join("../testdata", "subscriber_bulk_response.json"), &resp)

		w.WriteHeader(http.StatusInternalServerError)
		bb, _ := json.Marshal(resp)
		w.Write(bb)
	}))

	defer subscriberService.Close()

	ctx := context.Background()
	fileToStruct(filepath.Join("../testdata", "subscriber_bulk.json"), &subscriberBulkPayload)

	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(subscriberService.URL)})

	resp, err := c.SubscriberApi.BulkCreate(ctx, subscriberBulkPayload)
	require.NotNil(t, err)
	assert.NotNil(t, resp)

	//idempotency and retry tests
	assert.Equal(t, reqCount, 1)
	assert.True(t, allElementsSame(idempotencyHeader))
	assert.Equal(t, len(idempotencyHeader), 1)
}
