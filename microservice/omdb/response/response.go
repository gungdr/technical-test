package response

func Success(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"data":  data,
		"error": nil,
	}
}

func Error(err error) map[string]interface{} {
	return map[string]interface{}{
		"data":  nil,
		"error": err.Error(),
	}
}
