package Controllers

import (
	"github.com/dapsframework/daps/Routing"
	"github.com/dapsframework/dapsskeleton/app/Http/Middleware"
)

type WelcomeController struct {
	Routing.Controller
}

func (controller *WelcomeController) Init() {
	controller.AddMiddleware(new(Middleware.RequestLogger))
	controller.AddMiddleware(new(Middleware.ExtraContent))
}
func (controller *WelcomeController) Index() {

}
