package errors


type errorResponse struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Fields map[string]string `json:"fields,omitempty"`
}
