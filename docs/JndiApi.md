# \JndiApi

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMsgVpnJndiConnectionFactory**](JndiApi.md#CreateMsgVpnJndiConnectionFactory) | **Post** /msgVpns/{msgVpnName}/jndiConnectionFactories | Creates a JNDI Connection Factory object.
[**CreateMsgVpnJndiQueue**](JndiApi.md#CreateMsgVpnJndiQueue) | **Post** /msgVpns/{msgVpnName}/jndiQueues | Creates a JNDI Queue object.
[**CreateMsgVpnJndiTopic**](JndiApi.md#CreateMsgVpnJndiTopic) | **Post** /msgVpns/{msgVpnName}/jndiTopics | Creates a JNDI Topic object.
[**DeleteMsgVpnJndiConnectionFactory**](JndiApi.md#DeleteMsgVpnJndiConnectionFactory) | **Delete** /msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName} | Deletes a JNDI Connection Factory object.
[**DeleteMsgVpnJndiQueue**](JndiApi.md#DeleteMsgVpnJndiQueue) | **Delete** /msgVpns/{msgVpnName}/jndiQueues/{queueName} | Deletes a JNDI Queue object.
[**DeleteMsgVpnJndiTopic**](JndiApi.md#DeleteMsgVpnJndiTopic) | **Delete** /msgVpns/{msgVpnName}/jndiTopics/{topicName} | Deletes a JNDI Topic object.
[**GetMsgVpnJndiConnectionFactories**](JndiApi.md#GetMsgVpnJndiConnectionFactories) | **Get** /msgVpns/{msgVpnName}/jndiConnectionFactories | Gets a list of JNDI Connection Factory objects.
[**GetMsgVpnJndiConnectionFactory**](JndiApi.md#GetMsgVpnJndiConnectionFactory) | **Get** /msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName} | Gets a JNDI Connection Factory object.
[**GetMsgVpnJndiQueue**](JndiApi.md#GetMsgVpnJndiQueue) | **Get** /msgVpns/{msgVpnName}/jndiQueues/{queueName} | Gets a JNDI Queue object.
[**GetMsgVpnJndiQueues**](JndiApi.md#GetMsgVpnJndiQueues) | **Get** /msgVpns/{msgVpnName}/jndiQueues | Gets a list of JNDI Queue objects.
[**GetMsgVpnJndiTopic**](JndiApi.md#GetMsgVpnJndiTopic) | **Get** /msgVpns/{msgVpnName}/jndiTopics/{topicName} | Gets a JNDI Topic object.
[**GetMsgVpnJndiTopics**](JndiApi.md#GetMsgVpnJndiTopics) | **Get** /msgVpns/{msgVpnName}/jndiTopics | Gets a list of JNDI Topic objects.
[**ReplaceMsgVpnJndiConnectionFactory**](JndiApi.md#ReplaceMsgVpnJndiConnectionFactory) | **Put** /msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName} | Replaces a JNDI Connection Factory object.
[**ReplaceMsgVpnJndiQueue**](JndiApi.md#ReplaceMsgVpnJndiQueue) | **Put** /msgVpns/{msgVpnName}/jndiQueues/{queueName} | Replaces a JNDI Queue object.
[**ReplaceMsgVpnJndiTopic**](JndiApi.md#ReplaceMsgVpnJndiTopic) | **Put** /msgVpns/{msgVpnName}/jndiTopics/{topicName} | Replaces a JNDI Topic object.
[**UpdateMsgVpnJndiConnectionFactory**](JndiApi.md#UpdateMsgVpnJndiConnectionFactory) | **Patch** /msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName} | Updates a JNDI Connection Factory object.
[**UpdateMsgVpnJndiQueue**](JndiApi.md#UpdateMsgVpnJndiQueue) | **Patch** /msgVpns/{msgVpnName}/jndiQueues/{queueName} | Updates a JNDI Queue object.
[**UpdateMsgVpnJndiTopic**](JndiApi.md#UpdateMsgVpnJndiTopic) | **Patch** /msgVpns/{msgVpnName}/jndiTopics/{topicName} | Updates a JNDI Topic object.


# **CreateMsgVpnJndiConnectionFactory**
> MsgVpnJndiConnectionFactoryResponse CreateMsgVpnJndiConnectionFactory($msgVpnName, $body, $select_)

Creates a JNDI Connection Factory object.

Creates a JNDI Connection Factory object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: connectionFactoryName|x|x||| msgVpnName|x||x||    A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.2.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **body** | [**MsgVpnJndiConnectionFactory**](MsgVpnJndiConnectionFactory.md)| The JNDI Connection Factory object&#39;s attributes. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnJndiConnectionFactoryResponse**](MsgVpnJndiConnectionFactoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnJndiQueue**
> MsgVpnJndiQueueResponse CreateMsgVpnJndiQueue($msgVpnName, $body, $select_)

Creates a JNDI Queue object.

Creates a JNDI Queue object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x|| queueName|x|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.2.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **body** | [**MsgVpnJndiQueue**](MsgVpnJndiQueue.md)| The JNDI Queue object&#39;s attributes. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnJndiQueueResponse**](MsgVpnJndiQueueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnJndiTopic**
> MsgVpnJndiTopicResponse CreateMsgVpnJndiTopic($msgVpnName, $body, $select_)

Creates a JNDI Topic object.

Creates a JNDI Topic object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x|| topicName|x|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.2.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **body** | [**MsgVpnJndiTopic**](MsgVpnJndiTopic.md)| The JNDI Topic object&#39;s attributes. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnJndiTopicResponse**](MsgVpnJndiTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnJndiConnectionFactory**
> SempMetaOnlyResponse DeleteMsgVpnJndiConnectionFactory($msgVpnName, $connectionFactoryName)

Deletes a JNDI Connection Factory object.

Deletes a JNDI Connection Factory object.  A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.2.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **connectionFactoryName** | **string**| The connectionFactoryName of the JNDI Connection Factory. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnJndiQueue**
> SempMetaOnlyResponse DeleteMsgVpnJndiQueue($msgVpnName, $queueName)

Deletes a JNDI Queue object.

Deletes a JNDI Queue object.  A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.2.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **queueName** | **string**| The queueName of the JNDI Queue. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnJndiTopic**
> SempMetaOnlyResponse DeleteMsgVpnJndiTopic($msgVpnName, $topicName)

Deletes a JNDI Topic object.

Deletes a JNDI Topic object.  A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.2.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **topicName** | **string**| The topicName of the JNDI Topic. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnJndiConnectionFactories**
> MsgVpnJndiConnectionFactoriesResponse GetMsgVpnJndiConnectionFactories($msgVpnName, $count, $cursor, $where, $select_)

Gets a list of JNDI Connection Factory objects.

Gets a list of JNDI Connection Factory objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: connectionFactoryName|x|| msgVpnName|x||    A SEMP client authorized with a minimum access scope/level of \"vpn/readonly\" is required to perform this operation.  This has been available since 2.2.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **count** | **int32**| Limit the count of objects in the response. See [Count](#count \&quot;Description of the syntax of the &#x60;count&#x60; parameter\&quot;). | [optional] [default to 10]
 **cursor** | **string**| The cursor, or position, for the next page of objects. See [Cursor](#cursor \&quot;Description of the syntax of the &#x60;cursor&#x60; parameter\&quot;). | [optional] 
 **where** | [**[]string**](string.md)| Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;). | [optional] 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnJndiConnectionFactoriesResponse**](MsgVpnJndiConnectionFactoriesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnJndiConnectionFactory**
> MsgVpnJndiConnectionFactoryResponse GetMsgVpnJndiConnectionFactory($msgVpnName, $connectionFactoryName, $select_)

Gets a JNDI Connection Factory object.

Gets a JNDI Connection Factory object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: connectionFactoryName|x|| msgVpnName|x||    A SEMP client authorized with a minimum access scope/level of \"vpn/readonly\" is required to perform this operation.  This has been available since 2.2.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **connectionFactoryName** | **string**| The connectionFactoryName of the JNDI Connection Factory. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnJndiConnectionFactoryResponse**](MsgVpnJndiConnectionFactoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnJndiQueue**
> MsgVpnJndiQueueResponse GetMsgVpnJndiQueue($msgVpnName, $queueName, $select_)

Gets a JNDI Queue object.

Gets a JNDI Queue object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: msgVpnName|x|| queueName|x||    A SEMP client authorized with a minimum access scope/level of \"vpn/readonly\" is required to perform this operation.  This has been available since 2.2.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **queueName** | **string**| The queueName of the JNDI Queue. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnJndiQueueResponse**](MsgVpnJndiQueueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnJndiQueues**
> MsgVpnJndiQueuesResponse GetMsgVpnJndiQueues($msgVpnName, $where, $select_)

Gets a list of JNDI Queue objects.

Gets a list of JNDI Queue objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: msgVpnName|x|| queueName|x||    A SEMP client authorized with a minimum access scope/level of \"vpn/readonly\" is required to perform this operation.  This has been available since 2.2.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **where** | [**[]string**](string.md)| Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;). | [optional] 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnJndiQueuesResponse**](MsgVpnJndiQueuesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnJndiTopic**
> MsgVpnJndiTopicResponse GetMsgVpnJndiTopic($msgVpnName, $topicName, $select_)

Gets a JNDI Topic object.

Gets a JNDI Topic object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: msgVpnName|x|| topicName|x||    A SEMP client authorized with a minimum access scope/level of \"vpn/readonly\" is required to perform this operation.  This has been available since 2.2.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **topicName** | **string**| The topicName of the JNDI Topic. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnJndiTopicResponse**](MsgVpnJndiTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnJndiTopics**
> MsgVpnJndiTopicsResponse GetMsgVpnJndiTopics($msgVpnName, $count, $cursor, $where, $select_)

Gets a list of JNDI Topic objects.

Gets a list of JNDI Topic objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: msgVpnName|x|| topicName|x||    A SEMP client authorized with a minimum access scope/level of \"vpn/readonly\" is required to perform this operation.  This has been available since 2.2.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **count** | **int32**| Limit the count of objects in the response. See [Count](#count \&quot;Description of the syntax of the &#x60;count&#x60; parameter\&quot;). | [optional] [default to 10]
 **cursor** | **string**| The cursor, or position, for the next page of objects. See [Cursor](#cursor \&quot;Description of the syntax of the &#x60;cursor&#x60; parameter\&quot;). | [optional] 
 **where** | [**[]string**](string.md)| Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;). | [optional] 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnJndiTopicsResponse**](MsgVpnJndiTopicsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnJndiConnectionFactory**
> MsgVpnJndiConnectionFactoryResponse ReplaceMsgVpnJndiConnectionFactory($msgVpnName, $connectionFactoryName, $body, $select_)

Replaces a JNDI Connection Factory object.

Replaces a JNDI Connection Factory object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: connectionFactoryName|x|x||| msgVpnName|x|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.2.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **connectionFactoryName** | **string**| The connectionFactoryName of the JNDI Connection Factory. | 
 **body** | [**MsgVpnJndiConnectionFactory**](MsgVpnJndiConnectionFactory.md)| The JNDI Connection Factory object&#39;s attributes. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnJndiConnectionFactoryResponse**](MsgVpnJndiConnectionFactoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnJndiQueue**
> MsgVpnJndiQueueResponse ReplaceMsgVpnJndiQueue($msgVpnName, $queueName, $body, $select_)

Replaces a JNDI Queue object.

Replaces a JNDI Queue object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: msgVpnName|x|x||| queueName|x|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.2.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **queueName** | **string**| The queueName of the JNDI Queue. | 
 **body** | [**MsgVpnJndiQueue**](MsgVpnJndiQueue.md)| The JNDI Queue object&#39;s attributes. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnJndiQueueResponse**](MsgVpnJndiQueueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnJndiTopic**
> MsgVpnJndiTopicResponse ReplaceMsgVpnJndiTopic($msgVpnName, $topicName, $body, $select_)

Replaces a JNDI Topic object.

Replaces a JNDI Topic object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: msgVpnName|x|x||| topicName|x|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.2.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **topicName** | **string**| The topicName of the JNDI Topic. | 
 **body** | [**MsgVpnJndiTopic**](MsgVpnJndiTopic.md)| The JNDI Topic object&#39;s attributes. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnJndiTopicResponse**](MsgVpnJndiTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnJndiConnectionFactory**
> MsgVpnJndiConnectionFactoryResponse UpdateMsgVpnJndiConnectionFactory($msgVpnName, $connectionFactoryName, $body, $select_)

Updates a JNDI Connection Factory object.

Updates a JNDI Connection Factory object. Any attribute missing from the request will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: connectionFactoryName|x|x||| msgVpnName|x|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.2.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **connectionFactoryName** | **string**| The connectionFactoryName of the JNDI Connection Factory. | 
 **body** | [**MsgVpnJndiConnectionFactory**](MsgVpnJndiConnectionFactory.md)| The JNDI Connection Factory object&#39;s attributes. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnJndiConnectionFactoryResponse**](MsgVpnJndiConnectionFactoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnJndiQueue**
> MsgVpnJndiQueueResponse UpdateMsgVpnJndiQueue($msgVpnName, $queueName, $body, $select_)

Updates a JNDI Queue object.

Updates a JNDI Queue object. Any attribute missing from the request will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: msgVpnName|x|x||| queueName|x|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.2.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **queueName** | **string**| The queueName of the JNDI Queue. | 
 **body** | [**MsgVpnJndiQueue**](MsgVpnJndiQueue.md)| The JNDI Queue object&#39;s attributes. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnJndiQueueResponse**](MsgVpnJndiQueueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnJndiTopic**
> MsgVpnJndiTopicResponse UpdateMsgVpnJndiTopic($msgVpnName, $topicName, $body, $select_)

Updates a JNDI Topic object.

Updates a JNDI Topic object. Any attribute missing from the request will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: msgVpnName|x|x||| topicName|x|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.2.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **topicName** | **string**| The topicName of the JNDI Topic. | 
 **body** | [**MsgVpnJndiTopic**](MsgVpnJndiTopic.md)| The JNDI Topic object&#39;s attributes. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnJndiTopicResponse**](MsgVpnJndiTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

