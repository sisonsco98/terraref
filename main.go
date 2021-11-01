package main

import (
//	"fmt"
	"KSCD/parser"
	"KSCD/mapper"
//	"KSCD/validator"
)

func main() {

	// run parser.go
	parser.Parser()

	// run mapper.go
	mapper.Mapper()

}
