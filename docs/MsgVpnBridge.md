# MsgVpnBridge

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BridgeName** | **string** | The name of the Bridge. | [optional] [default to null]
**BridgeVirtualRouter** | **string** | The Virtual Router of the Bridge. The allowed values and their meaning are:  &lt;pre&gt; \&quot;primary\&quot; - Bridge belongs to the Primary Virtual Router. \&quot;backup\&quot; - Bridge belongs to the Backup Virtual Router. &lt;/pre&gt;  | [optional] [default to null]
**Enabled** | **bool** | Enable or disable the bridge. The default value is &#x60;false&#x60;. | [optional] [default to null]
**MaxTtl** | **int64** | The max-ttl value for the bridge, in hops. When a bridge is sending a message to the remote router, the TTL value for the message becomes the lower of its current TTL value or this value. The default value is &#x60;8&#x60;. | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN. | [optional] [default to null]
**RemoteAuthenticationBasicClientUsername** | **string** | The client username the bridge uses to login to the Remote Message VPN. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] [default to null]
**RemoteAuthenticationBasicPassword** | **string** | The password for the client username the bridge uses to login to the Remote Message VPN. The default is to have no &#x60;remoteAuthenticationBasicPassword&#x60;. | [optional] [default to null]
**RemoteAuthenticationScheme** | **string** | The authentication scheme for the Remote Message VPN. The default value is &#x60;\&quot;basic\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;basic\&quot; - Basic Authentication Scheme (via username and password). \&quot;client-certificate\&quot; - Client Certificate Authentication Scheme (via certificate-file). &lt;/pre&gt;  | [optional] [default to null]
**RemoteConnectionRetryCount** | **int64** | The number of retries that are attempted for a router name before the next remote router alternative is attempted. The default value is &#x60;0&#x60;. | [optional] [default to null]
**RemoteConnectionRetryDelay** | **int64** | The number of seconds that must pass before retrying a connection. The default value is &#x60;3&#x60;. | [optional] [default to null]
**RemoteDeliverToOnePriority** | **string** | The deliver-to-one priority for the bridge used on the remote router. The default value is &#x60;\&quot;p1\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;p1\&quot; - Priority 1 (highest). \&quot;p2\&quot; - Priority 2. \&quot;p3\&quot; - Priority 3. \&quot;p4\&quot; - Priority 4 (lowest). \&quot;da\&quot; - Deliver Always. &lt;/pre&gt;  | [optional] [default to null]
**TlsCipherSuiteList** | **string** | The colon-separated list of of cipher suites for the TLS authentication mechanism. The suite selected will be the first suite in the list that is supported by the remote router. The default value is &#x60;\&quot;ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-SHA384:ECDHE-RSA-AES256-SHA:AES256-GCM-SHA384:AES256-SHA256:AES256-SHA:ECDHE-RSA-DES-CBC3-SHA:DES-CBC3-SHA:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-SHA256:ECDHE-RSA-AES128-SHA:AES128-GCM-SHA256:AES128-SHA256:AES128-SHA\&quot;&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


