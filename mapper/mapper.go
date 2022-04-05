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

/*** GLOBAL SLICES FOR ELEMENTS AND ARROWS ***/

var Elements []TerraNavigator
var Arrows []RelationNavigator

/*** GLOBAL STRUCTS TO STORE AND ACCESS INFO ABOUT ELEMENTS AND ARROWS ***/

var TerraNav TerraNavigator

type TerraNavigator struct {
	Name					string
	Project					string
	ObjectShape				string
	HiddenId				int
	XPosCenter, YPosCenter	int
	Width, Height			int
}

type RelationNavigator struct {
	ArrowID					int
	SourceID				int
	XPosSource, YPosSource	int
	TargetID				int
	XPosTarget, YPosTarget	int
}

/*** CREATE GLOBAL XML TREE ***/

var XML = etree.NewDocument()

/*** DEFAULT DIMENSIONS ***/

var DimX, DimY = 850, 1100				// diagram
var CardWidth, CardHeight = 250, 60		// normal cards
var ZoneWidth, ZoneHeight = 350, 380	// project zones

var GlobalID = 0

func Mapper(outFileLocation string) {

	/*** CREATE THE .drawio FILE ***/

	outFile, errCreate := os.Create(outFileLocation)
	// error creating file
	if errCreate != nil {
		log.Println("Error creating file.", errCreate)
		os.Exit(1)
	}
	// keep open
	defer outFile.Close()

	/*** CREATE THE GRID FOR PLACING ELEMENTS ***/

	fmt.Println()
	fmt.Println("****************************************************************************************************")
	fmt.Println("*                                   G R I D    L O C A T I O N S                                   *")
	fmt.Println("****************************************************************************************************")
	fmt.Println()

	// grid locations
	type location struct {
		x, y int
	}
	var grid []location

	// set grid dimensions
	rows, cols := len(parser.T.Resources) + 1, len(parser.T.Resources)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			// store grid locations
			tempX, tempY := 50 + (CardWidth * 2 * c), 50 + (CardHeight * 2 * r)
			grid = append(grid, location{tempX, tempY})
			// print grid locations
			fmt.Print("(", grid[(r * len(parser.T.Resources)) + c].x, ", ", grid[(r * len(parser.T.Resources)) + c].y, ")", "\t")
		}
		fmt.Println()
	}

	/*** DETERMINE THE DEPENDENCIES AND DEPENDENTS OF EACH RESOURCE ***/

	fmt.Println()
	fmt.Println("****************************************************************************************************")
	fmt.Println("*            R E S O U R C E    D E P E N D E N C I E S    A N D    D E P E N D E N T S            *")
	fmt.Println("****************************************************************************************************")
	fmt.Println()

	// map each resource name to resource index
	nameToIndex := make(map[string]int)
	for i := 0; i < len(parser.T.Resources); i++ {
		if parser.T.Resources[i].Name != "default" {
			nameToIndex[parser.T.Resources[i].Name] = i
		}
	}

	// number of dependents and dependencies for each resource
	numDependents := make([]int, rows * cols)
	numDependencies := make([]int, rows * cols)

	// list of dependencies and dependents for each resource
	dependencyNames := make(map[int][]string)
	dependencyIndices := make(map[int][]int)
	dependentNames := make(map[int][]string)
	dependentIndices := make(map[int][]int)

	// iterate through each resource -> instance -> dependency to count numDependencies and numDependents
	for r := 0; r < len(parser.T.Resources); r++ {
		for i := 0; i < len(parser.T.Resources[r].Instances); i++ {
			for d := 0; d < len(parser.T.Resources[r].Instances[i].Dependencies); d++ {
				// save dependency info
				dependency := parser.T.Resources[r].Instances[i].Dependencies[d]
				dependencyName := strings.Split(dependency, ".")
				dependencyIndex := nameToIndex[dependencyName[1]]
				// increment numDependencies and numDependents
				numDependencies[r] += 1
				numDependents[dependencyIndex] += 1
			}
		}
	}

	// iterate through each resource
	for r := 0; r < len(parser.T.Resources); r++ {

		// temp list of dependencies and dependents for current resource
		var tempDependencyNames []string
		var tempDependencyIndices []int
		var tempDependentNames []string
		var tempDependentIndices []int

		// find the name and index of each dependency of the current resource
		if numDependencies[r] > 0 {
			for i := 0; i < len(parser.T.Resources[r].Instances); i++ {
				for d := 0; d < len(parser.T.Resources[r].Instances[i].Dependencies); d++ {
					// save dependency info
					dependency := parser.T.Resources[r].Instances[i].Dependencies[d]
					dependencyName := strings.Split(dependency, ".")
					dependencyIndex := nameToIndex[dependencyName[1]]
					// append dependency
					tempDependencyNames = append(tempDependencyNames, dependencyName[0])
					tempDependencyIndices = append(tempDependencyIndices, dependencyIndex)
				}
			}
			// store dependencies for current resource
			dependencyNames[r] = tempDependencyNames
			dependencyIndices[r] = tempDependencyIndices
		}

		// find the name and index of each dependent of the current resource
		if numDependents[r] > 0 {
			rName := parser.T.Resources[r].Name
			for resource := 0; resource < len(parser.T.Resources); resource++ {
				resourceName := parser.T.Resources[resource].Instances[0].Attributes.Name
				for i := 0; i < len(parser.T.Resources[resource].Instances); i++ {
					for d := 0; d < len(parser.T.Resources[resource].Instances[i].Dependencies); d++ {
						if len(parser.T.Resources[resource].Instances[i].Dependencies) > 0 {
							// save dependent info
							dependency := parser.T.Resources[resource].Instances[i].Dependencies[d]
							dependencyName := strings.Split(dependency, ".")
							if rName == dependencyName[1] {
								// append dependent
								tempDependentNames = append(tempDependentNames, resourceName)
								tempDependentIndices = append(tempDependentIndices, resource)
							}
						}
					}
				}
			}
			// store dependents for resource
			dependentNames[r] = tempDependentNames
			dependentIndices[r] = tempDependentIndices
		}

	}

	// list dependencies of each resource
	for r := 0; r < len(parser.T.Resources); r++ {
		fmt.Print("(", r, ") has ", numDependencies[r], " dependencies:")
		for d := 0; d < len(dependencyIndices[r]); d++ {
			fmt.Print(" (", (dependencyIndices[r])[d], " ", (dependencyNames[r])[d], ")")
		}
		fmt.Println()
	}
	fmt.Println()

	// list dependents of each resource
	for r := 0; r < len(parser.T.Resources); r++ {
		fmt.Print("(", r, ") has ", numDependents[r], " dependents:")
		for d := 0; d < len(dependentIndices[r]); d++ {
			fmt.Print(" (", (dependentIndices[r])[d], " ", (dependentNames[r])[d], ")")
		}
		fmt.Println()
	}
	fmt.Println()
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	/*** CREATE ELEMENT TREE WITH PARSED DATA ***/

	XML.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	mxGraphModel := XML.CreateElement("mxGraphModel")
	mxGraphModel.CreateAttr("gridSize", "10")
	mxGraphModel.CreateAttr("pageWidth", fmt.Sprint(DimX))
	mxGraphModel.CreateAttr("pageHeight", fmt.Sprint(DimY))

	root := mxGraphModel.CreateElement("root")

	mxCell := root.CreateElement("mxCell")
	mxCell.CreateAttr("id", fmt.Sprint(GlobalID))
	GlobalID++

	mxCell = root.CreateElement("mxCell")
	mxCell.CreateAttr("id", fmt.Sprint(GlobalID))
	mxCell.CreateAttr("parent", fmt.Sprint(GlobalID - 1))
	GlobalID++

	/*** CREATING PROJECT REGIONS ***/

	projectX := 20
	projectY := 380
	subX := 5

	// iterate through all resoureces (elements)
	elementID := 0
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
			nameToIndex[parser.T.Resources[r].Name] = elementID
		}

		// if name is network, create project area
		if parser.T.Resources[r].Name == "network" {

			minX := projectX
			minY := 60
			maxX := 350
			maxY := projectY

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(GlobalID))
			mxCell.CreateAttr("parent", fmt.Sprint(1))
			GlobalID = GlobalID + 1

			if len(resourceName) > 0 {
				mxCell.CreateAttr("value", fmt.Sprintf("%s	%s", resourceName, resourceType))
			} else {
				mxCell.CreateAttr("value", resourceType)
			}

			mxCell.CreateAttr("style", fmt.Sprint("whiteSpace=wrap;sketch=0;points=[[0,0,0],[0.25,0,0],[0.5,0,0],[0.75,0,0],[1,0,0],[1,0.25,0],[1,0.5,0],[1,0.75,0],[1,1,0],[0.75,1,0],[0.5,1,0],[0.25,1,0],[0,1,0],[0,0.75,0],[0,0.5,0],[0,0.25,0]];rounded=1;absoluteArcSize=1;arcSize=2;html=1;strokeColor=none;gradientColor=none;shadow=0;dashed=0;fontSize=12;fontColor=#9E9E9E;align=left;verticalAlign=top;spacing=10;spacingTop=-4;" + objectShape))
			mxCell.CreateAttr("vertex", "1")

			// set current elements location based off grid (x, y) locations
			currentRow, currentCol := len(parser.T.Resources), r
			// SHOULD NOT USE LINE BELOW
			currentRow = 1
			xLocation, yLocation := grid[(len(parser.T.Resources) * currentRow) + currentCol].x, grid[(len(parser.T.Resources) * currentRow) + currentCol].y

			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", fmt.Sprint(xLocation - minX))
			mxGeometry.CreateAttr("y", fmt.Sprint(yLocation - minY))
			mxGeometry.CreateAttr("width", fmt.Sprint(maxX))
			mxGeometry.CreateAttr("height", fmt.Sprint(maxY))
			mxGeometry.CreateAttr("as", "geometry")

			tmp := new(TerraNavigator)
			tmp.Name = parser.T.Resources[r].Name
			tmp.HiddenId = GlobalID - 2
			tmp.XPosCenter = minX + (maxX / 2)
			tmp.YPosCenter = minY + (maxY / 2)
			tmp.Width = maxX
			tmp.Height = maxY
			tmp.Project = parser.T.Resources[r].Instances[0].Attributes.Project
			tmp.ObjectShape = objectShape
			Elements = append(Elements, *tmp)

			projectX = projectX + 550
		}

		// if name is subnetwork, create project area
		if parser.T.Resources[r].Name == "subnetwork" {

			minX := projectX - subX
			minY := 30
			maxX := 330
			maxY := projectY - 40

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(GlobalID))
			mxCell.CreateAttr("parent", fmt.Sprint(1))
			GlobalID = GlobalID + 1

			if len(resourceName) > 0 {
				mxCell.CreateAttr("value", fmt.Sprintf("%s	%s", resourceName, resourceType))
			} else {
				mxCell.CreateAttr("value", resourceType)
			}

			mxCell.CreateAttr("style", fmt.Sprint("whiteSpace=wrap;sketch=0;points=[[0,0,0],[0.25,0,0],[0.5,0,0],[0.75,0,0],[1,0,0],[1,0.25,0],[1,0.5,0],[1,0.75,0],[1,1,0],[0.75,1,0],[0.5,1,0],[0.25,1,0],[0,1,0],[0,0.75,0],[0,0.5,0],[0,0.25,0]];rounded=1;absoluteArcSize=1;arcSize=2;html=1;strokeColor=none;gradientColor=none;shadow=0;dashed=0;fontSize=12;fontColor=#9E9E9E;align=left;verticalAlign=top;spacing=10;spacingTop=-4;" + objectShape))
			mxCell.CreateAttr("vertex", "1")

			// set current elements location based off grid (x, y) locations
			currentRow, currentCol := len(parser.T.Resources), r
			// SHOULD NOT USE LINE BELOW
			currentRow, currentCol = 1, 0
			xLocation, yLocation := grid[(len(parser.T.Resources) * currentRow) + currentCol].x, grid[(len(parser.T.Resources) * currentRow) + currentCol].y

			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", fmt.Sprint(xLocation - 10))
			mxGeometry.CreateAttr("y", fmt.Sprint(yLocation - minY))
			mxGeometry.CreateAttr("width", fmt.Sprint(maxX))
			mxGeometry.CreateAttr("height", fmt.Sprint(maxY))
			mxGeometry.CreateAttr("as", "geometry")

			tmp := new(TerraNavigator)
			tmp.Name = parser.T.Resources[r].Name
			tmp.HiddenId = GlobalID - 2
			tmp.XPosCenter = minX + (maxX / 2)
			tmp.YPosCenter = minY + (maxY / 2)
			tmp.Width = maxX
			tmp.Height = maxY
			tmp.Project = parser.T.Resources[r].Instances[0].Attributes.Project
			tmp.ObjectShape = objectShape
			Elements = append(Elements, *tmp)
		}
	}

	rowOffset := -1

	/*** ITERATE THROUGH RESOURCES ***/

	for i := 0; i < len(parser.T.Resources); i++ {

		// (1) store resource type (ex: google_storage_bucket)
		resourceType := parser.T.Resources[i].Type

		if parser.T.Resources[i].Name != "default" {
			// store the name and id of dependency elements
			nameToIndex[parser.T.Resources[i].Name] = elementID
		}

		// (2) use resource type to lookup the draw.io name (ex: Bucket)
		objectName := utility.LookupName(resourceType)

		// (3) use object name to lookup the draw.io shape (ex: shape=mxgraph.gcp2.bucket)
		objectShape := utility.LookupShape(objectName)

		// (4) use object name to lookup the correct case of creating the draw.io shape (ex: 1)
		t := utility.LookupCase(objectName)

		// (5) use specific resource name for main text (ex: example-storage-bucket)
		resourceName := parser.T.Resources[i].Instances[0].Attributes.Name

		// if network or subnetwork, skip and tick rowOffset
		if parser.T.Resources[i].Name == "network" || parser.T.Resources[i].Name == "subnetwork" {
			rowOffset++
			continue
		}

		// set current elements location based off grid (x, y) locations
		currentRow, currentCol := i, numDependents[i]
		xLocation, yLocation := grid[(len(parser.T.Resources) * (currentRow - rowOffset)) + currentCol].x, grid[(len(parser.T.Resources) * (currentRow - rowOffset)) + currentCol].y

		/*** DETERMINE WHICH XML STRUCTURE IS NEEDED ***/

		switch t {

		/*** GCP / PATHS ***/

		case 0:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(GlobalID))
			mxCell.CreateAttr("parent", fmt.Sprint(GlobalID - 1))
			GlobalID++

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
			mxCell.CreateAttr("id", fmt.Sprint(GlobalID))
			mxCell.CreateAttr("parent", fmt.Sprint(1))
			GlobalID++

			mxCell.CreateAttr("value", "")
			mxCell.CreateAttr("style", "whiteSpace=wrap;html=1;strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;")
			mxCell.CreateAttr("vertex", "1")

			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", fmt.Sprint(xLocation))
			mxGeometry.CreateAttr("y", fmt.Sprint(yLocation))
			mxGeometry.CreateAttr("width", fmt.Sprint(CardWidth))
			mxGeometry.CreateAttr("height", fmt.Sprint(CardHeight))
			mxGeometry.CreateAttr("as", "geometry")

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(GlobalID))
			mxCell.CreateAttr("parent", fmt.Sprint(GlobalID - 1))
			GlobalID++

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

			tmp := new(TerraNavigator)
			tmp.Name = parser.T.Resources[i].Name
			tmp.HiddenId = GlobalID - 2
			tmp.XPosCenter = xLocation + (CardWidth / 2)
			tmp.YPosCenter = yLocation + (CardHeight / 2)
			tmp.Width = CardWidth
			tmp.Height = CardHeight
			tmp.ObjectShape = objectShape
			Elements = append(Elements, *tmp)

		/****************************************************************************************************/

		/*** GCP / USER AND DEVICE CARDS ***/

		case 2:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(GlobalID))
			mxCell.CreateAttr("parent", fmt.Sprint(1))
			GlobalID++

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
			mxGeometry.CreateAttr("width", fmt.Sprint(CardWidth))
			mxGeometry.CreateAttr("height", fmt.Sprint(CardHeight))
			mxGeometry.CreateAttr("as", "geometry")

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(GlobalID))
			mxCell.CreateAttr("parent", fmt.Sprint(GlobalID - 1))
			GlobalID++

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

			tmp := new(TerraNavigator)
			tmp.Name = parser.T.Resources[i].Name
			tmp.HiddenId = GlobalID - 2
			tmp.XPosCenter = xLocation + (CardWidth / 2)
			tmp.YPosCenter = yLocation + (CardHeight / 2)
			tmp.Width = CardWidth
			tmp.Height = CardHeight
			tmp.ObjectShape = objectShape
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
			mxCell.CreateAttr("id", fmt.Sprint(GlobalID))
			mxCell.CreateAttr("parent", fmt.Sprint(1))
			GlobalID++

			mxCell.CreateAttr("value", "")
			mxCell.CreateAttr("style", "whiteSpace=wrap;html=1;strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;")
			mxCell.CreateAttr("vertex", "1")

			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", fmt.Sprint(xLocation))
			mxGeometry.CreateAttr("y", fmt.Sprint(yLocation))
			mxGeometry.CreateAttr("width", fmt.Sprint(CardWidth))
			mxGeometry.CreateAttr("height", fmt.Sprint(CardHeight))
			mxGeometry.CreateAttr("as", "geometry")

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(GlobalID))
			mxCell.CreateAttr("parent", fmt.Sprint(GlobalID - 1))
			GlobalID++

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

			tmp := new(TerraNavigator)
			tmp.Name = parser.T.Resources[i].Name
			tmp.HiddenId = GlobalID - 2
			tmp.XPosCenter = xLocation + (CardWidth / 2)
			tmp.YPosCenter = yLocation + (CardHeight / 2)
			tmp.Width = CardWidth
			tmp.Height = CardHeight
			tmp.ObjectShape = objectShape
			Elements = append(Elements, *tmp)

		/****************************************************************************************************/

		/*** GCP / PRODUCT CARDS ***/

		case 4:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(GlobalID))
			mxCell.CreateAttr("parent", fmt.Sprint(1))
			GlobalID++

			mxCell.CreateAttr("value", "")
			mxCell.CreateAttr("style", "whiteSpace=wrap;html=1;strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;")
			mxCell.CreateAttr("vertex", "1")

			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", fmt.Sprint(xLocation))
			mxGeometry.CreateAttr("y", fmt.Sprint(yLocation))
			mxGeometry.CreateAttr("width", fmt.Sprint(CardWidth))
			mxGeometry.CreateAttr("height", fmt.Sprint(CardHeight))
			mxGeometry.CreateAttr("as", "geometry")

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(GlobalID))
			mxCell.CreateAttr("parent", fmt.Sprint(GlobalID - 1))
			GlobalID++

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

			tmp := new(TerraNavigator)
			tmp.Name = parser.T.Resources[i].Name
			tmp.HiddenId = GlobalID - 2
			tmp.XPosCenter = xLocation + (CardWidth / 2)
			tmp.YPosCenter = yLocation + (CardHeight / 2)
			tmp.Width = CardWidth
			tmp.Height = CardHeight
			tmp.ObjectShape = objectShape
			Elements = append(Elements, *tmp)

		/****************************************************************************************************/

		/*** GCP ICONS ***/

		case 5:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(GlobalID))
			mxCell.CreateAttr("parent", fmt.Sprint(1))
			GlobalID++

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
			mxGeometry.CreateAttr("width", fmt.Sprint(CardWidth))
			mxGeometry.CreateAttr("height", fmt.Sprint(CardHeight))
			mxGeometry.CreateAttr("as", "geometry")

			tmp := new(TerraNavigator)
			tmp.Name = parser.T.Resources[i].Name
			tmp.HiddenId = GlobalID - 2
			tmp.XPosCenter = xLocation + (CardWidth / 2)
			tmp.YPosCenter = yLocation + (CardHeight / 2)
			tmp.Width = CardWidth
			tmp.Height = CardHeight
			tmp.ObjectShape = objectShape
			Elements = append(Elements, *tmp)

		case 6: // Cloud Scheduler

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(GlobalID))
			mxCell.CreateAttr("parent", fmt.Sprint(1))
			GlobalID++

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
			mxGeometry.CreateAttr("width", fmt.Sprint(CardWidth))
			mxGeometry.CreateAttr("height", fmt.Sprint(CardHeight))
			mxGeometry.CreateAttr("as", "geometry")

			tmp := new(TerraNavigator)
			tmp.Name = parser.T.Resources[i].Name
			tmp.HiddenId = GlobalID - 2
			tmp.XPosCenter = xLocation + (CardWidth / 2)
			tmp.YPosCenter = yLocation + (CardHeight / 2)
			tmp.Width = CardWidth
			tmp.Height = CardHeight
			tmp.ObjectShape = objectShape
			Elements = append(Elements, *tmp)

		/****************************************************************************************************/

		/*** GCP / ZONES ***/

		case 7:

			fmt.Println("hits for", parser.T.Resources[i].Name)
			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(GlobalID))
			mxCell.CreateAttr("parent", fmt.Sprint(1))
			GlobalID++

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
			mxGeometry.CreateAttr("width", fmt.Sprint(CardWidth))
			mxGeometry.CreateAttr("height", fmt.Sprint(CardHeight))
			mxGeometry.CreateAttr("as", "geometry")

			tmp := new(TerraNavigator)
			tmp.Name = parser.T.Resources[i].Name
			tmp.HiddenId = GlobalID - 2
			tmp.XPosCenter = xLocation + (CardWidth / 2)
			tmp.YPosCenter = yLocation + (CardHeight / 2)
			tmp.Width = CardWidth
			tmp.Height = CardHeight
			tmp.ObjectShape = objectShape
			Elements = append(Elements, *tmp)

		case 8:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", fmt.Sprint(GlobalID))
			mxCell.CreateAttr("parent", fmt.Sprint(1))
			GlobalID++

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

	/*** USE DEPENDENCIES TO CREATE ARROWS ***/

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
						mxCell.CreateAttr("id", fmt.Sprint(GlobalID))
						mxCell.CreateAttr("parent", fmt.Sprint(1))
						GlobalID++

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

						tmp := new(RelationNavigator)
						tmp.ArrowID = GlobalID - 1
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

	/*** PRINT TO THE outFileLocation FILE ***/

	XML.Indent(4)
	XML.WriteToFile(outFileLocation)

	// close file
	outFile.Close()
}
