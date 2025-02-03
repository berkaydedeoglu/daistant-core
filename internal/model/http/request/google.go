package request

type GoogleAuthCallbackRequest struct {
	Code string `form:"code"`
}
