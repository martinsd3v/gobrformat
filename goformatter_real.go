package formatter

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	valid "github.com/martinsd3v/gobrvalid"
)

//RealFromNumber responsible for formatter number to Real patter
func RealFromNumber(value interface{}, precision int) (formated string, isValid bool) {
	v := reflect.ValueOf(value)
	isValid = true

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		formated = fmt.Sprintf("%d,", v.Int()) + strings.Repeat("0", precision)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		formated = fmt.Sprintf("%d,", v.Uint()) + strings.Repeat("0", precision)
	case reflect.Float32, reflect.Float64:
		formated = fmt.Sprintf(fmt.Sprintf("%%.%df", precision), v.Float())
	case reflect.String:
		clean := valid.ReplacePattern(v.String(), "[^0-9,.]", "")

		if valid.Contains(clean, ",") && valid.Contains(clean, ".") {
			clean = valid.ReplacePattern(clean, "[.]", "")
			clean = valid.ReplacePattern(clean, "[,]", ".")
		}

		if valid.Contains(clean, ".") {
			value, _ := strconv.ParseFloat(clean, 64)
			formated = fmt.Sprintf(fmt.Sprintf("%%.%df", precision), value)
		} else {
			formated = fmt.Sprintf("%s,", clean) + strings.Repeat("0", precision)
		}
	default:
		isValid = false
	}

	if isValid {
		//Change '.' per ','
		formated = valid.ReplacePattern(formated, "[.]", ",")

		//remove decimal numbers from string
		lenFormated := len(formated)
		decimal := formated[lenFormated-(precision+1) : lenFormated]
		r := valid.Reverse(formated[0 : lenFormated-(precision+1)])

		switch l := len(r); {
		case l > 12:
			formated = fmt.Sprintf("%s.%s.%s.%s.%s", r[0:3], r[3:6], r[6:9], r[9:12], r[12:l])
		case l > 9:
			formated = fmt.Sprintf("%s.%s.%s.%s", r[0:3], r[3:6], r[6:9], r[9:l])
		case l > 6:
			formated = fmt.Sprintf("%s.%s.%s", r[0:3], r[3:6], r[6:l])
		case l > 3:
			formated = fmt.Sprintf("%s.%s", r[0:3], r[3:l])
		default:
			formated = r
		}

		formated = valid.Reverse(formated) + decimal
	}

	return
}
