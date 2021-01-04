package gobrformat

import (
	"fmt"

	valid "github.com/martinsd3v/gobrvalid"
)

//DateBR responsible for formatter dates in Brazil pattern
func DateBR(date string, hour bool) (formatted string, isValid bool) {
	isValid, dt := valid.IsDate(date)

	if isValid {
		formatted = fmt.Sprintf("%02d/%02d/%04d", dt.Day(), dt.Month(), dt.Year())

		if hour {
			formatted = fmt.Sprintf("%s %02d:%02d:%02d", formatted, dt.Hour(), dt.Minute(), dt.Second())
		}
	}
	return
}

//DateDB responsible for formatter dates in database pattern
func DateDB(date string, hour bool) (formatted string, isValid bool) {
	isValid, dt := valid.IsDate(date)

	if isValid {
		formatted = fmt.Sprintf("%04d-%02d-%02d", dt.Year(), dt.Month(), dt.Day())

		if hour {
			formatted = fmt.Sprintf("%s %02d:%02d:%02d", formatted, dt.Hour(), dt.Minute(), dt.Second())
		}
	}
	return
}
