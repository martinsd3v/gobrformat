package gobrformat

import (
	"fmt"

	valid "github.com/martinsd3v/gobrvalid"
)

//CPF responsible for formatter CPF
func CPF(cpf string) (isValid bool, formatted string) {
	isValid, clean := valid.IsCPF(cpf)

	if isValid {
		formatted = fmt.Sprintf("%s.%s.%s-%s", clean[0:3], clean[3:6], clean[6:9], clean[9:11])
	}
	return
}

//CNPJ responsible for formatter CNPJ
func CNPJ(cpf string) (isValid bool, formatted string) {
	isValid, clean := valid.IsCNPJ(cpf)

	if isValid {
		formatted = fmt.Sprintf("%s.%s.%s/%s-%s", clean[0:2], clean[2:5], clean[5:8], clean[8:12], clean[12:14])
	}
	return
}

//CPForCNPJ responsible for formatter CPF or CNPJ
func CPForCNPJ(doc string) (isValid bool, formatted string) {
	isValid, clean := valid.IsCPForCNPJ(doc)

	if isValid {
		if len(clean) == 11 {
			_, formatted = CPF(clean)
		} else {
			_, formatted = CNPJ(clean)
		}
	}

	return
}
