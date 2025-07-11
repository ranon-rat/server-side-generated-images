package router

import (
	"os"

	"github.com/labstack/echo"
	"github.com/ranon-rat/server-side-generated-images/src/controllers"
)

func SetupRouter() {
	e := echo.New()
	e.GET("/simple-fractal", controllers.SimpleFractal)
	e.GET("/mandelbrot", controllers.MandelbrotSet)
	e.GET("/lorenz", controllers.LorenzSystem)
	/*e.GET("/image-zy", controllers.ZY)
	e.GET("/image-yx", controllers.ZX)*/
	e.File("/", "src/view/main.html")
	/*e.GET("/", func(c echo.Context) error {
		c.String(200, "<img src=\"/image\">")
		//c.Response().Write([]byte(`<img src="/image">`))
		return nil
	})*/

	port, exist := os.LookupEnv("PORT")
	if !exist {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}

