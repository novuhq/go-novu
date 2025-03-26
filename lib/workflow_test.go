package lib_test

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"testing"

	"github.com/novuhq/go-novu/lib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const workflowId = "workflowId"
const WORKFLOW_PATH = "/v1/workflows"

func TestWorkflowService_List_Workflows_Success(t *testing.T) {
	body := map[string]string{}
	var expectedResponse *lib.WorkflowGetResponse = &lib.WorkflowGetResponse{
		Data: lib.WorkflowData{
			Id:                  workflowId,
			Name:                "workflowName",
			NotificationGroupId: "notificationGroupId",
			Tags:                []string{"tag1", "tag2"},
			Description:         "workflowDescription",
			Steps:               []interface{}(nil),
			Active:              true,
			Critical:            true,
			PreferenceSettings: lib.Channel{
				Email: true,
				Sms:   true,
				Chat:  true,
				InApp: true,
				Push:  true,
			},
			BlueprintId: "blueprintId",
			Data:        nil,
		},
	}

	httpServer := createTestServer(t, TestServerOptions[map[string]string, lib.WorkflowGetResponse]{
		expectedURLPath:    WORKFLOW_PATH,
		expectedSentMethod: http.MethodGet,
		expectedSentBody:   body,
		responseStatusCode: http.StatusOK,
		responseBody:       *expectedResponse,
	})

	ctx := context.Background()
	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(httpServer.URL)})
	resp, err := c.WorkflowApi.List(ctx, nil)

	require.NoError(t, err)
	require.Equal(t, expectedResponse, resp)
}

func TestWorkflowService_Create_Workflow_Success(t *testing.T) {

	var createWorkflowRequest *lib.WorkflowData = &lib.WorkflowData{
		Name:                "workflowName",
		NotificationGroupId: "notificationGroupId",
		Tags:                []string{"tag1", "tag2"},
		Description:         "workflowDescription",
		Steps:               []interface{}(nil),
		Active:              true,
		Critical:            true,
		PreferenceSettings: lib.Channel{
			Email: true,
			Sms:   true,
			Chat:  true,
			InApp: true,
			Push:  true,
		},
		BlueprintId: "blueprintId",
		Data:        nil,
	}

	var response *lib.WorkflowData
	fileToStruct(filepath.Join("../testdata", "create_workflow_response.json"), &response)

	httpServer := createTestServer(t, TestServerOptions[lib.WorkflowData, lib.WorkflowData]{
		expectedURLPath:    WORKFLOW_PATH,
		expectedSentMethod: http.MethodPost,
		expectedSentBody:   *createWorkflowRequest,
		responseStatusCode: http.StatusCreated,
		responseBody:       *response,
	})

	ctx := context.Background()
	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(httpServer.URL)})
	res, err := c.WorkflowApi.Create(ctx, *createWorkflowRequest)

	require.Equal(t, response, res)
	require.NoError(t, err)

}

func TestWorkflowService_Update_Success(t *testing.T) {

	var updateWorkflowRequest *lib.WorkflowData = &lib.WorkflowData{
		Name:                "workflowName",
		NotificationGroupId: "notificationGroupId",
		Tags:                []string{"tag1", "tag2"},
		Description:         "workflowDescription",
		Steps:               []interface{}(nil),
		Active:              false,
		Critical:            true,
		PreferenceSettings: lib.Channel{
			Email: true,
			Sms:   true,
			Chat:  true,
			InApp: true,
			Push:  true,
		},
		Data: nil,
	}
	var response *lib.WorkflowData
	fileToStruct(filepath.Join("../testdata", "update_workflow_response.json"), &response)

	httpServer := createTestServer(t, TestServerOptions[lib.WorkflowData, lib.WorkflowData]{
		expectedURLPath:    fmt.Sprintf("/v1/workflows/%s", workflowId),
		expectedSentMethod: http.MethodPut,
		expectedSentBody:   *updateWorkflowRequest,
		responseStatusCode: http.StatusOK,
		responseBody:       *response,
	})

	ctx := context.Background()
	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(httpServer.URL)})
	res, err := c.WorkflowApi.Update(ctx, workflowId, *updateWorkflowRequest)

	assert.Equal(t, response, res)
	require.NoError(t, err)

}

func TestWorkflowService_Delete_Success(t *testing.T) {

	var DeleteResponse *lib.WorkflowDeleteResponse = &lib.WorkflowDeleteResponse{
		Data: true,
	}

	httpServer := createTestServer(t, TestServerOptions[map[string]string, lib.WorkflowDeleteResponse]{
		expectedURLPath:    fmt.Sprintf("/v1/workflows/%s", workflowId),
		expectedSentMethod: http.MethodDelete,
		responseStatusCode: http.StatusOK,
		responseBody:       *DeleteResponse,
	})

	ctx := context.Background()
	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(httpServer.URL)})
	res, err := c.WorkflowApi.Delete(ctx, workflowId)

	require.Equal(t, DeleteResponse, res)
	require.NoError(t, err)
}

func TestWorkflowService_Get_Success(t *testing.T) {

	var response *lib.WorkflowData
	fileToStruct(filepath.Join("../testdata", "get_workflow_response.json"), &response)

	httpServer := createTestServer(t, TestServerOptions[map[string]string, lib.WorkflowData]{
		expectedURLPath:    fmt.Sprintf("/v1/workflows/%s", workflowId),
		expectedSentMethod: http.MethodGet,
		responseStatusCode: http.StatusOK,
		responseBody:       *response,
	})

	ctx := context.Background()
	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(httpServer.URL)})
	res, err := c.WorkflowApi.Get(ctx, workflowId)

	require.Equal(t, response, res)
	require.NoError(t, err)
}

func TestWorkflowService_UpdateStatus_Success(t *testing.T) {

	var response *lib.WorkflowData
	fileToStruct(filepath.Join("../testdata", "update_workflow_status_response.json"), &response)
	httpServer := createTestServer(t, TestServerOptions[map[string]string, lib.WorkflowData]{
		expectedURLPath:    fmt.Sprintf("/v1/workflows/%s/status", workflowId),
		expectedSentMethod: http.MethodPut,
		responseStatusCode: http.StatusOK,
		responseBody:       *response,
	})

	ctx := context.Background()
	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(httpServer.URL)})
	res, err := c.WorkflowApi.UpdateStatus(ctx, workflowId)

	require.Equal(t, response, res)
	require.NoError(t, err)
}
