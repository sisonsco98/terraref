package main

import (
	"flag"

	"KSCD/mapper"
	"KSCD/parser"
	"KSCD/validator"
)

func main() {
	// Format -> Flag format, --argument, default value, help message.
	fileLoc := flag.String("in", "inputs/terraform.tfstate", "Filename that we're parsing from. This should be a .tfstate file.")
	outfileLoc := flag.String("out", "outputs/out.drawio", "Destination for draw.io output file.")
	// Parse the flag arguments~!
	flag.Parse()

	// Return value fileLoc is actually the address of a string variable with the value I'm looking for.
	var fileLocation = *fileLoc
	var outputLocation = *outfileLoc
	// run parser.go
	parser.Parser(fileLocation)

	// run mapper.go
	mapper.Mapper(outputLocation)

	// run validator.go
	validator.Validator(outputLocation)

}
