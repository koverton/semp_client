# \ClientProfileApi

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMsgVpnClientProfile**](ClientProfileApi.md#CreateMsgVpnClientProfile) | **Post** /msgVpns/{msgVpnName}/clientProfiles | Creates a Client Profile object.
[**DeleteMsgVpnClientProfile**](ClientProfileApi.md#DeleteMsgVpnClientProfile) | **Delete** /msgVpns/{msgVpnName}/clientProfiles/{clientProfileName} | Deletes a Client Profile object.
[**GetMsgVpnClientProfile**](ClientProfileApi.md#GetMsgVpnClientProfile) | **Get** /msgVpns/{msgVpnName}/clientProfiles/{clientProfileName} | Gets a Client Profile object.
[**GetMsgVpnClientProfiles**](ClientProfileApi.md#GetMsgVpnClientProfiles) | **Get** /msgVpns/{msgVpnName}/clientProfiles | Gets a list of Client Profile objects.
[**ReplaceMsgVpnClientProfile**](ClientProfileApi.md#ReplaceMsgVpnClientProfile) | **Put** /msgVpns/{msgVpnName}/clientProfiles/{clientProfileName} | Replaces a Client Profile object.
[**UpdateMsgVpnClientProfile**](ClientProfileApi.md#UpdateMsgVpnClientProfile) | **Patch** /msgVpns/{msgVpnName}/clientProfiles/{clientProfileName} | Updates a Client Profile object.


# **CreateMsgVpnClientProfile**
> MsgVpnClientProfileResponse CreateMsgVpnClientProfile($msgVpnName, $body, $select_)

Creates a Client Profile object.

Creates a Client Profile object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: clientProfileName|x|x||| msgVpnName|x||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent EventThresholdByPercent|clearPercent|setPercent| EventThresholdByPercent|setPercent|clearPercent|    A SEMP client authorized with a minimum access scope/level of \"global/readwrite\" is required to perform this operation.  This has been available since 2.0.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **body** | [**MsgVpnClientProfile**](MsgVpnClientProfile.md)| The Client Profile object&#39;s attributes. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnClientProfileResponse**](MsgVpnClientProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnClientProfile**
> SempMetaOnlyResponse DeleteMsgVpnClientProfile($msgVpnName, $clientProfileName)

Deletes a Client Profile object.

Deletes a Client Profile object.  A SEMP client authorized with a minimum access scope/level of \"global/readwrite\" is required to perform this operation.  This has been available since 2.0.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **clientProfileName** | **string**| The clientProfileName of the Client Profile. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnClientProfile**
> MsgVpnClientProfileResponse GetMsgVpnClientProfile($msgVpnName, $clientProfileName, $select_)

Gets a Client Profile object.

Gets a Client Profile object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: clientProfileName|x|| msgVpnName|x||    A SEMP client authorized with a minimum access scope/level of \"vpn/readonly\" is required to perform this operation.  This has been available since 2.0.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **clientProfileName** | **string**| The clientProfileName of the Client Profile. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnClientProfileResponse**](MsgVpnClientProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnClientProfiles**
> MsgVpnClientProfilesResponse GetMsgVpnClientProfiles($msgVpnName, $count, $cursor, $where, $select_)

Gets a list of Client Profile objects.

Gets a list of Client Profile objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: clientProfileName|x|| msgVpnName|x||    A SEMP client authorized with a minimum access scope/level of \"vpn/readonly\" is required to perform this operation.  This has been available since 2.0.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **count** | **int32**| Limit the count of objects in the response. See [Count](#count \&quot;Description of the syntax of the &#x60;count&#x60; parameter\&quot;). | [optional] [default to 10]
 **cursor** | **string**| The cursor, or position, for the next page of objects. See [Cursor](#cursor \&quot;Description of the syntax of the &#x60;cursor&#x60; parameter\&quot;). | [optional] 
 **where** | [**[]string**](string.md)| Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;). | [optional] 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnClientProfilesResponse**](MsgVpnClientProfilesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnClientProfile**
> MsgVpnClientProfileResponse ReplaceMsgVpnClientProfile($msgVpnName, $clientProfileName, $body, $select_)

Replaces a Client Profile object.

Replaces a Client Profile object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: clientProfileName|x|x||| msgVpnName|x|x|||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent EventThresholdByPercent|clearPercent|setPercent| EventThresholdByPercent|setPercent|clearPercent|    A SEMP client authorized with a minimum access scope/level of \"global/readwrite\" is required to perform this operation.  This has been available since 2.0.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **clientProfileName** | **string**| The clientProfileName of the Client Profile. | 
 **body** | [**MsgVpnClientProfile**](MsgVpnClientProfile.md)| The Client Profile object&#39;s attributes. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnClientProfileResponse**](MsgVpnClientProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnClientProfile**
> MsgVpnClientProfileResponse UpdateMsgVpnClientProfile($msgVpnName, $clientProfileName, $body, $select_)

Updates a Client Profile object.

Updates a Client Profile object. Any attribute missing from the request will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: clientProfileName|x|x||| msgVpnName|x|x|||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent EventThresholdByPercent|clearPercent|setPercent| EventThresholdByPercent|setPercent|clearPercent|    A SEMP client authorized with a minimum access scope/level of \"global/readwrite\" is required to perform this operation.  This has been available since 2.0.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **clientProfileName** | **string**| The clientProfileName of the Client Profile. | 
 **body** | [**MsgVpnClientProfile**](MsgVpnClientProfile.md)| The Client Profile object&#39;s attributes. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnClientProfileResponse**](MsgVpnClientProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

