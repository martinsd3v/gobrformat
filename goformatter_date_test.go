package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDateBR(t *testing.T) {
	t.Run("Should be able to fomatted a date", func(t *testing.T) {
		var (
			dateOriginal string
			dateExpected string
			dateFomatted string
			isValid      bool
		)

		dateOriginal = "2020-12-28 15:07:29.590049 -0300"
		dateExpected = "28/12/2020 15:07:29"
		dateFomatted, isValid = DateBR(dateOriginal, true)

		assert.True(t, isValid)
		assert.EqualValues(t, dateExpected, dateFomatted)

		dateOriginal = "2020-5-10"
		dateExpected = "10/05/2020"
		dateFomatted, isValid = DateBR(dateOriginal, false)

		assert.True(t, isValid)
		assert.EqualValues(t, dateExpected, dateFomatted)

		dateOriginal = "2020-10-1 23:55:00"
		dateExpected = "01/10/2020 23:55:00"
		dateFomatted, isValid = DateBR(dateOriginal, true)

		assert.True(t, isValid)
		assert.EqualValues(t, dateExpected, dateFomatted)

		dateOriginal = "2020-15-1 25:55:00"
		_, isValid = DateBR(dateOriginal, true)

		assert.False(t, isValid)

	})
}

func TestDateBD(t *testing.T) {
	t.Run("Should be able to fomatted a date", func(t *testing.T) {
		var (
			dateOriginal string
			dateExpected string
			dateFomatted string
			isValid      bool
		)

		dateOriginal = "2020-5-1"
		dateExpected = "2020-05-01"
		dateFomatted, isValid = DateDB(dateOriginal, false)

		assert.True(t, isValid)
		assert.EqualValues(t, dateExpected, dateFomatted)

		dateOriginal = "01/10/2020 23:55:00"
		dateExpected = "2020-10-01 23:55:00"
		dateFomatted, isValid = DateDB(dateOriginal, true)

		assert.True(t, isValid)
		assert.EqualValues(t, dateExpected, dateFomatted)

		dateOriginal = "2020-15-1 25:55:00"
		_, isValid = DateDB(dateOriginal, true)

		assert.False(t, isValid)

	})
}
