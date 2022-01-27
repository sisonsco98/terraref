package mapper

import (
	"fmt"
	"log" // logging errors
	"os"  // create and open files
	"strings"

	"KSCD/libraries" // libraries.go
	"KSCD/parser"    // parser.go

	"github.com/beevik/etree" // creating xml file (go get github.com/beevik/etree)
)

// to access (x,y) position of elements on map
var terraNav terraNavigator

type terraNavigator struct {
	hiddenID   int
	Name       string
	XPosCenter int
	YPosCenter int
	Width      int
	Height     int
}

type relationNavigator struct {
	ArrowID    int
	SourceID   int
	TargetID   int
	XPosSource int
	YPosSource int
	XPosTarget int
	YPosTarget int
}

var xml = etree.NewDocument()

// xml element IDs
var globalID int = 0
var elementID int = 0

// dimensions of diagram
var globalXBound = 850
var globalYBound = 1100

// starting (x,y) position
var currentX = 50
var currentY = 50

// slice (array?) of elements
var Pizza []terraNavigator

var ArrowRelationships []relationNavigator

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

	// dependency map
	nameDependencyMap := make(map[string]int)

	/*** CREATE ELEMENT TREE WITH PARSED DATA ***/

	xml.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	mxGraphModel := xml.CreateElement("mxGraphModel")
	mxGraphModel.CreateAttr("gridSize", "10")
	mxGraphModel.CreateAttr("pageWidth", fmt.Sprint(globalXBound))
	mxGraphModel.CreateAttr("pageHeight", fmt.Sprint(globalYBound))

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
		objectName := libraries.LookupName(resourceType)

		// (3) use object name to lookup the draw.io shape (ex: shape=mxgraph.gcp2.gateway)
		objectShape := libraries.LookupShape(objectName)

		// (4) use object name to lookup the correct case of creating the draw.io shape
		t := libraries.LookupCase(objectName)

		if parser.T.Resources[i].Name != "default" {
			nameDependencyMap[parser.T.Resources[i].Name] = elementID
		}

		// (5) Grab the object's name in case it's on a dependency.

		// ???

		// set object's width, height and (x,y) location
		var shapeWidth, shapeHeight = libraries.Dimensions(t)
		var xLocation, yLocation = coordinateFinder(t)

		/*** DETERMINE WHICH XML STRUCTURE IS NEEDED ***/

		switch t {

		/*** GCP / PATHS ***/

		case 0:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(globalID-1))
			globalID = globalID + 1
			mxCell.CreateAttr("value", "")
			mxCell.CreateAttr("style", "edgeStyle=orthogonalEdgeStyle;fontSize=12;html=1;endArrow=blockThin;endFill=1;rounded=0;strokeWidth=2;endSize=4;startSize=4;")
			mxCell.CreateAttr("edge", "1")

			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("relative", "1")
			mxGeometry.CreateAttr("as", "geometry")

			mxPoint := mxGeometry.CreateElement("mxPoint")
			mxPoint.CreateAttr("x", "80")
			mxPoint.CreateAttr("y", "160")
			mxPoint.CreateAttr("as", "sourcePoint")

			mxPoint = mxGeometry.CreateElement("mxPoint")
			mxPoint.CreateAttr("x", "180")
			mxPoint.CreateAttr("y", "160")
			mxPoint.CreateAttr("as", "targetPoint")

		/****************************************************************************************************/

		/*** GCP / SERVICE CARDS ***/

		case 1:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(1))
			globalID = globalID + 1
			mxCell.CreateAttr("value", "")
			mxCell.CreateAttr("style", "strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;")
			mxCell.CreateAttr("vertex", "1")

			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", fmt.Sprint(xLocation))
			mxGeometry.CreateAttr("y", fmt.Sprint(yLocation))
			mxGeometry.CreateAttr("width", fmt.Sprint(shapeWidth))
			mxGeometry.CreateAttr("height", fmt.Sprint(shapeHeight))
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

			var tmp = new(terraNavigator)
			tmp.Name = parser.T.Resources[i].Name
			tmp.hiddenID = globalID - 2
			tmp.XPosCenter = xLocation + (shapeWidth / 2)
			tmp.YPosCenter = yLocation + (shapeHeight / 2)
			tmp.Width = shapeWidth
			tmp.Height = shapeHeight
			Pizza = append(Pizza, *tmp)

		/****************************************************************************************************/

		/*** GCP / USER AND DEVICE CARDS ***/

		case 2:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(1))
			globalID = globalID + 1
			mxCell.CreateAttr("value", parser.T.Resources[i].Type)
			mxCell.CreateAttr("style", fmt.Sprintln("strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;labelPosition=center;verticalLabelPosition=middle;align=center;verticalAlign=bottom;spacingLeft=0;fontColor=#999999;fontSize=12;whiteSpace=wrap;spacingBottom=2;"+objectShape))
			mxCell.CreateAttr("vertex", "1")

			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", fmt.Sprint(xLocation))
			mxGeometry.CreateAttr("y", fmt.Sprint(yLocation))
			mxGeometry.CreateAttr("width", fmt.Sprint(shapeWidth))
			mxGeometry.CreateAttr("height", fmt.Sprint(shapeHeight))
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
			mxGeometry.CreateAttr("width", "50")
			mxGeometry.CreateAttr("height", "50")
			mxGeometry.CreateAttr("relative", "1")
			mxGeometry.CreateAttr("as", "geometry")

			mxPoint := mxGeometry.CreateElement("mxPoint")
			mxPoint.CreateAttr("x", "-25")
			mxPoint.CreateAttr("y", "15")
			mxPoint.CreateAttr("as", "offset")

			var tmp = new(terraNavigator)
			tmp.Name = parser.T.Resources[i].Name
			tmp.hiddenID = globalID - 2
			tmp.XPosCenter = xLocation + (shapeWidth / 2)
			tmp.YPosCenter = yLocation + (shapeHeight / 2)
			tmp.Width = shapeWidth
			tmp.Height = shapeHeight
			Pizza = append(Pizza, *tmp)

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
			mxCell.CreateAttr("parent", fmt.Sprint(1))
			globalID = globalID + 1
			mxCell.CreateAttr("value", "")
			mxCell.CreateAttr("style", "strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;")
			mxCell.CreateAttr("vertex", "1")

			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", fmt.Sprint(xLocation))
			mxGeometry.CreateAttr("y", fmt.Sprint(yLocation))
			mxGeometry.CreateAttr("width", fmt.Sprint(shapeWidth))
			mxGeometry.CreateAttr("height", fmt.Sprint(shapeHeight))
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

			var tmp = new(terraNavigator)
			tmp.Name = parser.T.Resources[i].Name
			tmp.hiddenID = globalID - 2
			tmp.XPosCenter = xLocation + (shapeWidth / 2)
			tmp.YPosCenter = yLocation + (shapeHeight / 2)
			tmp.Width = shapeWidth
			tmp.Height = shapeHeight
			Pizza = append(Pizza, *tmp)

		/****************************************************************************************************/

		/*** GCP / PRODUCT CARDS ***/

		case 4:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(1))
			globalID = globalID + 1
			mxCell.CreateAttr("value", "")
			mxCell.CreateAttr("style", "strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;")
			mxCell.CreateAttr("vertex", "1")

			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", fmt.Sprint(xLocation))
			mxGeometry.CreateAttr("y", fmt.Sprint(yLocation))
			mxGeometry.CreateAttr("width", fmt.Sprint(shapeWidth))
			mxGeometry.CreateAttr("height", fmt.Sprint(shapeHeight))
			mxGeometry.CreateAttr("as", "geometry")

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(globalID-1))
			globalID = globalID + 1
			mxCell.CreateAttr("value", parser.T.Resources[i].Type)
			mxCell.CreateAttr("style", fmt.Sprint("sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;"+objectShape))
			mxCell.CreateAttr("vertex", "1")

			mxGeometry = mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("width", "45")
			mxGeometry.CreateAttr("height", "45")
			mxGeometry.CreateAttr("relative", "1")
			mxGeometry.CreateAttr("as", "geometry")

			mxPoint := mxGeometry.CreateElement("mxPoint")
			mxPoint.CreateAttr("x", "5")
			mxPoint.CreateAttr("y", "7")
			mxPoint.CreateAttr("as", "offset")

			var tmp = new(terraNavigator)
			tmp.Name = parser.T.Resources[i].Name
			tmp.hiddenID = globalID - 2
			tmp.XPosCenter = xLocation + (shapeWidth / 2)
			tmp.YPosCenter = yLocation + (shapeHeight / 2)
			tmp.Width = shapeWidth
			tmp.Height = shapeHeight
			Pizza = append(Pizza, *tmp)

		/****************************************************************************************************/

		/*** GCP ICONS ***/

		case 5:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(1))
			globalID = globalID + 1
			mxCell.CreateAttr("value", parser.T.Resources[i].Type)
			mxCell.CreateAttr("style", fmt.Sprint("sketch=0;html=1;fillColor=#5184F3;strokeColor=none;verticalAlign=top;labelPosition=center;verticalLabelPosition=bottom;align=center;spacingTop=-6;fontSize=11;fontStyle=1;fontColor=#999999;"+objectShape))
			mxCell.CreateAttr("vertex", "1")

			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", fmt.Sprint(xLocation))
			mxGeometry.CreateAttr("y", fmt.Sprint(yLocation))
			mxGeometry.CreateAttr("width", fmt.Sprint(shapeWidth))
			mxGeometry.CreateAttr("height", fmt.Sprint(shapeHeight))
			mxGeometry.CreateAttr("as", "geometry")

			var tmp = new(terraNavigator)
			tmp.Name = parser.T.Resources[i].Name
			tmp.hiddenID = globalID - 2
			tmp.XPosCenter = xLocation + (shapeWidth / 2)
			tmp.YPosCenter = yLocation + (shapeHeight / 2)
			tmp.Width = shapeWidth
			tmp.Height = shapeHeight
			Pizza = append(Pizza, *tmp)

		case 6: // Cloud Scheduler

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(1))
			globalID = globalID + 1
			mxCell.CreateAttr("value", parser.T.Resources[i].Type)
			mxCell.CreateAttr("style", fmt.Sprint("html=1;fillColor=#5184F3;strokeColor=none;verticalAlign=top;labelPosition=center;verticalLabelPosition=bottom;align=center;fontSize=11;fontStyle=1;fontColor=#999999;"+objectShape))
			mxCell.CreateAttr("vertex", "1")

			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", fmt.Sprint(xLocation))
			mxGeometry.CreateAttr("y", fmt.Sprint(yLocation))
			mxGeometry.CreateAttr("width", fmt.Sprint(shapeWidth))
			mxGeometry.CreateAttr("height", fmt.Sprint(shapeHeight))
			mxGeometry.CreateAttr("as", "geometry")

			var tmp = new(terraNavigator)
			tmp.Name = parser.T.Resources[i].Name
			tmp.hiddenID = globalID - 2
			tmp.XPosCenter = xLocation + (shapeWidth / 2)
			tmp.YPosCenter = yLocation + (shapeHeight / 2)
			tmp.Width = shapeWidth
			tmp.Height = shapeHeight
			Pizza = append(Pizza, *tmp)

		/****************************************************************************************************/

		/*** GCP / ZONES ***/

		case 7:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(1))
			globalID = globalID + 1
			mxCell.CreateAttr("value", parser.T.Resources[i].Type)
			mxCell.CreateAttr("style", fmt.Sprint("sketch=0;points=[[0,0,0],[0.25,0,0],[0.5,0,0],[0.75,0,0],[1,0,0],[1,0.25,0],[1,0.5,0],[1,0.75,0],[1,1,0],[0.75,1,0],[0.5,1,0],[0.25,1,0],[0,1,0],[0,0.75,0],[0,0.5,0],[0,0.25,0]];rounded=1;absoluteArcSize=1;arcSize=2;html=1;strokeColor=none;gradientColor=none;shadow=0;dashed=0;fontSize=12;fontColor=#9E9E9E;align=left;verticalAlign=top;spacing=10;spacingTop=-4;"+objectShape))
			mxCell.CreateAttr("vertex", "1")

			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", fmt.Sprint(xLocation))
			mxGeometry.CreateAttr("y", fmt.Sprint(yLocation))
			mxGeometry.CreateAttr("width", fmt.Sprint(shapeWidth))
			mxGeometry.CreateAttr("height", fmt.Sprint(shapeHeight))
			mxGeometry.CreateAttr("as", "geometry")

			var tmp = new(terraNavigator)
			tmp.Name = parser.T.Resources[i].Name
			tmp.hiddenID = globalID - 2
			tmp.XPosCenter = xLocation + (shapeWidth / 2)
			tmp.YPosCenter = yLocation + (shapeHeight / 2)
			tmp.Width = shapeWidth
			tmp.Height = shapeHeight
			Pizza = append(Pizza, *tmp)

		/****************************************************************************************************/

		/*** GCP / EXPANDED PRODUCT CARDS ***/

		// skip for now

		/****************************************************************************************************/

		/*** GCP / GENERAL ICONS ***/

		// skip for now

		/****************************************************************************************************/

		// Error case

		default:
			log.Println("Error: No match.", errCreate)
			os.Exit(1)
		}

		elementID++
	}

	// iterate through all resources
	for r := 0; r < len(parser.T.Resources); r++ {

		// iterate through all instances of resource
		for i := 0; i < len(parser.T.Resources[r].Instances); i++ {

			// iterate through all dependencies of each instance
			for d := 0; d < len(parser.T.Resources[r].Instances[i].Dependencies); d++ {

				// save dependency
				resourceName := parser.T.Resources[r].Instances[i].Dependencies[d]
				dependencyName := strings.Split(resourceName, ".")

				// testing outputs
				// fmt.Println("Parent Resource Name : ", Pizza[r].Name)
				// fmt.Println("Dependency Name : ", dependencyName[1])

				ctr := 0
				for range Pizza {

					// dependencyName[1] since we want the second name
					if Pizza[ctr].Name == dependencyName[1] {

						// testing outputs
						fmt.Println("We've matched the elements.")
						fmt.Println("We need to draw an arrow from element ", Pizza[r].Name, " to element ", Pizza[ctr].Name)
						fmt.Println(Pizza[r].Name, " is located at (", Pizza[r].XPosCenter, ",", Pizza[r].YPosCenter, ")")
						fmt.Println(Pizza[ctr].Name, " is located at (", Pizza[ctr].XPosCenter, ",", Pizza[ctr].YPosCenter, ")")
						fmt.Println(Pizza[r].Name, "'s ID is ", Pizza[r].hiddenID)
						fmt.Println(Pizza[ctr].Name, "'s ID is ", Pizza[ctr].hiddenID)

						/*** CREATE XML ELEMENT FOR ARROW TO CONNECT DEPENDENCIES ***/

						mxCell = root.CreateElement("mxCell")
						mxCell.CreateAttr("id", fmt.Sprint(globalID))
						mxCell.CreateAttr("parent", fmt.Sprint(1))
						fmt.Println(mxCell.GetPath())
						globalID = globalID + 1
						mxCell.CreateAttr("value", "")
						mxCell.CreateAttr("style", "edgeStyle=orthogonalEdgeStyle;fontSize=12;html=1;endArrow=blockThin;endFill=1;rounded=0;strokeWidth=2;endSize=4;startSize=4;")
						mxCell.CreateAttr("edge", "1")
						mxCell.CreateAttr("target", fmt.Sprintf("%d", Pizza[ctr].hiddenID))
						mxCell.CreateAttr("source", fmt.Sprintf("%d", Pizza[r].hiddenID))

						mxGeometry := mxCell.CreateElement("mxGeometry")
						mxGeometry.CreateAttr("relative", "1")
						mxGeometry.CreateAttr("as", "geometry")

						mxPoint := mxGeometry.CreateElement("mxPoint")
						mxPoint.CreateAttr("x", fmt.Sprint(Pizza[r].XPosCenter))
						mxPoint.CreateAttr("y", fmt.Sprint(Pizza[r].YPosCenter))
						mxPoint.CreateAttr("as", "sourcePoint")

						mxPoint = mxGeometry.CreateElement("mxPoint")
						mxPoint.CreateAttr("x", fmt.Sprint(Pizza[ctr].XPosCenter))
						mxPoint.CreateAttr("y", fmt.Sprint(Pizza[ctr].YPosCenter))
						mxPoint.CreateAttr("as", "targetPoint")

						// Creating ArrowNavigator for Validator
						var tmp = new(relationNavigator)
						tmp.ArrowID = globalID - 1
						tmp.SourceID = Pizza[r].hiddenID
						tmp.TargetID = Pizza[ctr].hiddenID
						tmp.XPosSource = Pizza[r].XPosCenter
						tmp.YPosSource = Pizza[r].YPosCenter
						tmp.XPosTarget = Pizza[ctr].XPosCenter
						tmp.YPosTarget = Pizza[ctr].YPosCenter
						ArrowRelationships = append(ArrowRelationships, *tmp)
					}

					ctr++
				}
			}
		}
	}

	/*** PRINT TO THE terraform.drawio FILE ***/

	xml.Indent(4)
	xml.WriteToFile("terraform.drawio")

	//	testing outputs
	//	for key, element := range nameDependencyMap {
	//		fmt.Println(key + " is the element with index " + fmt.Sprint(element))
	//	}

	// close file
	outFile.Close()
}

/*** RETURNS COORDINATES FOR PLACING OBJECTS ***/

func coordinateFinder(class int)(int, int) {

	// get shapeWidth and shapeHeight from libraries by class
	var shapeWidth, shapeHeight = libraries.Dimensions(class)

	// offset objects by 50
	offsetX := shapeWidth + 50
	offsetY := shapeHeight + 50

	// set objects (x,y) position using previously defined offset
	// first fill out row (left -> right), then move to new row
	if (currentX + offsetX + shapeWidth) > globalXBound {
		currentX = 50
		currentY += offsetY
		return currentX, currentY
	} else {
		currentX += offsetX
		return currentX, currentY
	}
}