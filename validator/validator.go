package validator

import (
	"fmt"
	"log"	// logging errors
	"os"	// create and open files
	"bufio"	// scanning files

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
		for _, slice := range mapper.Pizza {

			// Check if arrows overlap boxes at all
			if arrow.SourceID != slice.HiddenId && arrow.TargetID != slice.HiddenId {

				fmt.Println(arrow, arrow.XPosSource, arrow.YPosSource, arrow.XPosTarget, arrow.YPosTarget, "test arrow")
				// fmt.Println(slice, slice.XPosCenter, slice.YPosCenter, "test slice")

				if Ymatch == true {
					// Horizontal lines
					if arrow.YPosSource > arrow.YPosTarget {
						// Source below target
						if arrow.YPosTarget <= slice.YPosCenter && slice.YPosCenter <= arrow.YPosSource {
							ArrowsOverlap = true
						}
					} else {
						// Target above source
						if arrow.YPosTarget >= slice.YPosCenter && slice.YPosCenter >= arrow.YPosSource {
							ArrowsOverlap = true
						}
					}

				} else {
					// Vertical lines
					if arrow.XPosSource > arrow.XPosTarget {
						// Source below target
						if arrow.XPosTarget <= slice.XPosCenter && slice.XPosCenter <= arrow.XPosSource {
							ArrowsOverlap = true
						}
					} else {
						// Target above source
						if arrow.XPosTarget >= slice.XPosCenter && slice.XPosCenter >= arrow.XPosSource {
							ArrowsOverlap = true
						}
					}

				}
			}

//			if (arrow.XPosSource == arrow.XPosTarget) {
			// if (ArrowsOverlap == true) && (arrow.XPosSource == arrow.XPosTarget) {
			// 		// vertical arrow
			// 	fmt.Println("VERTICAL ARROW")

			// 	if (arrow.YPosSource + (60 * 2) == arrow.YPosTarget) {
			// 		// source directly above target
			// 		fmt.Println("SOURCE DIRECTLY ABOVE TARGET")
			// 	} else if (arrow.YPosSource == arrow.YPosTarget + (60 * 2)) {
			// 		// source directly below target
			// 		fmt.Println("SOURCE DIRECTLY BELOW TARGET")
			// 	} else { //////////////////////////////////////////////////////////////////////////////////////////
			//// }

			if (ArrowsOverlap == true) && !((arrow.XPosSource == arrow.XPosTarget) && (arrow.YPosSource + (60 * 2) == arrow.YPosTarget)) {
			// 					fmt.Println(arrow.XPosSource, arrow.YPosSource, arrow.XPosTarget, arrow.YPosTarget, "VERTICAL")
			// 				}

				// If the arrow does overlap AND it has not run through this function yet, create it's bending array
//			if ArrowsOverlap == true && arrow.HasMoved == false {
				var newX int
				var newY int
				var stillOverlaps bool = true

				// Finding an X or Y Coordinate that is unoccupied
				// Checking first if we need to move along x or y axis
				if arrow.YPosSource == arrow.YPosTarget {

					for stillOverlaps == true {
						newY = arrow.YPosSource + slice.Height
						stillOverlaps = false
						arrow.HasMoved = true
					}

						// Check if still overlaps
	
						if arrow.SourceID != slice.HiddenId && arrow.TargetID != slice.HiddenId {
	
							if Ymatch == true {
								// Horizontal lines
								if arrow.YPosSource > arrow.YPosTarget {
									// Source below target
									if arrow.YPosTarget <= slice.YPosCenter && slice.YPosCenter <= arrow.YPosSource {
										stillOverlaps = true
									}
								} else {
									// Target above source
									if arrow.YPosTarget >= slice.YPosCenter && slice.YPosCenter >= arrow.YPosSource {
										stillOverlaps = true
									}
								}
			
							} else {
								// Vertical lines
								if arrow.XPosSource > arrow.XPosTarget {
									// Source below target
									if arrow.XPosTarget <= slice.XPosCenter && slice.XPosCenter <= arrow.XPosSource {
										stillOverlaps = true
									}
								} else {
									// Target above source
									if arrow.XPosTarget >= slice.XPosCenter && slice.XPosCenter >= arrow.XPosSource {
										stillOverlaps = true
									}
								}
			
							}
						}

						for stillOverlaps == true {
							newY = arrow.YPosSource - slice.Height
							stillOverlaps = false
							arrow.HasMoved = true
						}

					// Writing of necessary XML code to create bends
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

				} else {

					for stillOverlaps == true {
						newX = arrow.XPosSource + slice.Width
						stillOverlaps = false
						arrow.HasMoved = true
					}

						// Check if still overlaps
	
						if arrow.SourceID != slice.HiddenId && arrow.TargetID != slice.HiddenId {
	
							if Ymatch == true {
								// Horizontal lines
								if arrow.YPosSource > arrow.YPosTarget {
									// Source below target
									if arrow.YPosTarget <= slice.YPosCenter && slice.YPosCenter <= arrow.YPosSource {
										stillOverlaps = true
									}
								} else {
									// Target above source
									if arrow.YPosTarget >= slice.YPosCenter && slice.YPosCenter >= arrow.YPosSource {
										stillOverlaps = true
									}
								}
			
							} else {
								// Vertical lines
								if arrow.XPosSource > arrow.XPosTarget {
									// Source below target
									if arrow.XPosTarget <= slice.XPosCenter && slice.XPosCenter <= arrow.XPosSource {
										stillOverlaps = true
									}
								} else {
									// Target above source
									if arrow.XPosTarget >= slice.XPosCenter && slice.XPosCenter >= arrow.XPosSource {
										stillOverlaps = true
									}
								}
			
							}
						}

						for stillOverlaps == true {
							newX = arrow.XPosSource - slice.Width
							stillOverlaps = false
							arrow.HasMoved = true
						}

					// Writing of necessary XML code to create bends
					path := fmt.Sprintf("/mxGraphModel/root/mxCell[%d]/mxGeometry", arrow.ArrowID+1)
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
			}//}
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
