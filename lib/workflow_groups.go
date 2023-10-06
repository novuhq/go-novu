package lib

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type IWorkflowGroup interface {
	Create(ctx context.Context, name string) (*GetWorkflowResponse, error)
	Get(ctx context.Context, workflowGroupID string) (*GetWorkflowResponse, error)
	List(ctx context.Context) (*ListWorkflowGroupsResponse, error)
	Delete(ctx context.Context, workflowGroupID string) (*DeleteWorkflowGroupResponse, error)
	Update(ctx context.Context, workflowGroupID string, request UpdateWorkflowGroupRequest) (*GetWorkflowResponse, error)
}

type WorkflowGroupService service

const WORKFLOW_GROUP_PATH = "notification-groups"

func (t *WorkflowGroupService) Create(ctx context.Context, name string) (*GetWorkflowResponse, error) {
	var resp GetWorkflowResponse

	URL := t.client.config.BackendURL.JoinPath(WORKFLOW_GROUP_PATH)
	reqBody := CreateWorkflowGroupRequest{
		Name: name,
	}

	jsonBody, _ := json.Marshal(reqBody)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, URL.String(), bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	httpResponse, err := t.client.sendRequest(req, &resp)
	if err != nil {
		return nil, err
	}

	if httpResponse.StatusCode != HTTPStatusCreated {
		return nil, errors.Wrap(err, "unable to create workflow group")
	}

	return &resp, nil
}

func (t *WorkflowGroupService) List(ctx context.Context) (*ListWorkflowGroupsResponse, error) {
	var resp ListWorkflowGroupsResponse
	URL := t.client.config.BackendURL.JoinPath(WORKFLOW_GROUP_PATH)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL.String(), http.NoBody)
	if err != nil {
		return nil, err
	}

	_, err = t.client.sendRequest(req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (t *WorkflowGroupService) Get(ctx context.Context, workflowGroupID string) (*GetWorkflowResponse, error) {
	var resp GetWorkflowResponse

	URL := t.client.config.BackendURL.JoinPath(WORKFLOW_GROUP_PATH, workflowGroupID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL.String(), bytes.NewBuffer([]byte{}))
	if err != nil {
		return nil, err
	}

	_, err = t.client.sendRequest(req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (t *WorkflowGroupService) Delete(ctx context.Context, workflowGroupID string) (*DeleteWorkflowGroupResponse, error) {
	var resp DeleteWorkflowGroupResponse
	URL := t.client.config.BackendURL.JoinPath(WORKFLOW_GROUP_PATH, workflowGroupID)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, URL.String(), http.NoBody)
	if err != nil {
		return nil, err
	}

	_, err = t.client.sendRequest(req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (t *WorkflowGroupService) Update(ctx context.Context, workflowGroupID string, request UpdateWorkflowGroupRequest) (*GetWorkflowResponse, error) {
	var resp GetWorkflowResponse

	URL := t.client.config.BackendURL.JoinPath(WORKFLOW_GROUP_PATH, workflowGroupID)
	jsonBody, _ := json.Marshal(request)

	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, URL.String(), bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	_, err = t.client.sendRequest(req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
