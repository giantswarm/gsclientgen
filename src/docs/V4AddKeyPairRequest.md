# V4AddKeyPairRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Free text information about the key pair | 
**TtlHours** | **int32** | Expiration time (from creation) in hours | [optional] 
**CnPrefix** | **string** | The common name prefix of the certificate subject. This only allows characters that are usable in domain names (&#x60;a-z&#x60;, &#x60;0-9&#x60;, and &#x60;.-&#x60;, where &#x60;.-&#x60; must not occur at either the start or the end). | [optional] 
**CertificateOrganizations** | **string** | This will set the certificate subject&#39;s &#x60;organization&#x60; fields. Use a comma seperated list of values.  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


