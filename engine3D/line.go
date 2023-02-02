package engine3D

type Line struct {
	Start Point
	End   Point
}

func (l *Line) Translate(vector Point) {
	l.Start.Translate(vector)
	l.End.Translate(vector)
}

func (l Line) Rotate(angle float64) Line {
	return Line{
		Start: l.Start.Rotate(angle),
		End:   l.End.Rotate(angle),
	}
}
