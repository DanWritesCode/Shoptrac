package response

import (
	"net/http"

	"github.com/unrolled/render"
)

var r = render.New()

// Error will display an error.
func Error(w http.ResponseWriter, errRes *ErrorResponse) {
	r.JSON(w, errRes.Code, errRes)
}

// BadRequest will render a Bad Request response.
func BadRequest(w http.ResponseWriter, msg string) {
	r.JSON(w, 400, NewError(400, msg))
}

// JSON will render a JSON response.
func JSON(w http.ResponseWriter, code int, payload interface{}) {
	r.JSON(w, code, payload)
}

// DefaultResponse will render a JSON response, with the default API format.
func DefaultResponse(w http.ResponseWriter, code int, data interface{}) {
	JSON(w, code, NewAPIResponse(code, data))
}
