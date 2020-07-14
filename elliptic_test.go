package ellipticcrypto

import (
	"testing"
)

func TestEllipticCurve_SumPoints(t *testing.T) {
	tests := []struct {
		name string
		P             point
		Q             point
		curve         ellipticCurve
		expectedPoint point
	}{
		{
			name: "simple sum",
			curve: ellipticCurve{a: -36},
			P: point{-3, 9},
			Q: point{-2, 8},
			expectedPoint: point{
				X: 6,
				Y: 0,
			},
		},
		{
			name: "2P",
			curve: ellipticCurve{a: -36},
			P: point{-3, 9},
			Q: point{-3, 9},
			expectedPoint: point{
				X: float64(25)/float64(4),
				Y: -float64(35)/float64(8),
			},
		},
	}
	for _, tt := range tests{
		t.Run(tt.name, func(t *testing.T) {
			res := tt.curve.SumPoints(tt.P, tt.Q)
			if res != tt.expectedPoint{
				t.Fatalf("not expected result %s != %s", res, tt.expectedPoint)
			}
		})
	}
}
