package engine3D

type Surface struct {
	A Point
	B Point
	C Point
	D Point
}

func (s *Surface) translate(vector Point) {
	s.A.Translate(vector)
	s.B.Translate(vector)
	s.C.Translate(vector)
	s.D.Translate(vector)
}
