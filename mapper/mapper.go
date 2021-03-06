package mapper

import (
	"KSCD/libraries/providers/GCP/utility"
	"KSCD/parser"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/beevik/etree"
)

/*** GLOBAL STRUCT TO STORE AND ACCESS GRID LOCATIONS ***/
type Location struct {
	x, y int
}

var Grid []Location

/*** GLOBAL SLICES FOR ELEMENTS AND ARROWS ***/
var Elements []TerraNavigator
var Arrows []RelationNavigator

/*** GLOBAL STRUCTS TO STORE AND ACCESS INFO ABOUT ELEMENTS AND ARROWS ***/
var TerraNav TerraNavigator

type TerraNavigator struct {
	Name                   string
	Project                string
	ObjectShape            string
	HiddenId               int
	XPosCenter, YPosCenter int
	Width, Height          int
}

type RelationNavigator struct {
	ArrowID                int
	SourceID               int
	XPosSource, YPosSource int
	TargetID               int
	XPosTarget, YPosTarget int
}

/*** CREATE GLOBAL XML TREE ***/
var XML = etree.NewDocument()
var MXGraphModel, Root, MXCell, MXGeometry, MXPoint *etree.Element

/*** DEFAULT DIMENSIONS ***/
var DimX, DimY = 850, 1100           // diagram
var CardWidth, CardHeight = 250, 60  // normal cards
var ZoneWidth, ZoneHeight = 350, 380 // project zones

var GlobalID = 2

