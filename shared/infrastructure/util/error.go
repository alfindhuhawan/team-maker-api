package util

import (
	"team-maker-api/shared/model/errorenum"
)

func GetErrorFormat(err error) (string, string) {
	errorCode := "UNDEFINED"
	errorMessage := err.Error()

	et, ok := err.(errorenum.ErrorType)
	if ok {
		errorCode = et.Code()
		errorMessage = et.Error()
	}

	return errorCode, errorMessage
}
