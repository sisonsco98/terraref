package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

/*** GLOBAL STRUCT USED TO UNMARSHAL THE .tfstate FILE ***/
var T Terraform

type Terraform struct {
	Resources []struct {
		Type      string `json:"type"`
		Name      string `json:"name"`
		Provider  string `json:"provider"`
		Instances []struct {
			Attributes struct {
				ID      string `json:"id"`
				Name    string `json:"name"`
				Project string `json:"project"`
			} `json:"attributes"`
			Dependencies []string `json:"dependencies"`
		} `json:"instances"`
	} `json:"resources"`
}

/*** GLOBAL SLICES FOR OUTPUTS AND PROVIDERS FROM THE .tfstate FILE ***/
var Outputs []string
var Providers []string

/*** GLOBAL SLICES AND MAPS FOR RELATIONSHIPS ***/
var NameToIndex = make(map[string]int)
var NumDependencies []int
var NumDependents []int

var DependencyNames = make(map[int][]string)
var DependencyIndices = make(map[int][]int)
var DependentNames = make(map[int][]string)
var DependentIndices = make(map[int][]int)

func Parser(inFileLocation string) {
	inFile, errRead := ioutil.ReadFile(inFileLocation)
	if errRead != nil {
		log.Println("Error reading file.", errRead)
		os.Exit(1)
	}

	/*** UNMARSHAL THE .tfstate FILE ***/
	unmarshalFile(inFile)
	getOutputs(inFile, unmarshalFile(inFile))
	getProviders(inFile, unmarshalFile(inFile))

	/*** FIND THE DEPENDENCIES AND DEPENDENTS OF EACH RESOURCE ***/
	fmt.Println()
	fmt.Println("****************************************************************************************************")
	fmt.Println("*            R E S O U R C E    D E P E N D E N C I E S    A N D    D E P E N D E N T S            *")
	fmt.Println("****************************************************************************************************")
	fmt.Println()

	for i := 0; i < len(T.Resources); i++ {
		if T.Resources[i].Name != "default" {
			NameToIndex[T.Resources[i].Name] = i
		}
	}

	NumDependencies = make([]int, len(T.Resources))
	NumDependents = make([]int, len(T.Resources))

	countDependenciesDependents()

	storeDependencies()
	printDependencies()

	storeDependents()
	printDependents()

}

/*** FUNCTIONS ***/

// Unmarshal "outputs" block and "resources" block
func unmarshalFile(inFile []byte) map[string]interface{} {
	var outputBlock map[string]interface{}
	json.Unmarshal(inFile, &outputBlock)
	json.Unmarshal(inFile, &T)
	return outputBlock
}

// Parse the output string to get the outputs
func getOutputs(inFile []byte, outputBlock map[string]interface{}) {
	outputStr := fmt.Sprintln(outputBlock["outputs"])
	outputRegex := regexp.MustCompile(`(([a-z_]*):map)+`)
	output := outputRegex.FindAllStringSubmatch(outputStr, -1)

	for i := range output {
		Outputs = append(Outputs, output[i][2])
	}
}

// Parse provider string to get just the provider within the quotes
func getProviders(inFile []byte, outputBlock map[string]interface{}) {
	providerRegex := regexp.MustCompile(`[^"]*()[^"]*`)

	for i := 0; i < len(T.Resources); i++ {
		provider := providerRegex.FindAllStringSubmatch(T.Resources[i].Provider, -1)
		Providers = append(Providers, provider[1][0])
	}
}

// Counts number of dependencies and depedents for reources
func countDependenciesDependents() {
	for r := 0; r < len(T.Resources); r++ {
		for i := 0; i < len(T.Resources[r].Instances); i++ {
			for d := 0; d < len(T.Resources[r].Instances[i].Dependencies); d++ {
				dependency := T.Resources[r].Instances[i].Dependencies[d]
				dependencyName := strings.Split(dependency, ".")
				dependencyIndex := NameToIndex[dependencyName[1]]

				NumDependencies[r] += 1
				NumDependents[dependencyIndex] += 1
			}
		}
	}
}

// Stores dependencies into struct
func storeDependencies() {
	for r := 0; r < len(T.Resources); r++ {
		var tempDependencyNames []string
		var tempDependencyIndices []int

		// Triggers only if element has dependencies
		if NumDependencies[r] > 0 {
			for i := 0; i < len(T.Resources[r].Instances); i++ {
				for d := 0; d < len(T.Resources[r].Instances[i].Dependencies); d++ {
					dependency := T.Resources[r].Instances[i].Dependencies[d]
					dependencyName := strings.Split(dependency, ".")
					dependencyIndex := NameToIndex[dependencyName[1]]

					tempDependencyNames = append(tempDependencyNames, dependencyName[0])
					tempDependencyIndices = append(tempDependencyIndices, dependencyIndex)
				}
			}

			DependencyNames[r] = tempDependencyNames
			DependencyIndices[r] = tempDependencyIndices
		}
	}
}

// Prints dependencies to console
func printDependencies() {
	for r := 0; r < len(T.Resources); r++ {
		fmt.Print("(", r, ") has ", NumDependencies[r], " dependencies:")
		for d := 0; d < len(DependencyIndices[r]); d++ {
			fmt.Print(" (", (DependencyIndices[r])[d], " ", (DependencyNames[r])[d], ")")
		}
		fmt.Println()
	}
	fmt.Println()
}

// Stores dependents into struct
func storeDependents() {
	for r := 0; r < len(T.Resources); r++ {
		var tempDependentNames []string
		var tempDependentIndices []int

		// Triggers if an element is a dependent
		if NumDependents[r] > 0 {
			rName := T.Resources[r].Name
			for resource := 0; resource < len(T.Resources); resource++ {
				resourceName := T.Resources[resource].Instances[0].Attributes.Name
				for i := 0; i < len(T.Resources[resource].Instances); i++ {
					for d := 0; d < len(T.Resources[resource].Instances[i].Dependencies); d++ {
						if len(T.Resources[resource].Instances[i].Dependencies) > 0 {

							dependency := T.Resources[resource].Instances[i].Dependencies[d]
							dependencyName := strings.Split(dependency, ".")
							if rName == dependencyName[1] {

								tempDependentNames = append(tempDependentNames, resourceName)
								tempDependentIndices = append(tempDependentIndices, resource)
							}
						}
					}
				}
			}

			DependentNames[r] = tempDependentNames
			DependentIndices[r] = tempDependentIndices
		}
	}
}

// Prints dependents to console
func printDependents() {
	for r := 0; r < len(T.Resources); r++ {
		fmt.Print("(", r, ") has ", NumDependents[r], " dependents:")
		for d := 0; d < len(DependentIndices[r]); d++ {
			fmt.Print(" (", (DependentIndices[r])[d], " ", (DependentNames[r])[d], ")")
		}
		fmt.Println()
	}
}
