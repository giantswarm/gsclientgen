# V5GetNodePoolResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Node pool identifier. Unique within a tenant cluster. | [optional] 
**Name** | **string** | Node pool name | [optional] 
**AvailabilityZones** | **[]string** | Names of the availability zones used by the nodes of this pool.  | [optional] 
**Scaling** | [**V5GetNodePoolResponseScaling**](V5GetNodePoolResponse_scaling.md) |  | [optional] 
**Subnet** | **string** | IP address block used by the node pool | [optional] 
**NodeSpec** | [**V5GetNodePoolResponseNodeSpec**](V5GetNodePoolResponseNodeSpec.md) |  | [optional] 
**Status** | [**V5GetNodePoolResponseStatus**](V5GetNodePoolResponse_status.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


