package response

import "net/http"

type APIResponse struct {
	StatusCode int                    `json:"-"`
	Body       map[string]interface{} `json:"body"`
}

func BadRequest(message string) APIResponse {
	return APIResponse{
		StatusCode: http.StatusBadRequest,
		Body: map[string]interface{}{
			"status":  "error",
			"message": message,
		},
	}
}

func ValidationError(errors map[string]string) APIResponse {
	return APIResponse{
		StatusCode: http.StatusUnprocessableEntity,
		Body: map[string]interface{}{
			"status": "validation_error",
			"errors": errors,
		},
	}
}

func InternalServerError(message string) APIResponse {
	return APIResponse{
		StatusCode: http.StatusInternalServerError,
		Body: map[string]interface{}{
			"status":  "error",
			"message": message,
		},
	}
}

func Unauthorized(message string) APIResponse {
	return APIResponse{
		StatusCode: http.StatusUnauthorized,
		Body: map[string]interface{}{
			"status":  "unauthorized",
			"message": message,
		},
	}
}

func Created(message string, data interface{}) APIResponse {
	return APIResponse{
		StatusCode: http.StatusCreated,
		Body: map[string]interface{}{
			"status":  "success",
			"message": message,
			"data":    data,
		},
	}
}

func Success(message string, data interface{}) APIResponse {
	return APIResponse{
		StatusCode: http.StatusOK,
		Body: map[string]interface{}{
			"status":  "success",
			"message": message,
			"data":    data,
		},
	}
}
