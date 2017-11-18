# \BridgeApi

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMsgVpnBridge**](BridgeApi.md#CreateMsgVpnBridge) | **Post** /msgVpns/{msgVpnName}/bridges | Creates a Bridge object.
[**CreateMsgVpnBridgeRemoteMsgVpn**](BridgeApi.md#CreateMsgVpnBridgeRemoteMsgVpn) | **Post** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns | Creates a Remote Message VPN object.
[**CreateMsgVpnBridgeRemoteSubscription**](BridgeApi.md#CreateMsgVpnBridgeRemoteSubscription) | **Post** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions | Creates a Remote Subscription object.
[**CreateMsgVpnBridgeTlsTrustedCommonName**](BridgeApi.md#CreateMsgVpnBridgeTlsTrustedCommonName) | **Post** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames | Creates a Trusted Common Name object.
[**DeleteMsgVpnBridge**](BridgeApi.md#DeleteMsgVpnBridge) | **Delete** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Deletes a Bridge object.
[**DeleteMsgVpnBridgeRemoteMsgVpn**](BridgeApi.md#DeleteMsgVpnBridgeRemoteMsgVpn) | **Delete** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface} | Deletes a Remote Message VPN object.
[**DeleteMsgVpnBridgeRemoteSubscription**](BridgeApi.md#DeleteMsgVpnBridgeRemoteSubscription) | **Delete** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions/{remoteSubscriptionTopic} | Deletes a Remote Subscription object.
[**DeleteMsgVpnBridgeTlsTrustedCommonName**](BridgeApi.md#DeleteMsgVpnBridgeTlsTrustedCommonName) | **Delete** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Deletes a Trusted Common Name object.
[**GetMsgVpnBridge**](BridgeApi.md#GetMsgVpnBridge) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Gets a Bridge object.
[**GetMsgVpnBridgeRemoteMsgVpn**](BridgeApi.md#GetMsgVpnBridgeRemoteMsgVpn) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface} | Gets a Remote Message VPN object.
[**GetMsgVpnBridgeRemoteMsgVpns**](BridgeApi.md#GetMsgVpnBridgeRemoteMsgVpns) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns | Gets a list of Remote Message VPN objects.
[**GetMsgVpnBridgeRemoteSubscription**](BridgeApi.md#GetMsgVpnBridgeRemoteSubscription) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions/{remoteSubscriptionTopic} | Gets a Remote Subscription object.
[**GetMsgVpnBridgeRemoteSubscriptions**](BridgeApi.md#GetMsgVpnBridgeRemoteSubscriptions) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions | Gets a list of Remote Subscription objects.
[**GetMsgVpnBridgeTlsTrustedCommonName**](BridgeApi.md#GetMsgVpnBridgeTlsTrustedCommonName) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Gets a Trusted Common Name object.
[**GetMsgVpnBridgeTlsTrustedCommonNames**](BridgeApi.md#GetMsgVpnBridgeTlsTrustedCommonNames) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames | Gets a list of Trusted Common Name objects.
[**GetMsgVpnBridges**](BridgeApi.md#GetMsgVpnBridges) | **Get** /msgVpns/{msgVpnName}/bridges | Gets a list of Bridge objects.
[**ReplaceMsgVpnBridge**](BridgeApi.md#ReplaceMsgVpnBridge) | **Put** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Replaces a Bridge object.
[**ReplaceMsgVpnBridgeRemoteMsgVpn**](BridgeApi.md#ReplaceMsgVpnBridgeRemoteMsgVpn) | **Put** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface} | Replaces a Remote Message VPN object.
[**UpdateMsgVpnBridge**](BridgeApi.md#UpdateMsgVpnBridge) | **Patch** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Updates a Bridge object.
[**UpdateMsgVpnBridgeRemoteMsgVpn**](BridgeApi.md#UpdateMsgVpnBridgeRemoteMsgVpn) | **Patch** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface} | Updates a Remote Message VPN object.


# **CreateMsgVpnBridge**
> MsgVpnBridgeResponse CreateMsgVpnBridge($msgVpnName, $body, $select_)

Creates a Bridge object.

Creates a Bridge object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: bridgeName|x|x||| bridgeVirtualRouter|x|x||| msgVpnName|x||x|| remoteAuthenticationBasicPassword||||x|    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnBridge|remoteAuthenticationBasicClientUsername|remoteAuthenticationBasicPassword| MsgVpnBridge|remoteAuthenticationBasicPassword|remoteAuthenticationBasicClientUsername|    A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation. Requests which include the following attributes require greater access scope/level:   Attribute|Access Scope/Level :---|:---: tlsCipherSuiteList|global/readwrite    This has been available since 2.0.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **body** | [**MsgVpnBridge**](MsgVpnBridge.md)| The Bridge object&#39;s attributes. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnBridgeResponse**](MsgVpnBridgeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnBridgeRemoteMsgVpn**
> MsgVpnBridgeRemoteMsgVpnResponse CreateMsgVpnBridgeRemoteMsgVpn($msgVpnName, $bridgeName, $bridgeVirtualRouter, $body, $select_)

Creates a Remote Message VPN object.

Creates a Remote Message VPN object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: bridgeName|x||x|| bridgeVirtualRouter|x||x|| msgVpnName|x||x|| password||||x| remoteMsgVpnInterface|x|||| remoteMsgVpnLocation|x|x||| remoteMsgVpnName|x|x|||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnBridgeRemoteMsgVpn|clientUsername|password| MsgVpnBridgeRemoteMsgVpn|password|clientUsername|    A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.0.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **bridgeName** | **string**| The bridgeName of the Bridge. | 
 **bridgeVirtualRouter** | **string**| The bridgeVirtualRouter of the Bridge. | 
 **body** | [**MsgVpnBridgeRemoteMsgVpn**](MsgVpnBridgeRemoteMsgVpn.md)| The Remote Message VPN object&#39;s attributes. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnBridgeRemoteMsgVpnResponse**](MsgVpnBridgeRemoteMsgVpnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnBridgeRemoteSubscription**
> MsgVpnBridgeRemoteSubscriptionResponse CreateMsgVpnBridgeRemoteSubscription($msgVpnName, $bridgeName, $bridgeVirtualRouter, $body, $select_)

Creates a Remote Subscription object.

Creates a Remote Subscription object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: bridgeName|x||x|| bridgeVirtualRouter|x||x|| deliverAlwaysEnabled||x||| msgVpnName|x||x|| remoteSubscriptionTopic|x|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.0.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **bridgeName** | **string**| The bridgeName of the Bridge. | 
 **bridgeVirtualRouter** | **string**| The bridgeVirtualRouter of the Bridge. | 
 **body** | [**MsgVpnBridgeRemoteSubscription**](MsgVpnBridgeRemoteSubscription.md)| The Remote Subscription object&#39;s attributes. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnBridgeRemoteSubscriptionResponse**](MsgVpnBridgeRemoteSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnBridgeTlsTrustedCommonName**
> MsgVpnBridgeTlsTrustedCommonNameResponse CreateMsgVpnBridgeTlsTrustedCommonName($msgVpnName, $bridgeName, $bridgeVirtualRouter, $body, $select_)

Creates a Trusted Common Name object.

Creates a Trusted Common Name object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: bridgeName|x||x|| bridgeVirtualRouter|x||x|| msgVpnName|x||x|| tlsTrustedCommonName|x|x|||    A SEMP client authorized with a minimum access scope/level of \"global/readwrite\" is required to perform this operation.  This has been available since 2.0.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **bridgeName** | **string**| The bridgeName of the Bridge. | 
 **bridgeVirtualRouter** | **string**| The bridgeVirtualRouter of the Bridge. | 
 **body** | [**MsgVpnBridgeTlsTrustedCommonName**](MsgVpnBridgeTlsTrustedCommonName.md)| The Trusted Common Name object&#39;s attributes. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnBridgeTlsTrustedCommonNameResponse**](MsgVpnBridgeTlsTrustedCommonNameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnBridge**
> SempMetaOnlyResponse DeleteMsgVpnBridge($msgVpnName, $bridgeName, $bridgeVirtualRouter)

Deletes a Bridge object.

Deletes a Bridge object.  A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation. Requests which include the following attributes require greater access scope/level:   Attribute|Access Scope/Level :---|:---: tlsCipherSuiteList|global/readwrite    This has been available since 2.0.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **bridgeName** | **string**| The bridgeName of the Bridge. | 
 **bridgeVirtualRouter** | **string**| The bridgeVirtualRouter of the Bridge. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnBridgeRemoteMsgVpn**
> SempMetaOnlyResponse DeleteMsgVpnBridgeRemoteMsgVpn($msgVpnName, $bridgeName, $bridgeVirtualRouter, $remoteMsgVpnName, $remoteMsgVpnLocation, $remoteMsgVpnInterface)

Deletes a Remote Message VPN object.

Deletes a Remote Message VPN object.  A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.0.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **bridgeName** | **string**| The bridgeName of the Bridge. | 
 **bridgeVirtualRouter** | **string**| The bridgeVirtualRouter of the Bridge. | 
 **remoteMsgVpnName** | **string**| The remoteMsgVpnName of the Remote Message VPN. | 
 **remoteMsgVpnLocation** | **string**| The remoteMsgVpnLocation of the Remote Message VPN. | 
 **remoteMsgVpnInterface** | **string**| The remoteMsgVpnInterface of the Remote Message VPN. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnBridgeRemoteSubscription**
> SempMetaOnlyResponse DeleteMsgVpnBridgeRemoteSubscription($msgVpnName, $bridgeName, $bridgeVirtualRouter, $remoteSubscriptionTopic)

Deletes a Remote Subscription object.

Deletes a Remote Subscription object.  A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.0.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **bridgeName** | **string**| The bridgeName of the Bridge. | 
 **bridgeVirtualRouter** | **string**| The bridgeVirtualRouter of the Bridge. | 
 **remoteSubscriptionTopic** | **string**| The remoteSubscriptionTopic of the Remote Subscription. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnBridgeTlsTrustedCommonName**
> SempMetaOnlyResponse DeleteMsgVpnBridgeTlsTrustedCommonName($msgVpnName, $bridgeName, $bridgeVirtualRouter, $tlsTrustedCommonName)

Deletes a Trusted Common Name object.

Deletes a Trusted Common Name object.  A SEMP client authorized with a minimum access scope/level of \"global/readwrite\" is required to perform this operation.  This has been available since 2.0.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **bridgeName** | **string**| The bridgeName of the Bridge. | 
 **bridgeVirtualRouter** | **string**| The bridgeVirtualRouter of the Bridge. | 
 **tlsTrustedCommonName** | **string**| The tlsTrustedCommonName of the Trusted Common Name. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnBridge**
> MsgVpnBridgeResponse GetMsgVpnBridge($msgVpnName, $bridgeName, $bridgeVirtualRouter, $select_)

Gets a Bridge object.

Gets a Bridge object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: bridgeName|x|| bridgeVirtualRouter|x|| msgVpnName|x|| remoteAuthenticationBasicPassword||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/readonly\" is required to perform this operation.  This has been available since 2.0.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **bridgeName** | **string**| The bridgeName of the Bridge. | 
 **bridgeVirtualRouter** | **string**| The bridgeVirtualRouter of the Bridge. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnBridgeResponse**](MsgVpnBridgeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnBridgeRemoteMsgVpn**
> MsgVpnBridgeRemoteMsgVpnResponse GetMsgVpnBridgeRemoteMsgVpn($msgVpnName, $bridgeName, $bridgeVirtualRouter, $remoteMsgVpnName, $remoteMsgVpnLocation, $remoteMsgVpnInterface, $select_)

Gets a Remote Message VPN object.

Gets a Remote Message VPN object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: bridgeName|x|| bridgeVirtualRouter|x|| msgVpnName|x|| password||x| remoteMsgVpnInterface|x|| remoteMsgVpnLocation|x|| remoteMsgVpnName|x||    A SEMP client authorized with a minimum access scope/level of \"vpn/readonly\" is required to perform this operation.  This has been available since 2.0.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **bridgeName** | **string**| The bridgeName of the Bridge. | 
 **bridgeVirtualRouter** | **string**| The bridgeVirtualRouter of the Bridge. | 
 **remoteMsgVpnName** | **string**| The remoteMsgVpnName of the Remote Message VPN. | 
 **remoteMsgVpnLocation** | **string**| The remoteMsgVpnLocation of the Remote Message VPN. | 
 **remoteMsgVpnInterface** | **string**| The remoteMsgVpnInterface of the Remote Message VPN. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnBridgeRemoteMsgVpnResponse**](MsgVpnBridgeRemoteMsgVpnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnBridgeRemoteMsgVpns**
> MsgVpnBridgeRemoteMsgVpnsResponse GetMsgVpnBridgeRemoteMsgVpns($msgVpnName, $bridgeName, $bridgeVirtualRouter, $where, $select_)

Gets a list of Remote Message VPN objects.

Gets a list of Remote Message VPN objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: bridgeName|x|| bridgeVirtualRouter|x|| msgVpnName|x|| password||x| remoteMsgVpnInterface|x|| remoteMsgVpnLocation|x|| remoteMsgVpnName|x||    A SEMP client authorized with a minimum access scope/level of \"vpn/readonly\" is required to perform this operation.  This has been available since 2.0.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **bridgeName** | **string**| The bridgeName of the Bridge. | 
 **bridgeVirtualRouter** | **string**| The bridgeVirtualRouter of the Bridge. | 
 **where** | [**[]string**](string.md)| Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;). | [optional] 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnBridgeRemoteMsgVpnsResponse**](MsgVpnBridgeRemoteMsgVpnsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnBridgeRemoteSubscription**
> MsgVpnBridgeRemoteSubscriptionResponse GetMsgVpnBridgeRemoteSubscription($msgVpnName, $bridgeName, $bridgeVirtualRouter, $remoteSubscriptionTopic, $select_)

Gets a Remote Subscription object.

Gets a Remote Subscription object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: bridgeName|x|| bridgeVirtualRouter|x|| msgVpnName|x|| remoteSubscriptionTopic|x||    A SEMP client authorized with a minimum access scope/level of \"vpn/readonly\" is required to perform this operation.  This has been available since 2.0.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **bridgeName** | **string**| The bridgeName of the Bridge. | 
 **bridgeVirtualRouter** | **string**| The bridgeVirtualRouter of the Bridge. | 
 **remoteSubscriptionTopic** | **string**| The remoteSubscriptionTopic of the Remote Subscription. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnBridgeRemoteSubscriptionResponse**](MsgVpnBridgeRemoteSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnBridgeRemoteSubscriptions**
> MsgVpnBridgeRemoteSubscriptionsResponse GetMsgVpnBridgeRemoteSubscriptions($msgVpnName, $bridgeName, $bridgeVirtualRouter, $count, $cursor, $where, $select_)

Gets a list of Remote Subscription objects.

Gets a list of Remote Subscription objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: bridgeName|x|| bridgeVirtualRouter|x|| msgVpnName|x|| remoteSubscriptionTopic|x||    A SEMP client authorized with a minimum access scope/level of \"vpn/readonly\" is required to perform this operation.  This has been available since 2.0.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **bridgeName** | **string**| The bridgeName of the Bridge. | 
 **bridgeVirtualRouter** | **string**| The bridgeVirtualRouter of the Bridge. | 
 **count** | **int32**| Limit the count of objects in the response. See [Count](#count \&quot;Description of the syntax of the &#x60;count&#x60; parameter\&quot;). | [optional] [default to 10]
 **cursor** | **string**| The cursor, or position, for the next page of objects. See [Cursor](#cursor \&quot;Description of the syntax of the &#x60;cursor&#x60; parameter\&quot;). | [optional] 
 **where** | [**[]string**](string.md)| Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;). | [optional] 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnBridgeRemoteSubscriptionsResponse**](MsgVpnBridgeRemoteSubscriptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnBridgeTlsTrustedCommonName**
> MsgVpnBridgeTlsTrustedCommonNameResponse GetMsgVpnBridgeTlsTrustedCommonName($msgVpnName, $bridgeName, $bridgeVirtualRouter, $tlsTrustedCommonName, $select_)

Gets a Trusted Common Name object.

Gets a Trusted Common Name object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: bridgeName|x|| bridgeVirtualRouter|x|| msgVpnName|x|| tlsTrustedCommonName|x||    A SEMP client authorized with a minimum access scope/level of \"vpn/readonly\" is required to perform this operation.  This has been available since 2.0.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **bridgeName** | **string**| The bridgeName of the Bridge. | 
 **bridgeVirtualRouter** | **string**| The bridgeVirtualRouter of the Bridge. | 
 **tlsTrustedCommonName** | **string**| The tlsTrustedCommonName of the Trusted Common Name. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnBridgeTlsTrustedCommonNameResponse**](MsgVpnBridgeTlsTrustedCommonNameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnBridgeTlsTrustedCommonNames**
> MsgVpnBridgeTlsTrustedCommonNamesResponse GetMsgVpnBridgeTlsTrustedCommonNames($msgVpnName, $bridgeName, $bridgeVirtualRouter, $where, $select_)

Gets a list of Trusted Common Name objects.

Gets a list of Trusted Common Name objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: bridgeName|x|| bridgeVirtualRouter|x|| msgVpnName|x|| tlsTrustedCommonName|x||    A SEMP client authorized with a minimum access scope/level of \"vpn/readonly\" is required to perform this operation.  This has been available since 2.0.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **bridgeName** | **string**| The bridgeName of the Bridge. | 
 **bridgeVirtualRouter** | **string**| The bridgeVirtualRouter of the Bridge. | 
 **where** | [**[]string**](string.md)| Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;). | [optional] 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnBridgeTlsTrustedCommonNamesResponse**](MsgVpnBridgeTlsTrustedCommonNamesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnBridges**
> MsgVpnBridgesResponse GetMsgVpnBridges($msgVpnName, $count, $cursor, $where, $select_)

Gets a list of Bridge objects.

Gets a list of Bridge objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: bridgeName|x|| bridgeVirtualRouter|x|| msgVpnName|x|| remoteAuthenticationBasicPassword||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/readonly\" is required to perform this operation.  This has been available since 2.0.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **count** | **int32**| Limit the count of objects in the response. See [Count](#count \&quot;Description of the syntax of the &#x60;count&#x60; parameter\&quot;). | [optional] [default to 10]
 **cursor** | **string**| The cursor, or position, for the next page of objects. See [Cursor](#cursor \&quot;Description of the syntax of the &#x60;cursor&#x60; parameter\&quot;). | [optional] 
 **where** | [**[]string**](string.md)| Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;). | [optional] 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnBridgesResponse**](MsgVpnBridgesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnBridge**
> MsgVpnBridgeResponse ReplaceMsgVpnBridge($msgVpnName, $bridgeName, $bridgeVirtualRouter, $body, $select_)

Replaces a Bridge object.

Replaces a Bridge object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: bridgeName|x|x||| bridgeVirtualRouter|x|x||| maxTtl||||x| msgVpnName|x|x||| remoteAuthenticationBasicClientUsername||||x| remoteAuthenticationBasicPassword|||x|x| remoteAuthenticationScheme||||x| remoteDeliverToOnePriority||||x|    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnBridge|remoteAuthenticationBasicClientUsername|remoteAuthenticationBasicPassword| MsgVpnBridge|remoteAuthenticationBasicPassword|remoteAuthenticationBasicClientUsername|    A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation. Requests which include the following attributes require greater access scope/level:   Attribute|Access Scope/Level :---|:---: tlsCipherSuiteList|global/readwrite    This has been available since 2.0.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **bridgeName** | **string**| The bridgeName of the Bridge. | 
 **bridgeVirtualRouter** | **string**| The bridgeVirtualRouter of the Bridge. | 
 **body** | [**MsgVpnBridge**](MsgVpnBridge.md)| The Bridge object&#39;s attributes. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnBridgeResponse**](MsgVpnBridgeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnBridgeRemoteMsgVpn**
> MsgVpnBridgeRemoteMsgVpnResponse ReplaceMsgVpnBridgeRemoteMsgVpn($msgVpnName, $bridgeName, $bridgeVirtualRouter, $remoteMsgVpnName, $remoteMsgVpnLocation, $remoteMsgVpnInterface, $body, $select_)

Replaces a Remote Message VPN object.

Replaces a Remote Message VPN object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: bridgeName|x|x||| bridgeVirtualRouter|x|x||| clientUsername||||x| compressedDataEnabled||||x| egressFlowWindowSize||||x| msgVpnName|x|x||| password|||x|x| remoteMsgVpnInterface|x|x||| remoteMsgVpnLocation|x|x||| remoteMsgVpnName|x|x||| tlsEnabled||||x|    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnBridgeRemoteMsgVpn|clientUsername|password| MsgVpnBridgeRemoteMsgVpn|password|clientUsername|    A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.0.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **bridgeName** | **string**| The bridgeName of the Bridge. | 
 **bridgeVirtualRouter** | **string**| The bridgeVirtualRouter of the Bridge. | 
 **remoteMsgVpnName** | **string**| The remoteMsgVpnName of the Remote Message VPN. | 
 **remoteMsgVpnLocation** | **string**| The remoteMsgVpnLocation of the Remote Message VPN. | 
 **remoteMsgVpnInterface** | **string**| The remoteMsgVpnInterface of the Remote Message VPN. | 
 **body** | [**MsgVpnBridgeRemoteMsgVpn**](MsgVpnBridgeRemoteMsgVpn.md)| The Remote Message VPN object&#39;s attributes. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnBridgeRemoteMsgVpnResponse**](MsgVpnBridgeRemoteMsgVpnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnBridge**
> MsgVpnBridgeResponse UpdateMsgVpnBridge($msgVpnName, $bridgeName, $bridgeVirtualRouter, $body, $select_)

Updates a Bridge object.

Updates a Bridge object. Any attribute missing from the request will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: bridgeName|x|x||| bridgeVirtualRouter|x|x||| maxTtl||||x| msgVpnName|x|x||| remoteAuthenticationBasicClientUsername||||x| remoteAuthenticationBasicPassword|||x|x| remoteAuthenticationScheme||||x| remoteDeliverToOnePriority||||x|    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnBridge|remoteAuthenticationBasicClientUsername|remoteAuthenticationBasicPassword| MsgVpnBridge|remoteAuthenticationBasicPassword|remoteAuthenticationBasicClientUsername|    A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation. Requests which include the following attributes require greater access scope/level:   Attribute|Access Scope/Level :---|:---: tlsCipherSuiteList|global/readwrite    This has been available since 2.0.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **bridgeName** | **string**| The bridgeName of the Bridge. | 
 **bridgeVirtualRouter** | **string**| The bridgeVirtualRouter of the Bridge. | 
 **body** | [**MsgVpnBridge**](MsgVpnBridge.md)| The Bridge object&#39;s attributes. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnBridgeResponse**](MsgVpnBridgeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnBridgeRemoteMsgVpn**
> MsgVpnBridgeRemoteMsgVpnResponse UpdateMsgVpnBridgeRemoteMsgVpn($msgVpnName, $bridgeName, $bridgeVirtualRouter, $remoteMsgVpnName, $remoteMsgVpnLocation, $remoteMsgVpnInterface, $body, $select_)

Updates a Remote Message VPN object.

Updates a Remote Message VPN object. Any attribute missing from the request will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: bridgeName|x|x||| bridgeVirtualRouter|x|x||| clientUsername||||x| compressedDataEnabled||||x| egressFlowWindowSize||||x| msgVpnName|x|x||| password|||x|x| remoteMsgVpnInterface|x|x||| remoteMsgVpnLocation|x|x||| remoteMsgVpnName|x|x||| tlsEnabled||||x|    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnBridgeRemoteMsgVpn|clientUsername|password| MsgVpnBridgeRemoteMsgVpn|password|clientUsername|    A SEMP client authorized with a minimum access scope/level of \"vpn/readwrite\" is required to perform this operation.  This has been available since 2.0.0.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgVpnName** | **string**| The msgVpnName of the Message VPN. | 
 **bridgeName** | **string**| The bridgeName of the Bridge. | 
 **bridgeVirtualRouter** | **string**| The bridgeVirtualRouter of the Bridge. | 
 **remoteMsgVpnName** | **string**| The remoteMsgVpnName of the Remote Message VPN. | 
 **remoteMsgVpnLocation** | **string**| The remoteMsgVpnLocation of the Remote Message VPN. | 
 **remoteMsgVpnInterface** | **string**| The remoteMsgVpnInterface of the Remote Message VPN. | 
 **body** | [**MsgVpnBridgeRemoteMsgVpn**](MsgVpnBridgeRemoteMsgVpn.md)| The Remote Message VPN object&#39;s attributes. | 
 **select_** | [**[]string**](string.md)| Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;). | [optional] 

### Return type

[**MsgVpnBridgeRemoteMsgVpnResponse**](MsgVpnBridgeRemoteMsgVpnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

