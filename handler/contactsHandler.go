package handler

import (
	"github.com/ZafirProjects/portfolio/view/contact"
	"github.com/labstack/echo/v4"
)

type ContactHandler struct{}

func (h *ContactHandler) HandleContactRender(c echo.Context) error {
	return render(c, contact.ContactRender())
}
