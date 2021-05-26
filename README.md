# ogc-specifications

![GitHub license](https://img.shields.io/github/license/PDOK/ogc-specifications)
[![GitHub release](https://img.shields.io/github/release/PDOK/ogc-specifications.svg)](https://github.com/PDOK/ogc-specifications/releases)
[![Go Report Card](https://goreportcard.com/badge/PDOK/ogc-specifications)](https://goreportcard.com/report/PDOK/ogc-specifications)

The package ogc-specifications is a implementation of the OGC Web Service Specifications as defined by the [OGC](https://www.ogc.org/).
This package has support for the following OGC Web Service Specifications:

| Spec | Version | Operation | Request | Response |
| --- | --- | --- | --- | --- |
| WMS | 1.3.0 | GetCapabilities | :heavy_check_mark:  | :grey_exclamation: |
| WMS | 1.3.0 | GetMap | :heavy_check_mark: | |
| WMS | 1.3.0 | GetFeatureInfo | :heavy_check_mark: | |
| WFS | 2.0.0 | GetCapabilities | :heavy_check_mark: | :grey_exclamation: |
| WFS | 2.0.0 | DescribeFeatureType | :heavy_check_mark: | |
| WFS | 2.0.0 | GetFeature | :heavy_check_mark: | |
| WMTS | 1.0.0 | GetCapabilities | :heavy_check_mark: | :grey_exclamation: |
| WCS | 2.0.1 | GetCapabilities | :heavy_check_mark: | :grey_exclamation: |

It will provide the user with OperationRequest, OperationResponse and Capabilities structs for the different OGC specifications that can be used with in a developers application, so one doesn't needs to create/build those complex structs for 'every' application that has more then 'simple' interaction with a OGC Web Service. It will allow the developer to parse XML documents and query strings like they are defined in the OGC specification an build go structs with it and it will generate XML documents and KVP query strings based on those structs.

## Notice

This is still a 'work-in-progress' with the following major to do's:

- [ ] WFS StoredQuery support
- [ ] WMTS GetTile support
- [ ] WCS GetCoverage support
- [ ] OGC response support for metadata calls like GetCapabilities and DescribeFeatureType
- [ ] Sufficent validation support
- [ ] cleanup YAML parser
- [ ] WMS Time & Elevation parameters

## Installation

```go
go get github.com/pdok/ogc-specifications
```

```import
import ows "github.com/pdok/ogc-specifications/pkg/ows"
```

## Test

```go
go test ./... -covermode=atomic
```

And for benchmarks:

```go
go test -bench=. ./...
```

And generate CPU & MEM profiles:

```go
go test -cpuprofile cpu.prof -memprofile mem.prof -bench=.  ./...
```

For visualization use [pprof](https://github.com/google/pprof)

## Usage

- [Simple BBOX](./examples/simple-bbox/main.go)

## How to Contribute

Make a pull request...

## License

Distributed under MIT License, please see license file within the code for more details.
