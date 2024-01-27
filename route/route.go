package route

import (
	"tipen-demo/handler"
	"tipen-demo/middleware"
	"tipen-demo/pkg"

	"github.com/gofiber/fiber/v2"
)

type Route struct {
	Router     fiber.Router
	Handler    *handler.Handler
	Validator  *pkg.Validator
	Middleware *middleware.Middleware
}

type RouteParams struct {
	App        *fiber.App
	Handler    *handler.Handler
	Middleware *middleware.Middleware
}

func NewRoute(r RouteParams) *Route {
	return &Route{
		Router:     r.App.Group("/api"),
		Handler:    r.Handler,
		Validator:  &pkg.Validator{},
		Middleware: r.Middleware,
	}
}

func (r *Route) Init() {
	r.InitUser()
	r.InitSeller()
}
