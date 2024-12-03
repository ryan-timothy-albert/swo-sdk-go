// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/solarwinds/swo-sdk-go/solarwinds/models/components"
)

type DeleteWebsiteMonitorRequest struct {
	EntityID string `pathParam:"style=simple,explode=false,name=entityId"`
}

func (o *DeleteWebsiteMonitorRequest) GetEntityID() string {
	if o == nil {
		return ""
	}
	return o.EntityID
}

// DeleteWebsiteMonitorResponseBody - The request has succeeded.
type DeleteWebsiteMonitorResponseBody struct {
	ID string `json:"id"`
}

func (o *DeleteWebsiteMonitorResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteWebsiteMonitorResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The request has succeeded.
	Object *DeleteWebsiteMonitorResponseBody
}

func (o *DeleteWebsiteMonitorResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteWebsiteMonitorResponse) GetObject() *DeleteWebsiteMonitorResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
