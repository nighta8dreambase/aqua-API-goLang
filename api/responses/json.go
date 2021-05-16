package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	CODE    int                    `json:"code"`
	MESSAGE string                 `json:"message"`
	DATA    map[string]interface{} `json:"data"`
}

type ResponseList struct {
	CODE    int                      `json:"code"`
	MESSAGE string                   `json:"message"`
	DATA    []map[string]interface{} `json:"data"`
}

func Unauthorized() Response {
	res := Response{CODE: 0, MESSAGE: "access_denied", DATA: nil}
	return res
}

func Success(data map[string]interface{}) Response {
	res := Response{CODE: 1, MESSAGE: "success", DATA: data}
	return res
}

func SuccessList(data []map[string]interface{}) ResponseList {
	res := ResponseList{CODE: 1, MESSAGE: "success", DATA: data}
	return res
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
