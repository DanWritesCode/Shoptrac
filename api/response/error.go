package response

// ErrorResponse is the default error response for the API.
type ErrorResponse struct {
	Code  int    `json:"code,omitempty"`
	Error string `json:"error,omitempty"`
}

// Default errors.
var (
	BadRequestError = NewError(400, "400 Bad request.")
	InternalError   = NewError(500, "An internal server error occurred.")
	NotFoundError   = NewError(404, "Endpoint Not Found")
  NotAllowedError   = NewError(405, "Method Not Allowed")
)

// NewError returns a new ErrorResponse struct.
func NewError(code int, err string) *ErrorResponse {
	return &ErrorResponse{
		Code:  code,
		Error: err,
	}
}
