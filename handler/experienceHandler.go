package handler

import (
	"github.com/ZafirProjects/portfolio/model"
	"github.com/ZafirProjects/portfolio/view/experience"
	"github.com/labstack/echo/v4"
)

type ExperienceHandler struct{}

var experiences = []model.Experience{
	{
		Role:        "Software Engineering Intern",
		Company:     "Mimecast",
		Duration:    "August 2021 - September 2022",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
	},
	{
		Role:        "Full Stack Developer",
		Company:     "Brunel University London",
		Duration:    "February 2023 - March 2023",
		Description: "Aliquam ut porttitor leo a diam sollicitudin tempor. Volutpat commodo sed egestas egestas fringilla phasellus. Id neque aliquam vestibulum morbi blandit cursus. Ultricies lacus sed turpis tincidunt id aliquet risus feugiat. Eget velit aliquet sagittis id consectetur. Senectus et netus et malesuada fames. Auctor eu augue ut lectus arcu. Dui sapien eget mi proin sed libero enim sed faucibus. Feugiat in ante metus dictum at tempor commodo. Aliquam nulla facilisi cras fermentum. Lorem donec massa sapien faucibus et. Non nisi est sit amet. Vestibulum rhoncus est pellentesque elit. Non tellus orci ac auctor.",
	},
}

var experienceIndex int = 0

func (h *ExperienceHandler) RenderExperience(c echo.Context) error {
	experienceIndex = 0
	return render(c, experience.RenderExperience(experiences, experienceIndex))
}

func (h *ExperienceHandler) RenderNextExperience(c echo.Context) error {
	experienceIndex++
	if experienceIndex%len(experiences) == 0 {
		experienceIndex = 0
	}
	return render(c, experience.ExperienceCard(experiences, experienceIndex))
}

func (h *ExperienceHandler) RenderPreviousExperience(c echo.Context) error {
	experienceIndex--
	if experienceIndex < 0 {
		experienceIndex = len(experiences) - 1
	}
	return render(c, experience.ExperienceCard(experiences, experienceIndex))
}
