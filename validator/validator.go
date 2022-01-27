package validator

import (
	"bufio" // scanning files
	"fmt"
	"log" // logging errors
	"os"  // create and open files

	"KSCD/mapper" // Mapper File

	"github.com/beevik/etree" // creating xml file (go get github.com/beevik/etree)
)

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
	for _, arrow := range mapper.ArrowRelationships {

		var ArrowsOverlap bool = false
		var Ymatch bool

		if arrow.YPosSource == arrow.YPosTarget {
			Ymatch = true
		} else {
			Ymatch = false
		}

		// Loops through Pizza to get the slices
		for _, pizza := range mapper.Pizza {
			fmt.Println(pizza)

			// Check if arrows overlap boxes at all

			// Should not be equal to either options
			if arrow.SourceID != pizza.HiddenId && arrow.TargetID != arrow.SourceID {
				if Ymatch == true {

				}
			}

			// Check if arrow has already been resized

			// If the arrow does overlap AND it has not run through this function yet, create it's bending array (NOTE: Functions below are being performed above due to need to go through pizza function)
			if ArrowsOverlap == true {

				var newX int
				var newY int
				var stillOverlaps bool = true

				// Finding an X or Y Coordinate that is unoccupied
				// Checking first if we need to move along x or y axis
				if arrow.YPosSource == arrow.YPosTarget {
					for stillOverlaps == true {
						newX = arrow.XPosSource + pizza.Width
						newY = arrow.YPosSource

						for _, slice := range mapper.Pizza {
							fmt.Println(slice)

						}

					}

					// Check if still overlaps

				} else {
					for stillOverlaps == true {
						newY = arrow.YPosSource + pizza.Height
						newX = arrow.XPosSource

						for _, slice := range mapper.Pizza {
							fmt.Println(slice)

						}

					}
				}

				// Writing of necessary XML code to create bends
				path := fmt.Sprintf("/mxGraphModel/root/mxCell[%d]/mxGeometry", arrow.ArrowID+1)
				arrowGeom := xml.FindElement(path)

				array := arrowGeom.CreateElement("Array")
				array.CreateAttr("as", "points")

				mxPoint := array.CreateElement("mxPoint")
				mxPoint.CreateAttr("x", fmt.Sprint(newX))
				mxPoint.CreateAttr("y", fmt.Sprint(newY))

				mxPoint = array.CreateElement("mxPoint")
				mxPoint.CreateAttr("x", fmt.Sprint(newX))
				mxPoint.CreateAttr("y", fmt.Sprint(newY))
			}
		}
	}

	// iterate through lines of file
	//	for scanner.Scan() {
	//		fmt.Println(scanner.Text())
	//	}

	/* ITERATE THROUGH RESOURCES -> INSTANCES -> DEPENDENCIES */

	// Writing back to terraform.drawio
	xml.Indent(4)
	xml.WriteToFile("terraform.drawio")
	inFile.Close()

}
