package validator

import (
	"fmt"
	"log"				// logging errors
	"os"				// create and open files
	"bufio"				// scanning files
	"KSCD/parser"		// parser.go
//	"KSCD/mapper"		// mapper.go
)

func Validator() {

	/*** OPEN THE terraform.drawio FILE ***/

	inFile, errRead := os.Open("terraform.drawio")
	// error reading file
	if errRead != nil {
		log.Println("Error opening file.", errRead)
		os.Exit(1)
	}
	// keep file open
	defer inFile.Close()

	/*** SCAN THE terraform.drawio FILE (line by line) ***/

	scanner := bufio.NewScanner(inFile)
	// error scanning file
	errScan := scanner.Err()
	if errScan != nil {
		log.Println("Error scanning file.", errRead)
		os.Exit(1)
	}

	// iterate through lines of file
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	fmt.Println()

	/* ITERATE THROUGH RESOURCES -> INSTANCES -> DEPENDENCIES */

	for r := 0; r < len(parser.T.Resources); r++ {
		for i := 0; i < len(parser.T.Resources[r].Instances); i++ {
			for d := 0; d < len(parser.T.Resources[r].Instances[i].Dependencies); d++ {
				// prints resource type followed by it's dependencies
				fmt.Println("Resource Type:", parser.T.Resources[r].Type, "Dependency:", parser.T.Resources[r].Instances[i].Dependencies[d])

				// NOTE: prints dependencies as type.name, might need just type ??
			}
		}
	}

//	mapper.Xml.WriteToFile("terraform.drawio")
}
