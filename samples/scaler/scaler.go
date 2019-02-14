package scaler

type scaler interface {
	scale(float64)
}

type rect struct {
	a, b int
}

func (r *rect) scale(s float64) {
	r.a = int(float64(r.a) * s)
	r.b = int(float64(r.b) * s)
}

type circle struct {
	r int
}

func (c *circle) scale(s float64) {
	c.r = int(float64(c.r) * s)
}
