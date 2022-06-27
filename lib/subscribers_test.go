package lib_test

import (
	"context"
	"encoding/json"
	"github.com/novuhq/go-novu/lib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"
)

const subscriberID = "62b51a44da1af31d109f5da7"

func TestSubscriberService_Identify_Success(t *testing.T) {
	var (
		subscriberPayload lib.SubscriberPayload
		receivedBody      lib.SubscriberPayload
		expectedRequest   lib.SubscriberPayload
		expectedResponse  lib.SubscriberResponse
	)

	subscriberService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if err := json.NewDecoder(req.Body).Decode(&receivedBody); err != nil {
			log.Printf("error in unmarshalling %+v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		t.Run("Header must contain ApiKey", func(t *testing.T) {
			authKey := req.Header.Get("Authorization")
			assert.True(t, strings.Contains(authKey, novuApiKey))
			assert.True(t, strings.HasPrefix(authKey, "ApiKey"))
		})

		t.Run("URL and request method is as expected", func(t *testing.T) {
			expectedURL := "/v1/subscribers"
			assert.Equal(t, http.MethodPost, req.Method)
			assert.Equal(t, expectedURL, req.RequestURI)
		})

		t.Run("Request is as expected", func(t *testing.T) {
			fileToStruct(filepath.Join("../testdata", "identify_subscriber.json"), &expectedRequest)
			assert.Equal(t, expectedRequest, receivedBody)
		})

		var resp lib.SubscriberResponse
		fileToStruct(filepath.Join("../testdata", "subscriber_response.json"), &resp)

		w.WriteHeader(http.StatusOK)
		bb, _ := json.Marshal(resp)
		w.Write(bb)
	}))

	defer subscriberService.Close()

	ctx := context.Background()
	fileToStruct(filepath.Join("../testdata", "identify_subscriber.json"), &subscriberPayload)

	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: subscriberService.URL})

	resp, err := c.SubscriberApi.Identify(ctx, subscriberID, subscriberPayload)
	require.Nil(t, err)
	assert.NotNil(t, resp)

	t.Run("Response is as expected", func(t *testing.T) {
		fileToStruct(filepath.Join("../testdata", "subscriber_response.json"), &expectedResponse)
		assert.Equal(t, expectedResponse, resp)
	})
}

func TestSubscriberService_Update_Success(t *testing.T) {
	var (
		updateSubscriber lib.SubscriberPayload
		receivedBody     lib.SubscriberPayload
		expectedRequest  lib.SubscriberPayload
		expectedResponse lib.SubscriberResponse
	)

	subscriberService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if err := json.NewDecoder(req.Body).Decode(&receivedBody); err != nil {
			log.Printf("error in unmarshalling %+v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		t.Run("Header must contain ApiKey", func(t *testing.T) {
			authKey := req.Header.Get("Authorization")
			assert.True(t, strings.Contains(authKey, novuApiKey))
			assert.True(t, strings.HasPrefix(authKey, "ApiKey"))
		})

		t.Run("URL and request method is as expected", func(t *testing.T) {
			expectedURL := "/v1/subscribers/" + subscriberID
			assert.Equal(t, http.MethodPut, req.Method)
			assert.Equal(t, expectedURL, req.RequestURI)
		})

		t.Run("Request is as expected", func(t *testing.T) {
			fileToStruct(filepath.Join("../testdata", "update_subscriber.json"), &expectedRequest)
			assert.Equal(t, expectedRequest, receivedBody)
		})

		var resp lib.SubscriberResponse
		fileToStruct(filepath.Join("../testdata", "subscriber_response.json"), &resp)

		w.WriteHeader(http.StatusOK)
		bb, _ := json.Marshal(resp)
		w.Write(bb)
	}))

	ctx := context.Background()
	fileToStruct(filepath.Join("../testdata", "update_subscriber.json"), &updateSubscriber)

	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: subscriberService.URL})

	resp, err := c.SubscriberApi.Update(ctx, subscriberID, updateSubscriber)
	require.Nil(t, err)
	assert.NotNil(t, resp)

	t.Run("Response is as expected", func(t *testing.T) {
		fileToStruct(filepath.Join("../testdata", "subscriber_response.json"), &expectedResponse)
		assert.Equal(t, expectedResponse, resp)
	})
}

func TestSubscriberService_Delete_Success(t *testing.T) {
	var expectedResponse lib.SubscriberResponse

	ctx := context.Background()

	subscriberService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		t.Run("Header must contain ApiKey", func(t *testing.T) {
			authKey := req.Header.Get("Authorization")
			assert.True(t, strings.Contains(authKey, novuApiKey))
			assert.True(t, strings.HasPrefix(authKey, "ApiKey"))
		})

		t.Run("URL and request method is as expected", func(t *testing.T) {
			expectedURL := "/v1/subscribers/" + subscriberID
			assert.Equal(t, http.MethodDelete, req.Method)
			assert.Equal(t, expectedURL, req.RequestURI)
		})

		var resp lib.SubscriberResponse
		fileToStruct(filepath.Join("../testdata", "subscriber_response.json"), &resp)

		w.WriteHeader(http.StatusOK)
		bb, _ := json.Marshal(resp)
		w.Write(bb)
	}))

	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: subscriberService.URL})

	resp, err := c.SubscriberApi.Delete(ctx, subscriberID)
	require.Nil(t, err)
	assert.NotNil(t, resp)

	t.Run("Response is as expected", func(t *testing.T) {
		fileToStruct(filepath.Join("../testdata", "subscriber_response.json"), &expectedResponse)
		assert.Equal(t, expectedResponse, resp)
	})
}