func Mapper(outFileLocation string) {
	outFile, errCreate := os.Create(outFileLocation)

	// File creation error
	if errCreate != nil {
		log.Println("We weren't able to create an output file named "+outFileLocation+" in mapper.go. Terminating...", errCreate)
		os.Exit(1)
	}

	defer outFile.Close()

	/*** CREATE THE GRID FOR PLACING ELEMENTS ***/
	fmt.Println()
	fmt.Println("****************************************************************************************************")
	fmt.Println("*                                   G R I D    L O C A T I O N S                                   *")
	fmt.Println("****************************************************************************************************")
	fmt.Println()
	createGrid()
	printGrid()

	/*** CREATE DEFAULT ELEMENTS ***/
	XML.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	MXGraphModel = XML.CreateElement("mxGraphModel")
	MXGraphModel.CreateAttr("gridSize", "10")
	MXGraphModel.CreateAttr("pageWidth", fmt.Sprint(DimX))
	MXGraphModel.CreateAttr("pageHeight", fmt.Sprint(DimY))

	Root = MXGraphModel.CreateElement("root")

	MXCell = Root.CreateElement("mxCell")
	MXCell.CreateAttr("id", fmt.Sprint(0))

	MXCell = Root.CreateElement("mxCell")
	MXCell.CreateAttr("id", fmt.Sprint(1))
	MXCell.CreateAttr("parent", fmt.Sprint(0))

	/*** PROJECT REGION SCANNING ***/
	projectX := 20
	projectY := 380
	subX := 5

	elementID := 0
	for r := 0; r < len(parser.T.Resources); r++ {

		resourceType := parser.T.Resources[r].Type
		objectName := utility.LookupName(resourceType)
		objectShape := utility.LookupShape(objectName)
		resourceName := parser.T.Resources[r].Instances[0].Attributes.Name

		if parser.T.Resources[r].Name != "default" {
			parser.NameToIndex[parser.T.Resources[r].Name] = elementID
		}

		if parser.T.Resources[r].Name == "network" {
			minX := projectX
			minY := 60
			maxX := 350
			maxY := projectY

			MXCell = Root.CreateElement("mxCell")
			MXCell.CreateAttr("id", fmt.Sprint(GlobalID))
			MXCell.CreateAttr("parent", fmt.Sprint(1))
			GlobalID = GlobalID + 1

			if len(resourceName) > 0 {
				MXCell.CreateAttr("value", fmt.Sprintf("%s	%s", resourceName, resourceType))
			} else {
				MXCell.CreateAttr("value", resourceType)
			}

			MXCell.CreateAttr("style", fmt.Sprint("whiteSpace=wrap;sketch=0;points=[[0,0,0],[0.25,0,0],[0.5,0,0],[0.75,0,0],[1,0,0],[1,0.25,0],[1,0.5,0],[1,0.75,0],[1,1,0],[0.75,1,0],[0.5,1,0],[0.25,1,0],[0,1,0],[0,0.75,0],[0,0.5,0],[0,0.25,0]];rounded=1;absoluteArcSize=1;arcSize=2;html=1;strokeColor=none;gradientColor=none;shadow=0;dashed=0;fontSize=12;fontColor=#9E9E9E;align=left;verticalAlign=top;spacing=10;spacingTop=-4;"+objectShape))
			MXCell.CreateAttr("vertex", "1")

			// set current elements location based off grid (x, y) locations
			currentRow, currentCol := len(parser.T.Resources), r
			// SHOULD NOT USE LINE BELOW
			currentRow = 1
			xPos, yPos := Grid[(len(parser.T.Resources)*currentRow)+currentCol].x, Grid[(len(parser.T.Resources)*currentRow)+currentCol].y

			MXGeometry = MXCell.CreateElement("mxGeometry")
			MXGeometry.CreateAttr("x", fmt.Sprint(xPos-minX))
			MXGeometry.CreateAttr("y", fmt.Sprint(yPos-minY))
			MXGeometry.CreateAttr("width", fmt.Sprint(maxX))
			MXGeometry.CreateAttr("height", fmt.Sprint(maxY))
			MXGeometry.CreateAttr("as", "geometry")

			zoneTerraNavigator(r, minX, minY, maxX, maxY, objectShape)

			projectX = projectX + 550
		}

		if parser.T.Resources[r].Name == "subnetwork" {

			minX := projectX - subX
			minY := 30
			maxX := 330
			maxY := projectY - 40

			MXCell = Root.CreateElement("mxCell")
			MXCell.CreateAttr("id", fmt.Sprint(GlobalID))
			MXCell.CreateAttr("parent", fmt.Sprint(1))
			GlobalID = GlobalID + 1

			if len(resourceName) > 0 {
				MXCell.CreateAttr("value", fmt.Sprintf("%s	%s", resourceName, resourceType))
			} else {
				MXCell.CreateAttr("value", resourceType)
			}

			MXCell.CreateAttr("style", fmt.Sprint("whiteSpace=wrap;sketch=0;points=[[0,0,0],[0.25,0,0],[0.5,0,0],[0.75,0,0],[1,0,0],[1,0.25,0],[1,0.5,0],[1,0.75,0],[1,1,0],[0.75,1,0],[0.5,1,0],[0.25,1,0],[0,1,0],[0,0.75,0],[0,0.5,0],[0,0.25,0]];rounded=1;absoluteArcSize=1;arcSize=2;html=1;strokeColor=none;gradientColor=none;shadow=0;dashed=0;fontSize=12;fontColor=#9E9E9E;align=left;verticalAlign=top;spacing=10;spacingTop=-4;"+objectShape))
			MXCell.CreateAttr("vertex", "1")

			// set current elements location based off grid (x, y) locations
			currentRow, currentCol := len(parser.T.Resources), r
			// SHOULD NOT USE LINE BELOW
			currentRow, currentCol = 1, 0
			xPos, yPos := Grid[(len(parser.T.Resources)*currentRow)+currentCol].x, Grid[(len(parser.T.Resources)*currentRow)+currentCol].y

			MXGeometry = MXCell.CreateElement("mxGeometry")
			MXGeometry.CreateAttr("x", fmt.Sprint(xPos-10))
			MXGeometry.CreateAttr("y", fmt.Sprint(yPos-minY))
			MXGeometry.CreateAttr("width", fmt.Sprint(maxX))
			MXGeometry.CreateAttr("height", fmt.Sprint(maxY))
			MXGeometry.CreateAttr("as", "geometry")

			zoneTerraNavigator(r, minX, minY, maxX, maxY, objectShape)
		}
	}

	rowOffset := -1

	/*** NORMAL RESOURCE SCANNING ***/
	for i := 0; i < len(parser.T.Resources); i++ {

		resourceType := parser.T.Resources[i].Type

		if parser.T.Resources[i].Name != "default" {
			parser.NameToIndex[parser.T.Resources[i].Name] = elementID
		}

		objectName := utility.LookupName(resourceType)
		objectShape := utility.LookupShape(objectName)
		t := utility.LookupCase(objectName)
		resourceName := parser.T.Resources[i].Instances[0].Attributes.Name

		if parser.T.Resources[i].Name == "network" || parser.T.Resources[i].Name == "subnetwork" {
			rowOffset++
			continue
		}

		currentRow, currentCol := i, parser.NumDependents[i]
		xPos, yPos := Grid[(len(parser.T.Resources)*(currentRow-rowOffset))+currentCol].x, Grid[(len(parser.T.Resources)*(currentRow-rowOffset))+currentCol].y

		/*** DETERMINE WHICH XML STRUCTURE IS NEEDED ***/
		switch t {

		/*** GCP / PATHS ***/
		case 0:

			MXCell = Root.CreateElement("mxCell")
			MXCell.CreateAttr("id", fmt.Sprint(GlobalID))
			MXCell.CreateAttr("parent", fmt.Sprint(GlobalID-1))
			GlobalID++

			MXCell.CreateAttr("value", "")
			MXCell.CreateAttr("style", "whiteSpace=wrap;html=1;edgeStyle=orthogonalEdgeStyle;fontSize=12;html=1;endArrow=blockThin;endFill=1;rounded=0;strokeWidth=2;endSize=4;startSize=4;")
			MXCell.CreateAttr("edge", "1")

			MXGeometry = MXCell.CreateElement("mxGeometry")
			MXGeometry.CreateAttr("relative", "1")
			MXGeometry.CreateAttr("as", "geometry")

			mxPoint := MXGeometry.CreateElement("mxPoint")
			mxPoint.CreateAttr("x", "80")
			mxPoint.CreateAttr("y", "160")
			mxPoint.CreateAttr("as", "sourcePoint")

			mxPoint = MXGeometry.CreateElement("mxPoint")
			mxPoint.CreateAttr("x", "180")
			mxPoint.CreateAttr("y", "160")
			mxPoint.CreateAttr("as", "targetPoint")

		/****************************************************************************************************/

		/*** GCP / SERVICE CARDS ***/
		case 1:

			MXCell = Root.CreateElement("mxCell")
			MXCell.CreateAttr("id", fmt.Sprint(GlobalID))
			MXCell.CreateAttr("parent", fmt.Sprint(1))
			GlobalID++

			MXCell.CreateAttr("value", "")
			MXCell.CreateAttr("style", "whiteSpace=wrap;html=1;strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;")
			MXCell.CreateAttr("vertex", "1")

			MXGeometry = MXCell.CreateElement("mxGeometry")
			MXGeometry.CreateAttr("x", fmt.Sprint(xPos))
			MXGeometry.CreateAttr("y", fmt.Sprint(yPos))
			MXGeometry.CreateAttr("width", fmt.Sprint(CardWidth))
			MXGeometry.CreateAttr("height", fmt.Sprint(CardHeight))
			MXGeometry.CreateAttr("as", "geometry")

			MXCell = Root.CreateElement("mxCell")
			MXCell.CreateAttr("id", fmt.Sprint(GlobalID))
			MXCell.CreateAttr("parent", fmt.Sprint(GlobalID-1))
			GlobalID++

			if len(resourceName) > 0 {
				MXCell.CreateAttr("value", fmt.Sprintf("%s	%s", resourceName, resourceType))
			} else {
				MXCell.CreateAttr("value", resourceType)
			}

			MXCell.CreateAttr("style", fmt.Sprint("whiteSpace=wrap;sketch=0;dashed=0;connectable=0;html=1;fillColor=#757575;strokeColor=none;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontSize=12;"+objectShape))
			MXCell.CreateAttr("vertex", "1")

			MXGeometry = MXCell.CreateElement("mxGeometry")
			MXGeometry.CreateAttr("y", "0.5")
			MXGeometry.CreateAttr("width", "32")
			MXGeometry.CreateAttr("height", "32")
			MXGeometry.CreateAttr("relative", "1")
			MXGeometry.CreateAttr("as", "geometry")

			mxPoint := MXGeometry.CreateElement("mxPoint")
			mxPoint.CreateAttr("x", "5")
			mxPoint.CreateAttr("y", "-16")
			mxPoint.CreateAttr("as", "offset")

			cardTerraNavigator(i, xPos, yPos, objectShape)

		/****************************************************************************************************/

		/*** GCP / USER AND DEVICE CARDS ***/
		case 2:

			style1 := fmt.Sprint("strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;labelPosition=center;verticalLabelPosition=middle;align=center;verticalAlign=bottom;spacingLeft=0;fontColor=#999999;fontSize=12;whiteSpace=wrap;spacingBottom=2;" + objectShape)
			style2 := fmt.Sprint("whiteSpace=wrap;sketch=0;dashed=0;connectable=0;html=1;fillColor=#757575;strokeColor=none;part=1;" + objectShape)
			attr := fmt.Sprint("x")
			var width, height float64 = 50, 50
			var x, y float64 = -25, 15

			elementXML(resourceName, resourceType, xPos, yPos, style1, style2, attr, width, height, x, y)
			cardTerraNavigator(i, xPos, yPos, objectShape)

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

			style1 := fmt.Sprint("whiteSpace=wrap;html=1;strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;")
			style2 := fmt.Sprint("whiteSpace=wrap;sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" + objectShape)
			attr := fmt.Sprint("y")
			var width, height float64 = 44, 39
			var x, y float64 = 5, -19.5

			elementXML(resourceName, resourceType, xPos, yPos, style1, style2, attr, width, height, x, y)
			cardTerraNavigator(i, xPos, yPos, objectShape)

		/****************************************************************************************************/

		/*** GCP / PRODUCT CARDS ***/
		case 4:

			style1 := fmt.Sprint("whiteSpace=wrap;html=1;strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;")
			style2 := fmt.Sprint("whiteSpace=wrap;html=1;sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" + objectShape)
			attr := fmt.Sprint("")
			var width, height float64 = 45, 45
			var x, y float64 = 5, 7

			elementXML(resourceName, resourceType, xPos, yPos, style1, style2, attr, width, height, x, y)
			cardTerraNavigator(i, xPos, yPos, objectShape)

		/****************************************************************************************************/

		/*** GCP ICONS ***/
		case 5:

			MXCell = Root.CreateElement("mxCell")
			MXCell.CreateAttr("id", fmt.Sprint(GlobalID))
			MXCell.CreateAttr("parent", fmt.Sprint(1))
			GlobalID++

			if len(resourceName) > 0 {
				MXCell.CreateAttr("value", fmt.Sprintf("%s	%s", resourceName, resourceType))
			} else {
				MXCell.CreateAttr("value", resourceType)
			}

			MXCell.CreateAttr("style", fmt.Sprint("whiteSpace=wrap;sketch=0;html=1;fillColor=#5184F3;strokeColor=none;verticalAlign=top;labelPosition=center;verticalLabelPosition=bottom;align=center;spacingTop=-6;fontSize=11;fontStyle=1;fontColor=#999999;"+objectShape))
			MXCell.CreateAttr("vertex", "1")

			MXGeometry = MXCell.CreateElement("mxGeometry")
			MXGeometry.CreateAttr("x", fmt.Sprint(xPos))
			MXGeometry.CreateAttr("y", fmt.Sprint(yPos))
			MXGeometry.CreateAttr("width", fmt.Sprint(CardWidth))
			MXGeometry.CreateAttr("height", fmt.Sprint(CardHeight))
			MXGeometry.CreateAttr("as", "geometry")

			cardTerraNavigator(i, xPos, yPos, objectShape)

		case 6: // Cloud Scheduler

			MXCell = Root.CreateElement("mxCell")
			MXCell.CreateAttr("id", fmt.Sprint(GlobalID))
			MXCell.CreateAttr("parent", fmt.Sprint(1))
			GlobalID++

			if len(resourceName) > 0 {
				MXCell.CreateAttr("value", fmt.Sprintf("%s	%s", resourceName, resourceType))
			} else {
				MXCell.CreateAttr("value", resourceType)
			}

			MXCell.CreateAttr("style", fmt.Sprint("whiteSpace=wrap;html=1;fillColor=#5184F3;strokeColor=none;verticalAlign=top;labelPosition=center;verticalLabelPosition=bottom;align=center;fontSize=11;fontStyle=1;fontColor=#999999;"+objectShape))
			MXCell.CreateAttr("vertex", "1")

			MXGeometry = MXCell.CreateElement("mxGeometry")
			MXGeometry.CreateAttr("x", fmt.Sprint(xPos))
			MXGeometry.CreateAttr("y", fmt.Sprint(yPos))
			MXGeometry.CreateAttr("width", fmt.Sprint(CardWidth))
			MXGeometry.CreateAttr("height", fmt.Sprint(CardHeight))
			MXGeometry.CreateAttr("as", "geometry")

			cardTerraNavigator(i, xPos, yPos, objectShape)

		/****************************************************************************************************/

		/*** GCP / ZONES ***/
		case 7:

			fmt.Println("hits for", parser.T.Resources[i].Name)
			MXCell = Root.CreateElement("mxCell")
			MXCell.CreateAttr("id", fmt.Sprint(GlobalID))
			MXCell.CreateAttr("parent", fmt.Sprint(1))
			GlobalID++

			if len(resourceName) > 0 {
				MXCell.CreateAttr("value", fmt.Sprintf("%s	%s", resourceName, resourceType))
			} else {
				MXCell.CreateAttr("value", resourceType)
			}

			MXCell.CreateAttr("style", fmt.Sprint("whiteSpace=wrap;sketch=0;points=[[0,0,0],[0.25,0,0],[0.5,0,0],[0.75,0,0],[1,0,0],[1,0.25,0],[1,0.5,0],[1,0.75,0],[1,1,0],[0.75,1,0],[0.5,1,0],[0.25,1,0],[0,1,0],[0,0.75,0],[0,0.5,0],[0,0.25,0]];rounded=1;absoluteArcSize=1;arcSize=2;html=1;strokeColor=none;gradientColor=none;shadow=0;dashed=0;fontSize=12;fontColor=#9E9E9E;align=left;verticalAlign=top;spacing=10;spacingTop=-4;"+objectShape))
			MXCell.CreateAttr("vertex", "1")

			MXGeometry = MXCell.CreateElement("mxGeometry")
			MXGeometry.CreateAttr("x", fmt.Sprint(xPos))
			MXGeometry.CreateAttr("y", fmt.Sprint(yPos))
			MXGeometry.CreateAttr("width", fmt.Sprint(CardWidth))
			MXGeometry.CreateAttr("height", fmt.Sprint(CardHeight))
			MXGeometry.CreateAttr("as", "geometry")

			cardTerraNavigator(i, xPos, yPos, objectShape)

		case 8:

			MXCell = Root.CreateElement("mxCell")
			MXCell.CreateAttr("id", fmt.Sprint(GlobalID))
			MXCell.CreateAttr("parent", fmt.Sprint(1))
			GlobalID++

			MXCell.CreateAttr("value", resourceType)
			MXCell.CreateAttr("vertex", fmt.Sprint(1))
			MXCell.CreateAttr("style", fmt.Sprint(utility.LookupZone(parser.T.Resources[i].Name)))

			MXGeometry = MXCell.CreateElement("mxGeometry")
			MXGeometry.CreateAttr("x", "160")
			MXGeometry.CreateAttr("y", "120")
			MXGeometry.CreateAttr("width", "160")
			MXGeometry.CreateAttr("height", "120")
			MXGeometry.CreateAttr("as", "geometry")

		/****************************************************************************************************/

		/*** GCP / EXPANDED PRODUCT CARDS ***/

		// Unimplemented

		/****************************************************************************************************/

		/*** GCP / GENERAL ICONS ***/

		// Unimplemented

		/****************************************************************************************************/

		// Error case
		default:
			log.Println("Error: No match.", errCreate)
			os.Exit(1)
		}

		elementID++
	}

	/*** ARROW CREATION ***/
	for r := 0; r < len(parser.T.Resources); r++ {
		for i := 0; i < len(parser.T.Resources[r].Instances); i++ {
			for d := 0; d < len(parser.T.Resources[r].Instances[i].Dependencies); d++ {
				resourceName := parser.T.Resources[r].Instances[i].Dependencies[d]
				dependencyName := strings.Split(resourceName, ".")

				ctr := 0
				for range Elements {
					if Elements[ctr].Name == dependencyName[1] {
						createArrow(r, ctr)
						cardRelationNavigator(r, ctr)
					}
					ctr++
				}
			}
		}
	}

	// Writing to file and close
	XML.Indent(4)
	XML.WriteToFile(outFileLocation)
	outFile.Close()

}

