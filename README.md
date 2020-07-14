![GitHub license](https://img.shields.io/github/license/PDOK/ogc-specifications)
![GitHub release](https://img.shields.io/github/release/PDOK/ogc-specifications.svg)
[![Go Report Card](https://goreportcard.com/badge/PDOK/ogc-specifications)](https://goreportcard.com/report/PDOK/ogc-specifications) 

# ogc-specifications

The package ogc-specifications is a implementation of the OGC Webservice Specifications as defined by the [OGC](https://www.ogc.org/).
This package has support for the following OGC Webservice Specifications:

| Spec | Operation | Request | Reponse |
| --- | --- | --- | --- |
| WMS | GetCapabilities | :heavy_check_mark:  | :grey_exclamation: |
| WMS | GetMap | :heavy_check_mark: | |
| WMS | GetFeatureInfo | :heavy_check_mark: | |
| WFS | GetCapabilities | :heavy_check_mark: | :grey_exclamation: |
| WFS | DescribeFeatureType | :heavy_check_mark: | |
| WFS | GetFeature | :heavy_check_mark: | |
| WMTS | GetCapabilities | :heavy_check_mark: | :grey_exclamation: |
| WCS | GetCapabilities | :heavy_check_mark: | :grey_exclamation: |

It will provide the user with structs that can be used with in a developers application, so one doesn't needs to create/build those complex structs for 'every' application that has more then 'simple' interaction with a OGC Webservice. It will allow the developer to parse XML documents and query strings like they are defined in the OGC specification an build go structs with it and it will generate XML documents and query strings based on those structs.

## Notice

This is still a 'work in progres' with the following major todo's:

- [ ] Validation support
- [ ] YAML parser
- [ ] WMTS support
- [ ] WCS support
- [ ] OGC response support (at least for the metadata calls like DescribeFeatureType)
- [ ] WFS StoredQuery
- [ ] WMS Time & Elevation parameters

## Installation

```go
go get github.com/pdok/ogc-specifications
```

```import
import "github.com/pdok/ogc-specifications"
```

## Test

```go
go test ./... -covermode=atomic
```

And for benchmarks:

```go
go test -bench=. ./...
```

## Usage

- [Simple BBOX](./examples/simple-bbox/main.go)

## How to Contribute

Make a pull request...

## License

Distributed under MIT License, please see license file within the code for more details.
