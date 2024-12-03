// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/solarwinds/swo-sdk-go/solarwinds/models/components"
)

// ListEntityTypesResponseBody - List of entity types
type ListEntityTypesResponseBody struct {
	Types []string `json:"types"`
}

func (o *ListEntityTypesResponseBody) GetTypes() []string {
	if o == nil {
		return []string{}
	}
	return o.Types
}

type ListEntityTypesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// List of entity types
	Object *ListEntityTypesResponseBody
}

func (o *ListEntityTypesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListEntityTypesResponse) GetObject() *ListEntityTypesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
