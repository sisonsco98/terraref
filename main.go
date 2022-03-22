package main

import (
	"KSCD/mapper"
	"KSCD/parser"
	"KSCD/validator"
	"flag"
)

func main() {

	fileLoc := flag.String("statefile", "terraform.tfstate", "Filename that we're parsing from." )

	flag.Parse()

	var fileLocation = *fileLoc

	// run parser.go
	parser.Parser(fileLocation)

	// run mapper.go
	mapper.Mapper()

	// run validator.go
	validator.Validator()

}
