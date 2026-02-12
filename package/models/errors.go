package models

var (
	UnexpError          = "Unexpected error occurred while processing your request."
	ValidationError     = "Unable to validate input(s). Please try again."
	LoginError          = "These credentials do not match our records."
	ThrottleError       = "Too many login attempts. Please try again."
	RecordNotFoundError = "Records not found."
	ErrorOccurred       = "Error occurred. Please try again."
	ErrorDecoding       = "Error occurred while decoding the response. Please try again."
	Success             = "Success"
	CreatedBy           = int32(1) //default system user
	ParseError          = "Error parsing date"
)
