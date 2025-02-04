package request

type GoogleAuthCallbackRequest struct {
	Code  string `form:"code"`
	State string `form:"state"`
	Scope string `form:"scope"`
}
