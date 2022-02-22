package mapper

import (
	"fmt"
	"log" // logging errors
	"os"  // create and open files
	"strings"

	"KSCD/libraries/providers/GCP/utility" //utility.go
	"KSCD/parser"                          // parser.go

	"github.com/beevik/etree" // creating xml file (go get github.com/beevik/etree)
)

// to access (x,y) position of elements on map
var terraNav terraNavigator

type terraNavigator struct {
	HiddenId    int
	Name        string
	XPosCenter  int
	YPosCenter  int
	Width       int
	Height      int
	Project     string
	ObjectShape string
}

type relationNavigator struct {
	ArrowID    int
	SourceID   int
	TargetID   int
	XPosSource int
	YPosSource int
	XPosTarget int
	YPosTarget int
	HasMoved   bool
}

var xml = etree.NewDocument()

// xml element IDs
var globalID int = 0
var elementID int = 0

// dimensions of diagram
var globalXBound = 850
var globalYBound = 1100

// starting (x,y) position
var currentX = 50 - (250 * 2)
var currentY = 50

// slice (array?) of elements
var Pizza []terraNavigator

var ArrowRelationships []relationNavigator

type location struct {
	x int
	y int
}

// Grid is a variably sized array containing the x, y coordinates for the needed number of elements.
// Please keep in mind this is a 1 dimensional array - so there's no natural distinction between rows.
var grid []location

// Grid should be constant - we shouldn't be modifying that. It's just a reference.
// calculatedLocations is an int array where calculatedLocations[i] = x, where i is the resource index
// and x is the grid[x] where we're placing the elements.

// A proper call might look like tempX, tempY := grid[calculatedLocations[i]], where i is the index.
var calculatedLocations []int // What spot on the grid is it assigned to?

