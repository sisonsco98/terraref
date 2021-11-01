package mapper

import (
	"fmt"
	"log"	// logging errors
	"os"	// create and open file
	"KSCD/parser"
)

func Mapper() {

	/*** CREATE THE terraform.drawio FILE ***/

	outFile, errCreate := os.Create("terraform.drawio")
	// error creating file
	if errCreate != nil {
		log.Println("Error creating file.", errCreate)
		os.Exit(1)
	}
	// keep open
	defer outFile.Close()

	/*** PRINT PARSED DATA TO THE terraform.drawio FILE ***/

	// iterates through matches and prints each output
	for i := range parser.Outputs {
		fmt.Fprintln(outFile, "Outputs:", parser.Outputs[i])
	}
	fmt.Fprintln(outFile)

	fmt.Fprintln(outFile, "Resources:")
	fmt.Fprintln(outFile)
	
	// iterate through the resources and prints
	for i := 0; i < len(parser.T.Resources); i++ {
		fmt.Fprintln(outFile, "\t", "Type:", parser.T.Resources[i].Type)
		fmt.Fprintln(outFile, "\t", "Name:", parser.T.Resources[i].Name)
		fmt.Fprintln(outFile, "\t", "Provider:", parser.T.Resources[i].Provider)

		fmt.Fprintln(outFile, "\t", "Instances:")
		fmt.Fprintln(outFile, "\t\t", "Attributes:")

		// iterate through the instances and prints
		for j := 0; j < len(parser.T.Resources[i].Instances); j++ {
			if len(parser.T.Resources[i].Instances[j].Attributes.ID) > 0 {
				fmt.Fprintln(outFile, "\t\t\t", "ID:", parser.T.Resources[i].Instances[j].Attributes.ID)
			}
			if len(parser.T.Resources[i].Instances[j].Attributes.Name) > 0 {
				fmt.Fprintln(outFile, "\t\t\t", "Name:", parser.T.Resources[i].Instances[j].Attributes.Name)
			}
			if len(parser.T.Resources[i].Instances[j].Attributes.Project) > 0 {
				fmt.Fprintln(outFile, "\t\t\t", "Project:", parser.T.Resources[i].Instances[j].Attributes.Project)
			}

			// iterate through dependencies and prints
			for k := 0; k < len(parser.T.Resources[i].Instances[j].Dependencies); k++ {
				fmt.Fprintln(outFile, "\t\t", "Dependencies:", parser.T.Resources[i].Instances[j].Dependencies[k])
			}
		}
	}
	
	// close file
	outFile.Close()
}
