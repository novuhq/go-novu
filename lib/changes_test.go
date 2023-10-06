package lib_test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/novuhq/go-novu/lib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var bulkApplyPayload = `{
  "changeIds": [
    "string"
  ]
}`

var applyResponse = `{
	"data": [
		{
			"_id": "string",
			"_creatorId": "string",
			"_environmentId": "string",
			"_organizationId": "string",
			"_entityId": "string",
			"enabled": true,
			"type": "Feed",
			"change": {},
			"createdAt": "string",
			"_parentId": "string"
		}
	]
}`

var getResponse = `{
  "totalCount": 0,
  "data": [
    {
      "_id": "string",
      "_creatorId": "string",
      "_environmentId": "string",
      "_organizationId": "string",
      "_entityId": "string",
      "enabled": true,
      "type": "Feed",
      "change": {},
      "createdAt": "string",
      "_parentId": "string"
    }
  ],
  "pageSize": 0,
  "page": 0
}
`

var getCountResponse = `{
  "data": 0
}`

func payloadStringToStruct(str string, s interface{}) error {
	bb := []byte(str)
	err := json.Unmarshal(bb, s)
	if err != nil {
		return err
	}
	return nil
}

func TestChangesService_GetCount_Success(t *testing.T) {
	var (
		expectedResponse lib.ChangesCountResponse
	)

	ChangesService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		t.Run("Header must contain ApiKey", func(t *testing.T) {
			authKey := req.Header.Get("Authorization")
			assert.True(t, strings.Contains(authKey, novuApiKey))
			assert.True(t, strings.HasPrefix(authKey, "ApiKey"))
		})

		t.Run("URL and request method is as expected", func(t *testing.T) {
			expectedURL := fmt.Sprintf("/v1/changes/count")
			assert.Equal(t, http.MethodGet, req.Method)
			assert.Equal(t, expectedURL, req.RequestURI)
		})

		var resp lib.ChangesCountResponse
		err := payloadStringToStruct(getCountResponse, &resp)
		require.Nil(t, err)

		w.WriteHeader(http.StatusOK)
		bb, _ := json.Marshal(resp)
		w.Write(bb)
	}))

	defer ChangesService.Close()

	ctx := context.Background()

	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(ChangesService.URL)})
	resp, err := c.ChangesApi.GetChangesCount(ctx)
	require.Nil(t, err)
	assert.NotNil(t, resp)

	t.Run("Response is as expected", func(t *testing.T) {
		err := payloadStringToStruct(getCountResponse, &expectedResponse)
		require.Nil(t, err)
		assert.Equal(t, expectedResponse, resp)
	})
}

func TestChangesService_Get_Success(t *testing.T) {
	var (
		expectedResponse lib.ChangesGetResponse
	)
	promoted := "yes"
	page := 1
	limit := 10

	ChangesService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		t.Run("Header must contain ApiKey", func(t *testing.T) {
			authKey := req.Header.Get("Authorization")
			assert.True(t, strings.Contains(authKey, novuApiKey))
			assert.True(t, strings.HasPrefix(authKey, "ApiKey"))
		})

		t.Run("URL and request method is as expected", func(t *testing.T) {
			expectedURL := fmt.Sprintf("/v1/changes?promoted=%s&page=%s&limit=%s", promoted, strconv.Itoa(page), strconv.Itoa(limit))
			assert.Equal(t, http.MethodGet, req.Method)
			assert.Equal(t, expectedURL, req.RequestURI)
		})

		var resp lib.ChangesGetResponse
		err := payloadStringToStruct(getResponse, &resp)
		require.Nil(t, err)

		w.WriteHeader(http.StatusOK)
		bb, _ := json.Marshal(resp)
		w.Write(bb)
	}))

	defer ChangesService.Close()

	ctx := context.Background()

	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(ChangesService.URL)})
	q := lib.ChangesGetQuery{Page: 1, Limit: 10, Promoted: "yes"}
	resp, err := c.ChangesApi.GetChanges(ctx, q)
	require.Nil(t, err)
	assert.NotNil(t, resp)

	t.Run("Response is as expected", func(t *testing.T) {
		err := payloadStringToStruct(getResponse, &expectedResponse)
		require.Nil(t, err)
		assert.Equal(t, expectedResponse, resp)
	})
}

func TestChangesService_Apply_Success(t *testing.T) {
	const changeID = "62b51a44da1af31d109f5da7"
	var (
		expectedResponse lib.ChangesApplyResponse
	)

	ChangesService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		t.Run("Header must contain ApiKey", func(t *testing.T) {
			authKey := req.Header.Get("Authorization")
			assert.True(t, strings.Contains(authKey, novuApiKey))
			assert.True(t, strings.HasPrefix(authKey, "ApiKey"))
		})

		t.Run("URL and request method is as expected", func(t *testing.T) {
			expectedURL := "/v1/changes/" + changeID + "/apply"
			assert.Equal(t, http.MethodPost, req.Method)
			assert.Equal(t, expectedURL, req.RequestURI)
		})

		var resp lib.ChangesApplyResponse
		payloadStringToStruct(applyResponse, &resp)

		w.WriteHeader(http.StatusOK)
		bb, _ := json.Marshal(resp)
		w.Write(bb)
	}))

	defer ChangesService.Close()

	ctx := context.Background()

	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(ChangesService.URL)})

	resp, err := c.ChangesApi.ApplyChange(ctx, changeID)
	require.Nil(t, err)
	assert.NotNil(t, resp)

	t.Run("Response is as expected", func(t *testing.T) {
		// fileToStruct(filepath.Join("../testdata", "changes_apply_response.json"), &expectedResponse)
		payloadStringToStruct(applyResponse, &expectedResponse)
		assert.Equal(t, expectedResponse, resp)
	})
}

func TestChangesService_BulkApply_Success(t *testing.T) {
	var (
		changesBulkApplyPayload lib.ChangesBulkApplyPayload
		receivedBody            lib.ChangesBulkApplyPayload
		expectedRequest         lib.ChangesBulkApplyPayload
		expectedResponse        lib.ChangesApplyResponse
	)

	ChangesService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
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
			expectedURL := "/v1/changes/bulk/apply"
			assert.Equal(t, http.MethodPost, req.Method)
			assert.Equal(t, expectedURL, req.RequestURI)
		})

		t.Run("Request is as expected", func(t *testing.T) {
			payloadStringToStruct(bulkApplyPayload, &expectedRequest)
			assert.Equal(t, expectedRequest, receivedBody)
		})

		var resp lib.ChangesApplyResponse
		payloadStringToStruct(applyResponse, &resp)

		w.WriteHeader(http.StatusOK)
		bb, _ := json.Marshal(resp)
		w.Write(bb)
	}))

	defer ChangesService.Close()

	ctx := context.Background()
	payloadStringToStruct(bulkApplyPayload, &changesBulkApplyPayload)

	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(ChangesService.URL)})

	resp, err := c.ChangesApi.ApplyBulkChanges(ctx, changesBulkApplyPayload)
	require.Nil(t, err)
	assert.NotNil(t, resp)

	t.Run("Response is as expected", func(t *testing.T) {
		payloadStringToStruct(applyResponse, &expectedResponse)
		assert.Equal(t, expectedResponse, resp)
	})
}
