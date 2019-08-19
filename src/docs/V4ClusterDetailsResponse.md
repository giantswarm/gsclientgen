# V4ClusterDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique cluster identifier | [optional] 
**ApiEndpoint** | **string** | URI of the Kubernetes API endpoint | [optional] 
**CreateDate** | **string** | Date/time of cluster creation | [optional] 
**Owner** | **string** | Name of the organization owning the cluster | [optional] 
**Name** | **string** | Cluster name | [optional] 
**CredentialId** | **string** | ID of the credentials used to operate the cluster (only on AWS and Azure). See [Set credentials](#operation/addCredentials) for details.  | [optional] 
**ReleaseVersion** | **string** | The [release](https://docs.giantswarm.io/api/#tag/releases) version currently running this cluster.  | [optional] 
**Scaling** | [**V4ClusterDetailsResponseScaling**](V4ClusterDetailsResponse_scaling.md) |  | [optional] 
**AvailabilityZones** | **[]string** | List of availability zones a cluster is spread across. | [optional] 
**Workers** | [**[]V4NodeDefinition**](V4NodeDefinition.md) | Information about worker nodes in the cluster (deprecated) | [optional] 
**Kvm** | [**V4ClusterDetailsResponseKvm**](V4ClusterDetailsResponse_kvm.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


