package parser

import (
	"fmt"
	"io/ioutil"		// read files
	"log"			// logging errors
	"os"			// create and open files
	"encoding/json"	// json.Unmarshal
	"regexp"		// matching regex
)

/*** GLOBAL SLICES FOR OUTPUTS AND PROVIDERS FROM THE .tfstate FILE ***/

var Outputs []string
var Providers []string
var Elements []string

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

func Parser(filename string) {

	/*** READ IN THE .tfstate FILE (entire file) ***/

	inFile, errRead := ioutil.ReadFile(filename)
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

	// parser provider string to get just the provider within the quotes
	// ex. provider["registry.terraform.io/hashicorp/google"]
	providerRegex := regexp.MustCompile(`[^"]*()[^"]*`)

	// iterates through matches and stores each output in Providers[] slice
	for i := 0; i < len(T.Resources); i++ {
		provider := providerRegex.FindAllStringSubmatch(T.Resources[i].Provider, -1)
		Providers = append(Providers, provider[1][0])
	}

}
