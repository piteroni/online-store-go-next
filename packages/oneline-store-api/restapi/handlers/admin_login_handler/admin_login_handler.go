package admin_login_handler

import (
	common "piteroni/online-store-api"
	"piteroni/online-store-api/restapi/auth"
	"piteroni/online-store-api/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

var _ operations.PostAdminLoginHandler = &AdminLoginHandler{}

type AdminLoginHandler struct {
	*AdminAuthenticator
	*auth.JWTTokenGenerator
	*common.AppLogger
}

func (h *AdminLoginHandler) Handle(params operations.PostAdminLoginParams) middleware.Responder {
	err := params.Body.Validate(nil)
	if err != nil {
		h.AppLogger.Print(err.Error())
		return operations.NewPostAdminLoginBadRequest()
	}

	u, err := h.AdminAuthenticator.Authenticate(*params.Body.LoginID, *params.Body.Password)
	if err != nil {
		_, ok := err.(*AuthenticationError)
		if ok {
			h.AppLogger.Print(err.Error())
			return operations.NewPostAdminLoginUnauthorized()
		} else {
			h.AppLogger.Error(err)
			return operations.NewPostAdminLoginInternalServerError()
		}
	}

	token, err := h.JWTTokenGenerator.GenerateJWTToken(u.ID)
	if err != nil {
		h.AppLogger.Error(err)
		return operations.NewPostAdminLoginInternalServerError()
	}

	p := operations.PostAdminLoginOKBody{
		Token: &token,
	}

	return operations.NewPostAdminLoginOK().WithPayload(&p)
}
