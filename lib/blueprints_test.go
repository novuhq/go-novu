package lib_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"testing"

	"github.com/novuhq/go-novu/lib"
	"github.com/stretchr/testify/require"
)

func responseStringToStruct(str string, s interface{}) io.Reader {
	bb := []byte(str)
	json.Unmarshal(bb, s)
	return bytes.NewReader(bb)
}

const templateID = "650ae12e61c036a0a5419480"

var groupByCategory = `{
  "general": [
    {}
  ],
  "popular": {}
}`

func TestBlueprintService_GetGroupByCategory_Success(t *testing.T) {
	var expectedResponse lib.BlueprintGroupByCategoryResponse
	responseStringToStruct(groupByCategory, &expectedResponse)

	httpServer := createTestServer(t, TestServerOptions[io.Reader, *lib.BlueprintGroupByCategoryResponse]{
		expectedURLPath:    fmt.Sprintf("/v1/blueprints/group-by-category"),
		expectedSentMethod: http.MethodGet,
		expectedSentBody:   http.NoBody,
		responseStatusCode: http.StatusOK,
		responseBody:       &expectedResponse,
	})

	ctx := context.Background()
	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(httpServer.URL)})
	resp, err := c.BlueprintApi.GetGroupByCategory(ctx)

	require.NoError(t, err)
	require.Equal(t, resp, expectedResponse)
}

func TestBlueprintService_GetByTemplateID_Success(t *testing.T) {
	var expectedResponse lib.BlueprintByTemplateIdResponse
	fileToStruct(filepath.Join("../testdata", "blueprint_response.json"), &expectedResponse)

	httpServer := createTestServer(t, TestServerOptions[io.Reader, *lib.BlueprintByTemplateIdResponse]{
		expectedURLPath:    fmt.Sprintf("/v1/blueprints/%s", templateID),
		expectedSentMethod: http.MethodGet,
		expectedSentBody:   http.NoBody,
		responseStatusCode: http.StatusOK,
		responseBody:       &expectedResponse,
	})

	ctx := context.Background()
	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(httpServer.URL)})
	resp, err := c.BlueprintApi.GetByTemplateID(ctx, templateID)

	require.NoError(t, err)
	require.Equal(t, resp, expectedResponse)
}
