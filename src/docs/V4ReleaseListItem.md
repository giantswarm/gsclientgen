# V4ReleaseListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | The semantic version number | 
**Timestamp** | **string** | Date and time of the release creation | 
**Active** | **bool** | If true, the version is available for new clusters and cluster upgrades. Older versions become unavailable and thus have the value &#x60;false&#x60; here.  | [optional] 
**Changelog** | [**[]V4ReleaseListItemChangelog**](V4ReleaseListItem_changelog.md) | Structured list of changes in this release, in comparison to the previous version, with respect to the contained components.  | 
**Components** | [**[]V4ReleaseListItemComponents**](V4ReleaseListItem_components.md) | List of components and their version contained in the release  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


