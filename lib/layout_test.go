package lib_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/novuhq/go-novu/lib"
	"github.com/stretchr/testify/require"
)

const LayoutId = "2222"

func TestLayoutService_Create_Layout_Success(t *testing.T) {
	var createLayoutRequest *lib.CreateLayoutRequest = &lib.CreateLayoutRequest{
		Name:        "layoutName",
		Identifier:  "layoutIdentifier",
		Description: "layoutDescription",
		Content:     "layoutContent",
		Variables:   []interface{}(nil),
		IsDefault:   true,
	}
	res, _ := json.Marshal(createLayoutRequest)
	fmt.Println(string(res))
	var expectedResponse *lib.CreateLayoutResponse = &lib.CreateLayoutResponse{
		Data: struct {
			Id string `json:"_id"`
		}{
			Id: "2222",
		},
	}

	httpServer := createTestServer(t, TestServerOptions[lib.CreateLayoutRequest, lib.CreateLayoutResponse]{
		expectedURLPath:    "/v1/layouts",
		expectedSentBody:   *createLayoutRequest,
		expectedSentMethod: http.MethodPost,
		responseStatusCode: http.StatusCreated,
		responseBody:       *expectedResponse,
	})

	ctx := context.Background()
	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(httpServer.URL)})
	resp, err := c.LayoutApi.Create(ctx, *createLayoutRequest)

	require.NoError(t, err)
	require.Equal(t, expectedResponse, resp)
}

func TestLayoutService_List_Layouts_Success(t *testing.T) {
	body := map[string]string{}
	var expectedResponse *lib.LayoutsResponse = &lib.LayoutsResponse{
		Page:       0,
		PageSize:   20,
		TotalCount: 1,
		Data: []lib.LayoutResponse{{
			Id:             "id",
			OrganizationId: "orgId",
			EnvironmentId:  "envId",
			CreatorId:      "creatorId",
			Name:           "layoutName",
			Identifier:     "layoutIdentifier",
			Description:    "layoutDescription",
			Channel:        "in_app",
			Content:        "layoutContent",
			ContentType:    "layoutContentType",
			Variables:      []interface{}{},
			IsDefault:      true,
			IsDeleted:      false,
			CreatedAt:      "createdAt",
			UpdatedAt:      "updatedAt",
			ParentId:       "parentId",
		}},
	}
	httpServer := createTestServer(t, TestServerOptions[map[string]string, lib.LayoutsResponse]{
		expectedURLPath:    "/v1/layouts",
		expectedSentMethod: http.MethodGet,
		expectedSentBody:   body,
		responseStatusCode: http.StatusCreated,
		responseBody:       *expectedResponse,
	})

	ctx := context.Background()
	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(httpServer.URL)})
	resp, err := c.LayoutApi.List(ctx, nil)

	require.NoError(t, err)
	require.Equal(t, expectedResponse, resp)
}

func TestLayoutService_Get_Layout_Success(t *testing.T) {

	var expectedResponse *lib.LayoutResponse = &lib.LayoutResponse{
		Id:             "id",
		OrganizationId: "orgId",
		EnvironmentId:  "envId",
		CreatorId:      "creatorId",
		Name:           "layoutName",
		Identifier:     "layoutIdentifier",
		Description:    "layoutDescription",
		Channel:        "in_app",
		Content:        "layoutContent",
		ContentType:    "layoutContentType",
		Variables:      []interface{}{},
		IsDefault:      true,
		IsDeleted:      false,
		CreatedAt:      "createdAt",
		UpdatedAt:      "updatedAt",
		ParentId:       "parentId",
	}

	httpServer := createTestServer(t, TestServerOptions[map[string]string, lib.LayoutResponse]{
		expectedURLPath:    fmt.Sprintf("/v1/layouts/%s", LayoutId),
		expectedSentMethod: http.MethodGet,
		responseStatusCode: http.StatusOK,
		responseBody:       *expectedResponse,
	})

	ctx := context.Background()
	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(httpServer.URL)})
	resp, err := c.LayoutApi.Get(ctx, "2222")

	require.NoError(t, err)
	require.Equal(t, expectedResponse, resp)
}

func TestLayoutService_Delete_Layout_Success(t *testing.T) {

	body := map[string]string{}

	httpServer := createTestServer(t, TestServerOptions[map[string]string, map[string]string]{
		expectedURLPath:    fmt.Sprintf("/v1/layouts/%s", LayoutId),
		expectedSentMethod: http.MethodDelete,
		responseStatusCode: http.StatusOK,
		responseBody:       body,
	})

	ctx := context.Background()
	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(httpServer.URL)})
	err := c.LayoutApi.Delete(ctx, "2222")

	require.NoError(t, err)
}

func TestLayoutService_Update_Layout_Success(t *testing.T) {

	var updateLayoutRequest *lib.CreateLayoutRequest = &lib.CreateLayoutRequest{
		Name:        "layoutName",
		Identifier:  "layoutIdentifier",
		Description: "layoutDescription",
		Content:     "layoutContent",
		Variables:   []interface{}(nil),
		IsDefault:   false,
	}
	res, _ := json.Marshal(updateLayoutRequest)
	fmt.Println(string(res))
	var expectedResponse *lib.LayoutResponse = &lib.LayoutResponse{
		Id:             "id",
		OrganizationId: "orgId",
		EnvironmentId:  "envId",
		CreatorId:      "creatorId",
		Name:           "layoutName",
		Identifier:     "layoutIdentifier",
		Description:    "layoutDescription",
		Channel:        "in_app",
		Content:        "layoutContent",
		ContentType:    "layoutContentType",
		Variables:      []interface{}{},
		IsDefault:      true,
		IsDeleted:      false,
		CreatedAt:      "createdAt",
		UpdatedAt:      "updatedAt",
		ParentId:       "parentId",
	}
	httpServer := createTestServer[lib.CreateLayoutRequest, lib.LayoutResponse](t, TestServerOptions[lib.CreateLayoutRequest, lib.LayoutResponse]{
		expectedURLPath:    fmt.Sprintf("/v1/layouts/%s", LayoutId),
		expectedSentBody:   *updateLayoutRequest,
		expectedSentMethod: http.MethodPatch,
		responseStatusCode: http.StatusOK,
		responseBody:       *expectedResponse,
	})
	ctx := context.Background()
	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(httpServer.URL)})
	resp, err := c.LayoutApi.Update(ctx, "2222", *updateLayoutRequest)
	require.NoError(t, err)
	require.Equal(t, expectedResponse, resp)
}

func TestLayoutService_Layout_SetDefault(t *testing.T) {

	body := map[string]string{}

	httpServer := createTestServer(t, TestServerOptions[map[string]string, map[string]string]{
		expectedURLPath:    fmt.Sprintf("/v1/layouts/%s/default", LayoutId),
		expectedSentBody:   body,
		expectedSentMethod: http.MethodPost,
		responseStatusCode: http.StatusNoContent,
		responseBody:       body,
	})

	ctx := context.Background()
	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(httpServer.URL)})
	err := c.LayoutApi.SetDefault(ctx, LayoutId)

	require.NoError(t, err)
}
