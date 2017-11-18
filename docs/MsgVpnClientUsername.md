# MsgVpnClientUsername

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclProfileName** | **string** | The acl-profile of this client-username. The default value is &#x60;\&quot;default\&quot;&#x60;. | [optional] [default to null]
**ClientProfileName** | **string** | The client-profile of this client-username. The default value is &#x60;\&quot;default\&quot;&#x60;. | [optional] [default to null]
**ClientUsername** | **string** | The name of the Client Username. | [optional] [default to null]
**Enabled** | **bool** | Enables or disables a client-username. When disabled all clients currently connected referencing this client username are disconnected. The default value is &#x60;false&#x60;. | [optional] [default to null]
**GuaranteedEndpointPermissionOverrideEnabled** | **bool** | Enables or disables guaranteed endpoint permission override for a client-username. When enabled all guaranteed endpoints may be accessed, modified or deleted with the same permission as the owner. The default value is &#x60;false&#x60;. | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN. | [optional] [default to null]
**Password** | **string** | The password of this client-username for internal authentication. The default is to have no &#x60;password&#x60;. | [optional] [default to null]
**SubscriptionManagerEnabled** | **bool** | Enables or disables subscription management capability. This is the ability to manage subscriptions on behalf of other client names. The default value is &#x60;false&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


