# MsgVpnReplicatedTopic

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MsgVpnName** | **string** | The name of the Message VPN. | [optional] [default to null]
**ReplicatedTopic** | **string** | Topic for applying replication. Published messages matching this topic will be replicated to the standby site. | [optional] [default to null]
**ReplicationMode** | **string** | Choose the replication-mode for the Replicated Topic. The default value is &#x60;\&quot;async\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;sync\&quot; - Synchronous replication-mode. Published messages are acknowledged when they are spooled on the standby site. \&quot;async\&quot; - Asynchronous replication-mode. Published messages are acknowledged when they are spooled locally. &lt;/pre&gt;  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


