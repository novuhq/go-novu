package lib_test

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/novuhq/go-novu/utils"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/novuhq/go-novu/lib"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

const (
	novuApiKey  = "test-API-key"
	novuEventId = "test-novu"
)

func fileToStruct(filepath string, s interface{}) io.Reader {
	bb, _ := os.ReadFile(filepath)
	json.Unmarshal(bb, s)
	return bytes.NewReader(bb)
}

func TestEventServiceTrigger_Success(t *testing.T) {
	var (
		receivedBody         lib.ITriggerPayloadOptions
		expectedTokenRequest lib.ITriggerPayloadOptions
		triggerPayload       lib.ITriggerPayloadOptions
	)

	eventService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
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
			expectedURL := "/v1/events/trigger"
			assert.Equal(t, http.MethodPost, req.Method)
			assert.Equal(t, expectedURL, req.RequestURI)
		})

		t.Run("Request is as expected", func(t *testing.T) {
			fileToStruct(filepath.Join("../testdata", "novu_send_trigger.json"), &expectedTokenRequest)
			assert.Equal(t, expectedTokenRequest, receivedBody)
		})

		var resp lib.EventResponse
		fileToStruct(filepath.Join("../testdata", "novu_send_trigger_response.json"), &resp)

		w.WriteHeader(http.StatusOK)
		bb, _ := json.Marshal(resp)
		w.Write(bb)
	}))

	defer eventService.Close()

	ctx := context.Background()
	fileToStruct(filepath.Join("../testdata", "novu_send_trigger.json"), &triggerPayload)

	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: utils.MustParseURL(eventService.URL)})
	_, err := c.EventApi.Trigger(ctx, novuEventId, triggerPayload)

	require.Nil(t, err)
}

func TestEventServiceTriggerForTopic_Success(t *testing.T) {
	var (
		receivedBody         lib.ITriggerPayloadOptions
		expectedTokenRequest lib.ITriggerPayloadOptions
		triggerPayload       lib.ITriggerPayloadOptions
	)

	eventService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
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
			expectedURL := "/v1/events/trigger"
			assert.Equal(t, http.MethodPost, req.Method)
			assert.Equal(t, expectedURL, req.RequestURI)
		})

		t.Run("Request is as expected", func(t *testing.T) {
			fileToStruct(filepath.Join("../testdata", "novu_send_trigger_topic_recipient.json"), &expectedTokenRequest)
			assert.Equal(t, expectedTokenRequest, receivedBody)
		})

		var resp lib.EventResponse
		fileToStruct(filepath.Join("../testdata", "novu_send_trigger_response.json"), &resp)

		w.WriteHeader(http.StatusOK)
		bb, _ := json.Marshal(resp)
		w.Write(bb)
	}))

	defer eventService.Close()

	ctx := context.Background()
	fileToStruct(filepath.Join("../testdata", "novu_send_trigger_topic_recipient.json"), &triggerPayload)

	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: utils.MustParseURL(eventService.URL)})
	_, err := c.EventApi.Trigger(ctx, novuEventId, triggerPayload)

	require.Nil(t, err)
}
