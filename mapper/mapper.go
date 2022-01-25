package mapper

import (
	"KSCD/libraries" // parser.go
	"KSCD/parser"    // parser.go
	"fmt"
	"log" // logging errors
	"os"  // create and open files
	"strings"
	//"testing/quick"

	"github.com/beevik/etree" // creating xml file (go get github.com/beevik/etree)
)

// The intention here - the primary wall at the moment is DEPENDENCIES.
// This is WHERE things are on the map, that's why we have an X and Y position.
// Arrows need an x/y origin + x/y target, so I figured we could just access this.
var terraNav terraNavigator
type terraNavigator struct {
	hiddenID int
	Name string
	XPos int
	YPos int
}

var globalID int = 0
var elementID int = 0
var xml = etree.NewDocument()



//Making this global just because I can for now
var globalXBound = 850
var globalYBound = 1100

var currentX = 50
var currentY = 50

// This is a slice (array?) of elements - we generate four boxes per element, but only really one element.
var Pizza []terraNavigator

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

	// Add a new map - Where's
	nameDependencyMap := make(map[string]int)

	/*** CREATE ELEMENT TREE WITH PARSED DATA ***/

	xml.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	mxGraphModel := xml.CreateElement("mxGraphModel")
	mxGraphModel.CreateAttr("gridSize", "10")
	mxGraphModel.CreateAttr("pageWidth", fmt.Sprint(globalXBound))
	mxGraphModel.CreateAttr("pageHeight",fmt.Sprint(globalYBound))

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



		// There's two distinct things we're grabbing here - (x, y) for the size of the object in question.
		// The other thing we're grabbing is a place to place it.
		var shapeWidth, shapeHeight = libraries.Dimensions(t)
		var xLocation, yLocation = coordinateFinder(t)

