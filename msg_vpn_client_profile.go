/* 
 * SEMP (Solace Element Management Protocol)
 *
 * SEMP (starting in `v2`, see [note 1](#notes)) is a RESTful API for configuring, monitoring, and administering a Solace router.  SEMP uses URIs to address manageable **resources** of the Solace router.  Resources are either individual **objects**, or **collections** of objects.  This document applies to the following API:   API|Base Path|Purpose|Comments :---|:---|:---|:--- Configuration|/SEMP/v2/config|Reading and writing config state|See [note 2](#notes)    Resources are always nouns, with individual objects being singular and  collections being plural. Objects within a collection are identified by an  `obj-id`, which follows the collection name with the form  `collection-name/obj-id`. Some examples:  <pre> /SEMP/v2/config/msgVpns                       ; MsgVpn collection /SEMP/v2/config/msgVpns/finance               ; MsgVpn object named \"finance\" /SEMP/v2/config/msgVpns/finance/queues        ; Queue collection within MsgVpn \"finance\" /SEMP/v2/config/msgVpns/finance/queues/orderQ ; Queue object named \"orderQ\" within MsgVpn \"finance\" </pre>  ## Collection Resources  Collections are unordered lists of objects (unless described as otherwise), and  are described by JSON arrays. Each item in the array represents an object in  the same manner as the individual object would normally be represented. The creation of a new object is done through its collection  resource.   ## Object Resources  Objects are composed of attributes and collections, and are described by JSON  content as name/value pairs. The collections of an object are not contained  directly in the object's JSON content, rather the content includes a URI  attribute which points to the collection. This contained collection resource  must be managed as a separate resource through this URI.  At a minimum, every object has 1 or more identifying attributes, and its own  `uri` attribute which contains the URI to itself. Attributes may have any  (non-exclusively) of the following properties:   Property|Meaning|Comments :---|:---|:--- Identifying|Attribute is involved in unique identification of the object, and appears in its URI| Required|Attribute must be provided in the request| Read-Only|Attribute can only be read, not written|See [note 3](#notes) Write-Only|Attribute can only be written, not read| Requires-Disable|Attribute can only be changed when object is disabled| Deprecated|Attribute is deprecated, and will disappear in the next SEMP version|    In some requests, certain attributes may only be provided in  certain combinations with other attributes:   Relationship|Meaning :---|:--- Requires|Attribute may only be changed by a request if a particular attribute or combination of attributes is also provided in the request Conflicts|Attribute may only be provided in a request if a particular attribute or combination of attributes is not also provided in the request     ## HTTP Methods  The following HTTP methods manipulate resources in accordance with these  general principles:   Method|Resource|Meaning|Request Body|Response Body|Missing Request Attributes :---|:---|:---|:---|:---|:--- POST|Collection|Create object|Initial attribute values|Object attributes and metadata|Set to default PUT|Object|Create or replace object|New attribute values|Object attributes and metadata|Set to default (but see [note 4](#notes)) PATCH|Object|Update object|New attribute values|Object attributes and metadata|unchanged DELETE|Object|Delete object|Empty|Object metadata|N/A GET|Object|Get object|Empty|Object attributes and metadata|N/A GET|Collection|Get collection|Empty|Object attributes and collection metadata|N/A    ## Common Query Parameters  The following are some common query parameters that are supported by many  method/URI combinations. Individual URIs may document additional parameters.  Note that multiple query parameters can be used together in a single URI,  separated by the ampersand character. For example:  <pre> ; Request for the MsgVpns collection using two hypothetical query parameters ; \"q1\" and \"q2\" with values \"val1\" and \"val2\" respectively /SEMP/v2/config/msgVpns?q1=val1&q2=val2 </pre>  ### select  Include in the response only selected attributes of the object. Use this query  parameter to limit the size of the returned data for each returned object, or  return only those fields that are desired.  The value of `select` is a comma-separated list of attribute names. Names may  include the `*` wildcard (zero or more characters). Nested attribute names  are supported using periods (e.g. `parentName.childName`). If the list is  empty (i.e. `select=`) no attributes are returned; otherwise the list must  match at least one attribute name of the object. Some examples:  <pre> ; List of all MsgVpn names /SEMP/v2/config/msgVpns?select=msgVpnName  ; Authentication attributes of MsgVpn \"finance\" /SEMP/v2/config/msgVpns/finance?select=authentication*  ; Access related attributes of Queue \"orderQ\" of MsgVpn \"finance\" /SEMP/v2/config/msgVpns/finance/queues/orderQ?select=owner,permission </pre>  ### where  Include in the response only objects where certain conditions are true. Use  this query parameter to limit which objects are returned to those whose  attribute values meet the given conditions.  The value of `where` is a comma-separated list of expressions. All expressions  must be true for the object to be included in the response. Each expression  takes the form:  <pre> expression  = attribute-name OP value OP          = '==' | '!=' | '&lt;' | '&gt;' | '&lt;=' | '&gt;=' </pre>  `value` may be a number, string, `true`, or `false`, as appropriate for the  type of `attribute-name`. Greater-than and less-than comparisons only work for  numbers. A `*` in a string `value` is interpreted as a wildcard (zero or more  characters). Some examples:  <pre> ; Only enabled MsgVpns /SEMP/v2/config/msgVpns?where=enabled==true  ; Only MsgVpns using basic non-LDAP authentication /SEMP/v2/config/msgVpns?where=authenticationBasicEnabled==true,authenticationBasicType!=ldap  ; Only MsgVpns that allow more than 100 client connections /SEMP/v2/config/msgVpns?where=maxConnectionCount&gt;100  ; Only MsgVpns with msgVpnName starting with \"B\": /SEMP/v2/config/msgVpns?where=msgVpnName==B* </pre>  ### count  Limit the count of objects in the response. This can be useful to limit the  size of the response for large collections. The minimum value for `count` is  `1` and the default is `10`. There is a hidden maximum  as to prevent overloading the system. For example:  <pre> ; Up to 25 MsgVpns /SEMP/v2/config/msgVpns?count=25 </pre>  ### cursor  The cursor, or position, for the next page of objects. Cursors are opaque data  that should not be created or interpreted by SEMP clients, and should only be  used as described below.  When a request is made for a collection and there may be additional objects  available for retrieval that are not included in the initial response, the  response will include a `cursorQuery` field containing a cursor. The value  of this field can be specified in the `cursor` query parameter of a  subsequent request to retrieve the next page of objects. For convenience,  an appropriate URI is constructed automatically by the router and included  in the `nextPageUri` field of the response. This URI can be used directly  to retrieve the next page of objects.  ## Notes  Note|Description :---:|:--- 1|This specification defines SEMP starting in \"v2\", and not the original SEMP \"v1\" interface. Request and response formats between \"v1\" and \"v2\" are entirely incompatible, although both protocols share a common port configuration on the Solace router. They are differentiated by the initial portion of the URI path, one of either \"/SEMP/\" or \"/SEMP/v2/\" 2|This API is partially implemented. Only a subset of all objects are available. 3|Read-only attributes may appear in POST and PUT/PATCH requests. However, if a read-only attribute is not marked as identifying, it will be ignored during a PUT/PATCH. 4|For PUT, if the SEMP user is not authorized to modify the attribute, its value is left unchanged rather than set to default. In addition, the values of write-only attributes are not set to their defaults on a PUT. If the object does not exist, it is created first. 5|For DELETE, the body of the request currently serves no purpose and will cause an error if not empty.    
 *
 * OpenAPI spec version: 2.3.0
 * Contact: support_request@solace.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package semp_client

type MsgVpnClientProfile struct {

	// Enable or disable allowing bridge connections to login. The default value is `false`.
	AllowBridgeConnectionsEnabled bool `json:"allowBridgeConnectionsEnabled,omitempty"`

	// Enable or disable allowing a client to bind to topic endpoints or queues with cut-through forwarding. Changing this value does not affect existing sessions. The default value is `false`.
	AllowCutThroughForwardingEnabled bool `json:"allowCutThroughForwardingEnabled,omitempty"`

	// Enable or disable allowing a client to create topic endponts or queues for the receiving of persistent or non-persistent messages. Changing this value does not affect existing sessions. The default value is `false`.
	AllowGuaranteedEndpointCreateEnabled bool `json:"allowGuaranteedEndpointCreateEnabled,omitempty"`

	// Enable or disable allowing a client to bind to topic endpoints or queues for the receiving of persistent or non-persistent messages. Changing this value does not affect existing sessions. The default value is `false`.
	AllowGuaranteedMsgReceiveEnabled bool `json:"allowGuaranteedMsgReceiveEnabled,omitempty"`

	// Enable or disable allowing a client to send persistent and non-persistent messages. Changing this value does not affect existing sessions. The default value is `false`.
	AllowGuaranteedMsgSendEnabled bool `json:"allowGuaranteedMsgSendEnabled,omitempty"`

	// Enable or disable allowing a client to use trasacted sessions to bundle persistent or non-persistent message send and receives. Changing this value does not affect existing sessions. The default value is `false`.
	AllowTransactedSessionsEnabled bool `json:"allowTransactedSessionsEnabled,omitempty"`

	// The name of a queue to copy settings from when a new queue is created by an API. The referenced queue must exist. The default value is `\"\"`.
	ApiQueueManagementCopyFromOnCreateName string `json:"apiQueueManagementCopyFromOnCreateName,omitempty"`

	// The name of a topic-endpoint to copy settings from when a new topic-endpoint is created by an API. The referenced topic-endpoint must exist. The default value is `\"\"`.
	ApiTopicEndpointManagementCopyFromOnCreateName string `json:"apiTopicEndpointManagementCopyFromOnCreateName,omitempty"`

	// The name of the Client Profile.
	ClientProfileName string `json:"clientProfileName,omitempty"`

	// The eliding delay interval (in milliseconds). 0 means no delay in delivering the message to the client. The default value is `0`.
	ElidingDelay int64 `json:"elidingDelay,omitempty"`

	// Enables or disables eliding. The default value is `false`.
	ElidingEnabled bool `json:"elidingEnabled,omitempty"`

	// The maximum number of topics that can be tracked for eliding on a per client basis. The default value is `256`.
	ElidingMaxTopicCount int64 `json:"elidingMaxTopicCount,omitempty"`

	EventClientProvisionedEndpointSpoolUsageThreshold EventThresholdByPercent `json:"eventClientProvisionedEndpointSpoolUsageThreshold,omitempty"`

	EventConnectionCountPerClientUsernameThreshold EventThreshold `json:"eventConnectionCountPerClientUsernameThreshold,omitempty"`

	EventEgressFlowCountThreshold EventThreshold `json:"eventEgressFlowCountThreshold,omitempty"`

	EventEndpointCountPerClientUsernameThreshold EventThreshold `json:"eventEndpointCountPerClientUsernameThreshold,omitempty"`

	EventIngressFlowCountThreshold EventThreshold `json:"eventIngressFlowCountThreshold,omitempty"`

	EventServiceSmfConnectionCountPerClientUsernameThreshold EventThreshold `json:"eventServiceSmfConnectionCountPerClientUsernameThreshold,omitempty"`

	EventServiceWebConnectionCountPerClientUsernameThreshold EventThreshold `json:"eventServiceWebConnectionCountPerClientUsernameThreshold,omitempty"`

	EventSubscriptionCountThreshold EventThreshold `json:"eventSubscriptionCountThreshold,omitempty"`

	EventTransactedSessionCountThreshold EventThreshold `json:"eventTransactedSessionCountThreshold,omitempty"`

	EventTransactionCountThreshold EventThreshold `json:"eventTransactionCountThreshold,omitempty"`

	// The maximum number of client connections that can be simultaneously connected with the same client-username. This value may be higher than supported by the hardware. The default is the max value supported by the hardware.
	MaxConnectionCountPerClientUsername int64 `json:"maxConnectionCountPerClientUsername,omitempty"`

	// The maximum number of egress flows that can be created by a single client associated with this client-profile. The default is the max value supported by the hardware.
	MaxEgressFlowCount int64 `json:"maxEgressFlowCount,omitempty"`

	// The maximum number of queues and topic endpoints that can be created across clients using the same client-username associated with this client-profile. The default is the max value supported by the hardware.
	MaxEndpointCountPerClientUsername int64 `json:"maxEndpointCountPerClientUsername,omitempty"`

	// The maximum number of ingress flows that can be created by a single client associated with this client-profile. The default is the max value supported by the hardware.
	MaxIngressFlowCount int64 `json:"maxIngressFlowCount,omitempty"`

	// The maximum number of subscriptions for a single client associated with this client-profile. The default varies by platform.
	MaxSubscriptionCount int64 `json:"maxSubscriptionCount,omitempty"`

	// The maximum number of transacted sessions that can be created by a single client associated with this client-profile. The default value is `10`.
	MaxTransactedSessionCount int64 `json:"maxTransactedSessionCount,omitempty"`

	// The maximum number of transacted sessions that can be created by a single client associated with this client-profile. The default varies by platform.
	MaxTransactionCount int64 `json:"maxTransactionCount,omitempty"`

	// The name of the Message VPN.
	MsgVpnName string `json:"msgVpnName,omitempty"`

	// The maximum depth of the C-1 queue measured in work units. Each work unit is 2048 bytes of data. The default value is `20000`.
	QueueControl1MaxDepth int32 `json:"queueControl1MaxDepth,omitempty"`

	// The minimum number of messages that must be on the C-1 queue before its depth is checked against the `queueControl1MaxDepth` setting. The default value is `4`.
	QueueControl1MinMsgBurst int32 `json:"queueControl1MinMsgBurst,omitempty"`

	// The maximum depth of the D-1 queue measured in work units. Each work unit is 2048 bytes of data. The default value is `20000`.
	QueueDirect1MaxDepth int32 `json:"queueDirect1MaxDepth,omitempty"`

	// The minimum number of messages that must be on the D-1 queue before its depth is checked against the `queueDirect1MaxDepth` setting. The default value is `4`.
	QueueDirect1MinMsgBurst int32 `json:"queueDirect1MinMsgBurst,omitempty"`

	// The maximum depth of the D-2 queue measured in work units. Each work unit is 2048 bytes of data. The default value is `20000`.
	QueueDirect2MaxDepth int32 `json:"queueDirect2MaxDepth,omitempty"`

	// The minimum number of messages that must be on the D-2 queue before its depth is checked against the `queueDirect2MaxDepth` setting. The default value is `4`.
	QueueDirect2MinMsgBurst int32 `json:"queueDirect2MinMsgBurst,omitempty"`

	// The maximum depth of the D-3 queue measured in work units. Each work unit is 2048 bytes of data. The default value is `20000`.
	QueueDirect3MaxDepth int32 `json:"queueDirect3MaxDepth,omitempty"`

	// The minimum number of messages that must be on the D-3 queue before its depth is checked against the `queueDirect3MaxDepth` setting. The default value is `4`.
	QueueDirect3MinMsgBurst int32 `json:"queueDirect3MinMsgBurst,omitempty"`

	// The maximum depth of the G-1 queue measured in work units. Each work unit is 2048 bytes of data. The default value is `20000`.
	QueueGuaranteed1MaxDepth int32 `json:"queueGuaranteed1MaxDepth,omitempty"`

	// The minimum number of messages that must be on the G-1 queue before its depth is checked against the `queueGuaranteed1MaxDepth` setting. The default value is `255`.
	QueueGuaranteed1MinMsgBurst int32 `json:"queueGuaranteed1MinMsgBurst,omitempty"`

	// Enable or disable the sending of a negative acknowledgement on the discard of a guaranteed message. When a guaranteed message is published to a topic, and the router does not have any guaranteed subscriptions matching the message topic, the message can either be silently discarded, or a negative acknowledgement can be returned to the sender indicating that no guaranteed subscriptions were found for the message. It should be noted that even if the message is rejected to the publisher, it will still be delivered to any clients who have direct subscriptions to the topic. This configuration option does not affect the behavior of messages published to unknown queue names. That always results in the message being rejected to the sender. The default value is `false`. Available since 2.1.0.
	RejectMsgToSenderOnNoSubscriptionMatchEnabled bool `json:"rejectMsgToSenderOnNoSubscriptionMatchEnabled,omitempty"`

	// Enable or disable whether clients using this client profile are allowed to connect to the Message VPN if its replication is in standby state. The default value is `false`.
	ReplicationAllowClientConnectWhenStandbyEnabled bool `json:"replicationAllowClientConnectWhenStandbyEnabled,omitempty"`

	// The maximum number of SMF client connections that can be simultaneously connected with the same client-username. This value may be higher than supported by the hardware. The default is the max value supported by the hardware.
	ServiceSmfMaxConnectionCountPerClientUsername int64 `json:"serviceSmfMaxConnectionCountPerClientUsername,omitempty"`

	// The number of seconds during which the client must send a request or else the session is terminated. The default value is `30`.
	ServiceWebInactiveTimeout int64 `json:"serviceWebInactiveTimeout,omitempty"`

	// The maximum number of web-transport connections that can be simultaneously connected with the same client-username. This value may be higher than supported by the hardware. The default is the max value supported by the hardware.
	ServiceWebMaxConnectionCountPerClientUsername int64 `json:"serviceWebMaxConnectionCountPerClientUsername,omitempty"`

	// The maximum number of bytes allowed in a single web transport payload before fragmentation occurs, not including the header. The default value is `1000000`.
	ServiceWebMaxPayload int64 `json:"serviceWebMaxPayload,omitempty"`

	// The TCP initial congestion window size for clients belonging to this profile.   The initial congestion window size is used when starting up a TCP connection or recovery from idle (that is, no traffic). It is the number of segments TCP sends before waiting for an acknowledgement from the peer. Larger values of initial window allows a connection to come up to speed quickly. However, care must be taken for if this parameter's value is too high, it may cause congestion in the network. For further details on initial window, refer to RFC 2581. Changing this parameter changes all clients matching this profile, whether already connected or not.   Changing the initial window from its default of 2 results in non-compliance with RFC 2581. Contact Solace Support personnel before changing this parameter. The default value is `2`.
	TcpCongestionWindowSize int64 `json:"tcpCongestionWindowSize,omitempty"`

	// The number of keepalive probes TCP should send before dropping the connection. The default value is `5`.
	TcpKeepaliveCount int64 `json:"tcpKeepaliveCount,omitempty"`

	// The time (in seconds) a connection needs to remain idle before TCP begins sending keepalive probes. The default value is `3`.
	TcpKeepaliveIdleTime int64 `json:"tcpKeepaliveIdleTime,omitempty"`

	// The time between individual keepalive probes, when no response is received. The default value is `1`.
	TcpKeepaliveInterval int64 `json:"tcpKeepaliveInterval,omitempty"`

	// The TCP maximum segment size for the clients belonging to this profile. The default value is `1460`.
	TcpMaxSegmentSize int64 `json:"tcpMaxSegmentSize,omitempty"`

	// The TCP maximum window size (in KB) for clients belonging to this profile. Changes are applied to all existing connections. The maximum window should be at least the bandwidth-delay product of the link between the TCP peers. If the maximum window is less than the bandwidth-delay product, then the TCP connection operates below its maximum potential throughput. If the maximum window is less than about twice the bandwidth-delay product, then occasional packet loss causes TCP connection to operate below its maximum potential throughput as it handles the missing ACKs and retransmissions. There are also problems with a maximum window that's too large. In the presence of a high offered load, TCP gradually increases its congestion window until either (a) the congestion window reaches the maximum window, or (b) packet loss occurs in the network. Initially, when the congestion window is small, the network's physical bandwidth-delay acts as a memory buffer for packets in flight. As the congestion window crosses the bandwidth-delay product, though, the buffering of in-flight packets moves to queues in various switches, routers, etc. in the network. As the congestion window continues to increase, some such queue in some equipment overflows, causing packet loss and TCP back-off. The default value is `256`.
	TcpMaxWindowSize int64 `json:"tcpMaxWindowSize,omitempty"`
}
