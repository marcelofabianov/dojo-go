package handler

// ErrorContext provides contextual information about a validation error.
type ErrorContext struct {
	Field string `json:"field"`
	Param string `json:"param"`
	Tag   string `json:"tag"`
}

// ErrorDetail provides detailed information about a single validation error.
type ErrorDetail struct {
	Message string       `json:"message"`
	Code    string       `json:"code"`
	Context ErrorContext `json:"context"`
}

// ErrorResponse defines the standard error response structure for the API,
// including validation details.
type ErrorResponse struct {
	Message string        `json:"message" example:"Request validation failed"`
	Code    string        `json:"code" example:"invalid_input"`
	Details []ErrorDetail `json:"details,omitempty"`
}
