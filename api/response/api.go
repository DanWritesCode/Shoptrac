package response

// APIResponse is the default API response type, if the API
// isn't returning a data structure.
type APIResponse struct {
	Code int         `json:"code,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

// NewAPIResponse will return a new APIResponse struct
func NewAPIResponse(code int, data interface{}) *APIResponse {
	return &APIResponse{
		Code: code,
		Data: data,
	}
}
