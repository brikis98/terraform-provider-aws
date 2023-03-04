// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []func(context.Context) (datasource.DataSourceWithConfigure, error) {
	return []func(context.Context) (datasource.DataSourceWithConfigure, error){
		newDataSourceSecurityGroupRule,
		newDataSourceSecurityGroupRules,
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []func(context.Context) (resource.ResourceWithConfigure, error) {
	return []func(context.Context) (resource.ResourceWithConfigure, error){
		newResourceSecurityGroupEgressRule,
		newResourceSecurityGroupIngressRule,
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_ami":                                        DataSourceAMI,
		"aws_ami_ids":                                    DataSourceAMIIDs,
		"aws_availability_zone":                          DataSourceAvailabilityZone,
		"aws_availability_zones":                         DataSourceAvailabilityZones,
		"aws_customer_gateway":                           DataSourceCustomerGateway,
		"aws_ebs_default_kms_key":                        DataSourceEBSDefaultKMSKey,
		"aws_ebs_encryption_by_default":                  DataSourceEBSEncryptionByDefault,
		"aws_ebs_snapshot":                               DataSourceEBSSnapshot,
		"aws_ebs_snapshot_ids":                           DataSourceEBSSnapshotIDs,
		"aws_ebs_volume":                                 DataSourceEBSVolume,
		"aws_ebs_volumes":                                DataSourceEBSVolumes,
		"aws_ec2_client_vpn_endpoint":                    DataSourceClientVPNEndpoint,
		"aws_ec2_coip_pool":                              DataSourceCoIPPool,
		"aws_ec2_coip_pools":                             DataSourceCoIPPools,
		"aws_ec2_host":                                   DataSourceHost,
		"aws_ec2_instance_type":                          DataSourceInstanceType,
		"aws_ec2_instance_type_offering":                 DataSourceInstanceTypeOffering,
		"aws_ec2_instance_type_offerings":                DataSourceInstanceTypeOfferings,
		"aws_ec2_instance_types":                         DataSourceInstanceTypes,
		"aws_ec2_local_gateway":                          DataSourceLocalGateway,
		"aws_ec2_local_gateway_route_table":              DataSourceLocalGatewayRouteTable,
		"aws_ec2_local_gateway_route_tables":             DataSourceLocalGatewayRouteTables,
		"aws_ec2_local_gateway_virtual_interface":        DataSourceLocalGatewayVirtualInterface,
		"aws_ec2_local_gateway_virtual_interface_group":  DataSourceLocalGatewayVirtualInterfaceGroup,
		"aws_ec2_local_gateway_virtual_interface_groups": DataSourceLocalGatewayVirtualInterfaceGroups,
		"aws_ec2_local_gateways":                         DataSourceLocalGateways,
		"aws_ec2_managed_prefix_list":                    DataSourceManagedPrefixList,
		"aws_ec2_managed_prefix_lists":                   DataSourceManagedPrefixLists,
		"aws_ec2_network_insights_analysis":              DataSourceNetworkInsightsAnalysis,
		"aws_ec2_network_insights_path":                  DataSourceNetworkInsightsPath,
		"aws_ec2_serial_console_access":                  DataSourceSerialConsoleAccess,
		"aws_ec2_spot_price":                             DataSourceSpotPrice,
		"aws_ec2_transit_gateway":                        DataSourceTransitGateway,
		"aws_ec2_transit_gateway_attachment":             DataSourceTransitGatewayAttachment,
		"aws_ec2_transit_gateway_connect":                DataSourceTransitGatewayConnect,
		"aws_ec2_transit_gateway_connect_peer":           DataSourceTransitGatewayConnectPeer,
		"aws_ec2_transit_gateway_dx_gateway_attachment":  DataSourceTransitGatewayDxGatewayAttachment,
		"aws_ec2_transit_gateway_multicast_domain":       DataSourceTransitGatewayMulticastDomain,
		"aws_ec2_transit_gateway_peering_attachment":     DataSourceTransitGatewayPeeringAttachment,
		"aws_ec2_transit_gateway_route_table":            DataSourceTransitGatewayRouteTable,
		"aws_ec2_transit_gateway_route_tables":           DataSourceTransitGatewayRouteTables,
		"aws_ec2_transit_gateway_vpc_attachment":         DataSourceTransitGatewayVPCAttachment,
		"aws_ec2_transit_gateway_vpc_attachments":        DataSourceTransitGatewayVPCAttachments,
		"aws_ec2_transit_gateway_vpn_attachment":         DataSourceTransitGatewayVPNAttachment,
		"aws_eip":                                        DataSourceEIP,
		"aws_eips":                                       DataSourceEIPs,
		"aws_instance":                                   DataSourceInstance,
		"aws_instances":                                  DataSourceInstances,
		"aws_internet_gateway":                           DataSourceInternetGateway,
		"aws_key_pair":                                   DataSourceKeyPair,
		"aws_launch_template":                            DataSourceLaunchTemplate,
		"aws_nat_gateway":                                DataSourceNATGateway,
		"aws_nat_gateways":                               DataSourceNATGateways,
		"aws_network_acls":                               DataSourceNetworkACLs,
		"aws_network_interface":                          DataSourceNetworkInterface,
		"aws_network_interfaces":                         DataSourceNetworkInterfaces,
		"aws_prefix_list":                                DataSourcePrefixList,
		"aws_route":                                      DataSourceRoute,
		"aws_route_table":                                DataSourceRouteTable,
		"aws_route_tables":                               DataSourceRouteTables,
		"aws_security_group":                             DataSourceSecurityGroup,
		"aws_security_groups":                            DataSourceSecurityGroups,
		"aws_subnet":                                     DataSourceSubnet,
		"aws_subnet_ids":                                 DataSourceSubnetIDs,
		"aws_subnets":                                    DataSourceSubnets,
		"aws_vpc":                                        DataSourceVPC,
		"aws_vpc_dhcp_options":                           DataSourceVPCDHCPOptions,
		"aws_vpc_endpoint":                               DataSourceVPCEndpoint,
		"aws_vpc_endpoint_service":                       DataSourceVPCEndpointService,
		"aws_vpc_ipam_pool":                              DataSourceIPAMPool,
		"aws_vpc_ipam_pool_cidrs":                        DataSourceIPAMPoolCIDRs,
		"aws_vpc_ipam_pools":                             DataSourceIPAMPools,
		"aws_vpc_ipam_preview_next_cidr":                 DataSourceIPAMPreviewNextCIDR,
		"aws_vpc_peering_connection":                     DataSourceVPCPeeringConnection,
		"aws_vpc_peering_connections":                    DataSourceVPCPeeringConnections,
		"aws_vpcs":                                       DataSourceVPCs,
		"aws_vpn_gateway":                                DataSourceVPNGateway,
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_ami":                                              ResourceAMI,
		"aws_ami_copy":                                         ResourceAMICopy,
		"aws_ami_from_instance":                                ResourceAMIFromInstance,
		"aws_ami_launch_permission":                            ResourceAMILaunchPermission,
		"aws_customer_gateway":                                 ResourceCustomerGateway,
		"aws_default_network_acl":                              ResourceDefaultNetworkACL,
		"aws_default_route_table":                              ResourceDefaultRouteTable,
		"aws_default_security_group":                           ResourceDefaultSecurityGroup,
		"aws_default_subnet":                                   ResourceDefaultSubnet,
		"aws_default_vpc":                                      ResourceDefaultVPC,
		"aws_default_vpc_dhcp_options":                         ResourceDefaultVPCDHCPOptions,
		"aws_ebs_default_kms_key":                              ResourceEBSDefaultKMSKey,
		"aws_ebs_encryption_by_default":                        ResourceEBSEncryptionByDefault,
		"aws_ebs_snapshot":                                     ResourceEBSSnapshot,
		"aws_ebs_snapshot_copy":                                ResourceEBSSnapshotCopy,
		"aws_ebs_snapshot_import":                              ResourceEBSSnapshotImport,
		"aws_ebs_volume":                                       ResourceEBSVolume,
		"aws_ec2_availability_zone_group":                      ResourceAvailabilityZoneGroup,
		"aws_ec2_capacity_reservation":                         ResourceCapacityReservation,
		"aws_ec2_carrier_gateway":                              ResourceCarrierGateway,
		"aws_ec2_client_vpn_authorization_rule":                ResourceClientVPNAuthorizationRule,
		"aws_ec2_client_vpn_endpoint":                          ResourceClientVPNEndpoint,
		"aws_ec2_client_vpn_network_association":               ResourceClientVPNNetworkAssociation,
		"aws_ec2_client_vpn_route":                             ResourceClientVPNRoute,
		"aws_ec2_fleet":                                        ResourceFleet,
		"aws_ec2_host":                                         ResourceHost,
		"aws_ec2_instance_state":                               ResourceInstanceState,
		"aws_ec2_local_gateway_route":                          ResourceLocalGatewayRoute,
		"aws_ec2_local_gateway_route_table_vpc_association":    ResourceLocalGatewayRouteTableVPCAssociation,
		"aws_ec2_managed_prefix_list":                          ResourceManagedPrefixList,
		"aws_ec2_managed_prefix_list_entry":                    ResourceManagedPrefixListEntry,
		"aws_ec2_network_insights_analysis":                    ResourceNetworkInsightsAnalysis,
		"aws_ec2_network_insights_path":                        ResourceNetworkInsightsPath,
		"aws_ec2_serial_console_access":                        ResourceSerialConsoleAccess,
		"aws_ec2_subnet_cidr_reservation":                      ResourceSubnetCIDRReservation,
		"aws_ec2_tag":                                          ResourceTag,
		"aws_ec2_traffic_mirror_filter":                        ResourceTrafficMirrorFilter,
		"aws_ec2_traffic_mirror_filter_rule":                   ResourceTrafficMirrorFilterRule,
		"aws_ec2_traffic_mirror_session":                       ResourceTrafficMirrorSession,
		"aws_ec2_traffic_mirror_target":                        ResourceTrafficMirrorTarget,
		"aws_ec2_transit_gateway":                              ResourceTransitGateway,
		"aws_ec2_transit_gateway_connect":                      ResourceTransitGatewayConnect,
		"aws_ec2_transit_gateway_connect_peer":                 ResourceTransitGatewayConnectPeer,
		"aws_ec2_transit_gateway_multicast_domain":             ResourceTransitGatewayMulticastDomain,
		"aws_ec2_transit_gateway_multicast_domain_association": ResourceTransitGatewayMulticastDomainAssociation,
		"aws_ec2_transit_gateway_multicast_group_member":       ResourceTransitGatewayMulticastGroupMember,
		"aws_ec2_transit_gateway_multicast_group_source":       ResourceTransitGatewayMulticastGroupSource,
		"aws_ec2_transit_gateway_peering_attachment":           ResourceTransitGatewayPeeringAttachment,
		"aws_ec2_transit_gateway_peering_attachment_accepter":  ResourceTransitGatewayPeeringAttachmentAccepter,
		"aws_ec2_transit_gateway_policy_table":                 ResourceTransitGatewayPolicyTable,
		"aws_ec2_transit_gateway_policy_table_association":     ResourceTransitGatewayPolicyTableAssociation,
		"aws_ec2_transit_gateway_prefix_list_reference":        ResourceTransitGatewayPrefixListReference,
		"aws_ec2_transit_gateway_route":                        ResourceTransitGatewayRoute,
		"aws_ec2_transit_gateway_route_table":                  ResourceTransitGatewayRouteTable,
		"aws_ec2_transit_gateway_route_table_association":      ResourceTransitGatewayRouteTableAssociation,
		"aws_ec2_transit_gateway_route_table_propagation":      ResourceTransitGatewayRouteTablePropagation,
		"aws_ec2_transit_gateway_vpc_attachment":               ResourceTransitGatewayVPCAttachment,
		"aws_ec2_transit_gateway_vpc_attachment_accepter":      ResourceTransitGatewayVPCAttachmentAccepter,
		"aws_egress_only_internet_gateway":                     ResourceEgressOnlyInternetGateway,
		"aws_eip":                                              ResourceEIP,
		"aws_eip_association":                                  ResourceEIPAssociation,
		"aws_flow_log":                                         ResourceFlowLog,
		"aws_instance":                                         ResourceInstance,
		"aws_internet_gateway":                                 ResourceInternetGateway,
		"aws_internet_gateway_attachment":                      ResourceInternetGatewayAttachment,
		"aws_key_pair":                                         ResourceKeyPair,
		"aws_launch_template":                                  ResourceLaunchTemplate,
		"aws_main_route_table_association":                     ResourceMainRouteTableAssociation,
		"aws_nat_gateway":                                      ResourceNATGateway,
		"aws_network_acl":                                      ResourceNetworkACL,
		"aws_network_acl_association":                          ResourceNetworkACLAssociation,
		"aws_network_acl_rule":                                 ResourceNetworkACLRule,
		"aws_network_interface":                                ResourceNetworkInterface,
		"aws_network_interface_attachment":                     ResourceNetworkInterfaceAttachment,
		"aws_network_interface_sg_attachment":                  ResourceNetworkInterfaceSGAttachment,
		"aws_placement_group":                                  ResourcePlacementGroup,
		"aws_route":                                            ResourceRoute,
		"aws_route_table":                                      ResourceRouteTable,
		"aws_route_table_association":                          ResourceRouteTableAssociation,
		"aws_security_group":                                   ResourceSecurityGroup,
		"aws_security_group_rule":                              ResourceSecurityGroupRule,
		"aws_snapshot_create_volume_permission":                ResourceSnapshotCreateVolumePermission,
		"aws_spot_datafeed_subscription":                       ResourceSpotDataFeedSubscription,
		"aws_spot_fleet_request":                               ResourceSpotFleetRequest,
		"aws_spot_instance_request":                            ResourceSpotInstanceRequest,
		"aws_subnet":                                           ResourceSubnet,
		"aws_volume_attachment":                                ResourceVolumeAttachment,
		"aws_vpc":                                              ResourceVPC,
		"aws_vpc_dhcp_options":                                 ResourceVPCDHCPOptions,
		"aws_vpc_dhcp_options_association":                     ResourceVPCDHCPOptionsAssociation,
		"aws_vpc_endpoint":                                     ResourceVPCEndpoint,
		"aws_vpc_endpoint_connection_accepter":                 ResourceVPCEndpointConnectionAccepter,
		"aws_vpc_endpoint_connection_notification":             ResourceVPCEndpointConnectionNotification,
		"aws_vpc_endpoint_policy":                              ResourceVPCEndpointPolicy,
		"aws_vpc_endpoint_route_table_association":             ResourceVPCEndpointRouteTableAssociation,
		"aws_vpc_endpoint_security_group_association":          ResourceVPCEndpointSecurityGroupAssociation,
		"aws_vpc_endpoint_service":                             ResourceVPCEndpointService,
		"aws_vpc_endpoint_service_allowed_principal":           ResourceVPCEndpointServiceAllowedPrincipal,
		"aws_vpc_endpoint_subnet_association":                  ResourceVPCEndpointSubnetAssociation,
		"aws_vpc_ipam":                                         ResourceIPAM,
		"aws_vpc_ipam_organization_admin_account":              ResourceIPAMOrganizationAdminAccount,
		"aws_vpc_ipam_pool":                                    ResourceIPAMPool,
		"aws_vpc_ipam_pool_cidr":                               ResourceIPAMPoolCIDR,
		"aws_vpc_ipam_pool_cidr_allocation":                    ResourceIPAMPoolCIDRAllocation,
		"aws_vpc_ipam_preview_next_cidr":                       ResourceIPAMPreviewNextCIDR,
		"aws_vpc_ipam_resource_discovery":                      ResourceIPAMResourceDiscovery,
		"aws_vpc_ipam_resource_discovery_association":          ResourceIPAMResourceDiscoveryAssociation,
		"aws_vpc_ipam_scope":                                   ResourceIPAMScope,
		"aws_vpc_ipv4_cidr_block_association":                  ResourceVPCIPv4CIDRBlockAssociation,
		"aws_vpc_ipv6_cidr_block_association":                  ResourceVPCIPv6CIDRBlockAssociation,
		"aws_vpc_network_performance_metric_subscription":      ResourceNetworkPerformanceMetricSubscription,
		"aws_vpc_peering_connection":                           ResourceVPCPeeringConnection,
		"aws_vpc_peering_connection_accepter":                  ResourceVPCPeeringConnectionAccepter,
		"aws_vpc_peering_connection_options":                   ResourceVPCPeeringConnectionOptions,
		"aws_vpn_connection":                                   ResourceVPNConnection,
		"aws_vpn_connection_route":                             ResourceVPNConnectionRoute,
		"aws_vpn_gateway":                                      ResourceVPNGateway,
		"aws_vpn_gateway_attachment":                           ResourceVPNGatewayAttachment,
		"aws_vpn_gateway_route_propagation":                    ResourceVPNGatewayRoutePropagation,
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.EC2
}

var ServicePackage = &servicePackage{}
