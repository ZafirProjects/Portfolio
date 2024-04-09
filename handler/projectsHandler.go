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
		Description:  "Aliquam ut porttitor leo a diam sollicitudin tempor. Volutpat commodo sed egestas egestas fringilla phasellus. Id neque aliquam vestibulum morbi blandit cursus. Ultricies lacus sed turpis tincidunt id aliquet risus feugiat. Eget velit aliquet sagittis id consectetur. Senectus et netus et malesuada fames. Auctor eu augue ut lectus arcu. Dui sapien eget mi proin sed libero enim sed faucibus. Feugiat in ante metus dictum at tempor commodo. Aliquam nulla facilisi cras fermentum. Lorem donec massa sapien faucibus et. Non nisi est sit amet. Vestibulum rhoncus est pellentesque elit. Non tellus orci ac auctor.",
	},
	{
		Title:        "Go + HTMX SPA",
		Technologies: "Go, Templ, HTMX",
		Description:  "Faucibus interdum posuere lorem ipsum dolor sit amet. Donec pretium vulputate sapien nec sagittis aliquam. Nulla pharetra diam sit amet nisl suscipit adipiscing bibendum. Fringilla est ullamcorper eget nulla facilisi etiam dignissim diam. Consectetur adipiscing elit duis tristique. Porttitor rhoncus dolor purus non. Morbi tristique senectus et netus. Dis parturient montes nascetur ridiculus mus mauris vitae ultricies. Eu scelerisque felis imperdiet proin fermentum leo. Adipiscing vitae proin sagittis nisl rhoncus. Molestie a iaculis at erat pellentesque. Felis bibendum ut tristique et egestas quis ipsum suspendisse ultrices. Proin nibh nisl condimentum id venenatis.",
	},
}

var projectIndex int = 0

func (h *ProjectsHandler) HandleProjectRender(c echo.Context) error {
	projectIndex = 0
	return render(c, projects.RenderProjects(projectsList, projectIndex))
}

func (h *ProjectsHandler) RenderNextProject(c echo.Context) error {
	projectIndex++
	if projectIndex%len(projectsList) == 0 {
		projectIndex = 0
	}
	return render(c, projects.ProjectCard(projectsList, projectIndex))
}

func (h *ProjectsHandler) RenderPreviousProject(c echo.Context) error {
	projectIndex--
	if projectIndex < 0 {
		projectIndex = len(projectsList) - 1
	}
	return render(c, projects.ProjectCard(projectsList, projectIndex))
}