var nameList []string

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
	var dependencyOccurences []int

	//Calculate the boundaries in the x and y direction for the purposes of establishing the grid.
	xItemLimit := ((globalXBound-50)/250)/2 + (((globalXBound - 50) / 250) % 2)
	yItemLimit := len(parser.T.Resources) / xItemLimit

	//We're allocating coordinates on the grid based on the above parameters.
	for i := 0; i < yItemLimit; i++ {
		for j := 0; j < xItemLimit; j++ {
			tempX, tempY := coordinateFinder()
			tempObj := location{tempX, tempY}
			grid = append(grid, tempObj)
			calculatedLocations = append(calculatedLocations, 999)
			dependencyOccurences = append(dependencyOccurences, 0)
		}
	}

	// Display the grid - this should display coordinates in columns and rows based on their actual position.
	for i := 0; i < len(parser.T.Resources)-1; i++ {
		if i%xItemLimit != 0 {
			fmt.Println(grid[i].x, grid[i].y)
		} else {
			fmt.Print(grid[i].x, grid[i].y)
			fmt.Print("\t\t")
		}
	}
	fmt.Println()

	// Display it, but with more detail.
	for i := 0; i < len(parser.T.Resources)-1; i++ {
		fmt.Println("Element ", i, " is located at ", grid[i].x, grid[i].y)
	}
	fmt.Println()

	// iterate through all resources and grab unusual ones.
	for i := 0; i < len(parser.T.Resources); i++ {
		if parser.T.Resources[i].Name != "default" {
			fmt.Println("Found an unusual resource called", parser.T.Resources[i].Name, "at index ", i)
			nameDependencyMap[parser.T.Resources[i].Name] = i
		}
	}

	fmt.Println("~Calculating free spaces needed~")

	var containedProjects []string
	subnetNumber := 1

	// Iterate through all resources and fetch COUNT. and
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

				dependencyIndex := nameDependencyMap[dependencyName[1]]

				fmt.Println("Element", r, "needs element", dependencyIndex, "as a dependency.")

				dependencyOccurences[dependencyIndex] += 1
			}

			// Finding dinstinct projects in the file, then creating zones for them
			projectName := parser.T.Resources[r].Instances[i].Attributes.Project

			// If does not exist, add into list
			if doesProjectExist(containedProjects, projectName) == false {
				containedProjects = append(containedProjects, projectName)
			}

			// Checking to see if the project is part of a subnetwork
			if strings.Contains(parser.T.Resources[r].Instances[i].Attributes.ID, "subnet") == true {
				containedProjects = append(containedProjects, fmt.Sprint("subnet ", subnetNumber))
				subnetNumber++
				// TODO: This isnt necessarily correct, but cant fix until we add
			}

		}
	}

	fmt.Println("Contained Projects: ", containedProjects)

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

	/**		CREATING PROJECT REGIONS		**/
	projectX := 30
	projectY := 350
	subX := 30

	// iterate through all resoureces
	for r := 0; r < len(parser.T.Resources); r++ {

		// (1) store resource type (ex: google_api_gateway_gateway)
		resourceType := parser.T.Resources[r].Type

		// (2) use resource type to lookup the draw.io name (ex: Gateway)
		objectName := utility.LookupName(resourceType)

		// (3) use object name to lookup the draw.io shape (ex: shape=mxgraph.gcp2.gateway)
		objectShape := utility.LookupShape(objectName)

		// (5) use specific resource name for main text
		resourceName := parser.T.Resources[r].Instances[0].Attributes.Name

		if parser.T.Resources[r].Name != "default" {
			nameDependencyMap[parser.T.Resources[r].Name] = elementID
		}

		// if name is network, create project area
		if parser.T.Resources[r].Name == "network" {

			minX := projectX
			minY := projectY
			maxX := 375
			maxY := minY + 100

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(1))
			globalID = globalID + 1

			if len(resourceName) > 0 {
				mxCell.CreateAttr("value", fmt.Sprintf("%s	%s", resourceName, resourceType))
			} else {
				mxCell.CreateAttr("value", resourceType)
			}

			mxCell.CreateAttr("style", fmt.Sprint("whiteSpace=wrap;sketch=0;points=[[0,0,0],[0.25,0,0],[0.5,0,0],[0.75,0,0],[1,0,0],[1,0.25,0],[1,0.5,0],[1,0.75,0],[1,1,0],[0.75,1,0],[0.5,1,0],[0.25,1,0],[0,1,0],[0,0.75,0],[0,0.5,0],[0,0.25,0]];rounded=1;absoluteArcSize=1;arcSize=2;html=1;strokeColor=none;gradientColor=none;shadow=0;dashed=0;fontSize=12;fontColor=#9E9E9E;align=left;verticalAlign=top;spacing=10;spacingTop=-4;"+objectShape))
			mxCell.CreateAttr("vertex", "1")

			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", fmt.Sprint(minX))
			mxGeometry.CreateAttr("y", fmt.Sprint(minY))
			mxGeometry.CreateAttr("width", fmt.Sprint(maxX))
			mxGeometry.CreateAttr("height", fmt.Sprint(maxY))
			mxGeometry.CreateAttr("as", "geometry")

			var tmp = new(terraNavigator)
			tmp.Name = parser.T.Resources[r].Name
			tmp.HiddenId = globalID - 2
			tmp.XPosCenter = minX + (maxX / 2)
			tmp.YPosCenter = minY + (maxY / 2)
			tmp.Width = maxX
			tmp.Height = maxY
			tmp.Project = parser.T.Resources[r].Instances[0].Attributes.Project
			Pizza = append(Pizza, *tmp)

			projectX = projectX + 500
		}

		// if name is subnetwork, create project area
		if parser.T.Resources[r].Name == "subnetwork" {

			minX := subX + 5
			minY := projectY + 30
			maxX := 350
			maxY := projectY + 60

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(1))
			globalID = globalID + 1

			if len(resourceName) > 0 {
				mxCell.CreateAttr("value", fmt.Sprintf("%s	%s", resourceName, resourceType))
			} else {
				mxCell.CreateAttr("value", resourceType)
			}

			mxCell.CreateAttr("style", fmt.Sprint("whiteSpace=wrap;sketch=0;points=[[0,0,0],[0.25,0,0],[0.5,0,0],[0.75,0,0],[1,0,0],[1,0.25,0],[1,0.5,0],[1,0.75,0],[1,1,0],[0.75,1,0],[0.5,1,0],[0.25,1,0],[0,1,0],[0,0.75,0],[0,0.5,0],[0,0.25,0]];rounded=1;absoluteArcSize=1;arcSize=2;html=1;strokeColor=none;gradientColor=none;shadow=0;dashed=0;fontSize=12;fontColor=#9E9E9E;align=left;verticalAlign=top;spacing=10;spacingTop=-4;"+objectShape))
			mxCell.CreateAttr("vertex", "1")

			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", fmt.Sprint(minX))
			mxGeometry.CreateAttr("y", fmt.Sprint(minY))
			mxGeometry.CreateAttr("width", fmt.Sprint(maxX))
			mxGeometry.CreateAttr("height", fmt.Sprint(maxY))
			mxGeometry.CreateAttr("as", "geometry")

			var tmp = new(terraNavigator)
			tmp.Name = parser.T.Resources[r].Name
			tmp.HiddenId = globalID - 2
			tmp.XPosCenter = minX + (maxX / 2)
			tmp.YPosCenter = minY + (maxY / 2)
			tmp.Width = maxX
			tmp.Height = maxY
			tmp.Project = parser.T.Resources[r].Instances[0].Attributes.Project
			Pizza = append(Pizza, *tmp)
		}
	}

	/* ITERATE THROUGH RESOURCES */

	test := utility.LookupZone("User 1 (Default)")
	fmt.Println(test)

	for i := 0; i < len(parser.T.Resources); i++ {

		// (1) store resource type (ex: google_api_gateway_gateway)
		resourceType := parser.T.Resources[i].Type

		// (2) use resource type to lookup the draw.io name (ex: Gateway)
		objectName := utility.LookupName(resourceType)

		// (3) use object name to lookup the draw.io shape (ex: shape=mxgraph.gcp2.gateway)
		objectShape := utility.LookupShape(objectName)

		// (4) use object name to lookup the correct case of creating the draw.io shape
		t := utility.LookupCase(objectName)

		// (5) use specific resource name for main text
		resourceName := parser.T.Resources[i].Instances[0].Attributes.Name

		if parser.T.Resources[i].Name != "default" {
			nameDependencyMap[parser.T.Resources[i].Name] = elementID
		}

		// (5) Grab the object's name in case it's on a dependency.

		// ???

		// set object's width, height and (x,y) location
		var shapeWidth, shapeHeight = utility.Dimensions(t)
		var xLocation, yLocation = coordinateFinder()

		/*** DETERMINE WHICH XML STRUCTURE IS NEEDED ***/

		if parser.T.Resources[i].Name == "network" || parser.T.Resources[i].Name == "subnetwork" {
			continue
		}

		switch t {

		/*** GCP / PATHS ***/

		case 0:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(globalID-1))
			globalID = globalID + 1
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
			globalID = globalID + 1
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
			mxCell.CreateAttr("parent", fmt.Sprint(globalID-1))
			globalID = globalID + 1

			if len(resourceName) > 0 {
				mxCell.CreateAttr("value", fmt.Sprintf("%s	%s", resourceName, resourceType))
			} else {
				mxCell.CreateAttr("value", resourceType)
			}

			mxCell.CreateAttr("style", fmt.Sprint("whiteSpace=wrap;sketch=0;dashed=0;connectable=0;html=1;fillColor=#757575;strokeColor=none;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontSize=12;"+objectShape))
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
			tmp.Project = parser.T.Resources[i].Instances[0].Attributes.Project
			Pizza = append(Pizza, *tmp)

		/****************************************************************************************************/

		/*** GCP / USER AND DEVICE CARDS ***/

		case 2:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(1))
			globalID = globalID + 1

			if len(resourceName) > 0 {
				mxCell.CreateAttr("value", fmt.Sprintf("%s	%s", resourceName, resourceType))
			} else {
				mxCell.CreateAttr("value", resourceType)
			}

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
			mxCell.CreateAttr("style", fmt.Sprint("whiteSpace=wrap;sketch=0;dashed=0;connectable=0;html=1;fillColor=#757575;strokeColor=none;part=1;"+objectShape))
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
			tmp.Project = parser.T.Resources[i].Instances[0].Attributes.Project
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
			mxCell.CreateAttr("parent", fmt.Sprint(globalID-1))
			globalID = globalID + 1

			if len(resourceName) > 0 {
				mxCell.CreateAttr("value", fmt.Sprintf("%s	%s", resourceName, resourceType))
			} else {
				mxCell.CreateAttr("value", resourceType)
			}

			mxCell.CreateAttr("style", fmt.Sprint("whiteSpace=wrap;sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;"+objectShape))
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
			tmp.Project = parser.T.Resources[i].Instances[0].Attributes.Project
			tmp.ObjectShape = objectShape
			Pizza = append(Pizza, *tmp)

		/****************************************************************************************************/

		/*** GCP / PRODUCT CARDS ***/

		case 4:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(1))
			globalID = globalID + 1
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
			mxCell.CreateAttr("parent", fmt.Sprint(globalID-1))
			globalID = globalID + 1

			if len(resourceName) > 0 {
				mxCell.CreateAttr("value", fmt.Sprintf("%s	%s", resourceName, resourceType))
			} else {
				mxCell.CreateAttr("value", resourceType)
			}

			mxCell.CreateAttr("style", fmt.Sprint("whiteSpace=wrap;html=1;sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;"+objectShape))
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
			tmp.Project = parser.T.Resources[i].Instances[0].Attributes.Project
			Pizza = append(Pizza, *tmp)

		/****************************************************************************************************/

		/*** GCP ICONS ***/

		case 5:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(1))
			globalID = globalID + 1

			if len(resourceName) > 0 {
				mxCell.CreateAttr("value", fmt.Sprintf("%s	%s", resourceName, resourceType))
			} else {
				mxCell.CreateAttr("value", resourceType)
			}

			mxCell.CreateAttr("style", fmt.Sprint("whiteSpace=wrap;sketch=0;html=1;fillColor=#5184F3;strokeColor=none;verticalAlign=top;labelPosition=center;verticalLabelPosition=bottom;align=center;spacingTop=-6;fontSize=11;fontStyle=1;fontColor=#999999;"+objectShape))
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
			tmp.Project = parser.T.Resources[i].Instances[0].Attributes.Project
			Pizza = append(Pizza, *tmp)

		case 6: // Cloud Scheduler

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(1))
			globalID = globalID + 1

			if len(resourceName) > 0 {
				mxCell.CreateAttr("value", fmt.Sprintf("%s	%s", resourceName, resourceType))
			} else {
				mxCell.CreateAttr("value", resourceType)
			}

			mxCell.CreateAttr("style", fmt.Sprint("whiteSpace=wrap;html=1;fillColor=#5184F3;strokeColor=none;verticalAlign=top;labelPosition=center;verticalLabelPosition=bottom;align=center;fontSize=11;fontStyle=1;fontColor=#999999;"+objectShape))
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
			tmp.Project = parser.T.Resources[i].Instances[0].Attributes.Project
			Pizza = append(Pizza, *tmp)

		/****************************************************************************************************/

		/*** GCP / ZONES ***/

		case 7:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(1))
			globalID = globalID + 1

			if len(resourceName) > 0 {
				mxCell.CreateAttr("value", fmt.Sprintf("%s	%s", resourceName, resourceType))
			} else {
				mxCell.CreateAttr("value", resourceType)
			}

			mxCell.CreateAttr("style", fmt.Sprint("whiteSpace=wrap;sketch=0;points=[[0,0,0],[0.25,0,0],[0.5,0,0],[0.75,0,0],[1,0,0],[1,0.25,0],[1,0.5,0],[1,0.75,0],[1,1,0],[0.75,1,0],[0.5,1,0],[0.25,1,0],[0,1,0],[0,0.75,0],[0,0.5,0],[0,0.25,0]];rounded=1;absoluteArcSize=1;arcSize=2;html=1;strokeColor=none;gradientColor=none;shadow=0;dashed=0;fontSize=12;fontColor=#9E9E9E;align=left;verticalAlign=top;spacing=10;spacingTop=-4;"+objectShape))
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
			tmp.Project = parser.T.Resources[i].Instances[0].Attributes.Project
			Pizza = append(Pizza, *tmp)

		//ID, Value, Style, Vertex, Parent
		case 8:
			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(globalID))
			mxCell.CreateAttr("parent", fmt.Sprint(1))
			mxCell.CreateAttr("value", resourceType) //?
			mxCell.CreateAttr("vertex", fmt.Sprint(1))
			mxCell.CreateAttr("style", fmt.Sprint(utility.LookupZone(parser.T.Resources[i].Name)))

			//x, y, wid, hei, as
			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", "160")
			mxGeometry.CreateAttr("y", "120")
			mxGeometry.CreateAttr("width", "160")
			mxGeometry.CreateAttr("height", "120")
			mxGeometry.CreateAttr("as", "geometry")

			globalID = globalID + 1
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

	/**		USING DEPENDENCIES TO GET ARROWS DRAWN		**/
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
						// fmt.Println("We've matched the elements.")
						// fmt.Println("We need to draw an arrow from element ", Pizza[r].Name, " to element ", Pizza[ctr].Name)
						// fmt.Println(Pizza[r].Name, " is located at (", Pizza[r].XPosCenter, ",", Pizza[r].YPosCenter, ")")
						// fmt.Println(Pizza[ctr].Name, " is located at (", Pizza[ctr].XPosCenter, ",", Pizza[ctr].YPosCenter, ")")
						// fmt.Println(Pizza[r].Name, "'s ID is ", Pizza[r].HiddenId)
						// fmt.Println(Pizza[ctr].Name, "'s ID is ", Pizza[ctr].HiddenId)

						/*** CREATE XML ELEMENT FOR ARROW TO CONNECT DEPENDENCIES ***/

						mxCell = root.CreateElement("mxCell")
						mxCell.CreateAttr("id", fmt.Sprint(globalID))
						mxCell.CreateAttr("parent", fmt.Sprint(1))
						// fmt.Println(mxCell.GetPath())
						globalID = globalID + 1
						mxCell.CreateAttr("value", "")
						mxCell.CreateAttr("style", "whiteSpace=wrap;html=1;edgeStyle=orthogonalEdgeStyle;fontSize=12;html=1;endArrow=blockThin;endFill=1;rounded=0;strokeWidth=2;endSize=4;startSize=4;")
						mxCell.CreateAttr("edge", "1")
						mxCell.CreateAttr("target", fmt.Sprintf("%d", Pizza[ctr].HiddenId))
						mxCell.CreateAttr("source", fmt.Sprintf("%d", Pizza[r].HiddenId))

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
						tmp.SourceID = Pizza[r].HiddenId
						tmp.TargetID = Pizza[ctr].HiddenId
						tmp.XPosSource = Pizza[r].XPosCenter
						tmp.YPosSource = Pizza[r].YPosCenter
						tmp.XPosTarget = Pizza[ctr].XPosCenter
						tmp.YPosTarget = Pizza[ctr].YPosCenter
						tmp.HasMoved = false
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

func coordinateFinder() (int, int) {

	// get shapeWidth and shapeHeight from libraries by class
	var shapeWidth, shapeHeight = 250, 60

	// offset objects by 50
	offsetX := shapeWidth * 2
	offsetY := shapeHeight * 2

	// set objects (x,y) position using previously defined offset
	// first fill out row (left -> right), then move to new row
	if (currentX + offsetX + shapeWidth) > globalXBound {
		//fmt.Sprint("The current X is ", currentX)

		currentX = 50
		currentY += offsetY
		return currentX, currentY
	} else {

		currentX += offsetX
		return currentX, currentY
	}
}

/*** RETURNS WHETHER OR NOT A PROJECT EXISTS ***/

func doesProjectExist(s []string, str string) bool {

	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
