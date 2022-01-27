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
	var ArrowsOverlap bool = false
	var ArrowResized bool = true

	fmt.Println(ArrowsOverlap)
	fmt.Println(ArrowResized)

	// Loops through Pizza to get the slices
	for _, pizza := range mapper.Pizza {
		fmt.Println(pizza)

		// Goes through all the arrows existing
		for _, arrow := range mapper.ArrowRelationships {

			// Check if arrows overlap at all

			// Check if arrow has already been resized

			// If the arrow does overlap AND it has not run through this function yet, create it's bending array (NOTE: Functions below are being performed above due to need to go through pizza function)
			if ArrowOverlap(arrow.XPosSource, arrow.YPosSource, arrow.XPosTarget, arrow.YPosTarget) == true && ArrowResize(arrow.ArrowID+1) == false {

				// Finds a Y Coordinate that is unoccupied
				var newY int = FindFreeSpace() // Will be taken out, and newY set to y coordinate found below
				for _, slice := range mapper.Pizza {
					fmt.Println(slice)
				}

				// Writing of necessary XML code to create bend
				path := fmt.Sprintf("/mxGraphModel/root/mxCell[%d]/mxGeometry", arrow.ArrowID+1)
				arrowGeom := xml.FindElement(path)

				array := arrowGeom.CreateElement("Array")
				array.CreateAttr("as", "points")

				mxPoint := array.CreateElement("mxPoint")
				mxPoint.CreateAttr("x", fmt.Sprint(arrow.XPosSource))
				mxPoint.CreateAttr("y", fmt.Sprint(newY))

				mxPoint = array.CreateElement("mxPoint")
				mxPoint.CreateAttr("x", fmt.Sprint(arrow.XPosTarget))
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

// Function will take in the arrow coordinates, then needs to loop through Pizza and find if arrow intersects (might not work as its own function due to [])
func ArrowOverlap(XPosSource int, YPosSource int, XPosTarget int, YPosTarget int) bool {

	return true
}

// Checks if the arrow has been resized before
func ArrowResize(ArrowID int) bool {
	return false
}

// Will check for a y coordinate that has a free space
func FindFreeSpace() int {

	return 15
}
