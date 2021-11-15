package mapper

import (
	"fmt"
	"log"				// logging errors
	"os"				// create and open file
//	"KSCD/parser"		// parser.go
	"KSCD/libraries"	// parser.go

	"github.com/beevik/etree"	// creating xml file (go get github.com/beevik/etree)
)

func Mapper() {

	/*** CREATE THE terraform.drawio FILE ***/

	outFile, errCreate := os.Create("terraform.drawio")
	// error creating file
	if errCreate != nil {
		log.Println("Error creating file.", errCreate)
		os.Exit(1)
	}
	// keep open
	defer outFile.Close()

	/*** CREATE ELEMENT TREE WITH PARSED DATA ***/
	
	xml := etree.NewDocument()
	xml.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	mxGraphModel := xml.CreateElement("mxGraphModel")
	mxGraphModel.CreateAttr("gridSize", "10")
	mxGraphModel.CreateAttr("pageWidth", "850")
	mxGraphModel.CreateAttr("pageHeight", "1100")

	root := mxGraphModel.CreateElement("root")

	mxCell := root.CreateElement("mxCell")
	mxCell.CreateAttr("id", "0")

	mxCell = root.CreateElement("mxCell")
	mxCell.CreateAttr("id", "1")
	mxCell.CreateAttr("parent", "0")

	/* ITERATE THROUGH RESOURCES */

//	for i := 0; i < len(parser.T.Resources); i++ {

		// (1) store resource type (ex: google_api_gateway_gateway)
//		resourceType := parser.T.Resources[i].Type

		// (2) use resource type to lookup the draw.io name (ex: Gateway)
//		objectName := "drawio name" use resourceType as lookup value

//		objectName := "Gateway"
//		objectName := "AdMob"
//		objectName := "Avere Physical Appliance"
//		objectName := "Campaign Manager"
//		objectName := "Fastly"
//		objectName := "Firebase"
//		objectName := "Forseti Security"
//		objectName := "Google Ad Manager"
//		objectName := "Google Ads"
//		objectName := "Google Analytics 360"
//		objectName := "Google Analytics"
//		objectName := "Google Play Game Services"
//		objectName := "Istio"
//		objectName := "Kubernetes"
//		objectName := "TensorFlow"

		// (3) use object name to lookup the draw.io shape (ex: shape=mxgraph.gcp2.gateway)
		objectShape := libraries.ShapeLookup[objectName]

		t := libraries.Lookup(objectName)

		switch t {
	
		/*** GCP / SERVICE CARDS ***/

			/*
				<mxCell id="2" value="" style="strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;" vertex="1" parent="1">
					<mxGeometry x="40" y="40" width="100" height="44" as="geometry" />
				</mxCell>
				<mxCell id="3" value="Gateway" style="sketch=0;dashed=0;connectable=0;html=1;fillColor=#757575;strokeColor=none;shape=mxgraph.gcp2.gateway;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontSize=12;" vertex="1" parent="2">
					<mxGeometry y="0.5" width="32" height="32" relative="1" as="geometry">
						<mxPoint x="5" y="-16" as="offset" />
					</mxGeometry>
				</mxCell>

				"Gateway"	:	"shape=mxgraph.gcp2.gateway"
			*/

		case 1:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", "2")			// fmt.Sprint(i + 2)
			mxCell.CreateAttr("value", "")
			mxCell.CreateAttr("style", "strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;")
			mxCell.CreateAttr("vertex", "1")
			mxCell.CreateAttr("parent", "1")
	
			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", "360")		// DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("y", "120")		// DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("width", "150")
			mxGeometry.CreateAttr("height", "56")
			mxGeometry.CreateAttr("as", "geometry")

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", "3")			// fmt.Sprint(i + 3)
			mxCell.CreateAttr("value", objectName)
			mxCell.CreateAttr("style", fmt.Sprint("sketch=0;dashed=0;connectable=0;html=1;fillColor=#757575;strokeColor=none;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontSize=12;" + objectShape))
			mxCell.CreateAttr("vertex", "1")
			mxCell.CreateAttr("parent", "2")		// fmt.Sprint(i + 2)
	
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

		/****************************************************************************************************/

		/*** GCP / USER AND DEVICE CARDS ***/

			// "shape=mxgraph.gcp2.application;"

			/*
				<mxCell id="2" value="Application" style="strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;labelPosition=center;verticalLabelPosition=middle;align=center;verticalAlign=bottom;spacingLeft=0;fontColor=#999999;fontSize=12;whiteSpace=wrap;spacingBottom=2;" vertex="1" parent="1">
					<mxGeometry x="40" y="40" width="70" height="85" as="geometry" />
				</mxCell>
				<mxCell id="3" value="" style="sketch=0;dashed=0;connectable=0;html=1;fillColor=#757575;strokeColor=none;shape=mxgraph.gcp2.application;part=1;" vertex="1" parent="2">
					<mxGeometry x="0.5" width="50" height="40" relative="1" as="geometry">
					    <mxPoint x="-25" y="15" as="offset" />
					</mxGeometry>
				</mxCell>
			*/

		case 2:
			
			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", "2")			// fmt.Sprint(i + 2)
			mxCell.CreateAttr("value", objectName)
			mxCell.CreateAttr("style", "strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;labelPosition=center;verticalLabelPosition=middle;align=center;verticalAlign=bottom;spacingLeft=0;fontColor=#999999;fontSize=12;whiteSpace=wrap;spacingBottom=2;")
			mxCell.CreateAttr("vertex", "1")
			mxCell.CreateAttr("parent", "1")
	
			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", "360")		// DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("y", "120")		// DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("width", "70")
			mxGeometry.CreateAttr("height", "100")
			mxGeometry.CreateAttr("as", "geometry")

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", "3")			// fmt.Sprint(i + 3)
			mxCell.CreateAttr("value", "")
			mxCell.CreateAttr("style", fmt.Sprint("sketch=0;dashed=0;connectable=0;html=1;fillColor=#757575;strokeColor=none;part=1;" + objectShape))
			mxCell.CreateAttr("vertex", "1")
			mxCell.CreateAttr("parent", "2")		// fmt.Sprint(i + 2)
	
			mxGeometry = mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", "0.5")
			mxGeometry.CreateAttr("width", "50")	// INCONSISTENT, SOME MAY LOOK OFF
			mxGeometry.CreateAttr("height", "50")	// INCONSISTENT, SOME MAY LOOK OFF
			mxGeometry.CreateAttr("relative", "1")
			mxGeometry.CreateAttr("as", "geometry")

			mxPoint := mxGeometry.CreateElement("mxPoint")
			mxPoint.CreateAttr("x", "-25")			// INCONSISTENT, SOME MAY LOOK OFF
			mxPoint.CreateAttr("y", "15")			// INCONSISTENT, SOME MAY LOOK OFF
			mxPoint.CreateAttr("as", "offset")

		/****************************************************************************************************/
			
		/*** GCP / COMPUTE ***/

			// "shape=mxgraph.gcp2.hexIcon;prIcon=compute_engine;"

			/*
				<mxCell id="2" value="" style="strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;" vertex="1" parent="1">
					<mxGeometry x="40" y="40" width="110" height="60" as="geometry" />
				</mxCell>
				<mxCell id="3" value="Compute&#xa;Engine" style="sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;shape=mxgraph.gcp2.hexIcon;prIcon=compute_engine;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" vertex="1" parent="2">
					<mxGeometry y="0.5" width="44" height="39" relative="1" as="geometry">
						<mxPoint x="5" y="-19.5" as="offset" />
					</mxGeometry>
				</mxCell>
			*/

		/*** GCP / API MANAGEMENT ***/

			// "shape=mxgraph.gcp2.hexIcon;prIcon=api_analytics;"

			/*
				<mxCell id="2" value="" style="strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;" vertex="1" parent="1">
					<mxGeometry x="40" y="40" width="110" height="60" as="geometry" />
				</mxCell>
				<mxCell id="3" value="API&#xa;Analytics" style="sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;shape=mxgraph.gcp2.hexIcon;prIcon=api_analytics;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" vertex="1" parent="2">
					<mxGeometry y="0.5" width="44" height="39" relative="1" as="geometry">
						<mxPoint x="5" y="-19.5" as="offset" />
					</mxGeometry>
				</mxCell>
			*/

		/*** GCP / SECURITY ***/

			// "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_iam;"

			/*
				<mxCell id="2" value="" style="strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;" vertex="1" parent="1">
					<mxGeometry x="40" y="40" width="100" height="60" as="geometry" />
				</mxCell>
				<mxCell id="3" value="Cloud&#xa;IAM" style="sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;shape=mxgraph.gcp2.hexIcon;prIcon=cloud_iam;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" vertex="1" parent="2">
					<mxGeometry y="0.5" width="44" height="39" relative="1" as="geometry">
						<mxPoint x="5" y="-19.5" as="offset" />
					</mxGeometry>
				</mxCell>
			*/

		/*** GCP / DATA ANALYTICS ***/

			// "shape=mxgraph.gcp2.hexIcon;prIcon=bigquery;"

			/*
				<mxCell id="2" value="" style="strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;" vertex="1" parent="1">
					<mxGeometry x="40" y="40" width="120" height="60" as="geometry" />
				</mxCell>
				<mxCell id="3" value="BigQuery" style="sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;shape=mxgraph.gcp2.hexIcon;prIcon=bigquery;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" vertex="1" parent="2">
					<mxGeometry y="0.5" width="44" height="39" relative="1" as="geometry">
						<mxPoint x="5" y="-19.5" as="offset" />
					</mxGeometry>
				</mxCell>
			*/

		/*** GCP / DATA TRANSFER ***/

			// "shape=mxgraph.gcp2.hexIcon;prIcon=transfer_appliance;"

			/*
				<mxCell id="2" value="" style="strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;" vertex="1" parent="1">
					<mxGeometry x="40" y="40" width="120" height="60" as="geometry" />
				</mxCell>
				<mxCell id="3" value="Transfer&#xa;Appliance" style="sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;shape=mxgraph.gcp2.hexIcon;prIcon=transfer_appliance;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" vertex="1" parent="2">
					<mxGeometry y="0.5" width="44" height="39" relative="1" as="geometry">
						<mxPoint x="5" y="-19.5" as="offset" />
					</mxGeometry>
				</mxCell>
			*/

		/*** GCP / CLOUD AI ***/

			// "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_machine_learning;"

			/*
				<mxCell id="2" value="" style="strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;" vertex="1" parent="1">
					<mxGeometry x="40" y="40" width="150" height="60" as="geometry" />
				</mxCell>
				<mxCell id="3" value="Cloud Machine&#xa;Learning" style="sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;shape=mxgraph.gcp2.hexIcon;prIcon=cloud_machine_learning;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" vertex="1" parent="2">
					<mxGeometry y="0.5" width="44" height="39" relative="1" as="geometry">
						<mxPoint x="5" y="-19.5" as="offset" />
					</mxGeometry>
				</mxCell>
			*/

		/*** GCP / INTERNET OF THINGS ***/

			// "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_iot_core;"

			/*
				<mxCell id="2" value="" style="strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;" vertex="1" parent="1">
					<mxGeometry x="40" y="40" width="110" height="60" as="geometry" />
				</mxCell>
				<mxCell id="3" value="Cloud&#xa;IoT Core" style="sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;shape=mxgraph.gcp2.hexIcon;prIcon=cloud_iot_core;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" vertex="1" parent="2">
					<mxGeometry y="0.5" width="44" height="39" relative="1" as="geometry">
						<mxPoint x="5" y="-19.5" as="offset" />
					</mxGeometry>
				</mxCell>
			*/

		/*** GCP / DATABASES ***/

			// "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_sql;"

			/*
				<mxCell id="2" value="" style="strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;" vertex="1" parent="1">
					<mxGeometry x="40" y="40" width="100" height="60" as="geometry" />
				</mxCell>
				<mxCell id="3" value="Cloud&#xa;SQL" style="sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;shape=mxgraph.gcp2.hexIcon;prIcon=cloud_sql;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" vertex="1" parent="2">
					<mxGeometry y="0.5" width="44" height="39" relative="1" as="geometry">
						<mxPoint x="5" y="-19.5" as="offset" />
					</mxGeometry>
				</mxCell>
			*/

		/*** GCP / STORAGE ***/

			// "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_storage;"

			/*
				<mxCell id="2" value="" style="strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;" vertex="1" parent="1">
					<mxGeometry x="40" y="40" width="110" height="60" as="geometry" />
				</mxCell>
				<mxCell id="3" value="Cloud&#xa;Storage" style="sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;shape=mxgraph.gcp2.hexIcon;prIcon=cloud_storage;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" vertex="1" parent="2">
					<mxGeometry y="0.5" width="44" height="39" relative="1" as="geometry">
						<mxPoint x="5" y="-19.5" as="offset" />
					</mxGeometry>
				</mxCell>
			*/

		/*** GCP / MANAGEMENT TOOLS ***/

			// "shape=mxgraph.gcp2.hexIcon;prIcon=stackdriver;"

			/*
				<mxCell id="2" value="" style="strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;" vertex="1" parent="1">
					<mxGeometry x="40" y="40" width="130" height="60" as="geometry" />
				</mxCell>
				<mxCell id="3" value="Stackdriver" style="sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;shape=mxgraph.gcp2.hexIcon;prIcon=stackdriver;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" vertex="1" parent="2">
					<mxGeometry y="0.5" width="44" height="39" relative="1" as="geometry">
						<mxPoint x="5" y="-19.5" as="offset" />
					</mxGeometry>
				</mxCell>
			*/

		/*** GCP / NETWORKING ***/

			// "shape=mxgraph.gcp2.hexIcon;prIcon=virtual_private_cloud;"

			/*
				<mxCell id="2" value="" style="strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;" vertex="1" parent="1">
					<mxGeometry x="40" y="40" width="140" height="60" as="geometry" />
				</mxCell>
				<mxCell id="3" value="Virtual&#xa;Private Cloud" style="sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;shape=mxgraph.gcp2.hexIcon;prIcon=virtual_private_cloud;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" vertex="1" parent="2">
					<mxGeometry y="0.5" width="44" height="39" relative="1" as="geometry">
						<mxPoint x="5" y="-19.5" as="offset" />
					</mxGeometry>
				</mxCell>
			*/
	
		/*** GCP / DEVELOPER TOOLS ***/

			// "shape=mxgraph.gcp2.hexIcon;prIcon=placeholder;"

			/*
				<mxCell id="2" value="" style="strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;" vertex="1" parent="1">
					<mxGeometry x="40" y="40" width="100" height="60" as="geometry" />
				</mxCell>
				<mxCell id="3" value="Cloud&#xa;SDK" style="sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;shape=mxgraph.gcp2.hexIcon;prIcon=placeholder;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" vertex="1" parent="2">
					<mxGeometry y="0.5" width="44" height="39" relative="1" as="geometry">
						<mxPoint x="5" y="-19.5" as="offset" />
					</mxGeometry>
				</mxCell>
			*/

		case 3:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", "2")			// fmt.Sprint(i + 2)
			mxCell.CreateAttr("value", "")
			mxCell.CreateAttr("style", "strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;")
			mxCell.CreateAttr("vertex", "1")
			mxCell.CreateAttr("parent", "1")
	
			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", "360")		// DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("y", "120")		// DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("width", "175")
			mxGeometry.CreateAttr("height", "60")
			mxGeometry.CreateAttr("as", "geometry")

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", "3")			// fmt.Sprint(i + 3)
			mxCell.CreateAttr("value", objectName)
			mxCell.CreateAttr("style", fmt.Sprint("sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" + objectShape))
			mxCell.CreateAttr("vertex", "1")
			mxCell.CreateAttr("parent", "2")		// fmt.Sprint(i + 2)
	
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

		/****************************************************************************************************/
	
		/*** GCP / PRODUCT CARDS ***/

			// "shape=mxgraph.gcp2.kubernetes_logo"

			/*
				<mxCell id="2" value="" style="strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;" vertex="1" parent="1">
					<mxGeometry x="40" y="40" width="130" height="60" as="geometry" />
				</mxCell>
				<mxCell id="3" value="Kubernetes" style="sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;shape=mxgraph.gcp2.kubernetes_logo;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" vertex="1" parent="2">
					<mxGeometry width="45" height="43.65" relative="1" as="geometry">
						<mxPoint x="5" y="7.675000000000001" as="offset" />
					</mxGeometry>
				</mxCell>
			*/
		
		case 4:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", "2")			// fmt.Sprint(i + 2)
			mxCell.CreateAttr("value", "")
			mxCell.CreateAttr("style", "strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;")
			mxCell.CreateAttr("vertex", "1")
			mxCell.CreateAttr("parent", "1")
	
			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", "360")		// DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("y", "120")		// DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("width", "150")
			mxGeometry.CreateAttr("height", "75")
			mxGeometry.CreateAttr("as", "geometry")

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", "3")			// fmt.Sprint(i + 3)
			mxCell.CreateAttr("value", objectName)
			mxCell.CreateAttr("style", fmt.Sprint("sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" + objectShape))
			mxCell.CreateAttr("vertex", "1")
			mxCell.CreateAttr("parent", "2")		// fmt.Sprint(i + 2)
	
			mxGeometry = mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("width", "45")	// INCONSISTENT, SOME MAY LOOK OFF (esp. Avere Physical Appliance, Fastly, Google Play Game Services)
			mxGeometry.CreateAttr("height", "45")	// INCONSISTENT, SOME MAY LOOK OFF (esp. Avere Physical Appliance, Fastly, Google Play Game Services)
			mxGeometry.CreateAttr("relative", "1")
			mxGeometry.CreateAttr("as", "geometry")

			mxPoint := mxGeometry.CreateElement("mxPoint")
			mxPoint.CreateAttr("x", "5")			// INCONSISTENT, SOME MAY LOOK OFF (esp. Avere Physical Appliance, Fastly, Google Play Game Services)
			mxPoint.CreateAttr("y", "7")			// INCONSISTENT, SOME MAY LOOK OFF (esp. Avere Physical Appliance, Fastly, Google Play Game Services)
			mxPoint.CreateAttr("as", "offset")

		/****************************************************************************************************/

		/*** GCP ICONS ***/

			// "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_machine_learning;"

			/*
				<mxCell id="2" value="AI Platform" style="sketch=0;html=1;fillColor=#5184F3;strokeColor=none;verticalAlign=top;labelPosition=center;verticalLabelPosition=bottom;align=center;spacingTop=-6;fontSize=11;fontStyle=1;fontColor=#999999;shape=mxgraph.gcp2.hexIcon;prIcon=cloud_machine_learning" vertex="1" parent="1">
					<mxGeometry x="40" y="40" width="66" height="58.5" as="geometry" />
				</mxCell>
			*/

		case 5:

			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", "2")			// fmt.Sprint(i + 2)
			mxCell.CreateAttr("value", objectName)
			mxCell.CreateAttr("style", fmt.Sprint("sketch=0;html=1;fillColor=#5184F3;strokeColor=none;verticalAlign=top;labelPosition=center;verticalLabelPosition=bottom;align=center;spacingTop=-6;fontSize=11;fontStyle=1;fontColor=#999999;" + objectShape))
			mxCell.CreateAttr("vertex", "1")
			mxCell.CreateAttr("parent", "1")
	
			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", "360")		// DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("y", "120")		// DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("width", "66")
			mxGeometry.CreateAttr("height", "58.5")
			mxGeometry.CreateAttr("as", "geometry")

		// GCP ICONS / Developer Tools / Cloud Scheduler
		case 6:
			
			mxCell = root.CreateElement("mxCell")
			mxCell.CreateAttr("id", "2")			// fmt.Sprint(i + 2)
			mxCell.CreateAttr("value", objectName)
			mxCell.CreateAttr("style", fmt.Sprint("html=1;fillColor=#5184F3;strokeColor=none;verticalAlign=top;labelPosition=center;verticalLabelPosition=bottom;align=center;fontSize=11;fontStyle=1;fontColor=#999999;" + objectShape))
			mxCell.CreateAttr("vertex", "1")
			mxCell.CreateAttr("parent", "1")
	
			mxGeometry := mxCell.CreateElement("mxGeometry")
			mxGeometry.CreateAttr("x", "360")		// DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("y", "120")		// DETERMINE METHOD FOR SETTING THIS
			mxGeometry.CreateAttr("width", "30")
			mxGeometry.CreateAttr("height", "34.5")
			mxGeometry.CreateAttr("as", "geometry")

		/****************************************************************************************************/
		
		/*** GCP / PATHS ***/
	
			// skip for now

		/****************************************************************************************************/
	
		/*** GCP / ZONES ***/
	
			// skip for now

		/****************************************************************************************************/
	
		/*** GCP / EXPANDED PRODUCT CARDS ***/
	
			// skip for now

		/****************************************************************************************************/
	
		/*** GCP / GENERAL ICONS ***/

			// skip for now

		/****************************************************************************************************/

		// error case
		default:
			log.Println("Error: No match.", errCreate)
			os.Exit(1)
		}

//	}

	/*** PRINT TO THE terraform.drawio FILE ***/

	xml.Indent(4)
	xml.WriteToFile("terraform.drawio")
	
	// close file
	outFile.Close()
}
