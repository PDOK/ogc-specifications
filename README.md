# ogc-specifications

The package ogc-specifications is a implementation of the OGC Webservice Specifications as defined by the [OGC](https://www.ogc.org/).
This package has support for the following OGC Webservice Specifications:

| Spec | Request or Operation |
| --- | --- |
| WMS | GetCapabilities |
| WMS | GetMap |
| WMS | GetFeatureInfo |
| WFS | GetCapabilities |
| WFS | DescribeFeatureType |
| WFS | GetFeature |

It will provide the user with structs that can be used with in a developers application, so one doesn't needs to create/build those complex structs for 'every' application that has more then 'simple' interaction with a OGC Webservice.

## Installation

```go
go get github.com/pdok/ogc-specifications
```

```import
import "github.com/pdok/ogc-specifications"
```

## How to Contribute

Make a pull request...

## License

Distributed under MIT License, please see license file within the code for more details.
