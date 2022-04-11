package main

import (
	"flag"

	"KSCD/parser"
	"KSCD/mapper"
	"KSCD/validator"
)

func main() {

	// Format -> Flag format, --argument, default value, help message.
	inFileLoc := flag.String("in", "inputs/terraform.tfstate", "Filename that we're parsing from. This should be a .tfstate file.")
	outfileLoc := flag.String("out", "outputs/out.drawio", "Destination for draw.io output file.")

	// parse the flag arguments
	flag.Parse()

	// Return value inFileLoc is actually the address of a string variable with the value I'm looking for.
	var inFileLocation = *inFileLoc
	var outFileLocation = *outfileLoc

	// run parser.go
	parser.Parser(inFileLocation)

	// run mapper.go
	mapper.Mapper(outFileLocation)

	// run validator.go
	validator.Validator(outFileLocation)

}
