package main

import (
	"KSCD/mapper"
	"KSCD/parser"
	"KSCD/validator"
	"flag"
)

func main() {
	// Format -> Flag format, --argument, default value, help message.
	fileLoc := flag.String("statefile", "terraform.tfstate", "Filename that we're parsing from." )

	// Parse the flag arguments~!
	flag.Parse()

	// Return value fileLoc is actually the address of a string variable with the value I'm looking for.
	var fileLocation = *fileLoc

	// run parser.go
	parser.Parser(fileLocation)

	// run mapper.go
	mapper.Mapper()

	// run validator.go
	validator.Validator()

}
