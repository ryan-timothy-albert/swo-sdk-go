// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/solarwinds/swo-sdk-go/solarwinds/models/components"
)

type CreateCompositeMetricResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The request has succeeded and a new resource has been created as a result.
	CompositeMetric *components.CompositeMetric
}

func (o *CreateCompositeMetricResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateCompositeMetricResponse) GetCompositeMetric() *components.CompositeMetric {
	if o == nil {
		return nil
	}
	return o.CompositeMetric
}
