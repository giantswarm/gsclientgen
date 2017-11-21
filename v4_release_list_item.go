/*
 * Giant Swarm legacy API
 *
 * Caution: This is an incomplete description of some legacy API functions.
 *
 * OpenAPI spec version: legacy
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package gsclientgen

type V4ReleaseListItem struct {

	// The semantic version number
	Version string `json:"version"`

	// Date and time of the release creation
	Timestamp string `json:"timestamp"`

	// If true, the version is available for new clusters and cluster upgrades. Older versions become unavailable and thus have the value `false` here.
	Active bool `json:"active,omitempty"`

	// Structured list of changes in this release, in comparison to the previous version, with respect to the contained components.
	Changelog []V4ReleaseChangelogItem `json:"changelog"`

	// List of components and their version contained in the release
	Components []V4ReleaseComponent `json:"components"`
}
