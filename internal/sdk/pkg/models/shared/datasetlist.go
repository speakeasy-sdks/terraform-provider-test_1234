// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DataSetListApis struct {
	// A URL to the API console for each API
	APIDocumentationURL *string `json:"apiDocumentationUrl,omitempty"`
	// To be used as a dataset parameter value
	APIKey *string `json:"apiKey,omitempty"`
	// The URL describing the dataset's fields
	APIURL *string `json:"apiUrl,omitempty"`
	// To be used as a version parameter value
	APIVersionNumber *string `json:"apiVersionNumber,omitempty"`
}

type DataSetList struct {
	Apis  []DataSetListApis `json:"apis,omitempty"`
	Total *int64            `json:"total,omitempty"`
}