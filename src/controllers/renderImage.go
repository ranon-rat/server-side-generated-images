package controllers

import (
	"log"
	"math/rand"
	"time"

	//"github.com/fogleman/gg"

	"github.com/fogleman/gg"
	"github.com/labstack/echo"
)

const (
	S            = 600
	boundaryWord = "MJPEGBOUNDARY"
	headerf      = "\r\n" +
		"--" + boundaryWord + "\r\n" +
		"X-Timestamp: 0.000000\r\n" +
		"\r\n"
)

func RenderImage(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	c.Response().Header().Add("Content-Type", "multipart/x-mixed-replace;boundary="+boundaryWord)
	log.Println(echo.MIMEApplicationForm)

	dc := gg.NewContext(600, 600)

	for {

		for i := 0; i < 360; i += 15 {
			dc.SetRGBA(float64(rand.Intn(255)), float64(rand.Intn(255)), float64(rand.Intn(255)), 1)
			dc.Push()
			dc.RotateAbout(gg.Radians(float64(i)), S/2, S/2)
			dc.DrawEllipse(S/2, S/2, S*7/16, S/8)
			dc.Fill()
			dc.Pop()
			time.Sleep(time.Millisecond * 200)
			if _, err := c.Response().Write([]byte(headerf)); err != nil {
				return err
			}
			if err := dc.EncodePNG(c.Response()); err != nil {
				return err
			}
			if _, err := c.Response().Write([]byte("\r")); err != nil {
				return err
			}
			c.Response().Flush()
		}

	}

}
