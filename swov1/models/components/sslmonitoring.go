// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type SslMonitoring struct {
	//   Whether SSL monitoring is enabled for the website.
	//   If set to false, SSL monitoring will be disabled, but the other settings will be remembered in case you re-enable it later.
	//   If omitted, the previous setting will stay in effect. If there is no previous setting, the value will default to false.
	Enabled *bool `json:"enabled,omitempty"`
	// Number of days before the expiration date an SSL certificate will be considered 'expiring.'
	DaysPriorToExpiration *int `json:"daysPriorToExpiration,omitempty"`
	//   Use this option to limit the certificate expiration check to only the first certificate in the chain (normally the host certificate).
	//   This way you will not be warned about impending expiration of intermediate or root Certification Authority certificates in the chain.
	//   This option does not affect any other certificate validity checks besides expiration.
	//   If omitted, the previous setting will stay in effect. If there is no previous setting, the value will default to false.
	IgnoreIntermediateCertificates *bool `json:"ignoreIntermediateCertificates,omitempty"`
}

func (o *SslMonitoring) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *SslMonitoring) GetDaysPriorToExpiration() *int {
	if o == nil {
		return nil
	}
	return o.DaysPriorToExpiration
}

func (o *SslMonitoring) GetIgnoreIntermediateCertificates() *bool {
	if o == nil {
		return nil
	}
	return o.IgnoreIntermediateCertificates
}
