package models

import (
	"net/http"
	"strings"
	"time"
)

type ResponseDataModel struct {
	TimeStamp  time.Time `json:"timestamp"`
	StatusCode int32     `json:"status_code"`
	HttpStatus string    `json:"http_status"`
	Message    string    `json:"message"`
	Count      int32     `json:"count"`
	Data       interface {
	} `json:"data,omitempty"`
}

func ResponseSuccessData(data interface{}) *ResponseDataModel {
	//val := reflect.ValueOf(data)
	msgDataResp := &ResponseDataModel{
		TimeStamp:  time.Now(),
		StatusCode: http.StatusOK,
		HttpStatus: http.StatusText(http.StatusOK),
		Message:    Success,
		//Count:      int32(val.Len()),
		Data: data,
	}
	return msgDataResp
}

func ResponseErrorData(data interface{}, errMsg string) *ResponseDataModel {
	msgDataResp := &ResponseDataModel{
		TimeStamp:  time.Now(),
		StatusCode: http.StatusInternalServerError,
		HttpStatus: http.StatusText(http.StatusInternalServerError),
		Message:    errMsg,
		Data:       data,
	}
	return msgDataResp
}

func JsonResponseMessage(message, controllerName string, isCreated, isDeleted bool) string {
	if strings.Contains(strings.ToLower(message), strings.ToLower("Success")) {
		if isCreated {
			message = controllerName + " created successfully"
			return message
		} else if isDeleted {
			message = controllerName + " deleted successfully"
			return message
		}
		message = controllerName + " updated successfully"
		return message
	} else {
		if isCreated {
			message = "Error occurred. " + controllerName + " not created"
			return message
		} else if isDeleted {
			message = "Error occurred. " + controllerName + " not deleted"
			return message
		}
		message = "Error occurred. " + controllerName + " not updated"
		return message
	}
}

func Booleans() map[string]string {
	myBooleans := map[string]string{
		"True":  "True",
		"False": "False",
	}
	return myBooleans
}
