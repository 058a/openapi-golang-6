package items_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"openapi/internal/infrastructure/env"
	oapicodegen "openapi/internal/infrastructure/oapicodegen/stock"

	"github.com/google/uuid"
)

type RequestHelper struct {
	client *http.Client
}

func (h *RequestHelper) Post(reqBody *oapicodegen.PostStockItemJSONRequestBody) (*http.Response, error) {
	reqBodyJson, _ := json.Marshal(reqBody)
	req, err := http.NewRequest(
		http.MethodPost,
		env.GetServiceUrl()+"/stock/items",
		bytes.NewBuffer(reqBodyJson),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *RequestHelper) Put( stockItemId uuid.UUID, reqBody *oapicodegen.PutStockItemJSONRequestBody) (*http.Response, error) {
	reqBodyJson, _ := json.Marshal(reqBody)
	req, err := http.NewRequest(
		http.MethodPut,
		env.GetServiceUrl()+"/stock/items/"+stockItemId.String(),
		bytes.NewBuffer(reqBodyJson),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *RequestHelper) Delete(stockItemId uuid.UUID) (*http.Response, error) {
	req, err := http.NewRequest(
		http.MethodDelete,
		env.GetServiceUrl()+"/stock/items/"+stockItemId.String(),
		nil,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type ResponseConvertHelper struct{}

func (h *ResponseConvertHelper) AsCreated(res *http.Response) (*oapicodegen.Created, error) {
	resBodyByte, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resBody := &oapicodegen.Created{}
	json.Unmarshal(resBodyByte, &resBody)
	if resBody.Id == uuid.Nil {
		return nil, fmt.Errorf("expected not empty, actual empty")
	}
	return resBody, nil
}

func (h *ResponseConvertHelper) AsBadRequest(res *http.Response) (*oapicodegen.BadRequestResponse, error) {
	resBodyByte, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resBody := &oapicodegen.BadRequest{}
	json.Unmarshal(resBodyByte, &resBody)

	return resBody, nil
}