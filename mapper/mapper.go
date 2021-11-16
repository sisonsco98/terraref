package mapper

import (
	"fmt"
	"log"				// logging errors
	"os"				// create and open file
	"KSCD/parser"		// parser.go
	"KSCD/libraries"	// parser.go

	"github.com/beevik/etree"	// creating xml file (go get github.com/beevik/etree)
)

var globalID int = 0

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
	mxGraphModel.CreateAttr("gridSize", "10")
	mxGraphModel.CreateAttr("pageWidth", "850")
	mxGraphModel.CreateAttr("pageHeight", "1100")

	root := mxGraphModel.CreateElement("root")

	mxCell := root.CreateElement("mxCell")
	mxCell.CreateAttr("id", fmt.Sprint(globalID))
	globalID = globalID + 1

	mxCell = root.CreateElement("mxCell")
	mxCell.CreateAttr("id", fmt.Sprint(globalID))
	mxCell.CreateAttr("parent", fmt.Sprint(globalID-1))
	globalID = globalID + 1

	/* ITERATE THROUGH RESOURCES */

	for i := 0; i < len(parser.T.Resources); i++ {

		// (1) store resource type (ex: google_api_gateway_gateway)
		resourceType := parser.T.Resources[i].Type

		// (2) use resource type to lookup the draw.io name (ex: Gateway)
		objectName := libraries.NameLookup[resourceType]

		// (3) use object name to lookup the draw.io shape (ex: shape=mxgraph.gcp2.gateway)
		objectShape := libraries.ShapeLookup[objectName]

		t := libraries.Lookup(objectName)

		switch t {

		/*** GCP / SERVICE CARDS ***/

		case 1:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(globalID-1))
			globalID = globalID + 1
			mxCell.CreateAttr("value", "")
			mxCell.CreateAttr("style", "strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;")
			mxCell.CreateAttr("vertex", "1")

			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", "360") // DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("y", "120") // DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("width", "150")
			mxGeometry.CreateAttr("height", "56")
			mxGeometry.CreateAttr("as", "geometry")

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(globalID-1))
			globalID = globalID + 1
			mxCell.CreateAttr("value", parser.T.Resources[i].Type)
			mxCell.CreateAttr("style", fmt.Sprint("sketch=0;dashed=0;connectable=0;html=1;fillColor=#757575;strokeColor=none;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontSize=12;"+objectShape))
			mxCell.CreateAttr("vertex", "1")

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

		/****************************************************************************************************/

		/*** GCP / USER AND DEVICE CARDS ***/

		case 2:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(globalID-1))
			globalID = globalID + 1
			mxCell.CreateAttr("value", parser.T.Resources[i].Type)
			mxCell.CreateAttr("style", "strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;labelPosition=center;verticalLabelPosition=middle;align=center;verticalAlign=bottom;spacingLeft=0;fontColor=#999999;fontSize=12;whiteSpace=wrap;spacingBottom=2;")
			mxCell.CreateAttr("vertex", "1")
			mxCell.CreateAttr("parent", "1")

			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", "360") // DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("y", "120") // DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("width", "70")
			mxGeometry.CreateAttr("height", "100")
			mxGeometry.CreateAttr("as", "geometry")

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(globalID-1))
			globalID = globalID + 1
			mxCell.CreateAttr("value", "")
			mxCell.CreateAttr("style", fmt.Sprint("sketch=0;dashed=0;connectable=0;html=1;fillColor=#757575;strokeColor=none;part=1;"+objectShape))
			mxCell.CreateAttr("vertex", "1")

			mxGeometry = mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", "0.5")
			mxGeometry.CreateAttr("width", "50")  // INCONSISTENT, SOME MAY LOOK OFF
			mxGeometry.CreateAttr("height", "50") // INCONSISTENT, SOME MAY LOOK OFF
			mxGeometry.CreateAttr("relative", "1")
			mxGeometry.CreateAttr("as", "geometry")

			mxPoint := mxGeometry.CreateElement("mxPoint")
			mxPoint.CreateAttr("x", "-25") // INCONSISTENT, SOME MAY LOOK OFF
			mxPoint.CreateAttr("y", "15")  // INCONSISTENT, SOME MAY LOOK OFF
			mxPoint.CreateAttr("as", "offset")

		/****************************************************************************************************/

		/*** GCP / COMPUTE ***/
		/*** GCP / API MANAGEMENT ***/
		/*** GCP / SECURITY ***/
		/*** GCP / DATA ANALYTICS ***/
		/*** GCP / DATA TRANSFER ***/
		/*** GCP / CLOUD AI ***/
		/*** GCP / INTERNET OF THINGS ***/
		/*** GCP / DATABASES ***/
		/*** GCP / STORAGE ***/
		/*** GCP / MANAGEMENT TOOLS ***/
		/*** GCP / NETWORKING ***/
		/*** GCP / DEVELOPER TOOLS ***/

		case 3:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(globalID-1))
			globalID = globalID + 1
			mxCell.CreateAttr("value", "")
			mxCell.CreateAttr("style", "strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;")
			mxCell.CreateAttr("vertex", "1")

			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", "360") // DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("y", "120") // DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("width", "175")
			mxGeometry.CreateAttr("height", "60")
			mxGeometry.CreateAttr("as", "geometry")

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(globalID-1))
			globalID = globalID + 1
			mxCell.CreateAttr("value", parser.T.Resources[i].Type)
			mxCell.CreateAttr("style", fmt.Sprint("sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;"+objectShape))
			mxCell.CreateAttr("vertex", "1")

			mxGeometry = mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("y", "0.5")
			mxGeometry.CreateAttr("width", "44")
			mxGeometry.CreateAttr("height", "39")
			mxGeometry.CreateAttr("relative", "1")
			mxGeometry.CreateAttr("as", "geometry")

			mxPoint := mxGeometry.CreateElement("mxPoint")
			mxPoint.CreateAttr("x", "5")
			mxPoint.CreateAttr("y", "-19.5")
			mxPoint.CreateAttr("as", "offset")

		/****************************************************************************************************/

		/*** GCP / PRODUCT CARDS ***/

		case 4:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(globalID-1))
			globalID = globalID + 1
			mxCell.CreateAttr("value", "")
			mxCell.CreateAttr("style", "strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;")
			mxCell.CreateAttr("vertex", "1")

			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", "360") // DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("y", "120") // DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("width", "150")
			mxGeometry.CreateAttr("height", "60")
			mxGeometry.CreateAttr("as", "geometry")

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(globalID-1))
			globalID = globalID + 1
			mxCell.CreateAttr("value", parser.T.Resources[i].Type)
			mxCell.CreateAttr("style", fmt.Sprint("sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;"+objectShape))
			mxCell.CreateAttr("vertex", "1")

			mxGeometry = mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("width", "45")  // INCONSISTENT, SOME MAY LOOK OFF (esp. Avere Physical Appliance, Fastly, Google Play Game Services)
			mxGeometry.CreateAttr("height", "45") // INCONSISTENT, SOME MAY LOOK OFF (esp. Avere Physical Appliance, Fastly, Google Play Game Services)
			mxGeometry.CreateAttr("relative", "1")
			mxGeometry.CreateAttr("as", "geometry")

			mxPoint := mxGeometry.CreateElement("mxPoint")
			mxPoint.CreateAttr("x", "5")
			mxPoint.CreateAttr("y", "7") // INCONSISTENT, SOME MAY LOOK OFF (esp. Avere Physical Appliance, Fastly, Google Play Game Services)
			mxPoint.CreateAttr("as", "offset")

		/****************************************************************************************************/

		/*** GCP ICONS ***/

		case 5:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(globalID-1))
			globalID = globalID + 1
			mxCell.CreateAttr("value", parser.T.Resources[i].Type)
			mxCell.CreateAttr("style", fmt.Sprint("sketch=0;html=1;fillColor=#5184F3;strokeColor=none;verticalAlign=top;labelPosition=center;verticalLabelPosition=bottom;align=center;spacingTop=-6;fontSize=11;fontStyle=1;fontColor=#999999;"+objectShape))
			mxCell.CreateAttr("vertex", "1")

			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", "360") // DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("y", "120") // DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("width", "66")
			mxGeometry.CreateAttr("height", "58.5")
			mxGeometry.CreateAttr("as", "geometry")

		case 6: // Cloud Scheduler

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(globalID-1))
			globalID = globalID + 1
			mxCell.CreateAttr("value", parser.T.Resources[i].Type)
			mxCell.CreateAttr("style", fmt.Sprint("html=1;fillColor=#5184F3;strokeColor=none;verticalAlign=top;labelPosition=center;verticalLabelPosition=bottom;align=center;fontSize=11;fontStyle=1;fontColor=#999999;"+objectShape))
			mxCell.CreateAttr("vertex", "1")

			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", "360") // DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("y", "120") // DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("width", "30")
			mxGeometry.CreateAttr("height", "34.5")
			mxGeometry.CreateAttr("as", "geometry")

		/****************************************************************************************************/

		/*** GCP / PATHS ***/

		//case 7:

		// skip for now

		/****************************************************************************************************/

		/*** GCP / ZONES ***/

		case 8:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(globalID-1))
			globalID = globalID + 1
			mxCell.CreateAttr("value", parser.T.Resources[i].Type)
			mxCell.CreateAttr("style", fmt.Sprint("sketch=0;points=[[0,0,0],[0.25,0,0],[0.5,0,0],[0.75,0,0],[1,0,0],[1,0.25,0],[1,0.5,0],[1,0.75,0],[1,1,0],[0.75,1,0],[0.5,1,0],[0.25,1,0],[0,1,0],[0,0.75,0],[0,0.5,0],[0,0.25,0]];rounded=1;absoluteArcSize=1;arcSize=2;html=1;strokeColor=none;gradientColor=none;shadow=0;dashed=0;fontSize=12;fontColor=#9E9E9E;align=left;verticalAlign=top;spacing=10;spacingTop=-4;"+objectShape))
			mxCell.CreateAttr("vertex", "1")

			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", "360") // DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("y", "120") // DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("width", "220")
			mxGeometry.CreateAttr("height", "190")
			mxGeometry.CreateAttr("as", "geometry")

		/****************************************************************************************************/

		/*** GCP / EXPANDED PRODUCT CARDS ***/

		//	case 9:
		// skip for now

		/****************************************************************************************************/

		/*** GCP / GENERAL ICONS ***/

		//	case 10:
		// skip for now

		/****************************************************************************************************/

		// General case

		case 99:

		/****************************************************************************************************/

		// Error case

		default:
			log.Println("Error: No match.", errCreate)
			os.Exit(1)
		}

	}

	/*** PRINT TO THE terraform.drawio FILE ***/

	xml.Indent(4)
	xml.WriteToFile("terraform.drawio")

	// close file
	outFile.Close()
}