//		currentX += 40	// offset?
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
			mxGeometry.CreateAttr("x", fmt.Sprint(xLocation)) // DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("y", fmt.Sprint(yLocation)) // DETERMINE METHOD FOR SETTING THIS
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


			// So we're creating a new
			var tmp = new(terraNavigator)
			tmp.Name = parser.T.Resources[i].Name
			tmp.hiddenID = elementID
			tmp.XPos = xLocation
			tmp.YPos = yLocation
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
//			mxCell.CreateAttr("parent", "1")

			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", fmt.Sprint(xLocation)) // DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("y", fmt.Sprint(yLocation)) // DETERMINE METHOD FOR SETTING THIS
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
			mxGeometry.CreateAttr("width", "50")  // INCONSISTENT, SOME MAY LOOK OFF
			mxGeometry.CreateAttr("height", "50") // INCONSISTENT, SOME MAY LOOK OFF
			mxGeometry.CreateAttr("relative", "1")
			mxGeometry.CreateAttr("as", "geometry")

			mxPoint := mxGeometry.CreateElement("mxPoint")
			mxPoint.CreateAttr("x", "-25") // INCONSISTENT, SOME MAY LOOK OFF
			mxPoint.CreateAttr("y", "15")  // INCONSISTENT, SOME MAY LOOK OFF
			mxPoint.CreateAttr("as", "offset")

			var tmp = new(terraNavigator)
			tmp.Name = parser.T.Resources[i].Name
			tmp.hiddenID = elementID
			tmp.XPos = xLocation
			tmp.YPos = yLocation
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
			mxGeometry.CreateAttr("x", fmt.Sprint(xLocation)) // DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("y", fmt.Sprint(yLocation)) // DETERMINE METHOD FOR SETTING THIS
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
			tmp.hiddenID = elementID
			tmp.XPos = xLocation
			tmp.YPos = yLocation
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
			mxGeometry.CreateAttr("x", fmt.Sprint(xLocation)) // DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("y", fmt.Sprint(yLocation)) // DETERMINE METHOD FOR SETTING THIS
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
			mxGeometry.CreateAttr("width", "45")  // INCONSISTENT, SOME MAY LOOK OFF (esp. Avere Physical Appliance, Fastly, Google Play Game Services)
			mxGeometry.CreateAttr("height", "45") // INCONSISTENT, SOME MAY LOOK OFF (esp. Avere Physical Appliance, Fastly, Google Play Game Services)
			mxGeometry.CreateAttr("relative", "1")
			mxGeometry.CreateAttr("as", "geometry")

			mxPoint := mxGeometry.CreateElement("mxPoint")
			mxPoint.CreateAttr("x", "5")
			mxPoint.CreateAttr("y", "7") // INCONSISTENT, SOME MAY LOOK OFF (esp. Avere Physical Appliance, Fastly, Google Play Game Services)
			mxPoint.CreateAttr("as", "offset")

			var tmp = new(terraNavigator)
			tmp.Name = parser.T.Resources[i].Name
			tmp.hiddenID = elementID
			tmp.XPos = xLocation
			tmp.YPos = yLocation
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
			mxGeometry.CreateAttr("x", fmt.Sprint(xLocation)) // DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("y", fmt.Sprint(yLocation)) // DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("width", fmt.Sprint(shapeWidth))
			mxGeometry.CreateAttr("height", fmt.Sprint(shapeHeight))
			mxGeometry.CreateAttr("as", "geometry")

			var tmp = new(terraNavigator)
			tmp.Name = parser.T.Resources[i].Name
			tmp.hiddenID = elementID
			tmp.XPos = xLocation
			tmp.YPos = yLocation
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
			mxGeometry.CreateAttr("x", fmt.Sprint(xLocation)) // DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("y", fmt.Sprint(yLocation)) // DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("width", fmt.Sprint(shapeWidth))
			mxGeometry.CreateAttr("height", fmt.Sprint(shapeHeight))
			mxGeometry.CreateAttr("as", "geometry")

			var tmp = new(terraNavigator)
			tmp.Name = parser.T.Resources[i].Name
			tmp.hiddenID = elementID
			tmp.XPos = xLocation
			tmp.YPos = yLocation
			Pizza = append(Pizza, *tmp)

		/****************************************************************************************************/

		/*** GCP / PATHS ***/

		//case 7:
		// skip for now

		/****************************************************************************************************/

		/*** GCP / ZONES ***/

		case 8:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(1))
			globalID = globalID + 1
			mxCell.CreateAttr("value", parser.T.Resources[i].Type)
			mxCell.CreateAttr("style", fmt.Sprint("sketch=0;points=[[0,0,0],[0.25,0,0],[0.5,0,0],[0.75,0,0],[1,0,0],[1,0.25,0],[1,0.5,0],[1,0.75,0],[1,1,0],[0.75,1,0],[0.5,1,0],[0.25,1,0],[0,1,0],[0,0.75,0],[0,0.5,0],[0,0.25,0]];rounded=1;absoluteArcSize=1;arcSize=2;html=1;strokeColor=none;gradientColor=none;shadow=0;dashed=0;fontSize=12;fontColor=#9E9E9E;align=left;verticalAlign=top;spacing=10;spacingTop=-4;"+objectShape))
			mxCell.CreateAttr("vertex", "1")

			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", fmt.Sprint(xLocation)) // DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("y", fmt.Sprint(yLocation)) // DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("width", fmt.Sprint(shapeWidth))
			mxGeometry.CreateAttr("height", fmt.Sprint(shapeHeight))
			mxGeometry.CreateAttr("as", "geometry")

			var tmp = new(terraNavigator)
			tmp.Name = parser.T.Resources[i].Name
			tmp.hiddenID = elementID
			tmp.XPos = xLocation
			tmp.YPos = yLocation
			Pizza = append(Pizza, *tmp)
		/****************************************************************************************************/

		/*** GCP / EXPANDED PRODUCT CARDS ***/

		//	case 9:
		// skip for now

		/****************************************************************************************************/

		/*** GCP / GENERAL ICONS ***/

		//	case 10:
		// skip for now

		/****************************************************************************************************/

		// Error case

		default:
			log.Println("Error: No match.", errCreate)
			os.Exit(1)
		}

		elementID++
	}

	/*** PRINT TO THE terraform.drawio FILE ***/

	// Checking the resources now.

	// This loops over all the resources parsed from the file initially.
	for r := 0; r < len(parser.T.Resources); r++ {

		// Dependencies are nested inside the instances block, so we go deeper in this loop.
		// 1/13 - Can remove this? Probably? Don't break what doesn't work atm.
		for i := 0; i < len(parser.T.Resources[r].Instances); i++ {

			// If they exist, dependencies are in the instances block. We're looping over the Dependencies.
			for d := 0; d < len(parser.T.Resources[r].Instances[i].Dependencies); d++ {

				// Grab the resource we're LOOKING for
				resourceName := parser.T.Resources[r].Instances[i].Dependencies[d]
				dependencyName := strings.Split(resourceName, ".") //turns it into a slice?

				fmt.Println("Parent Resource Name : ", Pizza[r].Name)
				fmt.Println("Dependency Name : ", dependencyName[1])

				ctr := 0
				for range Pizza {
					if (Pizza[ctr].Name == dependencyName[1]){ //dependencyName[1] since we want the second name.

						fmt.Println("We've matched the elements.")

						fmt.Println("We need to draw an arrow from element ", Pizza[r].Name, " to element ", Pizza[ctr].Name)

						fmt.Println(Pizza[r].Name, " is located at (", Pizza[r].XPos, ",", Pizza[r].YPos, ")")

						fmt.Println(Pizza[ctr].Name, " is located at (", Pizza[ctr].XPos, ",", Pizza[ctr].YPos , ")")

						mxCell = root.CreateElement("mxCell")
						mxCell.CreateAttr("id", fmt.Sprint(globalID))
						mxCell.CreateAttr("parent", fmt.Sprint(1))
						globalID = globalID + 1
						mxCell.CreateAttr("value", "")
						mxCell.CreateAttr("style", "edgeStyle=orthogonalEdgeStyle;fontSize=12;html=1;endArrow=blockThin;endFill=1;rounded=0;strokeWidth=2;endSize=4;startSize=4;")
						mxCell.CreateAttr("edge", "1")

						mxGeometry := mxCell.CreateElement("mxGeometry")
						mxGeometry.CreateAttr("relative", "1")
						mxGeometry.CreateAttr("as", "geometry")

						mxPoint := mxGeometry.CreateElement("mxPoint")
						mxPoint.CreateAttr("x", fmt.Sprint(Pizza[r].XPos))
						mxPoint.CreateAttr("y", fmt.Sprint(Pizza[r].YPos))
						mxPoint.CreateAttr("as", "sourcePoint")

						mxPoint = mxGeometry.CreateElement("mxPoint")
						mxPoint.CreateAttr("x", fmt.Sprint(Pizza[ctr].XPos))
						mxPoint.CreateAttr("y", fmt.Sprint(Pizza[ctr].YPos))
						mxPoint.CreateAttr("as", "targetPoint")
					}


					ctr++
				}

				// NOTE: prints dependencies as type.name, might need just type ??
			}
		}
	}

	xml.Indent(4)
	xml.WriteToFile("terraform.drawio")

	// 1/13 Reworded debugging loop over the nameDependencyMap
	for key, element := range nameDependencyMap {
		fmt.Println(key + " is the element with index " + fmt.Sprint(element))
	}


	// Removed 1/13/21 - This is just for debugging. We list the elements and where to find them.

	//ctr := 0
	//for range Pizza {
	//	fmt.Println(Pizza[ctr].Name)
	//	fmt.Println(Pizza[ctr].hiddenID)
	//	fmt.Println(Pizza[ctr].XPos)
	//	fmt.Println(Pizza[ctr].YPos)
	//	ctr++
	//}


	// close file
	outFile.Close()
}

func coordinateFinder(class int) (int, int){
	var shapeWidth, shapeHeight = libraries.Dimensions(class)
	offsetX := shapeWidth + 50
	offsetY := shapeHeight + 50

	if ((currentX + offsetX) > globalXBound) {
		currentX = 50
		currentY += offsetY
		return currentX, currentY
	} else {
		currentX += offsetX
		return currentX - offsetX, currentY
	}
}