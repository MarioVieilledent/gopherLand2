package engine3D

type Camera struct {
	Name  string
	Focal float64 // Distance between camera and screen
	Angle float64 // Angle on Y axis
	Pos   Point   // Position
}

func (c *Camera) translate(vector Point) {
	c.Pos.Translate(vector)
}
