package lib_test

import (
	"context"
	"net/http"
	"path/filepath"
	"testing"

	"github.com/novuhq/go-novu/lib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const WORKFLOW_GROUP_PATH = "/v1/notification-groups"

var ctx = context.Background()

func TestCreateWorkflowGroup_Success(t *testing.T) {
	var response *lib.GetWorkflowResponse
	fileToStruct(filepath.Join("../testdata", "create_workflow_group_response.json"), &response)
	name := "workflow-group-name"
	httpServer := createTestServer(t, TestServerOptions[lib.CreateWorkflowGroupRequest, *lib.GetWorkflowResponse]{
		expectedURLPath:    WORKFLOW_GROUP_PATH,
		expectedSentMethod: http.MethodPost,
		expectedSentBody: lib.CreateWorkflowGroupRequest{
			Name: name,
		},
		responseStatusCode: http.StatusCreated,
		responseBody:       response,
	})

	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(httpServer.URL)})
	res, err := c.WorkflowGroupApi.Create(ctx, name)

	assert.Equal(t, response, res)
	require.NoError(t, err)
}

func TestGetWorkflowGroup_Success(t *testing.T) {
	var response *lib.GetWorkflowResponse
	fileToStruct(filepath.Join("../testdata", "get_workflow_group_response.json"), &response)
	workflowId := "6425cb064a2a919bc61b0365"
	httpServer := createTestServer(t, TestServerOptions[lib.CreateWorkflowGroupRequest, *lib.GetWorkflowResponse]{
		expectedURLPath:    WORKFLOW_GROUP_PATH + "/" + workflowId,
		expectedSentMethod: http.MethodGet,
		responseStatusCode: http.StatusOK,
		responseBody:       response,
	})

	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(httpServer.URL)})
	res, err := c.WorkflowGroupApi.Get(ctx, workflowId)

	assert.Equal(t, response, res)
	require.NoError(t, err)
}

func TestListWorkflowGroup_Success(t *testing.T) {
	var response *lib.ListWorkflowGroupsResponse
	fileToStruct(filepath.Join("../testdata", "list_workflow_groups_response.json"), &response)
	httpServer := createTestServer(t, TestServerOptions[lib.CreateWorkflowGroupRequest, *lib.ListWorkflowGroupsResponse]{
		expectedURLPath:    WORKFLOW_GROUP_PATH,
		expectedSentMethod: http.MethodGet,
		responseStatusCode: http.StatusOK,
		responseBody:       response,
	})

	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(httpServer.URL)})
	res, err := c.WorkflowGroupApi.List(ctx)

	assert.Equal(t, response, res)
	require.NoError(t, err)
}

func TestDeleteWorkflowGroup_Success(t *testing.T) {
	var response *lib.DeleteWorkflowGroupResponse
	fileToStruct(filepath.Join("../testdata", "delete_workflow_group_response.json"), &response)
	workflowId := "6425cb064a2a919bc61b0365"
	httpServer := createTestServer(t, TestServerOptions[any, *lib.DeleteWorkflowGroupResponse]{
		expectedURLPath:    WORKFLOW_GROUP_PATH + "/" + workflowId,
		expectedSentMethod: http.MethodDelete,
		responseStatusCode: http.StatusOK,
		responseBody:       response,
	})

	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(httpServer.URL)})
	res, err := c.WorkflowGroupApi.Delete(ctx, workflowId)

	assert.Equal(t, response, res)
	require.NoError(t, err)
}

func TestUpdateWorkflowGroup_Success(t *testing.T) {
	var response *lib.GetWorkflowResponse
	fileToStruct(filepath.Join("../testdata", "update_workflow_group_response.json"), &response)
	workflowId := "6425cb064a2a919bc61b0365"
	updatedRequest := lib.UpdateWorkflowGroupRequest{
		Name: "updated-name",
	}
	httpServer := createTestServer(t, TestServerOptions[lib.UpdateWorkflowGroupRequest, *lib.GetWorkflowResponse]{
		expectedURLPath:    WORKFLOW_GROUP_PATH + "/" + workflowId,
		expectedSentMethod: http.MethodPatch,
		expectedSentBody:   updatedRequest,
		responseStatusCode: http.StatusOK,
		responseBody:       response,
	})

	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(httpServer.URL)})
	res, err := c.WorkflowGroupApi.Update(ctx, workflowId, updatedRequest)

	assert.Equal(t, response, res)
	require.NoError(t, err)
}
