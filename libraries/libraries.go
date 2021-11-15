package libraries

import (
	"fmt"
	"os"
)

var ShapeLookup = map[string]string{

	/*** GCP / PATHS ***/

	/*** GCP / ZONES ***/

	/*** GCP / SERVICE CARDS ***/

	"Gateway"								:	"shape=mxgraph.gcp2.gateway",
	"Memcache"								:	"shape=mxgraph.gcp2.memcache",
	"Logs API"								:	"shape=mxgraph.gcp2.logs_api",
	"Cluster"								:	"shape=mxgraph.gcp2.cluster",
	"NAT"									:	"shape=mxgraph.gcp2.nat",
	"Squid Proxy"							:	"shape=mxgraph.gcp2.squid_proxy",
	"Bucket"								:	"shape=mxgraph.gcp2.bucket",
	"Service Discovery"						:	"shape=mxgraph.gcp2.service_discovery",
	"Task Queues"							:	"shape=mxgraph.gcp2.task_queues",
	"Image Services"						:	"shape=mxgraph.gcp2.image_services",
	"Dedicated Game Server"					:	"shape=mxgraph.gcp2.dedicated_game_server",
	"Frontend Platform Services"			:	"shape=mxgraph.gcp2.frontend_platform_services",
	"Google Edge POP"						:	"shape=mxgraph.gcp2.google_network_edge_cache",
	"External Payment Form"					:	"shape=mxgraph.gcp2.external_payment_form",
	"Internal Payment Authorization"		:	"shape=mxgraph.gcp2.internal_payment_authorization",
	"VPN Gateway"							:	"shape=mxgraph.gcp2.gateway",
	"Application System(s)"					:	"shape=mxgraph.gcp2.application_system",
	"Virtual File System"					:	"shape=mxgraph.gcp2.virtual_file_system",
	"CDN Interconnect"						:	"shape=mxgraph.gcp2.google_network_edge_cache",
	"Scheduled Tasks"						:	"shape=mxgraph.gcp2.scheduled_tasks",
	"HTTPS Load Balancer"					:	"shape=mxgraph.gcp2.network_load_balancer",
	"Persistent Disk Snapshot"				:	"shape=mxgraph.gcp2.persistent_disk_snapshot",
	"Persistent Disk"						:	"shape=mxgraph.gcp2.persistent_disk_snapshot",
	"Network Load Balancer"					:	"shape=mxgraph.gcp2.network_load_balancer",
	"Google Network w/ Edge Cache"			:	"shape=mxgraph.gcp2.google_network_edge_cache",
	"Push Notification Service"				:	"shape=mxgraph.gcp2.push_notification_service",
	"Blank One Line (w/ bubble)"			:	"shape=mxgraph.gcp2.blank",
	"Blank One Line (w/o bubble)"			:	"",
	"Blank Two and Three Line (w/ Bubble)"	:	"shape=mxgraph.gcp2.blank",
	"Blank Two and Three Line (w/o bubble)"	:	"",

	/*** GCP / USER AND DEVICE CARDS ***/

	"Application"							:	"",
	"Beacon"								:	"",
	"Circuit Board"							:	"",
	"Database"								:	"",
	"Desktop"								:	"",
	"Desktop and Mobile"					:	"",
	"Game"									:	"",
	"Gateway"								:	"",
	"Laptop"								:	"",
	"Lightbulb"								:	"",
	"List"									:	"",
	"Live"									:	"",
	"Local-Compute"							:	"",
	"Mobile Devices"						:	"",
	"Payment"								:	"",
	"Phone"									:	"",
	"Record"								:	"",
	"Report"								:	"",
	"Retail"								:	"",
	"Speaker"								:	"",
	"Storage"								:	"",
	"Stream"								:	"",
	"Users"									:	"",
	"Webcam"								:	"",

	// NOTE CASE FOR MULTIPLE CARDS OPTIONS ???

	/*** GCP / COMPUTE ***/

	"Compute Engine"						:	"",
	"GPU"									:	"",
	"App Engine"							:	"",
	"Cloud Functions"						:	"",
	"Kubernetes Engine"						:	"",
	"Container-Optimized OS"				:	"",
	"Cloud Run"								:	"",
	"GKE-on-Prem"							:	"",

	/*** GCP / API MANAGEMENT ***/

	"API Analytics"							:	"",
	"Apigee Sense"							:	"",
	"API Monetization"						:	"",
	"Cloud Endpoints"						:	"",
	"Apigee API Platform"					:	"",
	"Developer Portal"						:	"",

	/*** GCP / SECURITY ***/

	"Cloud IAM"								:	"",
	"BeyondCorp"							:	"",
	"Cloud Resource Manager"				:	"",
	"Data Loss Prevention API"				:	"",
	"Cloud Security Scanner"				:	"",
	"Key Management Service"				:	"",
	"Identity-Aware Proxy"					:	"",
	"Cloud Security Command Center"			:	"",
	"Security Key Enforcement"				:	"",

	/*** GCP / DATA ANALYTICS ***/

	"BigQuery"								:	"",
	"Cloud Datalab"							:	"",
	"Cloud Dataflow"						:	"",
	"Cloud Pub/Sub"							:	"",
	"Cloud Dataproc"						:	"",
	"Genomics"								:	"",
	"Cloud Dataprep"						:	"",
	"Cloud Composer"						:	"",
	"Cloud Data Catalog"					:	"",
	"Cloud Data Fusion"						:	"",

	/*** GCP / DATA TRANSFER ***/

	"Transfer Appliance"					:	"",

	/*** GCP / CLOUD AI ***/

	"Cloud Machine Learning"				:	"",
	"Natural Language API"					:	"",
	"Vision API"							:	"",
	"Translation API"						:	"",
	"Speech API"							:	"",
	"Jobs API"								:	"",
	"Cloud Video Intelligence API"			:	"",
	"Advanced Solutions Lab"				:	"",

	/*** GCP / INTERNET OF THINGS ***/

	"Cloud IoT Core"						:	"",

	/*** GCP / DATABASES ***/

	"Cloud SQL"								:	"",
	"Cloud Bigtable"						:	"",
	"Cloud Spanner"							:	"",
	"Cloud Memorystore"						:	"",
	"Cloud Firestore"						:	"",
	"Cloud Datastore"						:	"",

	/*** GCP / STORAGE ***/

	"Cloud Storage"							:	"",
	"Persistent Disk"						:	"",
	"Cloud Firestore"						:	"",

	/*** GCP / MANAGEMENT TOOLS ***/

	"Stackdriver"							:	"",
	"Debugger"								:	"",
	"Monitoring"							:	"",
	"Deployment Manager"					:	"",
	"Logging"								:	"",
	"Cloud Console"							:	"",
	"Error Reporting"						:	"",
	"Cloud Shell"							:	"",
	"Trace"									:	"",
	"Cloud Mobile App"						:	"",
	"Profiler"								:	"",
	"Billing API"							:	"",
	"Cloud APIs"							:	"",

	/*** GCP / NETWORKING ***/

	"Virtual Private Cloud"					:	"",
	"Dedicated Interconnect"				:	"",
	"Cloud Load Balancing"					:	"",
	"Cloud DNS"								:	"",
	"Cloud CDN"								:	"",
	"Cloud Network"							:	"",
	"Cloud External IP Addresses"			:	"",
	"Cloud Routes"							:	"",
	"Cloud Firewall Rules"					:	"",
	"Cloud VPN"								:	"",
	"Cloud Router"							:	"",
	"Cloud Armor"							:	"",
	"Standard Network Tier"					:	"",
	"Premium Network Tier"					:	"",
	"Partner Interconnect"					:	"",
	
	/*** GCP / DEVELOPER TOOLS ***/

	"Cloud SDK"								:	"",
	"Cloud Build"							:	"",
	"Gradle App Engine Plugin"				:	"",
	"Cloud Tools for Visual Studio"			:	"",
	"Cloud Source Repositories"				:	"",
	"Maven App Engine Plugin"				:	"",
	"Cloud Tools for Eclipse"				:	"",
	"Cloud Tools for IntelliJ"				:	"",
	"Cloud Test Lab"						:	"",
	"Cloud Tools for PowerShell"			:	"",
	"IDE Plugins"							:	"",
	"Container Registry"					:	"",
	
	/*** GCP / EXPANDED PRODUCT CARDS ***/

	// skip for now
	
	/*** GCP / PRODUCT CARDS ***/

	"Kubernetes"							:	"",
	"TensorFlow"							:	"",
	"Forseti Security"						:	"",
	"Istio"									:	"",
	"Firebase"								:	"",
	"Fastly"								:	"",
	"AdMob"									:	"",
	"Google Play Game Services"				:	"",
	"Campaign Manager"						:	"",
	"Google Analytics"						:	"",
	"Google Ads"							:	"",
	"Avere Physical Appliance"				:	"",
	"Google Analytics 360"					:	"",
	"Google Ad Manager"						:	"",
	
	/*** GCP / GENERAL ICONS ***/

	// skip for now
	
	/*** GCP ICONS / ... ***/

	// ???
}

