package mapper

import (
	"fmt"
	"log"				// logging errors
	"os"				// create and open file
//	"KSCD/parser"		// parser.go
	"KSCD/libraries"	// parser.go

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

/*********************************************************************************************************************************************************************************************************************************/

//	objectType := "NAT"
	objectType := "Gateway"
	t := libraries.Lookup(objectType)
	gcpShape := libraries.ShapeLookup[objectType]

//	// generates GCP icons; does not work for cards yet
//	for i := 0; i < len(parser.T.Resources); i++ {

//		// lookup value for switch
//		t := libraries.Lookup(fmt.Sprint(parser.T.Resources[i].Type))

		switch t {
		
		/* <<< GCP / SERVICE CARDS */

		case 1:

			mxCell = root.CreateElement("mxCell")
//			mxCell.CreateAttr("id", fmt.Sprint(i + 2)) ////
			mxCell.CreateAttr("id", fmt.Sprint(2))
			mxCell.CreateAttr("value", "")
			mxCell.CreateAttr("style", "strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;")
			mxCell.CreateAttr("vertex", fmt.Sprint(1))
			mxCell.CreateAttr("parent", fmt.Sprint(1))
	
			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", "360")			// ??
			mxGeometry.CreateAttr("y", "120")			// ??
			mxGeometry.CreateAttr("width", "150")		// ??
			mxGeometry.CreateAttr("height", "75")		// ??
			mxGeometry.CreateAttr("as", "geometry")

			mxCell = root.CreateElement("mxCell")
//			mxCell.CreateAttr("id", fmt.Sprint(i + 2)) ////
			mxCell.CreateAttr("id", fmt.Sprint(3)) ////
//			mxCell.CreateAttr("value", parser.T.Resources[i].Type)
			mxCell.CreateAttr("value", objectType)
			mxCell.CreateAttr("style", fmt.Sprint("sketch=0;dashed=0;connectable=0;html=1;fillColor=#757575;strokeColor=none;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontSize=12;" + gcpShape))
			mxCell.CreateAttr("vertex", "1")	// this seems to be set for GCP icons
//			mxCell.CreateAttr("parent", fmt.Sprint(i + 1)) ////
			mxCell.CreateAttr("parent", fmt.Sprint(2)) ////
	
			mxGeometry = mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("y", "0.5")
			mxGeometry.CreateAttr("width", "32")
			mxGeometry.CreateAttr("height", "32")
			mxGeometry.CreateAttr("relative", "1")
			mxGeometry.CreateAttr("as", "geometry")

			mxPoint := mxGeometry.CreateElement("mxPoint")
			mxPoint.CreateAttr("x", "5")
			mxPoint.CreateAttr("y", "-16")
			mxPoint.CreateAttr("as", "offset")
		
		/* GCP / SERVICE CARDS >>> */

		/* <<< GCP ICONS */

		// case 2:
		// 	mxCell = root.CreateElement("mxCell")
		// 	mxCell.CreateAttr("id", fmt.Sprint(i + 2))
		// 	mxCell.CreateAttr("value", parser.T.Resources[i].Type)
		// 	mxCell.CreateAttr("style", fmt.Sprint("sketch=0;html=1;fillColor=#5184F3;strokeColor=none;verticalAlign=top;labelPosition=center;verticalLabelPosition=bottom;align=center;spacingTop=-6;fontSize=11;fontStyle=1;fontColor=#999999;" + gcpShape))
		// 	mxCell.CreateAttr("vertex", fmt.Sprint(1))	// this seems to be set for GCP icons
		// 	mxCell.CreateAttr("parent", fmt.Sprint(1))	// this seems to be set for GCP icons
	
		// 	mxGeometry := mxCell.CreateElement("mxGeometry")
		// 	mxGeometry.CreateAttr("x", fmt.Sprint(225 + (i * 200)))		// 
		// 	mxGeometry.CreateAttr("y", fmt.Sprint(550))
		// 	mxGeometry.CreateAttr("width", fmt.Sprint(66))		// size for GCP icons
		// 	mxGeometry.CreateAttr("height", fmt.Sprint(58.5))	// size for GCP icons
		// 	mxGeometry.CreateAttr("as", "geometry")

		/* GCP ICONS >>> */

		// error case
		default:
			log.Println("Error: No match.", errCreate)
			os.Exit(1)
		}

//	}

	/*** PRINT TO THE terraform.drawio FILE ***/

	xml.Indent(4)
	xml.WriteToFile("terraform.drawio")
	
	// close file
	outFile.Close()
}
