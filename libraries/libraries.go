package libraries

import (
	"fmt"
	"os"
)

var CaseMap = map[string]int{

	/*** GCP / SERVICE CARDS ***/

	"Application System(s)":          1,
	"Bucket":                         1,
	"CDN Interconnect":               1,
	"Cluster":                        1,
	"Dedicated Game Server":          1,
	"External Payment Form":          1,
	"Frontend Platform Services":     1,
	"Gateway":                        1,
	"Google Edge POP":                1,
	"Google Network w/ Edge Cache":   1,
	"HTTPS Load Balancer":            1,
	"Image Services":                 1,
	"Internal Payment Authorization": 1,
	"Logs API":                       1,
	"Memcache":                       1,
	"NAT":                            1,
	"Network Load Balancer":          1,
	"Persistent Disk Snapshot":       1,
	// "Persistent Disk":                1,
	"Push Notification Service": 1,
	"Scheduled Tasks":           1,
	"Service Discovery":         1,
	"Squid Proxy":               1,
	"Task Queues":               1,
	"VPN Gateway":               1,
	"Virtual File System":       1,

	/****************************************************************************************************/

	/*** GCP / USER AND DEVICE CARDS ***/

	"Application":        2,
	"Beacon":             2,
	"Circuit Board":      2,
	"Database":           2,
	"Desktop":            2,
	"Desktop and Mobile": 2,
	"Game":               2,
	// "Gateway":            2,
	"Laptop":         2,
	"Lightbulb":      2,
	"List":           2,
	"Live":           2,
	"Local-Compute":  2,
	"Mobile Devices": 2,
	"Payment":        2,
	"Phone":          2,
	"Record":         2,
	"Report":         2,
	"Retail":         2,
	"Speaker":        2,
	"Storage":        2,
	"Stream":         2,
	"Users":          2,
	"Webcam":         2,

	/****************************************************************************************************/

	/*** GCP / API MANAGEMENT ***/
	/*** GCP / CLOUD AI ***/
	/*** GCP / COMPUTE ***/
	/*** GCP / DATA ANALYTICS ***/
	/*** GCP / DATA TRANSFER ***/
	/*** GCP / DATABASES ***/
	/*** GCP / DEVELOPER TOOLS ***/
	/*** GCP / INTERNET OF THINGS ***/
	/*** GCP / MANAGEMENT TOOLS ***/
	/*** GCP / NETWORKING ***/
	/*** GCP / SECURITY ***/
	/*** GCP / STORAGE ***/

	"API Analytics":                 3,
	"API Monetization":              3,
	"Advanced Solutions Lab":        3,
	"Apigee API Platform":           3,
	"Apigee Sense":                  3,
	"App Engine":                    3,
	"BeyondCorp":                    3,
	"BigQuery":                      3,
	"Billing API":                   3,
	"Cloud APIs":                    3,
	"Cloud Armor":                   3,
	"Cloud Bigtable":                3,
	"Cloud Build":                   3,
	"Cloud CDN":                     3,
	"Cloud Composer":                3,
	"Cloud Console":                 3,
	"Cloud DNS":                     3,
	"Cloud Data Catalog":            3,
	"Cloud Data Fusion":             3,
	"Cloud Dataflow":                3,
	"Cloud Datalab":                 3,
	"Cloud Dataprep":                3,
	"Cloud Dataproc":                3,
	"Cloud Datastore":               3,
	"Cloud Endpoints":               3,
	"Cloud External IP Addresses":   3,
	"Cloud Firestore":               3,
	"Cloud Firewall Rules":          3,
	"Cloud Functions":               3,
	"Cloud IAM":                     3,
	"Cloud IoT Core":                3,
	"Cloud Load Balancing":          3,
	"Cloud Machine Learning":        3,
	"Cloud Memorystore":             3,
	"Cloud Mobile App":              3,
	"Cloud Network":                 3,
	"Cloud Pub/Sub":                 3,
	"Cloud Resource Manager":        3,
	"Cloud Router":                  3,
	"Cloud Routes":                  3,
	"Cloud Run":                     3,
	"Cloud SDK":                     3,
	"Cloud SQL":                     3,
	"Cloud Security Command Center": 3,
	"Cloud Security Scanner":        3,
	"Cloud Shell":                   3,
	"Cloud Source Repositories":     3,
	"Cloud Spanner":                 3,
	"Cloud Storage":                 3,
	"Cloud Test Lab":                3,
	"Cloud Tools for Eclipse":       3,
	"Cloud Tools for IntelliJ":      3,
	"Cloud Tools for PowerShell":    3,
	"Cloud Tools for Visual Studio": 3,
	"Cloud VPN":                     3,
	"Cloud Video Intelligence API":  3,
	"Compute Engine":                3,
	"Container Registry":            3,
	"Container-Optimized OS":        3,
	"Data Loss Prevention API":      3,
	"Debugger":                      3,
	"Dedicated Interconnect":        3,
	"Deployment Manager":            3,
	"Developer Portal":              3,
	"Error Reporting":               3,
	"GKE-on-Prem":                   3,
	"GPU":                           3,
	"Genomics":                      3,
	"Gradle App Engine Plugin":      3,
	"IDE Plugins":                   3,
	"Identity-Aware Proxy":          3,
	"Jobs API":                      3,
	"Key Management Service":        3,
	"Kubernetes Engine":             3,
	"Logging":                       3,
	"Maven App Engine Plugin":       3,
	"Monitoring":                    3,
	"Natural Language API":          3,
	"Partner Interconnect":          3,
	"Persistent Disk":               3,
	"Premium Network Tier":          3,
	"Profiler":                      3,
	"Security Key Enforcement":      3,
	"Speech API":                    3,
	"Stackdriver":                   3,
	"Standard Network Tier":         3,
	"Trace":                         3,
	"Transfer Appliance":            3,
	"Translation API":               3,
	"Virtual Private Cloud":         3,
	"Vision API":                    3,

	/****************************************************************************************************/

	/*** GCP / PRODUCT CARDS ***/

	"AdMob":                     4,
	"Avere Physical Appliance":  4,
	"Campaign Manager":          4,
	"Fastly":                    4,
	"Firebase":                  4,
	"Forseti Security":          4,
	"Google Ad Manager":         4,
	"Google Ads":                4,
	"Google Analytics 360":      4,
	"Google Analytics":          4,
	"Google Play Game Services": 4,
	"Istio":                     4,
	"Kubernetes":                4,
	"TensorFlow":                4,

	/****************************************************************************************************/

	/*** GCP ICONS ***/

	// "AI Hub":                            5,
	// "AI Platform Data Labeling Service": 5,
	// "AI Platform":                       5,
	// "API Analytics":                     5,
	// "API Monetization":                  5,
	// "Advanced Solutions Lab":            5,
	// "Apigee API Platform":               5,
	// "Apigee Sense":                      5,
	// "App Engine":                        5,
	// "AutoML Natural Language":           5,
	// "AutoML Tables":                     5,
	// "AutoML Translation":                5,
	// "AutoML Video Intelligence":         5,
	// "AutoML Vision":                     5,
	// "BigQuery":                          5,
	// "Cloud APIs":                        5,
	// "Cloud Armor":                       5,
	// "Cloud AutoML":                      5,
	// "Cloud Bigtable":                    5,
	// "Cloud Billing API":                 5,
	// "Cloud Build":                       5,
	// "Cloud CDN":                         5,
	// "Cloud Code":                        5,
	// "Cloud Composer":                    5,
	// "Cloud Console":                     5,
	// "Cloud DNS":                         5,
	// "Cloud Data Catalog":                5,
	// "Cloud Data Fusion":                 5,
	// "Cloud Dataflow":                    5,
	// "Cloud Datalab":                     5,
	// "Cloud Dataprep":                    5,
	// "Cloud Dataproc":                    5,
	// "Cloud Datastore":                   5,
	// "Cloud Deployment Manager":          5,
	// "Cloud Endpoints":                   5,
	// "Cloud External IP Addresses":       5,
	// "Cloud Firestore":                   5,
	// "Cloud Firewall Rules":              5,
	// "Cloud Functions":                   5,
	// "Cloud IAM":                         5,
	// "Cloud Interface API":               5,
	// "Cloud IoT Core":                    5,
	// "Cloud Jobs API":                    5,
	// "Cloud Load Balancing":              5,
	// "Cloud Memorystore":                 5,
	// "Cloud Mobile App":                  5,
	// "Cloud NAT":                         5,
	// "Cloud Natural Language API":        5,
	// "Cloud Network":                     5,
	// "Cloud Pub/Sub":                     5,
	// "Cloud Resource Manager":            5,
	// "Cloud Router":                      5,
	// "Cloud Routes":                      5,
	// "Cloud Run":                         5,
	// "Cloud Run (On Prem)":               5,
	// "Cloud SDK":                         5,
	// "Cloud SQL":                         5,
	// "Cloud Security Command Center":     5,
	// "Cloud Security Scanner":            5,
	// "Cloud Service Mesh":                5,
	// "Cloud Shell":                       5,
	// "Cloud Source Repositories":         5,
	// "Cloud Spanner":                     5,
	// "Cloud Speech-to-Text":              5,
	// "Cloud Storage":                     5,
	// "Cloud TPU":                         5,
	// "Cloud Tasks":                       5,
	// "Cloud Test Lab":                    5,
	// "Cloud Text-to-Speech":              5,
	// "Cloud Tools for Eclipse":           5,
	// "Cloud Tools for IntelliJ":          5,
	// "Cloud Tools for PowerShell":        5,
	// "Cloud Tools for Visual Studio":     5,
	// "Cloud Transition API":              5,
	// "Cloud VPN":                         5,
	// "Cloud Video Intelligence API":      5,
	// "Cloud Vision API":                  5,
	// "Compute Engine":                    5,
	// "Container Registry":                5,
	// "Container-Optimized OS":            5,
	// "Debugger":                          5,
	// "Dedicated Interconnect":            5,
	// "Developer Portal":                  5,
	// "Dialogflow Enterprise Edition":     5,
	// "Error Reporting":                   5,
	// "GPU":                               5,
	// "Generic GCP Product":               5,
	// "Genomics":                          5,
	// "Gradle Engine Plugin":              5,
	// "IDE Plugins":                       5,
	// "Key Management Service":            5,
	// "Kubernetes Engine":                 5,
	// "Logging":                           5,
	// "Maven App Engine Plugin":           5,
	// "Monitoring":                        5,
	// "Partner Interconnect":              5,
	// "Persistent Disk":                   5,
	// "Premium Network Tier":              5,
	// "Profiler":                          5,
	// "Recommendations AI":                5,
	// "Stackdriver":                       5,
	// "Standard Network Tier":             5,
	// "Trace":                             5,
	// "Traffic Director":                  5,
	// "Transfer Appliance":                5,
	// "Virtual Private Cloud":             5,

	// "Cloud Scheduler": 6,

	/****************************************************************************************************/

	/*** GCP / PATHS ***/

	// skip for now

	/****************************************************************************************************/

	/*** GCP / ZONES ***/

	"Firewall": 8,

	/****************************************************************************************************/

	/*** GCP / EXPANDED PRODUCT CARDS ***/

	// skip for now

	/****************************************************************************************************/

	/*** GCP / GENERAL ICONS ***/

	"General": 1,
}

