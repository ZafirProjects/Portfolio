package handler

import (
	"github.com/ZafirProjects/portfolio/model"
	"github.com/ZafirProjects/portfolio/view/projects"
	"github.com/labstack/echo/v4"
)

type ProjectsHandler struct{}

var projectsList = []model.Project{
	{
		Title:        "Dissertation",
		Technologies: "Python, Scki-Kit Learn, Data Science",
		Description:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
	},
	{
		Title:        "Full Stack Application",
		Technologies: "Java, JavaScript, Spring, React, MySQL, Azure Cloud",
		Description:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
	},
	{
		Title:        "Go + HTMX SPA",
		Technologies: "Go, Templ, HTMX",
		Description:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
	},
}

func (h *ProjectsHandler) HandleProjectRender(c echo.Context) error {
	return render(c, projects.RenderProjects(projectsList))
}
