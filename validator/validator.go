package validator

import (
	"fmt"
	"log"			// logging errors
	"os"			// create and open files
	"bufio"			// scanning files

	"KSCD/mapper"

	// creating xml file (go get github.com/beevik/etree)
	"github.com/beevik/etree"
)

/*** CREATE GLOBAL XML TREE ***/

var XML = etree.NewDocument()

func Validator(outFileLocation string) {

	/*** OPEN THE outFileLocation FILE ***/

	inFile, errRead := os.Open(outFileLocation)
	// error reading file
	if errRead != nil {
		log.Println("Error opening file.", errRead)
		os.Exit(1)
	}

	if err := XML.ReadFromFile(outFileLocation); err != nil {
		panic(err)
	}
	// keep open
	defer inFile.Close()

	/*** SCAN THE outFileLocation FILE (line by line) ***/

	scanner := bufio.NewScanner(inFile)
	// error scanning file
	errScan := scanner.Err()
	if errScan != nil {
		log.Println("Error scanning file.", errRead)
		os.Exit(1)
	}

	/*** UPDATE THE outFileLocation FILE ***/

	removeInvalidShapes()

	/*** ARROW COLLISION ***/

	// iterate through all arrows and check for overlap
	for _, arrow := range mapper.Arrows {

		var xBend int

		// iterate through all elements to get the slices
		for _, slice := range mapper.Elements {

			/*** VERTICAL ARROWS ***/

			if arrow.XPosSource == arrow.XPosTarget {
				if (arrow.YPosSource - (slice.Height * 2) == arrow.YPosTarget) || (arrow.YPosSource + (slice.Height * 2) == arrow.YPosTarget) {
					// NO BENDING, target is directly above / below source
					// [source]--[target]
				} else if (arrow.YPosSource - (slice.Height * 2) > arrow.YPosTarget) || (arrow.YPosSource + (slice.Height * 2) < arrow.YPosTarget) {
					// NEED BENDING, target is not directly above / below source
					// [source]-x-[target]
					if arrow.XPosSource == 50 + slice.Width / 2 {
						// left row, bend left
						xBend = arrow.XPosSource - slice.Width / 2 - 25
					} else if arrow.XPosSource == 50 + slice.Width / 2 + (slice.Width * 2) {
						// right row, bend right
						xBend = arrow.XPosSource + slice.Width / 2 + 25
					}
					createArrowBend(arrow.ArrowID, xBend, arrow.YPosSource, xBend, arrow.YPosTarget)
				}
			}

			/*** HORIZONTAL ARROWS ***/

			if arrow.YPosSource == arrow.YPosTarget {
				// NO BENDING, target is directly left / right of source
				// [source]--[target]
			}

			/*** DIAGONAL ARROWS ***/

			if (arrow.XPosSource != arrow.XPosTarget) && (arrow.YPosSource != arrow.YPosTarget) {
				// NEED BENDING, target not directly above / below / left / right of source
				// [source]-x-[target]
				xBend = (arrow.XPosSource + arrow.XPosTarget) / 2
				createArrowBend(arrow.ArrowID, xBend, arrow.YPosSource, xBend, arrow.YPosTarget)
			}

		}

	}

	/*** WRITE BACK TO THE outFileLocation FILE ***/

	XML.Indent(4)
	XML.WriteToFile(outFileLocation)

	// close file
	inFile.Close()

}

/*** FUNCTIONS ***/

func removeInvalidShapes() {
	// iterate through all elements
	for _, slice := range mapper.Elements {
		// remove any invalid shapes from the xml tree
		if slice.ObjectShape == "shape=mxgraph.gcp2.blank" {
			path := fmt.Sprintf("/mxGraphModel/root/mxCell[%d]", slice.HiddenId + 2)
			removeElement := XML.FindElement(path)
			removeParent := removeElement.Parent()
			removeParent.RemoveChildAt(slice.HiddenId + 10)
			removeParent.RemoveChildAt(slice.HiddenId + 8)
			fmt.Printf("ERROR: Element %s has either an invalid shape or is not implemented yet by terraref. Removing...\n", slice.Name)
		}
	}
}

func createArrowBend(id int, xSource int, ySource int, xTarget int, yTarget int) {
	path := fmt.Sprintf("/mxGraphModel/root/mxCell[%d]/mxGeometry", id + 1)
	arrowGeom := XML.FindElement(path)
	array := arrowGeom.CreateElement("Array")
	array.CreateAttr("as", "points")
	mxPoint := array.CreateElement("mxPoint")
	mxPoint.CreateAttr("x", fmt.Sprint(xSource))
	mxPoint.CreateAttr("y", fmt.Sprint(ySource))
	mxPoint = array.CreateElement("mxPoint")
	mxPoint.CreateAttr("x", fmt.Sprint(xTarget))
	mxPoint.CreateAttr("y", fmt.Sprint(yTarget))
}
