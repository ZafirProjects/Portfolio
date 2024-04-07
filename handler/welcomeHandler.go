package handler

import (
	"github.com/ZafirProjects/portfolio/view/welcome"
	"github.com/labstack/echo/v4"
)

type WelcomeHandler struct{}

func (h *WelcomeHandler) RenderWelcome(c echo.Context) error {
	return render(c, welcome.RenderWelcome())
}
