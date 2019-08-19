# V5AddClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | **string** | Name of the organization owning the cluster  | 
**Name** | **string** | Cluster name | [optional] 
**ReleaseVersion** | **string** | The [release](https://docs.giantswarm.io/api/#tag/releases) version to use in the new cluster. If not given, the latest release will be used.  | [optional] 
**Master** | [**V5AddClusterRequestMaster**](V5AddClusterRequest_master.md) |  | [optional] 
**Nodepools** | [**[]V5AddNodePoolRequest**](V5AddNodePoolRequest.md) | Specification of node pools to be created immediately with the cluster.  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


