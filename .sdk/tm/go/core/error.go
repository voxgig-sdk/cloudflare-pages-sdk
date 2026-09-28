package core

type CloudflarePagesError struct {
	IsCloudflarePagesError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewCloudflarePagesError(code string, msg string, ctx *Context) *CloudflarePagesError {
	return &CloudflarePagesError{
		IsCloudflarePagesError: true,
		Sdk:              "CloudflarePages",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *CloudflarePagesError) Error() string {
	return e.Msg
}
