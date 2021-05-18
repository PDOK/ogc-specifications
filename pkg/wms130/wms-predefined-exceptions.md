# WMS Exception codes

Taken from OGC WMS 1.3.0 Document: [#06-042](http://portal.opengeospatial.org/files/?artifact_id=14416)

## Table E.1 - Service exception codes

| Exception code | Meaning |
| ---| --- |
| InvalidFormat |Request contains a Format not offered by the server. |
| InvalidCRS | Request contains a CRS not offered by the server for one or more of the Layers in the request. |
| LayerNotDefined | GetMap request is for a Layer not offered by the server, or GetFeatureInfo request is for a Layer not shown on the map. |
| StyleNotDefined | Request is for a Layer in a Style not offered by the server. |
| LayerNotQueryable | GetFeatureInfo request is applied to a Layer which is not declared queryable. |
| InvalidPoint | GetFeatureInfo request contains invalid I or J value. |
| CurrentUpdateSequence | Value of (optional) UpdateSequence parameter in GetCapabilities request is equal to current value of service metadata update sequence number. |
| InvalidUpdateSequence | Value of (optional) UpdateSequence parameter in GetCapabilities request is greater than current value of service metadata update sequence number. |
| MissingDimensionValue | Request does not include a sample dimension value, and the server did not declare a default value for that dimension. |
| InvalidDimensionValue | Request contains an invalid sample dimension value. OperationNotSupported Request is for an optional operation that is not supported by the server. |
