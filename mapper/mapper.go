package mapper

import (
	"fmt"
	"log"			// logging errors
	"os"			// create and open file
	"github.com/beevik/etree"	// creating xml file (go get github.com/beevik/etree)
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

	/*** CREATE ELEMENT TREE WITH PARSED DATA ***/
	
	xml := etree.NewDocument()
	xml.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	mxGraphModel := xml.CreateElement("mxGraphModel")
	root := mxGraphModel.CreateElement("root")

	mxCell := root.CreateElement("mxCell")
	mxCell.CreateAttr("id", "0")
	mxCell = root.CreateElement("mxCell")
	mxCell.CreateAttr("id", "1")
	mxCell.CreateAttr("parent", "0")

	for i := 0; i < len(parser.T.Resources); i++ {
		mxCell = root.CreateElement("mxCell")
		mxCell.CreateAttr("id", fmt.Sprint(i + 2))
		mxCell.CreateAttr("value", parser.T.Resources[i].Type)
		mxCell.CreateAttr("style", "shape=mxgraph.gcp2.hexIcon;prIcon=gpu")		// style is where we'd set the icon, need to create string somehow then put here
		mxCell.CreateAttr("vertex", fmt.Sprint(1))
		mxCell.CreateAttr("parent", fmt.Sprint(1))

		mxGeometry := mxCell.CreateElement("mxGeometry")
		mxGeometry.CreateAttr("x", fmt.Sprint(260))
		mxGeometry.CreateAttr("y", fmt.Sprint(260))
		mxGeometry.CreateAttr("width", fmt.Sprint(66))
		mxGeometry.CreateAttr("height", fmt.Sprint(58.5))
		mxGeometry.CreateAttr("as", "geometry")
	}

	// // iterates through outputs slice
	// for i := range parser.Outputs {
	// 	output := xml.CreateElement("Output")
	// 	output.CreateAttr("output", parser.Outputs[i])
	// }

	// // iterate through the resources
	// for i := 0; i < len(parser.T.Resources); i++ {
	// 	resource := xml.CreateElement("Resource")
	// 	resource.CreateAttr("type", parser.T.Resources[i].Type)
	// 	resource.CreateAttr("name", parser.T.Resources[i].Name)
	// 	resource.CreateAttr("provider", parser.Providers[i])

	// 	// iterate through the instances
	// 	for j := 0; j < len(parser.T.Resources[i].Instances); j++ {

	// 		instance := resource.CreateElement("Instance")
	// 		attribute := instance.CreateElement("Attribute")

	// 		if len(parser.T.Resources[i].Instances[j].Attributes.ID) > 0 {
	// 			attribute.CreateAttr("id", parser.T.Resources[i].Instances[j].Attributes.ID)
	// 		}
	// 		if len(parser.T.Resources[i].Instances[j].Attributes.Name) > 0 {
	// 			attribute.CreateAttr("name", parser.T.Resources[i].Instances[j].Attributes.Name)
	// 		}
	// 		if len(parser.T.Resources[i].Instances[j].Attributes.Project) > 0 {
	// 			attribute.CreateAttr("project", parser.T.Resources[i].Instances[j].Attributes.Project)
	// 		}

	// 		// iterate through dependencies
	// 		for k := 0; k < len(parser.T.Resources[i].Instances[j].Dependencies); k++ {
	// 			instance.CreateAttr("dependency", parser.T.Resources[i].Instances[j].Dependencies[k])
	// 		}
	// 	}
	// }

	/*** PRINT TO THE terraform.drawio FILE ***/

	xml.Indent(4)
//	xml.WriteTo(os.Stdout)
	xml.WriteToFile("terraform.drawio")
	
	// close file
	outFile.Close()
}
