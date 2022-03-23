package validator

import (
	"fmt"
	"log"			// logging errors
	"os"			// create and open files
	"bufio"			// scanning files

	"KSCD/mapper"

	// creating xml file (go get github.com/beevik/etree)
	"github.com/beevik/etree")

var xml = etree.NewDocument()

func Validator() {

	/*** OPEN THE terraform.drawio FILE ***/

	inFile, errRead := os.Open("terraform.drawio")
	// error reading file
	if errRead != nil {
		log.Println("Error opening file.", errRead)
		os.Exit(1)
	}

	if err := xml.ReadFromFile("terraform.drawio"); err != nil {
		panic(err)
	}

	// keep file open
	defer inFile.Close()

	/*** SCAN THE terraform.drawio FILE (line by line) ***/

	scanner := bufio.NewScanner(inFile)
	// error scanning file
	errScan := scanner.Err()
	if errScan != nil {
		log.Println("Error scanning file.", errRead)
		os.Exit(1)
	}

	// Checking for arrows overlapping boxes
	// Goes through all the arrows existing
	for _, arrow := range mapper.Arrows {

		var newX int

		// Loops through Elements to get the slices //?????
		for _, slice := range mapper.Elements {

			/*** VERTICAL ARROWS ***/

			if (arrow.XPosSource == arrow.XPosTarget) {

				if ((arrow.YPosSource - (slice.Height * 2) == arrow.YPosTarget) || (arrow.YPosSource + (slice.Height * 2) == arrow.YPosTarget)) {

					// NO BENDING, target is directly above / below source
					// [source]--[target]

				} else if (arrow.YPosSource - (slice.Height * 2) > arrow.YPosTarget) || (arrow.YPosSource + (slice.Height * 2) < arrow.YPosTarget) {

					// NEED BENDING, target is not directly above / below source
					// [source]-x-[target]

					fmt.Println(arrow.XPosSource, arrow.XPosTarget)
					if (arrow.XPosSource == 50 + slice.Width / 2) {
						// left row, bend left
						newX = arrow.XPosSource - slice.Width / 2 - 25
					} else if (arrow.XPosSource == 50 + slice.Width / 2 + (slice.Width * 2)) {
						// right row, bend right
						newX = arrow.XPosSource + slice.Width / 2 + 25
					}

					// XML for creating bends
					path := fmt.Sprintf("/mxGraphModel/root/mxCell[%d]/mxGeometry", arrow.ArrowID + 1)
					arrowGeom := xml.FindElement(path)

					array := arrowGeom.CreateElement("Array")
					array.CreateAttr("as", "points")

					mxPoint := array.CreateElement("mxPoint")
					mxPoint.CreateAttr("x", fmt.Sprint(newX))
					mxPoint.CreateAttr("y", fmt.Sprint(arrow.YPosSource))

					mxPoint = array.CreateElement("mxPoint")
					mxPoint.CreateAttr("x", fmt.Sprint(newX))
					mxPoint.CreateAttr("y", fmt.Sprint(arrow.YPosTarget))

				}
			}

			/*** HORIZONTAL ARROWS ***/

			if (arrow.YPosSource == arrow.YPosTarget) {

				// NO BENDING, target is directly left / right of source
				// [source]--[target]

			}

			/*** DIAGONAL ARROWS ***/

			if ((arrow.XPosSource != arrow.XPosTarget) && (arrow.YPosSource != arrow.YPosTarget)) {

				// NEED BENDING, target not directly above / below / left / right of source
				// [source]-x-[target]

				newX = (arrow.XPosSource + arrow.XPosTarget) / 2

				// XML for creating bends
				path := fmt.Sprintf("/mxGraphModel/root/mxCell[%d]/mxGeometry", arrow.ArrowID + 1)
				arrowGeom := xml.FindElement(path)

				array := arrowGeom.CreateElement("Array")
				array.CreateAttr("as", "points")

				mxPoint := array.CreateElement("mxPoint")
				mxPoint.CreateAttr("x", fmt.Sprint(newX))
				mxPoint.CreateAttr("y", fmt.Sprint(arrow.YPosSource))

				mxPoint = array.CreateElement("mxPoint")
				mxPoint.CreateAttr("x", fmt.Sprint(newX))
				mxPoint.CreateAttr("y", fmt.Sprint(arrow.YPosTarget))

			}

		}
	
	}

	// Writing back to terraform.drawio
	xml.Indent(4)
	xml.WriteToFile("terraform.drawio")
	inFile.Close()

}
