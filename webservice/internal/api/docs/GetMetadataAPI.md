# \GetMetadataAPI

All URIs are relative to *https://data.api.abs.gov.au*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLatestStructure**](GetMetadataAPI.md#GetLatestStructure) | **Get** /rest/{structureType}/{agencyId}/{structureId} | Get the latest structure of a specific type, for a given agency and structure id.
[**GetStructures**](GetMetadataAPI.md#GetStructures) | **Get** /rest/{structureType}/{agencyId}/{structureId}/{structureVersion} | Get a specific version of a structure of a specific type, for a given agency and structure id.
[**GetStructuresByAgencyId**](GetMetadataAPI.md#GetStructuresByAgencyId) | **Get** /rest/{structureType}/{agencyId} | Get all structures of a specific type.



## GetLatestStructure

> string GetLatestStructure(ctx, structureType, agencyId, structureId).References(references).Detail(detail).Execute()

Get the latest structure of a specific type, for a given agency and structure id.

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
	structureType := "structureType_example" // string | The type of structure to retrieve.
	agencyId := "agencyId_example" // string | The id of the agency maintaining the structures. Eg. \"ABS\". (default to "ABS")
	structureId := "structureId_example" // string | The structure's id. \"all\" will return all artefacts of the selected structure type.
	references := "references_example" // string | Instruct the web service to return (or not) the artefacts referenced by the artefact(s) you are querying. Eg. the codelists and concepts used by the data structure you are querying. You can also retrieve the artefacts that use the artefact you are querying, eg. the dataflows that use the data structure definition queried.  * **none**: No references will be returned. This is the default. * **parents**: The artefacts that use the artefact matching the query (for example, the dataflows that use the data structure definition matching the query) will be returned. * **parentsandsiblings**: The artefacts that use the artefact matching the query, as well as the artefacts referenced by these artefacts will be returned. * **children**: The artefacts referenced by the matching artefact will be returned (for example, the concept schemes and code lists used in a DSD). * **descendants**: References of references, up to any level, will also be returned. * **all**: The combination of parentsandsiblings and descendants. * In addition, a specific structure type may also be used (e.g. codelist, dataflow, etc.).  (optional)
	detail := "detail_example" // string | Specify the desired amount of detail to be returned. For example, it is possible to instruct the web service to return only basic information about the resource, this is known in SDMX as a stub.  * **allstubs**: All artefacts will be returned as stubs. * **referencestubs**: The referenced artefacts will be returned as stubs. * **referencepartial**: The referenced item schemes should only include items used by the artefact to be returned. For example, a concept scheme would only contain the concepts used in a DSD, and its isPartial flag would be set to true. As another example, if a dataflow is constrained, then the codelists returned should only contain the subset of codes allowed by that constraint. * **allcompletestubs**: All artefacts should be returned as complete stubs, containing identification information, the artefacts' name, description, annotations and isFinal information. * **referencecompletestubs**: The referenced artefacts should be returned as complete stubs, containing identification information, the artefacts' name, description, annotations and isFinal information. * **full**: All available information for all artefacts will be returned. This is the default.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetMetadataAPI.GetLatestStructure(context.Background(), structureType, agencyId, structureId).References(references).Detail(detail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetMetadataAPI.GetLatestStructure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLatestStructure`: string
	fmt.Fprintf(os.Stdout, "Response from `GetMetadataAPI.GetLatestStructure`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**structureType** | **string** | The type of structure to retrieve. | 
**agencyId** | **string** | The id of the agency maintaining the structures. Eg. \&quot;ABS\&quot;. | [default to &quot;ABS&quot;]
**structureId** | **string** | The structure&#39;s id. \&quot;all\&quot; will return all artefacts of the selected structure type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestStructureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **references** | **string** | Instruct the web service to return (or not) the artefacts referenced by the artefact(s) you are querying. Eg. the codelists and concepts used by the data structure you are querying. You can also retrieve the artefacts that use the artefact you are querying, eg. the dataflows that use the data structure definition queried.  * **none**: No references will be returned. This is the default. * **parents**: The artefacts that use the artefact matching the query (for example, the dataflows that use the data structure definition matching the query) will be returned. * **parentsandsiblings**: The artefacts that use the artefact matching the query, as well as the artefacts referenced by these artefacts will be returned. * **children**: The artefacts referenced by the matching artefact will be returned (for example, the concept schemes and code lists used in a DSD). * **descendants**: References of references, up to any level, will also be returned. * **all**: The combination of parentsandsiblings and descendants. * In addition, a specific structure type may also be used (e.g. codelist, dataflow, etc.).  | 
 **detail** | **string** | Specify the desired amount of detail to be returned. For example, it is possible to instruct the web service to return only basic information about the resource, this is known in SDMX as a stub.  * **allstubs**: All artefacts will be returned as stubs. * **referencestubs**: The referenced artefacts will be returned as stubs. * **referencepartial**: The referenced item schemes should only include items used by the artefact to be returned. For example, a concept scheme would only contain the concepts used in a DSD, and its isPartial flag would be set to true. As another example, if a dataflow is constrained, then the codelists returned should only contain the subset of codes allowed by that constraint. * **allcompletestubs**: All artefacts should be returned as complete stubs, containing identification information, the artefacts&#39; name, description, annotations and isFinal information. * **referencecompletestubs**: The referenced artefacts should be returned as complete stubs, containing identification information, the artefacts&#39; name, description, annotations and isFinal information. * **full**: All available information for all artefacts will be returned. This is the default.  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml, application/vnd.sdmx.structure+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStructures

> string GetStructures(ctx, structureType, agencyId, structureId, structureVersion).References(references).Detail(detail).Execute()

Get a specific version of a structure of a specific type, for a given agency and structure id.

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
	structureType := "structureType_example" // string | The type of structure to retrieve.
	agencyId := "agencyId_example" // string | The id of the agency maintaining the structures. Eg. \"ABS\". (default to "ABS")
	structureId := "structureId_example" // string | The structure's id. \"all\" will return all artefacts of the selected structure type.
	structureVersion := "structureVersion_example" // string | The version of the structure to retrieve. Three numbers separated by points, eg. \"1.0.0\".
	references := "references_example" // string | Instruct the web service to return (or not) the artefacts referenced by the artefact(s) you are querying. Eg. the codelists and concepts used by the data structure you are querying. You can also retrieve the artefacts that use the artefact you are querying, eg. the dataflows that use the data structure definition queried.  * **none**: No references will be returned. This is the default. * **parents**: The artefacts that use the artefact matching the query (for example, the dataflows that use the data structure definition matching the query) will be returned. * **parentsandsiblings**: The artefacts that use the artefact matching the query, as well as the artefacts referenced by these artefacts will be returned. * **children**: The artefacts referenced by the matching artefact will be returned (for example, the concept schemes and code lists used in a DSD). * **descendants**: References of references, up to any level, will also be returned. * **all**: The combination of parentsandsiblings and descendants. * In addition, a specific structure type may also be used (e.g. codelist, dataflow, etc.).  (optional)
	detail := "detail_example" // string | Specify the desired amount of detail to be returned. For example, it is possible to instruct the web service to return only basic information about the resource, this is known in SDMX as a stub.  * **allstubs**: All artefacts will be returned as stubs. * **referencestubs**: The referenced artefacts will be returned as stubs. * **referencepartial**: The referenced item schemes should only include items used by the artefact to be returned. For example, a concept scheme would only contain the concepts used in a DSD, and its isPartial flag would be set to true. As another example, if a dataflow is constrained, then the codelists returned should only contain the subset of codes allowed by that constraint. * **allcompletestubs**: All artefacts should be returned as complete stubs, containing identification information, the artefacts' name, description, annotations and isFinal information. * **referencecompletestubs**: The referenced artefacts should be returned as complete stubs, containing identification information, the artefacts' name, description, annotations and isFinal information. * **full**: All available information for all artefacts will be returned. This is the default.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetMetadataAPI.GetStructures(context.Background(), structureType, agencyId, structureId, structureVersion).References(references).Detail(detail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetMetadataAPI.GetStructures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStructures`: string
	fmt.Fprintf(os.Stdout, "Response from `GetMetadataAPI.GetStructures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**structureType** | **string** | The type of structure to retrieve. | 
**agencyId** | **string** | The id of the agency maintaining the structures. Eg. \&quot;ABS\&quot;. | [default to &quot;ABS&quot;]
**structureId** | **string** | The structure&#39;s id. \&quot;all\&quot; will return all artefacts of the selected structure type. | 
**structureVersion** | **string** | The version of the structure to retrieve. Three numbers separated by points, eg. \&quot;1.0.0\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStructuresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **references** | **string** | Instruct the web service to return (or not) the artefacts referenced by the artefact(s) you are querying. Eg. the codelists and concepts used by the data structure you are querying. You can also retrieve the artefacts that use the artefact you are querying, eg. the dataflows that use the data structure definition queried.  * **none**: No references will be returned. This is the default. * **parents**: The artefacts that use the artefact matching the query (for example, the dataflows that use the data structure definition matching the query) will be returned. * **parentsandsiblings**: The artefacts that use the artefact matching the query, as well as the artefacts referenced by these artefacts will be returned. * **children**: The artefacts referenced by the matching artefact will be returned (for example, the concept schemes and code lists used in a DSD). * **descendants**: References of references, up to any level, will also be returned. * **all**: The combination of parentsandsiblings and descendants. * In addition, a specific structure type may also be used (e.g. codelist, dataflow, etc.).  | 
 **detail** | **string** | Specify the desired amount of detail to be returned. For example, it is possible to instruct the web service to return only basic information about the resource, this is known in SDMX as a stub.  * **allstubs**: All artefacts will be returned as stubs. * **referencestubs**: The referenced artefacts will be returned as stubs. * **referencepartial**: The referenced item schemes should only include items used by the artefact to be returned. For example, a concept scheme would only contain the concepts used in a DSD, and its isPartial flag would be set to true. As another example, if a dataflow is constrained, then the codelists returned should only contain the subset of codes allowed by that constraint. * **allcompletestubs**: All artefacts should be returned as complete stubs, containing identification information, the artefacts&#39; name, description, annotations and isFinal information. * **referencecompletestubs**: The referenced artefacts should be returned as complete stubs, containing identification information, the artefacts&#39; name, description, annotations and isFinal information. * **full**: All available information for all artefacts will be returned. This is the default.  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml, application/vnd.sdmx.structure+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStructuresByAgencyId

> string GetStructuresByAgencyId(ctx, structureType, agencyId).Detail(detail).Execute()

Get all structures of a specific type.

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
	structureType := "structureType_example" // string | The type of structure to retrieve.
	agencyId := "agencyId_example" // string | The id of the agency maintaining the structures. Eg. \"ABS\". (default to "ABS")
	detail := "detail_example" // string | Specify the desired amount of detail to be returned. For example, it is possible to instruct the web service to return only basic information about the resource, this is known in SDMX as a stub.  * **allstubs**: All artefacts will be returned as stubs. * **referencestubs**: The referenced artefacts will be returned as stubs. * **referencepartial**: The referenced item schemes should only include items used by the artefact to be returned. For example, a concept scheme would only contain the concepts used in a DSD, and its isPartial flag would be set to true. As another example, if a dataflow is constrained, then the codelists returned should only contain the subset of codes allowed by that constraint. * **allcompletestubs**: All artefacts should be returned as complete stubs, containing identification information, the artefacts' name, description, annotations and isFinal information. * **referencecompletestubs**: The referenced artefacts should be returned as complete stubs, containing identification information, the artefacts' name, description, annotations and isFinal information. * **full**: All available information for all artefacts will be returned. This is the default.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetMetadataAPI.GetStructuresByAgencyId(context.Background(), structureType, agencyId).Detail(detail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetMetadataAPI.GetStructuresByAgencyId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStructuresByAgencyId`: string
	fmt.Fprintf(os.Stdout, "Response from `GetMetadataAPI.GetStructuresByAgencyId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**structureType** | **string** | The type of structure to retrieve. | 
**agencyId** | **string** | The id of the agency maintaining the structures. Eg. \&quot;ABS\&quot;. | [default to &quot;ABS&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetStructuresByAgencyIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **detail** | **string** | Specify the desired amount of detail to be returned. For example, it is possible to instruct the web service to return only basic information about the resource, this is known in SDMX as a stub.  * **allstubs**: All artefacts will be returned as stubs. * **referencestubs**: The referenced artefacts will be returned as stubs. * **referencepartial**: The referenced item schemes should only include items used by the artefact to be returned. For example, a concept scheme would only contain the concepts used in a DSD, and its isPartial flag would be set to true. As another example, if a dataflow is constrained, then the codelists returned should only contain the subset of codes allowed by that constraint. * **allcompletestubs**: All artefacts should be returned as complete stubs, containing identification information, the artefacts&#39; name, description, annotations and isFinal information. * **referencecompletestubs**: The referenced artefacts should be returned as complete stubs, containing identification information, the artefacts&#39; name, description, annotations and isFinal information. * **full**: All available information for all artefacts will be returned. This is the default.  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml, application/vnd.sdmx.structure+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

