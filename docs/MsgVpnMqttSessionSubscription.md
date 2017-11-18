# MsgVpnMqttSessionSubscription

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MqttSessionClientId** | **string** | The client-id of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet. | [optional] [default to null]
**MqttSessionVirtualRouter** | **string** | The Virtual Router of the MQTT Session. The allowed values and their meaning are:  &lt;pre&gt; \&quot;primary\&quot; - MQTT Session belongs to the Primary Virtual Router. \&quot;backup\&quot; - MQTT Session belongs to the Backup Virtual Router. &lt;/pre&gt;  | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN. | [optional] [default to null]
**SubscriptionQos** | **int64** | The quality of service for the subscription. The value is either &#x60;0&#x60; (deliver at most once) or &#x60;1&#x60; (deliver at least once). The default value is &#x60;0&#x60;. | [optional] [default to null]
**SubscriptionTopic** | **string** | An MQTT topic string. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


