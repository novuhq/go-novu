package lib

import (
	"context"
	"net/http"
)

type BlueprintService service

func (b *BlueprintService) GetGroupByCategory(ctx context.Context) (BlueprintGroupByCategoryResponse, error) {
	var resp BlueprintGroupByCategoryResponse
	URL := b.client.config.BackendURL.JoinPath("blueprints", "group-by-category")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL.String(), http.NoBody)
	if err != nil {
		return resp, err
	}

	_, err = b.client.sendRequest(req, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (b *BlueprintService) GetByTemplateID(ctx context.Context, templateID string) (BlueprintByTemplateIdResponse, error) {
	var resp BlueprintByTemplateIdResponse
	URL := b.client.config.BackendURL.JoinPath("blueprints", templateID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL.String(), http.NoBody)
	if err != nil {
		return resp, err
	}

	_, err = b.client.sendRequest(req, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