var ShapeMap = map[string]string{

	//// DUPLICATES IN GROUPS REMOVED ///// ******************************************************************************************************************************************
	//// STILL NEED TO CHECK FOR DUPLICATES ALTOGETHER ///// ****************************************************************************************************************************************

	/*** GCP / SERVICE CARDS ***/

	"Application System(s)":          "shape=mxgraph.gcp2.application_system",             // "Application&#xa;System(s)"
	"Bucket":                         "shape=mxgraph.gcp2.bucket",                         // "Bucket"
	"CDN Interconnect":               "shape=mxgraph.gcp2.google_network_edge_cache",      // "CDN&#xa;Interconnect"
	"Cluster":                        "shape=mxgraph.gcp2.cluster",                        // "Cluster"
	"Dedicated Game Server":          "shape=mxgraph.gcp2.dedicated_game_server",          // "Dedicated&#xa;Game Server"
	"External Payment Form":          "shape=mxgraph.gcp2.external_payment_form",          // "External&#xa;Payment Form"
	"Frontend Platform Services":     "shape=mxgraph.gcp2.frontend_platform_services",     // "Frontend&#xa;Platform Services"
	"Gateway":                        "shape=mxgraph.gcp2.gateway",                        // "Gateway"
	"Google Edge POP":                "shape=mxgraph.gcp2.google_network_edge_cache",      // "Google&#xa;Edge POP"
	"Google Network w/ Edge Cache":   "shape=mxgraph.gcp2.google_network_edge_cache",      // "Google&#xa; Network W/&#xa;Edge Cache"
	"HTTPS Load Balancer":            "shape=mxgraph.gcp2.network_load_balancer",          // "HTTPS&#xa;Load Balancer"
	"Image Services":                 "shape=mxgraph.gcp2.image_services",                 // "Image&#xa;Services"
	"Internal Payment Authorization": "shape=mxgraph.gcp2.internal_payment_authorization", // "Internal Payment&#xa;Authorization"
	"Logs API":                       "shape=mxgraph.gcp2.logs_api",                       // "Logs API"
	"Memcache":                       "shape=mxgraph.gcp2.memcache",                       // "Memcache"
	"NAT":                            "shape=mxgraph.gcp2.nat",                            // "NAT"
	"Network Load Balancer":          "shape=mxgraph.gcp2.network_load_balancer",          // "Network&#xa;Load&#xa;Balancer"
	"Persistent Disk Snapshot":       "shape=mxgraph.gcp2.persistent_disk_snapshot",       // "Persistent&#xa;Disk Snapshot"
	// "Persistent Disk":                "shape=mxgraph.gcp2.persistent_disk_snapshot",       // "Persistent&#xa;Disk"
	"Push Notification Service": "shape=mxgraph.gcp2.push_notification_service", // "Push&#xa;Notification&#xa;Service"
	"Scheduled Tasks":           "shape=mxgraph.gcp2.scheduled_tasks",           // "Scheduled&#xa;Tasks"
	"Service Discovery":         "shape=mxgraph.gcp2.service_discovery",         // "Service Discovery"
	"Squid Proxy":               "shape=mxgraph.gcp2.squid_proxy",               // "Squid Proxy"
	"Task Queues":               "shape=mxgraph.gcp2.task_queues",               // "Task&#xa;Queues"
	"VPN Gateway":               "shape=mxgraph.gcp2.gateway",                   // "VPN Gateway"
	"Virtual File System":       "shape=mxgraph.gcp2.virtual_file_system",       // "Virtual&#xa;File System"

	/****************************************************************************************************/

	/*** GCP / USER AND DEVICE CARDS ***/

	"Application":        "shape=mxgraph.gcp2.application",        // "Application"
	"Beacon":             "shape=mxgraph.gcp2.beacon",             // "Beacon"
	"Circuit-Board":      "shape=mxgraph.gcp2.circuit_board",      // "Circuit-Board"
	"Database":           "shape=mxgraph.gcp2.database",           // "Database"
	"Desktop":            "shape=mxgraph.gcp2.desktop",            // "Desktop"
	"Desktop and Mobile": "shape=mxgraph.gcp2.desktop_and_mobile", // "Desktop and Mobile"
	"Game":               "shape=mxgraph.gcp2.game",               // "Game"
	// "Gateway":            "shape=mxgraph.gcp2.gateway_icon",        // "Gateway"
	"Laptop":         "shape=mxgraph.gcp2.laptop",              // "Laptop"
	"Lightbulb":      "shape=mxgraph.gcp2.lightbulb",           // "Lightbulb"
	"List":           "shape=mxgraph.gcp2.list",                // "List"
	"Live":           "shape=mxgraph.gcp2.live",                // "Live"
	"Local-Compute":  "shape=mxgraph.gcp2.compute_engine_icon", // "Local-Compute"
	"Mobile Devices": "shape=mxgraph.gcp2.mobile_devices",      // "Mobile Devices"
	"Payment":        "shape=mxgraph.gcp2.payment",             // "Payment"
	"Phone":          "shape=mxgraph.gcp2.phone",               // "Phone"
	"Record":         "shape=mxgraph.gcp2.record",              // "Record"
	"Report":         "shape=mxgraph.gcp2.report",              // "Report"
	"Retail":         "shape=mxgraph.gcp2.retail",              // "Retail"
	"Speaker":        "shape=mxgraph.gcp2.speaker",             // "Speaker"
	"Storage":        "shape=mxgraph.gcp2.storage",             // "Storage"
	"Stream":         "shape=mxgraph.gcp2.stream",              // "Stream"
	"Users":          "shape=mxgraph.gcp2.users",               // "Users"
	"Webcam":         "shape=mxgraph.gcp2.webcam",              // "Webcam"

	/****************************************************************************************************/

	/*** GCP / API MANAGEMENT ***/
	/*** GCP / CLOUD AI ***/
	/*** GCP / COMPUTE ***/
	/*** GCP / DATA ANALYTICS ***/
	/*** GCP / DATA TRANSFER ***/
	/*** GCP / DATABASES ***/
	/*** GCP / DEVELOPER TOOLS ***/
	/*** GCP / INTERNET OF THINGS ***/
	/*** GCP / MANAGEMENT TOOLS ***/
	/*** GCP / NETWORKING ***/
	/*** GCP / SECURITY ***/
	/*** GCP / STORAGE ***/

	"API Analytics":                 "shape=mxgraph.gcp2.hexIcon;prIcon=api_analytics",                 // "API&#xa;Analytics"
	"API Monetization":              "shape=mxgraph.gcp2.hexIcon;prIcon=api_monetization",              // "API&#xa;Monetization"
	"Advanced Solutions Lab":        "shape=mxgraph.gcp2.hexIcon;prIcon=advanced_solutions_lab",        // "Advanced&#xa;Solutions Lab"
	"Apigee API Platform":           "shape=mxgraph.gcp2.hexIcon;prIcon=apigee_api_platform",           // "Apigee API&#xa;Platform"
	"Apigee Sense":                  "shape=mxgraph.gcp2.hexIcon;prIcon=apigee_sense",                  // "Apigee&#xa;Sense"
	"App Engine":                    "shape=mxgraph.gcp2.hexIcon;prIcon=app_engine",                    // "App&#xa;Engine"
	"BeyondCorp":                    "shape=mxgraph.gcp2.hexIcon;prIcon=beyondcorp",                    // "BeyondCorp"
	"BigQuery":                      "shape=mxgraph.gcp2.hexIcon;prIcon=bigquery",                      // "BigQuery"
	"Billing API":                   "shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",                   // "Billing&#xa;API"
	"Cloud APIs":                    "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_apis",                    // "Cloud&#xa;APIs"
	"Cloud Armor":                   "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_armor",                   // "Cloud&#xa;Armor"
	"Cloud Bigtable":                "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_bigtable",                // "Cloud&#xa;Bigtable"
	"Cloud Build":                   "shape=mxgraph.gcp2.hexIcon;prIcon=container_builder",             // "Cloud&#xa;Build"
	"Cloud CDN":                     "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_cdn",                     // "Cloud &#xa;CDN"
	"Cloud Composer":                "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_composer",                // "Cloud&#xa;Composer"
	"Cloud Console":                 "shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",                   // "Cloud&#xa;Console"
	"Cloud DNS":                     "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_dns",                     // "Cloud &#xa;DNS"
	"Cloud Data Catalog":            "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_data_catalog",            // "Cloud Data Catalog"
	"Cloud Data Fusion":             "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_data_fusion",             // "Cloud Data Fusion"
	"Cloud Dataflow":                "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_dataflow",                // "Cloud&#xa;Dataflow"
	"Cloud Datalab":                 "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_datalab",                 // "Cloud&#xa;Datalab"
	"Cloud Dataprep":                "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_dataprep",                // "Cloud&#xa;Dataprep"
	"Cloud Dataproc":                "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_dataproc",                // "Cloud&#xa;Dataproc"
	"Cloud Datastore":               "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_datastore",               // "Cloud&#xa;Datastore"
	"Cloud Endpoints":               "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_endpoints",               // "Cloud&#xa;Endpoints"
	"Cloud External IP Addresses":   "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_external_ip_addresses",   // "Cloud External&#xa;IP Addresses"
	"Cloud Firestore":               "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_filestore",               // "Cloud&#xa;Filestore"
	"Cloud Firewall Rules":          "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_firewall_rules",          // "Cloud&#xa;Firewall Rules"
	"Cloud Functions":               "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_functions",               // "Cloud&#xa;Functions"
	"Cloud IAM":                     "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_iam",                     // "Cloud&#xa;IAM"
	"Cloud IoT Core":                "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_iot_core",                // "Cloud&#xa;IoT Core"
	"Cloud Load Balancing":          "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_load_balancing",          // "Cloud Load&#xa;Balancing"
	"Cloud Machine Learning":        "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_machine_learning",        // "Cloud Machine&#xa;Learning"
	"Cloud Memorystore":             "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_memorystore",             // "Cloud&#xa;Memorystore"
	"Cloud Mobile App":              "shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",                   // "Cloud&#xa;Mobile App"
	"Cloud Network":                 "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_network",                 // "Cloud&#xa;Network"
	"Cloud Pub/Sub":                 "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_pubsub",                  // "Cloud&#xa;Pub/Sub"
	"Cloud Resource Manager":        "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_iam",                     // "Cloud Resource&#xa;Manager"
	"Cloud Router":                  "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_router",                  // "Cloud&#xa;Router"
	"Cloud Routes":                  "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_routes",                  // "Cloud&#xa;Routes"
	"Cloud Run":                     "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_run",                     // "Cloud Run"
	"Cloud SDK":                     "shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",                   // "Cloud&#xa;SDK"
	"Cloud SQL":                     "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_sql",                     // "Cloud&#xa;SQL"
	"Cloud Security Command Center": "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_security_command_center", // "Cloud Security&#xa;Command Center"
	"Cloud Security Scanner":        "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_security_scanner",        // "Cloud Security&#xa;Scanner"
	"Cloud Shell":                   "shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",                   // "Cloud&#xa;Shell"
	"Cloud Source Repositories":     "shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",                   // "Cloud Source&#xa;Repositories"
	"Cloud Spanner":                 "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_spanner",                 // "Cloud&#xa;Spanner"
	"Cloud Storage":                 "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_storage",                 // "Cloud&#xa;Storage"
	"Cloud Test Lab":                "shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",                   // "Cloud&#xa;Test Lab"
	"Cloud Tools for Eclipse":       "shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",                   // "Cloud Tools&#xa;for Eclipse"
	"Cloud Tools for IntelliJ":      "shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",                   // "Cloud Tools&#xa;for IntelliJ"
	"Cloud Tools for PowerShell":    "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_tools_for_powershell",    // "Cloud Tools for&#xa;PowerShell"
	"Cloud Tools for Visual Studio": "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_tools_for_powershell",    // "Cloud Tools for&#xa;Visual Studio"
	"Cloud VPN":                     "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_vpn",                     // "Cloud&#xa;VPN"
	"Cloud Video Intelligence API":  "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_video_intelligence_api",  // "Cloud Video&#xa;Intelligence API"
	"Compute Engine":                "shape=mxgraph.gcp2.hexIcon;prIcon=compute_engine",                // "Compute&#xa;Engine"
	"Container Registry":            "shape=mxgraph.gcp2.hexIcon;prIcon=container_registry",            // "Container&#xa;Registry"
	"Container-Optimized OS":        "shape=mxgraph.gcp2.hexIcon;prIcon=container_optimized_os",        // "Container-&#xa;Optimized OS"
	"Data Loss Prevention API":      "shape=mxgraph.gcp2.hexIcon;prIcon=data_loss_prevention_api",      // "Data Loss&#xa;Prevention API"
	"Debugger":                      "shape=mxgraph.gcp2.hexIcon;prIcon=debugger",                      // "Debugger"
	"Dedicated Interconnect":        "shape=mxgraph.gcp2.hexIcon;prIcon=dedicated_interconnect",        // "Dedicated&#xa;Interconnect"
	"Deployment Manager":            "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_deployment_manager",      // "Deployment&#xa;Manager"
	"Developer Portal":              "shape=mxgraph.gcp2.hexIcon;prIcon=developer_portal",              // "Developer&#xa;Portal"
	"Error Reporting":               "shape=mxgraph.gcp2.hexIcon;prIcon=error_reporting",               // "Error&#xa;Reporting"
	"GKE-on-Prem":                   "shape=mxgraph.gcp2.hexIcon;prIcon=gke_on_prem",                   // "GKE-on-Prem"
	"GPU":                           "shape=mxgraph.gcp2.hexIcon;prIcon=gpu",                           // "GPU"
	"Genomics":                      "shape=mxgraph.gcp2.hexIcon;prIcon=genomics",                      // "Genomics"
	"Gradle App Engine Plugin":      "shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",                   // "Gradle App&#xa;Engine Plugin"
	"IDE Plugins":                   "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_tools_for_powershell",    // "IDE Plugins"
	"Identity-Aware Proxy":          "shape=mxgraph.gcp2.hexIcon;prIcon=identity_aware_proxy",          // "Identity-Aware&#xa;Proxy"
	"Jobs API":                      "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_jobs_api",                // "Jobs&#xa;API"
	"Key Management Service":        "shape=mxgraph.gcp2.hexIcon;prIcon=key_management_service",        // "Key Management&#xa;Service"
	"Kubernetes Engine":             "shape=mxgraph.gcp2.hexIcon;prIcon=container_engine",              // "Kubernetes&#xa;Engine"
	"Logging":                       "shape=mxgraph.gcp2.hexIcon;prIcon=logging",                       // "Logging"
	"Maven App Engine Plugin":       "shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",                   // "Maven App&#xa;Engine Plugin"
	"Monitoring":                    "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_deployment_manager",      // "Monitoring"
	"Natural Language API":          "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_natural_language_api",    // "Natural&#xa;Language API"
	"Partner Interconnect":          "shape=mxgraph.gcp2.hexIcon;prIcon=partner_interconnect",          // "Partner&#xa;Interconnect"
	"Persistent Disk":               "shape=mxgraph.gcp2.hexIcon;prIcon=persistent_disk",               // "Persistent&#xa;Disk"
	"Premium Network Tier":          "shape=mxgraph.gcp2.hexIcon;prIcon=premium_network_tier",          // "Premium&#xa;Network Tier"
	"Profiler":                      "shape=mxgraph.gcp2.hexIcon;prIcon=profiler",                      // "Profiler"
	"Security Key Enforcement":      "shape=mxgraph.gcp2.hexIcon;prIcon=security_key_enforcement",      // "Security Key&#xa;Enforcement"
	"Speech API":                    "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_speech_api",              // "Speech&#xa;API"
	"Stackdriver":                   "shape=mxgraph.gcp2.hexIcon;prIcon=stackdriver",                   // "Stackdriver"
	"Standard Network Tier":         "shape=mxgraph.gcp2.hexIcon;prIcon=standard_network_tier",         // "Standard&#xa;Network Tier"
	"Trace":                         "shape=mxgraph.gcp2.hexIcon;prIcon=trace",                         // "Trace"
	"Transfer Appliance":            "shape=mxgraph.gcp2.hexIcon;prIcon=transfer_appliance",            // "Transfer&#xa;Appliance"
	"Translation API":               "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_translation_api",         // "Translation&#xa;API"
	"Virtual Private Cloud":         "shape=mxgraph.gcp2.hexIcon;prIcon=virtual_private_cloud",         // "Virtual&#xa;Private Cloud"
	"Vision API":                    "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_vision_api",              // "Vision&#xa;API"

	/****************************************************************************************************/

	/*** GCP / PRODUCT CARDS ***/

	"AdMob":                     "shape=mxgraph.gcp2.admob",                    // "AdMob"
	"Avere Physical Appliance":  "shape=mxgraph.gcp2.avere",                    // "Avere Physical&#xa;Appliance"
	"Campaign Manager":          "shape=mxgraph.gcp2.campaign_manager",         // "Campaign&#xa;Manager"
	"Fastly":                    "shape=mxgraph.gcp2.fastly",                   // "Fastly"
	"Firebase":                  "shape=mxgraph.gcp2.firebase",                 // "Firebase"
	"Forseti Security":          "shape=mxgraph.gcp2.forseti_logo",             // "Forseti&#xa;Security"
	"Google Ad Manager":         "shape=mxgraph.gcp2.google_ad_manager",        // "Google Ad&#xa;Manager"
	"Google Ads":                "shape=mxgraph.gcp2.google_ads",               // "Google&#xa;Ads"
	"Google Analytics 360":      "shape=mxgraph.gcp2.google_analytics_360",     // "Google&#xa;Analytics 360"
	"Google Analytics":          "shape=mxgraph.gcp2.google_analytics",         // "Google&#xa;Analytics"
	"Google Play Game Services": "shape=mxgraph.gcp2.google_play_game_service", // "Google Play&#xa;Game Services"
	"Istio":                     "shape=mxgraph.gcp2.istio_logo",               // "Istio"
	"Kubernetes":                "shape=mxgraph.gcp2.kubernetes_logo",          // "Kubernetes"
	"TensorFlow":                "shape=mxgraph.gcp2.tensorflow_logo",          // "TensorFlow"

	/****************************************************************************************************/

	/*** GCP ICONS ***/

	// "AI Hub":                            "shape=mxgraph.gcp2.hexIcon;prIcon=ai_hub",                        // "AI Hub"
	// "AI Platform Data Labeling Service": "shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",                   // "AI Platform&#xa;Data Labeling&#xa;Service"
	// "AI Platform":                       "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_machine_learning",        // "AI Platform"
	// "API Analytics":                     "shape=mxgraph.gcp2.hexIcon;prIcon=api_analytics",                 // "API&#xa;Analytics"
	// "API Monetization":                  "shape=mxgraph.gcp2.hexIcon;prIcon=api_monetization",              // "API&#xa;Monetization"
	// "Advanced Solutions Lab":            "shape=mxgraph.gcp2.hexIcon;prIcon=advanced_solutions_lab",        // "Advanced&#xa;Solutions Lab"
	// "Apigee API Platform":               "shape=mxgraph.gcp2.hexIcon;prIcon=apigee_api_platform",           // "Apigee API&#xa;Platform"
	// "Apigee Sense":                      "shape=mxgraph.gcp2.hexIcon;prIcon=apigee_sense",                  // "Apigee&#xa;Sense"
	// "App Engine":                        "shape=mxgraph.gcp2.hexIcon;prIcon=app_engine",                    // "App&#xa;Engine"
	// "AutoML Natural Language":           "shape=mxgraph.gcp2.hexIcon;prIcon=automl_natural_language",       // "AutoML Natural&#xa;Language"
	// "AutoML Tables":                     "shape=mxgraph.gcp2.hexIcon;prIcon=automl_tables",                 // "AutoML Tables"
	// "AutoML Translation":                "shape=mxgraph.gcp2.hexIcon;prIcon=automl_translation",            // "AutoML&#xa;Translation"
	// "AutoML Video Intelligence":         "shape=mxgraph.gcp2.hexIcon;prIcon=automl_video_intelligence",     // "AutoML Video&#xa;Intelligence"
	// "AutoML Vision":                     "shape=mxgraph.gcp2.hexIcon;prIcon=automl_vision",                 // "AutoML Vision"
	// "BigQuery":                          "shape=mxgraph.gcp2.hexIcon;prIcon=bigquery",                      // "BigQuery"
	// "Cloud APIs":                        "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_apis",                    // "Cloud&#xa;APIs"
	// "Cloud Armor":                       "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_armor",                   // "Cloud Armor"
	// "Cloud AutoML":                      "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_automl",                  // "Cloud&#xa;AutoML"
	// "Cloud Bigtable":                    "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_bigtable",                // "Cloud&#xa;Bigtable"
	// "Cloud Billing API":                 "shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",                   // "Cloud&#xa;Billing API"
	// "Cloud Build":                       "shape=mxgraph.gcp2.hexIcon;prIcon=container_builder",             // "Cloud Build"
	// "Cloud CDN":                         "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_cdn",                     // "Cloud&#xa;CDN"
	// "Cloud Code":                        "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_code",                    // "Cloud Code"
	// "Cloud Composer":                    "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_composer",                // "Cloud&#xa;Composer"
	// "Cloud Console":                     "shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",                   // "Cloud&#xa;Console"
	// "Cloud DNS":                         "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_dns",                     // "Cloud&#xa;DNS"
	// "Cloud Data Catalog":                "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_data_catalog",            // "Cloud Data Catalog"
	// "Cloud Data Fusion":                 "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_data_fusion",             // "Cloud Data&#xa;Fusion"
	// "Cloud Dataflow":                    "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_dataflow",                // "Cloud&#xa;Dataflow"
	// "Cloud Datalab":                     "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_datalab",                 // "Cloud&#xa;Datalab"
	// "Cloud Dataprep":                    "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_dataprep",                // "Cloud&#xa;Dataprep"
	// "Cloud Dataproc":                    "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_dataproc",                // "Cloud&#xa;Dataproc"
	// "Cloud Datastore":                   "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_datastore",               // "Cloud&#xa;Datastore"
	// "Cloud Deployment Manager":          "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_deployment_manager",      // "Cloud&#xa;Deployment&#xa;Manager"
	// "Cloud Endpoints":                   "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_endpoints",               // "Cloud&#xa;Endpoints"
	// "Cloud External IP Addresses":       "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_external_ip_addresses",   // "Cloud&#xa;External IP&#xa;Addresses"
	// "Cloud Firestore":                   "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_filestore",               // "Cloud&#xa;Filestore"
	// "Cloud Firewall Rules":              "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_firewall_rules",          // "Cloud&#xa;Firewall Rules"
	// "Cloud Functions":                   "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_functions",               // "Cloud&#xa;Functions"
	// "Cloud IAM":                         "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_iam",                     // "Cloud IAM"
	// "Cloud Interface API":               "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_inference_api",           // "Cloud&#xa;Inference API"
	// "Cloud IoT Core":                    "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_iot_core",                // "Cloud IoT&#xa;Core"
	// "Cloud Jobs API":                    "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_jobs_api",                // "Cloud&#xa;Jobs API"
	// "Cloud Load Balancing":              "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_load_balancing",          // "Cloud Load&#xa;Balancing"
	// "Cloud Memorystore":                 "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_memorystore",             // "Cloud&#xa;Memorystore"
	// "Cloud Mobile App":                  "shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",                   // "Cloud Mobile&#xa;App"
	// "Cloud NAT":                         "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_nat",                     // "Cloud NAT"
	// "Cloud Natural Language API":        "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_natural_language_api",    // "Cloud Natural&#xa;Language API"
	// "Cloud Network":                     "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_network",                 // "Cloud&#xa;Network"
	// "Cloud Pub/Sub":                     "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_pubsub",                  // "Cloud&#xa;Pub/Sub"
	// "Cloud Resource Manager":            "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_iam",                     // "Cloud Resource&#xa;Manager"
	// "Cloud Router":                      "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_router",                  // "Cloud&#xa;Router"
	// "Cloud Routes":                      "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_routes",                  // "Cloud&#xa;Routes"
	// "Cloud Run":                         "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_run",                     // "Cloud Run"
	// "Cloud Run (On Prem)":               "shape=mxgraph.gcp2.hexIcon;prIcon=gke_on_prem",                   // "Cloud Run&#xa;(On Prem)"
	// "Cloud SDK":                         "shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",                   // "Cloud SDK"
	// "Cloud SQL":                         "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_sql",                     // "Cloud SQL"
	// "Cloud Security Command Center":     "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_security_command_center", // "Cloud Security&#xa;Command&#xa;Center"
	// "Cloud Security Scanner":            "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_security_scanner",        // "Cloud Security&#xa;Scanner"
	// "Cloud Service Mesh":                "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_service_mesh",            // "Cloud Service Mesh"
	// "Cloud Shell":                       "shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",                   // "Cloud&#xa;Shell"
	// "Cloud Source Repositories":         "shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",                   // "Cloud Source&#xa;Repositories"
	// "Cloud Spanner":                     "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_spanner",                 // "Cloud&#xa;Spanner"
	// "Cloud Speech-to-Text":              "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_speech_api",              // "Cloud&#xa;Speech-to-Text"
	// "Cloud Storage":                     "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_storage",                 // "Cloud&#xa;Storage"
	// "Cloud TPU":                         "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_tpu",                     // "Cloud TPU"
	// "Cloud Tasks":                       "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_tasks",                   // "Cloud Tasks"
	// "Cloud Test Lab":                    "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_test_lab",                // "Cloud&#xa;Test Lab"
	// "Cloud Text-to-Speech":              "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_text_to_speech",          // "Cloud&#xa;Text-to-Speech"
	// "Cloud Tools for Eclipse":           "shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",                   // "Cloud Tools&#xa;for Eclipse"
	// "Cloud Tools for IntelliJ":          "shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",                   // "Cloud Tools&#xa;for IntelliJ"
	// "Cloud Tools for PowerShell":        "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_tools_for_powershell",    // "Cloud&#xa;Tools for&#xa;PowerShell"
	// "Cloud Tools for Visual Studio":     "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_tools_for_powershell",    // "Cloud&#xa;Tools for&#xa;Visual Studio"
	// "Cloud Transition API":              "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_translation_api",         // "Cloud&#xa;Translation&#xa;API"
	// "Cloud VPN":                         "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_vpn",                     // "Cloud VPN"
	// "Cloud Video Intelligence API":      "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_video_intelligence_api",  // "Cloud Video&#xa;Intelligence&#xa;API"
	// "Cloud Vision API":                  "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_vision_api",              // "Cloud&#xa;Vision API"
	// "Compute Engine":                    "shape=mxgraph.gcp2.hexIcon;prIcon=compute_engine",                // "Compute&#xa;Engine"
	// "Container Registry":                "shape=mxgraph.gcp2.hexIcon;prIcon=container_registry",            // "Container&#xa;Registry"
	// "Container-Optimized OS":            "shape=mxgraph.gcp2.hexIcon;prIcon=container_optimized_os",        // "Container-&#xa;Optimized OS"
	// "Debugger":                          "shape=mxgraph.gcp2.hexIcon;prIcon=debugger",                      // "Debugger"
	// "Dedicated Interconnect":            "shape=mxgraph.gcp2.hexIcon;prIcon=dedicated_interconnect",        // "Dedicated&#xa;Interconnect"
	// "Developer Portal":                  "shape=mxgraph.gcp2.hexIcon;prIcon=developer_portal",              // "Developer&#xa;Portal"
	// "Dialogflow Enterprise Edition":     "shape=mxgraph.gcp2.hexIcon;prIcon=dialogflow_enterprise_edition", // "Dialogflow&#xa;Enterprise&#xa;Edition"
	// "Error Reporting":                   "shape=mxgraph.gcp2.hexIcon;prIcon=error_reporting",               // "Error&#xa;Reporting"
	// "GPU":                               "shape=mxgraph.gcp2.hexIcon;prIcon=gpu",                           // "GPU"
	// "Generic GCP Product":               "shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",                   // "Generic GCP&#xa;Product"
	// "Genomics":                          "shape=mxgraph.gcp2.hexIcon;prIcon=genomics",                      // "Genomics"
	// "Gradle Engine Plugin":              "shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",                   // "Gradle App&#xa;Engine Plugin"
	// "IDE Plugins":                       "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_tools_for_powershell",    // "IDE Plugins"
	// "Key Management Service":            "shape=mxgraph.gcp2.hexIcon;prIcon=key_management_service",        // "Key&#xa;Management&#xa;Service"
	// "Kubernetes Engine":                 "shape=mxgraph.gcp2.hexIcon;prIcon=container_engine",              // "Kubernetes&#xa;Engine"
	// "Logging":                           "shape=mxgraph.gcp2.hexIcon;prIcon=logging",                       // "Logging"
	// "Maven App Engine Plugin":           "shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",                   // "Maven App&#xa;Engine Plugin"
	// "Monitoring":                        "shape=mxgraph.gcp2.hexIcon;prIcon=cloud_deployment_manager",      // "Monitoring"
	// "Partner Interconnect":              "shape=mxgraph.gcp2.hexIcon;prIcon=partner_interconnect",          // "Partner&#xa;Interconnect"
	// "Persistent Disk":                   "shape=mxgraph.gcp2.hexIcon;prIcon=persistent_disk",               // "Persistent&#xa;Disk"
	// "Premium Network Tier":              "shape=mxgraph.gcp2.hexIcon;prIcon=premium_network_tier",          // "Premium&#xa;Network Tier"
	// "Profiler":                          "shape=mxgraph.gcp2.hexIcon;prIcon=profiler",                      // "Profiler"
	// "Recommendations AI":                "shape=mxgraph.gcp2.hexIcon;prIcon=recommendations_ai",            // "Recommendations&#xa;AI"
	// "Stackdriver":                       "shape=mxgraph.gcp2.hexIcon;prIcon=stackdriver",                   // "Stackdriver"
	// "Standard Network Tier":             "shape=mxgraph.gcp2.hexIcon;prIcon=standard_network_tier",         // "Standard&#xa;Network Tier"
	// "Trace":                             "shape=mxgraph.gcp2.hexIcon;prIcon=trace",                         // "Trace"
	// "Traffic Director":                  "shape=mxgraph.gcp2.hexIcon;prIcon=traffic_director",              // "Traffic Director"
	// "Transfer Appliance":                "shape=mxgraph.gcp2.hexIcon;prIcon=transfer_appliance",            // "Transfer&#xa;Appliance"
	// "Virtual Private Cloud":             "shape=mxgraph.gcp2.hexIcon;prIcon=virtual_private_cloud",         // "Virtual&#xa;Private Cloud"

	// "Cloud Scheduler": "shape=mxgraph.gcp2.cloud_scheduler", // "Cloud Scheduler"

	/****************************************************************************************************/

	/*** GCP / SERVICE CARDS (Blank One Line | Blank Two & Three Line) ***/

	// skip for now

	/****************************************************************************************************/

	/*** GCP / PATHS ***/

	"Firewall": "fillColor=#FBE9E7;", // "Firewall"

	/****************************************************************************************************/

	/*** GCP / ZONES ***/

	// skip for now

	/****************************************************************************************************/

	/*** GCP / EXPANDED PRODUCT CARDS ***/

	// skip for now

	/****************************************************************************************************/

	/*** GCP / GENERAL ICONS ***/

	"General": "shape=mxgraph.gcp2.blank", // "Blank Line"
}

