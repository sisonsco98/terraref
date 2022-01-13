package main

import (
	"KSCD/mapper"
	"KSCD/parser"
	"KSCD/validator"
)

func main() {

	// run parser.go
	parser.Parser()

	// run mapper.go
	mapper.Mapper()

	// run validator.go
	validator.Validator()

}
