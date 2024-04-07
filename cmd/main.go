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
	app.GET("/experience", experienceHandler.RenderExperience)

	projectsHandler := handler.ProjectsHandler{}
	app.GET("/projects", projectsHandler.HandleProjectRender)

	contactHandler := handler.ContactHandler{}
	app.GET("/contact", contactHandler.HandleContactRender)

	//start server
	app.Logger.Fatal(app.Start("localhost:3000"))
}
