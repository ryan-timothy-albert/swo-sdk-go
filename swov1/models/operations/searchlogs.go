// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/solarwinds/swo-sdk-go/swov1/internal/utils"
	"github.com/solarwinds/swo-sdk-go/swov1/models/components"
	"time"
)

type SearchLogsRequest struct {
	// A search query string
	Filter *string `queryParam:"style=form,explode=false,name=filter"`
	// Filter logs by a specific group name
	Group *string `queryParam:"style=form,explode=false,name=group"`
	// Filter logs by a specific entity id
	EntityID *string `queryParam:"style=form,explode=false,name=entityId"`
	// Timestamp in ISO 8601 format in UTC timezone: yyyy-MM-ddTHH:mm:ssZ
	StartTime *time.Time `queryParam:"style=form,explode=false,name=startTime"`
	// Timestamp in ISO 8601 format in UTC timezone: yyyy-MM-ddTHH:mm:ssZ
	EndTime *time.Time `queryParam:"style=form,explode=false,name=endTime"`
	// Search direction: backward, forward, or tail. Backward sorts logs from oldest to newest, forward sorts logs from newest to oldest, and tail sorts from oldest to newest.
	Direction *string `default:"backward" queryParam:"style=form,explode=false,name=direction"`
	// Number of items in a response page. Default varies by API.
	PageSize *int `queryParam:"style=form,explode=false,name=pageSize"`
	// Token for the requested page
	SkipToken *string `queryParam:"style=form,explode=false,name=skipToken"`
}

func (s SearchLogsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SearchLogsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SearchLogsRequest) GetFilter() *string {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *SearchLogsRequest) GetGroup() *string {
	if o == nil {
		return nil
	}
	return o.Group
}

func (o *SearchLogsRequest) GetEntityID() *string {
	if o == nil {
		return nil
	}
	return o.EntityID
}

func (o *SearchLogsRequest) GetStartTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartTime
}

func (o *SearchLogsRequest) GetEndTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndTime
}

func (o *SearchLogsRequest) GetDirection() *string {
	if o == nil {
		return nil
	}
	return o.Direction
}

func (o *SearchLogsRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *SearchLogsRequest) GetSkipToken() *string {
	if o == nil {
		return nil
	}
	return o.SkipToken
}

// SearchLogsResponseBody - The request has succeeded.
type SearchLogsResponseBody struct {
	Logs     []components.LogsEvent    `json:"logs"`
	PageInfo components.CommonPageInfo `json:"pageInfo"`
}

func (o *SearchLogsResponseBody) GetLogs() []components.LogsEvent {
	if o == nil {
		return []components.LogsEvent{}
	}
	return o.Logs
}

func (o *SearchLogsResponseBody) GetPageInfo() components.CommonPageInfo {
	if o == nil {
		return components.CommonPageInfo{}
	}
	return o.PageInfo
}

type SearchLogsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The request has succeeded.
	Object *SearchLogsResponseBody

	Next func() (*SearchLogsResponse, error)
}

func (o *SearchLogsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *SearchLogsResponse) GetObject() *SearchLogsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
