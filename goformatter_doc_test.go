package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCPF(t *testing.T) {
	t.Run("Should be able to fomatted a cpf", func(t *testing.T) {
		var (
			cpfOriginal string
			cpfExpected string
			cpfFomatted string
			isValid     bool
		)

		cpfOriginal = "51890317055"
		cpfExpected = "518.903.170-55"
		isValid, cpfFomatted = CPF(cpfOriginal)

		assert.True(t, isValid)
		assert.EqualValues(t, cpfExpected, cpfFomatted)

		cpfOriginal = "51890317066"
		isValid, _ = CPF(cpfOriginal)

		assert.False(t, isValid)
	})
}

func TestCNPJ(t *testing.T) {
	t.Run("Should be able to fomatted a cnpj", func(t *testing.T) {
		var (
			cnpjOriginal string
			cnpjExpected string
			cnpjFomatted string
			isValid      bool
		)

		cnpjOriginal = "33.992998/000107"
		cnpjExpected = "33.992.998/0001-07"
		isValid, cnpjFomatted = CNPJ(cnpjOriginal)

		assert.True(t, isValid)
		assert.EqualValues(t, cnpjExpected, cnpjFomatted)

		cnpjOriginal = "33.992.998/000104"
		isValid, _ = CNPJ(cnpjOriginal)

		assert.False(t, isValid)
	})
}

func TestCPForCNPJ(t *testing.T) {
	t.Run("Should be able to fomatted a cpf", func(t *testing.T) {
		var (
			docOriginal string
			docExpected string
			docFomatted string
			isValid     bool
		)

		docOriginal = "51890317055"
		docExpected = "518.903.170-55"
		isValid, docFomatted = CPForCNPJ(docOriginal)

		assert.True(t, isValid)
		assert.EqualValues(t, docExpected, docFomatted)

		docOriginal = "33.992.998/000107"
		docExpected = "33.992.998/0001-07"
		isValid, docFomatted = CPForCNPJ(docOriginal)

		assert.True(t, isValid)
		assert.EqualValues(t, docExpected, docFomatted)

		docOriginal = "51890317044"
		isValid, _ = CPForCNPJ(docOriginal)

		assert.False(t, isValid)
	})
}
