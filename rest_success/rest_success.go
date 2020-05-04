package rest_success

//RestSuccess interface rest success
type RestSuccess interface {
	Message() string
}
type restSuccess struct {
	SuccessMessage string `json:"message"`
}

func (s restSuccess) Message() string {
	return s.SuccessMessage
}

//NewRestSuccess create new success message, input string message to make custom success message
func NewRestSuccess(message string) RestSuccess {
	return restSuccess{
		SuccessMessage: message,
	}
}
