package util

func SuccessResponse(statusCode int, message string, data interface{}) map[string]interface{} {
	response := map[string]interface{}{
		"Status":  statusCode,
		"Message": message,
		"Data":    data,
	}

	return response

}

func FailResponse(statusCode int, message string) map[string]interface{} {
	response := map[string]interface{}{
		"Status":  statusCode,
		"Message": message,
	}

	return response
}

func MessageResponse(statusCode int, messageString string) map[string]interface{}{
	response := map[string]interface{}{
		"Status" : statusCode,
		"Message": messageString,
	}
	return response
}