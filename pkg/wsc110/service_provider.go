package wsc110

type ServiceProvider struct {
	ProviderName string `xml:"ows:ProviderName" yaml:"providerName"`
	ProviderSite struct {
		Type string `xml:"xlink:type,attr" yaml:"type"`
		Href string `xml:"xlink:href,attr" yaml:"href"`
	} `xml:"ows:ProviderSite" yaml:"providerSite"`
	ServiceContact struct {
		IndividualName string `xml:"ows:IndividualName,omitempty" yaml:"individualName"`
		PositionName   string `xml:"ows:PositionName,omitempty" yaml:"positionName"`
		ContactInfo    struct {
			Phone struct {
				Voice     string `xml:"ows:Voice" yaml:"voice"`
				Facsimile string `xml:"ows:Facsimile" yaml:"facsimile"`
			} `xml:"ows:Phone" yaml:"phone"`
			Address struct {
				DeliveryPoint         string `xml:"ows:DeliveryPoint" yaml:"deliveryPoint"`
				City                  string `xml:"ows:City" yaml:"city"`
				AdministrativeArea    string `xml:"ows:AdministrativeArea" yaml:"administrativeArea"`
				PostalCode            string `xml:"ows:PostalCode" yaml:"postalCode"`
				Country               string `xml:"ows:Country" yaml:"country"`
				ElectronicMailAddress string `xml:"ows:ElectronicMailAddress" yaml:"electronicMailAddress"`
			} `xml:"ows:Address" yaml:"address"`
			OnlineResource *struct {
				Type string `xml:"xlink:type,attr,omitempty" yaml:"type"`
				Href string `xml:"xlink:href,attr,omitempty" yaml:"href"`
			} `xml:"ows:OnlineResource,omitempty" yaml:"onlineResource"`
			HoursOfService      string `xml:"ows:HoursOfService,omitempty" yaml:"hoursOfService"`
			ContactInstructions string `xml:"ows:ContactInstructions,omitempty" yaml:"contactInstructions"`
		} `xml:"ows:ContactInfo" yaml:"contactInfo"`
		Role string `xml:"ows:Role,omitempty" yaml:"role"`
	} `xml:"ows:ServiceContact" yaml:"serviceContact"`
}
