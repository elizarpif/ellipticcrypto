package ellipticcrypto

import (
	"math/big"
	"reflect"
	"testing"
)

func TestEllipticCurve_SumPoints(t *testing.T) {
	tests := []struct {
		name          string
		P             point
		Q             point
		curve         ellipticCurve
		expectedPoint point
	}{
		{
			name:  "simple sum",
			curve: ellipticCurve{a: -36},
			P:     point{-3, 9},
			Q:     point{-2, 8},
			expectedPoint: point{
				X: 6,
				Y: 0,
			},
		},
		{
			name:  "2P",
			curve: ellipticCurve{a: -36},
			P:     point{-3, 9},
			Q:     point{-3, 9},
			expectedPoint: point{
				X: float64(25) / float64(4),
				Y: -float64(35) / float64(8),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.curve.SumPoints(tt.P, tt.Q)
			if res != tt.expectedPoint {
				t.Fatalf("not expected result %s != %s", res, tt.expectedPoint)
			}
		})
	}
}

func Test_ellipticCurveField_Double(t *testing.T) {
	type fields struct {
		a *big.Int
		b *big.Int
		p *big.Int
	}
	type args struct {
		p1 pointField
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   pointField
	}{
		{
			name: "test double",
			fields: fields{
				a: big.NewInt(2),
				b: big.NewInt(3),
				p: big.NewInt(97),
			},
			args: args{
				p1: pointField{
					X: big.NewInt(3),
					Y: big.NewInt(6),
				},
			},
			want: pointField{
				X: big.NewInt(80),
				Y: big.NewInt(10),
			},
		},
		{
			name: "test double",
			fields: fields{
				a: big.NewInt(2),
				b: big.NewInt(3),
				p: big.NewInt(97),
			},
			args: args{
				p1: pointField{
					X: big.NewInt(3),
					Y: big.NewInt(91),
				},
			},
			want: pointField{
				X: big.NewInt(0),
				Y: big.NewInt(0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ec := ellipticCurveField{
				a: tt.fields.a,
				b: tt.fields.b,
				p: tt.fields.p,
			}
			if got := ec.Double(tt.args.p1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Double() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ellipticCurveField_Sum(t *testing.T) {
	type fields struct {
		a *big.Int
		b *big.Int
		p *big.Int
	}
	type args struct {
		p1 pointField
		p2 pointField
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   pointField
	}{
		{
			name: "check sum with equal x coordinates",
			fields: fields{
				a: big.NewInt(2),
				b: big.NewInt(3),
				p: big.NewInt(97),
			},
			args: args{
				p1: pointField{
					X: big.NewInt(3),
					Y: big.NewInt(91),
				},
				p2: pointField{
					X: big.NewInt(3),
					Y: big.NewInt(6),
				},
			},
			want: pointField{
				X: big.NewInt(0),
				Y: big.NewInt(0),
			},
		},
		{
			name: "check simple sum",
			fields: fields{
				a: big.NewInt(2),
				b: big.NewInt(3),
				p: big.NewInt(97),
			},
			args: args{
				p1: pointField{
					X: big.NewInt(3),
					Y: big.NewInt(6),
				},
				p2: pointField{
					X: big.NewInt(80),
					Y: big.NewInt(10),
				},
			},
			want: pointField{
				X: big.NewInt(80),
				Y: big.NewInt(87),
			},
		},
		{
			name: "check double",
			fields: fields{
				a: big.NewInt(2),
				b: big.NewInt(3),
				p: big.NewInt(97),
			},
			args: args{
				p1: pointField{
					X: big.NewInt(3),
					Y: big.NewInt(6),
				},
				p2: pointField{
					X: big.NewInt(3),
					Y: big.NewInt(6),
				},
			},
			want: pointField{
				X: big.NewInt(80),
				Y: big.NewInt(10),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ec := ellipticCurveField{
				a: tt.fields.a,
				b: tt.fields.b,
				p: tt.fields.p,
			}
			if got := ec.Sum(tt.args.p1, tt.args.p2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ellipticCurveField_Mul(t *testing.T) {
	type fields struct {
		a *big.Int
		b *big.Int
		p *big.Int
	}
	type args struct {
		p1 pointField
		n  int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   pointField
	}{
		{

			name: "check simple sum",
			fields: fields{
				a: big.NewInt(2),
				b: big.NewInt(3),
				p: big.NewInt(97),
			},
			args: args{
				p1: pointField{
					X: big.NewInt(3),
					Y: big.NewInt(6),
				},
				n: 6,
			},
			want: pointField{
				X: big.NewInt(3),
				Y: big.NewInt(6),
			},
		},
		{

			name: "check cofactor",
			fields: fields{
				a: big.NewInt(-1),
				b: big.NewInt(3),
				p: big.NewInt(37),
			},
			args: args{
				p1: pointField{
					X: big.NewInt(2),
					Y: big.NewInt(3),
				},
				n: 6,
			},
			want: pointField{
				X: big.NewInt(2),
				Y: big.NewInt(34),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ec := ellipticCurveField{
				a: tt.fields.a,
				b: tt.fields.b,
				p: tt.fields.p,
			}
			if got := ec.Mul(tt.args.p1, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}
