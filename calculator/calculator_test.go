package calculator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSuccessWork(t *testing.T) {
	var testStruct = []struct {
		mes string
		in  string
		out float64
	}{
		{"check plus operation", "1+1", 2},
		{"check multiplication operation", "2*2", 4},
		{"check division operation", "1/2", 0.5},
		{"checking expression with brackets", "(1+2)-3", 0},
		{"checking expression with brackets and float", "(1.1 + 5)*10", 61},
		{"checking big expression", "15/(7-(1+1))*3-(2+(1+1))*15/(7-(200+1))*3-(2+(1+1))*(15/(7-(1+1))*3-(2+(1+1))+15/(7-(1+1))*3-(2+(1+1)))", -30.072166},
	}
	for _, tt := range testStruct {
		t.Run(tt.in, func(t *testing.T) {
			actual, _ := Calculate(tt.in)
			require.Equal(t, tt.out, actual, tt.mes)
		})
	}
}

func TestFalseWork(t *testing.T) {
	var testStruct = []struct {
		mes string
		in  string
		out float64
	}{
		{"check division zero", "1/0", 0},
		{"check no correct symbol", "1-", 0},
		{"check no parametr's", "()", 0},
	}
	for _, tt := range testStruct {
		_, err := Calculate(tt.in)
		require.Error(t, err, tt.mes)
	}
}
