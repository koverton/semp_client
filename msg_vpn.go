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

type MsgVpn struct {

	// Enable or disable Basic Authentication for clients connecting to the Message VPN. The default value is `true`.
	AuthenticationBasicEnabled bool `json:"authenticationBasicEnabled,omitempty"`

	// The name of the RADIUS or LDAP Profile to use when \"authenticationBasicType\" is \"radius\" or \"ldap\" respectively. The default value is `\"default\"`.
	AuthenticationBasicProfileName string `json:"authenticationBasicProfileName,omitempty"`

	// The RADIUS domain string to use when \"authenticationBasicType\" is \"radius\". The default value is `\"\"`.
	AuthenticationBasicRadiusDomain string `json:"authenticationBasicRadiusDomain,omitempty"`

	// Authentication mechanism to be used for Basic authentication of clients connecting to the Message VPN. The default value is `\"radius\"`. The allowed values and their meaning are:  <pre> \"radius\" - RADIUS authentication. A RADIUS profile name must be provided. \"ldap\" - LDAP authentication. An LDAP profile name must be provided. \"internal\" - Internal database. Authentication is against Client Usernames. \"none\" - No authentication. Anonymous login allowed. </pre> 
	AuthenticationBasicType string `json:"authenticationBasicType,omitempty"`

	// When enabled, if the client specifies a Client Username via the API connect method, the client provided Username is used instead of the CN (Common Name) field of the certificate\"s subject. When disabled, the certificate CN is always used as the Client Username. The default value is `false`.
	AuthenticationClientCertAllowApiProvidedUsernameEnabled bool `json:"authenticationClientCertAllowApiProvidedUsernameEnabled,omitempty"`

	// Enable or disable the Client Certificate client Authentication for the Message VPN. The default value is `false`.
	AuthenticationClientCertEnabled bool `json:"authenticationClientCertEnabled,omitempty"`

	// The maximum depth for the client certificate chain. The depth of the chain is defined as the number of signing CA certificates that are present in the chain back to the trusted self-signed root CA certificate. The default value is `3`.
	AuthenticationClientCertMaxChainDepth int64 `json:"authenticationClientCertMaxChainDepth,omitempty"`

	// Enable or disable validation of the \"Not Before\" and \"Not After\" validity dates in the client certificate. When disabled, a certificate will be accepted even if the certificate is not valid according to the \"Not Before\" and \"Not After\" validity dates in the certificate. The default value is `true`.
	AuthenticationClientCertValidateDateEnabled bool `json:"authenticationClientCertValidateDateEnabled,omitempty"`

	// When enabled, if the client specifies a Client Username via the API connect method, the client provided Username is used instead of the Kerberos Principal name in Kerberos token. When disabled, the Kerberos Principal name is always used as the Client Username. The default value is `false`.
	AuthenticationKerberosAllowApiProvidedUsernameEnabled bool `json:"authenticationKerberosAllowApiProvidedUsernameEnabled,omitempty"`

	// Enable or disable Kerberos Authentication for clients in the Message VPN. If a user provides credentials for a different authentication scheme, this setting is not applicable. The default value is `false`.
	AuthenticationKerberosEnabled bool `json:"authenticationKerberosEnabled,omitempty"`

	// The name of the attribute that should be retrieved from the LDAP server as part of the LDAP search when authorizing a client. It indicates that the client belongs to a particular group (i.e. the value associated with this attribute). The default value is `\"memberOf\"`.
	AuthorizationLdapGroupMembershipAttributeName string `json:"authorizationLdapGroupMembershipAttributeName,omitempty"`

	// The LDAP Profile name to be used when \"authorizationType\" is \"ldap\". The default value is `\"\"`.
	AuthorizationProfileName string `json:"authorizationProfileName,omitempty"`

	// Authorization mechanism to be used for clients connecting to the Message VPN. The default value is `\"internal\"`. The allowed values and their meaning are:  <pre> \"ldap\" - LDAP authorization. \"internal\" - Internal authorization. </pre> 
	AuthorizationType string `json:"authorizationType,omitempty"`

	// Enable or disable validation of the Common Name (CN) in the server certificate from the Remote Router. If enabled, the Common Name is checked against the list of Trusted Common Names configured for the Bridge. The default value is `true`.
	BridgingTlsServerCertEnforceTrustedCommonNameEnabled bool `json:"bridgingTlsServerCertEnforceTrustedCommonNameEnabled,omitempty"`

	// The maximum depth for the server certificate chain. The depth of the chain is defined as the number of signing CA certificates that are present in the chain back to the trusted self-signed root CA certificate. The default value is `3`.
	BridgingTlsServerCertMaxChainDepth int64 `json:"bridgingTlsServerCertMaxChainDepth,omitempty"`

	// Enable or disable validation of the \"Not Before\" and \"Not After\" validity dates in the server certificate. When disabled, a certificate will be accepted even if the certificate is not valid according to the \"Not Before\" and \"Not After\" validity dates in the certificate. The default value is `true`.
	BridgingTlsServerCertValidateDateEnabled bool `json:"bridgingTlsServerCertValidateDateEnabled,omitempty"`

	// Enable or disable managing of Cache Instances over the Message Bus. For a given Message VPN only one router in the network should have this attribute enabled. The default value is `true`.
	DistributedCacheManagementEnabled bool `json:"distributedCacheManagementEnabled,omitempty"`

	// Enable or disable the Message VPN. The default value is `false`.
	Enabled bool `json:"enabled,omitempty"`

	EventConnectionCountThreshold EventThreshold `json:"eventConnectionCountThreshold,omitempty"`

	EventEgressFlowCountThreshold EventThreshold `json:"eventEgressFlowCountThreshold,omitempty"`

	EventEgressMsgRateThreshold EventThresholdByValue `json:"eventEgressMsgRateThreshold,omitempty"`

	EventEndpointCountThreshold EventThreshold `json:"eventEndpointCountThreshold,omitempty"`

	EventIngressFlowCountThreshold EventThreshold `json:"eventIngressFlowCountThreshold,omitempty"`

	EventIngressMsgRateThreshold EventThresholdByValue `json:"eventIngressMsgRateThreshold,omitempty"`

	// Size in KB for what is being considered a large message for the Message VPN. The default value is `1024`.
	EventLargeMsgThreshold int64 `json:"eventLargeMsgThreshold,omitempty"`

	// A prefix applied to all published Events in the Message VPN. The default value is `\"\"`.
	EventLogTag string `json:"eventLogTag,omitempty"`

	EventMsgSpoolUsageThreshold EventThreshold `json:"eventMsgSpoolUsageThreshold,omitempty"`

	// Enable or disable Client level Event message publishing. The default value is `false`.
	EventPublishClientEnabled bool `json:"eventPublishClientEnabled,omitempty"`

	// Enable or disable Message VPN level Event message publishing. The default value is `false`.
	EventPublishMsgVpnEnabled bool `json:"eventPublishMsgVpnEnabled,omitempty"`

	// Subscription level Event message publishing mode. The default value is `\"off\"`. The allowed values and their meaning are:  <pre> \"off\" - Disable client level event message publishing. \"on-with-format-v1\" - Enable client level event message publishing with format v1. \"on-with-no-unsubscribe-events-on-disconnect-format-v1\" - As \"on-with-format-v1\", but unsubscribe events are not generated when a client disconnects. Unsubscribe events are still raised when a client explicitly unsubscribes from its subscriptions. \"on-with-format-v2\" - Enable client level event message publishing with format v2. \"on-with-no-unsubscribe-events-on-disconnect-format-v2\" - As \"on-with-format-v2\", but unsubscribe events are not generated when a client disconnects. Unsubscribe events are still raised when a client explicitly unsubscribes from its subscriptions. </pre> 
	EventPublishSubscriptionMode string `json:"eventPublishSubscriptionMode,omitempty"`

	// Enable or disable Event publish topics in MQTT format. The default value is `false`.
	EventPublishTopicFormatMqttEnabled bool `json:"eventPublishTopicFormatMqttEnabled,omitempty"`

	// Enable or disable Event publish topics in SMF format. The default value is `true`.
	EventPublishTopicFormatSmfEnabled bool `json:"eventPublishTopicFormatSmfEnabled,omitempty"`

	EventServiceMqttConnectionCountThreshold EventThreshold `json:"eventServiceMqttConnectionCountThreshold,omitempty"`

	EventServiceRestIncomingConnectionCountThreshold EventThreshold `json:"eventServiceRestIncomingConnectionCountThreshold,omitempty"`

	EventServiceSmfConnectionCountThreshold EventThreshold `json:"eventServiceSmfConnectionCountThreshold,omitempty"`

	EventServiceWebConnectionCountThreshold EventThreshold `json:"eventServiceWebConnectionCountThreshold,omitempty"`

	EventSubscriptionCountThreshold EventThreshold `json:"eventSubscriptionCountThreshold,omitempty"`

	EventTransactedSessionCountThreshold EventThreshold `json:"eventTransactedSessionCountThreshold,omitempty"`

	EventTransactionCountThreshold EventThreshold `json:"eventTransactionCountThreshold,omitempty"`

	// Enable or disable the export of subscriptions in the Message VPN to other routers in the network over Neighbor links. The default value is `false`.
	ExportSubscriptionsEnabled bool `json:"exportSubscriptionsEnabled,omitempty"`

	// Enable or disable JNDI access for clients in the Message VPN. The default value is `false`. Available since 2.2.0.
	JndiEnabled bool `json:"jndiEnabled,omitempty"`

	// The maximum number of client connections that can be simultaneously connected to the Message VPN. This value may be higher than supported by the hardware. The default is the maximum value supported by the hardware. The default is the max value supported by the hardware.
	MaxConnectionCount int64 `json:"maxConnectionCount,omitempty"`

	// The maximum number of Publisher (egress) flows that can be created on the Message VPN. The default value is `16000`.
	MaxEgressFlowCount int64 `json:"maxEgressFlowCount,omitempty"`

	// The maximum number of Queues and Topic Endpoints that can be created on the Message VPN. The default value is `16000`.
	MaxEndpointCount int64 `json:"maxEndpointCount,omitempty"`

	// The maximum number of Subscriber (ingress) flows that can be created on the Message VPN. The default value is `16000`.
	MaxIngressFlowCount int64 `json:"maxIngressFlowCount,omitempty"`

	// Max spool usage (in MB) allowed for the Message VPN. The default value is `0`.
	MaxMsgSpoolUsage int64 `json:"maxMsgSpoolUsage,omitempty"`

	// The maximum number of local client subscriptions (both primary and backup) that can be added to the Message VPN. The default varies by platform. The default varies by platform.
	MaxSubscriptionCount int64 `json:"maxSubscriptionCount,omitempty"`

	// The maximum number of transacted sessions for the Message VPN. The default varies by platform. The default varies by platform.
	MaxTransactedSessionCount int64 `json:"maxTransactedSessionCount,omitempty"`

	// The maximum number of transactions for the Message VPN. The default varies by platform. The default varies by platform.
	MaxTransactionCount int64 `json:"maxTransactionCount,omitempty"`

	// The name of the Message VPN.
	MsgVpnName string `json:"msgVpnName,omitempty"`

	// The acknowledgement (ACK) propagation interval for the Replication Bridge, in number of replicated messages. The default value is `20`.
	ReplicationAckPropagationIntervalMsgCount int64 `json:"replicationAckPropagationIntervalMsgCount,omitempty"`

	// The Client Username the Replication Bridge uses to login to the Remote Message VPN on the Replication mate. The default value is `\"\"`.
	ReplicationBridgeAuthenticationBasicClientUsername string `json:"replicationBridgeAuthenticationBasicClientUsername,omitempty"`

	// The password the Replication Bridge uses to login to the Remote Message VPN on the Replication mate. The default is to have no password. The default is to have no `replicationBridgeAuthenticationBasicPassword`.
	ReplicationBridgeAuthenticationBasicPassword string `json:"replicationBridgeAuthenticationBasicPassword,omitempty"`

	// The Authentication Scheme for the Replication Bridge. The default value is `\"basic\"`. The allowed values and their meaning are:  <pre> \"basic\" - Basic Authentication Scheme (via username and password). \"client-certificate\" - Client Certificate Authentication Scheme (via certificate-file). </pre> 
	ReplicationBridgeAuthenticationScheme string `json:"replicationBridgeAuthenticationScheme,omitempty"`

	// Whether compression is used for the Replication Bridge. The default value is `false`.
	ReplicationBridgeCompressedDataEnabled bool `json:"replicationBridgeCompressedDataEnabled,omitempty"`

	// The size of the window used for guaranteed messages on the Replication Bridge, in messages. The default value is `255`.
	ReplicationBridgeEgressFlowWindowSize int64 `json:"replicationBridgeEgressFlowWindowSize,omitempty"`

	// Number of seconds that must pass before retrying the Replication Bridge connection. The default value is `3`.
	ReplicationBridgeRetryDelay int64 `json:"replicationBridgeRetryDelay,omitempty"`

	// Enable or disable use of TLS for the Replication Bridge connection. The default value is `false`.
	ReplicationBridgeTlsEnabled bool `json:"replicationBridgeTlsEnabled,omitempty"`

	// The Client Profile for the Unidirectional Replication Bridge. The Client Profile must exist in the local Message VPN, and it is used only for the TCP parameters. The default value is `\"#client-profile\"`.
	ReplicationBridgeUnidirectionalClientProfileName string `json:"replicationBridgeUnidirectionalClientProfileName,omitempty"`

	// Enable or disable the Replication feature for the Message VPN. The default value is `false`.
	ReplicationEnabled bool `json:"replicationEnabled,omitempty"`

	// The behavior to take when enabling the Replication feature for the Message VPN, depending on the existence of the Replication Queue. The default value is `\"fail-on-existing-queue\"`. The allowed values and their meaning are:  <pre> \"fail-on-existing-queue\" - The data replication queue must not already exist. \"force-use-existing-queue\" - The data replication queue must already exist. Any data messages on the queue will be forwarded to interested applications. IMPORTANT: Before using this mode be certain that the messages are not stale or otherwise unsuitable to be forwarded. This mode can only be specified when the existing queue is configured the same as is currently specified under replication configuration otherwise the enabling of replication will fail. \"force-recreate-queue\" - The data replication queue must already exist. Any data messages on the queue will be discarded. IMPORTANT: Before using this mode be certain that the messages on the existing data replication queue are not needed by interested applications. </pre> 
	ReplicationEnabledQueueBehavior string `json:"replicationEnabledQueueBehavior,omitempty"`

	// The maximum amount of the Message Spool space that can be used by the Replication Bridge Queue, in megabytes (MB). The default value is `60000`.
	ReplicationQueueMaxMsgSpoolUsage int64 `json:"replicationQueueMaxMsgSpoolUsage,omitempty"`

	// Assign the message discard behavior, that is the circumstances under which a NACK is sent to the Client on the Replication Bridge Queue discards. The default value is `true`.
	ReplicationQueueRejectMsgToSenderOnDiscardEnabled bool `json:"replicationQueueRejectMsgToSenderOnDiscardEnabled,omitempty"`

	// Enable or disable the synchronously replicated topics ineligible behavior of the Replication Bridge. If enabled and the synchronous replication becomes ineligible, guaranteed messages published to synchronously replicated topics will be rejected back to the sender. If disabled, the synchronous replication will revert to the asynchronous one. The default value is `false`.
	ReplicationRejectMsgWhenSyncIneligibleEnabled bool `json:"replicationRejectMsgWhenSyncIneligibleEnabled,omitempty"`

	// The replication role for the Message VPN. The default value is `\"standby\"`. The allowed values and their meaning are:  <pre> \"active\" - Assume the Active role in Replication for the Message VPN. \"standby\" - Assume the Standby role in Replication for the Message VPN. </pre> 
	ReplicationRole string `json:"replicationRole,omitempty"`

	// The transaction replication mode for all transactions within the Message VPN. When mode is asynchronous, all transactions originated by clients are replicated to the standby site using the asynchronous replication. When mode is synchronous, all transactions originated by clients are replicated to the standby site using the synchronous replication. Changing this value during operation will not affect existing transactions, it is only used upon starting a transaction. The default value is `\"async\"`. The allowed values and their meaning are:  <pre> \"sync\" - Synchronous replication-mode. Published messages are acknowledged when they are spooled on the standby site. \"async\" - Asynchronous replication-mode. Published messages are acknowledged when they are spooled locally. </pre> 
	ReplicationTransactionMode string `json:"replicationTransactionMode,omitempty"`

	// Enable or disable validation of the Common Name (CN) in the server certificate from the remote REST Consumer. If enabled, the Common Name is checked against the list of Trusted Common Names configured for the REST Consumer. The default value is `true`.
	RestTlsServerCertEnforceTrustedCommonNameEnabled bool `json:"restTlsServerCertEnforceTrustedCommonNameEnabled,omitempty"`

	// The maximum depth for the server certificate from the remote REST Consumer chain. The depth of the chain is defined as the number of signing CA certificates that are present in the chain back to the trusted self-signed root CA certificate. The default value is `3`.
	RestTlsServerCertMaxChainDepth int64 `json:"restTlsServerCertMaxChainDepth,omitempty"`

	// Enable or disable validation of the \"Not Before\" and \"Not After\" validity dates in the server certificate from the remote REST Consumer. When disabled, a certificate will be accepted even if the certificate is not valid according to the \"Not Before\" and \"Not After\" validity dates in the certificate. The default value is `true`.
	RestTlsServerCertValidateDateEnabled bool `json:"restTlsServerCertValidateDateEnabled,omitempty"`

	// Enable or disable \"admin client\" SEMP over Message Bus for the current Message VPN. This applies only to SEMPv1. The default value is `false`.
	SempOverMsgBusAdminClientEnabled bool `json:"sempOverMsgBusAdminClientEnabled,omitempty"`

	// Enable or disable \"admin distributed-cache\" SEMP over Message Bus for the current Message VPN. This applies only to SEMPv1. The default value is `false`.
	SempOverMsgBusAdminDistributedCacheEnabled bool `json:"sempOverMsgBusAdminDistributedCacheEnabled,omitempty"`

	// Enable or disable \"admin\" SEMP over Message Bus for the current Message VPN. This applies only to SEMPv1. The default value is `false`.
	SempOverMsgBusAdminEnabled bool `json:"sempOverMsgBusAdminEnabled,omitempty"`

	// Enable or disable SEMP over Message Bus for the current Message VPN. This applies only to SEMPv1. The default value is `true`.
	SempOverMsgBusEnabled bool `json:"sempOverMsgBusEnabled,omitempty"`

	// Enable or disable \"show\" SEMP over Message Bus for the current Message VPN. This applies only to SEMPv1. The default value is `false`.
	SempOverMsgBusShowEnabled bool `json:"sempOverMsgBusShowEnabled,omitempty"`

	// The maximum number of MQTT client connections that can be simultaneously connected to the Message VPN. The default is the max value supported by the hardware. Available since 2.1.0.
	ServiceMqttMaxConnectionCount int64 `json:"serviceMqttMaxConnectionCount,omitempty"`

	// Enable or disable the plain-text MQTT service in the Message VPN. Disabling causes clients currently connected to be disconnected. The default value is `false`. Available since 2.1.0.
	ServiceMqttPlainTextEnabled bool `json:"serviceMqttPlainTextEnabled,omitempty"`

	// The port number for plain-text MQTT clients that connect to the Message VPN. The default value is `0`. Available since 2.1.0.
	ServiceMqttPlainTextListenPort int64 `json:"serviceMqttPlainTextListenPort,omitempty"`

	// Enable or disable the use of TLS for the MQTT service in the Message VPN. Disabling causes clients currently connected over TLS to be disconnected. The default value is `false`. Available since 2.1.0.
	ServiceMqttTlsEnabled bool `json:"serviceMqttTlsEnabled,omitempty"`

	// The port number for MQTT clients that connect to the Message VPN over TLS. The default value is `0`. Available since 2.1.0.
	ServiceMqttTlsListenPort int64 `json:"serviceMqttTlsListenPort,omitempty"`

	// Enable or disable the use of WebSocket over TLS for the MQTT service in the Message VPN. Disabling causes clients currently connected by WebSocket over TLS to be disconnected. The default value is `false`. Available since 2.1.0.
	ServiceMqttTlsWebSocketEnabled bool `json:"serviceMqttTlsWebSocketEnabled,omitempty"`

	// The port number for MQTT clients that connect to the Message VPN using WebSocket over TLS. The default value is `0`. Available since 2.1.0.
	ServiceMqttTlsWebSocketListenPort int64 `json:"serviceMqttTlsWebSocketListenPort,omitempty"`

	// Enable or disable the use of WebSocket for the MQTT service in the Message VPN. Disabling causes clients currently connected by WebSocket to be disconnected. The default value is `false`. Available since 2.1.0.
	ServiceMqttWebSocketEnabled bool `json:"serviceMqttWebSocketEnabled,omitempty"`

	// The port number for plain-text MQTT clients that connect to the Message VPN using WebSocket. The default value is `0`. Available since 2.1.0.
	ServiceMqttWebSocketListenPort int64 `json:"serviceMqttWebSocketListenPort,omitempty"`

	// The maximum number of REST incoming client connections that can be simultaneously connected to the Message VPN. The default is the max value supported by the hardware.
	ServiceRestIncomingMaxConnectionCount int64 `json:"serviceRestIncomingMaxConnectionCount,omitempty"`

	// Enable or disable the plain-text REST service for incoming clients in the Message VPN. Disabling causes clients currently connected to be disconnected. The default value is `false`.
	ServiceRestIncomingPlainTextEnabled bool `json:"serviceRestIncomingPlainTextEnabled,omitempty"`

	// The port number for incoming plain-text REST clients that connect to the Message VPN. The default value is `0`.
	ServiceRestIncomingPlainTextListenPort int64 `json:"serviceRestIncomingPlainTextListenPort,omitempty"`

	// Enable or disable the use of TLS for the REST service for incoming clients in the Message VPN. Disabling causes clients currently connected over TLS to be disconnected. The default value is `false`.
	ServiceRestIncomingTlsEnabled bool `json:"serviceRestIncomingTlsEnabled,omitempty"`

	// The port number for incoming REST clients that connect to the Message VPN over TLS. The default value is `0`.
	ServiceRestIncomingTlsListenPort int64 `json:"serviceRestIncomingTlsListenPort,omitempty"`

	// The maximum number of REST Consumer (outgoing) client connections that can be simultaneously connected to the Message VPN. The default varies by platform.
	ServiceRestOutgoingMaxConnectionCount int64 `json:"serviceRestOutgoingMaxConnectionCount,omitempty"`

	// The maximum number of SMF client connections that can be simultaneously connected to the Message VPN. The default is the max value supported by the hardware.
	ServiceSmfMaxConnectionCount int64 `json:"serviceSmfMaxConnectionCount,omitempty"`

	// Enable or disable the plain-text SMF service in the Message VPN. Disabling causes clients currently connected to be disconnected. The default value is `true`.
	ServiceSmfPlainTextEnabled bool `json:"serviceSmfPlainTextEnabled,omitempty"`

	// Enable or disable the use of TLS for the SMF service in the Message VPN. Disabling causes clients currently connected over TLS to be disconnected. The default value is `true`.
	ServiceSmfTlsEnabled bool `json:"serviceSmfTlsEnabled,omitempty"`

	// The maximum number of Web Transport client connections that can be simultaneously connected to the Message VPN. The default is the max value supported by the hardware.
	ServiceWebMaxConnectionCount int64 `json:"serviceWebMaxConnectionCount,omitempty"`

	// Enable or disable the plain-text Web Transport service in the Message VPN. Disabling causes clients currently connected to be disconnected. The default value is `true`.
	ServiceWebPlainTextEnabled bool `json:"serviceWebPlainTextEnabled,omitempty"`

	// Enable or disable the use of TLS for the Web Transport service in the Message VPN. Disabling causes clients currently connected over TLS to be disconnected. The default value is `true`.
	ServiceWebTlsEnabled bool `json:"serviceWebTlsEnabled,omitempty"`

	// Enable or disable the allowing of TLS SMF clients to downgrade their connections to plain-text connections. Changing this will not affect existing connections. The default value is `false`.
	TlsAllowDowngradeToPlainTextEnabled bool `json:"tlsAllowDowngradeToPlainTextEnabled,omitempty"`
}
