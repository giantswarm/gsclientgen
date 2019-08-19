# V4AddClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | **string** | Name of the organization owning the cluster | 
**Name** | **string** | Cluster name | [optional] 
**ReleaseVersion** | **string** | The [release](https://docs.giantswarm.io/api/#tag/releases) version to use in the new cluster  | [optional] 
**AvailabilityZones** | **int32** | Number of availability zones a cluster should be spread across. The default is provided via the [info](#operation/getInfo) endpoint.  | [optional] 
**Scaling** | [**V4AddClusterRequestScaling**](V4AddClusterRequest_scaling.md) |  | [optional] 
**Workers** | [**[]V4NodeDefinition**](V4NodeDefinition.md) | Worker node definition on KVM and Azure. If present, the first item of the array is expected to contain the specification for all worker nodes. Not available on AWS.  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