// class 1 objects are those in the GCP/Service Cards Category.
var classLookup = map[string]int{

	/*** GCP / Service Cards ***/

	"Gateway"								:	1,
	"Memcache"								:	1,
	"Logs API"								:	1,
	"Cluster"								:	1,
	"NAT"									:	1,
	"Squid Proxy"							:	1,
	"Bucket"								:	1,
	"Service Discovery"						:	1,
	"Task Queues"							:	1,
	"Image Services"						:	1,
	"Dedicated Game Server"					:	1,
	"Frontend Platform Services"			:	1,
	"Google Edge POP"						:	1,
	"External Payment Form"					:	1,
	"Internal Payment Authorization"		:	1,
	"VPN Gateway"							:	1,
	"Application System(s)"					:	1,
	"Virtual File System"					:	1,
	"CDN Interconnect"						:	1,
	"Scheduled Tasks"						:	1,
	"HTTPS Load Balancer"					:	1,
	"Persistent Disk Snapshot"				:	1,
	"Persistent Disk"						:	1,
	"Network Load Balancer"					:	1,
	"Google Network w/ Edge Cache"			:	1,
	"Push Notification Service"				:	1,
	"Blank One Line (w/ bubble)"			:	1,
	"Blank One Line (w/o bubble)"			:	1,
	"Blank Two and Three Line (w/ Bubble)"	:	1,
	"Blank Two and Three Line (w/o bubble)"	:	1,
}

func Lookup(name string) int {
	value, success := classLookup[name]

	if value, success := classLookup[name]; success {
		return value
	}

	fmt.Println(name)
	fmt.Println(value, success)
	fmt.Println("Something went wrong. Exiting.")
	os.Exit(1)

	return -1
}
