# \GetDataAPI

All URIs are relative to *https://data.api.abs.gov.au*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetData**](GetDataAPI.md#GetData) | **Get** /rest/data/{dataflowIdentifier}/{dataKey} | Get data from a dataflow in XML, JSON or CSV.



## GetData

> string GetData(ctx, dataflowIdentifier, dataKey).StartPeriod(startPeriod).EndPeriod(endPeriod).Format(format).Detail(detail).DimensionAtObservation(dimensionAtObservation).Execute()

Get data from a dataflow in XML, JSON or CSV.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	dataflowIdentifier := "dataflowIdentifier_example" // string | The dataflow identifier in {agencyId},{dataflowId},{version} format (eg. \"ABS,CPI,1.1.0\").  A list of all available dataflows can be returned using the *GET /rest/{structureType}/{agencyId}* operation. agencyId and version are optional. If agencyId is not specified it will default to “all”. If version is not specified it will default to “latest”. 
	dataKey := "dataKey_example" // string | The key to query data returned. Use \"all\" if you would like to return all data for the dataset. To filter data returned provide a data key, containing one or more coded values for each dimension, separated by a dot (dimensions must be in the order they are defined in the data structure).  For example “1.115486.10.50.Q”  filters for a single Consumer Price Index series  * *Measure*: 1 - Index Number * *Index*: 115486 - Health * *Adjustment Type*: 10 - Original * *Region*: 50 - Weighted average of eight capital cities * *Frequency*: Q - Quarterly   Wildcarding is supported by omitting values for the dimension. Eg. data for all regions: “1.115486.10..Q”. The OR operator is supported using the + character. Eg. data for 2 series “1.131188+131189.10.50.Q”.  You can combine wildcarding and the OR operator. Eg. “1.131188+131189.10..Q\".    The maximum allowed length of this parameter is currently 260 characters. Requests that exceed this limit will return a 400 error. We intend to increase this limit soon.  The dataKey parameter is case sensitive. 
	startPeriod := "startPeriod_example" // string | The start period (used to filter on time). This is inclusive. The value can be in the following formats:   * year: yyyy * year-semester: yyyy-S1 - yyyy-S2 * year-quarter: yyyy-Q1 - yyyy-Q4   * year-month: yyyy-01 - yyyy-12  (optional)
	endPeriod := "endPeriod_example" // string | The end period (used to filter on time). This is inclusive. The value can be in the following formats: * year: yyyy * year-semester: yyyy-S1 - yyyy-S2   * year-quarter: yyyy-Q1 - yyyy-Q4   * year-month: yyyy-01 - yyyy-12  (optional)
	format := "format_example" // string | The format of data to return. Refer to the accept header if omitted.   * **csvwithlabels**: csv format with columns for dimension codes and labels. * **csvfile**: csv format with columns for dimension codes. * **jsondata**: sdmx json format equivalent to application/vnd.sdmx.data+json * **genericdata**: sdmx xml format equivalent to application/vnd.sdmx.genericdata+xml * **structurespecificdata**: sdmx xml format equivalent to application/vnd.sdmx.structurespecificdata+xml  (optional)
	detail := "detail_example" // string | The detail of data to return.  * **full**: The data - series and observations - and the attributes will be returned. This is the default. * **dataonly**: The attributes will be excluded from the returned message. * **serieskeysonly**: Only the series, but without the attributes and the observations, will be returned. This can be useful for performance reasons, to return the series that match a certain query, without returning the actual data. * **nodata**: The series, including the attributes, will be returned, but the observations will not be returned.  (optional)
	dimensionAtObservation := "dimensionAtObservation_example" // string | Define the way data should be organized in the returned message. Possible options are:  * **TIME_PERIOD**: This will return a timeseries view of the data. This is the default value. * **AllDimensions**: This will return a flat view of the data, with no groupings. * **The ID of any other dimension**: This will return a cross-sectional view of the data.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetDataAPI.GetData(context.Background(), dataflowIdentifier, dataKey).StartPeriod(startPeriod).EndPeriod(endPeriod).Format(format).Detail(detail).DimensionAtObservation(dimensionAtObservation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetDataAPI.GetData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetData`: string
	fmt.Fprintf(os.Stdout, "Response from `GetDataAPI.GetData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataflowIdentifier** | **string** | The dataflow identifier in {agencyId},{dataflowId},{version} format (eg. \&quot;ABS,CPI,1.1.0\&quot;).  A list of all available dataflows can be returned using the *GET /rest/{structureType}/{agencyId}* operation. agencyId and version are optional. If agencyId is not specified it will default to “all”. If version is not specified it will default to “latest”.  | 
**dataKey** | **string** | The key to query data returned. Use \&quot;all\&quot; if you would like to return all data for the dataset. To filter data returned provide a data key, containing one or more coded values for each dimension, separated by a dot (dimensions must be in the order they are defined in the data structure).  For example “1.115486.10.50.Q”  filters for a single Consumer Price Index series  * *Measure*: 1 - Index Number * *Index*: 115486 - Health * *Adjustment Type*: 10 - Original * *Region*: 50 - Weighted average of eight capital cities * *Frequency*: Q - Quarterly   Wildcarding is supported by omitting values for the dimension. Eg. data for all regions: “1.115486.10..Q”. The OR operator is supported using the + character. Eg. data for 2 series “1.131188+131189.10.50.Q”.  You can combine wildcarding and the OR operator. Eg. “1.131188+131189.10..Q\&quot;.    The maximum allowed length of this parameter is currently 260 characters. Requests that exceed this limit will return a 400 error. We intend to increase this limit soon.  The dataKey parameter is case sensitive.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startPeriod** | **string** | The start period (used to filter on time). This is inclusive. The value can be in the following formats:   * year: yyyy * year-semester: yyyy-S1 - yyyy-S2 * year-quarter: yyyy-Q1 - yyyy-Q4   * year-month: yyyy-01 - yyyy-12  | 
 **endPeriod** | **string** | The end period (used to filter on time). This is inclusive. The value can be in the following formats: * year: yyyy * year-semester: yyyy-S1 - yyyy-S2   * year-quarter: yyyy-Q1 - yyyy-Q4   * year-month: yyyy-01 - yyyy-12  | 
 **format** | **string** | The format of data to return. Refer to the accept header if omitted.   * **csvwithlabels**: csv format with columns for dimension codes and labels. * **csvfile**: csv format with columns for dimension codes. * **jsondata**: sdmx json format equivalent to application/vnd.sdmx.data+json * **genericdata**: sdmx xml format equivalent to application/vnd.sdmx.genericdata+xml * **structurespecificdata**: sdmx xml format equivalent to application/vnd.sdmx.structurespecificdata+xml  | 
 **detail** | **string** | The detail of data to return.  * **full**: The data - series and observations - and the attributes will be returned. This is the default. * **dataonly**: The attributes will be excluded from the returned message. * **serieskeysonly**: Only the series, but without the attributes and the observations, will be returned. This can be useful for performance reasons, to return the series that match a certain query, without returning the actual data. * **nodata**: The series, including the attributes, will be returned, but the observations will not be returned.  | 
 **dimensionAtObservation** | **string** | Define the way data should be organized in the returned message. Possible options are:  * **TIME_PERIOD**: This will return a timeseries view of the data. This is the default value. * **AllDimensions**: This will return a flat view of the data, with no groupings. * **The ID of any other dimension**: This will return a cross-sectional view of the data.  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.sdmx.data+json, application/xml, application/vnd.sdmx.structurespecificdata+xml, text/csv, application/vnd.sdmx.data+csv, application/vnd.sdmx.data+csv;labels=both, application/vnd.sdmx.data+csv;file=true;labels=both

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

