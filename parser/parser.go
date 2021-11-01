package parser

import (
	"fmt"
	"log"			// logging errors
	"os"			// create and open file
	"io/ioutil"		// read in entire file
	"regexp"		// matching regex
	"encoding/json"	// json.Unmarshal
)

/*** STRUCTS / SLICE USED TO UNMARSHAL THE terraform.tfstate FILE ***/

var T Terraform
var Outputs[] string

type Terraform struct {
	Resources []struct {
		Type      string `json:"type"`
		Name      string `json:"name"`
		Provider  string `json:"provider"`
		Instances []struct {
			Attributes   struct {
				ID          string `json:"id"`
				Name        string `json:"name"`
				Project     string `json:"project"`
			} `json:"attributes"`
			Dependencies []string `json:"dependencies"`
		} `json:"instances"`
	} `json:"resources"`
}

func Parser() {

	/*** READ IN terraform.tfstate FILE (entire file) ***/

	inFile, errRead := ioutil.ReadFile("terraform.tfstate")
	// error reading file
	if errRead != nil {
		log.Println("Error reading file.", errRead)
		os.Exit(1)
	}

	/*** UNMARSHAL THE terraform.tfstate FILE ***/

	// "outputs" block
	var outputBlock map [string] interface{}
	json.Unmarshal(inFile, &outputBlock)

	// parse the string to get the outputs
	// ex: map[ab:map[type:string value:34.105.77.168] ip:map[type:string value:34.105.77.168]]
	str := fmt.Sprintln(outputBlock["outputs"])
	outputRegex := regexp.MustCompile(`(([a-z_]*):map)+`)
	output := outputRegex.FindAllStringSubmatch(str, -1)

	// iterates through matches and stores each output in Outputs[] slice
	for i := range output {
		Outputs = append(Outputs, output[i][2])
	}

	// "resources" block, using structs defined above
	json.Unmarshal(inFile, &T)
}
