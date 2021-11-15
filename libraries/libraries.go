package libraries

import (
	"fmt"
	"os"
)

var ShapeLookup = map [string] string {

	/*** GCP / SERVICE CARDS ***/

	// "Gateway"								:	"shape=mxgraph.gcp2.gateway",
	// "Memcache"								:	"shape=mxgraph.gcp2.memcache",
	// "Logs API"								:	"shape=mxgraph.gcp2.logs_api",
	// "Cluster"								:	"shape=mxgraph.gcp2.cluster",
	// "NAT"									:	"shape=mxgraph.gcp2.nat",
	// "Squid Proxy"							:	"shape=mxgraph.gcp2.squid_proxy",
	// "Bucket"								:	"shape=mxgraph.gcp2.bucket",
	// "Service Discovery"						:	"shape=mxgraph.gcp2.service_discovery",
	// "Task Queues"							:	"shape=mxgraph.gcp2.task_queues",
	// "Image Services"						:	"shape=mxgraph.gcp2.image_services",
	// "Dedicated Game Server"					:	"shape=mxgraph.gcp2.dedicated_game_server",
	// "Frontend Platform Services"			:	"shape=mxgraph.gcp2.frontend_platform_services",
	// "Google Edge POP"						:	"shape=mxgraph.gcp2.google_network_edge_cache",
	// "External Payment Form"					:	"shape=mxgraph.gcp2.external_payment_form",
	// "Internal Payment Authorization"		:	"shape=mxgraph.gcp2.internal_payment_authorization",
	// "VPN Gateway"							:	"shape=mxgraph.gcp2.gateway",
	// "Application System(s)"					:	"shape=mxgraph.gcp2.application_system",
	// "Virtual File System"					:	"shape=mxgraph.gcp2.virtual_file_system",
	// "CDN Interconnect"						:	"shape=mxgraph.gcp2.google_network_edge_cache",
	// "Scheduled Tasks"						:	"shape=mxgraph.gcp2.scheduled_tasks",
	// "HTTPS Load Balancer"					:	"shape=mxgraph.gcp2.network_load_balancer",
	// "Persistent Disk Snapshot"				:	"shape=mxgraph.gcp2.persistent_disk_snapshot",
	// "Persistent Disk"						:	"shape=mxgraph.gcp2.persistent_disk_snapshot",
	// "Network Load Balancer"					:	"shape=mxgraph.gcp2.network_load_balancer",
	// "Google Network w/ Edge Cache"			:	"shape=mxgraph.gcp2.google_network_edge_cache",
	// "Push Notification Service"				:	"shape=mxgraph.gcp2.push_notification_service",
	// "Blank One Line (w/ bubble)"			:	"shape=mxgraph.gcp2.blank",
	// "Blank One Line (w/o bubble)"			:	"",
	// "Blank Two and Three Line (w/ Bubble)"	:	"shape=mxgraph.gcp2.blank",
	// "Blank Two and Three Line (w/o bubble)"	:	"",

	// /*** GCP / USER AND DEVICE CARDS ***/

	// "Application"							:	"",
	// "Beacon"								:	"",
	// "Circuit Board"							:	"",
	// "Database"								:	"",
	// "Desktop"								:	"",
	// "Desktop and Mobile"					:	"",
	// "Game"									:	"",
	// "Gateway"								:	"",
	// "Laptop"								:	"",
	// "Lightbulb"								:	"",
	// "List"									:	"",
	// "Live"									:	"",
	// "Local-Compute"							:	"",
	// "Mobile Devices"						:	"",
	// "Payment"								:	"",
	// "Phone"									:	"",
	// "Record"								:	"",
	// "Report"								:	"",
	// "Retail"								:	"",
	// "Speaker"								:	"",
	// "Storage"								:	"",
	// "Stream"								:	"",
	// "Users"									:	"",
	// "Webcam"								:	"",

	// /*** GCP / COMPUTE ***/

	// "Compute Engine"						:	"",
	// "GPU"									:	"",
	// "App Engine"							:	"",
	// "Cloud Functions"						:	"",
	// "Kubernetes Engine"						:	"",
	// "Container-Optimized OS"				:	"",
	// "Cloud Run"								:	"",
	// "GKE-on-Prem"							:	"",

	// /*** GCP / API MANAGEMENT ***/

	// "API Analytics"							:	"",
	// "Apigee Sense"							:	"",
	// "API Monetization"						:	"",
	// "Cloud Endpoints"						:	"",
	// "Apigee API Platform"					:	"",
	// "Developer Portal"						:	"",

	// /*** GCP / SECURITY ***/

	// "Cloud IAM"								:	"",
	// "BeyondCorp"							:	"",
	// "Cloud Resource Manager"				:	"",
	// "Data Loss Prevention API"				:	"",
	// "Cloud Security Scanner"				:	"",
	// "Key Management Service"				:	"",
	// "Identity-Aware Proxy"					:	"",
	// "Cloud Security Command Center"			:	"",
	// "Security Key Enforcement"				:	"",

	// /*** GCP / DATA ANALYTICS ***/

	// "BigQuery"								:	"",
	// "Cloud Datalab"							:	"",
	// "Cloud Dataflow"						:	"",
	// "Cloud Pub/Sub"							:	"",
	// "Cloud Dataproc"						:	"",
	// "Genomics"								:	"",
	// "Cloud Dataprep"						:	"",
	// "Cloud Composer"						:	"",
	// "Cloud Data Catalog"					:	"",
	// "Cloud Data Fusion"						:	"",

	// /*** GCP / DATA TRANSFER ***/

	// "Transfer Appliance"					:	"",

	// /*** GCP / CLOUD AI ***/

	// "Cloud Machine Learning"				:	"",
	// "Natural Language API"					:	"",
	// "Vision API"							:	"",
	// "Translation API"						:	"",
	// "Speech API"							:	"",
	// "Jobs API"								:	"",
	// "Cloud Video Intelligence API"			:	"",
	// "Advanced Solutions Lab"				:	"",

	// /*** GCP / INTERNET OF THINGS ***/

	// "Cloud IoT Core"						:	"",

	// /*** GCP / DATABASES ***/

	// "Cloud SQL"								:	"",
	// "Cloud Bigtable"						:	"",
	// "Cloud Spanner"							:	"",
	// "Cloud Memorystore"						:	"",
	// "Cloud Firestore"						:	"",
	// "Cloud Datastore"						:	"",

	// /*** GCP / STORAGE ***/

	// "Cloud Storage"							:	"",
	// "Persistent Disk"						:	"",
	// "Cloud Firestore"						:	"",

	// /*** GCP / MANAGEMENT TOOLS ***/

	// "Stackdriver"							:	"",
	// "Debugger"								:	"",
	// "Monitoring"							:	"",
	// "Deployment Manager"					:	"",
	// "Logging"								:	"",
	// "Cloud Console"							:	"",
	// "Error Reporting"						:	"",
	// "Cloud Shell"							:	"",
	// "Trace"									:	"",
	// "Cloud Mobile App"						:	"",
	// "Profiler"								:	"",
	// "Billing API"							:	"",
	// "Cloud APIs"							:	"",

	// /*** GCP / NETWORKING ***/

	// "Virtual Private Cloud"					:	"",
	// "Dedicated Interconnect"				:	"",
	// "Cloud Load Balancing"					:	"",
	// "Cloud DNS"								:	"",
	// "Cloud CDN"								:	"",
	// "Cloud Network"							:	"",
	// "Cloud External IP Addresses"			:	"",
	// "Cloud Routes"							:	"",
	// "Cloud Firewall Rules"					:	"",
	// "Cloud VPN"								:	"",
	// "Cloud Router"							:	"",
	// "Cloud Armor"							:	"",
	// "Standard Network Tier"					:	"",
	// "Premium Network Tier"					:	"",
	// "Partner Interconnect"					:	"",
	
	// /*** GCP / DEVELOPER TOOLS ***/

	"Cloud Build"							:	"shape=mxgraph.gcp2.hexIcon;prIcon=container_builder",			//	"Cloud&#xa;Build"
	"Cloud SDK"								:	"shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",				//	"Cloud&#xa;SDK"
	"Cloud Source Repositories"				:	"shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",				//	"Cloud Source&#xa;Repositories"
	"Cloud Test Lab"						:	"shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",				//	"Cloud&#xa;Test Lab"
	"Cloud Tools for Eclipse"				:	"shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",				//	"Cloud Tools&#xa;for Eclipse"
	"Cloud Tools for IntelliJ"				:	"shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",				//	"Cloud Tools&#xa;for IntelliJ"
	"Cloud Tools for PowerShell"			:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_tools_for_powershell",	//	"Cloud Tools for&#xa;PowerShell"
	"Cloud Tools for Visual Studio"			:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_tools_for_powershell",	//	"Cloud Tools for&#xa;Visual Studio"
	"Container Registry"					:	"shape=mxgraph.gcp2.hexIcon;prIcon=container_registry",			//	"Container&#xa;Registry"
	"Gradle App Engine Plugin"				:	"shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",				//	"Gradle App&#xa;Engine Plugin"
	"IDE Plugins"							:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_tools_for_powershell",	//	"IDE Plugins"
	"Maven App Engine Plugin"				:	"shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",				//	"Maven App&#xa;Engine Plugin"
	
	/*** GCP / PRODUCT CARDS ***/

	"AdMob"									:   "shape=mxgraph.gcp2.admob",											// "AdMob"
	"Avere Physical Appliance"				:   "shape=mxgraph.gcp2.avere",											// "Avere Physical&#xa;Appliance"
	"Campaign Manager"						:   "shape=mxgraph.gcp2.campaign_manager",								// "Campaign&#xa;Manager"
	"Fastly"								:   "shape=mxgraph.gcp2.fastly",										// "Fastly"
	"Firebase"								:   "shape=mxgraph.gcp2.firebase",										// "Firebase"
	"Forseti Security"						:   "shape=mxgraph.gcp2.forseti_logo",									// "Forseti&#xa;Security"
	"Google Ad Manager"						:   "shape=mxgraph.gcp2.google_ad_manager",								// "Google Ad&#xa;Manager"
	"Google Ads"							:   "shape=mxgraph.gcp2.google_ads",									// "Google&#xa;Ads"
	"Google Analytics 360"					:   "shape=mxgraph.gcp2.google_analytics_360",							// "Google&#xa;Analytics 360"
	"Google Analytics"						:   "shape=mxgraph.gcp2.google_analytics",								// "Google&#xa;Analytics"
	"Google Play Game Services"				:   "shape=mxgraph.gcp2.google_play_game_service",						// "Google Play&#xa;Game Services"
	"Istio"									:   "shape=mxgraph.gcp2.istio_logo",									// "Istio"
	"Kubernetes"							:   "shape=mxgraph.gcp2.kubernetes_logo",								// "Kubernetes"
	"TensorFlow"							:   "shape=mxgraph.gcp2.tensorflow_logo",								// "TensorFlow"
	
	/*** GCP ICONS ***/

	"AI Hub"								:	"shape=mxgraph.gcp2.hexIcon;prIcon=ai_hub",							// "AI Hub"
	"AI Platform Data Labeling Service"		:	"shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",					// "AI Platform&#xa;Data Labeling&#xa;Service"
	"AI Platform"							:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_machine_learning",			// "AI Platform"
	"API Analytics"							:	"shape=mxgraph.gcp2.hexIcon;prIcon=api_analytics",					// "API&#xa;Analytics"
	"API Monetization"						:	"shape=mxgraph.gcp2.hexIcon;prIcon=api_monetization",				// "API&#xa;Monetization"
	"Advanced Solutions Lab"				:	"shape=mxgraph.gcp2.hexIcon;prIcon=advanced_solutions_lab",			// "Advanced&#xa;Solutions Lab"
	"Apigee API Platform"					:	"shape=mxgraph.gcp2.hexIcon;prIcon=apigee_api_platform",			// "Apigee API&#xa;Platform"
	"Apigee Sense"							:	"shape=mxgraph.gcp2.hexIcon;prIcon=apigee_sense",					// "Apigee&#xa;Sense"
	"App Engine"							:	"shape=mxgraph.gcp2.hexIcon;prIcon=app_engine",						// "App&#xa;Engine"
	"AutoML Natural Language"				:	"shape=mxgraph.gcp2.hexIcon;prIcon=automl_natural_language",		// "AutoML Natural&#xa;Language"
	"AutoML Tables"							:	"shape=mxgraph.gcp2.hexIcon;prIcon=automl_tables",					// "AutoML Tables"
	"AutoML Translation"					:	"shape=mxgraph.gcp2.hexIcon;prIcon=automl_translation",				// "AutoML&#xa;Translation"
	"AutoML Video Intelligence"				:	"shape=mxgraph.gcp2.hexIcon;prIcon=automl_video_intelligence",		// "AutoML Video&#xa;Intelligence"
	"AutoML Vision"							:	"shape=mxgraph.gcp2.hexIcon;prIcon=automl_vision",					// "AutoML Vision"
	"BigQuery"								:	"shape=mxgraph.gcp2.hexIcon;prIcon=bigquery",						// "BigQuery"
	"Cloud APIs"							:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_apis",						// "Cloud&#xa;APIs"
	"Cloud Armor"							:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_armor",					// "Cloud Armor"
	"Cloud AutoML"							:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_automl",					// "Cloud&#xa;AutoML"
	"Cloud Bigtable"						:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_bigtable",					// "Cloud&#xa;Bigtable"
	"Cloud Billing API"						:	"shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",					// "Cloud&#xa;Billing API"
	"Cloud Build"							:	"shape=mxgraph.gcp2.hexIcon;prIcon=container_builder",				// "Cloud Build"
	"Cloud CDN"								:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_cdn",						// "Cloud&#xa;CDN"
	"Cloud Code"							:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_code",						// "Cloud Code"
	"Cloud Composer"						:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_composer",					// "Cloud&#xa;Composer"
	"Cloud Console"							:	"shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",					// "Cloud&#xa;Console"
	"Cloud DNS"								:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_dns",						// "Cloud&#xa;DNS"
	"Cloud Data Catalog"					:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_data_catalog",				// "Cloud Data Catalog"
	"Cloud Data Fusion"						:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_data_fusion",				// "Cloud Data&#xa;Fusion"
	"Cloud Dataflow"						:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_dataflow",					// "Cloud&#xa;Dataflow"
	"Cloud Datalab"							:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_datalab",					// "Cloud&#xa;Datalab"
	"Cloud Dataprep"						:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_dataprep",					// "Cloud&#xa;Dataprep"
	"Cloud Dataproc"						:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_dataproc",					// "Cloud&#xa;Dataproc"
	"Cloud Datastore"						:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_datastore",				// "Cloud&#xa;Datastore"
	"Cloud Deployment Manager"				:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_deployment_manager",		// "Cloud&#xa;Deployment&#xa;Manager"
	"Cloud Endpoints"						:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_endpoints",				// "Cloud&#xa;Endpoints"
	"Cloud External IP Addresses"			:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_external_ip_addresses",	// "Cloud&#xa;External IP&#xa;Addresses"
	"Cloud Firestore"						:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_filestore",				// "Cloud&#xa;Filestore"
	"Cloud Firewall Rules"					:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_firewall_rules",			// "Cloud&#xa;Firewall Rules"
	"Cloud Functions"						:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_functions",				// "Cloud&#xa;Functions"
	"Cloud IAM"								:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_iam",						// "Cloud IAM"
	"Cloud Interface API"					:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_inference_api",			// "Cloud&#xa;Inference API"
	"Cloud IoT Core"						:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_iot_core",					// "Cloud IoT&#xa;Core"
	"Cloud Jobs API"						:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_jobs_api",					// "Cloud&#xa;Jobs API"
	"Cloud Load Balancing"					:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_load_balancing",			// "Cloud Load&#xa;Balancing"
	"Cloud Memorystore"						:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_memorystore",				// "Cloud&#xa;Memorystore"
	"Cloud Mobile App"						:	"shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",					// "Cloud Mobile&#xa;App"
	"Cloud NAT"								:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_nat",						// "Cloud NAT"
	"Cloud Natural Language API"			:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_natural_language_api",		// "Cloud Natural&#xa;Language API"
	"Cloud Network"							:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_network",					// "Cloud&#xa;Network"
	"Cloud Pub/Sub"							:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_pubsub",					// "Cloud&#xa;Pub/Sub"
	"Cloud Resource Manager"				:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_iam",						// "Cloud Resource&#xa;Manager"
	"Cloud Router"							:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_router",					// "Cloud&#xa;Router"
	"Cloud Routes"							:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_routes",					// "Cloud&#xa;Routes"
	"Cloud Run"								:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_run",						// "Cloud Run"
	"Cloud Run (On Prem)"					:	"shape=mxgraph.gcp2.hexIcon;prIcon=gke_on_prem",					// "Cloud Run&#xa;(On Prem)"
	"Cloud SDK"								:	"shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",					// "Cloud SDK"
	"Cloud SQL"								:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_sql",						// "Cloud SQL"
	"Cloud Security Command Center"			:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_security_command_center",	// "Cloud Security&#xa;Command&#xa;Center"
	"Cloud Security Scanner"				:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_security_scanner",			// "Cloud Security&#xa;Scanner"
	"Cloud Service Mesh"					:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_service_mesh",				// "Cloud Service Mesh"
	"Cloud Shell"							:	"shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",					// "Cloud&#xa;Shell"
	"Cloud Source Repositories"				:	"shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",					// "Cloud Source&#xa;Repositories"
	"Cloud Spanner"							:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_spanner",					// "Cloud&#xa;Spanner"
	"Cloud Speech-to-Text"					:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_speech_api",				// "Cloud&#xa;Speech-to-Text"
	"Cloud Storage"							:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_storage",					// "Cloud&#xa;Storage"
	"Cloud TPU"								:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_tpu",						// "Cloud TPU"
	"Cloud Tasks"							:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_tasks",					// "Cloud Tasks"
	"Cloud Test Lab"						:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_test_lab",					// "Cloud&#xa;Test Lab"
	"Cloud Text-to-Speech"					:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_text_to_speech",			// "Cloud&#xa;Text-to-Speech"
	"Cloud Tools for Eclipse"				:	"shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",					// "Cloud Tools&#xa;for Eclipse"
	"Cloud Tools for IntelliJ"				:	"shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",					// "Cloud Tools&#xa;for IntelliJ"
	"Cloud Tools for PowerShell"			:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_tools_for_powershell",		// "Cloud&#xa;Tools for&#xa;PowerShell"
	"Cloud Tools for Visual Studio"			:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_tools_for_powershell",		// "Cloud&#xa;Tools for&#xa;Visual Studio"
	"Cloud Transition API"					:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_translation_api",			// "Cloud&#xa;Translation&#xa;API"
	"Cloud VPN"								:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_vpn",						// "Cloud VPN"
	"Cloud Video Intelligence API"			:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_video_intelligence_api",	// "Cloud Video&#xa;Intelligence&#xa;API"
	"Cloud Vision API"						:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_vision_api",				// "Cloud&#xa;Vision API"
	"Compute Engine"						:	"shape=mxgraph.gcp2.hexIcon;prIcon=compute_engine",					// "Compute&#xa;Engine"
	"Container Registry"					:	"shape=mxgraph.gcp2.hexIcon;prIcon=container_registry",				// "Container&#xa;Registry"
	"Container-Optimized OS"				:	"shape=mxgraph.gcp2.hexIcon;prIcon=container_optimized_os",			// "Container-&#xa;Optimized OS"
	"Debugger"								:	"shape=mxgraph.gcp2.hexIcon;prIcon=debugger",						// "Debugger"
	"Dedicated Interconnect"				:	"shape=mxgraph.gcp2.hexIcon;prIcon=dedicated_interconnect",			// "Dedicated&#xa;Interconnect"
	"Developer Portal"						:	"shape=mxgraph.gcp2.hexIcon;prIcon=developer_portal",				// "Developer&#xa;Portal"
	"Dialogflow Enterprise Edition"			:	"shape=mxgraph.gcp2.hexIcon;prIcon=dialogflow_enterprise_edition",	// "Dialogflow&#xa;Enterprise&#xa;Edition"
	"Error Reporting"						:	"shape=mxgraph.gcp2.hexIcon;prIcon=error_reporting",				// "Error&#xa;Reporting"
	"GPU"									:	"shape=mxgraph.gcp2.hexIcon;prIcon=gpu",							// "GPU"
	"Generic GCP Product"					:	"shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",					// "Generic GCP&#xa;Product"
	"Genomics"								:	"shape=mxgraph.gcp2.hexIcon;prIcon=genomics",						// "Genomics"
	"Gradle Engine Plugin"					:	"shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",					// "Gradle App&#xa;Engine Plugin"
	"IDE Plugins"							:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_tools_for_powershell",		// "IDE Plugins"
	"Key Management Service"				:	"shape=mxgraph.gcp2.hexIcon;prIcon=key_management_service",			// "Key&#xa;Management&#xa;Service"
	"Kubernetes Engine"						:	"shape=mxgraph.gcp2.hexIcon;prIcon=container_engine",				// "Kubernetes&#xa;Engine"
	"Logging"								:	"shape=mxgraph.gcp2.hexIcon;prIcon=logging",						// "Logging"
	"Maven App Engine Plugin"				:	"shape=mxgraph.gcp2.hexIcon;prIcon=placeholder",					// "Maven App&#xa;Engine Plugin"
	"Monitoring"							:	"shape=mxgraph.gcp2.hexIcon;prIcon=cloud_deployment_manager",		// "Monitoring"
	"Partner Interconnect"					:	"shape=mxgraph.gcp2.hexIcon;prIcon=partner_interconnect",			// "Partner&#xa;Interconnect"
	"Persistent Disk"						:	"shape=mxgraph.gcp2.hexIcon;prIcon=persistent_disk",				// "Persistent&#xa;Disk"
	"Premium Network Tier"					:	"shape=mxgraph.gcp2.hexIcon;prIcon=premium_network_tier",			// "Premium&#xa;Network Tier"
	"Profiler"								:	"shape=mxgraph.gcp2.hexIcon;prIcon=profiler",						// "Profiler"
	"Recommendations AI"					:	"shape=mxgraph.gcp2.hexIcon;prIcon=recommendations_ai",				// "Recommendations&#xa;AI"
	"Stackdriver"							:	"shape=mxgraph.gcp2.hexIcon;prIcon=stackdriver",					// "Stackdriver"
	"Standard Network Tier"					:	"shape=mxgraph.gcp2.hexIcon;prIcon=standard_network_tier",			// "Standard&#xa;Network Tier"
	"Trace"									:	"shape=mxgraph.gcp2.hexIcon;prIcon=trace",							// "Trace"
	"Traffic Director"						:	"shape=mxgraph.gcp2.hexIcon;prIcon=traffic_director",				// "Traffic Director"
	"Transfer Appliance"					:	"shape=mxgraph.gcp2.hexIcon;prIcon=transfer_appliance",				// "Transfer&#xa;Appliance"
	"Virtual Private Cloud"					:	"shape=mxgraph.gcp2.hexIcon;prIcon=virtual_private_cloud",			// "Virtual&#xa;Private Cloud"

	"Cloud Scheduler"						:	"shape=mxgraph.gcp2.cloud_scheduler",								// "Cloud Scheduler"
		
	/*** GCP / PATHS ***/

		// skip for now

	/*** GCP / ZONES ***/

		// skip for now

	/*** GCP / EXPANDED PRODUCT CARDS ***/

		// skip for now

	/*** GCP / GENERAL ICONS ***/

		// skip for now
}

