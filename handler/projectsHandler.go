package handler

import (
	"github.com/ZafirProjects/portfolio/view/projects"
	"github.com/labstack/echo/v4"
)

type ProjectsHandler struct{}

func (h *ProjectsHandler) HandleProjectRender(c echo.Context) error {
	return render(c, projects.RenderProjects())
}
