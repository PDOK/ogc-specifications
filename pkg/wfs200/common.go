package wfs200

import (
	"net/url"

	"github.com/pdok/ogc-specifications/pkg/common"
	"github.com/pdok/ogc-specifications/pkg/wsc110"
)

//
const (
	getcapabilities     = `GetCapabilities`
	getfeature          = `GetFeature`
	describefeaturetype = `DescribeFeatureType`
	getpropertyvalue    = `GetPropertyValue`

	/* Unused/support transactional requesttypes
	lockfeature = `LockFeature`
	getfeaturewithlock = `GetFeatureWithLock`
	*/

	// TODO StoredQuery
	// liststoredqueries     = `ListStoredQueries`
	// describestoredqueries = `DescribeStoredQueries`
	/* Unused/support tranactional storedquery requesttypes
	createstoredquery = `CreateStoredQuery`
	dropstoredquery = `DropStoredQuery`
	*/

	Service = `WFS`
	Version = `2.0.0`
)

// WFS 2.0.0 Tokens
const (
	SERVICE = `SERVICE`
	REQUEST = `REQUEST`
	VERSION = `VERSION`

	OUTPUTFORMAT = `OUTPUTFORMAT`
)

// BaseRequestKVP struct
type BaseRequestKVP struct {
	Version string `yaml:"version,omitempty"`
	Request string `yaml:"request,omitempty"`
}

// BaseRequest based on Table 5 WFS2.0.0 spec
// Note: not usable for GetCapabilities request regarding deviation of Optional/Mandatory parameters SERVICE and VERSION
type BaseRequest struct {
	Service string              `xml:"service,attr" yaml:"service,omitempty"`
	Version string              `xml:"version,attr" yaml:"version"`
	Attr    common.XMLAttribute `xml:",attr"`
}

// ParseKVP builds a BaseRequest Struct based on the given parameters
func (b *BaseRequest) ParseKVP(query url.Values) wsc110.Exceptions {
	if len(query[SERVICE]) > 0 {
		// Service is optional, because it's implicit for a GetFeature/DescribeFeatureType request
		b.Service = query[SERVICE][0]
	}
	if len(query[VERSION]) > 0 {
		b.Version = query[VERSION][0]
	} else {
		// Version is mandatory
		return wsc110.Exceptions{wsc110.MissingParameterValue(VERSION)}
	}
	return nil
}
