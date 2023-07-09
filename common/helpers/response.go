package helpers

import "infra-book-crud/common/types"

func ApiResponse(status int, message string, data interface{}) types.ApiResponse {
	return types.ApiResponse{Status: status, Message: message, Data: data}
}
