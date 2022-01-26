package validator

import (
	"bufio"		// scanning files
	"log"		// logging errors
	"os"		// create and open files
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
//	for scanner.Scan() {
//		fmt.Println(scanner.Text())
//	}

	/* ITERATE THROUGH RESOURCES -> INSTANCES -> DEPENDENCIES */

}