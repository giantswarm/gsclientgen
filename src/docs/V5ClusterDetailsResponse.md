# V5ClusterDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique cluster identifier | [optional] 
**ApiEndpoint** | **string** | URI of the Kubernetes API endpoint | [optional] 
**CreateDate** | **string** | Date/time of cluster creation | [optional] 
**Owner** | **string** | Name of the organization owning the cluster  | [optional] 
**Name** | **string** | Cluster name | [optional] 
**CredentialId** | **string** | ID of the credentials used to operate the cluster. See [Set credentials](#operation/addCredentials) for details.  | [optional] 
**ReleaseVersion** | **string** | The [release](https://docs.giantswarm.io/api/#tag/releases) version used by the cluster  | [optional] 
**Master** | [**V5ClusterDetailsResponseMaster**](V5ClusterDetailsResponse_master.md) |  | [optional] 
**Conditions** | [**[]V5ClusterDetailsResponseConditions**](V5ClusterDetailsResponse_conditions.md) | List of conditions the cluster has gone through | [optional] 
**Versions** | [**[]V5ClusterDetailsResponseVersions**](V5ClusterDetailsResponse_versions.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


