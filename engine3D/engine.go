package engine3D

import "math"

type Engine3D struct {
	Cam            Camera
	Scale          float64
	Objs           []Line
	ScreenSettings ScreenSettings
}

type ScreenSettings struct {
	Width      int
	Height     int
	SemiWidth  int
	SemiHeight int
}

// Create ScreenSettings (width and height)
func CreateScreenSettings(w, h int) ScreenSettings {
	return ScreenSettings{
		Width:      w,
		Height:     h,
		SemiWidth:  w / 2, // Used to compute center of screen, faster
		SemiHeight: h / 2, // Used to compute center of screen, faster
	}
}

// Take a 3D point and plot it to a screen
func (e Engine3D) PlotPoint(point Point) (x, y float64) {
	// Mega danger ici :
	// x = e.plotPointDerivative(point, "x")
	// y = e.plotPointDerivative(point, "y")

	// Ca marche :
	x = float64(e.ScreenSettings.SemiWidth) + e.Scale*(e.Cam.Focal*point.X/(e.Cam.Focal+point.Z))
	y = float64(e.ScreenSettings.SemiHeight) + e.Scale*(e.Cam.Focal*point.Y/(e.Cam.Focal+point.Z))
	return
}

// Demerde toi
func (e Engine3D) plotPointDerivative(point Point, derivative string) float64 {
	// Point at focal length
	var d Point
	var alpha float64
	var beta float64
	if derivative == "x" {
		d = Point{
			e.Cam.Pos.X + math.Sin(e.Cam.Angle)*e.Cam.Focal,
			e.Cam.Pos.Y,
			e.Cam.Pos.Z,
		}

		// normN := math.Sqrt(d.X*d.X + d.Z*d.Z)

		beta = (d.X - point.Z) / (d.X*d.X + d.Z + d.Z)
		alpha = (point.X + d.Z*beta) / d.X
		return float64(e.ScreenSettings.SemiWidth) + e.Scale*(beta*e.Cam.Focal/alpha)
	} else if derivative == "y" {
		d = Point{
			e.Cam.Pos.X + math.Cos(e.Cam.Angle)*e.Cam.Focal,
			e.Cam.Pos.Y,
			e.Cam.Pos.Z + math.Sin(e.Cam.Angle)*e.Cam.Focal,
		}

		// normN := math.Sqrt(d.X*d.X + d.Z*d.Z)

		beta = (d.Y - point.Z) / (d.X*d.Y + d.Z + d.Z)
		alpha = (point.Y + d.Z*beta) / d.Y
		return float64(e.ScreenSettings.SemiHeight) + e.Scale*(beta*e.Cam.Focal/alpha)
	} else {
		return 0
	}
}

func norm(p Point) float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y + p.Z*p.Z)
}
