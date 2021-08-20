package wsc110

type ServiceProvider struct {
	ProviderName string `xml:"ows:ProviderName" yaml:"providername"`
	ProviderSite struct {
		Type string `xml:"xlink:type,attr" yaml:"type"`
		Href string `xml:"xlink:href,attr" yaml:"href"`
	} `xml:"ows:ProviderSite" yaml:"providersite"`
	ServiceContact struct {
		IndividualName string `xml:"ows:IndividualName,omitempty" yaml:"individualname"`
		PositionName   string `xml:"ows:PositionName,omitempty" yaml:"positionname"`
		ContactInfo    struct {
			Phone struct {
				Voice     string `xml:"ows:Voice" yaml:"voice"`
				Facsimile string `xml:"ows:Facsimile" yaml:"facsimile"`
			} `xml:"ows:Phone" yaml:"phone"`
			Address struct {
				DeliveryPoint         string `xml:"ows:DeliveryPoint" yaml:"deliverypoint"`
				City                  string `xml:"ows:City" yaml:"city"`
				AdministrativeArea    string `xml:"ows:AdministrativeArea" yaml:"administrativearea"`
				PostalCode            string `xml:"ows:PostalCode" yaml:"postalcode"`
				Country               string `xml:"ows:Country" yaml:"country"`
				ElectronicMailAddress string `xml:"ows:ElectronicMailAddress" yaml:"electronicmailaddress"`
			} `xml:"ows:Address" yaml:"address"`
			OnlineResource *struct {
				Type string `xml:"xlink:type,attr,omitempty" yaml:"type"`
				Href string `xml:"xlink:href,attr,omitempty" yaml:"href"`
			} `xml:"ows:OnlineResource,omitempty" yaml:"onlineresource"`
			HoursOfService      string `xml:"ows:HoursOfService,omitempty" yaml:"hoursofservice"`
			ContactInstructions string `xml:"ows:ContactInstructions,omitempty" yaml:"contactinstructions"`
		} `xml:"ows:ContactInfo" yaml:"contactinfo"`
		Role string `xml:"ows:Role,omitempty" yaml:"role"`
	} `xml:"ows:ServiceContact" yaml:"servicecontact"`
}
