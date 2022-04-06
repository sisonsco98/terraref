package parser

import (
	"fmt"
	"io/ioutil"		// read files
	"log"			// logging errors
	"os"			// create and open files
	"encoding/json"	// json.Unmarshal
	"regexp"		// matching regex
	"strings"
)

/*** GLOBAL SLICES FOR OUTPUTS AND PROVIDERS FROM THE .tfstate FILE ***/

var Outputs []string
var Providers []string

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

/*** ??? ***/

// map each resource name to resource index
var NameToIndex = make(map[string]int)

// number of dependents and dependencies for each resource
var NumDependents []int
var NumDependencies []int

// list of dependencies and dependents for each resource
var DependencyNames = make(map[int][]string)
var DependencyIndices = make(map[int][]int)
var DependentNames = make(map[int][]string)
var DependentIndices = make(map[int][]int)

func Parser(inFileLocation string) {

	/*** READ IN THE .tfstate FILE (entire file) ***/

	inFile, errRead := ioutil.ReadFile(inFileLocation)
	// error reading file
	if errRead != nil {
		log.Println("Error reading file.", errRead)
		os.Exit(1)
	}

	/*** UNMARSHAL THE .tfstate FILE ***/

	// unmarshal "outputs" block and "resources" block
	var outputBlock map[string]interface{}
	json.Unmarshal(inFile, &outputBlock)
	json.Unmarshal(inFile, &T)

	// parse the output string to get the outputs
	// ex: map[ab:map[type:string value:34.105.77.168] ip:map[type:string value:34.105.77.168]]
	outputStr := fmt.Sprintln(outputBlock["outputs"])
	outputRegex := regexp.MustCompile(`(([a-z_]*):map)+`)
	output := outputRegex.FindAllStringSubmatch(outputStr, -1)

	// iterates through matches and stores each output in Outputs[] slice
	for i := range output {
		Outputs = append(Outputs, output[i][2])
	}

	// parse provider string to get just the provider within the quotes
	// ex. provider["registry.terraform.io/hashicorp/google"]
	providerRegex := regexp.MustCompile(`[^"]*()[^"]*`)

	// iterates through matches and stores each output in Providers[] slice
	for i := 0; i < len(T.Resources); i++ {
		provider := providerRegex.FindAllStringSubmatch(T.Resources[i].Provider, -1)
		Providers = append(Providers, provider[1][0])
	}

	/*** COUNT AND STORE DEPENDENCIES AND DEPENDENTS ***/
	/*** DETERMINE THE DEPENDENCIES AND DEPENDENTS OF EACH RESOURCE ***/

	fmt.Println()
	fmt.Println("****************************************************************************************************")
	fmt.Println("*            R E S O U R C E    D E P E N D E N C I E S    A N D    D E P E N D E N T S            *")
	fmt.Println("****************************************************************************************************")
	fmt.Println()

	// number of dependents and dependencies for each resource
	NumDependents = make([]int, len(T.Resources) * len(T.Resources))
	NumDependencies = make([]int, len(T.Resources) * len(T.Resources))

	// store each resource name and resource index
	for i := 0; i < len(T.Resources); i++ {
		if T.Resources[i].Name != "default" {
			NameToIndex[T.Resources[i].Name] = i
		}
	}

	// iterate through each resource -> instance -> dependency to count NumDependencies and NumDependents
	for r := 0; r < len(T.Resources); r++ {
		for i := 0; i < len(T.Resources[r].Instances); i++ {
			for d := 0; d < len(T.Resources[r].Instances[i].Dependencies); d++ {
				// save dependency info
				dependency := T.Resources[r].Instances[i].Dependencies[d]
				dependencyName := strings.Split(dependency, ".")
				dependencyIndex := NameToIndex[dependencyName[1]]
				// increment NumDependencies and NumDependents
				NumDependencies[r] += 1
				NumDependents[dependencyIndex] += 1
			}
		}
	}

	// iterate through each resource
	for r := 0; r < len(T.Resources); r++ {

		// temp list of dependencies and dependents for current resource
		var tempDependencyNames []string
		var tempDependencyIndices []int
		var tempDependentNames []string
		var tempDependentIndices []int

		// find the name and index of each dependency of the current resource
		if NumDependencies[r] > 0 {
			for i := 0; i < len(T.Resources[r].Instances); i++ {
				for d := 0; d < len(T.Resources[r].Instances[i].Dependencies); d++ {
					// save dependency info
					dependency := T.Resources[r].Instances[i].Dependencies[d]
					dependencyName := strings.Split(dependency, ".")
					dependencyIndex := NameToIndex[dependencyName[1]]
					// append dependency
					tempDependencyNames = append(tempDependencyNames, dependencyName[0])
					tempDependencyIndices = append(tempDependencyIndices, dependencyIndex)
				}
			}
			// store dependencies for current resource
			DependencyNames[r] = tempDependencyNames
			DependencyIndices[r] = tempDependencyIndices
		}

		// find the name and index of each dependent of the current resource
		if NumDependents[r] > 0 {
			rName := T.Resources[r].Name
			for resource := 0; resource < len(T.Resources); resource++ {
				resourceName := T.Resources[resource].Instances[0].Attributes.Name
				for i := 0; i < len(T.Resources[resource].Instances); i++ {
					for d := 0; d < len(T.Resources[resource].Instances[i].Dependencies); d++ {
						if len(T.Resources[resource].Instances[i].Dependencies) > 0 {
							// save dependent info
							dependency := T.Resources[resource].Instances[i].Dependencies[d]
							dependencyName := strings.Split(dependency, ".")
							if rName == dependencyName[1] {
								// append dependent
								tempDependentNames = append(tempDependentNames, resourceName)
								tempDependentIndices = append(tempDependentIndices, resource)
							}
						}
					}
				}
			}
			// store dependents for resource
			DependentNames[r] = tempDependentNames
			DependentIndices[r] = tempDependentIndices
		}

	}

	/*** PRINT THE DEPENDENCIES AND DEPENDENTS OF EACH RESOURCE ***/

	// dependencies
	for r := 0; r < len(T.Resources); r++ {
		fmt.Print("(", r, ") has ", NumDependencies[r], " dependencies:")
		for d := 0; d < len(DependencyIndices[r]); d++ {
			fmt.Print(" (", (DependencyIndices[r])[d], " ", (DependencyNames[r])[d], ")")
		}
		fmt.Println()
	}
	fmt.Println()

	// dependents
	for r := 0; r < len(T.Resources); r++ {
		fmt.Print("(", r, ") has ", NumDependents[r], " dependents:")
		for d := 0; d < len(DependentIndices[r]); d++ {
			fmt.Print(" (", (DependentIndices[r])[d], " ", (DependentNames[r])[d], ")")
		}
		fmt.Println()
	}

}
