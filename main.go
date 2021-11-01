package main

import (
	"fmt"
	"KSCD/parser"
//	"KSCD/mapper"
//	"KSCD/validator"
)

func main() {

	// run parser.go
	parser.Parser()

	// iterates through matches and prints each output
	for i := range parser.Outputs {
		fmt.Println("Outputs:", parser.Outputs[i])
	}
	
	// iterate through the resources and prints
	for i := 0; i < len(parser.T.Resources); i++ {
		fmt.Println("\t", "Type:", parser.T.Resources[i].Type)
		fmt.Println("\t", "Name:", parser.T.Resources[i].Name)
		fmt.Println("\t", "Provider:", parser.T.Resources[i].Provider)

		fmt.Println("\t", "Instances:")
		fmt.Println("\t\t", "Attributes:")

		// iterate through the instances and prints
		for j := 0; j < len(parser.T.Resources[i].Instances); j++ {
			if len(parser.T.Resources[i].Instances[j].Attributes.ID) > 0 {
				fmt.Println("\t\t\t", "ID:", parser.T.Resources[i].Instances[j].Attributes.ID)
			}
			if len(parser.T.Resources[i].Instances[j].Attributes.Name) > 0 {
				fmt.Println("\t\t\t", "Name:", parser.T.Resources[i].Instances[j].Attributes.Name)
			}
			if len(parser.T.Resources[i].Instances[j].Attributes.Project) > 0 {
				fmt.Println("\t\t\t", "Project:", parser.T.Resources[i].Instances[j].Attributes.Project)
			}

			// iterate through dependencies and prints
			for k := 0; k < len(parser.T.Resources[i].Instances[j].Dependencies); k++ {
				fmt.Println("\t\t", "Dependencies:", parser.T.Resources[i].Instances[j].Dependencies[k])
			}
		}
	}

}
