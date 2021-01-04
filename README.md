# gobrformat

Golang Package for format brazilian documents, currency and dates

#### Installation
Make sure that Go is installed on your computer.
Type the following command in your terminal:
```bash
go get github.com/martinsd3v/gobrformat
```

#### Import package in your project
Add following line in your `*.go` file:
```go
import "github.com/martinsd3v/gobrformat"
```
If you are unhappy to use long `govalidator`, you can do something like this:
```go
import (
  fmt "github.com/martinsd3v/gobrformat"
)
```

#### List of functions:
```go
func CPF(cpf string) (isValid bool, formatted string)
func CNPJ(cpf string) (isValid bool, formatted string)
func CPForCNPJ(doc string) (isValid bool, formatted string)
func RealFromNumber(value interface{}, precision int) (formated string, isValid bool)
func DateBR(date string, hour bool) (formatted string, isValid bool)
func DateDB(date string, hour bool) (formatted string, isValid bool)
```

###### PT-BR

Pacote em golang para formatação de documentos brasileiros, moeda e datas

#### Instalação
Certifique-se de que Go está instalado em seu computador.
Digite o seguinte comando em seu terminal:

```bash
go get github.com/martinsd3v/gobrformat
```

#### Importar pacote em seu projeto
Adicione a seguinte linha em seu arquivo `*.go`:
```go
import "github.com/martinsd3v/gobrformat"
```
Se você não gosta de usar o `gobrformat` longo, pode fazer algo assim:
```go
import (
  fmt "github.com/martinsd3v/gobrformat"
)
```

#### Lista de funções:
```go
func CPF(cpf string) (isValid bool, formatted string)
func CNPJ(cpf string) (isValid bool, formatted string)
func CPForCNPJ(doc string) (isValid bool, formatted string)
func RealFromNumber(value interface{}, precision int) (formated string, isValid bool)
func DateBR(date string, hour bool) (formatted string, isValid bool)
func DateDB(date string, hour bool) (formatted string, isValid bool)
```