package view

import (
	"transfer-service/model"
)

const (
	ResponseCodeInvalidRequest = 400
	ResponseCodeGeneralError   = 500
)

type GenericResponse struct {
	Code    int64  `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

type AccountQueryData struct {
	AccountID int64  `json:"account_id"`
	Balance   string `json:"balance"`
}

func NewInvalidRequestResp(err error) GenericResponse {
	return GenericResponse{
		Code:    ResponseCodeInvalidRequest,
		Message: err.Error(),
	}
}

func NewGeneralErrorResp(err error) GenericResponse {
	return GenericResponse{
		Code:    ResponseCodeGeneralError,
		Message: err.Error(),
	}
}

func NewSuccessResp(data any) GenericResponse {
	return GenericResponse{
		Data:    data,
		Message: "success",
	}
}

func BuildAccountQueryData(input model.Account) AccountQueryData {
	return AccountQueryData{
		AccountID: input.ID,
		Balance:   input.Balance.String(),
	}
}
