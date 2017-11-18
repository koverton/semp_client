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

type MsgVpnJndiConnectionFactory struct {

	// Enable or disable whether new JMS connections can use the same Client identifier (ID) as an existing connection. The default value is `false`. Available since 2.3.0.
	AllowDuplicateClientIdEnabled bool `json:"allowDuplicateClientIdEnabled,omitempty"`

	// The description of the Client. The default value is `\"\"`.
	ClientDescription string `json:"clientDescription,omitempty"`

	// The Client identifier (ID). If not specified, a unique value for it will be generated. The default value is `\"\"`.
	ClientId string `json:"clientId,omitempty"`

	// The name of the JNDI Connection Factory.
	ConnectionFactoryName string `json:"connectionFactoryName,omitempty"`

	// Enable or disable overriding by the Subscriber (Consumer) of the deliver-to-one (DTO) property on messages received by it. The default value is `true`.
	DtoReceiveOverrideEnabled bool `json:"dtoReceiveOverrideEnabled,omitempty"`

	// The priority for receiving deliver-to-one (DTO) messages by the Subscriber (Consumer), if the messages are published on the Router that the Subscriber is directly connected to (messages are published locally). The default value is `1`.
	DtoReceiveSubscriberLocalPriority int32 `json:"dtoReceiveSubscriberLocalPriority,omitempty"`

	// The priority for receiving deliver-to-one (DTO) messages by the Subscriber (Consumer), if the messages are published on the Router other that the Subscriber is directly connected to (messages are published on a remote router). The default value is `1`.
	DtoReceiveSubscriberNetworkPriority int32 `json:"dtoReceiveSubscriberNetworkPriority,omitempty"`

	// Enable or disable the deliver-to-one (DTO) property on messages sent by the Publisher (Producer). The default value is `false`.
	DtoSendEnabled bool `json:"dtoSendEnabled,omitempty"`

	// Enable or disable whether a durable endpoint will be created on the Router when the corresponding \"Session.createDurableSubscriber()\" or \"Session.createQueue()\" is called. The created endpoint respects messages \"time-to-live\" (TTL) property according to the \"dynamicEndpointRespectTtlEnabled\" property value. The default value is `false`.
	DynamicEndpointCreateDurableEnabled bool `json:"dynamicEndpointCreateDurableEnabled,omitempty"`

	// Enable or disable whether dynamically created durable and non-durable endpoints respect messages \"time-to-live\" (TTL) property. The default value is `true`.
	DynamicEndpointRespectTtlEnabled bool `json:"dynamicEndpointRespectTtlEnabled,omitempty"`

	// The timeout for sending the acknowledgement (ACK) for guaranteed messages received by the Subscriber (Consumer), in milliseconds. The default value is `1000`.
	GuaranteedReceiveAckTimeout int32 `json:"guaranteedReceiveAckTimeout,omitempty"`

	// The size of the window for guaranteed messages received by the Subscriber (Consumer), in messages. The default value is `18`.
	GuaranteedReceiveWindowSize int32 `json:"guaranteedReceiveWindowSize,omitempty"`

	// The threshold for sending the acknowledgement (ACK) for guaranteed message receives by the Subscriber (Consumer) as percentage of its maximum value of \"guaranteedReceiveWindowSize\". The default value is `60`.
	GuaranteedReceiveWindowSizeAckThreshold int32 `json:"guaranteedReceiveWindowSizeAckThreshold,omitempty"`

	// The timeout for receiving the acknowledgement (ACK) for guaranteed messages sent by the Publisher (Producer), in milliseconds. The default value is `2000`.
	GuaranteedSendAckTimeout int32 `json:"guaranteedSendAckTimeout,omitempty"`

	// The size of the window for guaranteed messages sent by the Publisher (Producer), in messages. The default value is `255`.
	GuaranteedSendWindowSize int32 `json:"guaranteedSendWindowSize,omitempty"`

	// The default delivery mode for messages sent by the Publisher (Producer). The default value is `\"persistent\"`.
	MessagingDefaultDeliveryMode string `json:"messagingDefaultDeliveryMode,omitempty"`

	// Enable or disable whether messages sent by the Publisher (Producer) are Dead Message Queue (DMQ) eligible by default. The default value is `false`.
	MessagingDefaultDmqEligibleEnabled bool `json:"messagingDefaultDmqEligibleEnabled,omitempty"`

	// Enable or disable whether messages sent by the Publisher (Producer) are Eliding eligible by default. The default value is `false`.
	MessagingDefaultElidingEligibleEnabled bool `json:"messagingDefaultElidingEligibleEnabled,omitempty"`

	// Enable or disable adding or replacing of the JMSXUserID property in messages sent by the Publisher (Producer). The default value is `false`.
	MessagingJmsxUserIdEnabled bool `json:"messagingJmsxUserIdEnabled,omitempty"`

	// Enable or disable encoding of JMS text messages messages sent by the Publisher (Producer) in the message XML payload. When disabled it means that text messages are encoded in binary attachments. The default value is `true`.
	MessagingTextInXmlPayloadEnabled bool `json:"messagingTextInXmlPayloadEnabled,omitempty"`

	// The name of the Message VPN.
	MsgVpnName string `json:"msgVpnName,omitempty"`

	// The level of the ZLIB compression for the connection to the Router. \"0\" value means no compression, \"-1\" value means the compression level is specified in the JNDI Properties File. The default value is `-1`.
	TransportCompressionLevel int32 `json:"transportCompressionLevel,omitempty"`

	// The maximum number of attempts to establish an initial connection to the Router. \"0\" value means a single attempt, \"-1\" value means to retry forever. The default value is `0`.
	TransportConnectRetryCount int32 `json:"transportConnectRetryCount,omitempty"`

	// The maximum number of attempts to establish an initial connection to one host (Router) on the list of hosts (Routers). \"0\" value means a single attempt, \"-1\" value means to retry forever. The default value is `0`.
	TransportConnectRetryPerHostCount int32 `json:"transportConnectRetryPerHostCount,omitempty"`

	// The timeout for establishing an initial connection to the Router, in milliseconds. The default value is `30000`.
	TransportConnectTimeout int32 `json:"transportConnectTimeout,omitempty"`

	// Enable or disable usage of the Direct Transport Mode for sending non-persistent messages. Disabled value means that the Guaranteed Transport Mode is used. The default value is `true`.
	TransportDirectTransportEnabled bool `json:"transportDirectTransportEnabled,omitempty"`

	// The maximum number of consecutive application-level keepalive messages sent without the Router response before the connection to the Router is closed. The default value is `3`.
	TransportKeepaliveCount int32 `json:"transportKeepaliveCount,omitempty"`

	// Enable or disable usage of application-level keepalive messages to the Router. The default value is `true`.
	TransportKeepaliveEnabled bool `json:"transportKeepaliveEnabled,omitempty"`

	// The interval between application-level keepalive messages, in milliseconds. The default value is `3000`.
	TransportKeepaliveInterval int32 `json:"transportKeepaliveInterval,omitempty"`

	// Enable or disable delivery of asynchronous messages directly from the I/O thread. The default value is `false`.
	TransportMsgCallbackOnIoThreadEnabled bool `json:"transportMsgCallbackOnIoThreadEnabled,omitempty"`

	// Enable or disable optimization for the Direct Transport. If enabled, the application is limited by a single Subscriber (Consumer) and Publisher (Producer) per connection that, when combined with the topic dispatch turned off, improves latency when consuming messages. The default value is `false`.
	TransportOptimizeDirectEnabled bool `json:"transportOptimizeDirectEnabled,omitempty"`

	// The connection port number on the Router that is the port number for SMF clients. \"-1\" value means the port is specified in the JNDI Properties File. The default value is `-1`.
	TransportPort int32 `json:"transportPort,omitempty"`

	// The timeout for reading a reply from the Router, in milliseconds. The default value is `10000`.
	TransportReadTimeout int32 `json:"transportReadTimeout,omitempty"`

	// The size of the receive socket buffer, in bytes. It corresponds to the SO_RCVBUF socket option. The default value is `65536`.
	TransportReceiveBufferSize int32 `json:"transportReceiveBufferSize,omitempty"`

	// The maximum number of attempts to reestablish a connection to the Router after it has been lost. \"-1\" value means to retry forever. The default value is `3`.
	TransportReconnectRetryCount int32 `json:"transportReconnectRetryCount,omitempty"`

	// The amount of time before making another attempt to connect or reconnect to the Router after, in milliseconds. The default value is `3000`.
	TransportReconnectRetryWait int32 `json:"transportReconnectRetryWait,omitempty"`

	// The size of the send socket buffer, in bytes. It corresponds to the SO_SNDBUF socket option. The default value is `65536`.
	TransportSendBufferSize int32 `json:"transportSendBufferSize,omitempty"`

	// Enable or disable the TCP/IP Congestion Control (per RFC 896, the Nagles algorithm) known as \"the TCP_NODELAY option\". The default value is `true`.
	TransportTcpNoDelayEnabled bool `json:"transportTcpNoDelayEnabled,omitempty"`

	// Specify whether this is the XA Connection Factory. When enabled, the Connection Factory can be cast to \"XAConnectionFactory\", \"XAQueueConnectionFactory\" or \"XATopicConnectionFactory\". The default value is `false`.
	XaEnabled bool `json:"xaEnabled,omitempty"`
}
