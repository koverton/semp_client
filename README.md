# Go API client for semp_client

SEMP (starting in `v2`, see [note 1](#notes)) is a RESTful API for configuring, monitoring, and administering a Solace router.  SEMP uses URIs to address manageable **resources** of the Solace router.  Resources are either individual **objects**, or **collections** of objects.  This document applies to the following API:   API|Base Path|Purpose|Comments :---|:---|:---|:--- Configuration|/SEMP/v2/config|Reading and writing config state|See [note 2](#notes)    Resources are always nouns, with individual objects being singular and  collections being plural. Objects within a collection are identified by an  `obj-id`, which follows the collection name with the form  `collection-name/obj-id`. Some examples:  <pre> /SEMP/v2/config/msgVpns                       ; MsgVpn collection /SEMP/v2/config/msgVpns/finance               ; MsgVpn object named \"finance\" /SEMP/v2/config/msgVpns/finance/queues        ; Queue collection within MsgVpn \"finance\" /SEMP/v2/config/msgVpns/finance/queues/orderQ ; Queue object named \"orderQ\" within MsgVpn \"finance\" </pre>  ## Collection Resources  Collections are unordered lists of objects (unless described as otherwise), and  are described by JSON arrays. Each item in the array represents an object in  the same manner as the individual object would normally be represented. The creation of a new object is done through its collection  resource.   ## Object Resources  Objects are composed of attributes and collections, and are described by JSON  content as name/value pairs. The collections of an object are not contained  directly in the object's JSON content, rather the content includes a URI  attribute which points to the collection. This contained collection resource  must be managed as a separate resource through this URI.  At a minimum, every object has 1 or more identifying attributes, and its own  `uri` attribute which contains the URI to itself. Attributes may have any  (non-exclusively) of the following properties:   Property|Meaning|Comments :---|:---|:--- Identifying|Attribute is involved in unique identification of the object, and appears in its URI| Required|Attribute must be provided in the request| Read-Only|Attribute can only be read, not written|See [note 3](#notes) Write-Only|Attribute can only be written, not read| Requires-Disable|Attribute can only be changed when object is disabled| Deprecated|Attribute is deprecated, and will disappear in the next SEMP version|    In some requests, certain attributes may only be provided in  certain combinations with other attributes:   Relationship|Meaning :---|:--- Requires|Attribute may only be changed by a request if a particular attribute or combination of attributes is also provided in the request Conflicts|Attribute may only be provided in a request if a particular attribute or combination of attributes is not also provided in the request     ## HTTP Methods  The following HTTP methods manipulate resources in accordance with these  general principles:   Method|Resource|Meaning|Request Body|Response Body|Missing Request Attributes :---|:---|:---|:---|:---|:--- POST|Collection|Create object|Initial attribute values|Object attributes and metadata|Set to default PUT|Object|Create or replace object|New attribute values|Object attributes and metadata|Set to default (but see [note 4](#notes)) PATCH|Object|Update object|New attribute values|Object attributes and metadata|unchanged DELETE|Object|Delete object|Empty|Object metadata|N/A GET|Object|Get object|Empty|Object attributes and metadata|N/A GET|Collection|Get collection|Empty|Object attributes and collection metadata|N/A    ## Common Query Parameters  The following are some common query parameters that are supported by many  method/URI combinations. Individual URIs may document additional parameters.  Note that multiple query parameters can be used together in a single URI,  separated by the ampersand character. For example:  <pre> ; Request for the MsgVpns collection using two hypothetical query parameters ; \"q1\" and \"q2\" with values \"val1\" and \"val2\" respectively /SEMP/v2/config/msgVpns?q1=val1&q2=val2 </pre>  ### select  Include in the response only selected attributes of the object. Use this query  parameter to limit the size of the returned data for each returned object, or  return only those fields that are desired.  The value of `select` is a comma-separated list of attribute names. Names may  include the `*` wildcard (zero or more characters). Nested attribute names  are supported using periods (e.g. `parentName.childName`). If the list is  empty (i.e. `select=`) no attributes are returned; otherwise the list must  match at least one attribute name of the object. Some examples:  <pre> ; List of all MsgVpn names /SEMP/v2/config/msgVpns?select=msgVpnName  ; Authentication attributes of MsgVpn \"finance\" /SEMP/v2/config/msgVpns/finance?select=authentication*  ; Access related attributes of Queue \"orderQ\" of MsgVpn \"finance\" /SEMP/v2/config/msgVpns/finance/queues/orderQ?select=owner,permission </pre>  ### where  Include in the response only objects where certain conditions are true. Use  this query parameter to limit which objects are returned to those whose  attribute values meet the given conditions.  The value of `where` is a comma-separated list of expressions. All expressions  must be true for the object to be included in the response. Each expression  takes the form:  <pre> expression  = attribute-name OP value OP          = '==' | '!=' | '&lt;' | '&gt;' | '&lt;=' | '&gt;=' </pre>  `value` may be a number, string, `true`, or `false`, as appropriate for the  type of `attribute-name`. Greater-than and less-than comparisons only work for  numbers. A `*` in a string `value` is interpreted as a wildcard (zero or more  characters). Some examples:  <pre> ; Only enabled MsgVpns /SEMP/v2/config/msgVpns?where=enabled==true  ; Only MsgVpns using basic non-LDAP authentication /SEMP/v2/config/msgVpns?where=authenticationBasicEnabled==true,authenticationBasicType!=ldap  ; Only MsgVpns that allow more than 100 client connections /SEMP/v2/config/msgVpns?where=maxConnectionCount&gt;100  ; Only MsgVpns with msgVpnName starting with \"B\": /SEMP/v2/config/msgVpns?where=msgVpnName==B* </pre>  ### count  Limit the count of objects in the response. This can be useful to limit the  size of the response for large collections. The minimum value for `count` is  `1` and the default is `10`. There is a hidden maximum  as to prevent overloading the system. For example:  <pre> ; Up to 25 MsgVpns /SEMP/v2/config/msgVpns?count=25 </pre>  ### cursor  The cursor, or position, for the next page of objects. Cursors are opaque data  that should not be created or interpreted by SEMP clients, and should only be  used as described below.  When a request is made for a collection and there may be additional objects  available for retrieval that are not included in the initial response, the  response will include a `cursorQuery` field containing a cursor. The value  of this field can be specified in the `cursor` query parameter of a  subsequent request to retrieve the next page of objects. For convenience,  an appropriate URI is constructed automatically by the router and included  in the `nextPageUri` field of the response. This URI can be used directly  to retrieve the next page of objects.  ## Notes  Note|Description :---:|:--- 1|This specification defines SEMP starting in \"v2\", and not the original SEMP \"v1\" interface. Request and response formats between \"v1\" and \"v2\" are entirely incompatible, although both protocols share a common port configuration on the Solace router. They are differentiated by the initial portion of the URI path, one of either \"/SEMP/\" or \"/SEMP/v2/\" 2|This API is partially implemented. Only a subset of all objects are available. 3|Read-only attributes may appear in POST and PUT/PATCH requests. However, if a read-only attribute is not marked as identifying, it will be ignored during a PUT/PATCH. 4|For PUT, if the SEMP user is not authorized to modify the attribute, its value is left unchanged rather than set to default. In addition, the values of write-only attributes are not set to their defaults on a PUT. If the object does not exist, it is created first. 5|For DELETE, the body of the request currently serves no purpose and will cause an error if not empty.    

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 2.3.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen
For more information, please visit [http://www.solace.com](http://www.solace.com)

## Installation
Put the package under your project folder and add the following in import:
```
    "./semp_client"
```

## Documentation for API Endpoints

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AboutApi* | [**GetAboutApi**](docs/AboutApi.md#getaboutapi) | **Get** /about/api | Gets an API Description object.
*AboutApi* | [**GetAboutUser**](docs/AboutApi.md#getaboutuser) | **Get** /about/user | Gets a Current User object.
*AboutApi* | [**GetAboutUserMsgVpn**](docs/AboutApi.md#getaboutusermsgvpn) | **Get** /about/user/msgVpns/{msgVpnName} | Gets a Current User Message VPN object.
*AboutApi* | [**GetAboutUserMsgVpns**](docs/AboutApi.md#getaboutusermsgvpns) | **Get** /about/user/msgVpns | Gets a list of Current User Message VPN objects.
*AclProfileApi* | [**CreateMsgVpnAclProfile**](docs/AclProfileApi.md#createmsgvpnaclprofile) | **Post** /msgVpns/{msgVpnName}/aclProfiles | Creates an ACL Profile object.
*AclProfileApi* | [**CreateMsgVpnAclProfileClientConnectException**](docs/AclProfileApi.md#createmsgvpnaclprofileclientconnectexception) | **Post** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions | Creates a Client Connect Exception object.
*AclProfileApi* | [**CreateMsgVpnAclProfilePublishException**](docs/AclProfileApi.md#createmsgvpnaclprofilepublishexception) | **Post** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions | Creates a Publish Topic Exception object.
*AclProfileApi* | [**CreateMsgVpnAclProfileSubscribeException**](docs/AclProfileApi.md#createmsgvpnaclprofilesubscribeexception) | **Post** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions | Creates a Subscribe Topic Exception object.
*AclProfileApi* | [**DeleteMsgVpnAclProfile**](docs/AclProfileApi.md#deletemsgvpnaclprofile) | **Delete** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName} | Deletes an ACL Profile object.
*AclProfileApi* | [**DeleteMsgVpnAclProfileClientConnectException**](docs/AclProfileApi.md#deletemsgvpnaclprofileclientconnectexception) | **Delete** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions/{clientConnectExceptionAddress} | Deletes a Client Connect Exception object.
*AclProfileApi* | [**DeleteMsgVpnAclProfilePublishException**](docs/AclProfileApi.md#deletemsgvpnaclprofilepublishexception) | **Delete** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions/{topicSyntax},{publishExceptionTopic} | Deletes a Publish Topic Exception object.
*AclProfileApi* | [**DeleteMsgVpnAclProfileSubscribeException**](docs/AclProfileApi.md#deletemsgvpnaclprofilesubscribeexception) | **Delete** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions/{topicSyntax},{subscribeExceptionTopic} | Deletes a Subscribe Topic Exception object.
*AclProfileApi* | [**GetMsgVpnAclProfile**](docs/AclProfileApi.md#getmsgvpnaclprofile) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName} | Gets an ACL Profile object.
*AclProfileApi* | [**GetMsgVpnAclProfileClientConnectException**](docs/AclProfileApi.md#getmsgvpnaclprofileclientconnectexception) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions/{clientConnectExceptionAddress} | Gets a Client Connect Exception object.
*AclProfileApi* | [**GetMsgVpnAclProfileClientConnectExceptions**](docs/AclProfileApi.md#getmsgvpnaclprofileclientconnectexceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions | Gets a list of Client Connect Exception objects.
*AclProfileApi* | [**GetMsgVpnAclProfilePublishException**](docs/AclProfileApi.md#getmsgvpnaclprofilepublishexception) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions/{topicSyntax},{publishExceptionTopic} | Gets a Publish Topic Exception object.
*AclProfileApi* | [**GetMsgVpnAclProfilePublishExceptions**](docs/AclProfileApi.md#getmsgvpnaclprofilepublishexceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions | Gets a list of Publish Topic Exception objects.
*AclProfileApi* | [**GetMsgVpnAclProfileSubscribeException**](docs/AclProfileApi.md#getmsgvpnaclprofilesubscribeexception) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions/{topicSyntax},{subscribeExceptionTopic} | Gets a Subscribe Topic Exception object.
*AclProfileApi* | [**GetMsgVpnAclProfileSubscribeExceptions**](docs/AclProfileApi.md#getmsgvpnaclprofilesubscribeexceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions | Gets a list of Subscribe Topic Exception objects.
*AclProfileApi* | [**GetMsgVpnAclProfiles**](docs/AclProfileApi.md#getmsgvpnaclprofiles) | **Get** /msgVpns/{msgVpnName}/aclProfiles | Gets a list of ACL Profile objects.
*AclProfileApi* | [**ReplaceMsgVpnAclProfile**](docs/AclProfileApi.md#replacemsgvpnaclprofile) | **Put** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName} | Replaces an ACL Profile object.
*AclProfileApi* | [**UpdateMsgVpnAclProfile**](docs/AclProfileApi.md#updatemsgvpnaclprofile) | **Patch** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName} | Updates an ACL Profile object.
*AuthorizationGroupApi* | [**CreateMsgVpnAuthorizationGroup**](docs/AuthorizationGroupApi.md#createmsgvpnauthorizationgroup) | **Post** /msgVpns/{msgVpnName}/authorizationGroups | Creates an LDAP Authorization Group object.
*AuthorizationGroupApi* | [**DeleteMsgVpnAuthorizationGroup**](docs/AuthorizationGroupApi.md#deletemsgvpnauthorizationgroup) | **Delete** /msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName} | Deletes an LDAP Authorization Group object.
*AuthorizationGroupApi* | [**GetMsgVpnAuthorizationGroup**](docs/AuthorizationGroupApi.md#getmsgvpnauthorizationgroup) | **Get** /msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName} | Gets an LDAP Authorization Group object.
*AuthorizationGroupApi* | [**GetMsgVpnAuthorizationGroups**](docs/AuthorizationGroupApi.md#getmsgvpnauthorizationgroups) | **Get** /msgVpns/{msgVpnName}/authorizationGroups | Gets a list of LDAP Authorization Group objects.
*AuthorizationGroupApi* | [**ReplaceMsgVpnAuthorizationGroup**](docs/AuthorizationGroupApi.md#replacemsgvpnauthorizationgroup) | **Put** /msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName} | Replaces an LDAP Authorization Group object.
*AuthorizationGroupApi* | [**UpdateMsgVpnAuthorizationGroup**](docs/AuthorizationGroupApi.md#updatemsgvpnauthorizationgroup) | **Patch** /msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName} | Updates an LDAP Authorization Group object.
*BridgeApi* | [**CreateMsgVpnBridge**](docs/BridgeApi.md#createmsgvpnbridge) | **Post** /msgVpns/{msgVpnName}/bridges | Creates a Bridge object.
*BridgeApi* | [**CreateMsgVpnBridgeRemoteMsgVpn**](docs/BridgeApi.md#createmsgvpnbridgeremotemsgvpn) | **Post** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns | Creates a Remote Message VPN object.
*BridgeApi* | [**CreateMsgVpnBridgeRemoteSubscription**](docs/BridgeApi.md#createmsgvpnbridgeremotesubscription) | **Post** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions | Creates a Remote Subscription object.
*BridgeApi* | [**CreateMsgVpnBridgeTlsTrustedCommonName**](docs/BridgeApi.md#createmsgvpnbridgetlstrustedcommonname) | **Post** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames | Creates a Trusted Common Name object.
*BridgeApi* | [**DeleteMsgVpnBridge**](docs/BridgeApi.md#deletemsgvpnbridge) | **Delete** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Deletes a Bridge object.
*BridgeApi* | [**DeleteMsgVpnBridgeRemoteMsgVpn**](docs/BridgeApi.md#deletemsgvpnbridgeremotemsgvpn) | **Delete** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface} | Deletes a Remote Message VPN object.
*BridgeApi* | [**DeleteMsgVpnBridgeRemoteSubscription**](docs/BridgeApi.md#deletemsgvpnbridgeremotesubscription) | **Delete** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions/{remoteSubscriptionTopic} | Deletes a Remote Subscription object.
*BridgeApi* | [**DeleteMsgVpnBridgeTlsTrustedCommonName**](docs/BridgeApi.md#deletemsgvpnbridgetlstrustedcommonname) | **Delete** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Deletes a Trusted Common Name object.
*BridgeApi* | [**GetMsgVpnBridge**](docs/BridgeApi.md#getmsgvpnbridge) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Gets a Bridge object.
*BridgeApi* | [**GetMsgVpnBridgeRemoteMsgVpn**](docs/BridgeApi.md#getmsgvpnbridgeremotemsgvpn) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface} | Gets a Remote Message VPN object.
*BridgeApi* | [**GetMsgVpnBridgeRemoteMsgVpns**](docs/BridgeApi.md#getmsgvpnbridgeremotemsgvpns) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns | Gets a list of Remote Message VPN objects.
*BridgeApi* | [**GetMsgVpnBridgeRemoteSubscription**](docs/BridgeApi.md#getmsgvpnbridgeremotesubscription) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions/{remoteSubscriptionTopic} | Gets a Remote Subscription object.
*BridgeApi* | [**GetMsgVpnBridgeRemoteSubscriptions**](docs/BridgeApi.md#getmsgvpnbridgeremotesubscriptions) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions | Gets a list of Remote Subscription objects.
*BridgeApi* | [**GetMsgVpnBridgeTlsTrustedCommonName**](docs/BridgeApi.md#getmsgvpnbridgetlstrustedcommonname) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Gets a Trusted Common Name object.
*BridgeApi* | [**GetMsgVpnBridgeTlsTrustedCommonNames**](docs/BridgeApi.md#getmsgvpnbridgetlstrustedcommonnames) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames | Gets a list of Trusted Common Name objects.
*BridgeApi* | [**GetMsgVpnBridges**](docs/BridgeApi.md#getmsgvpnbridges) | **Get** /msgVpns/{msgVpnName}/bridges | Gets a list of Bridge objects.
*BridgeApi* | [**ReplaceMsgVpnBridge**](docs/BridgeApi.md#replacemsgvpnbridge) | **Put** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Replaces a Bridge object.
*BridgeApi* | [**ReplaceMsgVpnBridgeRemoteMsgVpn**](docs/BridgeApi.md#replacemsgvpnbridgeremotemsgvpn) | **Put** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface} | Replaces a Remote Message VPN object.
*BridgeApi* | [**UpdateMsgVpnBridge**](docs/BridgeApi.md#updatemsgvpnbridge) | **Patch** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Updates a Bridge object.
*BridgeApi* | [**UpdateMsgVpnBridgeRemoteMsgVpn**](docs/BridgeApi.md#updatemsgvpnbridgeremotemsgvpn) | **Patch** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface} | Updates a Remote Message VPN object.
*ClientProfileApi* | [**CreateMsgVpnClientProfile**](docs/ClientProfileApi.md#createmsgvpnclientprofile) | **Post** /msgVpns/{msgVpnName}/clientProfiles | Creates a Client Profile object.
*ClientProfileApi* | [**DeleteMsgVpnClientProfile**](docs/ClientProfileApi.md#deletemsgvpnclientprofile) | **Delete** /msgVpns/{msgVpnName}/clientProfiles/{clientProfileName} | Deletes a Client Profile object.
*ClientProfileApi* | [**GetMsgVpnClientProfile**](docs/ClientProfileApi.md#getmsgvpnclientprofile) | **Get** /msgVpns/{msgVpnName}/clientProfiles/{clientProfileName} | Gets a Client Profile object.
*ClientProfileApi* | [**GetMsgVpnClientProfiles**](docs/ClientProfileApi.md#getmsgvpnclientprofiles) | **Get** /msgVpns/{msgVpnName}/clientProfiles | Gets a list of Client Profile objects.
*ClientProfileApi* | [**ReplaceMsgVpnClientProfile**](docs/ClientProfileApi.md#replacemsgvpnclientprofile) | **Put** /msgVpns/{msgVpnName}/clientProfiles/{clientProfileName} | Replaces a Client Profile object.
*ClientProfileApi* | [**UpdateMsgVpnClientProfile**](docs/ClientProfileApi.md#updatemsgvpnclientprofile) | **Patch** /msgVpns/{msgVpnName}/clientProfiles/{clientProfileName} | Updates a Client Profile object.
*ClientUsernameApi* | [**CreateMsgVpnClientUsername**](docs/ClientUsernameApi.md#createmsgvpnclientusername) | **Post** /msgVpns/{msgVpnName}/clientUsernames | Creates a Client Username object.
*ClientUsernameApi* | [**DeleteMsgVpnClientUsername**](docs/ClientUsernameApi.md#deletemsgvpnclientusername) | **Delete** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername} | Deletes a Client Username object.
*ClientUsernameApi* | [**GetMsgVpnClientUsername**](docs/ClientUsernameApi.md#getmsgvpnclientusername) | **Get** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername} | Gets a Client Username object.
*ClientUsernameApi* | [**GetMsgVpnClientUsernames**](docs/ClientUsernameApi.md#getmsgvpnclientusernames) | **Get** /msgVpns/{msgVpnName}/clientUsernames | Gets a list of Client Username objects.
*ClientUsernameApi* | [**ReplaceMsgVpnClientUsername**](docs/ClientUsernameApi.md#replacemsgvpnclientusername) | **Put** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername} | Replaces a Client Username object.
*ClientUsernameApi* | [**UpdateMsgVpnClientUsername**](docs/ClientUsernameApi.md#updatemsgvpnclientusername) | **Patch** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername} | Updates a Client Username object.
*JndiApi* | [**CreateMsgVpnJndiConnectionFactory**](docs/JndiApi.md#createmsgvpnjndiconnectionfactory) | **Post** /msgVpns/{msgVpnName}/jndiConnectionFactories | Creates a JNDI Connection Factory object.
*JndiApi* | [**CreateMsgVpnJndiQueue**](docs/JndiApi.md#createmsgvpnjndiqueue) | **Post** /msgVpns/{msgVpnName}/jndiQueues | Creates a JNDI Queue object.
*JndiApi* | [**CreateMsgVpnJndiTopic**](docs/JndiApi.md#createmsgvpnjnditopic) | **Post** /msgVpns/{msgVpnName}/jndiTopics | Creates a JNDI Topic object.
*JndiApi* | [**DeleteMsgVpnJndiConnectionFactory**](docs/JndiApi.md#deletemsgvpnjndiconnectionfactory) | **Delete** /msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName} | Deletes a JNDI Connection Factory object.
*JndiApi* | [**DeleteMsgVpnJndiQueue**](docs/JndiApi.md#deletemsgvpnjndiqueue) | **Delete** /msgVpns/{msgVpnName}/jndiQueues/{queueName} | Deletes a JNDI Queue object.
*JndiApi* | [**DeleteMsgVpnJndiTopic**](docs/JndiApi.md#deletemsgvpnjnditopic) | **Delete** /msgVpns/{msgVpnName}/jndiTopics/{topicName} | Deletes a JNDI Topic object.
*JndiApi* | [**GetMsgVpnJndiConnectionFactories**](docs/JndiApi.md#getmsgvpnjndiconnectionfactories) | **Get** /msgVpns/{msgVpnName}/jndiConnectionFactories | Gets a list of JNDI Connection Factory objects.
*JndiApi* | [**GetMsgVpnJndiConnectionFactory**](docs/JndiApi.md#getmsgvpnjndiconnectionfactory) | **Get** /msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName} | Gets a JNDI Connection Factory object.
*JndiApi* | [**GetMsgVpnJndiQueue**](docs/JndiApi.md#getmsgvpnjndiqueue) | **Get** /msgVpns/{msgVpnName}/jndiQueues/{queueName} | Gets a JNDI Queue object.
*JndiApi* | [**GetMsgVpnJndiQueues**](docs/JndiApi.md#getmsgvpnjndiqueues) | **Get** /msgVpns/{msgVpnName}/jndiQueues | Gets a list of JNDI Queue objects.
*JndiApi* | [**GetMsgVpnJndiTopic**](docs/JndiApi.md#getmsgvpnjnditopic) | **Get** /msgVpns/{msgVpnName}/jndiTopics/{topicName} | Gets a JNDI Topic object.
*JndiApi* | [**GetMsgVpnJndiTopics**](docs/JndiApi.md#getmsgvpnjnditopics) | **Get** /msgVpns/{msgVpnName}/jndiTopics | Gets a list of JNDI Topic objects.
*JndiApi* | [**ReplaceMsgVpnJndiConnectionFactory**](docs/JndiApi.md#replacemsgvpnjndiconnectionfactory) | **Put** /msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName} | Replaces a JNDI Connection Factory object.
*JndiApi* | [**ReplaceMsgVpnJndiQueue**](docs/JndiApi.md#replacemsgvpnjndiqueue) | **Put** /msgVpns/{msgVpnName}/jndiQueues/{queueName} | Replaces a JNDI Queue object.
*JndiApi* | [**ReplaceMsgVpnJndiTopic**](docs/JndiApi.md#replacemsgvpnjnditopic) | **Put** /msgVpns/{msgVpnName}/jndiTopics/{topicName} | Replaces a JNDI Topic object.
*JndiApi* | [**UpdateMsgVpnJndiConnectionFactory**](docs/JndiApi.md#updatemsgvpnjndiconnectionfactory) | **Patch** /msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName} | Updates a JNDI Connection Factory object.
*JndiApi* | [**UpdateMsgVpnJndiQueue**](docs/JndiApi.md#updatemsgvpnjndiqueue) | **Patch** /msgVpns/{msgVpnName}/jndiQueues/{queueName} | Updates a JNDI Queue object.
*JndiApi* | [**UpdateMsgVpnJndiTopic**](docs/JndiApi.md#updatemsgvpnjnditopic) | **Patch** /msgVpns/{msgVpnName}/jndiTopics/{topicName} | Updates a JNDI Topic object.
*MqttSessionApi* | [**CreateMsgVpnMqttSession**](docs/MqttSessionApi.md#createmsgvpnmqttsession) | **Post** /msgVpns/{msgVpnName}/mqttSessions | Creates an MQTT Session object.
*MqttSessionApi* | [**CreateMsgVpnMqttSessionSubscription**](docs/MqttSessionApi.md#createmsgvpnmqttsessionsubscription) | **Post** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions | Creates an MQTT Session Subscription object.
*MqttSessionApi* | [**DeleteMsgVpnMqttSession**](docs/MqttSessionApi.md#deletemsgvpnmqttsession) | **Delete** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter} | Deletes an MQTT Session object.
*MqttSessionApi* | [**DeleteMsgVpnMqttSessionSubscription**](docs/MqttSessionApi.md#deletemsgvpnmqttsessionsubscription) | **Delete** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic} | Deletes an MQTT Session Subscription object.
*MqttSessionApi* | [**GetMsgVpnMqttSession**](docs/MqttSessionApi.md#getmsgvpnmqttsession) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter} | Gets an MQTT Session object.
*MqttSessionApi* | [**GetMsgVpnMqttSessionSubscription**](docs/MqttSessionApi.md#getmsgvpnmqttsessionsubscription) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic} | Gets an MQTT Session Subscription object.
*MqttSessionApi* | [**GetMsgVpnMqttSessionSubscriptions**](docs/MqttSessionApi.md#getmsgvpnmqttsessionsubscriptions) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions | Gets a list of MQTT Session Subscription objects.
*MqttSessionApi* | [**GetMsgVpnMqttSessions**](docs/MqttSessionApi.md#getmsgvpnmqttsessions) | **Get** /msgVpns/{msgVpnName}/mqttSessions | Gets a list of MQTT Session objects.
*MqttSessionApi* | [**ReplaceMsgVpnMqttSession**](docs/MqttSessionApi.md#replacemsgvpnmqttsession) | **Put** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter} | Replaces an MQTT Session object.
*MqttSessionApi* | [**ReplaceMsgVpnMqttSessionSubscription**](docs/MqttSessionApi.md#replacemsgvpnmqttsessionsubscription) | **Put** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic} | Replaces an MQTT Session Subscription object.
*MqttSessionApi* | [**UpdateMsgVpnMqttSession**](docs/MqttSessionApi.md#updatemsgvpnmqttsession) | **Patch** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter} | Updates an MQTT Session object.
*MqttSessionApi* | [**UpdateMsgVpnMqttSessionSubscription**](docs/MqttSessionApi.md#updatemsgvpnmqttsessionsubscription) | **Patch** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic} | Updates an MQTT Session Subscription object.
*MsgVpnApi* | [**CreateMsgVpn**](docs/MsgVpnApi.md#createmsgvpn) | **Post** /msgVpns | Creates a Message VPN object.
*MsgVpnApi* | [**CreateMsgVpnAclProfile**](docs/MsgVpnApi.md#createmsgvpnaclprofile) | **Post** /msgVpns/{msgVpnName}/aclProfiles | Creates an ACL Profile object.
*MsgVpnApi* | [**CreateMsgVpnAclProfileClientConnectException**](docs/MsgVpnApi.md#createmsgvpnaclprofileclientconnectexception) | **Post** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions | Creates a Client Connect Exception object.
*MsgVpnApi* | [**CreateMsgVpnAclProfilePublishException**](docs/MsgVpnApi.md#createmsgvpnaclprofilepublishexception) | **Post** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions | Creates a Publish Topic Exception object.
*MsgVpnApi* | [**CreateMsgVpnAclProfileSubscribeException**](docs/MsgVpnApi.md#createmsgvpnaclprofilesubscribeexception) | **Post** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions | Creates a Subscribe Topic Exception object.
*MsgVpnApi* | [**CreateMsgVpnAuthorizationGroup**](docs/MsgVpnApi.md#createmsgvpnauthorizationgroup) | **Post** /msgVpns/{msgVpnName}/authorizationGroups | Creates an LDAP Authorization Group object.
*MsgVpnApi* | [**CreateMsgVpnBridge**](docs/MsgVpnApi.md#createmsgvpnbridge) | **Post** /msgVpns/{msgVpnName}/bridges | Creates a Bridge object.
*MsgVpnApi* | [**CreateMsgVpnBridgeRemoteMsgVpn**](docs/MsgVpnApi.md#createmsgvpnbridgeremotemsgvpn) | **Post** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns | Creates a Remote Message VPN object.
*MsgVpnApi* | [**CreateMsgVpnBridgeRemoteSubscription**](docs/MsgVpnApi.md#createmsgvpnbridgeremotesubscription) | **Post** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions | Creates a Remote Subscription object.
*MsgVpnApi* | [**CreateMsgVpnBridgeTlsTrustedCommonName**](docs/MsgVpnApi.md#createmsgvpnbridgetlstrustedcommonname) | **Post** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames | Creates a Trusted Common Name object.
*MsgVpnApi* | [**CreateMsgVpnClientProfile**](docs/MsgVpnApi.md#createmsgvpnclientprofile) | **Post** /msgVpns/{msgVpnName}/clientProfiles | Creates a Client Profile object.
*MsgVpnApi* | [**CreateMsgVpnClientUsername**](docs/MsgVpnApi.md#createmsgvpnclientusername) | **Post** /msgVpns/{msgVpnName}/clientUsernames | Creates a Client Username object.
*MsgVpnApi* | [**CreateMsgVpnJndiConnectionFactory**](docs/MsgVpnApi.md#createmsgvpnjndiconnectionfactory) | **Post** /msgVpns/{msgVpnName}/jndiConnectionFactories | Creates a JNDI Connection Factory object.
*MsgVpnApi* | [**CreateMsgVpnJndiQueue**](docs/MsgVpnApi.md#createmsgvpnjndiqueue) | **Post** /msgVpns/{msgVpnName}/jndiQueues | Creates a JNDI Queue object.
*MsgVpnApi* | [**CreateMsgVpnJndiTopic**](docs/MsgVpnApi.md#createmsgvpnjnditopic) | **Post** /msgVpns/{msgVpnName}/jndiTopics | Creates a JNDI Topic object.
*MsgVpnApi* | [**CreateMsgVpnMqttSession**](docs/MsgVpnApi.md#createmsgvpnmqttsession) | **Post** /msgVpns/{msgVpnName}/mqttSessions | Creates an MQTT Session object.
*MsgVpnApi* | [**CreateMsgVpnMqttSessionSubscription**](docs/MsgVpnApi.md#createmsgvpnmqttsessionsubscription) | **Post** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions | Creates an MQTT Session Subscription object.
*MsgVpnApi* | [**CreateMsgVpnQueue**](docs/MsgVpnApi.md#createmsgvpnqueue) | **Post** /msgVpns/{msgVpnName}/queues | Creates a Queue object.
*MsgVpnApi* | [**CreateMsgVpnQueueSubscription**](docs/MsgVpnApi.md#createmsgvpnqueuesubscription) | **Post** /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions | Creates a Queue Subscription object.
*MsgVpnApi* | [**CreateMsgVpnReplicatedTopic**](docs/MsgVpnApi.md#createmsgvpnreplicatedtopic) | **Post** /msgVpns/{msgVpnName}/replicatedTopics | Creates a Replicated Topic object.
*MsgVpnApi* | [**CreateMsgVpnRestDeliveryPoint**](docs/MsgVpnApi.md#createmsgvpnrestdeliverypoint) | **Post** /msgVpns/{msgVpnName}/restDeliveryPoints | Creates a REST Delivery Point object.
*MsgVpnApi* | [**CreateMsgVpnRestDeliveryPointQueueBinding**](docs/MsgVpnApi.md#createmsgvpnrestdeliverypointqueuebinding) | **Post** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings | Creates a Queue Binding object.
*MsgVpnApi* | [**CreateMsgVpnRestDeliveryPointRestConsumer**](docs/MsgVpnApi.md#createmsgvpnrestdeliverypointrestconsumer) | **Post** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers | Creates a REST Consumer object.
*MsgVpnApi* | [**CreateMsgVpnSequencedTopic**](docs/MsgVpnApi.md#createmsgvpnsequencedtopic) | **Post** /msgVpns/{msgVpnName}/sequencedTopics | Creates a Sequenced Topic object.
*MsgVpnApi* | [**CreateMsgVpnTopicEndpoint**](docs/MsgVpnApi.md#createmsgvpntopicendpoint) | **Post** /msgVpns/{msgVpnName}/topicEndpoints | Creates a Topic Endpoint object.
*MsgVpnApi* | [**DeleteMsgVpn**](docs/MsgVpnApi.md#deletemsgvpn) | **Delete** /msgVpns/{msgVpnName} | Deletes a Message VPN object.
*MsgVpnApi* | [**DeleteMsgVpnAclProfile**](docs/MsgVpnApi.md#deletemsgvpnaclprofile) | **Delete** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName} | Deletes an ACL Profile object.
*MsgVpnApi* | [**DeleteMsgVpnAclProfileClientConnectException**](docs/MsgVpnApi.md#deletemsgvpnaclprofileclientconnectexception) | **Delete** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions/{clientConnectExceptionAddress} | Deletes a Client Connect Exception object.
*MsgVpnApi* | [**DeleteMsgVpnAclProfilePublishException**](docs/MsgVpnApi.md#deletemsgvpnaclprofilepublishexception) | **Delete** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions/{topicSyntax},{publishExceptionTopic} | Deletes a Publish Topic Exception object.
*MsgVpnApi* | [**DeleteMsgVpnAclProfileSubscribeException**](docs/MsgVpnApi.md#deletemsgvpnaclprofilesubscribeexception) | **Delete** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions/{topicSyntax},{subscribeExceptionTopic} | Deletes a Subscribe Topic Exception object.
*MsgVpnApi* | [**DeleteMsgVpnAuthorizationGroup**](docs/MsgVpnApi.md#deletemsgvpnauthorizationgroup) | **Delete** /msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName} | Deletes an LDAP Authorization Group object.
*MsgVpnApi* | [**DeleteMsgVpnBridge**](docs/MsgVpnApi.md#deletemsgvpnbridge) | **Delete** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Deletes a Bridge object.
*MsgVpnApi* | [**DeleteMsgVpnBridgeRemoteMsgVpn**](docs/MsgVpnApi.md#deletemsgvpnbridgeremotemsgvpn) | **Delete** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface} | Deletes a Remote Message VPN object.
*MsgVpnApi* | [**DeleteMsgVpnBridgeRemoteSubscription**](docs/MsgVpnApi.md#deletemsgvpnbridgeremotesubscription) | **Delete** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions/{remoteSubscriptionTopic} | Deletes a Remote Subscription object.
*MsgVpnApi* | [**DeleteMsgVpnBridgeTlsTrustedCommonName**](docs/MsgVpnApi.md#deletemsgvpnbridgetlstrustedcommonname) | **Delete** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Deletes a Trusted Common Name object.
*MsgVpnApi* | [**DeleteMsgVpnClientProfile**](docs/MsgVpnApi.md#deletemsgvpnclientprofile) | **Delete** /msgVpns/{msgVpnName}/clientProfiles/{clientProfileName} | Deletes a Client Profile object.
*MsgVpnApi* | [**DeleteMsgVpnClientUsername**](docs/MsgVpnApi.md#deletemsgvpnclientusername) | **Delete** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername} | Deletes a Client Username object.
*MsgVpnApi* | [**DeleteMsgVpnJndiConnectionFactory**](docs/MsgVpnApi.md#deletemsgvpnjndiconnectionfactory) | **Delete** /msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName} | Deletes a JNDI Connection Factory object.
*MsgVpnApi* | [**DeleteMsgVpnJndiQueue**](docs/MsgVpnApi.md#deletemsgvpnjndiqueue) | **Delete** /msgVpns/{msgVpnName}/jndiQueues/{queueName} | Deletes a JNDI Queue object.
*MsgVpnApi* | [**DeleteMsgVpnJndiTopic**](docs/MsgVpnApi.md#deletemsgvpnjnditopic) | **Delete** /msgVpns/{msgVpnName}/jndiTopics/{topicName} | Deletes a JNDI Topic object.
*MsgVpnApi* | [**DeleteMsgVpnMqttSession**](docs/MsgVpnApi.md#deletemsgvpnmqttsession) | **Delete** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter} | Deletes an MQTT Session object.
*MsgVpnApi* | [**DeleteMsgVpnMqttSessionSubscription**](docs/MsgVpnApi.md#deletemsgvpnmqttsessionsubscription) | **Delete** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic} | Deletes an MQTT Session Subscription object.
*MsgVpnApi* | [**DeleteMsgVpnQueue**](docs/MsgVpnApi.md#deletemsgvpnqueue) | **Delete** /msgVpns/{msgVpnName}/queues/{queueName} | Deletes a Queue object.
*MsgVpnApi* | [**DeleteMsgVpnQueueSubscription**](docs/MsgVpnApi.md#deletemsgvpnqueuesubscription) | **Delete** /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions/{subscriptionTopic} | Deletes a Queue Subscription object.
*MsgVpnApi* | [**DeleteMsgVpnReplicatedTopic**](docs/MsgVpnApi.md#deletemsgvpnreplicatedtopic) | **Delete** /msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic} | Deletes a Replicated Topic object.
*MsgVpnApi* | [**DeleteMsgVpnRestDeliveryPoint**](docs/MsgVpnApi.md#deletemsgvpnrestdeliverypoint) | **Delete** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName} | Deletes a REST Delivery Point object.
*MsgVpnApi* | [**DeleteMsgVpnRestDeliveryPointQueueBinding**](docs/MsgVpnApi.md#deletemsgvpnrestdeliverypointqueuebinding) | **Delete** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName} | Deletes a Queue Binding object.
*MsgVpnApi* | [**DeleteMsgVpnRestDeliveryPointRestConsumer**](docs/MsgVpnApi.md#deletemsgvpnrestdeliverypointrestconsumer) | **Delete** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName} | Deletes a REST Consumer object.
*MsgVpnApi* | [**DeleteMsgVpnSequencedTopic**](docs/MsgVpnApi.md#deletemsgvpnsequencedtopic) | **Delete** /msgVpns/{msgVpnName}/sequencedTopics/{sequencedTopic} | Deletes a Sequenced Topic object.
*MsgVpnApi* | [**DeleteMsgVpnTopicEndpoint**](docs/MsgVpnApi.md#deletemsgvpntopicendpoint) | **Delete** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName} | Deletes a Topic Endpoint object.
*MsgVpnApi* | [**GetMsgVpn**](docs/MsgVpnApi.md#getmsgvpn) | **Get** /msgVpns/{msgVpnName} | Gets a Message VPN object.
*MsgVpnApi* | [**GetMsgVpnAclProfile**](docs/MsgVpnApi.md#getmsgvpnaclprofile) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName} | Gets an ACL Profile object.
*MsgVpnApi* | [**GetMsgVpnAclProfileClientConnectException**](docs/MsgVpnApi.md#getmsgvpnaclprofileclientconnectexception) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions/{clientConnectExceptionAddress} | Gets a Client Connect Exception object.
*MsgVpnApi* | [**GetMsgVpnAclProfileClientConnectExceptions**](docs/MsgVpnApi.md#getmsgvpnaclprofileclientconnectexceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions | Gets a list of Client Connect Exception objects.
*MsgVpnApi* | [**GetMsgVpnAclProfilePublishException**](docs/MsgVpnApi.md#getmsgvpnaclprofilepublishexception) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions/{topicSyntax},{publishExceptionTopic} | Gets a Publish Topic Exception object.
*MsgVpnApi* | [**GetMsgVpnAclProfilePublishExceptions**](docs/MsgVpnApi.md#getmsgvpnaclprofilepublishexceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions | Gets a list of Publish Topic Exception objects.
*MsgVpnApi* | [**GetMsgVpnAclProfileSubscribeException**](docs/MsgVpnApi.md#getmsgvpnaclprofilesubscribeexception) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions/{topicSyntax},{subscribeExceptionTopic} | Gets a Subscribe Topic Exception object.
*MsgVpnApi* | [**GetMsgVpnAclProfileSubscribeExceptions**](docs/MsgVpnApi.md#getmsgvpnaclprofilesubscribeexceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions | Gets a list of Subscribe Topic Exception objects.
*MsgVpnApi* | [**GetMsgVpnAclProfiles**](docs/MsgVpnApi.md#getmsgvpnaclprofiles) | **Get** /msgVpns/{msgVpnName}/aclProfiles | Gets a list of ACL Profile objects.
*MsgVpnApi* | [**GetMsgVpnAuthorizationGroup**](docs/MsgVpnApi.md#getmsgvpnauthorizationgroup) | **Get** /msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName} | Gets an LDAP Authorization Group object.
*MsgVpnApi* | [**GetMsgVpnAuthorizationGroups**](docs/MsgVpnApi.md#getmsgvpnauthorizationgroups) | **Get** /msgVpns/{msgVpnName}/authorizationGroups | Gets a list of LDAP Authorization Group objects.
*MsgVpnApi* | [**GetMsgVpnBridge**](docs/MsgVpnApi.md#getmsgvpnbridge) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Gets a Bridge object.
*MsgVpnApi* | [**GetMsgVpnBridgeRemoteMsgVpn**](docs/MsgVpnApi.md#getmsgvpnbridgeremotemsgvpn) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface} | Gets a Remote Message VPN object.
*MsgVpnApi* | [**GetMsgVpnBridgeRemoteMsgVpns**](docs/MsgVpnApi.md#getmsgvpnbridgeremotemsgvpns) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns | Gets a list of Remote Message VPN objects.
*MsgVpnApi* | [**GetMsgVpnBridgeRemoteSubscription**](docs/MsgVpnApi.md#getmsgvpnbridgeremotesubscription) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions/{remoteSubscriptionTopic} | Gets a Remote Subscription object.
*MsgVpnApi* | [**GetMsgVpnBridgeRemoteSubscriptions**](docs/MsgVpnApi.md#getmsgvpnbridgeremotesubscriptions) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions | Gets a list of Remote Subscription objects.
*MsgVpnApi* | [**GetMsgVpnBridgeTlsTrustedCommonName**](docs/MsgVpnApi.md#getmsgvpnbridgetlstrustedcommonname) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Gets a Trusted Common Name object.
*MsgVpnApi* | [**GetMsgVpnBridgeTlsTrustedCommonNames**](docs/MsgVpnApi.md#getmsgvpnbridgetlstrustedcommonnames) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames | Gets a list of Trusted Common Name objects.
*MsgVpnApi* | [**GetMsgVpnBridges**](docs/MsgVpnApi.md#getmsgvpnbridges) | **Get** /msgVpns/{msgVpnName}/bridges | Gets a list of Bridge objects.
*MsgVpnApi* | [**GetMsgVpnClientProfile**](docs/MsgVpnApi.md#getmsgvpnclientprofile) | **Get** /msgVpns/{msgVpnName}/clientProfiles/{clientProfileName} | Gets a Client Profile object.
*MsgVpnApi* | [**GetMsgVpnClientProfiles**](docs/MsgVpnApi.md#getmsgvpnclientprofiles) | **Get** /msgVpns/{msgVpnName}/clientProfiles | Gets a list of Client Profile objects.
*MsgVpnApi* | [**GetMsgVpnClientUsername**](docs/MsgVpnApi.md#getmsgvpnclientusername) | **Get** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername} | Gets a Client Username object.
*MsgVpnApi* | [**GetMsgVpnClientUsernames**](docs/MsgVpnApi.md#getmsgvpnclientusernames) | **Get** /msgVpns/{msgVpnName}/clientUsernames | Gets a list of Client Username objects.
*MsgVpnApi* | [**GetMsgVpnJndiConnectionFactories**](docs/MsgVpnApi.md#getmsgvpnjndiconnectionfactories) | **Get** /msgVpns/{msgVpnName}/jndiConnectionFactories | Gets a list of JNDI Connection Factory objects.
*MsgVpnApi* | [**GetMsgVpnJndiConnectionFactory**](docs/MsgVpnApi.md#getmsgvpnjndiconnectionfactory) | **Get** /msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName} | Gets a JNDI Connection Factory object.
*MsgVpnApi* | [**GetMsgVpnJndiQueue**](docs/MsgVpnApi.md#getmsgvpnjndiqueue) | **Get** /msgVpns/{msgVpnName}/jndiQueues/{queueName} | Gets a JNDI Queue object.
*MsgVpnApi* | [**GetMsgVpnJndiQueues**](docs/MsgVpnApi.md#getmsgvpnjndiqueues) | **Get** /msgVpns/{msgVpnName}/jndiQueues | Gets a list of JNDI Queue objects.
*MsgVpnApi* | [**GetMsgVpnJndiTopic**](docs/MsgVpnApi.md#getmsgvpnjnditopic) | **Get** /msgVpns/{msgVpnName}/jndiTopics/{topicName} | Gets a JNDI Topic object.
*MsgVpnApi* | [**GetMsgVpnJndiTopics**](docs/MsgVpnApi.md#getmsgvpnjnditopics) | **Get** /msgVpns/{msgVpnName}/jndiTopics | Gets a list of JNDI Topic objects.
*MsgVpnApi* | [**GetMsgVpnMqttSession**](docs/MsgVpnApi.md#getmsgvpnmqttsession) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter} | Gets an MQTT Session object.
*MsgVpnApi* | [**GetMsgVpnMqttSessionSubscription**](docs/MsgVpnApi.md#getmsgvpnmqttsessionsubscription) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic} | Gets an MQTT Session Subscription object.
*MsgVpnApi* | [**GetMsgVpnMqttSessionSubscriptions**](docs/MsgVpnApi.md#getmsgvpnmqttsessionsubscriptions) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions | Gets a list of MQTT Session Subscription objects.
*MsgVpnApi* | [**GetMsgVpnMqttSessions**](docs/MsgVpnApi.md#getmsgvpnmqttsessions) | **Get** /msgVpns/{msgVpnName}/mqttSessions | Gets a list of MQTT Session objects.
*MsgVpnApi* | [**GetMsgVpnQueue**](docs/MsgVpnApi.md#getmsgvpnqueue) | **Get** /msgVpns/{msgVpnName}/queues/{queueName} | Gets a Queue object.
*MsgVpnApi* | [**GetMsgVpnQueueSubscription**](docs/MsgVpnApi.md#getmsgvpnqueuesubscription) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions/{subscriptionTopic} | Gets a Queue Subscription object.
*MsgVpnApi* | [**GetMsgVpnQueueSubscriptions**](docs/MsgVpnApi.md#getmsgvpnqueuesubscriptions) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions | Gets a list of Queue Subscription objects.
*MsgVpnApi* | [**GetMsgVpnQueues**](docs/MsgVpnApi.md#getmsgvpnqueues) | **Get** /msgVpns/{msgVpnName}/queues | Gets a list of Queue objects.
*MsgVpnApi* | [**GetMsgVpnReplicatedTopic**](docs/MsgVpnApi.md#getmsgvpnreplicatedtopic) | **Get** /msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic} | Gets a Replicated Topic object.
*MsgVpnApi* | [**GetMsgVpnReplicatedTopics**](docs/MsgVpnApi.md#getmsgvpnreplicatedtopics) | **Get** /msgVpns/{msgVpnName}/replicatedTopics | Gets a list of Replicated Topic objects.
*MsgVpnApi* | [**GetMsgVpnRestDeliveryPoint**](docs/MsgVpnApi.md#getmsgvpnrestdeliverypoint) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName} | Gets a REST Delivery Point object.
*MsgVpnApi* | [**GetMsgVpnRestDeliveryPointQueueBinding**](docs/MsgVpnApi.md#getmsgvpnrestdeliverypointqueuebinding) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName} | Gets a Queue Binding object.
*MsgVpnApi* | [**GetMsgVpnRestDeliveryPointQueueBindings**](docs/MsgVpnApi.md#getmsgvpnrestdeliverypointqueuebindings) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings | Gets a list of Queue Binding objects.
*MsgVpnApi* | [**GetMsgVpnRestDeliveryPointRestConsumer**](docs/MsgVpnApi.md#getmsgvpnrestdeliverypointrestconsumer) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName} | Gets a REST Consumer object.
*MsgVpnApi* | [**GetMsgVpnRestDeliveryPointRestConsumers**](docs/MsgVpnApi.md#getmsgvpnrestdeliverypointrestconsumers) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers | Gets a list of REST Consumer objects.
*MsgVpnApi* | [**GetMsgVpnRestDeliveryPoints**](docs/MsgVpnApi.md#getmsgvpnrestdeliverypoints) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints | Gets a list of REST Delivery Point objects.
*MsgVpnApi* | [**GetMsgVpnSequencedTopic**](docs/MsgVpnApi.md#getmsgvpnsequencedtopic) | **Get** /msgVpns/{msgVpnName}/sequencedTopics/{sequencedTopic} | Gets a Sequenced Topic object.
*MsgVpnApi* | [**GetMsgVpnSequencedTopics**](docs/MsgVpnApi.md#getmsgvpnsequencedtopics) | **Get** /msgVpns/{msgVpnName}/sequencedTopics | Gets a list of Sequenced Topic objects.
*MsgVpnApi* | [**GetMsgVpnTopicEndpoint**](docs/MsgVpnApi.md#getmsgvpntopicendpoint) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName} | Gets a Topic Endpoint object.
*MsgVpnApi* | [**GetMsgVpnTopicEndpoints**](docs/MsgVpnApi.md#getmsgvpntopicendpoints) | **Get** /msgVpns/{msgVpnName}/topicEndpoints | Gets a list of Topic Endpoint objects.
*MsgVpnApi* | [**GetMsgVpns**](docs/MsgVpnApi.md#getmsgvpns) | **Get** /msgVpns | Gets a list of Message VPN objects.
*MsgVpnApi* | [**ReplaceMsgVpn**](docs/MsgVpnApi.md#replacemsgvpn) | **Put** /msgVpns/{msgVpnName} | Replaces a Message VPN object.
*MsgVpnApi* | [**ReplaceMsgVpnAclProfile**](docs/MsgVpnApi.md#replacemsgvpnaclprofile) | **Put** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName} | Replaces an ACL Profile object.
*MsgVpnApi* | [**ReplaceMsgVpnAuthorizationGroup**](docs/MsgVpnApi.md#replacemsgvpnauthorizationgroup) | **Put** /msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName} | Replaces an LDAP Authorization Group object.
*MsgVpnApi* | [**ReplaceMsgVpnBridge**](docs/MsgVpnApi.md#replacemsgvpnbridge) | **Put** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Replaces a Bridge object.
*MsgVpnApi* | [**ReplaceMsgVpnBridgeRemoteMsgVpn**](docs/MsgVpnApi.md#replacemsgvpnbridgeremotemsgvpn) | **Put** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface} | Replaces a Remote Message VPN object.
*MsgVpnApi* | [**ReplaceMsgVpnClientProfile**](docs/MsgVpnApi.md#replacemsgvpnclientprofile) | **Put** /msgVpns/{msgVpnName}/clientProfiles/{clientProfileName} | Replaces a Client Profile object.
*MsgVpnApi* | [**ReplaceMsgVpnClientUsername**](docs/MsgVpnApi.md#replacemsgvpnclientusername) | **Put** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername} | Replaces a Client Username object.
*MsgVpnApi* | [**ReplaceMsgVpnJndiConnectionFactory**](docs/MsgVpnApi.md#replacemsgvpnjndiconnectionfactory) | **Put** /msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName} | Replaces a JNDI Connection Factory object.
*MsgVpnApi* | [**ReplaceMsgVpnJndiQueue**](docs/MsgVpnApi.md#replacemsgvpnjndiqueue) | **Put** /msgVpns/{msgVpnName}/jndiQueues/{queueName} | Replaces a JNDI Queue object.
*MsgVpnApi* | [**ReplaceMsgVpnJndiTopic**](docs/MsgVpnApi.md#replacemsgvpnjnditopic) | **Put** /msgVpns/{msgVpnName}/jndiTopics/{topicName} | Replaces a JNDI Topic object.
*MsgVpnApi* | [**ReplaceMsgVpnMqttSession**](docs/MsgVpnApi.md#replacemsgvpnmqttsession) | **Put** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter} | Replaces an MQTT Session object.
*MsgVpnApi* | [**ReplaceMsgVpnMqttSessionSubscription**](docs/MsgVpnApi.md#replacemsgvpnmqttsessionsubscription) | **Put** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic} | Replaces an MQTT Session Subscription object.
*MsgVpnApi* | [**ReplaceMsgVpnQueue**](docs/MsgVpnApi.md#replacemsgvpnqueue) | **Put** /msgVpns/{msgVpnName}/queues/{queueName} | Replaces a Queue object.
*MsgVpnApi* | [**ReplaceMsgVpnReplicatedTopic**](docs/MsgVpnApi.md#replacemsgvpnreplicatedtopic) | **Put** /msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic} | Replaces a Replicated Topic object.
*MsgVpnApi* | [**ReplaceMsgVpnRestDeliveryPoint**](docs/MsgVpnApi.md#replacemsgvpnrestdeliverypoint) | **Put** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName} | Replaces a REST Delivery Point object.
*MsgVpnApi* | [**ReplaceMsgVpnRestDeliveryPointQueueBinding**](docs/MsgVpnApi.md#replacemsgvpnrestdeliverypointqueuebinding) | **Put** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName} | Replaces a Queue Binding object.
*MsgVpnApi* | [**ReplaceMsgVpnRestDeliveryPointRestConsumer**](docs/MsgVpnApi.md#replacemsgvpnrestdeliverypointrestconsumer) | **Put** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName} | Replaces a REST Consumer object.
*MsgVpnApi* | [**ReplaceMsgVpnTopicEndpoint**](docs/MsgVpnApi.md#replacemsgvpntopicendpoint) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName} | Replaces a Topic Endpoint object.
*MsgVpnApi* | [**UpdateMsgVpn**](docs/MsgVpnApi.md#updatemsgvpn) | **Patch** /msgVpns/{msgVpnName} | Updates a Message VPN object.
*MsgVpnApi* | [**UpdateMsgVpnAclProfile**](docs/MsgVpnApi.md#updatemsgvpnaclprofile) | **Patch** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName} | Updates an ACL Profile object.
*MsgVpnApi* | [**UpdateMsgVpnAuthorizationGroup**](docs/MsgVpnApi.md#updatemsgvpnauthorizationgroup) | **Patch** /msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName} | Updates an LDAP Authorization Group object.
*MsgVpnApi* | [**UpdateMsgVpnBridge**](docs/MsgVpnApi.md#updatemsgvpnbridge) | **Patch** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Updates a Bridge object.
*MsgVpnApi* | [**UpdateMsgVpnBridgeRemoteMsgVpn**](docs/MsgVpnApi.md#updatemsgvpnbridgeremotemsgvpn) | **Patch** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface} | Updates a Remote Message VPN object.
*MsgVpnApi* | [**UpdateMsgVpnClientProfile**](docs/MsgVpnApi.md#updatemsgvpnclientprofile) | **Patch** /msgVpns/{msgVpnName}/clientProfiles/{clientProfileName} | Updates a Client Profile object.
*MsgVpnApi* | [**UpdateMsgVpnClientUsername**](docs/MsgVpnApi.md#updatemsgvpnclientusername) | **Patch** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername} | Updates a Client Username object.
*MsgVpnApi* | [**UpdateMsgVpnJndiConnectionFactory**](docs/MsgVpnApi.md#updatemsgvpnjndiconnectionfactory) | **Patch** /msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName} | Updates a JNDI Connection Factory object.
*MsgVpnApi* | [**UpdateMsgVpnJndiQueue**](docs/MsgVpnApi.md#updatemsgvpnjndiqueue) | **Patch** /msgVpns/{msgVpnName}/jndiQueues/{queueName} | Updates a JNDI Queue object.
*MsgVpnApi* | [**UpdateMsgVpnJndiTopic**](docs/MsgVpnApi.md#updatemsgvpnjnditopic) | **Patch** /msgVpns/{msgVpnName}/jndiTopics/{topicName} | Updates a JNDI Topic object.
*MsgVpnApi* | [**UpdateMsgVpnMqttSession**](docs/MsgVpnApi.md#updatemsgvpnmqttsession) | **Patch** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter} | Updates an MQTT Session object.
*MsgVpnApi* | [**UpdateMsgVpnMqttSessionSubscription**](docs/MsgVpnApi.md#updatemsgvpnmqttsessionsubscription) | **Patch** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic} | Updates an MQTT Session Subscription object.
*MsgVpnApi* | [**UpdateMsgVpnQueue**](docs/MsgVpnApi.md#updatemsgvpnqueue) | **Patch** /msgVpns/{msgVpnName}/queues/{queueName} | Updates a Queue object.
*MsgVpnApi* | [**UpdateMsgVpnReplicatedTopic**](docs/MsgVpnApi.md#updatemsgvpnreplicatedtopic) | **Patch** /msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic} | Updates a Replicated Topic object.
*MsgVpnApi* | [**UpdateMsgVpnRestDeliveryPoint**](docs/MsgVpnApi.md#updatemsgvpnrestdeliverypoint) | **Patch** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName} | Updates a REST Delivery Point object.
*MsgVpnApi* | [**UpdateMsgVpnRestDeliveryPointQueueBinding**](docs/MsgVpnApi.md#updatemsgvpnrestdeliverypointqueuebinding) | **Patch** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName} | Updates a Queue Binding object.
*MsgVpnApi* | [**UpdateMsgVpnRestDeliveryPointRestConsumer**](docs/MsgVpnApi.md#updatemsgvpnrestdeliverypointrestconsumer) | **Patch** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName} | Updates a REST Consumer object.
*MsgVpnApi* | [**UpdateMsgVpnTopicEndpoint**](docs/MsgVpnApi.md#updatemsgvpntopicendpoint) | **Patch** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName} | Updates a Topic Endpoint object.
*QueueApi* | [**CreateMsgVpnQueue**](docs/QueueApi.md#createmsgvpnqueue) | **Post** /msgVpns/{msgVpnName}/queues | Creates a Queue object.
*QueueApi* | [**CreateMsgVpnQueueSubscription**](docs/QueueApi.md#createmsgvpnqueuesubscription) | **Post** /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions | Creates a Queue Subscription object.
*QueueApi* | [**DeleteMsgVpnQueue**](docs/QueueApi.md#deletemsgvpnqueue) | **Delete** /msgVpns/{msgVpnName}/queues/{queueName} | Deletes a Queue object.
*QueueApi* | [**DeleteMsgVpnQueueSubscription**](docs/QueueApi.md#deletemsgvpnqueuesubscription) | **Delete** /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions/{subscriptionTopic} | Deletes a Queue Subscription object.
*QueueApi* | [**GetMsgVpnQueue**](docs/QueueApi.md#getmsgvpnqueue) | **Get** /msgVpns/{msgVpnName}/queues/{queueName} | Gets a Queue object.
*QueueApi* | [**GetMsgVpnQueueSubscription**](docs/QueueApi.md#getmsgvpnqueuesubscription) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions/{subscriptionTopic} | Gets a Queue Subscription object.
*QueueApi* | [**GetMsgVpnQueueSubscriptions**](docs/QueueApi.md#getmsgvpnqueuesubscriptions) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions | Gets a list of Queue Subscription objects.
*QueueApi* | [**GetMsgVpnQueues**](docs/QueueApi.md#getmsgvpnqueues) | **Get** /msgVpns/{msgVpnName}/queues | Gets a list of Queue objects.
*QueueApi* | [**ReplaceMsgVpnQueue**](docs/QueueApi.md#replacemsgvpnqueue) | **Put** /msgVpns/{msgVpnName}/queues/{queueName} | Replaces a Queue object.
*QueueApi* | [**UpdateMsgVpnQueue**](docs/QueueApi.md#updatemsgvpnqueue) | **Patch** /msgVpns/{msgVpnName}/queues/{queueName} | Updates a Queue object.
*ReplicatedTopicApi* | [**CreateMsgVpnReplicatedTopic**](docs/ReplicatedTopicApi.md#createmsgvpnreplicatedtopic) | **Post** /msgVpns/{msgVpnName}/replicatedTopics | Creates a Replicated Topic object.
*ReplicatedTopicApi* | [**DeleteMsgVpnReplicatedTopic**](docs/ReplicatedTopicApi.md#deletemsgvpnreplicatedtopic) | **Delete** /msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic} | Deletes a Replicated Topic object.
*ReplicatedTopicApi* | [**GetMsgVpnReplicatedTopic**](docs/ReplicatedTopicApi.md#getmsgvpnreplicatedtopic) | **Get** /msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic} | Gets a Replicated Topic object.
*ReplicatedTopicApi* | [**GetMsgVpnReplicatedTopics**](docs/ReplicatedTopicApi.md#getmsgvpnreplicatedtopics) | **Get** /msgVpns/{msgVpnName}/replicatedTopics | Gets a list of Replicated Topic objects.
*ReplicatedTopicApi* | [**ReplaceMsgVpnReplicatedTopic**](docs/ReplicatedTopicApi.md#replacemsgvpnreplicatedtopic) | **Put** /msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic} | Replaces a Replicated Topic object.
*ReplicatedTopicApi* | [**UpdateMsgVpnReplicatedTopic**](docs/ReplicatedTopicApi.md#updatemsgvpnreplicatedtopic) | **Patch** /msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic} | Updates a Replicated Topic object.
*RestDeliveryPointApi* | [**CreateMsgVpnRestDeliveryPoint**](docs/RestDeliveryPointApi.md#createmsgvpnrestdeliverypoint) | **Post** /msgVpns/{msgVpnName}/restDeliveryPoints | Creates a REST Delivery Point object.
*RestDeliveryPointApi* | [**CreateMsgVpnRestDeliveryPointQueueBinding**](docs/RestDeliveryPointApi.md#createmsgvpnrestdeliverypointqueuebinding) | **Post** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings | Creates a Queue Binding object.
*RestDeliveryPointApi* | [**CreateMsgVpnRestDeliveryPointRestConsumer**](docs/RestDeliveryPointApi.md#createmsgvpnrestdeliverypointrestconsumer) | **Post** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers | Creates a REST Consumer object.
*RestDeliveryPointApi* | [**DeleteMsgVpnRestDeliveryPoint**](docs/RestDeliveryPointApi.md#deletemsgvpnrestdeliverypoint) | **Delete** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName} | Deletes a REST Delivery Point object.
*RestDeliveryPointApi* | [**DeleteMsgVpnRestDeliveryPointQueueBinding**](docs/RestDeliveryPointApi.md#deletemsgvpnrestdeliverypointqueuebinding) | **Delete** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName} | Deletes a Queue Binding object.
*RestDeliveryPointApi* | [**DeleteMsgVpnRestDeliveryPointRestConsumer**](docs/RestDeliveryPointApi.md#deletemsgvpnrestdeliverypointrestconsumer) | **Delete** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName} | Deletes a REST Consumer object.
*RestDeliveryPointApi* | [**GetMsgVpnRestDeliveryPoint**](docs/RestDeliveryPointApi.md#getmsgvpnrestdeliverypoint) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName} | Gets a REST Delivery Point object.
*RestDeliveryPointApi* | [**GetMsgVpnRestDeliveryPointQueueBinding**](docs/RestDeliveryPointApi.md#getmsgvpnrestdeliverypointqueuebinding) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName} | Gets a Queue Binding object.
*RestDeliveryPointApi* | [**GetMsgVpnRestDeliveryPointQueueBindings**](docs/RestDeliveryPointApi.md#getmsgvpnrestdeliverypointqueuebindings) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings | Gets a list of Queue Binding objects.
*RestDeliveryPointApi* | [**GetMsgVpnRestDeliveryPointRestConsumer**](docs/RestDeliveryPointApi.md#getmsgvpnrestdeliverypointrestconsumer) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName} | Gets a REST Consumer object.
*RestDeliveryPointApi* | [**GetMsgVpnRestDeliveryPointRestConsumers**](docs/RestDeliveryPointApi.md#getmsgvpnrestdeliverypointrestconsumers) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers | Gets a list of REST Consumer objects.
*RestDeliveryPointApi* | [**GetMsgVpnRestDeliveryPoints**](docs/RestDeliveryPointApi.md#getmsgvpnrestdeliverypoints) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints | Gets a list of REST Delivery Point objects.
*RestDeliveryPointApi* | [**ReplaceMsgVpnRestDeliveryPoint**](docs/RestDeliveryPointApi.md#replacemsgvpnrestdeliverypoint) | **Put** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName} | Replaces a REST Delivery Point object.
*RestDeliveryPointApi* | [**ReplaceMsgVpnRestDeliveryPointQueueBinding**](docs/RestDeliveryPointApi.md#replacemsgvpnrestdeliverypointqueuebinding) | **Put** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName} | Replaces a Queue Binding object.
*RestDeliveryPointApi* | [**ReplaceMsgVpnRestDeliveryPointRestConsumer**](docs/RestDeliveryPointApi.md#replacemsgvpnrestdeliverypointrestconsumer) | **Put** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName} | Replaces a REST Consumer object.
*RestDeliveryPointApi* | [**UpdateMsgVpnRestDeliveryPoint**](docs/RestDeliveryPointApi.md#updatemsgvpnrestdeliverypoint) | **Patch** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName} | Updates a REST Delivery Point object.
*RestDeliveryPointApi* | [**UpdateMsgVpnRestDeliveryPointQueueBinding**](docs/RestDeliveryPointApi.md#updatemsgvpnrestdeliverypointqueuebinding) | **Patch** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName} | Updates a Queue Binding object.
*RestDeliveryPointApi* | [**UpdateMsgVpnRestDeliveryPointRestConsumer**](docs/RestDeliveryPointApi.md#updatemsgvpnrestdeliverypointrestconsumer) | **Patch** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName} | Updates a REST Consumer object.
*SystemInformationApi* | [**GetSystemInformation**](docs/SystemInformationApi.md#getsysteminformation) | **Get** /systemInformation | Gets SEMP API version and platform information.
*TopicEndpointApi* | [**CreateMsgVpnTopicEndpoint**](docs/TopicEndpointApi.md#createmsgvpntopicendpoint) | **Post** /msgVpns/{msgVpnName}/topicEndpoints | Creates a Topic Endpoint object.
*TopicEndpointApi* | [**DeleteMsgVpnTopicEndpoint**](docs/TopicEndpointApi.md#deletemsgvpntopicendpoint) | **Delete** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName} | Deletes a Topic Endpoint object.
*TopicEndpointApi* | [**GetMsgVpnTopicEndpoint**](docs/TopicEndpointApi.md#getmsgvpntopicendpoint) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName} | Gets a Topic Endpoint object.
*TopicEndpointApi* | [**GetMsgVpnTopicEndpoints**](docs/TopicEndpointApi.md#getmsgvpntopicendpoints) | **Get** /msgVpns/{msgVpnName}/topicEndpoints | Gets a list of Topic Endpoint objects.
*TopicEndpointApi* | [**ReplaceMsgVpnTopicEndpoint**](docs/TopicEndpointApi.md#replacemsgvpntopicendpoint) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName} | Replaces a Topic Endpoint object.
*TopicEndpointApi* | [**UpdateMsgVpnTopicEndpoint**](docs/TopicEndpointApi.md#updatemsgvpntopicendpoint) | **Patch** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName} | Updates a Topic Endpoint object.


## Documentation For Models

 - [AboutApi](docs/AboutApi.md)
 - [AboutApiLinks](docs/AboutApiLinks.md)
 - [AboutApiResponse](docs/AboutApiResponse.md)
 - [AboutUser](docs/AboutUser.md)
 - [AboutUserLinks](docs/AboutUserLinks.md)
 - [AboutUserMsgVpn](docs/AboutUserMsgVpn.md)
 - [AboutUserMsgVpnLinks](docs/AboutUserMsgVpnLinks.md)
 - [AboutUserMsgVpnResponse](docs/AboutUserMsgVpnResponse.md)
 - [AboutUserMsgVpnsResponse](docs/AboutUserMsgVpnsResponse.md)
 - [AboutUserResponse](docs/AboutUserResponse.md)
 - [EventThreshold](docs/EventThreshold.md)
 - [EventThresholdByPercent](docs/EventThresholdByPercent.md)
 - [EventThresholdByValue](docs/EventThresholdByValue.md)
 - [MsgVpn](docs/MsgVpn.md)
 - [MsgVpnAclProfile](docs/MsgVpnAclProfile.md)
 - [MsgVpnAclProfileClientConnectException](docs/MsgVpnAclProfileClientConnectException.md)
 - [MsgVpnAclProfileClientConnectExceptionLinks](docs/MsgVpnAclProfileClientConnectExceptionLinks.md)
 - [MsgVpnAclProfileClientConnectExceptionResponse](docs/MsgVpnAclProfileClientConnectExceptionResponse.md)
 - [MsgVpnAclProfileClientConnectExceptionsResponse](docs/MsgVpnAclProfileClientConnectExceptionsResponse.md)
 - [MsgVpnAclProfileLinks](docs/MsgVpnAclProfileLinks.md)
 - [MsgVpnAclProfilePublishException](docs/MsgVpnAclProfilePublishException.md)
 - [MsgVpnAclProfilePublishExceptionLinks](docs/MsgVpnAclProfilePublishExceptionLinks.md)
 - [MsgVpnAclProfilePublishExceptionResponse](docs/MsgVpnAclProfilePublishExceptionResponse.md)
 - [MsgVpnAclProfilePublishExceptionsResponse](docs/MsgVpnAclProfilePublishExceptionsResponse.md)
 - [MsgVpnAclProfileResponse](docs/MsgVpnAclProfileResponse.md)
 - [MsgVpnAclProfileSubscribeException](docs/MsgVpnAclProfileSubscribeException.md)
 - [MsgVpnAclProfileSubscribeExceptionLinks](docs/MsgVpnAclProfileSubscribeExceptionLinks.md)
 - [MsgVpnAclProfileSubscribeExceptionResponse](docs/MsgVpnAclProfileSubscribeExceptionResponse.md)
 - [MsgVpnAclProfileSubscribeExceptionsResponse](docs/MsgVpnAclProfileSubscribeExceptionsResponse.md)
 - [MsgVpnAclProfilesResponse](docs/MsgVpnAclProfilesResponse.md)
 - [MsgVpnAuthorizationGroup](docs/MsgVpnAuthorizationGroup.md)
 - [MsgVpnAuthorizationGroupLinks](docs/MsgVpnAuthorizationGroupLinks.md)
 - [MsgVpnAuthorizationGroupResponse](docs/MsgVpnAuthorizationGroupResponse.md)
 - [MsgVpnAuthorizationGroupsResponse](docs/MsgVpnAuthorizationGroupsResponse.md)
 - [MsgVpnBridge](docs/MsgVpnBridge.md)
 - [MsgVpnBridgeLinks](docs/MsgVpnBridgeLinks.md)
 - [MsgVpnBridgeRemoteMsgVpn](docs/MsgVpnBridgeRemoteMsgVpn.md)
 - [MsgVpnBridgeRemoteMsgVpnLinks](docs/MsgVpnBridgeRemoteMsgVpnLinks.md)
 - [MsgVpnBridgeRemoteMsgVpnResponse](docs/MsgVpnBridgeRemoteMsgVpnResponse.md)
 - [MsgVpnBridgeRemoteMsgVpnsResponse](docs/MsgVpnBridgeRemoteMsgVpnsResponse.md)
 - [MsgVpnBridgeRemoteSubscription](docs/MsgVpnBridgeRemoteSubscription.md)
 - [MsgVpnBridgeRemoteSubscriptionLinks](docs/MsgVpnBridgeRemoteSubscriptionLinks.md)
 - [MsgVpnBridgeRemoteSubscriptionResponse](docs/MsgVpnBridgeRemoteSubscriptionResponse.md)
 - [MsgVpnBridgeRemoteSubscriptionsResponse](docs/MsgVpnBridgeRemoteSubscriptionsResponse.md)
 - [MsgVpnBridgeResponse](docs/MsgVpnBridgeResponse.md)
 - [MsgVpnBridgeTlsTrustedCommonName](docs/MsgVpnBridgeTlsTrustedCommonName.md)
 - [MsgVpnBridgeTlsTrustedCommonNameLinks](docs/MsgVpnBridgeTlsTrustedCommonNameLinks.md)
 - [MsgVpnBridgeTlsTrustedCommonNameResponse](docs/MsgVpnBridgeTlsTrustedCommonNameResponse.md)
 - [MsgVpnBridgeTlsTrustedCommonNamesResponse](docs/MsgVpnBridgeTlsTrustedCommonNamesResponse.md)
 - [MsgVpnBridgesResponse](docs/MsgVpnBridgesResponse.md)
 - [MsgVpnClientProfile](docs/MsgVpnClientProfile.md)
 - [MsgVpnClientProfileLinks](docs/MsgVpnClientProfileLinks.md)
 - [MsgVpnClientProfileResponse](docs/MsgVpnClientProfileResponse.md)
 - [MsgVpnClientProfilesResponse](docs/MsgVpnClientProfilesResponse.md)
 - [MsgVpnClientUsername](docs/MsgVpnClientUsername.md)
 - [MsgVpnClientUsernameLinks](docs/MsgVpnClientUsernameLinks.md)
 - [MsgVpnClientUsernameResponse](docs/MsgVpnClientUsernameResponse.md)
 - [MsgVpnClientUsernamesResponse](docs/MsgVpnClientUsernamesResponse.md)
 - [MsgVpnJndiConnectionFactoriesResponse](docs/MsgVpnJndiConnectionFactoriesResponse.md)
 - [MsgVpnJndiConnectionFactory](docs/MsgVpnJndiConnectionFactory.md)
 - [MsgVpnJndiConnectionFactoryLinks](docs/MsgVpnJndiConnectionFactoryLinks.md)
 - [MsgVpnJndiConnectionFactoryResponse](docs/MsgVpnJndiConnectionFactoryResponse.md)
 - [MsgVpnJndiQueue](docs/MsgVpnJndiQueue.md)
 - [MsgVpnJndiQueueLinks](docs/MsgVpnJndiQueueLinks.md)
 - [MsgVpnJndiQueueResponse](docs/MsgVpnJndiQueueResponse.md)
 - [MsgVpnJndiQueuesResponse](docs/MsgVpnJndiQueuesResponse.md)
 - [MsgVpnJndiTopic](docs/MsgVpnJndiTopic.md)
 - [MsgVpnJndiTopicLinks](docs/MsgVpnJndiTopicLinks.md)
 - [MsgVpnJndiTopicResponse](docs/MsgVpnJndiTopicResponse.md)
 - [MsgVpnJndiTopicsResponse](docs/MsgVpnJndiTopicsResponse.md)
 - [MsgVpnLinks](docs/MsgVpnLinks.md)
 - [MsgVpnMqttSession](docs/MsgVpnMqttSession.md)
 - [MsgVpnMqttSessionLinks](docs/MsgVpnMqttSessionLinks.md)
 - [MsgVpnMqttSessionResponse](docs/MsgVpnMqttSessionResponse.md)
 - [MsgVpnMqttSessionSubscription](docs/MsgVpnMqttSessionSubscription.md)
 - [MsgVpnMqttSessionSubscriptionLinks](docs/MsgVpnMqttSessionSubscriptionLinks.md)
 - [MsgVpnMqttSessionSubscriptionResponse](docs/MsgVpnMqttSessionSubscriptionResponse.md)
 - [MsgVpnMqttSessionSubscriptionsResponse](docs/MsgVpnMqttSessionSubscriptionsResponse.md)
 - [MsgVpnMqttSessionsResponse](docs/MsgVpnMqttSessionsResponse.md)
 - [MsgVpnQueue](docs/MsgVpnQueue.md)
 - [MsgVpnQueueLinks](docs/MsgVpnQueueLinks.md)
 - [MsgVpnQueueResponse](docs/MsgVpnQueueResponse.md)
 - [MsgVpnQueueSubscription](docs/MsgVpnQueueSubscription.md)
 - [MsgVpnQueueSubscriptionLinks](docs/MsgVpnQueueSubscriptionLinks.md)
 - [MsgVpnQueueSubscriptionResponse](docs/MsgVpnQueueSubscriptionResponse.md)
 - [MsgVpnQueueSubscriptionsResponse](docs/MsgVpnQueueSubscriptionsResponse.md)
 - [MsgVpnQueuesResponse](docs/MsgVpnQueuesResponse.md)
 - [MsgVpnReplicatedTopic](docs/MsgVpnReplicatedTopic.md)
 - [MsgVpnReplicatedTopicLinks](docs/MsgVpnReplicatedTopicLinks.md)
 - [MsgVpnReplicatedTopicResponse](docs/MsgVpnReplicatedTopicResponse.md)
 - [MsgVpnReplicatedTopicsResponse](docs/MsgVpnReplicatedTopicsResponse.md)
 - [MsgVpnResponse](docs/MsgVpnResponse.md)
 - [MsgVpnRestDeliveryPoint](docs/MsgVpnRestDeliveryPoint.md)
 - [MsgVpnRestDeliveryPointLinks](docs/MsgVpnRestDeliveryPointLinks.md)
 - [MsgVpnRestDeliveryPointQueueBinding](docs/MsgVpnRestDeliveryPointQueueBinding.md)
 - [MsgVpnRestDeliveryPointQueueBindingLinks](docs/MsgVpnRestDeliveryPointQueueBindingLinks.md)
 - [MsgVpnRestDeliveryPointQueueBindingResponse](docs/MsgVpnRestDeliveryPointQueueBindingResponse.md)
 - [MsgVpnRestDeliveryPointQueueBindingsResponse](docs/MsgVpnRestDeliveryPointQueueBindingsResponse.md)
 - [MsgVpnRestDeliveryPointResponse](docs/MsgVpnRestDeliveryPointResponse.md)
 - [MsgVpnRestDeliveryPointRestConsumer](docs/MsgVpnRestDeliveryPointRestConsumer.md)
 - [MsgVpnRestDeliveryPointRestConsumerLinks](docs/MsgVpnRestDeliveryPointRestConsumerLinks.md)
 - [MsgVpnRestDeliveryPointRestConsumerResponse](docs/MsgVpnRestDeliveryPointRestConsumerResponse.md)
 - [MsgVpnRestDeliveryPointRestConsumersResponse](docs/MsgVpnRestDeliveryPointRestConsumersResponse.md)
 - [MsgVpnRestDeliveryPointsResponse](docs/MsgVpnRestDeliveryPointsResponse.md)
 - [MsgVpnSequencedTopic](docs/MsgVpnSequencedTopic.md)
 - [MsgVpnSequencedTopicLinks](docs/MsgVpnSequencedTopicLinks.md)
 - [MsgVpnSequencedTopicResponse](docs/MsgVpnSequencedTopicResponse.md)
 - [MsgVpnSequencedTopicsResponse](docs/MsgVpnSequencedTopicsResponse.md)
 - [MsgVpnTopicEndpoint](docs/MsgVpnTopicEndpoint.md)
 - [MsgVpnTopicEndpointLinks](docs/MsgVpnTopicEndpointLinks.md)
 - [MsgVpnTopicEndpointResponse](docs/MsgVpnTopicEndpointResponse.md)
 - [MsgVpnTopicEndpointsResponse](docs/MsgVpnTopicEndpointsResponse.md)
 - [MsgVpnsResponse](docs/MsgVpnsResponse.md)
 - [SempError](docs/SempError.md)
 - [SempMeta](docs/SempMeta.md)
 - [SempMetaOnlyResponse](docs/SempMetaOnlyResponse.md)
 - [SempPaging](docs/SempPaging.md)
 - [SempRequest](docs/SempRequest.md)
 - [SystemInformation](docs/SystemInformation.md)
 - [SystemInformationLinks](docs/SystemInformationLinks.md)
 - [SystemInformationResponse](docs/SystemInformationResponse.md)


## Documentation For Authorization


## basicAuth

- **Type**: HTTP basic authentication


## Author

support_request@solace.com