/*** FUNCTIONS ***/

// Creates and stores initial grid locations
func createGrid() {
	rows, cols := len(parser.T.Resources)+1, len(parser.T.Resources)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			tempX, tempY := 50+(CardWidth*2*c), 50+(CardHeight*2*r)
			Grid = append(Grid, Location{tempX, tempY})
		}
	}
}

// Prints grid location to console
func printGrid() {
	rows, cols := len(parser.T.Resources)+1, len(parser.T.Resources)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			fmt.Print("(", Grid[(r*len(parser.T.Resources))+c].x, ", ", Grid[(r*len(parser.T.Resources))+c].y, ")", "\t")
		}
		fmt.Println()
	}
	fmt.Println()
}

// General resource logic condensed to a single functions
func elementXML(resourceName string, resourceType string, xPos int, yPos int, style1 string, style2 string, attr string, width float64, height float64, x float64, y float64) {

	MXCell = Root.CreateElement("mxCell")
	MXCell.CreateAttr("id", fmt.Sprint(GlobalID))
	MXCell.CreateAttr("parent", fmt.Sprint(1))
	GlobalID++

	MXCell.CreateAttr("value", "")
	MXCell.CreateAttr("style", style1)
	MXCell.CreateAttr("vertex", "1")

	MXGeometry = MXCell.CreateElement("mxGeometry")
	MXGeometry.CreateAttr("x", fmt.Sprint(xPos))
	MXGeometry.CreateAttr("y", fmt.Sprint(yPos))
	MXGeometry.CreateAttr("width", fmt.Sprint(CardWidth))
	MXGeometry.CreateAttr("height", fmt.Sprint(CardHeight))
	MXGeometry.CreateAttr("as", "geometry")

	MXCell = Root.CreateElement("mxCell")
	MXCell.CreateAttr("id", fmt.Sprint(GlobalID))
	MXCell.CreateAttr("parent", fmt.Sprint(GlobalID-1))
	GlobalID++

	if len(resourceName) > 0 {
		MXCell.CreateAttr("value", fmt.Sprintf("%s	%s", resourceName, resourceType))
	} else {
		MXCell.CreateAttr("value", resourceType)
	}

	MXCell.CreateAttr("style", style2)
	MXCell.CreateAttr("vertex", "1")

	MXGeometry = MXCell.CreateElement("mxGeometry")

	if attr != "" {
		MXGeometry.CreateAttr(fmt.Sprint(attr), fmt.Sprint(0.5))
	}

	MXGeometry.CreateAttr("width", fmt.Sprint(width))
	MXGeometry.CreateAttr("height", fmt.Sprint(height))
	MXGeometry.CreateAttr("relative", "1")
	MXGeometry.CreateAttr("as", "geometry")

	mxPoint := MXGeometry.CreateElement("mxPoint")
	mxPoint.CreateAttr("x", fmt.Sprint(x))
	mxPoint.CreateAttr("y", fmt.Sprint(y))
	mxPoint.CreateAttr("as", "offset")

}

