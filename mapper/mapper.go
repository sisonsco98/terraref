package mapper

import (
	"fmt"
	"log"			// logging errors
	"os"			// create and open file
	"KSCD/parser"	// parser.go

	"github.com/beevik/etree"	// creating xml file (go get github.com/beevik/etree)
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
	mxGraphModel.CreateAttr("gridSize", fmt.Sprint(10))
	mxGraphModel.CreateAttr("pageWidth", fmt.Sprint(850))
	mxGraphModel.CreateAttr("pageHeight", fmt.Sprint(1100))

	root := mxGraphModel.CreateElement("root")

	mxCell := root.CreateElement("mxCell")
	mxCell.CreateAttr("id", "0")

	mxCell = root.CreateElement("mxCell")
	mxCell.CreateAttr("id", "1")
	mxCell.CreateAttr("parent", "0")

	// beginning part of style attribute for GCP icons
	gcpIconFormat := "sketch=0;html=1;fillColor=#5184F3;strokeColor=none;verticalAlign=top;labelPosition=center;verticalLabelPosition=bottom;align=center;spacingTop=-6;fontSize=11;fontStyle=1;fontColor=#999999;shape="

	// generates GCP icons; does not work for cards yet
	for i := 0; i < len(parser.T.Resources); i++ {
		mxCell = root.CreateElement("mxCell")
		mxCell.CreateAttr("id", fmt.Sprint(i + 2))
		mxCell.CreateAttr("value", parser.T.Resources[i].Type)

		// some loop or something to determine the icon we need; possibly based off of parser.T.Resources[i].Type ?
		gcpIconShape := "mxgraph.gcp2.hexIcon;prIcon=gpu"

		mxCell.CreateAttr("style", fmt.Sprint(gcpIconFormat + gcpIconShape))
		mxCell.CreateAttr("parent", fmt.Sprint(1))	// this seems to be set for GCP icons
		mxCell.CreateAttr("vertex", fmt.Sprint(1))	// this seems to be set for GCP icons

		mxGeometry := mxCell.CreateElement("mxGeometry")
		mxGeometry.CreateAttr("x", fmt.Sprint(225 + (i * 200)))		// 
		mxGeometry.CreateAttr("y", fmt.Sprint(550))
		mxGeometry.CreateAttr("width", fmt.Sprint(66))		// size for GCP icons
		mxGeometry.CreateAttr("height", fmt.Sprint(58.5))	// size for GCP icons
		mxGeometry.CreateAttr("as", "geometry")
	}

	/*** PRINT TO THE terraform.drawio FILE ***/

	xml.Indent(4)
	xml.WriteToFile("terraform.drawio")
	
	// close file
	outFile.Close()
}
