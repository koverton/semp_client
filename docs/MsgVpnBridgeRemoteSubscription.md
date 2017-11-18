# MsgVpnBridgeRemoteSubscription

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BridgeName** | **string** | The name of the Bridge. | [optional] [default to null]
**BridgeVirtualRouter** | **string** | The Virtual Router of the Bridge. The allowed values and their meaning are:  &lt;pre&gt; \&quot;primary\&quot; - Bridge belongs to the Primary Virtual Router. \&quot;backup\&quot; - Bridge belongs to the Backup Virtual Router. &lt;/pre&gt;  | [optional] [default to null]
**DeliverAlwaysEnabled** | **bool** | Flag the topic as deliver-always instead of with the configured deliver-to-one remote-priority value for the bridge. A given topic may be deliver-to-one or deliver-always but not both. | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN. | [optional] [default to null]
**RemoteSubscriptionTopic** | **string** | The Topic of the Remote Subscription. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


