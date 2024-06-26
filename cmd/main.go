package main

import (
	"github.com/ZafirProjects/portfolio/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func main() {
	app := echo.New()

	// Little bit of middlewares for housekeeping
	app.Use(middleware.Logger())
	app.Pre(middleware.RemoveTrailingSlash())
	app.Use(middleware.Recover())
	app.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
		rate.Limit(20),
	)))

	// routers with handlers
	testHandler := handler.TestHandler{}
	app.GET("/test", testHandler.HandleTestShow)
	app.POST("/count", testHandler.HandleTestCountShow)

	welcomeHandler := handler.WelcomeHandler{}
	app.GET("/", welcomeHandler.RenderWelcome)

	experienceHandler := handler.ExperienceHandler{}
	app.POST("/experience", experienceHandler.RenderExperience)
	app.POST("/experience/next", experienceHandler.RenderNextExperience)
	app.POST("/experience/previous", experienceHandler.RenderPreviousExperience)

	projectsHandler := handler.ProjectsHandler{}
	app.POST("/projects", projectsHandler.HandleProjectRender)
	app.POST("/projects/next", projectsHandler.RenderNextProject)
	app.POST("/projects/previous", projectsHandler.RenderPreviousProject)

	contactHandler := handler.ContactHandler{}
	app.POST("/contact", contactHandler.HandleContactRender)
	app.POST("/contact/sendEmail", contactHandler.HandleSendEmail)

	//start server
	app.Logger.Fatal(app.Start("localhost:3000"))
}
