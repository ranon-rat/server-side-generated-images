package controllers

import (
	"github.com/fogleman/gg"
	"github.com/labstack/echo"
)

func LorenzSystem(wr echo.Context) error {

	wr.Response().Header().Add("Content-Type", "multipart/x-mixed-replace;boundary="+boundaryWord)
	dc := gg.NewContext(S, S)
	x, y, z := 0.1, 0.0, 0.0
	a, b, c := 14.0, 45.0, float64(8/3)
	dx, dy, dz, dt := 0.0, 0.0, 0.0, 0.01
	floatSize := float64(S)
	for i := 0; i < 1000000; i++ {
		//=======MATHS=====\\
		dx, dy, dz = (a*(y-x))*dt, (x*(b-z)-y)*dt, (x*y-c*z)*dt
		x, y, z = x+dx, y+dy, z+dz
		dc.SetRGB255(int(z*255), int(y*255), 255)
		dc.DrawCircle((z * floatSize / 70), (y*floatSize/70)+floatSize/2, 5)
		dc.Fill()
		if _, err := wr.Response().Write([]byte(headerf)); err != nil {
			return err
		}
		if err := dc.EncodePNG(wr.Response()); err != nil {
			return err
		}

	}
	return nil
}
