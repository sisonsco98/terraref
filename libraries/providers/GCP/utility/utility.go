package utility

import (
	"KSCD/libraries/providers/GCP/GCPResources/Cases" //Additional resources
	"KSCD/libraries/providers/GCP/GCPResources/Names"
	"KSCD/libraries/providers/GCP/GCPResources/Shapes"
	"KSCD/libraries/providers/GCP/GCPResources/Zones"
	"fmt"
	"log"
	"os"
)

func LookupName(objectName string) string {
	objectShape, success := Names.NameMap[objectName]

	if objectShape, success := Names.NameMap[objectName]; success {
		return objectShape
	}



	// Improved error message

	log.Println("Utility.LookupName() wasn't able to find the draw.io name of " + objectName + ".")
	log.Println("Returning \"General\" and resuming.....")
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
	log.Println("Utility.LookupShape() wasn't able to find the draw.io shape of + " + resourceType + ".")
	log.Println("Returning a blank shape and resuming....")

	_, _, _ = resourceType, objectName, success
	return "shape=mxgraph.gcp2.blank"
}

func LookupZone(name string) string{
	zone, success := Zones.ZoneMap[name]
	if zone, success := Zones.ZoneMap[name]; success {
		return zone
	}

	log.Println("Utility.LookupZone() wasn't able to find the zone of + " + name + ".")
	log.Println("This error is unfortunately fatal. Terminating....")
	_, _ = zone, success

	os.Exit(1)
	return ""
}


func LookupCase(name string) int {
	value, success := Cases.CaseMap[name]

	if value, success := Cases.CaseMap[name]; success {
		return value
	}

	log.Println("Utility.LookupCase() wasn't able to find the case of + " + name + ".")
	log.Println("This error is unfortunately fatal. Terminating....")

	_ , _ = value, success
	os.Exit(1)

	return -1
}
