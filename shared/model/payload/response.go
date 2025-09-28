package payload

import (
	"encoding/json"
	"strings"
	"team-maker-api/shared/model/errorenum"
)

type Response struct {
	Code         int    `json:"code"`
	ErrorCode    string `json:"error_code"`
	CodeMessage  string `json:"code_message"`
	Message      string `json:"message"`
	ErrorMapping any    `json:"error_mapping"`
	Data         any    `json:"data"`
}

type ResponseSubmit struct {
	Code             int    `json:"code"`
	ErrorCode        string `json:"error_code"`
	CodeMessage      string `json:"code_message"`
	Message          string `json:"message"`
	ErrorMapping     any    `json:"error_mapping"`
	Data             any    `json:"data"`
	ShowErrorMessage bool   `json:"show_error_message"`
}

type ErrorMappingRes struct {
	ErrorCodes   string `json:"error_codes"`
	ErrorMessage string `json:"error_message"`
	ErrorFrom    string `json:"error_from"`
}

func OrderConsultResponse(data any, errorMapping any, isReadyCreateOrder bool) any {
	var res Response
	if isReadyCreateOrder {
		res.Code = 200
		res.ErrorCode = ""
		res.CodeMessage = "Success"
		res.Message = "Ready Create Order"
	} else {
		res.Code = 400
		res.ErrorCode = ""
		res.CodeMessage = "Failed"
		res.Message = "Not Ready Create Order"
	}

	res.ErrorMapping = errorMapping
	res.Data = data

	return res
}

func NewSuccessResponse(data any, errorMapping any) any {
	var res Response
	res.Code = 200
	res.ErrorCode = ""
	res.CodeMessage = "Success"
	res.ErrorMapping = errorMapping
	res.Message = "Success Created"
	res.Data = data

	return res
}

func NewErrorResponse(err error, traceID string) any {
	var res Response
	res.CodeMessage = "false"

	et, ok := err.(errorenum.ErrorType)
	if !ok {
		res.Code = 500
		res.CodeMessage = err.Error()
		res.Message = err.Error()
		return res
	}

	// ERROR REDIS RETURN THIS MESSAGE
	if et.Code() == "ERDS0001" {
		res.Message = "Something went wrong, try again later"
	}

	res.ErrorCode = et.Code()
	res.CodeMessage = et.Error()
	return res
}

func NewErrorSubmitResponse(err error, traceID string) any {
	var res Response
	res.CodeMessage = "false"

	et, ok := err.(errorenum.ErrorType)
	if !ok {
		res.Code = 400
		res.CodeMessage = err.Error()
		res.Message = err.Error()
		return res
	}

	res.ErrorCode = et.Code()
	res.CodeMessage = et.Error()
	return res
}

func NewErrorSubmitResponseCode(err error, traceID string) ResponseSubmit {
	var res ResponseSubmit
	res.CodeMessage = "false"

	// IF ERROR SHIPMENT PGW
	if strings.Contains(err.Error(), "ERRSHIPPGW") {
		type ResponseErrorShipment struct {
			Code             int    `json:"code"`
			CodeType         string `json:"code_type"`
			CodeMessage      string `json:"code_message"`
			ShowErrorMessage bool   `json:"show_error_message"`
		}

		var responseErrorShipment *ResponseErrorShipment

		resultSplit := strings.Split(err.Error(), ">>")
		_ = json.Unmarshal([]byte(resultSplit[1]), &responseErrorShipment)
		res.Code = responseErrorShipment.Code
		res.CodeMessage = responseErrorShipment.CodeType
		res.Message = responseErrorShipment.CodeMessage
		res.ShowErrorMessage = responseErrorShipment.ShowErrorMessage
		res.ErrorCode = "ERRSHIPPGW"
	} else {
		et, ok := err.(errorenum.ErrorType)
		if !ok {
			res.Code = 400
			res.CodeMessage = err.Error()
			res.Message = err.Error()
			return res
		}

		res.ErrorCode = et.Code()
		res.CodeMessage = et.Error()
	}

	return res
}
