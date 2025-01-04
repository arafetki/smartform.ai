package handlers

import "github.com/arafetki/smartform.ai/backend/internals/services"

type Handler struct {
	UsersService *services.UsersService
	FormsService *services.FormsService
}
