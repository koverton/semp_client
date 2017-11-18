# \MqttSessionApi

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMsgVpnMqttSession**](MqttSessionApi.md#CreateMsgVpnMqttSession) | **Post** /msgVpns/{msgVpnName}/mqttSessions | Creates an MQTT Session object.
[**CreateMsgVpnMqttSessionSubscription**](MqttSessionApi.md#CreateMsgVpnMqttSessionSubscription) | **Post** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions | Creates an MQTT Session Subscription object.
[**DeleteMsgVpnMqttSession**](MqttSessionApi.md#DeleteMsgVpnMqttSession) | **Delete** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter} | Deletes an MQTT Session object.
[**DeleteMsgVpnMqttSessionSubscription**](MqttSessionApi.md#DeleteMsgVpnMqttSessionSubscription) | **Delete** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic} | Deletes an MQTT Session Subscription object.
[**GetMsgVpnMqttSession**](MqttSessionApi.md#GetMsgVpnMqttSession) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter} | Gets an MQTT Session object.
[**GetMsgVpnMqttSessionSubscription**](MqttSessionApi.md#GetMsgVpnMqttSessionSubscription) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic} | Gets an MQTT Session Subscription object.
[**GetMsgVpnMqttSessionSubscriptions**](MqttSessionApi.md#GetMsgVpnMqttSessionSubscriptions) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions | Gets a list of MQTT Session Subscription objects.
[**GetMsgVpnMqttSessions**](MqttSessionApi.md#GetMsgVpnMqttSessions) | **Get** /msgVpns/{msgVpnName}/mqttSessions | Gets a list of MQTT Session objects.
[**ReplaceMsgVpnMqttSession**](MqttSessionApi.md#ReplaceMsgVpnMqttSession) | **Put** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter} | Replaces an MQTT Session object.
[**ReplaceMsgVpnMqttSessionSubscription**](MqttSessionApi.md#ReplaceMsgVpnMqttSessionSubscription) | **Put** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic} | Replaces an MQTT Session Subscription object.
[**UpdateMsgVpnMqttSession**](MqttSessionApi.md#UpdateMsgVpnMqttSession) | **Patch** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter} | Updates an MQTT Session object.
[**UpdateMsgVpnMqttSessionSubscription**](MqttSessionApi.md#UpdateMsgVpnMqttSessionSubscription) | **Patch** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic} | Updates an MQTT Session Subscription object.


# **CreateMsgVpnMqttSession**
> MsgVpnMqttSessionResponse CreateMsgVpnMqttSession($msgVpnName, $body, $select_)

Creates an MQTT Session object.

Creates an MQTT Session object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: mqttSessionClientId|x|x||| mqttSessionVirtualRouter|x|x||| msgVpnName|x||x||    A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.1.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **body** | [**MsgVpnMqttSession**](MsgVpnMqttSession.md)| The MQTT Session object&#39;s attributes. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnMqttSessionResponse**](MsgVpnMqttSessionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnMqttSessionSubscription**
> MsgVpnMqttSessionSubscriptionResponse CreateMsgVpnMqttSessionSubscription($msgVpnName, $mqttSessionClientId, $mqttSessionVirtualRouter, $body, $select_)

Creates an MQTT Session Subscription object.

Creates an MQTT Session Subscription object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: mqttSessionClientId|x||x|| mqttSessionVirtualRouter|x||x|| msgVpnName|x||x|| subscriptionTopic|x|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.1.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **mqttSessionClientId** | **string**| The mqttSessionClientId of the MQTT Session. | 
 **mqttSessionVirtualRouter** | **string**| The mqttSessionVirtualRouter of the MQTT Session. | 
 **body** | [**MsgVpnMqttSessionSubscription**](MsgVpnMqttSessionSubscription.md)| The MQTT Session Subscription object&#39;s attributes. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnMqttSessionSubscriptionResponse**](MsgVpnMqttSessionSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnMqttSession**
> SempMetaOnlyResponse DeleteMsgVpnMqttSession($msgVpnName, $mqttSessionClientId, $mqttSessionVirtualRouter)

Deletes an MQTT Session object.

Deletes an MQTT Session object.  A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.1.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **mqttSessionClientId** | **string**| The mqttSessionClientId of the MQTT Session. | 
 **mqttSessionVirtualRouter** | **string**| The mqttSessionVirtualRouter of the MQTT Session. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnMqttSessionSubscription**
> SempMetaOnlyResponse DeleteMsgVpnMqttSessionSubscription($msgVpnName, $mqttSessionClientId, $mqttSessionVirtualRouter, $subscriptionTopic)

Deletes an MQTT Session Subscription object.

Deletes an MQTT Session Subscription object.  A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.1.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **mqttSessionClientId** | **string**| The mqttSessionClientId of the MQTT Session. | 
 **mqttSessionVirtualRouter** | **string**| The mqttSessionVirtualRouter of the MQTT Session. | 
 **subscriptionTopic** | **string**| The subscriptionTopic of the MQTT Session Subscription. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnMqttSession**
> MsgVpnMqttSessionResponse GetMsgVpnMqttSession($msgVpnName, $mqttSessionClientId, $mqttSessionVirtualRouter, $select_)

Gets an MQTT Session object.

Gets an MQTT Session object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: mqttSessionClientId|x|| mqttSessionVirtualRouter|x|| msgVpnName|x||    A SEMP client authorized with a minimum access scope/level of \"vpn/readonly\" is required to perform this operation.  This has been available since 2.1.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **mqttSessionClientId** | **string**| The mqttSessionClientId of the MQTT Session. | 
 **mqttSessionVirtualRouter** | **string**| The mqttSessionVirtualRouter of the MQTT Session. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnMqttSessionResponse**](MsgVpnMqttSessionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnMqttSessionSubscription**
> MsgVpnMqttSessionSubscriptionResponse GetMsgVpnMqttSessionSubscription($msgVpnName, $mqttSessionClientId, $mqttSessionVirtualRouter, $subscriptionTopic, $select_)

Gets an MQTT Session Subscription object.

Gets an MQTT Session Subscription object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: mqttSessionClientId|x|| mqttSessionVirtualRouter|x|| msgVpnName|x|| subscriptionTopic|x||    A SEMP client authorized with a minimum access scope/level of \"vpn/readonly\" is required to perform this operation.  This has been available since 2.1.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **mqttSessionClientId** | **string**| The mqttSessionClientId of the MQTT Session. | 
 **mqttSessionVirtualRouter** | **string**| The mqttSessionVirtualRouter of the MQTT Session. | 
 **subscriptionTopic** | **string**| The subscriptionTopic of the MQTT Session Subscription. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnMqttSessionSubscriptionResponse**](MsgVpnMqttSessionSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnMqttSessionSubscriptions**
> MsgVpnMqttSessionSubscriptionsResponse GetMsgVpnMqttSessionSubscriptions($msgVpnName, $mqttSessionClientId, $mqttSessionVirtualRouter, $count, $cursor, $where, $select_)

Gets a list of MQTT Session Subscription objects.

Gets a list of MQTT Session Subscription objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: mqttSessionClientId|x|| mqttSessionVirtualRouter|x|| msgVpnName|x|| subscriptionTopic|x||    A SEMP client authorized with a minimum access scope/level of \"vpn/readonly\" is required to perform this operation.  This has been available since 2.1.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **mqttSessionClientId** | **string**| The mqttSessionClientId of the MQTT Session. | 
 **mqttSessionVirtualRouter** | **string**| The mqttSessionVirtualRouter of the MQTT Session. | 
 **count** | **int32**| Limit the count of objects in the response. See [Count](#count \&quot;Description of the syntax of the &#x60;count&#x60; parameter\&quot;). | [optional] [default to 10]
 **cursor** | **string**| The cursor, or position, for the next page of objects. See [Cursor](#cursor \&quot;Description of the syntax of the &#x60;cursor&#x60; parameter\&quot;). | [optional] 
 **where** | [**[]string**](string.md)| Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;). | [optional] 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnMqttSessionSubscriptionsResponse**](MsgVpnMqttSessionSubscriptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnMqttSessions**
> MsgVpnMqttSessionsResponse GetMsgVpnMqttSessions($msgVpnName, $count, $cursor, $where, $select_)

Gets a list of MQTT Session objects.

Gets a list of MQTT Session objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: mqttSessionClientId|x|| mqttSessionVirtualRouter|x|| msgVpnName|x||    A SEMP client authorized with a minimum access scope/level of \"vpn/readonly\" is required to perform this operation.  This has been available since 2.1.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **count** | **int32**| Limit the count of objects in the response. See [Count](#count \&quot;Description of the syntax of the &#x60;count&#x60; parameter\&quot;). | [optional] [default to 10]
 **cursor** | **string**| The cursor, or position, for the next page of objects. See [Cursor](#cursor \&quot;Description of the syntax of the &#x60;cursor&#x60; parameter\&quot;). | [optional] 
 **where** | [**[]string**](string.md)| Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;). | [optional] 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnMqttSessionsResponse**](MsgVpnMqttSessionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnMqttSession**
> MsgVpnMqttSessionResponse ReplaceMsgVpnMqttSession($msgVpnName, $mqttSessionClientId, $mqttSessionVirtualRouter, $body, $select_)

Replaces an MQTT Session object.

Replaces an MQTT Session object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: mqttSessionClientId|x|x||| mqttSessionVirtualRouter|x|x||| msgVpnName|x|x||| owner||||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.1.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **mqttSessionClientId** | **string**| The mqttSessionClientId of the MQTT Session. | 
 **mqttSessionVirtualRouter** | **string**| The mqttSessionVirtualRouter of the MQTT Session. | 
 **body** | [**MsgVpnMqttSession**](MsgVpnMqttSession.md)| The MQTT Session object&#39;s attributes. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnMqttSessionResponse**](MsgVpnMqttSessionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnMqttSessionSubscription**
> MsgVpnMqttSessionSubscriptionResponse ReplaceMsgVpnMqttSessionSubscription($msgVpnName, $mqttSessionClientId, $mqttSessionVirtualRouter, $subscriptionTopic, $body, $select_)

Replaces an MQTT Session Subscription object.

Replaces an MQTT Session Subscription object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: mqttSessionClientId|x|x||| mqttSessionVirtualRouter|x|x||| msgVpnName|x|x||| subscriptionTopic|x|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.1.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **mqttSessionClientId** | **string**| The mqttSessionClientId of the MQTT Session. | 
 **mqttSessionVirtualRouter** | **string**| The mqttSessionVirtualRouter of the MQTT Session. | 
 **subscriptionTopic** | **string**| The subscriptionTopic of the MQTT Session Subscription. | 
 **body** | [**MsgVpnMqttSessionSubscription**](MsgVpnMqttSessionSubscription.md)| The MQTT Session Subscription object&#39;s attributes. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnMqttSessionSubscriptionResponse**](MsgVpnMqttSessionSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnMqttSession**
> MsgVpnMqttSessionResponse UpdateMsgVpnMqttSession($msgVpnName, $mqttSessionClientId, $mqttSessionVirtualRouter, $body, $select_)

Updates an MQTT Session object.

Updates an MQTT Session object. Any attribute missing from the request will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: mqttSessionClientId|x|x||| mqttSessionVirtualRouter|x|x||| msgVpnName|x|x||| owner||||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.1.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **mqttSessionClientId** | **string**| The mqttSessionClientId of the MQTT Session. | 
 **mqttSessionVirtualRouter** | **string**| The mqttSessionVirtualRouter of the MQTT Session. | 
 **body** | [**MsgVpnMqttSession**](MsgVpnMqttSession.md)| The MQTT Session object&#39;s attributes. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnMqttSessionResponse**](MsgVpnMqttSessionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnMqttSessionSubscription**
> MsgVpnMqttSessionSubscriptionResponse UpdateMsgVpnMqttSessionSubscription($msgVpnName, $mqttSessionClientId, $mqttSessionVirtualRouter, $subscriptionTopic, $body, $select_)

Updates an MQTT Session Subscription object.

Updates an MQTT Session Subscription object. Any attribute missing from the request will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: mqttSessionClientId|x|x||| mqttSessionVirtualRouter|x|x||| msgVpnName|x|x||| subscriptionTopic|x|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.1.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **mqttSessionClientId** | **string**| The mqttSessionClientId of the MQTT Session. | 
 **mqttSessionVirtualRouter** | **string**| The mqttSessionVirtualRouter of the MQTT Session. | 
 **subscriptionTopic** | **string**| The subscriptionTopic of the MQTT Session Subscription. | 
 **body** | [**MsgVpnMqttSessionSubscription**](MsgVpnMqttSessionSubscription.md)| The MQTT Session Subscription object&#39;s attributes. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnMqttSessionSubscriptionResponse**](MsgVpnMqttSessionSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

