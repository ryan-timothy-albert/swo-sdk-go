// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/solarwinds/swo-sdk-go/solarwinds/models/components"
)

type GetWebsiteMonitorRequest struct {
	EntityID string `pathParam:"style=simple,explode=false,name=entityId"`
}

func (o *GetWebsiteMonitorRequest) GetEntityID() string {
	if o == nil {
		return ""
	}
	return o.EntityID
}

// AvailabilityCheckSettings -   Use this field to configure availability tests for the website.
//
//	You are required to configure at least availability monitoring or real user monitoring to be able to create website.
type AvailabilityCheckSettings struct {
	//   Use this field to configure whether availability tests should check for presence or absence of a particular string on a page.
	//   If the operator is DOES_NOT_CONTAIN and the value is found on the page, the availability test will fail.
	//   Likewise, if the operator is CONTAINS and the value is not found on the page, the availability test will fail.
	//   If omitted or set to null, the string checking functionality will be disabled.
	CheckForString *components.CheckForString `json:"checkForString,omitempty"`
	// Configure how often availability tests should be performed. Provide a number of seconds that is divisible by 60 and no greater than 14400 (4 hours).
	TestIntervalInSeconds int `json:"testIntervalInSeconds"`
	// Configure which protocols need availability tests to be performed. At least one protocol must be provided.
	Protocols []components.WebsiteProtocol `json:"protocols"`
	// Configure cloud platforms of the synthetic availability test probes. If omitted or set to null, no particular cloud platform will be enforced.
	PlatformOptions *components.ProbePlatformOptions `json:"platformOptions,omitempty"`
	//   Configure locations of the synthetic availability test probes.
	//   Acceptable values depend on the selected type and actual values of existing probes.
	TestFrom components.TestFrom `json:"testFrom"`
	//   Configure monitoring of SSL/TLS certificates validity. This option is relevant for HTTPS protocol only.
	//   If omitted or set to null, SSL monitoring will be disabled and its previous configuration discarded.
	Ssl *components.SslMonitoring `json:"ssl,omitempty"`
	//   Configure custom request headers to be sent with each availability test. It is possible to provide multiple headers with the same name.
	//   If omitted, set to null or set to an empty array, no custom headers will be sent.
	CustomHeaders []components.CustomHeaders `json:"customHeaders,omitempty"`
	//   Allow insecure SSL renegotiation which introduces a security risk in the communication process.
	//   Checking this option could lead to exposing credentials to unauthorized entities and the possibility of unauthorized access, interception, or manipulation of sensitive data, compromising the integrity and security of the communication channel.
	//   Available only with HTTPS check.
	//   If omitted or set to null, insecure SSL renegotiation won't be allowed.
	AllowInsecureRenegotiation *bool `json:"allowInsecureRenegotiation,omitempty"`
	//   Configure data that will be sent as POST request body by the synthetic probe.
	//   If omitted or set to null/empty string, the probe will send the usual GET requests.
	PostData *string `json:"postData,omitempty"`
}

func (o *AvailabilityCheckSettings) GetCheckForString() *components.CheckForString {
	if o == nil {
		return nil
	}
	return o.CheckForString
}

func (o *AvailabilityCheckSettings) GetTestIntervalInSeconds() int {
	if o == nil {
		return 0
	}
	return o.TestIntervalInSeconds
}

func (o *AvailabilityCheckSettings) GetProtocols() []components.WebsiteProtocol {
	if o == nil {
		return []components.WebsiteProtocol{}
	}
	return o.Protocols
}

func (o *AvailabilityCheckSettings) GetPlatformOptions() *components.ProbePlatformOptions {
	if o == nil {
		return nil
	}
	return o.PlatformOptions
}

func (o *AvailabilityCheckSettings) GetTestFrom() components.TestFrom {
	if o == nil {
		return components.TestFrom{}
	}
	return o.TestFrom
}

func (o *AvailabilityCheckSettings) GetSsl() *components.SslMonitoring {
	if o == nil {
		return nil
	}
	return o.Ssl
}

func (o *AvailabilityCheckSettings) GetCustomHeaders() []components.CustomHeaders {
	if o == nil {
		return nil
	}
	return o.CustomHeaders
}

func (o *AvailabilityCheckSettings) GetAllowInsecureRenegotiation() *bool {
	if o == nil {
		return nil
	}
	return o.AllowInsecureRenegotiation
}

func (o *AvailabilityCheckSettings) GetPostData() *string {
	if o == nil {
		return nil
	}
	return o.PostData
}

// Rum -   Use this field to configure real user monitoring (RUM) for the website.
//
//	You are required to configure at least availability monitoring or real user monitoring to be able to create website.
type Rum struct {
	ApdexTimeInSeconds *int    `json:"apdexTimeInSeconds,omitempty"`
	Snippet            *string `json:"snippet,omitempty"`
	Spa                bool    `json:"spa"`
}

func (o *Rum) GetApdexTimeInSeconds() *int {
	if o == nil {
		return nil
	}
	return o.ApdexTimeInSeconds
}

func (o *Rum) GetSnippet() *string {
	if o == nil {
		return nil
	}
	return o.Snippet
}

func (o *Rum) GetSpa() bool {
	if o == nil {
		return false
	}
	return o.Spa
}

// GetWebsiteMonitorResponseBody - The request has succeeded.
type GetWebsiteMonitorResponseBody struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	//   Name of the website, which must be unique within the organization.
	//   The website must also not contain any control characters, any white space other than space (U+0020), or any consecutive, leading or trailing spaces.
	Name string `json:"name"`
	// URL of the website. Must be a valid URL with no leading or trailing white space. Must not contain invalid port number (>65535).
	URL string `json:"url"`
	//   Use this field to configure availability tests for the website.
	//   You are required to configure at least availability monitoring or real user monitoring to be able to create website.
	AvailabilityCheckSettings *AvailabilityCheckSettings `json:"availabilityCheckSettings,omitempty"`
	// Entity tags. Tag is a key-value pair, where there may be only single tag value for the same key.
	Tags []components.Tag `json:"tags,omitempty"`
	//   Use this field to configure real user monitoring (RUM) for the website.
	//   You are required to configure at least availability monitoring or real user monitoring to be able to create website.
	Rum *Rum `json:"rum,omitempty"`
	// Timestamp for when the next on-demand check could be executed. If at '0', it means you can execute it anytime.
	NextOnDemandAvailabilityTime *int `json:"nextOnDemandAvailabilityTime,omitempty"`
}

func (o *GetWebsiteMonitorResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetWebsiteMonitorResponseBody) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *GetWebsiteMonitorResponseBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *GetWebsiteMonitorResponseBody) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *GetWebsiteMonitorResponseBody) GetAvailabilityCheckSettings() *AvailabilityCheckSettings {
	if o == nil {
		return nil
	}
	return o.AvailabilityCheckSettings
}

func (o *GetWebsiteMonitorResponseBody) GetTags() []components.Tag {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *GetWebsiteMonitorResponseBody) GetRum() *Rum {
	if o == nil {
		return nil
	}
	return o.Rum
}

func (o *GetWebsiteMonitorResponseBody) GetNextOnDemandAvailabilityTime() *int {
	if o == nil {
		return nil
	}
	return o.NextOnDemandAvailabilityTime
}

type GetWebsiteMonitorResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The request has succeeded.
	Object *GetWebsiteMonitorResponseBody
}

func (o *GetWebsiteMonitorResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetWebsiteMonitorResponse) GetObject() *GetWebsiteMonitorResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
