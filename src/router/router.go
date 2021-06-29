package router

import (
	"github.com/labstack/echo"
	"github.com/ranon-rat/video-transmission/src/controllers"
)

func SetupRouter() {
	e := echo.New()
	e.GET("/image", controllers.RenderImage)
	e.File("/", "view/main.html")
	/*e.GET("/", func(c echo.Context) error {
		c.String(200, "<img src=\"/image\">")
		//c.Response().Write([]byte(`<img src="/image">`))
		return nil
	})*/
	e.Logger.Fatal(e.Start(":8080"))
}
