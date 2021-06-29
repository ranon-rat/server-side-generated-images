package controllers

import (
	"log"

	//"github.com/fogleman/gg"

	"github.com/fogleman/gg"
	"github.com/labstack/echo"
)

const (
	S            = 600
	boundaryWord = "MJPEGBOUNDARY"
	headerf      = "\r\n" +
		"--" + boundaryWord + "\r\n" +
		"Content-Type: image/jpeg\r\n" +
		"X-Timestamp: 0.000000\r\n" +
		"\r\n"
)

var (
	width          = S
	height         = S
	div            = 2.5
	rad    float64 = S / 2
)

func drawCircle(x float64, y float64, radius float64, dc *gg.Context, c echo.Context) error {
	dc.SetRGB255(int(radius)%8*32, 0, int(radius)%64*128)
	dc.DrawCircle(x, y, radius)
	dc.Stroke()

	if radius > 1 {
		radius /= float64(div)
		//x
		if err := drawCircle((x + radius), y, radius, dc, c); err != nil {
			return err
		}
		if err := drawCircle((x - radius), y, radius, dc, c); err != nil {
			return err
		}

		//y
		if err := drawCircle(x, (y + radius), radius, dc, c); err != nil {
			return err
		}
		if err := drawCircle(x, (y - radius), radius, dc, c); err != nil {
			return err
		}

		if _, err := c.Response().Write([]byte(headerf)); err != nil {
			return err
		}
		if err := dc.EncodePNG(c.Response()); err != nil {
			return err
		}

	}
	return nil
}
func draw(c echo.Context) error {
	dc := gg.NewContext(width, height)
	dc.SetRGB(0, 0, 0)
	dc.DrawRectangle(0, 0, float64(width), float64(height))
	dc.Fill()
	return drawCircle(float64(width)/2, float64(height)/2, rad, dc, c)

}

func SimpleFractal(c echo.Context) error {

	c.Response().Header().Add("Content-Type", "multipart/x-mixed-replace;boundary="+boundaryWord)
	log.Println("new connection")

	return draw(c)

}