var NameMap = map[string]string{

	// TODO format so this looks nice

	"google_compute_instance": "Compute Engine",

	"google_compute_firewall": "Firewall",

	"google_compute_firewall_policy": "Cloud Firewall Rules",

	"google_storage_bucket": "Bucket",

	"google_sql_database": "Cloud SQL",

	"google_container_cluster": "Kubernetes",

	"google_app_engine_application": "App Engine",

	"google_api_gateway_api": "Gateway",
}

/*** Functions to look through file ***/

func LookupName(objectName string) string {
	objectShape, success := NameMap[objectName]

	if objectShape, success := NameMap[objectName]; success {
		return objectShape
	}

	objectShape = "General"
	fmt.Println(objectName, success)
	return objectShape
}

func LookupShape(resourceType string) string {
	objectName, success := NameMap[resourceType]

	if objectName, success := ShapeMap[resourceType]; success {
		return objectName
	}

	fmt.Println(resourceType)
	fmt.Println(objectName, success)

	return "shape=mxgraph.gcp2.blank"
}

func LookupCase(name string) int {
	value, success := CaseMap[name]

	if value, success := CaseMap[name]; success {
		return value
	}

	fmt.Println(name)
	fmt.Println(value, success)
	fmt.Println("Something went wrong. Exiting.")
	os.Exit(1)

	return -1
}
