# \ClientUsernameApi

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMsgVpnClientUsername**](ClientUsernameApi.md#CreateMsgVpnClientUsername) | **Post** /msgVpns/{msgVpnName}/clientUsernames | Creates a Client Username object.
[**DeleteMsgVpnClientUsername**](ClientUsernameApi.md#DeleteMsgVpnClientUsername) | **Delete** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername} | Deletes a Client Username object.
[**GetMsgVpnClientUsername**](ClientUsernameApi.md#GetMsgVpnClientUsername) | **Get** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername} | Gets a Client Username object.
[**GetMsgVpnClientUsernames**](ClientUsernameApi.md#GetMsgVpnClientUsernames) | **Get** /msgVpns/{msgVpnName}/clientUsernames | Gets a list of Client Username objects.
[**ReplaceMsgVpnClientUsername**](ClientUsernameApi.md#ReplaceMsgVpnClientUsername) | **Put** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername} | Replaces a Client Username object.
[**UpdateMsgVpnClientUsername**](ClientUsernameApi.md#UpdateMsgVpnClientUsername) | **Patch** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername} | Updates a Client Username object.


# **CreateMsgVpnClientUsername**
> MsgVpnClientUsernameResponse CreateMsgVpnClientUsername($msgVpnName, $body, $select_)

Creates a Client Username object.

Creates a Client Username object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: clientUsername|x|x||| msgVpnName|x||x|| password||||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.0.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **body** | [**MsgVpnClientUsername**](MsgVpnClientUsername.md)| The Client Username object&#39;s attributes. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnClientUsernameResponse**](MsgVpnClientUsernameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnClientUsername**
> SempMetaOnlyResponse DeleteMsgVpnClientUsername($msgVpnName, $clientUsername)

Deletes a Client Username object.

Deletes a Client Username object.  A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.0.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **clientUsername** | **string**| The clientUsername of the Client Username. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnClientUsername**
> MsgVpnClientUsernameResponse GetMsgVpnClientUsername($msgVpnName, $clientUsername, $select_)

Gets a Client Username object.

Gets a Client Username object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: clientUsername|x|| msgVpnName|x|| password||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/readonly\" is required to perform this operation.  This has been available since 2.0.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **clientUsername** | **string**| The clientUsername of the Client Username. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnClientUsernameResponse**](MsgVpnClientUsernameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnClientUsernames**
> MsgVpnClientUsernamesResponse GetMsgVpnClientUsernames($msgVpnName, $count, $cursor, $where, $select_)

Gets a list of Client Username objects.

Gets a list of Client Username objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: clientUsername|x|| msgVpnName|x|| password||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/readonly\" is required to perform this operation.  This has been available since 2.0.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **count** | **int32**| Limit the count of objects in the response. See [Count](#count \&quot;Description of the syntax of the &#x60;count&#x60; parameter\&quot;). | [optional] [default to 10]
 **cursor** | **string**| The cursor, or position, for the next page of objects. See [Cursor](#cursor \&quot;Description of the syntax of the &#x60;cursor&#x60; parameter\&quot;). | [optional] 
 **where** | [**[]string**](string.md)| Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;). | [optional] 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnClientUsernamesResponse**](MsgVpnClientUsernamesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnClientUsername**
> MsgVpnClientUsernameResponse ReplaceMsgVpnClientUsername($msgVpnName, $clientUsername, $body, $select_)

Replaces a Client Username object.

Replaces a Client Username object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: aclProfileName||||x| clientProfileName||||x| clientUsername|x|x||| msgVpnName|x|x||| password|||x||    A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.0.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **clientUsername** | **string**| The clientUsername of the Client Username. | 
 **body** | [**MsgVpnClientUsername**](MsgVpnClientUsername.md)| The Client Username object&#39;s attributes. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnClientUsernameResponse**](MsgVpnClientUsernameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnClientUsername**
> MsgVpnClientUsernameResponse UpdateMsgVpnClientUsername($msgVpnName, $clientUsername, $body, $select_)

Updates a Client Username object.

Updates a Client Username object. Any attribute missing from the request will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: aclProfileName||||x| clientProfileName||||x| clientUsername|x|x||| msgVpnName|x|x||| password|||x||    A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.0.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **clientUsername** | **string**| The clientUsername of the Client Username. | 
 **body** | [**MsgVpnClientUsername**](MsgVpnClientUsername.md)| The Client Username object&#39;s attributes. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnClientUsernameResponse**](MsgVpnClientUsernameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

