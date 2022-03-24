package utility

import (
	"KSCD/libraries/providers/GCP/GCPResources/Cases" //Additional resources
	"KSCD/libraries/providers/GCP/GCPResources/Names"
	"KSCD/libraries/providers/GCP/GCPResources/Shapes"
	"KSCD/libraries/providers/GCP/GCPResources/Zones"
	"fmt"
	//"log"
	"os"
	"syscall"
)

func LookupName(objectName string) string {
	objectShape, success := Names.NameMap[objectName]

	if objectShape, success := Names.NameMap[objectName]; success {
		return objectShape
	}

	// bug
	objectShape = "General"
	_ = success
	return objectShape
}

func LookupShape(resourceType string) string {
	objectName, success := Shapes.ShapeMap[resourceType]

	if objectName, success := Shapes.ShapeMap[resourceType]; success {
		return objectName
	}

	fmt.Println(resourceType)
	fmt.Println(objectName, success)

	return "shape=mxgraph.gcp2.blank"
}

func LookupZone(name string) string{
	zone, success := Zones.ZoneMap[name]
	if zone, success := Zones.ZoneMap[name]; success {
		return zone
	}

	fmt.Println("Something might have failed in the LookupZone function.")
	_ = zone
	_ = success

	syscall.Exit(1)
	return ""
}


func LookupCase(name string) int {
	value, success := Cases.CaseMap[name]

	if value, success := Cases.CaseMap[name]; success {
		return value
	}

	fmt.Println(name)
	fmt.Println(value, success)
	fmt.Println("Something went wrong. Exiting.")
	os.Exit(1)

	return -1
}
