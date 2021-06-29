package controllers

import (
	"math/cmplx"

	"github.com/fogleman/gg"
	"github.com/labstack/echo"
)

const (
	maxIteration = 30
)

func scale(in, inMin, inMax int, outMin, outMax float64) float64 {
	var n float64 = float64(in-inMin) / float64(inMax-inMin)
	var out float64 = n*(outMax-outMin) + outMin
	return out
}
func fractal(img *gg.Context, wr echo.Context) error {
	for px := 0; px < width; px++ {
		for py := 0; py < height; py++ {

			cx, cy := scale(px, 0, width, -2.511, 1), scale(py, 0, height, -1, 1)
			var c = complex(cx, cy)
			z, i := 0+0i, 0

			for cmplx.Abs(z) < 2 && i < maxIteration {
				z = z*z + c
				i++
			}
			if py%10 == 0 && px%10 == 0 {
				img.DrawRectangle(float64(px), float64(py), 10, 10)
				img.SetRGB(float64(i%64)*64, float64(i%8)*32, float64(i%64)*64)

				img.Fill()
				if _, err := wr.Response().Write([]byte(headerf)); err != nil {
					return err
				}
				if err := img.EncodePNG(wr.Response()); err != nil {
					return err
				}
				wr.Response().Flush()
			}
		}
	}
	return nil
}
func MandelbrotSet(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	c.Response().Header().Add("Content-Type", "multipart/x-mixed-replace;boundary="+boundaryWord)
	dc := gg.NewContext(width, height)

	return fractal(dc, c)

}
