# \ReplicatedTopicApi

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMsgVpnReplicatedTopic**](ReplicatedTopicApi.md#CreateMsgVpnReplicatedTopic) | **Post** /msgVpns/{msgVpnName}/replicatedTopics | Creates a Replicated Topic object.
[**DeleteMsgVpnReplicatedTopic**](ReplicatedTopicApi.md#DeleteMsgVpnReplicatedTopic) | **Delete** /msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic} | Deletes a Replicated Topic object.
[**GetMsgVpnReplicatedTopic**](ReplicatedTopicApi.md#GetMsgVpnReplicatedTopic) | **Get** /msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic} | Gets a Replicated Topic object.
[**GetMsgVpnReplicatedTopics**](ReplicatedTopicApi.md#GetMsgVpnReplicatedTopics) | **Get** /msgVpns/{msgVpnName}/replicatedTopics | Gets a list of Replicated Topic objects.
[**ReplaceMsgVpnReplicatedTopic**](ReplicatedTopicApi.md#ReplaceMsgVpnReplicatedTopic) | **Put** /msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic} | Replaces a Replicated Topic object.
[**UpdateMsgVpnReplicatedTopic**](ReplicatedTopicApi.md#UpdateMsgVpnReplicatedTopic) | **Patch** /msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic} | Updates a Replicated Topic object.


# **CreateMsgVpnReplicatedTopic**
> MsgVpnReplicatedTopicResponse CreateMsgVpnReplicatedTopic($msgVpnName, $body, $select_)

Creates a Replicated Topic object.

Creates a Replicated Topic object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x|| replicatedTopic|x|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.1.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **body** | [**MsgVpnReplicatedTopic**](MsgVpnReplicatedTopic.md)| The Replicated Topic object&#39;s attributes. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnReplicatedTopicResponse**](MsgVpnReplicatedTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnReplicatedTopic**
> SempMetaOnlyResponse DeleteMsgVpnReplicatedTopic($msgVpnName, $replicatedTopic)

Deletes a Replicated Topic object.

Deletes a Replicated Topic object.  A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.1.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **replicatedTopic** | **string**| The replicatedTopic of the Replicated Topic. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnReplicatedTopic**
> MsgVpnReplicatedTopicResponse GetMsgVpnReplicatedTopic($msgVpnName, $replicatedTopic, $select_)

Gets a Replicated Topic object.

Gets a Replicated Topic object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: msgVpnName|x|| replicatedTopic|x||    A SEMP client authorized with a minimum access scope/level of \"vpn/readonly\" is required to perform this operation.  This has been available since 2.1.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **replicatedTopic** | **string**| The replicatedTopic of the Replicated Topic. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnReplicatedTopicResponse**](MsgVpnReplicatedTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnReplicatedTopics**
> MsgVpnReplicatedTopicsResponse GetMsgVpnReplicatedTopics($msgVpnName, $count, $cursor, $where, $select_)

Gets a list of Replicated Topic objects.

Gets a list of Replicated Topic objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: msgVpnName|x|| replicatedTopic|x||    A SEMP client authorized with a minimum access scope/level of \"vpn/readonly\" is required to perform this operation.  This has been available since 2.1.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **count** | **int32**| Limit the count of objects in the response. See [Count](#count \&quot;Description of the syntax of the &#x60;count&#x60; parameter\&quot;). | [optional] [default to 10]
 **cursor** | **string**| The cursor, or position, for the next page of objects. See [Cursor](#cursor \&quot;Description of the syntax of the &#x60;cursor&#x60; parameter\&quot;). | [optional] 
 **where** | [**[]string**](string.md)| Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;). | [optional] 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnReplicatedTopicsResponse**](MsgVpnReplicatedTopicsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnReplicatedTopic**
> MsgVpnReplicatedTopicResponse ReplaceMsgVpnReplicatedTopic($msgVpnName, $replicatedTopic, $body, $select_)

Replaces a Replicated Topic object.

Replaces a Replicated Topic object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: msgVpnName|x|x||| replicatedTopic|x|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.1.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **replicatedTopic** | **string**| The replicatedTopic of the Replicated Topic. | 
 **body** | [**MsgVpnReplicatedTopic**](MsgVpnReplicatedTopic.md)| The Replicated Topic object&#39;s attributes. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnReplicatedTopicResponse**](MsgVpnReplicatedTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnReplicatedTopic**
> MsgVpnReplicatedTopicResponse UpdateMsgVpnReplicatedTopic($msgVpnName, $replicatedTopic, $body, $select_)

Updates a Replicated Topic object.

Updates a Replicated Topic object. Any attribute missing from the request will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: msgVpnName|x|x||| replicatedTopic|x|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.1.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **replicatedTopic** | **string**| The replicatedTopic of the Replicated Topic. | 
 **body** | [**MsgVpnReplicatedTopic**](MsgVpnReplicatedTopic.md)| The Replicated Topic object&#39;s attributes. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnReplicatedTopicResponse**](MsgVpnReplicatedTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

