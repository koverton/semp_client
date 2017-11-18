# \AboutApi

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAboutApi**](AboutApi.md#GetAboutApi) | **Get** /about/api | Gets an API Description object.
[**GetAboutUser**](AboutApi.md#GetAboutUser) | **Get** /about/user | Gets a Current User object.
[**GetAboutUserMsgVpn**](AboutApi.md#GetAboutUserMsgVpn) | **Get** /about/user/msgVpns/{msgVpnName} | Gets a Current User Message VPN object.
[**GetAboutUserMsgVpns**](AboutApi.md#GetAboutUserMsgVpns) | **Get** /about/user/msgVpns | Gets a list of Current User Message VPN objects.


# **GetAboutApi**
> AboutApiResponse GetAboutApi()

Gets an API Description object.

Gets an API Description object.  A SEMP client authorized with a minimum access scope/level of \"global/none\" is required to perform this operation.  This has been available since 2.2.0.


### Parameters
This endpoint does not need any parameter.

### Return type

[**AboutApiResponse**](AboutApiResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAboutUser**
> AboutUserResponse GetAboutUser($select_)

Gets a Current User object.

Gets a Current User object.    A SEMP client authorized with a minimum access scope/level of \"global/none\" is required to perform this operation.  This has been available since 2.2.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**AboutUserResponse**](AboutUserResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAboutUserMsgVpn**
> AboutUserMsgVpnResponse GetAboutUserMsgVpn($msgVpnName, $select_)

Gets a Current User Message VPN object.

Gets a Current User Message VPN object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: msgVpnName|x||    A SEMP client authorized with a minimum access scope/level of \"global/none\" is required to perform this operation.  This has been available since 2.2.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Current User Message VPN. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**AboutUserMsgVpnResponse**](AboutUserMsgVpnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAboutUserMsgVpns**
> AboutUserMsgVpnsResponse GetAboutUserMsgVpns($where, $select_)

Gets a list of Current User Message VPN objects.

Gets a list of Current User Message VPN objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: msgVpnName|x||    A SEMP client authorized with a minimum access scope/level of \"global/none\" is required to perform this operation.  This has been available since 2.2.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **where** | [**[]string**](string.md)| Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;). | [optional] 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**AboutUserMsgVpnsResponse**](AboutUserMsgVpnsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