// Writes element into TerraNavigator Struct
func cardTerraNavigator(index int, xPos int, yPos int, object string) {
	tmp := new(TerraNavigator)
	tmp.Name = parser.T.Resources[index].Name
	tmp.HiddenId = GlobalID - 2
	tmp.XPosCenter = xPos + (CardWidth / 2)
	tmp.YPosCenter = yPos + (CardHeight / 2)
	tmp.Width = CardWidth
	tmp.Height = CardHeight
	tmp.ObjectShape = object
	Elements = append(Elements, *tmp)
}

// Writes zone into TerraNavigator Struct
func zoneTerraNavigator(index int, minX int, minY int, maxX int, maxY int, object string) {
	tmp := new(TerraNavigator)
	tmp.Name = parser.T.Resources[index].Name
	tmp.HiddenId = GlobalID - 2
	tmp.XPosCenter = minX + (maxX / 2)
	tmp.YPosCenter = minY + (maxY / 2)
	tmp.Width = maxX
	tmp.Height = maxY
	tmp.Project = parser.T.Resources[index].Instances[0].Attributes.Project
	tmp.ObjectShape = object
	Elements = append(Elements, *tmp)
}

// Writes card into TerraNavigator Struct
func cardRelationNavigator(resource int, counter int) {
	tmp := new(RelationNavigator)
	tmp.ArrowID = GlobalID - 1
	tmp.SourceID = Elements[resource].HiddenId
	tmp.XPosSource = Elements[resource].XPosCenter
	tmp.YPosSource = Elements[resource].YPosCenter
	tmp.TargetID = Elements[counter].HiddenId
	tmp.XPosTarget = Elements[counter].XPosCenter
	tmp.YPosTarget = Elements[counter].YPosCenter
	Arrows = append(Arrows, *tmp)
}

