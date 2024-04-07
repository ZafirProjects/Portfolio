package handler

import (
	"github.com/ZafirProjects/portfolio/view/experience"
	"github.com/labstack/echo/v4"
)

type ExperienceHandler struct{}

func (h *ExperienceHandler) RenderExperience(c echo.Context) error {
	return render(c, experience.RenderExperience())
}
