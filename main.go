package main

import (
	"KSCD/mapper"
	"KSCD/parser"
	"KSCD/validator"
	"flag"
)

func main() {

	// Format -> Flag format, --argument, default value, help message.
	inFileLoc := flag.String("in", "inputs/shared_vpc.tfstate", "Filename that we're parsing from. This should be a .tfstate file.")
	outfileLoc := flag.String("out", "outputs/out.drawio", "Destination for draw.io output file.")

	flag.Parse()

	var inFileLocation = *inFileLoc
	var outFileLocation = *outfileLoc

	parser.Parser(inFileLocation)

	mapper.Mapper(outFileLocation)

	validator.Validator(outFileLocation)

}
