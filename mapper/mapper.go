package mapper

import (
	"fmt"
	"log"			// logging errors
	"os"			// create and open files
	"strings"

	"KSCD/parser"
	"KSCD/libraries/providers/GCP/utility"

	// creating xml file (go get github.com/beevik/etree)
	"github.com/beevik/etree"
)

/*** GLOBAL STRUCTS TO STORE AND ACCESS INFO ABOUT ELEMENTS AND ARROWS ***/

var terraNav terraNavigator

type terraNavigator struct {
	HiddenId   int
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

/*** GLOBAL SLICES FOR ELEMENTS AND ARROWS ***/

var Elements []terraNavigator
var Arrows []relationNavigator

/*** CREATE GLOBAL XML TREE ***/

var xml = etree.NewDocument()

// xml element IDs
var globalID int = 0
var elementID int = 0

// dimensions of diagram
var globalXBound, globalYBound = 850, 1100

// dimensions of element (normal card)
var shapeWidth, shapeHeight = 250, 60

// starting (x,y) position
var xPos, yPos = 50 - (2 * shapeWidth), 50

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

	/*** CREATE GRID FOR PLACING ELEMENTS ***/

	// elements to be placed on the (x, y) locations on the grid
	type location struct {
		x, y int
	}
	var grid []location

	// dependency map
	nameDependencyMap := make(map[string]int)
	var numDependents []int
	var numDependencies []int

	// determine the dimensions of the grid
	var rows, cols int
	cols = ((globalXBound - 50) / 250) / 2 + (((globalXBound - 50) / 250) % 2)
	if (len(parser.T.Resources) % 2 == 0) {
		rows = len(parser.T.Resources) / cols
	} else {
		rows = len(parser.T.Resources) / cols + 1
	}

	// allocate the (x, y) locations on the grid using coordinateFinder
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			tempX, tempY := coordinateFinder()
			tempObj := location{tempX, tempY}
			grid = append(grid, tempObj)
			numDependents = append(numDependents, 0)
			numDependencies = append(numDependencies, 0)
		}
	}

	fmt.Println()
	fmt.Println("/**************************************************/")
	fmt.Println("/*                 GRID LOCATIONS                 */")
	fmt.Println("/**************************************************/")
	fmt.Println()

	// display the grid locations and the element currently in each location
	index := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if index < len(parser.T.Resources) {
				fmt.Print("Element ", index, ": (", grid[index].x, ", ", grid[index].y, ")")
				fmt.Print("\t")
				index++
			}
		}
		fmt.Println()
	}

	// iterate through all resources and store dependencies (non default)
	for i := 0; i < len(parser.T.Resources); i++ {
		if parser.T.Resources[i].Name != "default" {
			nameDependencyMap[parser.T.Resources[i].Name] = i
		}
	}

	fmt.Println()
	fmt.Println("/**************************************************/")
	fmt.Println("/*                  DEPENDENCIES                  */")
	fmt.Println("/**************************************************/")
	fmt.Println()


	/*** FOR EACH RESOURCE, COUNT THE NUMBER OF DEPENDENCIES / DEPENDENTS  ***/

	var dependency, resourceName, rName string
	var dependencyName []string
	var dependencyIndex int

	// iterate through each resource -> instance -> dependency
	for r := 0; r < len(parser.T.Resources); r++ {
		for i := 0; i < len(parser.T.Resources[r].Instances); i++ {
			for d := 0; d < len(parser.T.Resources[r].Instances[i].Dependencies); d++ {

				// save any dependencies
				dependency = parser.T.Resources[r].Instances[i].Dependencies[d]
				dependencyName = strings.Split(dependency, ".")
				dependencyIndex = nameDependencyMap[dependencyName[1]]

				// increment numDependencies and numDependents
				numDependencies[r] += 1
				numDependents[dependencyIndex] += 1

			}
		}
	}

	/*** FOR EACH RESOURCE, FIND ITS DEPENDENCIES AND DEPENDENTS  ***/

	// iterate through each resource
	for r := 0; r < len(parser.T.Resources); r++ {

		// find and print the index and name of each resource that is a dependency of the current element
		fmt.Print("Element ", r, " has the ", numDependencies[r], " dependencies: \t")
		for i := 0; i < len(parser.T.Resources[r].Instances); i++ {
			if len(parser.T.Resources[r].Instances[i].Dependencies) > 0 {
				for d := 0; d < len(parser.T.Resources[r].Instances[i].Dependencies); d++ {
					dependency = parser.T.Resources[r].Instances[i].Dependencies[d]
					dependencyName = strings.Split(dependency, ".")
					dependencyIndex = nameDependencyMap[dependencyName[1]]
					fmt.Print(dependencyIndex, " (", dependencyName[1], ") / ")
				}
			}
		}
		fmt.Println()

		// find and print the index and name of each resource which has the current element as a dependency
		fmt.Print(numDependents[r], " elements are dependent on Element ", r, ": \t")
		if (numDependents[r] > 0) {
			rName = parser.T.Resources[r].Name
			for resource := 0; resource < len(parser.T.Resources); resource++ {
				resourceName = parser.T.Resources[resource].Type
				for i := 0; i < len(parser.T.Resources[resource].Instances); i++ {
					for d := 0; d < len(parser.T.Resources[resource].Instances[i].Dependencies); d++ {
						if len(parser.T.Resources[resource].Instances[i].Dependencies) > 0 {
							dependency = parser.T.Resources[resource].Instances[i].Dependencies[d]
							dependencyName = strings.Split(dependency, ".")
							dependencyIndex = nameDependencyMap[dependencyName[1]]
							if rName == dependencyName[1] {
								fmt.Print(resource, " (", resourceName, ") / ")
							}
						}
					}
				}
			}
		}
		fmt.Println()
		fmt.Println()

	}

	/*** CREATE ELEMENT TREE WITH PARSED DATA ***/

	xml.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	mxGraphModel := xml.CreateElement("mxGraphModel")
	mxGraphModel.CreateAttr("gridSize", "10")
	mxGraphModel.CreateAttr("pageWidth", fmt.Sprint(globalXBound))
	mxGraphModel.CreateAttr("pageHeight", fmt.Sprint(globalYBound))

	root := mxGraphModel.CreateElement("root")

	mxCell := root.CreateElement("mxCell")
	mxCell.CreateAttr("id", fmt.Sprint(globalID))
	globalID++

	mxCell = root.CreateElement("mxCell")
	mxCell.CreateAttr("id", fmt.Sprint(globalID))
	mxCell.CreateAttr("parent", fmt.Sprint(globalID - 1))
	globalID++

	/* ITERATE THROUGH RESOURCES */

	for i := 0; i < len(parser.T.Resources); i++ {

		// (1) store resource type (ex: google_storage_bucket)
		resourceType := parser.T.Resources[i].Type

		if parser.T.Resources[i].Name != "default" {
			// store the name and id of dependency elements
			nameDependencyMap[parser.T.Resources[i].Name] = elementID
		}

		// (2) use resource type to lookup the draw.io name (ex: Bucket)
		objectName := utility.LookupName(resourceType)

		// (3) use object name to lookup the draw.io shape (ex: shape=mxgraph.gcp2.bucket)
		objectShape := utility.LookupShape(objectName)

		// (4) use object name to lookup the correct case of creating the draw.io shape (ex: 1)
		t := utility.LookupCase(objectName)

		// (5) use specific resource name for main text (ex: example-storage-bucket)
		resourceName := parser.T.Resources[i].Instances[0].Attributes.Name

		// set current elements location based off grid (x, y) locations
		var xLocation, yLocation = grid[i].x, grid[i].y

		/*** DETERMINE WHICH XML STRUCTURE IS NEEDED ***/

		switch t {

		/*** GCP / PATHS ***/

		case 0:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(globalID - 1))
			globalID++

			mxCell.CreateAttr("value", "")
			mxCell.CreateAttr("style", "whiteSpace=wrap;html=1;edgeStyle=orthogonalEdgeStyle;fontSize=12;html=1;endArrow=blockThin;endFill=1;rounded=0;strokeWidth=2;endSize=4;startSize=4;")
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
			globalID++

			mxCell.CreateAttr("value", "")
			mxCell.CreateAttr("style", "whiteSpace=wrap;html=1;strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;")
			mxCell.CreateAttr("vertex", "1")

			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", fmt.Sprint(xLocation))
			mxGeometry.CreateAttr("y", fmt.Sprint(yLocation))
			mxGeometry.CreateAttr("width", fmt.Sprint(shapeWidth))
			mxGeometry.CreateAttr("height", fmt.Sprint(shapeHeight))
			mxGeometry.CreateAttr("as", "geometry")

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(globalID - 1))
			globalID++

			if len(resourceName) > 0 {
				mxCell.CreateAttr("value", fmt.Sprintf("%s	%s", resourceName, resourceType))
			} else {
				mxCell.CreateAttr("value", resourceType)
			}

			mxCell.CreateAttr("style", fmt.Sprint("whiteSpace=wrap;sketch=0;dashed=0;connectable=0;html=1;fillColor=#757575;strokeColor=none;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontSize=12;" + objectShape))
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
			tmp.HiddenId = globalID - 2
			tmp.XPosCenter = xLocation + (shapeWidth / 2)
			tmp.YPosCenter = yLocation + (shapeHeight / 2)
			tmp.Width = shapeWidth
			tmp.Height = shapeHeight
			Elements = append(Elements, *tmp)

		/****************************************************************************************************/

		/*** GCP / USER AND DEVICE CARDS ***/

		case 2:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(1))
			globalID++

			if len(resourceName) > 0 {
				mxCell.CreateAttr("value", fmt.Sprintf("%s	%s", resourceName, resourceType))
			} else {
				mxCell.CreateAttr("value", resourceType)
			}

			mxCell.CreateAttr("style", fmt.Sprintln("strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;labelPosition=center;verticalLabelPosition=middle;align=center;verticalAlign=bottom;spacingLeft=0;fontColor=#999999;fontSize=12;whiteSpace=wrap;spacingBottom=2;" + objectShape))
			mxCell.CreateAttr("vertex", "1")

			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", fmt.Sprint(xLocation))
			mxGeometry.CreateAttr("y", fmt.Sprint(yLocation))
			mxGeometry.CreateAttr("width", fmt.Sprint(shapeWidth))
			mxGeometry.CreateAttr("height", fmt.Sprint(shapeHeight))
			mxGeometry.CreateAttr("as", "geometry")

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(globalID - 1))
			globalID++

			mxCell.CreateAttr("value", "")
			mxCell.CreateAttr("style", fmt.Sprint("whiteSpace=wrap;sketch=0;dashed=0;connectable=0;html=1;fillColor=#757575;strokeColor=none;part=1;" + objectShape))
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
			tmp.HiddenId = globalID - 2
			tmp.XPosCenter = xLocation + (shapeWidth / 2)
			tmp.YPosCenter = yLocation + (shapeHeight / 2)
			tmp.Width = shapeWidth
			tmp.Height = shapeHeight
			Elements = append(Elements, *tmp)

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
			globalID++

			mxCell.CreateAttr("value", "")
			mxCell.CreateAttr("style", "whiteSpace=wrap;html=1;strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;")
			mxCell.CreateAttr("vertex", "1")

			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", fmt.Sprint(xLocation))
			mxGeometry.CreateAttr("y", fmt.Sprint(yLocation))
			mxGeometry.CreateAttr("width", fmt.Sprint(shapeWidth))
			mxGeometry.CreateAttr("height", fmt.Sprint(shapeHeight))
			mxGeometry.CreateAttr("as", "geometry")

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(globalID - 1))
			globalID++

			if len(resourceName) > 0 {
				mxCell.CreateAttr("value", fmt.Sprintf("%s	%s", resourceName, resourceType))
			} else {
				mxCell.CreateAttr("value", resourceType)
			}

			mxCell.CreateAttr("style", fmt.Sprint("whiteSpace=wrap;sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" + objectShape))
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
			tmp.HiddenId = globalID - 2
			tmp.XPosCenter = xLocation + (shapeWidth / 2)
			tmp.YPosCenter = yLocation + (shapeHeight / 2)
			tmp.Width = shapeWidth
			tmp.Height = shapeHeight
			Elements = append(Elements, *tmp)

		/****************************************************************************************************/

		/*** GCP / PRODUCT CARDS ***/

		case 4:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(1))
			globalID++

			mxCell.CreateAttr("value", "")
			mxCell.CreateAttr("style", "whiteSpace=wrap;html=1;strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;")
			mxCell.CreateAttr("vertex", "1")

			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", fmt.Sprint(xLocation))
			mxGeometry.CreateAttr("y", fmt.Sprint(yLocation))
			mxGeometry.CreateAttr("width", fmt.Sprint(shapeWidth))
			mxGeometry.CreateAttr("height", fmt.Sprint(shapeHeight))
			mxGeometry.CreateAttr("as", "geometry")

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(globalID - 1))
			globalID++

			if len(resourceName) > 0 {
				mxCell.CreateAttr("value", fmt.Sprintf("%s	%s", resourceName, resourceType))
			} else {
				mxCell.CreateAttr("value", resourceType)
			}

			mxCell.CreateAttr("style", fmt.Sprint("whiteSpace=wrap;html=1;sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" + objectShape))
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
			tmp.HiddenId = globalID - 2
			tmp.XPosCenter = xLocation + (shapeWidth / 2)
			tmp.YPosCenter = yLocation + (shapeHeight / 2)
			tmp.Width = shapeWidth
			tmp.Height = shapeHeight
			Elements = append(Elements, *tmp)

		/****************************************************************************************************/

		/*** GCP ICONS ***/

		case 5:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(1))
			globalID++

			if len(resourceName) > 0 {
				mxCell.CreateAttr("value", fmt.Sprintf("%s	%s", resourceName, resourceType))
			} else {
				mxCell.CreateAttr("value", resourceType)
			}

			mxCell.CreateAttr("style", fmt.Sprint("whiteSpace=wrap;sketch=0;html=1;fillColor=#5184F3;strokeColor=none;verticalAlign=top;labelPosition=center;verticalLabelPosition=bottom;align=center;spacingTop=-6;fontSize=11;fontStyle=1;fontColor=#999999;" + objectShape))
			mxCell.CreateAttr("vertex", "1")

			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", fmt.Sprint(xLocation))
			mxGeometry.CreateAttr("y", fmt.Sprint(yLocation))
			mxGeometry.CreateAttr("width", fmt.Sprint(shapeWidth))
			mxGeometry.CreateAttr("height", fmt.Sprint(shapeHeight))
			mxGeometry.CreateAttr("as", "geometry")

			var tmp = new(terraNavigator)
			tmp.Name = parser.T.Resources[i].Name
			tmp.HiddenId = globalID - 2
			tmp.XPosCenter = xLocation + (shapeWidth / 2)
			tmp.YPosCenter = yLocation + (shapeHeight / 2)
			tmp.Width = shapeWidth
			tmp.Height = shapeHeight
			Elements = append(Elements, *tmp)

		case 6: // Cloud Scheduler

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(1))
			globalID++

			if len(resourceName) > 0 {
				mxCell.CreateAttr("value", fmt.Sprintf("%s	%s", resourceName, resourceType))
			} else {
				mxCell.CreateAttr("value", resourceType)
			}

			mxCell.CreateAttr("style", fmt.Sprint("whiteSpace=wrap;html=1;fillColor=#5184F3;strokeColor=none;verticalAlign=top;labelPosition=center;verticalLabelPosition=bottom;align=center;fontSize=11;fontStyle=1;fontColor=#999999;" + objectShape))
			mxCell.CreateAttr("vertex", "1")

			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", fmt.Sprint(xLocation))
			mxGeometry.CreateAttr("y", fmt.Sprint(yLocation))
			mxGeometry.CreateAttr("width", fmt.Sprint(shapeWidth))
			mxGeometry.CreateAttr("height", fmt.Sprint(shapeHeight))
			mxGeometry.CreateAttr("as", "geometry")

			var tmp = new(terraNavigator)
			tmp.Name = parser.T.Resources[i].Name
			tmp.HiddenId = globalID - 2
			tmp.XPosCenter = xLocation + (shapeWidth / 2)
			tmp.YPosCenter = yLocation + (shapeHeight / 2)
			tmp.Width = shapeWidth
			tmp.Height = shapeHeight
			Elements = append(Elements, *tmp)

		/****************************************************************************************************/

		/*** GCP / ZONES ***/

		case 7:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(1))
			globalID++

			if len(resourceName) > 0 {
				mxCell.CreateAttr("value", fmt.Sprintf("%s	%s", resourceName, resourceType))
			} else {
				mxCell.CreateAttr("value", resourceType)
			}

			mxCell.CreateAttr("style", fmt.Sprint("whiteSpace=wrap;sketch=0;points=[[0,0,0],[0.25,0,0],[0.5,0,0],[0.75,0,0],[1,0,0],[1,0.25,0],[1,0.5,0],[1,0.75,0],[1,1,0],[0.75,1,0],[0.5,1,0],[0.25,1,0],[0,1,0],[0,0.75,0],[0,0.5,0],[0,0.25,0]];rounded=1;absoluteArcSize=1;arcSize=2;html=1;strokeColor=none;gradientColor=none;shadow=0;dashed=0;fontSize=12;fontColor=#9E9E9E;align=left;verticalAlign=top;spacing=10;spacingTop=-4;" + objectShape))
			mxCell.CreateAttr("vertex", "1")

			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", fmt.Sprint(xLocation))
			mxGeometry.CreateAttr("y", fmt.Sprint(yLocation))
			mxGeometry.CreateAttr("width", fmt.Sprint(shapeWidth))
			mxGeometry.CreateAttr("height", fmt.Sprint(shapeHeight))
			mxGeometry.CreateAttr("as", "geometry")

			var tmp = new(terraNavigator)
			tmp.Name = parser.T.Resources[i].Name
			tmp.HiddenId = globalID - 2
			tmp.XPosCenter = xLocation + (shapeWidth / 2)
			tmp.YPosCenter = yLocation + (shapeHeight / 2)
			tmp.Width = shapeWidth
			tmp.Height = shapeHeight
			Elements = append(Elements, *tmp)

		case 8:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(1))
			globalID++

			mxCell.CreateAttr("value", resourceType)
			mxCell.CreateAttr("vertex", fmt.Sprint(1))
			mxCell.CreateAttr("style", fmt.Sprint(utility.LookupZone(parser.T.Resources[i].Name)))

			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", "160")
			mxGeometry.CreateAttr("y", "120")
			mxGeometry.CreateAttr("width", "160")
			mxGeometry.CreateAttr("height", "120")
			mxGeometry.CreateAttr("as", "geometry")

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

				/*** CREATE XML ELEMENT FOR ARROW TO CONNECT DEPENDENCIES ***/

				ctr := 0
				for range Elements {

					// find elements that are dependencies
					if Elements[ctr].Name == dependencyName[1] {

						mxCell = root.CreateElement("mxCell")
						mxCell.CreateAttr("id", fmt.Sprint(globalID))
						mxCell.CreateAttr("parent", fmt.Sprint(1))						
						globalID++

						mxCell.CreateAttr("value", "")
						mxCell.CreateAttr("style", "whiteSpace=wrap;html=1;edgeStyle=orthogonalEdgeStyle;fontSize=12;html=1;endArrow=blockThin;endFill=1;rounded=0;strokeWidth=2;endSize=4;startSize=4;")
						mxCell.CreateAttr("edge", "1")
						mxCell.CreateAttr("target", fmt.Sprintf("%d", Elements[ctr].HiddenId))
						mxCell.CreateAttr("source", fmt.Sprintf("%d", Elements[r].HiddenId))

						mxGeometry := mxCell.CreateElement("mxGeometry")
						mxGeometry.CreateAttr("relative", "1")
						mxGeometry.CreateAttr("as", "geometry")

						mxPoint := mxGeometry.CreateElement("mxPoint")
						mxPoint.CreateAttr("x", fmt.Sprint(Elements[r].XPosCenter))
						mxPoint.CreateAttr("y", fmt.Sprint(Elements[r].YPosCenter))
						mxPoint.CreateAttr("as", "sourcePoint")

						mxPoint = mxGeometry.CreateElement("mxPoint")
						mxPoint.CreateAttr("x", fmt.Sprint(Elements[ctr].XPosCenter))
						mxPoint.CreateAttr("y", fmt.Sprint(Elements[ctr].YPosCenter))
						mxPoint.CreateAttr("as", "targetPoint")

						var tmp = new(relationNavigator)
						tmp.ArrowID = globalID - 1
						tmp.SourceID = Elements[r].HiddenId
						tmp.XPosSource = Elements[r].XPosCenter
						tmp.YPosSource = Elements[r].YPosCenter
						tmp.TargetID = Elements[ctr].HiddenId
						tmp.XPosTarget = Elements[ctr].XPosCenter
						tmp.YPosTarget = Elements[ctr].YPosCenter
						Arrows = append(Arrows, *tmp)
					}

					ctr++
				}
			}
		}
	}

	/*** PRINT TO THE terraform.drawio FILE ***/

	xml.Indent(4)
	xml.WriteToFile("terraform.drawio")

	// close file
	outFile.Close()
}

/*** RETURNS COORDINATES FOR PLACING OBJECTS ***/

func coordinateFinder() (int, int) {

	// offset objects by shapeWidth, shapeHeight
	xOffset := shapeWidth * 2
	yOffset := shapeHeight * 2

	// set objects (x,y) position using previously defined offset
	// first fill out row (left -> right), then move to new row
	if (xPos + xOffset + shapeWidth) > globalXBound {
		xPos = 50
		yPos += yOffset
		return xPos, yPos
	} else {
		xPos += xOffset
		return xPos, yPos
	}
}
