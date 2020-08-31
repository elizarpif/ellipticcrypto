package ellipticcrypto

import (
	"fmt"
	"math"
	"math/big"
	"reflect"
)

// точка на кривой
type pointField struct {
	X *big.Int
	Y *big.Int
}

func (p pointField) String() string {
	return fmt.Sprintf("(%d, %d)", p.X.Int64(), p.Y.Int64())
}

// эллиптическая кривая над полем порядка p
type ellipticCurveField struct {
	// y^2 = x^3 +ax + b
	a, b *big.Int
	// mod p
	p *big.Int
}

// реализация скалярного умножения nP = P+P+..+P
func (ec ellipticCurveField) Mul(p1 pointField, n int) pointField {
	p2 := p1
	for i := 1; i < n; i++ {
		// fmt.Println(p2)
		p2 = ec.Sum(p1, p2)
	}
	// fmt.Println(p2)
	return p2
}

// реализация скалярного сложения P1+P2
func (ec ellipticCurveField) Sum(p1, p2 pointField) pointField {
	if reflect.DeepEqual(p1, p2) {
		return ec.Double(p1)
	}

	o := pointField{
		X: new(big.Int),
		Y: new(big.Int),
	}

	if reflect.DeepEqual(p1, o) {
		return p2
	}
	if reflect.DeepEqual(o, p2) {
		return p1
	}

	dy := new(big.Int)
	dy = big.NewInt(p1.Y.Int64() - p2.Y.Int64())

	dx := &big.Int{}
	dx = big.NewInt(p1.X.Int64() - p2.X.Int64())

	if dx.Int64() == 0 {
		return pointField{
			X: big.NewInt(0),
			Y: big.NewInt(0),
		}
	}

	m := new(big.Int)
	tmp := new(big.Int)
	m.Mul(dy, tmp.ModInverse(dx, ec.p))
	m.Mod(m, ec.p)

	mInt := m.Int64()

	xRes := big.NewInt(mInt*mInt - p1.X.Int64() - p2.X.Int64())
	xRes = xRes.Mod(xRes, ec.p)

	yRes := big.NewInt(p2.Y.Int64() + mInt*(xRes.Int64()-p2.X.Int64()))
	yRes.Neg(yRes).Mod(yRes, ec.p)

	return pointField{
		X: xRes,
		Y: yRes,
	}
}

// реализация удвоения точки 2P=P+P
func (ec ellipticCurveField) Double(p1 pointField) pointField {
	x1 := p1.X.Int64()

	dy := x1*x1*3 + ec.a.Int64()

	dyModP := new(big.Int)
	dyModP.Mod(big.NewInt(dy), ec.p)

	dx := 2 * p1.Y.Int64()

	dxModP := new(big.Int)
	dxModP.Mod(big.NewInt(dx), ec.p)

	m := new(big.Int)
	m.Mul(dyModP, (&big.Int{}).ModInverse(dxModP, ec.p))
	m.Mod(m, ec.p)

	mInt := m.Int64()

	xRes := big.NewInt(mInt*mInt - 2*x1)
	xRes = xRes.Mod(xRes, ec.p)

	yRes := big.NewInt(m.Int64()*(x1-xRes.Int64()) - p1.Y.Int64())
	yRes = yRes.Mod(yRes, ec.p)

	return pointField{
		X: xRes,
		Y: yRes,
	}
}

// эллиптическая кривая в R2
type ellipticCurve struct {
	// y^2 = x^3 +ax + b
	a, b float64
}

// нахождение y зная x
func (ec ellipticCurve) Y(x float64) float64 {
	res := math.Sqrt(x*x*x + ec.a*x + ec.b)
	return res
}

// некая точка
type point struct {
	X float64
	Y float64
}

func (p point) String() string {
	return fmt.Sprintf("(%.5f,%.5f)", p.X, p.Y)
}

// реализация сложения двух точек на эллиптической кривой
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
