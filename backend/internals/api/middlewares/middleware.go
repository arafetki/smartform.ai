package middlewares

import "github.com/arafetki/smartform.ai/backend/internals/services"

type Middleware struct {
	UsersService *services.UsersService
}
