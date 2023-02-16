package lib_test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/novuhq/go-novu/lib"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func assureRequestHeaders(t *testing.T, req *http.Request, expectedURL string, expectedMethod string) {
	t.Run("Header must contain ApiKey", func(t *testing.T) {
		authKey := req.Header.Get("Authorization")
		assert.True(t, strings.Contains(authKey, novuApiKey))
		assert.True(t, strings.HasPrefix(authKey, "ApiKey"))
	})

	t.Run("URL and request method is as expected", func(t *testing.T) {
		assert.Equal(t, expectedMethod, req.Method)
		assert.Equal(t, expectedURL, req.RequestURI)
	})
}

type TestServerOptions[T any, K interface{}] struct {
	expectedURLPath    string
	expectedSentMethod string
	expectedSentBody   T

	responseStatusCode int
	responseBody       K
}

func createTestServer[T any, K interface{}](t *testing.T, options TestServerOptions[T, K]) *httptest.Server {
	var receivedBody T

	eventService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assureRequestHeaders(t, req, options.expectedURLPath, options.expectedSentMethod)

		if req.Body == http.NoBody {
			assert.Empty(t, options.expectedSentBody)
		} else {
			if err := json.NewDecoder(req.Body).Decode(&receivedBody); err != nil {
				log.Printf("error in unmarshalling %+v", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			t.Run("Request is as expected", func(t *testing.T) {
				assert.Equal(t, options.expectedSentBody, receivedBody)
			})
		}

		w.WriteHeader(options.responseStatusCode)

		bb, _ := json.Marshal(options.responseBody)
		w.Write(bb)
	}))
	t.Cleanup(func() {
		eventService.Close()
	})

	return eventService
}

func TestCreateTopic_Success(t *testing.T) {
	topicName := "topic"
	topicKey := "topicKey"
	httpServer := createTestServer(t, TestServerOptions[lib.CreateTopicRequest, map[string]string]{
		expectedURLPath:    "/v1/topics",
		expectedSentMethod: http.MethodPost,
		expectedSentBody: lib.CreateTopicRequest{
			Name: topicName,
			Key:  topicKey,
		},
		responseStatusCode: http.StatusCreated,
		responseBody:       map[string]string{},
	})

	ctx := context.Background()
	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: httpServer.URL})
	err := c.TopicsApi.Create(ctx, "topicKey", "topic")

	require.NoError(t, err)
}

func TestAddSubscription_Success(t *testing.T) {
	subs := []string{"subId"}
	key := "topicKey"

	httpServer := createTestServer(t, TestServerOptions[lib.SubscribersTopicRequest, map[string]string]{
		expectedURLPath:    fmt.Sprintf("/v1/topics/%s/subscribers", key),
		expectedSentMethod: http.MethodPost,
		expectedSentBody: lib.SubscribersTopicRequest{
			Subscribers: subs,
		},
		responseStatusCode: http.StatusNoContent,
		responseBody:       map[string]string{},
	})

	ctx := context.Background()
	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: httpServer.URL})
	err := c.TopicsApi.AddSubscribers(ctx, key, subs)

	require.NoError(t, err)
}

func TestAddSubscriptionRemoval_Success(t *testing.T) {
	subs := []string{"subId"}
	key := "topicKey"

	httpServer := createTestServer(t, TestServerOptions[lib.SubscribersTopicRequest, map[string]string]{
		expectedURLPath:    fmt.Sprintf("/v1/topics/%s/subscribers/removal", key),
		expectedSentMethod: http.MethodPost,
		expectedSentBody: lib.SubscribersTopicRequest{
			Subscribers: subs,
		},
		responseStatusCode: http.StatusNoContent,
		responseBody:       map[string]string{},
	})

	ctx := context.Background()
	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: httpServer.URL})
	err := c.TopicsApi.RemoveSubscribers(ctx, key, subs)

	require.NoError(t, err)
}

func TestGetTopic_Success(t *testing.T) {
	key := "topicKey"
	body := map[string]string{}
	var expectedResponse *lib.GetTopicResponse = &lib.GetTopicResponse{
		Id:             "id",
		OrganizationId: "orgId",
		EnvironmentId:  "envId",
		Key:            "topicKey",
		Name:           "topicName",
		Subscribers:    []string{"sibId"},
	}

	httpServer := createTestServer(t, TestServerOptions[map[string]string, *lib.GetTopicResponse]{
		expectedURLPath:    fmt.Sprintf("/v1/topics/%s", key),
		expectedSentMethod: http.MethodGet,
		expectedSentBody:   body,
		responseStatusCode: http.StatusOK,
		responseBody:       expectedResponse,
	})

	ctx := context.Background()
	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: httpServer.URL})
	resp, err := c.TopicsApi.Get(ctx, key)

	require.NoError(t, err)
	require.Equal(t, resp, expectedResponse)
}

func TestListTopics_Success(t *testing.T) {
	body := map[string]string{}
	var expectedResponse *lib.ListTopicsResponse = &lib.ListTopicsResponse{
		Page:       0,
		PageSize:   20,
		TotalCount: 1,
		Data: []lib.GetTopicResponse{{
			Id:             "id",
			OrganizationId: "orgId",
			EnvironmentId:  "envId",
			Key:            "topicKey",
			Name:           "topicName",
			Subscribers:    []string{"sibId"},
		}},
	}

	httpServer := createTestServer(t, TestServerOptions[map[string]string, *lib.ListTopicsResponse]{
		expectedURLPath:    "/v1/topics",
		expectedSentMethod: http.MethodGet,
		expectedSentBody:   body,
		responseStatusCode: http.StatusOK,
		responseBody:       expectedResponse,
	})

	ctx := context.Background()
	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: httpServer.URL})
	resp, err := c.TopicsApi.List(ctx, nil)

	require.NoError(t, err)
	require.Equal(t, resp, expectedResponse)
}

func TestRenameTopic_Success(t *testing.T) {
	topicKey := "topicKey"
	newName := "topicName"
	body := lib.RenameTopicRequest{
		Name: newName,
	}
	var expectedResponse *lib.GetTopicResponse = &lib.GetTopicResponse{
		Id:             "id",
		OrganizationId: "orgId",
		EnvironmentId:  "envId",
		Name:           newName,
		Key:            topicKey,
		Subscribers:    []string{"sibId"},
	}

	httpServer := createTestServer(t, TestServerOptions[lib.RenameTopicRequest, *lib.GetTopicResponse]{
		expectedURLPath:    fmt.Sprintf("/v1/topics/%s", topicKey),
		expectedSentMethod: http.MethodPatch,
		expectedSentBody:   body,
		responseStatusCode: http.StatusOK,
		responseBody:       expectedResponse,
	})

	ctx := context.Background()
	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: httpServer.URL})
	resp, err := c.TopicsApi.Rename(ctx, topicKey, newName)

	require.NoError(t, err)
	require.Equal(t, resp, expectedResponse)
}
