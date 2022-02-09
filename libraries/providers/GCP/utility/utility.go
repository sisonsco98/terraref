package utility

import (
	"KSCD/libraries/providers/GCP/GCPResources" //Additional resources
	"fmt"
	"log"
	"os"
	"syscall"
)

func LookupName(objectName string) string {
	objectShape, success := GCPResources.NameMap[objectName]

	if objectShape, success := GCPResources.NameMap[objectName]; success {
		return objectShape
	}

	// bug
	objectShape = "General"
	_ = success
	return objectShape
}

func LookupShape(resourceType string) string {
	objectName, success := GCPResources.ShapeMap[resourceType]

	if objectName, success := GCPResources.ShapeMap[resourceType]; success {
		return objectName
	}

	fmt.Println(resourceType)
	fmt.Println(objectName, success)

	return "shape=mxgraph.gcp2.blank"
}

func LookupZone(name string) string{
	zone, success := GCPResources.ZoneMap[name]
	if zone, success := GCPResources.ZoneMap[name]; success {
		return zone
	}

	fmt.Println(name)
	fmt.Println(zone, success)

	syscall.Exit(1)
	return ""
}


func LookupCase(name string) int {
	value, success := GCPResources.CaseMap[name]

	if value, success := GCPResources.CaseMap[name]; success {
		return value
	}

	fmt.Println(name)
	fmt.Println(value, success)
	fmt.Println("Something went wrong. Exiting.")
	os.Exit(1)

	return -1
}

func Dimensions(class int) (int, int) {
	switch class {
	case 1:
		return 250, 60 // 150, 56
	case 2:
		return 250, 60 // 70, 100
	case 3:
		return 250, 60 // 175, 60
	case 4:
		return 250, 60 // 150, 60
	case 5:
		return 250, 60 // 66, 59
	case 6:
		return 250, 60 // 30, 35
	case 7:
		return 250, 60 // 220, 190
	default:
		log.Println("Unable to find the dimensions.")
		os.Exit(0)
	}
	return 999, 999
}