package response

type errorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(message string) *errorResponse {
	return &errorResponse{Message: message}
}

func NewErrorResponseFrom(err error) *errorResponse {
	return &errorResponse{Message: err.Error()}
}
