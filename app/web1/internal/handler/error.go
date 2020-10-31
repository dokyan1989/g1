package handler

// Error ...
type Error interface {
	error
	Status() int
}

// StatusError ...
type StatusError struct {
	Code int
	Err  error
}

// Error ...
func (se StatusError) Error() string {
	return se.Err.Error()
}

// Status ...
func (se StatusError) Status() int {
	return se.Code
}

// FieldError ...
type FieldError struct {
	Message string `json:"message,omitempty"`
	Field   string `json:"field,omitempty"`
}

// ErrorResponse ...
type ErrorResponse struct {
	Message     string       `json:"message,omitempty"`
	FieldErrors []FieldError `json:"errors,omitempty"`
}
