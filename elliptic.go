package ellipticcrypto

import (
	"fmt"
	"math"
)

type ellipticCurve struct {
	// y^2 = x^3 +ax + b
	a, b float64
}

type point struct {
	X float64
	Y float64
}

func (p point) String() string {
	return fmt.Sprintf("(%.5f,%.5f)", p.X, p.Y)
}

func (ec ellipticCurve) SumPoints(p1 point, p2 point) point {
	if p1.Y == 0 {
		return p2
	}
	if p2.Y == 0 {
		return p1
	}

	var o point
	if p1.Y == p2.Y && p1.X == (-1)*p2.X {
		return o
	}

	if p1.X == p2.X {
		phi := (3*p1.X*p1.X + ec.a) / (2 * p1.Y)
		x3 := math.Pow(float64(phi), 2) - 2*p1.X
		return point{
			X: x3,
			Y: (-1)*p1.Y + phi*(p1.X-x3),
		}
	}

	alpha := (p2.Y - p1.Y) / (p2.X - p1.X)
	x3 := alpha*alpha - p2.X - p1.X

	res := point{
		X: x3,
		Y: (-1)*p1.Y + alpha*(p1.X-x3),
	}
	return res
}
