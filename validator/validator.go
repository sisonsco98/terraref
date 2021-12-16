package validator

import (
	"KSCD/mapper"
	"KSCD/parser" // parser.go
	"bufio"       // scanning files
	"fmt"
	"log" // logging errors
	"os"  // create and open files
	"strings" //Needed to use split
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
				//fmt.Println("Resource Type:", parser.T.Resources[r].Type, "Dependency:", parser.T.Resources[r].Instances[i].Dependencies[d])

				resourceName := parser.T.Resources[r].Instances[i].Dependencies[d]
				dependencyName := strings.Split(resourceName, ".")

				fmt.Println("Parent Resource Name : ", mapper.Pizza[r].Name)
				fmt.Println("Dependency Name : ", dependencyName[1])

				ctr := 0
				for range mapper.Pizza {
					if (mapper.Pizza[ctr].Name == dependencyName[1]){
						fmt.Println("We've matched the elements.")
						fmt.Println("We need to draw an arrow from element ", mapper.Pizza[r].Name, " to element ", mapper.Pizza[ctr].Name)
						fmt.Println(mapper.Pizza[r].Name, " is located at (", mapper.Pizza[r].XPos, ",", mapper.Pizza[r].YPos, ")")
						fmt.Println(mapper.Pizza[ctr].Name, " is located at (", mapper.Pizza[ctr].XPos, ",", mapper.Pizza[ctr].YPos , ")")

					}


					ctr++
				}

				// NOTE: prints dependencies as type.name, might need just type ??
			}
		}
	}
}
