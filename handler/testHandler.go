package handler

import (
	"strconv"

	"github.com/ZafirProjects/portfolio/view/test"
	"github.com/labstack/echo/v4"
)

type TestHandler struct{}

var count int = 0

func (h TestHandler) HandleTestShow(c echo.Context) error {
	return render(c, test.Show(strconv.Itoa(count)))
}

func (h TestHandler) HandleTestCountShow(c echo.Context) error {
	count++
	return render(c, test.Count(strconv.Itoa(count)))
}
