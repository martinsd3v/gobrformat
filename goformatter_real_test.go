package gobrformat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRealFromNumber(t *testing.T) {
	t.Run("Should be able to fomatted a cnpj", func(t *testing.T) {
		t.Helper()

		for _, tt := range []struct {
			test          interface{}
			precision     int
			expectValue   string
			expectIsValid bool
		}{
			{
				test:          nil,
				precision:     0,
				expectValue:   "",
				expectIsValid: false,
			},
			{
				test:          uint16(5.),
				precision:     2,
				expectValue:   "5,00",
				expectIsValid: true,
			},
			{
				test:          float32(.1),
				precision:     2,
				expectValue:   "0,10",
				expectIsValid: true,
			},
			{
				test:          int(1),
				precision:     2,
				expectValue:   "1,00",
				expectIsValid: true,
			},
			{
				test:          string("10"),
				precision:     2,
				expectValue:   "10,00",
				expectIsValid: true,
			},
			{
				test:          100.532,
				precision:     2,
				expectValue:   "100,53",
				expectIsValid: true,
			},
			{
				test:          1000.3656,
				precision:     3,
				expectValue:   "1.000,366",
				expectIsValid: true,
			},
			{
				test:          1000000.56,
				precision:     2,
				expectValue:   "1.000.000,56",
				expectIsValid: true,
			},
			{
				test:          "1000000.000,56",
				precision:     2,
				expectValue:   "1.000.000.000,56",
				expectIsValid: true,
			},
			{
				test:          "84000000000.000,56",
				precision:     2,
				expectValue:   "84.000.000.000.000,56",
				expectIsValid: true,
			},
		} {
			formatted, isValid := RealFromNumber(tt.test, tt.precision)
			assert.EqualValues(t, tt.expectValue, formatted)
			assert.Equal(t, tt.expectIsValid, isValid)
		}

	})
}
