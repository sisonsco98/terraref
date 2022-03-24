package Shapes

var ShapeMap = map[string]string{

	/*** GCP / PATHS ***/

	"Primary Path":            "dashed=0;strokeColor=#4284F3;",
	"Optional Primary Path":   "dashed=1;dashPattern=1 3;strokeColor=#4284F3;",
	"Secondary Path":          "dashed=0;strokeColor=#9E9E9E;",
	"Optional Secondary Path": "dashed=1;dashPattern=1 3;strokeColor=#9E9E9E;",
	"Success Status":          "strokeColor=#34A853;dashed=0;",
	"Failure Status":          "strokeColor=#EA4335;dashed=0;",

	/****************************************************************************************************/

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

	// "Cloud Scheduler":                   "shape=mxgraph.gcp2.cloud_scheduler",                              // "Cloud Scheduler"

	/****************************************************************************************************/

	/*** GCP / ZONES ***/

	"Firewall":        "fillColor=#FBE9E7;", // "Firewall"
	"Subnetwork":      "fillColor=#EDE7F6;",
	"Network":         "fillColor=#E3F2FD;",
	"Service Project": "fillColor=#F1F8E9;",

	/****************************************************************************************************/

	/*** GCP / EXPANDED PRODUCT CARDS ***/

	// skip for now

	/****************************************************************************************************/

	/*** GCP / GENERAL ICONS ***/

	"General": "shape=mxgraph.gcp2.blank", // "Blank Line"
}