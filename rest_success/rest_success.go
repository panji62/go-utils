package rest_success

//RestSuccess interface restsuccess
type RestSuccess interface {
	Message() string
}
type restSuccess struct {
	SuccessMessage string `json:"message"`
}

func (s restSuccess) Message() string {
	return s.SuccessMessage
}

//NewRestSuccess create new success message
func NewRestSuccess(message string) RestSuccess {
	return restSuccess{
		SuccessMessage: message,
	}
}
