# \SystemInformationApi

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSystemInformation**](SystemInformationApi.md#GetSystemInformation) | **Get** /systemInformation | Gets SEMP API version and platform information.


# **GetSystemInformation**
> SystemInformationResponse GetSystemInformation()

Gets SEMP API version and platform information.

Gets SEMP API version and platform information.  A SEMP client authorized with a minimum access scope/level of \"global/none\" is required to perform this operation.  This has been available since 2.1.0.  This has been deprecated since 2.2.0.


### Parameters
This endpoint does not need any parameter.

### Return type

[**SystemInformationResponse**](SystemInformationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

