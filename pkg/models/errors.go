package models

type ErrorDescription struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

type ErrorResponse struct {
	// Message - информация об ошибке
	Message string `json:"message"`
}

type ValidationErrorResponse struct {
	// Message - информация об ошибке
	Message string `json:"message"`
	// Errors - Массив полей с описанием ошибки
	Errors []ErrorDescription `json:"errors"`
}
