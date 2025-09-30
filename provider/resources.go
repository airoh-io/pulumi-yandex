// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package yandex

import (
	"fmt"
	"path/filepath"
	"unicode"

	"github.com/airoh-io/pulumi-yandex/provider/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/yandex-cloud/terraform-provider-yandex/yandex"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "yandex"
	// modules:
	mainMod = "index" // the y module
)

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeDataSource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the data source's
// first character.
func makeDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeMember(mod+"/"+fn, res)
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(yandex.NewSDKProvider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "yandex",
		Description: "A Pulumi package for creating and managing yandex cloud resources.",
		Keywords:    []string{"pulumi", "yandex"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/airoh-io/pulumi-yandex",
		GitHubOrg:   "yandex-cloud",
		Config:      map[string]*tfbridge.SchemaInfo{},
		Resources: map[string]*tfbridge.ResourceInfo{
			// ALB Resources
			"yandex_alb_target_group":    {Tok: makeResource(mainMod, "AlbTargetGroup")},
			"yandex_alb_backend_group":   {Tok: makeResource(mainMod, "AlbBackendGroup")},
			"yandex_alb_http_router":     {Tok: makeResource(mainMod, "AlbHttpRouter")},
			"yandex_alb_virtual_host":    {Tok: makeResource(mainMod, "AlbVirtualHost")},
			"yandex_alb_load_balancer":   {Tok: makeResource(mainMod, "AlbLoadBalancer")},
			"yandex_api_gateway":         {Tok: makeResource(mainMod, "ApiGateway")},
			"yandex_audit_trails_trail":  {Tok: makeResource(mainMod, "AuditTrailsTrail")},
			"yandex_backup_policy":       {Tok: makeResource(mainMod, "BackupPolicy")},
			"yandex_backup_policy_bindings": {Tok: makeResource(mainMod, "BackupPolicyBindings")},
			"yandex_cdn_origin_group":    {Tok: makeResource(mainMod, "CdnOriginGroup")},
			"yandex_cdn_resource":        {Tok: makeResource(mainMod, "CdnResource")},
			"yandex_cm_certificate":      {Tok: makeResource(mainMod, "CmCertificate")},
			// Compute Resources
			"yandex_compute_disk":                   {Tok: makeResource(mainMod, "ComputeDisk")},
			"yandex_compute_disk_placement_group":   {Tok: makeResource(mainMod, "ComputeDiskPlacementGroup")},
			"yandex_compute_filesystem":             {Tok: makeResource(mainMod, "ComputeFilesystem")},
			"yandex_compute_gpu_cluster":            {Tok: makeResource(mainMod, "ComputeGpuCluster")},
			"yandex_compute_image":                  {Tok: makeResource(mainMod, "ComputeImage")},
			"yandex_compute_instance":               {Tok: makeResource(mainMod, "ComputeInstance")},
			"yandex_compute_instance_group":         {Tok: makeResource(mainMod, "ComputeInstanceGroup")},
			"yandex_compute_placement_group":        {Tok: makeResource(mainMod, "ComputePlacementGroup")},
			"yandex_compute_snapshot":               {Tok: makeResource(mainMod, "ComputeSnapshot")},
			"yandex_compute_snapshot_schedule":      {Tok: makeResource(mainMod, "ComputeSnapshotSchedule")},
			// Container Resources
			"yandex_container_registry":                    {Tok: makeResource(mainMod, "ContainerRegistry")},
			"yandex_container_registry_ip_permission":      {Tok: makeResource(mainMod, "ContainerRegistryIpPermission")},
			"yandex_container_repository":                  {Tok: makeResource(mainMod, "ContainerRepository")},
			"yandex_container_repository_lifecycle_policy": {Tok: makeResource(mainMod, "ContainerRepositoryLifecyclePolicy")},
			"yandex_dataproc_cluster":                      {Tok: makeResource(mainMod, "DataprocCluster")},
			"yandex_datatransfer_endpoint":                 {Tok: makeResource(mainMod, "DatatransferEndpoint")},
			"yandex_datatransfer_transfer":                 {Tok: makeResource(mainMod, "DatatransferTransfer")},
			"yandex_dns_recordset":                         {Tok: makeResource(mainMod, "DnsRecordSet")},
			"yandex_dns_zone":                              {Tok: makeResource(mainMod, "DnsZone")},
			"yandex_function":                              {Tok: makeResource(mainMod, "Function")},
			"yandex_function_trigger":                      {Tok: makeResource(mainMod, "FunctionTrigger")},
			"yandex_function_scaling_policy":               {Tok: makeResource(mainMod, "FunctionScalingPolicy")},
			// IAM Resources
			"yandex_iam_service_account":                           {Tok: makeResource(mainMod, "IamServiceAccount")},
			"yandex_iam_service_account_api_key":                   {Tok: makeResource(mainMod, "IamServiceAccountApiKey")},
			"yandex_iam_service_account_iam_policy":                {Tok: makeResource(mainMod, "IamServiceAccountIamPolicy")},
			"yandex_iam_service_account_key":                       {Tok: makeResource(mainMod, "IamServiceAccountKey")},
			"yandex_iam_service_account_static_access_key":         {Tok: makeResource(mainMod, "IamServiceAccountStaticAccessKey")},
			"yandex_iam_workload_identity_federated_credential":    {Tok: makeResource(mainMod, "IamWorkloadIdentityFederatedCredential")},
			"yandex_iam_workload_identity_oidc_federation":         {Tok: makeResource(mainMod, "IamWorkloadIdentityOidcFederation")},
			// IoT Resources
			"yandex_iot_core_broker":   {Tok: makeResource(mainMod, "IotCoreBroker")},
			"yandex_iot_core_device":   {Tok: makeResource(mainMod, "IotCoreDevice")},
			"yandex_iot_core_registry": {Tok: makeResource(mainMod, "IotCoreRegistry")},
			// KMS Resources
			"yandex_kms_asymmetric_encryption_key": {Tok: makeResource(mainMod, "KmsAsymmetricEncryptionKey")},
			"yandex_kms_asymmetric_signature_key":  {Tok: makeResource(mainMod, "KmsAsymmetricSignatureKey")},
			"yandex_kms_secret_ciphertext":         {Tok: makeResource(mainMod, "KmsSecretCiphertext")},
			"yandex_kms_symmetric_key":             {Tok: makeResource(mainMod, "KmsSymmetricKey")},
			// Kubernetes Resources
			"yandex_kubernetes_cluster":    {Tok: makeResource(mainMod, "KubernetesCluster")},
			"yandex_kubernetes_node_group": {Tok: makeResource(mainMod, "KubernetesNodeGroup")},
			// Load Balancer Resources
			"yandex_lb_network_load_balancer": {Tok: makeResource(mainMod, "LbNetworkLoadBalancer")},
			"yandex_lb_target_group":          {Tok: makeResource(mainMod, "LbTargetGroup")},
			"yandex_loadtesting_agent":        {Tok: makeResource(mainMod, "LoadtestingAgent")},
			// Lockbox Resources
			"yandex_lockbox_secret":                 {Tok: makeResource(mainMod, "LockboxSecret")},
			"yandex_lockbox_secret_version":         {Tok: makeResource(mainMod, "LockboxSecretVersion")},
			"yandex_lockbox_secret_version_hashed":  {Tok: makeResource(mainMod, "LockboxSecretVersionHashed")},
			"yandex_logging_group":                  {Tok: makeResource(mainMod, "LoggingGroup")},
			// Managed Database Resources
			"yandex_mdb_clickhouse_cluster":  {Tok: makeResource(mainMod, "MdbClickhouseCluster")},
			"yandex_mdb_greenplum_cluster":   {Tok: makeResource(mainMod, "MdbGreenplumCluster")},
			"yandex_mdb_kafka_cluster":       {Tok: makeResource(mainMod, "MdbKafkaCluster")},
			"yandex_mdb_kafka_connector":     {Tok: makeResource(mainMod, "MdbKafkaConnector")},
			"yandex_mdb_kafka_topic":         {Tok: makeResource(mainMod, "MdbKafkaTopic")},
			"yandex_mdb_kafka_user":          {Tok: makeResource(mainMod, "MdbKafkaUser")},
			"yandex_mdb_mongodb_cluster":     {Tok: makeResource(mainMod, "MdbMongodbCluster")},
			"yandex_mdb_mysql_cluster":       {Tok: makeResource(mainMod, "MdbMysqlCluster")},
			"yandex_mdb_mysql_database":      {Tok: makeResource(mainMod, "MdbMysqlDatabase")},
			"yandex_mdb_mysql_user":          {Tok: makeResource(mainMod, "MdbMysqlUser")},
			"yandex_mdb_postgresql_cluster":  {Tok: makeResource(mainMod, "MdbPostgresqlCluster")},
			"yandex_mdb_postgresql_database": {Tok: makeResource(mainMod, "MdbPostgresqlDatabase")},
			"yandex_mdb_postgresql_user":     {Tok: makeResource(mainMod, "MdbPostgresqlUser")},
			"yandex_mdb_redis_cluster":       {Tok: makeResource(mainMod, "MdbRedisCluster")},
			"yandex_mdb_sqlserver_cluster":   {Tok: makeResource(mainMod, "MdbSqlServerCluster")},
			"yandex_message_queue":           {Tok: makeResource(mainMod, "MessageQueue")},
			"yandex_monitoring_dashboard":    {Tok: makeResource(mainMod, "MonitoringDashboard")},
			// Organization Manager Resources
			"yandex_organizationmanager_group":                        {Tok: makeResource(mainMod, "OrganizationmanagerGroup")},
			"yandex_organizationmanager_group_mapping":                {Tok: makeResource(mainMod, "OrganizationmanagerGroupMapping")},
			"yandex_organizationmanager_group_mapping_item":           {Tok: makeResource(mainMod, "OrganizationmanagerGroupMappingItem")},
			"yandex_organizationmanager_group_membership":             {Tok: makeResource(mainMod, "OrganizationmanagerGroupMembership")},
			"yandex_organizationmanager_os_login_settings":            {Tok: makeResource(mainMod, "OrganizationmanagerOsLoginSettings")},
			"yandex_organizationmanager_saml_federation":              {Tok: makeResource(mainMod, "OrganizationmanagerSamlFederation")},
			"yandex_organizationmanager_saml_federation_user_account": {Tok: makeResource(mainMod, "OrganizationmanagerSamlFederationUserAccount")},
			"yandex_organizationmanager_user_ssh_key":                 {Tok: makeResource(mainMod, "OrganizationmanagerUserSshKey")},
			// Resource Manager
			"yandex_resourcemanager_folder_iam_policy": {Tok: makeResource(mainMod, "ResourcemanagerFolderIamPolicy")},
			// Serverless Resources
			"yandex_serverless_container":         {Tok: makeResource(mainMod, "ServerlessContainer")},
			"yandex_serverless_eventrouter_bus":   {Tok: makeResource(mainMod, "ServerlessEventrouterBus")},
			"yandex_serverless_eventrouter_connector": {Tok: makeResource(mainMod, "ServerlessEventrouterConnector")},
			"yandex_serverless_eventrouter_rule":      {Tok: makeResource(mainMod, "ServerlessEventrouterRule")},
			"yandex_smartcaptcha_captcha":             {Tok: makeResource(mainMod, "SmartcaptchaCaptcha")},
			// Storage Resources
			"yandex_storage_bucket": {Tok: makeResource(mainMod, "StorageBucket")},
			"yandex_storage_object": {Tok: makeResource(mainMod, "StorageObject")},
			// SWS Resources
			"yandex_sws_advanced_rate_limiter_profile": {Tok: makeResource(mainMod, "SwsAdvancedRateLimiterProfile")},
			"yandex_sws_security_profile":              {Tok: makeResource(mainMod, "SwsSecurityProfile")},
			"yandex_sws_waf_profile":                   {Tok: makeResource(mainMod, "SwsWafProfile")},
			// VPC Resources
			"yandex_vpc_address":                {Tok: makeResource(mainMod, "VpcAddress")},
			"yandex_vpc_default_security_group": {Tok: makeResource(mainMod, "VpcDefaultSecurityGroup")},
			"yandex_vpc_gateway":                {Tok: makeResource(mainMod, "VpcGateway")},
			"yandex_vpc_network":                {Tok: makeResource(mainMod, "VpcNetwork")},
			"yandex_vpc_private_endpoint":       {Tok: makeResource(mainMod, "VpcPrivateEndpoint")},
			"yandex_vpc_route_table":            {Tok: makeResource(mainMod, "VpcRouteTable")},
			"yandex_vpc_security_group":         {Tok: makeResource(mainMod, "VpcSecurityGroup")},
			"yandex_vpc_subnet":                 {Tok: makeResource(mainMod, "VpcSubnet")},
			// YDB Resources
			"yandex_ydb_database_dedicated":  {Tok: makeResource(mainMod, "YdbDatabaseDedicated")},
			"yandex_ydb_database_serverless": {Tok: makeResource(mainMod, "YdbDatabaseServerless")},
			"yandex_ydb_table":               {Tok: makeResource(mainMod, "YdbTable")},
			"yandex_ydb_table_changefeed":    {Tok: makeResource(mainMod, "YdbTableChangefeed")},
			"yandex_ydb_table_index":         {Tok: makeResource(mainMod, "YdbTableIndex")},
			"yandex_ydb_topic":               {Tok: makeResource(mainMod, "YdbTopic")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"yandex_alb_target_group": {
				Tok: makeDataSource(mainMod, "getAlbTargetGroup"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_alb_target_group.html.markdown",
				},
			},
			"yandex_alb_backend_group": {
				Tok: makeDataSource(mainMod, "getAlbBackendGroup"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_alb_backend_group.html.markdown",
				},
			},
			"yandex_alb_http_router": {
				Tok: makeDataSource(mainMod, "getAlbHttpRouter"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_alb_http_router.html.markdown",
				},
			},
			"yandex_alb_virtual_host": {
				Tok: makeDataSource(mainMod, "getAlbVirtualHost"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_alb_virtual_host.html.markdown",
				},
			},
			"yandex_alb_load_balancer": {Tok: makeDataSource(mainMod, "getAlbLoadBalancer")},
			"yandex_api_gateway":       {Tok: makeDataSource(mainMod, "getApiGateway")},
			"yandex_client_config": {
				Tok: makeDataSource(mainMod, "getClientConfig"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_client_config.html.markdown",
				},
			},
			"yandex_compute_disk": {
				Tok: makeDataSource(mainMod, "getComputeDisk"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_compute_disk.html.markdown",
				},
			},
			"yandex_compute_disk_placement_group": {
				Tok: makeDataSource(mainMod, "getComputeDiskPlacementGroup"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_compute_disk_placement_group.html.markdown",
				},
			},
			"yandex_compute_image": {
				Tok: makeDataSource(mainMod, "getComputeImage"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_compute_image.html.markdown",
				},
			},
			"yandex_compute_instance": {
				Tok: makeDataSource(mainMod, "getComputeInstance"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_compute_instance.html.markdown",
				},
			},
			"yandex_compute_instance_group": {
				Tok: makeDataSource(mainMod, "getComputeInstanceGroup"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_compute_instance_group.html.markdown",
				},
			},
			"yandex_compute_placement_group": {
				Tok: makeDataSource(mainMod, "getComputePlacementGroup"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_compute_placement_group.html.markdown",
				},
			},
			"yandex_compute_snapshot": {
				Tok: makeDataSource(mainMod, "getComputeSnapshot"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_compute_snapshot.html.markdown",
				},
			},
			"yandex_container_registry": {
				Tok: makeDataSource(mainMod, "getContainerRegistry"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_container_registry.html.markdown",
				},
			},
			"yandex_container_repository": {
				Tok: makeDataSource(mainMod, "getContainerRepository"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_container_repository.html.markdown",
				},
			},
			"yandex_dataproc_cluster": {
				Tok: makeDataSource(mainMod, "getDataprocCluster"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_dataproc_cluster.html.markdown",
				},
			},
			"yandex_dns_zone": {
				Tok: makeDataSource(mainMod, "getDnsZone"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_dns_zone.html.markdown",
				},
			},
			"yandex_function": {
				Tok: makeDataSource(mainMod, "getFunction"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_function.html.markdown",
				},
			},
			"yandex_function_trigger": {
				Tok: makeDataSource(mainMod, "getFunctionTrigger"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_function_trigger.html.markdown",
				},
			},
			"yandex_function_scaling_policy": {Tok: makeDataSource(mainMod, "getFunctionScalingPolicy")},
			"yandex_iam_policy": {
				Tok: makeDataSource(mainMod, "getIamPolicy"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_iam_policy.html.markdown",
				},
			},
			"yandex_iam_role": {
				Tok: makeDataSource(mainMod, "getIamRole"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_iam_role.html.markdown",
				},
			},
			"yandex_iam_service_account": {
				Tok: makeDataSource(mainMod, "getIamServiceAccount"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_iam_service_account.html.markdown",
				},
			},
			"yandex_iam_user": {
				Tok: makeDataSource(mainMod, "getIamUser"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_iam_user.html.markdown",
				},
			},
			"yandex_iot_core_device": {
				Tok: makeDataSource(mainMod, "getIotCoreDevice"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_iot_core_device.html.markdown",
				},
			},
			"yandex_iot_core_registry": {
				Tok: makeDataSource(mainMod, "getIotCoreRegistry"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_iot_core_registry.html.markdown",
				},
			},
			"yandex_kubernetes_cluster": {
				Tok: makeDataSource(mainMod, "getKubernetesCluster"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_kubernetes_cluster.html.markdown",
				},
			},
			"yandex_kubernetes_node_group": {
				Tok: makeDataSource(mainMod, "getKubernetesNodeGroup"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_kubernetes_node_group.html.markdown",
				},
			},
			"yandex_lb_network_load_balancer": {
				Tok: makeDataSource(mainMod, "getLbNetworkLoadBalancer"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_lb_network_load_balancer.html.markdown",
				},
			},
			"yandex_lb_target_group": {
				Tok: makeDataSource(mainMod, "getLbTargetGroup"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_lb_target_group.html.markdown",
				},
			},
			"yandex_logging_group": {Tok: makeDataSource(mainMod, "getLoggingGroup")},
			"yandex_mdb_clickhouse_cluster": {
				Tok: makeDataSource(mainMod, "getMdbClickhouseCluster"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_mdb_clickhouse_cluster.html.markdown",
				},
			},
			"yandex_mdb_greenplum_cluster": {Tok: makeDataSource(mainMod, "getMdbGreenplumCluster")},
			"yandex_mdb_kafka_cluster": {
				Tok: makeDataSource(mainMod, "getMdbKafkaCluster"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_mdb_kafka_cluster.html.markdown",
				},
			},
			"yandex_mdb_kafka_topic": {Tok: makeDataSource(mainMod, "getMdbKafkaTopic")},
			"yandex_mdb_mongodb_cluster": {
				Tok: makeDataSource(mainMod, "getMdbMongodbCluster"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_mdb_mongodb_cluster.html.markdown",
				},
			},
			"yandex_mdb_mysql_cluster": {
				Tok: makeDataSource(mainMod, "getMdbMysqlCluster"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_mdb_mysql_cluster.html.markdown",
				},
			},
			"yandex_mdb_postgresql_cluster": {
				Tok: makeDataSource(mainMod, "getMdbPostgresqlCluster"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_mdb_postgresql_cluster.html.markdown",
				},
			},
			"yandex_mdb_redis_cluster": {
				Tok: makeDataSource(mainMod, "getMdbRedisCluster"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_mdb_redis_cluster.html.markdown",
				},
			},
			"yandex_mdb_sqlserver_cluster": {
				Tok: makeDataSource(mainMod, "getMdbSqlserverCluster"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_mdb_sqlserver_cluster.html.markdown",
				},
			},
			"yandex_message_queue": {
				Tok: makeDataSource(mainMod, "getMessageQueue"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_message_queue.html.markdown",
				},
			},

			"yandex_resourcemanager_cloud": {
				Tok: makeDataSource(mainMod, "getResourcemanagerCloud"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_resourcemanager_cloud.html.markdown",
				},
			},
			"yandex_resourcemanager_folder": {
				Tok: makeDataSource(mainMod, "getResourcemanagerFolder"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_resourcemanager_folder.html.markdown",
				},
			},
			"yandex_vpc_address": {
				Tok: makeDataSource(mainMod, "getVpcAddress"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_vpc_address.html.markdown",
				},
			},
			"yandex_vpc_network": {
				Tok: makeDataSource(mainMod, "getVpcNetwork"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_vpc_network.html.markdown",
				},
			},
			"yandex_vpc_route_table": {
				Tok: makeDataSource(mainMod, "getVpcRouteTable"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_vpc_route_table.html.markdown",
				},
			},
			"yandex_vpc_security_group": {
				Tok: makeDataSource(mainMod, "getVpcSecurityGroup"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_vpc_security_group.html.markdown",
				},
			},
			"yandex_vpc_subnet": {
				Tok: makeDataSource(mainMod, "getVpcSubnet"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_vpc_subnet.html.markdown",
				},
			},
			"yandex_ydb_database_dedicated": {
				Tok: makeDataSource(mainMod, "getYdbDatabaseDedicated"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_ydb_database_dedicated.html.markdown",
				},
			},
			"yandex_ydb_database_serverless": {
				Tok: makeDataSource(mainMod, "getYdbDatabaseServerless"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_ydb_database_serverless.html.markdown",
				},
			},
			"yandex_cdn_origin_group": {
				Tok: makeDataSource(mainMod, "getCdnOriginGroup"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_cdn_origin_group.html.markdown",
				},
			},
			"yandex_cdn_resource": {
				Tok: makeDataSource(mainMod, "getCdnResource"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_cdn_resource.html.markdown",
				},
			},
			"yandex_serverless_container": {
				Tok: makeDataSource(mainMod, "getServerlessContainer"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_serverless_container.html.markdown",
				},
			},
			"yandex_organizationmanager_saml_federation": {
				Tok: makeDataSource(mainMod, "getOrganizationmanagerSamlFederation"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_organizationmanager_saml_federation.html.markdown",
				},
			},
			"yandex_organizationmanager_saml_federation_user_account": {
				Tok: makeDataSource(mainMod, "getOrganizationmanagerSamlFederationUserAccount"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_organizationmanager_saml_federation_user_account.html.markdown",
				},
			},
			// New Data Sources
			"yandex_audit_trails_trail":                        {Tok: makeDataSource(mainMod, "getAuditTrailsTrail")},
			"yandex_backup_policy":                             {Tok: makeDataSource(mainMod, "getBackupPolicy")},
			"yandex_cm_certificate":                            {Tok: makeDataSource(mainMod, "getCmCertificate")},
			"yandex_cm_certificate_content":                    {Tok: makeDataSource(mainMod, "getCmCertificateContent")},
			"yandex_compute_filesystem":                        {Tok: makeDataSource(mainMod, "getComputeFilesystem")},
			"yandex_compute_gpu_cluster":                       {Tok: makeDataSource(mainMod, "getComputeGpuCluster")},
			"yandex_compute_snapshot_schedule":                 {Tok: makeDataSource(mainMod, "getComputeSnapshotSchedule")},
			"yandex_container_registry_ip_permission":          {Tok: makeDataSource(mainMod, "getContainerRegistryIpPermission")},
			"yandex_container_repository_lifecycle_policy":     {Tok: makeDataSource(mainMod, "getContainerRepositoryLifecyclePolicy")},
			"yandex_iam_service_agent":                         {Tok: makeDataSource(mainMod, "getIamServiceAgent")},
			"yandex_iam_workload_identity_federated_credential": {Tok: makeDataSource(mainMod, "getIamWorkloadIdentityFederatedCredential")},
			"yandex_iam_workload_identity_oidc_federation":     {Tok: makeDataSource(mainMod, "getIamWorkloadIdentityOidcFederation")},
			"yandex_iot_core_broker":                           {Tok: makeDataSource(mainMod, "getIotCoreBroker")},
			"yandex_kms_asymmetric_encryption_key":             {Tok: makeDataSource(mainMod, "getKmsAsymmetricEncryptionKey")},
			"yandex_kms_asymmetric_signature_key":              {Tok: makeDataSource(mainMod, "getKmsAsymmetricSignatureKey")},
			"yandex_kms_symmetric_key":                         {Tok: makeDataSource(mainMod, "getKmsSymmetricKey")},
			"yandex_loadtesting_agent":                         {Tok: makeDataSource(mainMod, "getLoadtestingAgent")},
			"yandex_lockbox_secret":                            {Tok: makeDataSource(mainMod, "getLockboxSecret")},
			"yandex_lockbox_secret_version":                    {Tok: makeDataSource(mainMod, "getLockboxSecretVersion")},
			"yandex_mdb_kafka_connector":                       {Tok: makeDataSource(mainMod, "getMdbKafkaConnector")},
			"yandex_mdb_kafka_user":                            {Tok: makeDataSource(mainMod, "getMdbKafkaUser")},
			"yandex_mdb_mysql_database":                        {Tok: makeDataSource(mainMod, "getMdbMysqlDatabase")},
			"yandex_mdb_mysql_user":                            {Tok: makeDataSource(mainMod, "getMdbMysqlUser")},
			"yandex_mdb_postgresql_database":                   {Tok: makeDataSource(mainMod, "getMdbPostgresqlDatabase")},
			"yandex_mdb_postgresql_user":                       {Tok: makeDataSource(mainMod, "getMdbPostgresqlUser")},
			"yandex_monitoring_dashboard":                      {Tok: makeDataSource(mainMod, "getMonitoringDashboard")},
			"yandex_organizationmanager_group":                 {Tok: makeDataSource(mainMod, "getOrganizationmanagerGroup")},
			"yandex_organizationmanager_os_login_settings":     {Tok: makeDataSource(mainMod, "getOrganizationmanagerOsLoginSettings")},
			"yandex_organizationmanager_user_ssh_key":          {Tok: makeDataSource(mainMod, "getOrganizationmanagerUserSshKey")},
			"yandex_serverless_eventrouter_bus":                {Tok: makeDataSource(mainMod, "getServerlessEventrouterBus")},
			"yandex_serverless_eventrouter_connector":          {Tok: makeDataSource(mainMod, "getServerlessEventrouterConnector")},
			"yandex_serverless_eventrouter_rule":               {Tok: makeDataSource(mainMod, "getServerlessEventrouterRule")},
			"yandex_smartcaptcha_captcha":                      {Tok: makeDataSource(mainMod, "getSmartcaptchaCaptcha")},
			"yandex_sws_advanced_rate_limiter_profile":         {Tok: makeDataSource(mainMod, "getSwsAdvancedRateLimiterProfile")},
			"yandex_sws_security_profile":                      {Tok: makeDataSource(mainMod, "getSwsSecurityProfile")},
			"yandex_sws_waf_profile":                           {Tok: makeDataSource(mainMod, "getSwsWafProfile")},
			"yandex_sws_waf_rule_set_descriptor":               {Tok: makeDataSource(mainMod, "getSwsWafRuleSetDescriptor")},
			"yandex_vpc_gateway":                               {Tok: makeDataSource(mainMod, "getVpcGateway")},
			"yandex_vpc_private_endpoint":                      {Tok: makeDataSource(mainMod, "getVpcPrivateEndpoint")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@airoh/pulumi-yandex",
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.142.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0",
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/airoh-io/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
