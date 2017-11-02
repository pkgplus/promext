package response

const (
	STATUS_SUC = "success"
	STATUS_ERR = "error"
)

type Response struct {
	Status    string      `json:"status"`
	Data      interface{} `json:"data"`
	ErrorType string      `json:"errorType,omitempty"`
	Error     string      `json:"error,omitempty"`
}

func NewSucResponse(data interface{}) Response {
	return Response{STATUS_SUC, data, "", ""}
}

func NewBadResponse(data interface{}, errType string, err error) Response {
	return Response{STATUS_ERR, data, errType, err.Error()}
}