// Creates arrow
func createArrow(resource int, counter int) {

	MXCell = Root.CreateElement("mxCell")
	MXCell.CreateAttr("id", fmt.Sprint(GlobalID))
	MXCell.CreateAttr("parent", fmt.Sprint(1))
	GlobalID++

	MXCell.CreateAttr("value", "")
	MXCell.CreateAttr("style", "whiteSpace=wrap;html=1;edgeStyle=orthogonalEdgeStyle;fontSize=12;html=1;endArrow=blockThin;endFill=1;rounded=0;strokeWidth=2;endSize=4;startSize=4;")
	MXCell.CreateAttr("edge", "1")
	MXCell.CreateAttr("target", fmt.Sprintf("%d", Elements[counter].HiddenId))
	MXCell.CreateAttr("source", fmt.Sprintf("%d", Elements[resource].HiddenId))

	MXGeometry = MXCell.CreateElement("mxGeometry")
	MXGeometry.CreateAttr("relative", "1")
	MXGeometry.CreateAttr("as", "geometry")

	mxPoint := MXGeometry.CreateElement("mxPoint")
	mxPoint.CreateAttr("x", fmt.Sprint(Elements[resource].XPosCenter))
	mxPoint.CreateAttr("y", fmt.Sprint(Elements[resource].YPosCenter))
	mxPoint.CreateAttr("as", "sourcePoint")

	mxPoint = MXGeometry.CreateElement("mxPoint")
	mxPoint.CreateAttr("x", fmt.Sprint(Elements[counter].XPosCenter))
	mxPoint.CreateAttr("y", fmt.Sprint(Elements[counter].YPosCenter))
	mxPoint.CreateAttr("as", "targetPoint")

}