var classLookup = map [string] int {

	/*** GCP / SERVICE CARDS ***/

	// "Gateway"								:	1,
	// "Memcache"								:	1,
	// "Logs API"								:	1,
	// "Cluster"								:	1,
	// "NAT"									:	1,
	// "Squid Proxy"							:	1,
	// "Bucket"								:	1,
	// "Service Discovery"						:	1,
	// "Task Queues"							:	1,
	// "Image Services"						:	1,
	// "Dedicated Game Server"					:	1,
	// "Frontend Platform Services"			:	1,
	// "Google Edge POP"						:	1,
	// "External Payment Form"					:	1,
	// "Internal Payment Authorization"		:	1,
	// "VPN Gateway"							:	1,
	// "Application System(s)"					:	1,
	// "Virtual File System"					:	1,
	// "CDN Interconnect"						:	1,
	// "Scheduled Tasks"						:	1,
	// "HTTPS Load Balancer"					:	1,
	// "Persistent Disk Snapshot"				:	1,
	// "Persistent Disk"						:	1,
	// "Network Load Balancer"					:	1,
	// "Google Network w/ Edge Cache"			:	1,
	// "Push Notification Service"				:	1,
	// "Blank One Line (w/ bubble)"			:	1,
	// "Blank One Line (w/o bubble)"			:	1,
	// "Blank Two and Three Line (w/ Bubble)"	:	1,
	// "Blank Two and Three Line (w/o bubble)"	:	1,

	// /*** GCP / USER AND DEVICE CARDS ***/

	// "Application"							:	2,
	// "Beacon"								:	2,
	// "Circuit Board"							:	2,
	// "Database"								:	2,
	// "Desktop"								:	2,
	// "Desktop and Mobile"					:	2,
	// "Game"									:	2,
	// "Gateway"								:	2,
	// "Laptop"								:	2,
	// "Lightbulb"								:	2,
	// "List"									:	2,
	// "Live"									:	2,
	// "Local-Compute"							:	2,
	// "Mobile Devices"						:	2,
	// "Payment"								:	2,
	// "Phone"									:	2,
	// "Record"								:	2,
	// "Report"								:	2,
	// "Retail"								:	2,
	// "Speaker"								:	2,
	// "Storage"								:	2,
	// "Stream"								:	2,
	// "Users"									:	2,
	// "Webcam"								:	2,

	// /*** GCP / COMPUTE ***/

	// "Compute Engine"						:	3,
	// "GPU"									:	3,
	// "App Engine"							:	3,
	// "Cloud Functions"						:	3,
	// "Kubernetes Engine"						:	3,
	// "Container-Optimized OS"				:	3,
	// "Cloud Run"								:	3,
	// "GKE-on-Prem"							:	3,

	// /*** GCP / API MANAGEMENT ***/

	// "API Analytics"							:	3,
	// "Apigee Sense"							:	3,
	// "API Monetization"						:	3,
	// "Cloud Endpoints"						:	3,
	// "Apigee API Platform"					:	3,
	// "Developer Portal"						:	3,

	// /*** GCP / SECURITY ***/

	// "Cloud IAM"								:	3,
	// "BeyondCorp"							:	3,
	// "Cloud Resource Manager"				:	3,
	// "Data Loss Prevention API"				:	3,
	// "Cloud Security Scanner"				:	3,
	// "Key Management Service"				:	3,
	// "Identity-Aware Proxy"					:	3,
	// "Cloud Security Command Center"			:	3,
	// "Security Key Enforcement"				:	3,

	// /*** GCP / DATA ANALYTICS ***/

	// "BigQuery"								:	3,
	// "Cloud Datalab"							:	3,
	// "Cloud Dataflow"						:	3,
	// "Cloud Pub/Sub"							:	3,
	// "Cloud Dataproc"						:	3,
	// "Genomics"								:	3,
	// "Cloud Dataprep"						:	3,
	// "Cloud Composer"						:	3,
	// "Cloud Data Catalog"					:	3,
	// "Cloud Data Fusion"						:	3,

	// /*** GCP / DATA TRANSFER ***/

	// "Transfer Appliance"					:	3,

	// /*** GCP / CLOUD AI ***/

	// "Cloud Machine Learning"				:	3,
	// "Natural Language API"					:	3,
	// "Vision API"							:	3,
	// "Translation API"						:	3,
	// "Speech API"							:	3,
	// "Jobs API"								:	3,
	// "Cloud Video Intelligence API"			:	3,
	// "Advanced Solutions Lab"				:	3,

	// /*** GCP / INTERNET OF THINGS ***/

	// "Cloud IoT Core"						:	3,

	// /*** GCP / DATABASES ***/

	// "Cloud SQL"								:	3,
	// "Cloud Bigtable"						:	3,
	// "Cloud Spanner"							:	3,
	// "Cloud Memorystore"						:	3,
	// "Cloud Firestore"						:	3,
	// "Cloud Datastore"						:	3,

	// /*** GCP / STORAGE ***/

	// "Cloud Storage"							:	3,
	// "Persistent Disk"						:	3,
	// "Cloud Firestore"						:	3,

	// /*** GCP / MANAGEMENT TOOLS ***/

	// "Stackdriver"							:	3,
	// "Debugger"								:	3,
	// "Monitoring"							:	3,
	// "Deployment Manager"					:	3,
	// "Logging"								:	3,
	// "Cloud Console"							:	3,
	// "Error Reporting"						:	3,
	// "Cloud Shell"							:	3,
	// "Trace"									:	3,
	// "Cloud Mobile App"						:	3,
	// "Profiler"								:	3,
	// "Billing API"							:	3,
	// "Cloud APIs"							:	3,

	// /*** GCP / NETWORKING ***/

	// "Virtual Private Cloud"					:	3,
	// "Dedicated Interconnect"				:	3,
	// "Cloud Load Balancing"					:	3,
	// "Cloud DNS"								:	3,
	// "Cloud CDN"								:	3,
	// "Cloud Network"							:	3,
	// "Cloud External IP Addresses"			:	3,
	// "Cloud Routes"							:	3,
	// "Cloud Firewall Rules"					:	3,
	// "Cloud VPN"								:	3,
	// "Cloud Router"							:	3,
	// "Cloud Armor"							:	3,
	// "Standard Network Tier"					:	3,
	// "Premium Network Tier"					:	3,
	// "Partner Interconnect"					:	3,
	
//// NEED TO CHECK FOR DUPLICATES /////

	// /*** GCP / DEVELOPER TOOLS ***/

	"Cloud Build"							:	3,
	"Cloud SDK"								:	3,
	"Cloud Source Repositories"				:	3,
	"Cloud Test Lab"						:	3,
	"Cloud Tools for Eclipse"				:	3,
	"Cloud Tools for IntelliJ"				:	3,
	"Cloud Tools for PowerShell"			:	3,
	"Cloud Tools for Visual Studio"			:	3,
	"Container Registry"					:	3,
	"Gradle App Engine Plugin"				:	3,
	"IDE Plugins"							:	3,
	"Maven App Engine Plugin"				:	3,
	
	/*** GCP / PRODUCT CARDS ***/

	"AdMob"									:	4,
	"Avere Physical Appliance"				:	4,
	"Campaign Manager"						:	4,
	"Fastly"								:	4,
	"Firebase"								:	4,
	"Forseti Security"						:	4,
	"Google Ad Manager"						:	4,
	"Google Ads"							:	4,
	"Google Analytics 360"					:	4,
	"Google Analytics"						:	4,
	"Google Play Game Services"				:	4,
	"Istio"									:	4,
	"Kubernetes"							:	4,
	"TensorFlow"							:	4,
	
	/*** GCP ICONS ***/

	"AI Hub"								:	5,
	"AI Platform Data Labeling Service"		:	5,
	"AI Platform"							:	5,
	"API Analytics"							:	5,
	"API Monetization"						:	5,
	"Advanced Solutions Lab"				:	5,
	"Apigee API Platform"					:	5,
	"Apigee Sense"							:	5,
	"App Engine"							:	5,
	"AutoML Natural Language"				:	5,
	"AutoML Tables"							:	5,
	"AutoML Translation"					:	5,
	"AutoML Video Intelligence"				:	5,
	"AutoML Vision"							:	5,
	"BigQuery"								:	5,
	"Cloud APIs"							:	5,
	"Cloud Armor"							:	5,
	"Cloud AutoML"							:	5,
	"Cloud Bigtable"						:	5,
	"Cloud Billing API"						:	5,
	"Cloud Build"							:	5,
	"Cloud CDN"								:	5,
	"Cloud Code"							:	5,
	"Cloud Composer"						:	5,
	"Cloud Console"							:	5,
	"Cloud DNS"								:	5,
	"Cloud Data Catalog"					:	5,
	"Cloud Data Fusion"						:	5,
	"Cloud Dataflow"						:	5,
	"Cloud Datalab"							:	5,
	"Cloud Dataprep"						:	5,
	"Cloud Dataproc"						:	5,
	"Cloud Datastore"						:	5,
	"Cloud Deployment Manager"				:	5,
	"Cloud Endpoints"						:	5,
	"Cloud External IP Addresses"			:	5,
	"Cloud Firestore"						:	5,
	"Cloud Firewall Rules"					:	5,
	"Cloud Functions"						:	5,
	"Cloud IAM"								:	5,
	"Cloud Interface API"					:	5,
	"Cloud IoT Core"						:	5,
	"Cloud Jobs API"						:	5,
	"Cloud Load Balancing"					:	5,
	"Cloud Memorystore"						:	5,
	"Cloud Mobile App"						:	5,
	"Cloud NAT"								:	5,
	"Cloud Natural Language API"			:	5,
	"Cloud Network"							:	5,
	"Cloud Pub/Sub"							:	5,
	"Cloud Resource Manager"				:	5,
	"Cloud Router"							:	5,
	"Cloud Routes"							:	5,
	"Cloud Run"								:	5,
	"Cloud Run (On Prem)"					:	5,
	"Cloud SDK"								:	5,
	"Cloud SQL"								:	5,
	"Cloud Security Command Center"			:	5,
	"Cloud Security Scanner"				:	5,
	"Cloud Service Mesh"					:	5,
	"Cloud Shell"							:	5,
	"Cloud Source Repositories"				:	5,
	"Cloud Spanner"							:	5,
	"Cloud Speech-to-Text"					:	5,
	"Cloud Storage"							:	5,
	"Cloud TPU"								:	5,
	"Cloud Tasks"							:	5,
	"Cloud Test Lab"						:	5,
	"Cloud Text-to-Speech"					:	5,
	"Cloud Tools for Eclipse"				:	5,
	"Cloud Tools for IntelliJ"				:	5,
	"Cloud Tools for PowerShell"			:	5,
	"Cloud Tools for Visual Studio"			:	5,
	"Cloud Transition API"					:	5,
	"Cloud VPN"								:	5,
	"Cloud Video Intelligence API"			:	5,
	"Cloud Vision API"						:	5,
	"Compute Engine"						:	5,
	"Container Registry"					:	5,
	"Container-Optimized OS"				:	5,
	"Debugger"								:	5,
	"Dedicated Interconnect"				:	5,
	"Developer Portal"						:	5,
	"Dialogflow Enterprise Edition"			:	5,
	"Error Reporting"						:	5,
	"GPU"									:	5,
	"Generic GCP Product"					:	5,
	"Genomics"								:	5,
	"Gradle Engine Plugin"					:	5,
	"IDE Plugins"							:	5,
	"Key Management Service"				:	5,
	"Kubernetes Engine"						:	5,
	"Logging"								:	5,
	"Maven App Engine Plugin"				:	5,
	"Monitoring"							:	5,
	"Partner Interconnect"					:	5,
	"Persistent Disk"						:	5,
	"Premium Network Tier"					:	5,
	"Profiler"								:	5,
	"Recommendations AI"					:	5,
	"Stackdriver"							:	5,
	"Standard Network Tier"					:	5,
	"Trace"									:	5,
	"Traffic Director"						:	5,
	"Transfer Appliance"					:	5,
	"Virtual Private Cloud"					:	5,

	"Cloud Scheduler"						:	6,

	/*** GCP / PATHS ***/

		// skip for now

	/*** GCP / ZONES ***/

		// skip for now

	/*** GCP / EXPANDED PRODUCT CARDS ***/

		// skip for now

	/*** GCP / GENERAL ICONS ***/

		// skip for now
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

/*	GCP DEVELOPER TOOLS


	*/

/*	GCP PRODUCT CARDS

	"Kubernetes"                            :   "shape=mxgraph.gcp2.kubernetes_logo",	// "Kubernetes"

	<mxCell id="2" value="" style="strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;" vertex="1" parent="1">
	<mxGeometry x="180" y="40" width="130" height="60" as="geometry" />
	</mxCell>
	<mxCell id="sksrVD8GqV6zCDYqazPW-154" value="Kubernetes" style="sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" vertex="1" parent="sksrVD8GqV6zCDYqazPW-153">
	<mxGeometry width="45" height="43.65" relative="1" as="geometry">
		<mxPoint x="5" y="7.675000000000001" as="offset" />
	</mxGeometry>
	</mxCell>

	"TensorFlow"                            :   "shape=mxgraph.gcp2.tensorflow_logo",	// "TensorFlow"

	<mxCell id="2" value="" style="strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;" vertex="1" parent="1">
	<mxGeometry x="180" y="80" width="130" height="60" as="geometry" />
	</mxCell>
	<mxCell id="sksrVD8GqV6zCDYqazPW-156" value="TensorFlow" style="sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" vertex="1" parent="sksrVD8GqV6zCDYqazPW-155">
	<mxGeometry width="42.3" height="45" relative="1" as="geometry">
		<mxPoint x="5" y="7" as="offset" />
	</mxGeometry>
	</mxCell>

	"Forseti Security"                      :   "shape=mxgraph.gcp2.forseti_logo",	// "Forseti&#xa;Security"

	<mxCell id="2" value="" style="strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;" vertex="1" parent="1">
	<mxGeometry x="250" y="50" width="110" height="60" as="geometry" />
	</mxCell>
	<mxCell id="sksrVD8GqV6zCDYqazPW-158" value="Forseti&#xa;Security" style="sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" vertex="1" parent="sksrVD8GqV6zCDYqazPW-157">
	<mxGeometry width="44.1" height="45" relative="1" as="geometry">
		<mxPoint x="5" y="7" as="offset" />
	</mxGeometry>
	</mxCell>

	"Istio"                                 :   "shape=mxgraph.gcp2.istio_logo",	// "Istio"

	<mxCell id="2" value="" style="strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;" vertex="1" parent="1">
	<mxGeometry x="220" y="140" width="80" height="60" as="geometry" />
	</mxCell>
	<mxCell id="sksrVD8GqV6zCDYqazPW-160" value="Istio" style="sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" vertex="1" parent="sksrVD8GqV6zCDYqazPW-159">
	<mxGeometry width="30.150000000000002" height="45" relative="1" as="geometry">
		<mxPoint x="5" y="7" as="offset" />
	</mxGeometry>
	</mxCell>

	"Firebase"                              :   "shape=mxgraph.gcp2.firebase",	// "Firebase"

	<mxCell id="2" value="" style="strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;" vertex="1" parent="1">
	<mxGeometry x="310" y="190" width="100" height="60" as="geometry" />
	</mxCell>
	<mxCell id="sksrVD8GqV6zCDYqazPW-162" value="Firebase" style="sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" vertex="1" parent="sksrVD8GqV6zCDYqazPW-161">
	<mxGeometry width="32.4" height="45" relative="1" as="geometry">
		<mxPoint x="5" y="7" as="offset" />
	</mxGeometry>
	</mxCell>

	"AdMob"                                 :   "shape=mxgraph.gcp2.admob",	// "AdMob"

	<mxCell id="2" value="" style="strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;" vertex="1" parent="1">
	<mxGeometry x="280" y="180" width="110" height="60" as="geometry" />
	</mxCell>
	<mxCell id="sksrVD8GqV6zCDYqazPW-166" value="AdMob" style="sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" vertex="1" parent="sksrVD8GqV6zCDYqazPW-165">
	<mxGeometry width="45" height="45" relative="1" as="geometry">
		<mxPoint x="5" y="7" as="offset" />
	</mxGeometry>
	</mxCell>

	"Campaign Manager"                      :   "shape=mxgraph.gcp2.campaign_manager",	// "Campaign&#xa;Manager"

	<mxCell id="2" value="" style="strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;" vertex="1" parent="1">
	<mxGeometry x="280" y="190" width="120" height="60" as="geometry" />
	</mxCell>
	<mxCell id="sksrVD8GqV6zCDYqazPW-170" value="Campaign&#xa;Manager" style="sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" vertex="1" parent="sksrVD8GqV6zCDYqazPW-169">
	<mxGeometry width="45" height="45" relative="1" as="geometry">
		<mxPoint x="5" y="7" as="offset" />
	</mxGeometry>
	</mxCell>

	"Google Analytics"                      :   "shape=mxgraph.gcp2.google_analytics",	// "Google&#xa;Analytics"

	<mxCell id="2" value="" style="strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;" vertex="1" parent="1">
	<mxGeometry x="305" y="170" width="120" height="60" as="geometry" />
	</mxCell>
	<mxCell id="sksrVD8GqV6zCDYqazPW-172" value="Google&#xa;Analytics" style="sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" vertex="1" parent="sksrVD8GqV6zCDYqazPW-171">
	<mxGeometry width="45" height="45" relative="1" as="geometry">
		<mxPoint x="5" y="7" as="offset" />
	</mxGeometry>
	</mxCell>

	"Google Ads"                            :   "shape=mxgraph.gcp2.google_ads",	// "Google&#xa;Ads"

	<mxCell id="2" value="" style="strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;" vertex="1" parent="1">
	<mxGeometry x="160" y="180" width="100" height="60" as="geometry" />
	</mxCell>
	<mxCell id="sksrVD8GqV6zCDYqazPW-174" value="Google&#xa;Ads" style="sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" vertex="1" parent="sksrVD8GqV6zCDYqazPW-173">
	<mxGeometry width="45" height="45" relative="1" as="geometry">
		<mxPoint x="5" y="7" as="offset" />
	</mxGeometry>
	</mxCell>

	"Google Analytics 360"                  :   "shape=mxgraph.gcp2.google_analytics_360",	// "Google&#xa;Analytics 360"

	<mxCell id="2" value="" style="strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;" vertex="1" parent="1">
	<mxGeometry x="170" y="200" width="140" height="60" as="geometry" />
	</mxCell>
	<mxCell id="sksrVD8GqV6zCDYqazPW-178" value="Google&#xa;Analytics 360" style="sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" vertex="1" parent="sksrVD8GqV6zCDYqazPW-177">
	<mxGeometry width="45" height="44.1" relative="1" as="geometry">
		<mxPoint x="5" y="7.45" as="offset" />
	</mxGeometry>
	</mxCell>

	"Google Ad Manager"                     :   "shape=mxgraph.gcp2.google_ad_manager",	// "Google Ad&#xa;Manager"

	<mxCell id="2" value="" style="strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;" vertex="1" parent="1">
	<mxGeometry x="340" y="230" width="120" height="60" as="geometry" />
	</mxCell>
	<mxCell id="sksrVD8GqV6zCDYqazPW-180" value="Google Ad&#xa;Manager" style="sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" vertex="1" parent="sksrVD8GqV6zCDYqazPW-179">
	<mxGeometry width="45" height="45" relative="1" as="geometry">
		<mxPoint x="5" y="7" as="offset" />
	</mxGeometry>
	</mxCell>

	"Avere Physical Appliance"              :   "shape=mxgraph.gcp2.avere",	// "Avere Physical&#xa;Appliance"

	<mxCell id="2" value="" style="strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;" vertex="1" parent="1">
	<mxGeometry x="260" y="160" width="150" height="60" as="geometry" />
	</mxCell>
	<mxCell id="sksrVD8GqV6zCDYqazPW-176" value="Avere Physical&#xa;Appliance" style="sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" vertex="1" parent="sksrVD8GqV6zCDYqazPW-175">
	<mxGeometry width="45" height="14.850000000000001" relative="1" as="geometry">
		<mxPoint x="5" y="22.075" as="offset" />
	</mxGeometry>
	</mxCell>

	"Fastly"                                :   "shape=mxgraph.gcp2.fastly",	// "Fastly"

	<mxCell id="2" value="" style="strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;" vertex="1" parent="1">
	<mxGeometry x="160" y="110" width="100" height="60" as="geometry" />
	</mxCell>
	<mxCell id="sksrVD8GqV6zCDYqazPW-182" value="Fastly" style="sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" vertex="1" parent="sksrVD8GqV6zCDYqazPW-181">
	<mxGeometry width="45" height="17.55" relative="1" as="geometry">
		<mxPoint x="5" y="20.725" as="offset" />
	</mxGeometry>
	</mxCell>

	"Google Play Game Services"             :   "shape=mxgraph.gcp2.google_play_game_service",	// "Google Play&#xa;Game Services"

	<mxCell id="2" value="" style="strokeColor=#dddddd;shadow=1;strokeWidth=1;rounded=1;absoluteArcSize=1;arcSize=2;" vertex="1" parent="1">
	<mxGeometry x="140" y="160" width="150" height="60" as="geometry" />
	</mxCell>
	<mxCell id="sksrVD8GqV6zCDYqazPW-168" value="Google Play&#xa;Game Services" style="sketch=0;dashed=0;connectable=0;html=1;fillColor=#5184F3;strokeColor=none;part=1;labelPosition=right;verticalLabelPosition=middle;align=left;verticalAlign=middle;spacingLeft=5;fontColor=#999999;fontSize=12;" vertex="1" parent="sksrVD8GqV6zCDYqazPW-167">
	<mxGeometry width="45" height="31.049999999999997" relative="1" as="geometry">
		<mxPoint x="5" y="13.975000000000001" as="offset" />
	</mxGeometry>
	</mxCell>
	
	*/
