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

import (
	"net/url"
	"strings"
	"encoding/json"
	"fmt"
)

type MsgVpnApi struct {
	Configuration *Configuration
}

func NewMsgVpnApi() *MsgVpnApi {
	configuration := NewConfiguration()
	return &MsgVpnApi{
		Configuration: configuration,
	}
}

func NewMsgVpnApiWithBasePath(basePath string) *MsgVpnApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &MsgVpnApi{
		Configuration: configuration,
	}
}

/**
 * Creates a Message VPN object.
 * Creates a Message VPN object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: msgVpnName|x|x||| replicationBridgeAuthenticationBasicPassword||||x| replicationEnabledQueueBehavior||||x|    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent EventThresholdByValue|clearValue|setValue| EventThresholdByValue|setValue|clearValue| MsgVpn|authenticationBasicProfileName|authenticationBasicType| MsgVpn|authorizationProfileName|authorizationType| MsgVpn|eventPublishTopicFormatMqttEnabled|eventPublishTopicFormatSmfEnabled| MsgVpn|eventPublishTopicFormatSmfEnabled|eventPublishTopicFormatMqttEnabled| MsgVpn|replicationBridgeAuthenticationBasicClientUsername|replicationBridgeAuthenticationBasicPassword| MsgVpn|replicationBridgeAuthenticationBasicPassword|replicationBridgeAuthenticationBasicClientUsername| MsgVpn|replicationEnabledQueueBehavior|replicationEnabled|    A SEMP client authorized with a minimum access scope/level of \&quot;global/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param body The Message VPN object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnResponse
 */
func (a MsgVpnApi) CreateMsgVpn(body MsgVpn, select_ []string) (*MsgVpnResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "CreateMsgVpn", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Creates an ACL Profile object.
 * Creates an ACL Profile object. Any attribute missing from the request will be set to its default value.  ACL profiles are a collection of user-defined profiles for specifying Access Control List properties.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: aclProfileName|x|x||| msgVpnName|x||x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param body The ACL Profile object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnAclProfileResponse
 */
func (a MsgVpnApi) CreateMsgVpnAclProfile(msgVpnName string, body MsgVpnAclProfile, select_ []string) (*MsgVpnAclProfileResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/aclProfiles"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnAclProfileResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "CreateMsgVpnAclProfile", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Creates a Client Connect Exception object.
 * Creates a Client Connect Exception object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: aclProfileName|x||x|| clientConnectExceptionAddress|x|x||| msgVpnName|x||x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param aclProfileName The aclProfileName of the ACL Profile.
 * @param body The Client Connect Exception object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnAclProfileClientConnectExceptionResponse
 */
func (a MsgVpnApi) CreateMsgVpnAclProfileClientConnectException(msgVpnName string, aclProfileName string, body MsgVpnAclProfileClientConnectException, select_ []string) (*MsgVpnAclProfileClientConnectExceptionResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"aclProfileName"+"}", fmt.Sprintf("%v", aclProfileName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnAclProfileClientConnectExceptionResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "CreateMsgVpnAclProfileClientConnectException", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Creates a Publish Topic Exception object.
 * Creates a Publish Topic Exception object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: aclProfileName|x||x|| msgVpnName|x||x|| publishExceptionTopic|x|x||| topicSyntax|x|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param aclProfileName The aclProfileName of the ACL Profile.
 * @param body The Publish Topic Exception object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnAclProfilePublishExceptionResponse
 */
func (a MsgVpnApi) CreateMsgVpnAclProfilePublishException(msgVpnName string, aclProfileName string, body MsgVpnAclProfilePublishException, select_ []string) (*MsgVpnAclProfilePublishExceptionResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"aclProfileName"+"}", fmt.Sprintf("%v", aclProfileName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnAclProfilePublishExceptionResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "CreateMsgVpnAclProfilePublishException", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Creates a Subscribe Topic Exception object.
 * Creates a Subscribe Topic Exception object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: aclProfileName|x||x|| msgVpnName|x||x|| subscribeExceptionTopic|x|x||| topicSyntax|x|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param aclProfileName The aclProfileName of the ACL Profile.
 * @param body The Subscribe Topic Exception object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnAclProfileSubscribeExceptionResponse
 */
func (a MsgVpnApi) CreateMsgVpnAclProfileSubscribeException(msgVpnName string, aclProfileName string, body MsgVpnAclProfileSubscribeException, select_ []string) (*MsgVpnAclProfileSubscribeExceptionResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"aclProfileName"+"}", fmt.Sprintf("%v", aclProfileName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnAclProfileSubscribeExceptionResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "CreateMsgVpnAclProfileSubscribeException", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Creates an LDAP Authorization Group object.
 * Creates an LDAP Authorization Group object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: authorizationGroupName|x|x||| msgVpnName|x||x|| orderAfterAuthorizationGroupName||||x| orderBeforeAuthorizationGroupName||||x|    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnAuthorizationGroup|orderAfterAuthorizationGroupName||orderBeforeAuthorizationGroupName MsgVpnAuthorizationGroup|orderBeforeAuthorizationGroupName||orderAfterAuthorizationGroupName    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param body The LDAP Authorization Group object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnAuthorizationGroupResponse
 */
func (a MsgVpnApi) CreateMsgVpnAuthorizationGroup(msgVpnName string, body MsgVpnAuthorizationGroup, select_ []string) (*MsgVpnAuthorizationGroupResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/authorizationGroups"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnAuthorizationGroupResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "CreateMsgVpnAuthorizationGroup", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Creates a Bridge object.
 * Creates a Bridge object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: bridgeName|x|x||| bridgeVirtualRouter|x|x||| msgVpnName|x||x|| remoteAuthenticationBasicPassword||||x|    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnBridge|remoteAuthenticationBasicClientUsername|remoteAuthenticationBasicPassword| MsgVpnBridge|remoteAuthenticationBasicPassword|remoteAuthenticationBasicClientUsername|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation. Requests which include the following attributes require greater access scope/level:   Attribute|Access Scope/Level :---|:---: tlsCipherSuiteList|global/readwrite    This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param body The Bridge object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnBridgeResponse
 */
func (a MsgVpnApi) CreateMsgVpnBridge(msgVpnName string, body MsgVpnBridge, select_ []string) (*MsgVpnBridgeResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/bridges"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnBridgeResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "CreateMsgVpnBridge", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Creates a Remote Message VPN object.
 * Creates a Remote Message VPN object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: bridgeName|x||x|| bridgeVirtualRouter|x||x|| msgVpnName|x||x|| password||||x| remoteMsgVpnInterface|x|||| remoteMsgVpnLocation|x|x||| remoteMsgVpnName|x|x|||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnBridgeRemoteMsgVpn|clientUsername|password| MsgVpnBridgeRemoteMsgVpn|password|clientUsername|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param bridgeName The bridgeName of the Bridge.
 * @param bridgeVirtualRouter The bridgeVirtualRouter of the Bridge.
 * @param body The Remote Message VPN object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnBridgeRemoteMsgVpnResponse
 */
func (a MsgVpnApi) CreateMsgVpnBridgeRemoteMsgVpn(msgVpnName string, bridgeName string, bridgeVirtualRouter string, body MsgVpnBridgeRemoteMsgVpn, select_ []string) (*MsgVpnBridgeRemoteMsgVpnResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeName"+"}", fmt.Sprintf("%v", bridgeName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeVirtualRouter"+"}", fmt.Sprintf("%v", bridgeVirtualRouter), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnBridgeRemoteMsgVpnResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "CreateMsgVpnBridgeRemoteMsgVpn", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Creates a Remote Subscription object.
 * Creates a Remote Subscription object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: bridgeName|x||x|| bridgeVirtualRouter|x||x|| deliverAlwaysEnabled||x||| msgVpnName|x||x|| remoteSubscriptionTopic|x|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param bridgeName The bridgeName of the Bridge.
 * @param bridgeVirtualRouter The bridgeVirtualRouter of the Bridge.
 * @param body The Remote Subscription object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnBridgeRemoteSubscriptionResponse
 */
func (a MsgVpnApi) CreateMsgVpnBridgeRemoteSubscription(msgVpnName string, bridgeName string, bridgeVirtualRouter string, body MsgVpnBridgeRemoteSubscription, select_ []string) (*MsgVpnBridgeRemoteSubscriptionResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeName"+"}", fmt.Sprintf("%v", bridgeName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeVirtualRouter"+"}", fmt.Sprintf("%v", bridgeVirtualRouter), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnBridgeRemoteSubscriptionResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "CreateMsgVpnBridgeRemoteSubscription", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Creates a Trusted Common Name object.
 * Creates a Trusted Common Name object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: bridgeName|x||x|| bridgeVirtualRouter|x||x|| msgVpnName|x||x|| tlsTrustedCommonName|x|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;global/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param bridgeName The bridgeName of the Bridge.
 * @param bridgeVirtualRouter The bridgeVirtualRouter of the Bridge.
 * @param body The Trusted Common Name object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnBridgeTlsTrustedCommonNameResponse
 */
func (a MsgVpnApi) CreateMsgVpnBridgeTlsTrustedCommonName(msgVpnName string, bridgeName string, bridgeVirtualRouter string, body MsgVpnBridgeTlsTrustedCommonName, select_ []string) (*MsgVpnBridgeTlsTrustedCommonNameResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeName"+"}", fmt.Sprintf("%v", bridgeName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeVirtualRouter"+"}", fmt.Sprintf("%v", bridgeVirtualRouter), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnBridgeTlsTrustedCommonNameResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "CreateMsgVpnBridgeTlsTrustedCommonName", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Creates a Client Profile object.
 * Creates a Client Profile object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: clientProfileName|x|x||| msgVpnName|x||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent EventThresholdByPercent|clearPercent|setPercent| EventThresholdByPercent|setPercent|clearPercent|    A SEMP client authorized with a minimum access scope/level of \&quot;global/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param body The Client Profile object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnClientProfileResponse
 */
func (a MsgVpnApi) CreateMsgVpnClientProfile(msgVpnName string, body MsgVpnClientProfile, select_ []string) (*MsgVpnClientProfileResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/clientProfiles"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnClientProfileResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "CreateMsgVpnClientProfile", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Creates a Client Username object.
 * Creates a Client Username object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: clientUsername|x|x||| msgVpnName|x||x|| password||||x|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param body The Client Username object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnClientUsernameResponse
 */
func (a MsgVpnApi) CreateMsgVpnClientUsername(msgVpnName string, body MsgVpnClientUsername, select_ []string) (*MsgVpnClientUsernameResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/clientUsernames"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnClientUsernameResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "CreateMsgVpnClientUsername", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Creates a JNDI Connection Factory object.
 * Creates a JNDI Connection Factory object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: connectionFactoryName|x|x||| msgVpnName|x||x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.2.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param body The JNDI Connection Factory object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnJndiConnectionFactoryResponse
 */
func (a MsgVpnApi) CreateMsgVpnJndiConnectionFactory(msgVpnName string, body MsgVpnJndiConnectionFactory, select_ []string) (*MsgVpnJndiConnectionFactoryResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/jndiConnectionFactories"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnJndiConnectionFactoryResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "CreateMsgVpnJndiConnectionFactory", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Creates a JNDI Queue object.
 * Creates a JNDI Queue object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x|| queueName|x|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.2.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param body The JNDI Queue object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnJndiQueueResponse
 */
func (a MsgVpnApi) CreateMsgVpnJndiQueue(msgVpnName string, body MsgVpnJndiQueue, select_ []string) (*MsgVpnJndiQueueResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/jndiQueues"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnJndiQueueResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "CreateMsgVpnJndiQueue", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Creates a JNDI Topic object.
 * Creates a JNDI Topic object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x|| topicName|x|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.2.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param body The JNDI Topic object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnJndiTopicResponse
 */
func (a MsgVpnApi) CreateMsgVpnJndiTopic(msgVpnName string, body MsgVpnJndiTopic, select_ []string) (*MsgVpnJndiTopicResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/jndiTopics"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnJndiTopicResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "CreateMsgVpnJndiTopic", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Creates an MQTT Session object.
 * Creates an MQTT Session object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: mqttSessionClientId|x|x||| mqttSessionVirtualRouter|x|x||| msgVpnName|x||x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.1.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param body The MQTT Session object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnMqttSessionResponse
 */
func (a MsgVpnApi) CreateMsgVpnMqttSession(msgVpnName string, body MsgVpnMqttSession, select_ []string) (*MsgVpnMqttSessionResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/mqttSessions"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnMqttSessionResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "CreateMsgVpnMqttSession", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Creates an MQTT Session Subscription object.
 * Creates an MQTT Session Subscription object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: mqttSessionClientId|x||x|| mqttSessionVirtualRouter|x||x|| msgVpnName|x||x|| subscriptionTopic|x|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.1.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param mqttSessionClientId The mqttSessionClientId of the MQTT Session.
 * @param mqttSessionVirtualRouter The mqttSessionVirtualRouter of the MQTT Session.
 * @param body The MQTT Session Subscription object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnMqttSessionSubscriptionResponse
 */
func (a MsgVpnApi) CreateMsgVpnMqttSessionSubscription(msgVpnName string, mqttSessionClientId string, mqttSessionVirtualRouter string, body MsgVpnMqttSessionSubscription, select_ []string) (*MsgVpnMqttSessionSubscriptionResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mqttSessionClientId"+"}", fmt.Sprintf("%v", mqttSessionClientId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mqttSessionVirtualRouter"+"}", fmt.Sprintf("%v", mqttSessionVirtualRouter), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnMqttSessionSubscriptionResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "CreateMsgVpnMqttSessionSubscription", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Creates a Queue object.
 * Creates a Queue object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x|| queueName|x|x|||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param body The Queue object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnQueueResponse
 */
func (a MsgVpnApi) CreateMsgVpnQueue(msgVpnName string, body MsgVpnQueue, select_ []string) (*MsgVpnQueueResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/queues"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnQueueResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "CreateMsgVpnQueue", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Creates a Queue Subscription object.
 * Creates a Queue Subscription object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x|| queueName|x||x|| subscriptionTopic|x|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param queueName The queueName of the Queue.
 * @param body The Queue Subscription object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnQueueSubscriptionResponse
 */
func (a MsgVpnApi) CreateMsgVpnQueueSubscription(msgVpnName string, queueName string, body MsgVpnQueueSubscription, select_ []string) (*MsgVpnQueueSubscriptionResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/queues/{queueName}/subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueName"+"}", fmt.Sprintf("%v", queueName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnQueueSubscriptionResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "CreateMsgVpnQueueSubscription", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Creates a Replicated Topic object.
 * Creates a Replicated Topic object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x|| replicatedTopic|x|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.1.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param body The Replicated Topic object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnReplicatedTopicResponse
 */
func (a MsgVpnApi) CreateMsgVpnReplicatedTopic(msgVpnName string, body MsgVpnReplicatedTopic, select_ []string) (*MsgVpnReplicatedTopicResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/replicatedTopics"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnReplicatedTopicResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "CreateMsgVpnReplicatedTopic", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Creates a REST Delivery Point object.
 * Creates a REST Delivery Point object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x|| restDeliveryPointName|x|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param body The REST Delivery Point object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnRestDeliveryPointResponse
 */
func (a MsgVpnApi) CreateMsgVpnRestDeliveryPoint(msgVpnName string, body MsgVpnRestDeliveryPoint, select_ []string) (*MsgVpnRestDeliveryPointResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnRestDeliveryPointResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "CreateMsgVpnRestDeliveryPoint", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Creates a Queue Binding object.
 * Creates a Queue Binding object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x|| queueBindingName|x|x||| restDeliveryPointName|x||x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param restDeliveryPointName The restDeliveryPointName of the REST Delivery Point.
 * @param body The Queue Binding object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnRestDeliveryPointQueueBindingResponse
 */
func (a MsgVpnApi) CreateMsgVpnRestDeliveryPointQueueBinding(msgVpnName string, restDeliveryPointName string, body MsgVpnRestDeliveryPointQueueBinding, select_ []string) (*MsgVpnRestDeliveryPointQueueBindingResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnRestDeliveryPointQueueBindingResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "CreateMsgVpnRestDeliveryPointQueueBinding", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Creates a REST Consumer object.
 * Creates a REST Consumer object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: authenticationHttpBasicPassword||||x| msgVpnName|x||x|| restConsumerName|x|x||| restDeliveryPointName|x||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnRestDeliveryPointRestConsumer|authenticationHttpBasicPassword|authenticationHttpBasicUsername| MsgVpnRestDeliveryPointRestConsumer|authenticationHttpBasicUsername|authenticationHttpBasicPassword| MsgVpnRestDeliveryPointRestConsumer|remotePort|tlsEnabled| MsgVpnRestDeliveryPointRestConsumer|tlsEnabled|remotePort|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param restDeliveryPointName The restDeliveryPointName of the REST Delivery Point.
 * @param body The REST Consumer object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnRestDeliveryPointRestConsumerResponse
 */
func (a MsgVpnApi) CreateMsgVpnRestDeliveryPointRestConsumer(msgVpnName string, restDeliveryPointName string, body MsgVpnRestDeliveryPointRestConsumer, select_ []string) (*MsgVpnRestDeliveryPointRestConsumerResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnRestDeliveryPointRestConsumerResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "CreateMsgVpnRestDeliveryPointRestConsumer", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Creates a Sequenced Topic object.
 * Creates a Sequenced Topic object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x|| sequencedTopic|x|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param body The Sequenced Topic object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnSequencedTopicResponse
 */
func (a MsgVpnApi) CreateMsgVpnSequencedTopic(msgVpnName string, body MsgVpnSequencedTopic, select_ []string) (*MsgVpnSequencedTopicResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/sequencedTopics"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnSequencedTopicResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "CreateMsgVpnSequencedTopic", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Creates a Topic Endpoint object.
 * Creates a Topic Endpoint object. Any attribute missing from the request will be set to its default value.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated :---|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x|| topicEndpointName|x|x|||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.1.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param body The Topic Endpoint object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnTopicEndpointResponse
 */
func (a MsgVpnApi) CreateMsgVpnTopicEndpoint(msgVpnName string, body MsgVpnTopicEndpoint, select_ []string) (*MsgVpnTopicEndpointResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/topicEndpoints"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnTopicEndpointResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "CreateMsgVpnTopicEndpoint", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Deletes a Message VPN object.
 * Deletes a Message VPN object.  A SEMP client authorized with a minimum access scope/level of \&quot;global/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @return *SempMetaOnlyResponse
 */
func (a MsgVpnApi) DeleteMsgVpn(msgVpnName string) (*SempMetaOnlyResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SempMetaOnlyResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DeleteMsgVpn", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Deletes an ACL Profile object.
 * Deletes an ACL Profile object.  ACL profiles are a collection of user-defined profiles for specifying Access Control List properties.  A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param aclProfileName The aclProfileName of the ACL Profile.
 * @return *SempMetaOnlyResponse
 */
func (a MsgVpnApi) DeleteMsgVpnAclProfile(msgVpnName string, aclProfileName string) (*SempMetaOnlyResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"aclProfileName"+"}", fmt.Sprintf("%v", aclProfileName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SempMetaOnlyResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DeleteMsgVpnAclProfile", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Deletes a Client Connect Exception object.
 * Deletes a Client Connect Exception object.  A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param aclProfileName The aclProfileName of the ACL Profile.
 * @param clientConnectExceptionAddress The clientConnectExceptionAddress of the Client Connect Exception.
 * @return *SempMetaOnlyResponse
 */
func (a MsgVpnApi) DeleteMsgVpnAclProfileClientConnectException(msgVpnName string, aclProfileName string, clientConnectExceptionAddress string) (*SempMetaOnlyResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions/{clientConnectExceptionAddress}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"aclProfileName"+"}", fmt.Sprintf("%v", aclProfileName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientConnectExceptionAddress"+"}", fmt.Sprintf("%v", clientConnectExceptionAddress), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SempMetaOnlyResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DeleteMsgVpnAclProfileClientConnectException", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Deletes a Publish Topic Exception object.
 * Deletes a Publish Topic Exception object.  A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param aclProfileName The aclProfileName of the ACL Profile.
 * @param topicSyntax The topicSyntax of the Publish Topic Exception.
 * @param publishExceptionTopic The publishExceptionTopic of the Publish Topic Exception.
 * @return *SempMetaOnlyResponse
 */
func (a MsgVpnApi) DeleteMsgVpnAclProfilePublishException(msgVpnName string, aclProfileName string, topicSyntax string, publishExceptionTopic string) (*SempMetaOnlyResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions/{topicSyntax},{publishExceptionTopic}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"aclProfileName"+"}", fmt.Sprintf("%v", aclProfileName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topicSyntax"+"}", fmt.Sprintf("%v", topicSyntax), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"publishExceptionTopic"+"}", fmt.Sprintf("%v", publishExceptionTopic), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SempMetaOnlyResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DeleteMsgVpnAclProfilePublishException", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Deletes a Subscribe Topic Exception object.
 * Deletes a Subscribe Topic Exception object.  A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param aclProfileName The aclProfileName of the ACL Profile.
 * @param topicSyntax The topicSyntax of the Subscribe Topic Exception.
 * @param subscribeExceptionTopic The subscribeExceptionTopic of the Subscribe Topic Exception.
 * @return *SempMetaOnlyResponse
 */
func (a MsgVpnApi) DeleteMsgVpnAclProfileSubscribeException(msgVpnName string, aclProfileName string, topicSyntax string, subscribeExceptionTopic string) (*SempMetaOnlyResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions/{topicSyntax},{subscribeExceptionTopic}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"aclProfileName"+"}", fmt.Sprintf("%v", aclProfileName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topicSyntax"+"}", fmt.Sprintf("%v", topicSyntax), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subscribeExceptionTopic"+"}", fmt.Sprintf("%v", subscribeExceptionTopic), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SempMetaOnlyResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DeleteMsgVpnAclProfileSubscribeException", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Deletes an LDAP Authorization Group object.
 * Deletes an LDAP Authorization Group object.  A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param authorizationGroupName The authorizationGroupName of the LDAP Authorization Group.
 * @return *SempMetaOnlyResponse
 */
func (a MsgVpnApi) DeleteMsgVpnAuthorizationGroup(msgVpnName string, authorizationGroupName string) (*SempMetaOnlyResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"authorizationGroupName"+"}", fmt.Sprintf("%v", authorizationGroupName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SempMetaOnlyResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DeleteMsgVpnAuthorizationGroup", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Deletes a Bridge object.
 * Deletes a Bridge object.  A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation. Requests which include the following attributes require greater access scope/level:   Attribute|Access Scope/Level :---|:---: tlsCipherSuiteList|global/readwrite    This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param bridgeName The bridgeName of the Bridge.
 * @param bridgeVirtualRouter The bridgeVirtualRouter of the Bridge.
 * @return *SempMetaOnlyResponse
 */
func (a MsgVpnApi) DeleteMsgVpnBridge(msgVpnName string, bridgeName string, bridgeVirtualRouter string) (*SempMetaOnlyResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeName"+"}", fmt.Sprintf("%v", bridgeName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeVirtualRouter"+"}", fmt.Sprintf("%v", bridgeVirtualRouter), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SempMetaOnlyResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DeleteMsgVpnBridge", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Deletes a Remote Message VPN object.
 * Deletes a Remote Message VPN object.  A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param bridgeName The bridgeName of the Bridge.
 * @param bridgeVirtualRouter The bridgeVirtualRouter of the Bridge.
 * @param remoteMsgVpnName The remoteMsgVpnName of the Remote Message VPN.
 * @param remoteMsgVpnLocation The remoteMsgVpnLocation of the Remote Message VPN.
 * @param remoteMsgVpnInterface The remoteMsgVpnInterface of the Remote Message VPN.
 * @return *SempMetaOnlyResponse
 */
func (a MsgVpnApi) DeleteMsgVpnBridgeRemoteMsgVpn(msgVpnName string, bridgeName string, bridgeVirtualRouter string, remoteMsgVpnName string, remoteMsgVpnLocation string, remoteMsgVpnInterface string) (*SempMetaOnlyResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeName"+"}", fmt.Sprintf("%v", bridgeName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeVirtualRouter"+"}", fmt.Sprintf("%v", bridgeVirtualRouter), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"remoteMsgVpnName"+"}", fmt.Sprintf("%v", remoteMsgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"remoteMsgVpnLocation"+"}", fmt.Sprintf("%v", remoteMsgVpnLocation), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"remoteMsgVpnInterface"+"}", fmt.Sprintf("%v", remoteMsgVpnInterface), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SempMetaOnlyResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DeleteMsgVpnBridgeRemoteMsgVpn", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Deletes a Remote Subscription object.
 * Deletes a Remote Subscription object.  A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param bridgeName The bridgeName of the Bridge.
 * @param bridgeVirtualRouter The bridgeVirtualRouter of the Bridge.
 * @param remoteSubscriptionTopic The remoteSubscriptionTopic of the Remote Subscription.
 * @return *SempMetaOnlyResponse
 */
func (a MsgVpnApi) DeleteMsgVpnBridgeRemoteSubscription(msgVpnName string, bridgeName string, bridgeVirtualRouter string, remoteSubscriptionTopic string) (*SempMetaOnlyResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions/{remoteSubscriptionTopic}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeName"+"}", fmt.Sprintf("%v", bridgeName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeVirtualRouter"+"}", fmt.Sprintf("%v", bridgeVirtualRouter), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"remoteSubscriptionTopic"+"}", fmt.Sprintf("%v", remoteSubscriptionTopic), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SempMetaOnlyResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DeleteMsgVpnBridgeRemoteSubscription", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Deletes a Trusted Common Name object.
 * Deletes a Trusted Common Name object.  A SEMP client authorized with a minimum access scope/level of \&quot;global/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param bridgeName The bridgeName of the Bridge.
 * @param bridgeVirtualRouter The bridgeVirtualRouter of the Bridge.
 * @param tlsTrustedCommonName The tlsTrustedCommonName of the Trusted Common Name.
 * @return *SempMetaOnlyResponse
 */
func (a MsgVpnApi) DeleteMsgVpnBridgeTlsTrustedCommonName(msgVpnName string, bridgeName string, bridgeVirtualRouter string, tlsTrustedCommonName string) (*SempMetaOnlyResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames/{tlsTrustedCommonName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeName"+"}", fmt.Sprintf("%v", bridgeName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeVirtualRouter"+"}", fmt.Sprintf("%v", bridgeVirtualRouter), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tlsTrustedCommonName"+"}", fmt.Sprintf("%v", tlsTrustedCommonName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SempMetaOnlyResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DeleteMsgVpnBridgeTlsTrustedCommonName", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Deletes a Client Profile object.
 * Deletes a Client Profile object.  A SEMP client authorized with a minimum access scope/level of \&quot;global/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param clientProfileName The clientProfileName of the Client Profile.
 * @return *SempMetaOnlyResponse
 */
func (a MsgVpnApi) DeleteMsgVpnClientProfile(msgVpnName string, clientProfileName string) (*SempMetaOnlyResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/clientProfiles/{clientProfileName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientProfileName"+"}", fmt.Sprintf("%v", clientProfileName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SempMetaOnlyResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DeleteMsgVpnClientProfile", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Deletes a Client Username object.
 * Deletes a Client Username object.  A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param clientUsername The clientUsername of the Client Username.
 * @return *SempMetaOnlyResponse
 */
func (a MsgVpnApi) DeleteMsgVpnClientUsername(msgVpnName string, clientUsername string) (*SempMetaOnlyResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/clientUsernames/{clientUsername}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientUsername"+"}", fmt.Sprintf("%v", clientUsername), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SempMetaOnlyResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DeleteMsgVpnClientUsername", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Deletes a JNDI Connection Factory object.
 * Deletes a JNDI Connection Factory object.  A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.2.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param connectionFactoryName The connectionFactoryName of the JNDI Connection Factory.
 * @return *SempMetaOnlyResponse
 */
func (a MsgVpnApi) DeleteMsgVpnJndiConnectionFactory(msgVpnName string, connectionFactoryName string) (*SempMetaOnlyResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"connectionFactoryName"+"}", fmt.Sprintf("%v", connectionFactoryName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SempMetaOnlyResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DeleteMsgVpnJndiConnectionFactory", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Deletes a JNDI Queue object.
 * Deletes a JNDI Queue object.  A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.2.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param queueName The queueName of the JNDI Queue.
 * @return *SempMetaOnlyResponse
 */
func (a MsgVpnApi) DeleteMsgVpnJndiQueue(msgVpnName string, queueName string) (*SempMetaOnlyResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/jndiQueues/{queueName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueName"+"}", fmt.Sprintf("%v", queueName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SempMetaOnlyResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DeleteMsgVpnJndiQueue", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Deletes a JNDI Topic object.
 * Deletes a JNDI Topic object.  A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.2.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param topicName The topicName of the JNDI Topic.
 * @return *SempMetaOnlyResponse
 */
func (a MsgVpnApi) DeleteMsgVpnJndiTopic(msgVpnName string, topicName string) (*SempMetaOnlyResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/jndiTopics/{topicName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topicName"+"}", fmt.Sprintf("%v", topicName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SempMetaOnlyResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DeleteMsgVpnJndiTopic", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Deletes an MQTT Session object.
 * Deletes an MQTT Session object.  A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.1.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param mqttSessionClientId The mqttSessionClientId of the MQTT Session.
 * @param mqttSessionVirtualRouter The mqttSessionVirtualRouter of the MQTT Session.
 * @return *SempMetaOnlyResponse
 */
func (a MsgVpnApi) DeleteMsgVpnMqttSession(msgVpnName string, mqttSessionClientId string, mqttSessionVirtualRouter string) (*SempMetaOnlyResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mqttSessionClientId"+"}", fmt.Sprintf("%v", mqttSessionClientId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mqttSessionVirtualRouter"+"}", fmt.Sprintf("%v", mqttSessionVirtualRouter), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SempMetaOnlyResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DeleteMsgVpnMqttSession", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Deletes an MQTT Session Subscription object.
 * Deletes an MQTT Session Subscription object.  A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.1.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param mqttSessionClientId The mqttSessionClientId of the MQTT Session.
 * @param mqttSessionVirtualRouter The mqttSessionVirtualRouter of the MQTT Session.
 * @param subscriptionTopic The subscriptionTopic of the MQTT Session Subscription.
 * @return *SempMetaOnlyResponse
 */
func (a MsgVpnApi) DeleteMsgVpnMqttSessionSubscription(msgVpnName string, mqttSessionClientId string, mqttSessionVirtualRouter string, subscriptionTopic string) (*SempMetaOnlyResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mqttSessionClientId"+"}", fmt.Sprintf("%v", mqttSessionClientId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mqttSessionVirtualRouter"+"}", fmt.Sprintf("%v", mqttSessionVirtualRouter), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionTopic"+"}", fmt.Sprintf("%v", subscriptionTopic), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SempMetaOnlyResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DeleteMsgVpnMqttSessionSubscription", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Deletes a Queue object.
 * Deletes a Queue object.  A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param queueName The queueName of the Queue.
 * @return *SempMetaOnlyResponse
 */
func (a MsgVpnApi) DeleteMsgVpnQueue(msgVpnName string, queueName string) (*SempMetaOnlyResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/queues/{queueName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueName"+"}", fmt.Sprintf("%v", queueName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SempMetaOnlyResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DeleteMsgVpnQueue", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Deletes a Queue Subscription object.
 * Deletes a Queue Subscription object.  A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param queueName The queueName of the Queue.
 * @param subscriptionTopic The subscriptionTopic of the Queue Subscription.
 * @return *SempMetaOnlyResponse
 */
func (a MsgVpnApi) DeleteMsgVpnQueueSubscription(msgVpnName string, queueName string, subscriptionTopic string) (*SempMetaOnlyResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/queues/{queueName}/subscriptions/{subscriptionTopic}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueName"+"}", fmt.Sprintf("%v", queueName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionTopic"+"}", fmt.Sprintf("%v", subscriptionTopic), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SempMetaOnlyResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DeleteMsgVpnQueueSubscription", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Deletes a Replicated Topic object.
 * Deletes a Replicated Topic object.  A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.1.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param replicatedTopic The replicatedTopic of the Replicated Topic.
 * @return *SempMetaOnlyResponse
 */
func (a MsgVpnApi) DeleteMsgVpnReplicatedTopic(msgVpnName string, replicatedTopic string) (*SempMetaOnlyResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"replicatedTopic"+"}", fmt.Sprintf("%v", replicatedTopic), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SempMetaOnlyResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DeleteMsgVpnReplicatedTopic", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Deletes a REST Delivery Point object.
 * Deletes a REST Delivery Point object.  A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param restDeliveryPointName The restDeliveryPointName of the REST Delivery Point.
 * @return *SempMetaOnlyResponse
 */
func (a MsgVpnApi) DeleteMsgVpnRestDeliveryPoint(msgVpnName string, restDeliveryPointName string) (*SempMetaOnlyResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SempMetaOnlyResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DeleteMsgVpnRestDeliveryPoint", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Deletes a Queue Binding object.
 * Deletes a Queue Binding object.  A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param restDeliveryPointName The restDeliveryPointName of the REST Delivery Point.
 * @param queueBindingName The queueBindingName of the Queue Binding.
 * @return *SempMetaOnlyResponse
 */
func (a MsgVpnApi) DeleteMsgVpnRestDeliveryPointQueueBinding(msgVpnName string, restDeliveryPointName string, queueBindingName string) (*SempMetaOnlyResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueBindingName"+"}", fmt.Sprintf("%v", queueBindingName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SempMetaOnlyResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DeleteMsgVpnRestDeliveryPointQueueBinding", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Deletes a REST Consumer object.
 * Deletes a REST Consumer object.  A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param restDeliveryPointName The restDeliveryPointName of the REST Delivery Point.
 * @param restConsumerName The restConsumerName of the REST Consumer.
 * @return *SempMetaOnlyResponse
 */
func (a MsgVpnApi) DeleteMsgVpnRestDeliveryPointRestConsumer(msgVpnName string, restDeliveryPointName string, restConsumerName string) (*SempMetaOnlyResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restConsumerName"+"}", fmt.Sprintf("%v", restConsumerName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SempMetaOnlyResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DeleteMsgVpnRestDeliveryPointRestConsumer", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Deletes a Sequenced Topic object.
 * Deletes a Sequenced Topic object.  A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param sequencedTopic The sequencedTopic of the Sequenced Topic.
 * @return *SempMetaOnlyResponse
 */
func (a MsgVpnApi) DeleteMsgVpnSequencedTopic(msgVpnName string, sequencedTopic string) (*SempMetaOnlyResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/sequencedTopics/{sequencedTopic}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sequencedTopic"+"}", fmt.Sprintf("%v", sequencedTopic), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SempMetaOnlyResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DeleteMsgVpnSequencedTopic", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Deletes a Topic Endpoint object.
 * Deletes a Topic Endpoint object.  A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.1.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param topicEndpointName The topicEndpointName of the Topic Endpoint.
 * @return *SempMetaOnlyResponse
 */
func (a MsgVpnApi) DeleteMsgVpnTopicEndpoint(msgVpnName string, topicEndpointName string) (*SempMetaOnlyResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topicEndpointName"+"}", fmt.Sprintf("%v", topicEndpointName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SempMetaOnlyResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DeleteMsgVpnTopicEndpoint", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a Message VPN object.
 * Gets a Message VPN object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: msgVpnName|x|| replicationBridgeAuthenticationBasicPassword||x| replicationEnabledQueueBehavior||x|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnResponse
 */
func (a MsgVpnApi) GetMsgVpn(msgVpnName string, select_ []string) (*MsgVpnResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpn", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets an ACL Profile object.
 * Gets an ACL Profile object.  ACL profiles are a collection of user-defined profiles for specifying Access Control List properties.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: aclProfileName|x|| msgVpnName|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param aclProfileName The aclProfileName of the ACL Profile.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnAclProfileResponse
 */
func (a MsgVpnApi) GetMsgVpnAclProfile(msgVpnName string, aclProfileName string, select_ []string) (*MsgVpnAclProfileResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"aclProfileName"+"}", fmt.Sprintf("%v", aclProfileName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnAclProfileResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnAclProfile", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a Client Connect Exception object.
 * Gets a Client Connect Exception object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: aclProfileName|x|| clientConnectExceptionAddress|x|| msgVpnName|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param aclProfileName The aclProfileName of the ACL Profile.
 * @param clientConnectExceptionAddress The clientConnectExceptionAddress of the Client Connect Exception.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnAclProfileClientConnectExceptionResponse
 */
func (a MsgVpnApi) GetMsgVpnAclProfileClientConnectException(msgVpnName string, aclProfileName string, clientConnectExceptionAddress string, select_ []string) (*MsgVpnAclProfileClientConnectExceptionResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions/{clientConnectExceptionAddress}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"aclProfileName"+"}", fmt.Sprintf("%v", aclProfileName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientConnectExceptionAddress"+"}", fmt.Sprintf("%v", clientConnectExceptionAddress), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnAclProfileClientConnectExceptionResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnAclProfileClientConnectException", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a list of Client Connect Exception objects.
 * Gets a list of Client Connect Exception objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: aclProfileName|x|| clientConnectExceptionAddress|x|| msgVpnName|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param aclProfileName The aclProfileName of the ACL Profile.
 * @param count Limit the count of objects in the response. See [Count](#count \&quot;Description of the syntax of the &#x60;count&#x60; parameter\&quot;).
 * @param cursor The cursor, or position, for the next page of objects. See [Cursor](#cursor \&quot;Description of the syntax of the &#x60;cursor&#x60; parameter\&quot;).
 * @param where Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;).
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnAclProfileClientConnectExceptionsResponse
 */
func (a MsgVpnApi) GetMsgVpnAclProfileClientConnectExceptions(msgVpnName string, aclProfileName string, count int32, cursor string, where []string, select_ []string) (*MsgVpnAclProfileClientConnectExceptionsResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"aclProfileName"+"}", fmt.Sprintf("%v", aclProfileName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("count", a.Configuration.APIClient.ParameterToString(count, ""))
	localVarQueryParams.Add("cursor", a.Configuration.APIClient.ParameterToString(cursor, ""))
	var whereCollectionFormat = "csv"
	localVarQueryParams.Add("where", a.Configuration.APIClient.ParameterToString(where, whereCollectionFormat))

	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnAclProfileClientConnectExceptionsResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnAclProfileClientConnectExceptions", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a Publish Topic Exception object.
 * Gets a Publish Topic Exception object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: aclProfileName|x|| msgVpnName|x|| publishExceptionTopic|x|| topicSyntax|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param aclProfileName The aclProfileName of the ACL Profile.
 * @param topicSyntax The topicSyntax of the Publish Topic Exception.
 * @param publishExceptionTopic The publishExceptionTopic of the Publish Topic Exception.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnAclProfilePublishExceptionResponse
 */
func (a MsgVpnApi) GetMsgVpnAclProfilePublishException(msgVpnName string, aclProfileName string, topicSyntax string, publishExceptionTopic string, select_ []string) (*MsgVpnAclProfilePublishExceptionResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions/{topicSyntax},{publishExceptionTopic}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"aclProfileName"+"}", fmt.Sprintf("%v", aclProfileName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topicSyntax"+"}", fmt.Sprintf("%v", topicSyntax), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"publishExceptionTopic"+"}", fmt.Sprintf("%v", publishExceptionTopic), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnAclProfilePublishExceptionResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnAclProfilePublishException", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a list of Publish Topic Exception objects.
 * Gets a list of Publish Topic Exception objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: aclProfileName|x|| msgVpnName|x|| publishExceptionTopic|x|| topicSyntax|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param aclProfileName The aclProfileName of the ACL Profile.
 * @param count Limit the count of objects in the response. See [Count](#count \&quot;Description of the syntax of the &#x60;count&#x60; parameter\&quot;).
 * @param cursor The cursor, or position, for the next page of objects. See [Cursor](#cursor \&quot;Description of the syntax of the &#x60;cursor&#x60; parameter\&quot;).
 * @param where Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;).
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnAclProfilePublishExceptionsResponse
 */
func (a MsgVpnApi) GetMsgVpnAclProfilePublishExceptions(msgVpnName string, aclProfileName string, count int32, cursor string, where []string, select_ []string) (*MsgVpnAclProfilePublishExceptionsResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"aclProfileName"+"}", fmt.Sprintf("%v", aclProfileName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("count", a.Configuration.APIClient.ParameterToString(count, ""))
	localVarQueryParams.Add("cursor", a.Configuration.APIClient.ParameterToString(cursor, ""))
	var whereCollectionFormat = "csv"
	localVarQueryParams.Add("where", a.Configuration.APIClient.ParameterToString(where, whereCollectionFormat))

	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnAclProfilePublishExceptionsResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnAclProfilePublishExceptions", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a Subscribe Topic Exception object.
 * Gets a Subscribe Topic Exception object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: aclProfileName|x|| msgVpnName|x|| subscribeExceptionTopic|x|| topicSyntax|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param aclProfileName The aclProfileName of the ACL Profile.
 * @param topicSyntax The topicSyntax of the Subscribe Topic Exception.
 * @param subscribeExceptionTopic The subscribeExceptionTopic of the Subscribe Topic Exception.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnAclProfileSubscribeExceptionResponse
 */
func (a MsgVpnApi) GetMsgVpnAclProfileSubscribeException(msgVpnName string, aclProfileName string, topicSyntax string, subscribeExceptionTopic string, select_ []string) (*MsgVpnAclProfileSubscribeExceptionResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions/{topicSyntax},{subscribeExceptionTopic}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"aclProfileName"+"}", fmt.Sprintf("%v", aclProfileName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topicSyntax"+"}", fmt.Sprintf("%v", topicSyntax), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subscribeExceptionTopic"+"}", fmt.Sprintf("%v", subscribeExceptionTopic), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnAclProfileSubscribeExceptionResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnAclProfileSubscribeException", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a list of Subscribe Topic Exception objects.
 * Gets a list of Subscribe Topic Exception objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: aclProfileName|x|| msgVpnName|x|| subscribeExceptionTopic|x|| topicSyntax|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param aclProfileName The aclProfileName of the ACL Profile.
 * @param count Limit the count of objects in the response. See [Count](#count \&quot;Description of the syntax of the &#x60;count&#x60; parameter\&quot;).
 * @param cursor The cursor, or position, for the next page of objects. See [Cursor](#cursor \&quot;Description of the syntax of the &#x60;cursor&#x60; parameter\&quot;).
 * @param where Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;).
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnAclProfileSubscribeExceptionsResponse
 */
func (a MsgVpnApi) GetMsgVpnAclProfileSubscribeExceptions(msgVpnName string, aclProfileName string, count int32, cursor string, where []string, select_ []string) (*MsgVpnAclProfileSubscribeExceptionsResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"aclProfileName"+"}", fmt.Sprintf("%v", aclProfileName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("count", a.Configuration.APIClient.ParameterToString(count, ""))
	localVarQueryParams.Add("cursor", a.Configuration.APIClient.ParameterToString(cursor, ""))
	var whereCollectionFormat = "csv"
	localVarQueryParams.Add("where", a.Configuration.APIClient.ParameterToString(where, whereCollectionFormat))

	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnAclProfileSubscribeExceptionsResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnAclProfileSubscribeExceptions", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a list of ACL Profile objects.
 * Gets a list of ACL Profile objects.  ACL profiles are a collection of user-defined profiles for specifying Access Control List properties.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: aclProfileName|x|| msgVpnName|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param count Limit the count of objects in the response. See [Count](#count \&quot;Description of the syntax of the &#x60;count&#x60; parameter\&quot;).
 * @param cursor The cursor, or position, for the next page of objects. See [Cursor](#cursor \&quot;Description of the syntax of the &#x60;cursor&#x60; parameter\&quot;).
 * @param where Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;).
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnAclProfilesResponse
 */
func (a MsgVpnApi) GetMsgVpnAclProfiles(msgVpnName string, count int32, cursor string, where []string, select_ []string) (*MsgVpnAclProfilesResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/aclProfiles"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("count", a.Configuration.APIClient.ParameterToString(count, ""))
	localVarQueryParams.Add("cursor", a.Configuration.APIClient.ParameterToString(cursor, ""))
	var whereCollectionFormat = "csv"
	localVarQueryParams.Add("where", a.Configuration.APIClient.ParameterToString(where, whereCollectionFormat))

	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnAclProfilesResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnAclProfiles", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets an LDAP Authorization Group object.
 * Gets an LDAP Authorization Group object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: authorizationGroupName|x|| msgVpnName|x|| orderAfterAuthorizationGroupName||x| orderBeforeAuthorizationGroupName||x|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param authorizationGroupName The authorizationGroupName of the LDAP Authorization Group.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnAuthorizationGroupResponse
 */
func (a MsgVpnApi) GetMsgVpnAuthorizationGroup(msgVpnName string, authorizationGroupName string, select_ []string) (*MsgVpnAuthorizationGroupResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"authorizationGroupName"+"}", fmt.Sprintf("%v", authorizationGroupName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnAuthorizationGroupResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnAuthorizationGroup", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a list of LDAP Authorization Group objects.
 * Gets a list of LDAP Authorization Group objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: authorizationGroupName|x|| msgVpnName|x|| orderAfterAuthorizationGroupName||x| orderBeforeAuthorizationGroupName||x|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param count Limit the count of objects in the response. See [Count](#count \&quot;Description of the syntax of the &#x60;count&#x60; parameter\&quot;).
 * @param cursor The cursor, or position, for the next page of objects. See [Cursor](#cursor \&quot;Description of the syntax of the &#x60;cursor&#x60; parameter\&quot;).
 * @param where Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;).
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnAuthorizationGroupsResponse
 */
func (a MsgVpnApi) GetMsgVpnAuthorizationGroups(msgVpnName string, count int32, cursor string, where []string, select_ []string) (*MsgVpnAuthorizationGroupsResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/authorizationGroups"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("count", a.Configuration.APIClient.ParameterToString(count, ""))
	localVarQueryParams.Add("cursor", a.Configuration.APIClient.ParameterToString(cursor, ""))
	var whereCollectionFormat = "csv"
	localVarQueryParams.Add("where", a.Configuration.APIClient.ParameterToString(where, whereCollectionFormat))

	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnAuthorizationGroupsResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnAuthorizationGroups", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a Bridge object.
 * Gets a Bridge object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: bridgeName|x|| bridgeVirtualRouter|x|| msgVpnName|x|| remoteAuthenticationBasicPassword||x|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param bridgeName The bridgeName of the Bridge.
 * @param bridgeVirtualRouter The bridgeVirtualRouter of the Bridge.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnBridgeResponse
 */
func (a MsgVpnApi) GetMsgVpnBridge(msgVpnName string, bridgeName string, bridgeVirtualRouter string, select_ []string) (*MsgVpnBridgeResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeName"+"}", fmt.Sprintf("%v", bridgeName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeVirtualRouter"+"}", fmt.Sprintf("%v", bridgeVirtualRouter), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnBridgeResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnBridge", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a Remote Message VPN object.
 * Gets a Remote Message VPN object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: bridgeName|x|| bridgeVirtualRouter|x|| msgVpnName|x|| password||x| remoteMsgVpnInterface|x|| remoteMsgVpnLocation|x|| remoteMsgVpnName|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param bridgeName The bridgeName of the Bridge.
 * @param bridgeVirtualRouter The bridgeVirtualRouter of the Bridge.
 * @param remoteMsgVpnName The remoteMsgVpnName of the Remote Message VPN.
 * @param remoteMsgVpnLocation The remoteMsgVpnLocation of the Remote Message VPN.
 * @param remoteMsgVpnInterface The remoteMsgVpnInterface of the Remote Message VPN.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnBridgeRemoteMsgVpnResponse
 */
func (a MsgVpnApi) GetMsgVpnBridgeRemoteMsgVpn(msgVpnName string, bridgeName string, bridgeVirtualRouter string, remoteMsgVpnName string, remoteMsgVpnLocation string, remoteMsgVpnInterface string, select_ []string) (*MsgVpnBridgeRemoteMsgVpnResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeName"+"}", fmt.Sprintf("%v", bridgeName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeVirtualRouter"+"}", fmt.Sprintf("%v", bridgeVirtualRouter), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"remoteMsgVpnName"+"}", fmt.Sprintf("%v", remoteMsgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"remoteMsgVpnLocation"+"}", fmt.Sprintf("%v", remoteMsgVpnLocation), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"remoteMsgVpnInterface"+"}", fmt.Sprintf("%v", remoteMsgVpnInterface), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnBridgeRemoteMsgVpnResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnBridgeRemoteMsgVpn", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a list of Remote Message VPN objects.
 * Gets a list of Remote Message VPN objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: bridgeName|x|| bridgeVirtualRouter|x|| msgVpnName|x|| password||x| remoteMsgVpnInterface|x|| remoteMsgVpnLocation|x|| remoteMsgVpnName|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param bridgeName The bridgeName of the Bridge.
 * @param bridgeVirtualRouter The bridgeVirtualRouter of the Bridge.
 * @param where Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;).
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnBridgeRemoteMsgVpnsResponse
 */
func (a MsgVpnApi) GetMsgVpnBridgeRemoteMsgVpns(msgVpnName string, bridgeName string, bridgeVirtualRouter string, where []string, select_ []string) (*MsgVpnBridgeRemoteMsgVpnsResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeName"+"}", fmt.Sprintf("%v", bridgeName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeVirtualRouter"+"}", fmt.Sprintf("%v", bridgeVirtualRouter), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var whereCollectionFormat = "csv"
	localVarQueryParams.Add("where", a.Configuration.APIClient.ParameterToString(where, whereCollectionFormat))

	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnBridgeRemoteMsgVpnsResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnBridgeRemoteMsgVpns", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a Remote Subscription object.
 * Gets a Remote Subscription object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: bridgeName|x|| bridgeVirtualRouter|x|| msgVpnName|x|| remoteSubscriptionTopic|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param bridgeName The bridgeName of the Bridge.
 * @param bridgeVirtualRouter The bridgeVirtualRouter of the Bridge.
 * @param remoteSubscriptionTopic The remoteSubscriptionTopic of the Remote Subscription.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnBridgeRemoteSubscriptionResponse
 */
func (a MsgVpnApi) GetMsgVpnBridgeRemoteSubscription(msgVpnName string, bridgeName string, bridgeVirtualRouter string, remoteSubscriptionTopic string, select_ []string) (*MsgVpnBridgeRemoteSubscriptionResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions/{remoteSubscriptionTopic}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeName"+"}", fmt.Sprintf("%v", bridgeName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeVirtualRouter"+"}", fmt.Sprintf("%v", bridgeVirtualRouter), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"remoteSubscriptionTopic"+"}", fmt.Sprintf("%v", remoteSubscriptionTopic), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnBridgeRemoteSubscriptionResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnBridgeRemoteSubscription", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a list of Remote Subscription objects.
 * Gets a list of Remote Subscription objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: bridgeName|x|| bridgeVirtualRouter|x|| msgVpnName|x|| remoteSubscriptionTopic|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param bridgeName The bridgeName of the Bridge.
 * @param bridgeVirtualRouter The bridgeVirtualRouter of the Bridge.
 * @param count Limit the count of objects in the response. See [Count](#count \&quot;Description of the syntax of the &#x60;count&#x60; parameter\&quot;).
 * @param cursor The cursor, or position, for the next page of objects. See [Cursor](#cursor \&quot;Description of the syntax of the &#x60;cursor&#x60; parameter\&quot;).
 * @param where Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;).
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnBridgeRemoteSubscriptionsResponse
 */
func (a MsgVpnApi) GetMsgVpnBridgeRemoteSubscriptions(msgVpnName string, bridgeName string, bridgeVirtualRouter string, count int32, cursor string, where []string, select_ []string) (*MsgVpnBridgeRemoteSubscriptionsResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeName"+"}", fmt.Sprintf("%v", bridgeName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeVirtualRouter"+"}", fmt.Sprintf("%v", bridgeVirtualRouter), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("count", a.Configuration.APIClient.ParameterToString(count, ""))
	localVarQueryParams.Add("cursor", a.Configuration.APIClient.ParameterToString(cursor, ""))
	var whereCollectionFormat = "csv"
	localVarQueryParams.Add("where", a.Configuration.APIClient.ParameterToString(where, whereCollectionFormat))

	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnBridgeRemoteSubscriptionsResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnBridgeRemoteSubscriptions", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a Trusted Common Name object.
 * Gets a Trusted Common Name object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: bridgeName|x|| bridgeVirtualRouter|x|| msgVpnName|x|| tlsTrustedCommonName|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param bridgeName The bridgeName of the Bridge.
 * @param bridgeVirtualRouter The bridgeVirtualRouter of the Bridge.
 * @param tlsTrustedCommonName The tlsTrustedCommonName of the Trusted Common Name.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnBridgeTlsTrustedCommonNameResponse
 */
func (a MsgVpnApi) GetMsgVpnBridgeTlsTrustedCommonName(msgVpnName string, bridgeName string, bridgeVirtualRouter string, tlsTrustedCommonName string, select_ []string) (*MsgVpnBridgeTlsTrustedCommonNameResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames/{tlsTrustedCommonName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeName"+"}", fmt.Sprintf("%v", bridgeName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeVirtualRouter"+"}", fmt.Sprintf("%v", bridgeVirtualRouter), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tlsTrustedCommonName"+"}", fmt.Sprintf("%v", tlsTrustedCommonName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnBridgeTlsTrustedCommonNameResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnBridgeTlsTrustedCommonName", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a list of Trusted Common Name objects.
 * Gets a list of Trusted Common Name objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: bridgeName|x|| bridgeVirtualRouter|x|| msgVpnName|x|| tlsTrustedCommonName|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param bridgeName The bridgeName of the Bridge.
 * @param bridgeVirtualRouter The bridgeVirtualRouter of the Bridge.
 * @param where Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;).
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnBridgeTlsTrustedCommonNamesResponse
 */
func (a MsgVpnApi) GetMsgVpnBridgeTlsTrustedCommonNames(msgVpnName string, bridgeName string, bridgeVirtualRouter string, where []string, select_ []string) (*MsgVpnBridgeTlsTrustedCommonNamesResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeName"+"}", fmt.Sprintf("%v", bridgeName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeVirtualRouter"+"}", fmt.Sprintf("%v", bridgeVirtualRouter), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var whereCollectionFormat = "csv"
	localVarQueryParams.Add("where", a.Configuration.APIClient.ParameterToString(where, whereCollectionFormat))

	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnBridgeTlsTrustedCommonNamesResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnBridgeTlsTrustedCommonNames", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a list of Bridge objects.
 * Gets a list of Bridge objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: bridgeName|x|| bridgeVirtualRouter|x|| msgVpnName|x|| remoteAuthenticationBasicPassword||x|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param count Limit the count of objects in the response. See [Count](#count \&quot;Description of the syntax of the &#x60;count&#x60; parameter\&quot;).
 * @param cursor The cursor, or position, for the next page of objects. See [Cursor](#cursor \&quot;Description of the syntax of the &#x60;cursor&#x60; parameter\&quot;).
 * @param where Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;).
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnBridgesResponse
 */
func (a MsgVpnApi) GetMsgVpnBridges(msgVpnName string, count int32, cursor string, where []string, select_ []string) (*MsgVpnBridgesResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/bridges"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("count", a.Configuration.APIClient.ParameterToString(count, ""))
	localVarQueryParams.Add("cursor", a.Configuration.APIClient.ParameterToString(cursor, ""))
	var whereCollectionFormat = "csv"
	localVarQueryParams.Add("where", a.Configuration.APIClient.ParameterToString(where, whereCollectionFormat))

	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnBridgesResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnBridges", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a Client Profile object.
 * Gets a Client Profile object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: clientProfileName|x|| msgVpnName|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param clientProfileName The clientProfileName of the Client Profile.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnClientProfileResponse
 */
func (a MsgVpnApi) GetMsgVpnClientProfile(msgVpnName string, clientProfileName string, select_ []string) (*MsgVpnClientProfileResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/clientProfiles/{clientProfileName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientProfileName"+"}", fmt.Sprintf("%v", clientProfileName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnClientProfileResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnClientProfile", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a list of Client Profile objects.
 * Gets a list of Client Profile objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: clientProfileName|x|| msgVpnName|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param count Limit the count of objects in the response. See [Count](#count \&quot;Description of the syntax of the &#x60;count&#x60; parameter\&quot;).
 * @param cursor The cursor, or position, for the next page of objects. See [Cursor](#cursor \&quot;Description of the syntax of the &#x60;cursor&#x60; parameter\&quot;).
 * @param where Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;).
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnClientProfilesResponse
 */
func (a MsgVpnApi) GetMsgVpnClientProfiles(msgVpnName string, count int32, cursor string, where []string, select_ []string) (*MsgVpnClientProfilesResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/clientProfiles"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("count", a.Configuration.APIClient.ParameterToString(count, ""))
	localVarQueryParams.Add("cursor", a.Configuration.APIClient.ParameterToString(cursor, ""))
	var whereCollectionFormat = "csv"
	localVarQueryParams.Add("where", a.Configuration.APIClient.ParameterToString(where, whereCollectionFormat))

	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnClientProfilesResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnClientProfiles", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a Client Username object.
 * Gets a Client Username object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: clientUsername|x|| msgVpnName|x|| password||x|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param clientUsername The clientUsername of the Client Username.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnClientUsernameResponse
 */
func (a MsgVpnApi) GetMsgVpnClientUsername(msgVpnName string, clientUsername string, select_ []string) (*MsgVpnClientUsernameResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/clientUsernames/{clientUsername}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientUsername"+"}", fmt.Sprintf("%v", clientUsername), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnClientUsernameResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnClientUsername", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a list of Client Username objects.
 * Gets a list of Client Username objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: clientUsername|x|| msgVpnName|x|| password||x|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param count Limit the count of objects in the response. See [Count](#count \&quot;Description of the syntax of the &#x60;count&#x60; parameter\&quot;).
 * @param cursor The cursor, or position, for the next page of objects. See [Cursor](#cursor \&quot;Description of the syntax of the &#x60;cursor&#x60; parameter\&quot;).
 * @param where Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;).
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnClientUsernamesResponse
 */
func (a MsgVpnApi) GetMsgVpnClientUsernames(msgVpnName string, count int32, cursor string, where []string, select_ []string) (*MsgVpnClientUsernamesResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/clientUsernames"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("count", a.Configuration.APIClient.ParameterToString(count, ""))
	localVarQueryParams.Add("cursor", a.Configuration.APIClient.ParameterToString(cursor, ""))
	var whereCollectionFormat = "csv"
	localVarQueryParams.Add("where", a.Configuration.APIClient.ParameterToString(where, whereCollectionFormat))

	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnClientUsernamesResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnClientUsernames", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a list of JNDI Connection Factory objects.
 * Gets a list of JNDI Connection Factory objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: connectionFactoryName|x|| msgVpnName|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.2.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param count Limit the count of objects in the response. See [Count](#count \&quot;Description of the syntax of the &#x60;count&#x60; parameter\&quot;).
 * @param cursor The cursor, or position, for the next page of objects. See [Cursor](#cursor \&quot;Description of the syntax of the &#x60;cursor&#x60; parameter\&quot;).
 * @param where Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;).
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnJndiConnectionFactoriesResponse
 */
func (a MsgVpnApi) GetMsgVpnJndiConnectionFactories(msgVpnName string, count int32, cursor string, where []string, select_ []string) (*MsgVpnJndiConnectionFactoriesResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/jndiConnectionFactories"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("count", a.Configuration.APIClient.ParameterToString(count, ""))
	localVarQueryParams.Add("cursor", a.Configuration.APIClient.ParameterToString(cursor, ""))
	var whereCollectionFormat = "csv"
	localVarQueryParams.Add("where", a.Configuration.APIClient.ParameterToString(where, whereCollectionFormat))

	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnJndiConnectionFactoriesResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnJndiConnectionFactories", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a JNDI Connection Factory object.
 * Gets a JNDI Connection Factory object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: connectionFactoryName|x|| msgVpnName|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.2.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param connectionFactoryName The connectionFactoryName of the JNDI Connection Factory.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnJndiConnectionFactoryResponse
 */
func (a MsgVpnApi) GetMsgVpnJndiConnectionFactory(msgVpnName string, connectionFactoryName string, select_ []string) (*MsgVpnJndiConnectionFactoryResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"connectionFactoryName"+"}", fmt.Sprintf("%v", connectionFactoryName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnJndiConnectionFactoryResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnJndiConnectionFactory", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a JNDI Queue object.
 * Gets a JNDI Queue object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: msgVpnName|x|| queueName|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.2.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param queueName The queueName of the JNDI Queue.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnJndiQueueResponse
 */
func (a MsgVpnApi) GetMsgVpnJndiQueue(msgVpnName string, queueName string, select_ []string) (*MsgVpnJndiQueueResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/jndiQueues/{queueName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueName"+"}", fmt.Sprintf("%v", queueName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnJndiQueueResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnJndiQueue", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a list of JNDI Queue objects.
 * Gets a list of JNDI Queue objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: msgVpnName|x|| queueName|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.2.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param where Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;).
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnJndiQueuesResponse
 */
func (a MsgVpnApi) GetMsgVpnJndiQueues(msgVpnName string, where []string, select_ []string) (*MsgVpnJndiQueuesResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/jndiQueues"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var whereCollectionFormat = "csv"
	localVarQueryParams.Add("where", a.Configuration.APIClient.ParameterToString(where, whereCollectionFormat))

	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnJndiQueuesResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnJndiQueues", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a JNDI Topic object.
 * Gets a JNDI Topic object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: msgVpnName|x|| topicName|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.2.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param topicName The topicName of the JNDI Topic.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnJndiTopicResponse
 */
func (a MsgVpnApi) GetMsgVpnJndiTopic(msgVpnName string, topicName string, select_ []string) (*MsgVpnJndiTopicResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/jndiTopics/{topicName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topicName"+"}", fmt.Sprintf("%v", topicName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnJndiTopicResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnJndiTopic", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a list of JNDI Topic objects.
 * Gets a list of JNDI Topic objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: msgVpnName|x|| topicName|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.2.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param count Limit the count of objects in the response. See [Count](#count \&quot;Description of the syntax of the &#x60;count&#x60; parameter\&quot;).
 * @param cursor The cursor, or position, for the next page of objects. See [Cursor](#cursor \&quot;Description of the syntax of the &#x60;cursor&#x60; parameter\&quot;).
 * @param where Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;).
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnJndiTopicsResponse
 */
func (a MsgVpnApi) GetMsgVpnJndiTopics(msgVpnName string, count int32, cursor string, where []string, select_ []string) (*MsgVpnJndiTopicsResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/jndiTopics"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("count", a.Configuration.APIClient.ParameterToString(count, ""))
	localVarQueryParams.Add("cursor", a.Configuration.APIClient.ParameterToString(cursor, ""))
	var whereCollectionFormat = "csv"
	localVarQueryParams.Add("where", a.Configuration.APIClient.ParameterToString(where, whereCollectionFormat))

	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnJndiTopicsResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnJndiTopics", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets an MQTT Session object.
 * Gets an MQTT Session object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: mqttSessionClientId|x|| mqttSessionVirtualRouter|x|| msgVpnName|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.1.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param mqttSessionClientId The mqttSessionClientId of the MQTT Session.
 * @param mqttSessionVirtualRouter The mqttSessionVirtualRouter of the MQTT Session.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnMqttSessionResponse
 */
func (a MsgVpnApi) GetMsgVpnMqttSession(msgVpnName string, mqttSessionClientId string, mqttSessionVirtualRouter string, select_ []string) (*MsgVpnMqttSessionResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mqttSessionClientId"+"}", fmt.Sprintf("%v", mqttSessionClientId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mqttSessionVirtualRouter"+"}", fmt.Sprintf("%v", mqttSessionVirtualRouter), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnMqttSessionResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnMqttSession", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets an MQTT Session Subscription object.
 * Gets an MQTT Session Subscription object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: mqttSessionClientId|x|| mqttSessionVirtualRouter|x|| msgVpnName|x|| subscriptionTopic|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.1.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param mqttSessionClientId The mqttSessionClientId of the MQTT Session.
 * @param mqttSessionVirtualRouter The mqttSessionVirtualRouter of the MQTT Session.
 * @param subscriptionTopic The subscriptionTopic of the MQTT Session Subscription.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnMqttSessionSubscriptionResponse
 */
func (a MsgVpnApi) GetMsgVpnMqttSessionSubscription(msgVpnName string, mqttSessionClientId string, mqttSessionVirtualRouter string, subscriptionTopic string, select_ []string) (*MsgVpnMqttSessionSubscriptionResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mqttSessionClientId"+"}", fmt.Sprintf("%v", mqttSessionClientId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mqttSessionVirtualRouter"+"}", fmt.Sprintf("%v", mqttSessionVirtualRouter), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionTopic"+"}", fmt.Sprintf("%v", subscriptionTopic), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnMqttSessionSubscriptionResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnMqttSessionSubscription", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a list of MQTT Session Subscription objects.
 * Gets a list of MQTT Session Subscription objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: mqttSessionClientId|x|| mqttSessionVirtualRouter|x|| msgVpnName|x|| subscriptionTopic|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.1.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param mqttSessionClientId The mqttSessionClientId of the MQTT Session.
 * @param mqttSessionVirtualRouter The mqttSessionVirtualRouter of the MQTT Session.
 * @param count Limit the count of objects in the response. See [Count](#count \&quot;Description of the syntax of the &#x60;count&#x60; parameter\&quot;).
 * @param cursor The cursor, or position, for the next page of objects. See [Cursor](#cursor \&quot;Description of the syntax of the &#x60;cursor&#x60; parameter\&quot;).
 * @param where Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;).
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnMqttSessionSubscriptionsResponse
 */
func (a MsgVpnApi) GetMsgVpnMqttSessionSubscriptions(msgVpnName string, mqttSessionClientId string, mqttSessionVirtualRouter string, count int32, cursor string, where []string, select_ []string) (*MsgVpnMqttSessionSubscriptionsResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mqttSessionClientId"+"}", fmt.Sprintf("%v", mqttSessionClientId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mqttSessionVirtualRouter"+"}", fmt.Sprintf("%v", mqttSessionVirtualRouter), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("count", a.Configuration.APIClient.ParameterToString(count, ""))
	localVarQueryParams.Add("cursor", a.Configuration.APIClient.ParameterToString(cursor, ""))
	var whereCollectionFormat = "csv"
	localVarQueryParams.Add("where", a.Configuration.APIClient.ParameterToString(where, whereCollectionFormat))

	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnMqttSessionSubscriptionsResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnMqttSessionSubscriptions", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a list of MQTT Session objects.
 * Gets a list of MQTT Session objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: mqttSessionClientId|x|| mqttSessionVirtualRouter|x|| msgVpnName|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.1.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param count Limit the count of objects in the response. See [Count](#count \&quot;Description of the syntax of the &#x60;count&#x60; parameter\&quot;).
 * @param cursor The cursor, or position, for the next page of objects. See [Cursor](#cursor \&quot;Description of the syntax of the &#x60;cursor&#x60; parameter\&quot;).
 * @param where Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;).
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnMqttSessionsResponse
 */
func (a MsgVpnApi) GetMsgVpnMqttSessions(msgVpnName string, count int32, cursor string, where []string, select_ []string) (*MsgVpnMqttSessionsResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/mqttSessions"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("count", a.Configuration.APIClient.ParameterToString(count, ""))
	localVarQueryParams.Add("cursor", a.Configuration.APIClient.ParameterToString(cursor, ""))
	var whereCollectionFormat = "csv"
	localVarQueryParams.Add("where", a.Configuration.APIClient.ParameterToString(where, whereCollectionFormat))

	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnMqttSessionsResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnMqttSessions", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a Queue object.
 * Gets a Queue object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: msgVpnName|x|| queueName|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param queueName The queueName of the Queue.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnQueueResponse
 */
func (a MsgVpnApi) GetMsgVpnQueue(msgVpnName string, queueName string, select_ []string) (*MsgVpnQueueResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/queues/{queueName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueName"+"}", fmt.Sprintf("%v", queueName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnQueueResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnQueue", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a Queue Subscription object.
 * Gets a Queue Subscription object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: msgVpnName|x|| queueName|x|| subscriptionTopic|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param queueName The queueName of the Queue.
 * @param subscriptionTopic The subscriptionTopic of the Queue Subscription.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnQueueSubscriptionResponse
 */
func (a MsgVpnApi) GetMsgVpnQueueSubscription(msgVpnName string, queueName string, subscriptionTopic string, select_ []string) (*MsgVpnQueueSubscriptionResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/queues/{queueName}/subscriptions/{subscriptionTopic}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueName"+"}", fmt.Sprintf("%v", queueName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionTopic"+"}", fmt.Sprintf("%v", subscriptionTopic), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnQueueSubscriptionResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnQueueSubscription", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a list of Queue Subscription objects.
 * Gets a list of Queue Subscription objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: msgVpnName|x|| queueName|x|| subscriptionTopic|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param queueName The queueName of the Queue.
 * @param count Limit the count of objects in the response. See [Count](#count \&quot;Description of the syntax of the &#x60;count&#x60; parameter\&quot;).
 * @param cursor The cursor, or position, for the next page of objects. See [Cursor](#cursor \&quot;Description of the syntax of the &#x60;cursor&#x60; parameter\&quot;).
 * @param where Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;).
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnQueueSubscriptionsResponse
 */
func (a MsgVpnApi) GetMsgVpnQueueSubscriptions(msgVpnName string, queueName string, count int32, cursor string, where []string, select_ []string) (*MsgVpnQueueSubscriptionsResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/queues/{queueName}/subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueName"+"}", fmt.Sprintf("%v", queueName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("count", a.Configuration.APIClient.ParameterToString(count, ""))
	localVarQueryParams.Add("cursor", a.Configuration.APIClient.ParameterToString(cursor, ""))
	var whereCollectionFormat = "csv"
	localVarQueryParams.Add("where", a.Configuration.APIClient.ParameterToString(where, whereCollectionFormat))

	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnQueueSubscriptionsResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnQueueSubscriptions", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a list of Queue objects.
 * Gets a list of Queue objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: msgVpnName|x|| queueName|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param count Limit the count of objects in the response. See [Count](#count \&quot;Description of the syntax of the &#x60;count&#x60; parameter\&quot;).
 * @param cursor The cursor, or position, for the next page of objects. See [Cursor](#cursor \&quot;Description of the syntax of the &#x60;cursor&#x60; parameter\&quot;).
 * @param where Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;).
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnQueuesResponse
 */
func (a MsgVpnApi) GetMsgVpnQueues(msgVpnName string, count int32, cursor string, where []string, select_ []string) (*MsgVpnQueuesResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/queues"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("count", a.Configuration.APIClient.ParameterToString(count, ""))
	localVarQueryParams.Add("cursor", a.Configuration.APIClient.ParameterToString(cursor, ""))
	var whereCollectionFormat = "csv"
	localVarQueryParams.Add("where", a.Configuration.APIClient.ParameterToString(where, whereCollectionFormat))

	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnQueuesResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnQueues", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a Replicated Topic object.
 * Gets a Replicated Topic object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: msgVpnName|x|| replicatedTopic|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.1.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param replicatedTopic The replicatedTopic of the Replicated Topic.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnReplicatedTopicResponse
 */
func (a MsgVpnApi) GetMsgVpnReplicatedTopic(msgVpnName string, replicatedTopic string, select_ []string) (*MsgVpnReplicatedTopicResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"replicatedTopic"+"}", fmt.Sprintf("%v", replicatedTopic), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnReplicatedTopicResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnReplicatedTopic", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a list of Replicated Topic objects.
 * Gets a list of Replicated Topic objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: msgVpnName|x|| replicatedTopic|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.1.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param count Limit the count of objects in the response. See [Count](#count \&quot;Description of the syntax of the &#x60;count&#x60; parameter\&quot;).
 * @param cursor The cursor, or position, for the next page of objects. See [Cursor](#cursor \&quot;Description of the syntax of the &#x60;cursor&#x60; parameter\&quot;).
 * @param where Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;).
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnReplicatedTopicsResponse
 */
func (a MsgVpnApi) GetMsgVpnReplicatedTopics(msgVpnName string, count int32, cursor string, where []string, select_ []string) (*MsgVpnReplicatedTopicsResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/replicatedTopics"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("count", a.Configuration.APIClient.ParameterToString(count, ""))
	localVarQueryParams.Add("cursor", a.Configuration.APIClient.ParameterToString(cursor, ""))
	var whereCollectionFormat = "csv"
	localVarQueryParams.Add("where", a.Configuration.APIClient.ParameterToString(where, whereCollectionFormat))

	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnReplicatedTopicsResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnReplicatedTopics", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a REST Delivery Point object.
 * Gets a REST Delivery Point object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: msgVpnName|x|| restDeliveryPointName|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param restDeliveryPointName The restDeliveryPointName of the REST Delivery Point.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnRestDeliveryPointResponse
 */
func (a MsgVpnApi) GetMsgVpnRestDeliveryPoint(msgVpnName string, restDeliveryPointName string, select_ []string) (*MsgVpnRestDeliveryPointResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnRestDeliveryPointResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnRestDeliveryPoint", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a Queue Binding object.
 * Gets a Queue Binding object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: msgVpnName|x|| queueBindingName|x|| restDeliveryPointName|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param restDeliveryPointName The restDeliveryPointName of the REST Delivery Point.
 * @param queueBindingName The queueBindingName of the Queue Binding.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnRestDeliveryPointQueueBindingResponse
 */
func (a MsgVpnApi) GetMsgVpnRestDeliveryPointQueueBinding(msgVpnName string, restDeliveryPointName string, queueBindingName string, select_ []string) (*MsgVpnRestDeliveryPointQueueBindingResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueBindingName"+"}", fmt.Sprintf("%v", queueBindingName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnRestDeliveryPointQueueBindingResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnRestDeliveryPointQueueBinding", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a list of Queue Binding objects.
 * Gets a list of Queue Binding objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: msgVpnName|x|| queueBindingName|x|| restDeliveryPointName|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param restDeliveryPointName The restDeliveryPointName of the REST Delivery Point.
 * @param count Limit the count of objects in the response. See [Count](#count \&quot;Description of the syntax of the &#x60;count&#x60; parameter\&quot;).
 * @param cursor The cursor, or position, for the next page of objects. See [Cursor](#cursor \&quot;Description of the syntax of the &#x60;cursor&#x60; parameter\&quot;).
 * @param where Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;).
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnRestDeliveryPointQueueBindingsResponse
 */
func (a MsgVpnApi) GetMsgVpnRestDeliveryPointQueueBindings(msgVpnName string, restDeliveryPointName string, count int32, cursor string, where []string, select_ []string) (*MsgVpnRestDeliveryPointQueueBindingsResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("count", a.Configuration.APIClient.ParameterToString(count, ""))
	localVarQueryParams.Add("cursor", a.Configuration.APIClient.ParameterToString(cursor, ""))
	var whereCollectionFormat = "csv"
	localVarQueryParams.Add("where", a.Configuration.APIClient.ParameterToString(where, whereCollectionFormat))

	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnRestDeliveryPointQueueBindingsResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnRestDeliveryPointQueueBindings", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a REST Consumer object.
 * Gets a REST Consumer object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: authenticationHttpBasicPassword||x| msgVpnName|x|| restConsumerName|x|| restDeliveryPointName|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param restDeliveryPointName The restDeliveryPointName of the REST Delivery Point.
 * @param restConsumerName The restConsumerName of the REST Consumer.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnRestDeliveryPointRestConsumerResponse
 */
func (a MsgVpnApi) GetMsgVpnRestDeliveryPointRestConsumer(msgVpnName string, restDeliveryPointName string, restConsumerName string, select_ []string) (*MsgVpnRestDeliveryPointRestConsumerResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restConsumerName"+"}", fmt.Sprintf("%v", restConsumerName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnRestDeliveryPointRestConsumerResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnRestDeliveryPointRestConsumer", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a list of REST Consumer objects.
 * Gets a list of REST Consumer objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: authenticationHttpBasicPassword||x| msgVpnName|x|| restConsumerName|x|| restDeliveryPointName|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param restDeliveryPointName The restDeliveryPointName of the REST Delivery Point.
 * @param count Limit the count of objects in the response. See [Count](#count \&quot;Description of the syntax of the &#x60;count&#x60; parameter\&quot;).
 * @param cursor The cursor, or position, for the next page of objects. See [Cursor](#cursor \&quot;Description of the syntax of the &#x60;cursor&#x60; parameter\&quot;).
 * @param where Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;).
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnRestDeliveryPointRestConsumersResponse
 */
func (a MsgVpnApi) GetMsgVpnRestDeliveryPointRestConsumers(msgVpnName string, restDeliveryPointName string, count int32, cursor string, where []string, select_ []string) (*MsgVpnRestDeliveryPointRestConsumersResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("count", a.Configuration.APIClient.ParameterToString(count, ""))
	localVarQueryParams.Add("cursor", a.Configuration.APIClient.ParameterToString(cursor, ""))
	var whereCollectionFormat = "csv"
	localVarQueryParams.Add("where", a.Configuration.APIClient.ParameterToString(where, whereCollectionFormat))

	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnRestDeliveryPointRestConsumersResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnRestDeliveryPointRestConsumers", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a list of REST Delivery Point objects.
 * Gets a list of REST Delivery Point objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: msgVpnName|x|| restDeliveryPointName|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param count Limit the count of objects in the response. See [Count](#count \&quot;Description of the syntax of the &#x60;count&#x60; parameter\&quot;).
 * @param cursor The cursor, or position, for the next page of objects. See [Cursor](#cursor \&quot;Description of the syntax of the &#x60;cursor&#x60; parameter\&quot;).
 * @param where Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;).
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnRestDeliveryPointsResponse
 */
func (a MsgVpnApi) GetMsgVpnRestDeliveryPoints(msgVpnName string, count int32, cursor string, where []string, select_ []string) (*MsgVpnRestDeliveryPointsResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("count", a.Configuration.APIClient.ParameterToString(count, ""))
	localVarQueryParams.Add("cursor", a.Configuration.APIClient.ParameterToString(cursor, ""))
	var whereCollectionFormat = "csv"
	localVarQueryParams.Add("where", a.Configuration.APIClient.ParameterToString(where, whereCollectionFormat))

	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnRestDeliveryPointsResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnRestDeliveryPoints", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a Sequenced Topic object.
 * Gets a Sequenced Topic object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: msgVpnName|x|| sequencedTopic|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param sequencedTopic The sequencedTopic of the Sequenced Topic.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnSequencedTopicResponse
 */
func (a MsgVpnApi) GetMsgVpnSequencedTopic(msgVpnName string, sequencedTopic string, select_ []string) (*MsgVpnSequencedTopicResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/sequencedTopics/{sequencedTopic}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sequencedTopic"+"}", fmt.Sprintf("%v", sequencedTopic), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnSequencedTopicResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnSequencedTopic", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a list of Sequenced Topic objects.
 * Gets a list of Sequenced Topic objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: msgVpnName|x|| sequencedTopic|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param count Limit the count of objects in the response. See [Count](#count \&quot;Description of the syntax of the &#x60;count&#x60; parameter\&quot;).
 * @param cursor The cursor, or position, for the next page of objects. See [Cursor](#cursor \&quot;Description of the syntax of the &#x60;cursor&#x60; parameter\&quot;).
 * @param where Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;).
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnSequencedTopicsResponse
 */
func (a MsgVpnApi) GetMsgVpnSequencedTopics(msgVpnName string, count int32, cursor string, where []string, select_ []string) (*MsgVpnSequencedTopicsResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/sequencedTopics"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("count", a.Configuration.APIClient.ParameterToString(count, ""))
	localVarQueryParams.Add("cursor", a.Configuration.APIClient.ParameterToString(cursor, ""))
	var whereCollectionFormat = "csv"
	localVarQueryParams.Add("where", a.Configuration.APIClient.ParameterToString(where, whereCollectionFormat))

	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnSequencedTopicsResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnSequencedTopics", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a Topic Endpoint object.
 * Gets a Topic Endpoint object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: msgVpnName|x|| topicEndpointName|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.1.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param topicEndpointName The topicEndpointName of the Topic Endpoint.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnTopicEndpointResponse
 */
func (a MsgVpnApi) GetMsgVpnTopicEndpoint(msgVpnName string, topicEndpointName string, select_ []string) (*MsgVpnTopicEndpointResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topicEndpointName"+"}", fmt.Sprintf("%v", topicEndpointName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnTopicEndpointResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnTopicEndpoint", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a list of Topic Endpoint objects.
 * Gets a list of Topic Endpoint objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: msgVpnName|x|| topicEndpointName|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.1.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param count Limit the count of objects in the response. See [Count](#count \&quot;Description of the syntax of the &#x60;count&#x60; parameter\&quot;).
 * @param cursor The cursor, or position, for the next page of objects. See [Cursor](#cursor \&quot;Description of the syntax of the &#x60;cursor&#x60; parameter\&quot;).
 * @param where Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;).
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnTopicEndpointsResponse
 */
func (a MsgVpnApi) GetMsgVpnTopicEndpoints(msgVpnName string, count int32, cursor string, where []string, select_ []string) (*MsgVpnTopicEndpointsResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/topicEndpoints"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("count", a.Configuration.APIClient.ParameterToString(count, ""))
	localVarQueryParams.Add("cursor", a.Configuration.APIClient.ParameterToString(cursor, ""))
	var whereCollectionFormat = "csv"
	localVarQueryParams.Add("where", a.Configuration.APIClient.ParameterToString(where, whereCollectionFormat))

	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnTopicEndpointsResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpnTopicEndpoints", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Gets a list of Message VPN objects.
 * Gets a list of Message VPN objects.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: msgVpnName|x|| replicationBridgeAuthenticationBasicPassword||x| replicationEnabledQueueBehavior||x|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param count Limit the count of objects in the response. See [Count](#count \&quot;Description of the syntax of the &#x60;count&#x60; parameter\&quot;).
 * @param cursor The cursor, or position, for the next page of objects. See [Cursor](#cursor \&quot;Description of the syntax of the &#x60;cursor&#x60; parameter\&quot;).
 * @param where Include in the response only objects where certain conditions are true. See [Where](#where \&quot;Description of the syntax of the &#x60;where&#x60; parameter\&quot;).
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnsResponse
 */
func (a MsgVpnApi) GetMsgVpns(count int32, cursor string, where []string, select_ []string) (*MsgVpnsResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("count", a.Configuration.APIClient.ParameterToString(count, ""))
	localVarQueryParams.Add("cursor", a.Configuration.APIClient.ParameterToString(cursor, ""))
	var whereCollectionFormat = "csv"
	localVarQueryParams.Add("where", a.Configuration.APIClient.ParameterToString(where, whereCollectionFormat))

	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(MsgVpnsResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetMsgVpns", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Replaces a Message VPN object.
 * Replaces a Message VPN object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: msgVpnName|x|x||| replicationBridgeAuthenticationBasicPassword|||x|| replicationEnabledQueueBehavior|||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent EventThresholdByValue|clearValue|setValue| EventThresholdByValue|setValue|clearValue| MsgVpn|authenticationBasicProfileName|authenticationBasicType| MsgVpn|authorizationProfileName|authorizationType| MsgVpn|eventPublishTopicFormatMqttEnabled|eventPublishTopicFormatSmfEnabled| MsgVpn|eventPublishTopicFormatSmfEnabled|eventPublishTopicFormatMqttEnabled| MsgVpn|replicationBridgeAuthenticationBasicClientUsername|replicationBridgeAuthenticationBasicPassword| MsgVpn|replicationBridgeAuthenticationBasicPassword|replicationBridgeAuthenticationBasicClientUsername| MsgVpn|replicationEnabledQueueBehavior|replicationEnabled|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation. Requests which include the following attributes require greater access scope/level:   Attribute|Access Scope/Level :---|:---: authenticationBasicEnabled|global/readwrite authenticationBasicProfileName|global/readwrite authenticationBasicRadiusDomain|global/readwrite authenticationBasicType|global/readwrite authenticationClientCertAllowApiProvidedUsernameEnabled|global/readwrite authenticationClientCertEnabled|global/readwrite authenticationClientCertMaxChainDepth|global/readwrite authenticationClientCertValidateDateEnabled|global/readwrite authenticationKerberosAllowApiProvidedUsernameEnabled|global/readwrite authenticationKerberosEnabled|global/readwrite bridgingTlsServerCertEnforceTrustedCommonNameEnabled|global/readwrite bridgingTlsServerCertMaxChainDepth|global/readwrite bridgingTlsServerCertValidateDateEnabled|global/readwrite exportSubscriptionsEnabled|global/readwrite maxConnectionCount|global/readwrite maxEgressFlowCount|global/readwrite maxEndpointCount|global/readwrite maxIngressFlowCount|global/readwrite maxMsgSpoolUsage|global/readwrite maxSubscriptionCount|global/readwrite maxTransactedSessionCount|global/readwrite maxTransactionCount|global/readwrite replicationBridgeAuthenticationBasicClientUsername|global/readwrite replicationBridgeAuthenticationBasicPassword|global/readwrite replicationBridgeAuthenticationScheme|global/readwrite replicationBridgeCompressedDataEnabled|global/readwrite replicationBridgeEgressFlowWindowSize|global/readwrite replicationBridgeRetryDelay|global/readwrite replicationBridgeTlsEnabled|global/readwrite replicationBridgeUnidirectionalClientProfileName|global/readwrite replicationEnabled|global/readwrite replicationEnabledQueueBehavior|global/readwrite replicationQueueMaxMsgSpoolUsage|global/readwrite replicationRole|global/readwrite restTlsServerCertEnforceTrustedCommonNameEnabled|global/readwrite restTlsServerCertMaxChainDepth|global/readwrite restTlsServerCertValidateDateEnabled|global/readwrite sempOverMsgBusAdminClientEnabled|global/readwrite sempOverMsgBusAdminDistributedCacheEnabled|global/readwrite sempOverMsgBusAdminEnabled|global/readwrite sempOverMsgBusEnabled|global/readwrite sempOverMsgBusShowEnabled|global/readwrite serviceRestIncomingMaxConnectionCount|global/readwrite serviceRestIncomingPlainTextListenPort|global/readwrite serviceRestIncomingTlsListenPort|global/readwrite serviceRestOutgoingMaxConnectionCount|global/readwrite serviceSmfMaxConnectionCount|global/readwrite serviceWebMaxConnectionCount|global/readwrite    This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param body The Message VPN object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnResponse
 */
func (a MsgVpnApi) ReplaceMsgVpn(msgVpnName string, body MsgVpn, select_ []string) (*MsgVpnResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Put")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ReplaceMsgVpn", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Replaces an ACL Profile object.
 * Replaces an ACL Profile object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.  ACL profiles are a collection of user-defined profiles for specifying Access Control List properties.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: aclProfileName|x|x||| msgVpnName|x|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param aclProfileName The aclProfileName of the ACL Profile.
 * @param body The ACL Profile object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnAclProfileResponse
 */
func (a MsgVpnApi) ReplaceMsgVpnAclProfile(msgVpnName string, aclProfileName string, body MsgVpnAclProfile, select_ []string) (*MsgVpnAclProfileResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Put")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"aclProfileName"+"}", fmt.Sprintf("%v", aclProfileName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnAclProfileResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ReplaceMsgVpnAclProfile", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Replaces an LDAP Authorization Group object.
 * Replaces an LDAP Authorization Group object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: aclProfileName||||x| authorizationGroupName|x|x||| clientProfileName||||x| msgVpnName|x|x||| orderAfterAuthorizationGroupName|||x|| orderBeforeAuthorizationGroupName|||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnAuthorizationGroup|orderAfterAuthorizationGroupName||orderBeforeAuthorizationGroupName MsgVpnAuthorizationGroup|orderBeforeAuthorizationGroupName||orderAfterAuthorizationGroupName    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param authorizationGroupName The authorizationGroupName of the LDAP Authorization Group.
 * @param body The LDAP Authorization Group object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnAuthorizationGroupResponse
 */
func (a MsgVpnApi) ReplaceMsgVpnAuthorizationGroup(msgVpnName string, authorizationGroupName string, body MsgVpnAuthorizationGroup, select_ []string) (*MsgVpnAuthorizationGroupResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Put")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"authorizationGroupName"+"}", fmt.Sprintf("%v", authorizationGroupName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnAuthorizationGroupResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ReplaceMsgVpnAuthorizationGroup", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Replaces a Bridge object.
 * Replaces a Bridge object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: bridgeName|x|x||| bridgeVirtualRouter|x|x||| maxTtl||||x| msgVpnName|x|x||| remoteAuthenticationBasicClientUsername||||x| remoteAuthenticationBasicPassword|||x|x| remoteAuthenticationScheme||||x| remoteDeliverToOnePriority||||x|    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnBridge|remoteAuthenticationBasicClientUsername|remoteAuthenticationBasicPassword| MsgVpnBridge|remoteAuthenticationBasicPassword|remoteAuthenticationBasicClientUsername|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation. Requests which include the following attributes require greater access scope/level:   Attribute|Access Scope/Level :---|:---: tlsCipherSuiteList|global/readwrite    This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param bridgeName The bridgeName of the Bridge.
 * @param bridgeVirtualRouter The bridgeVirtualRouter of the Bridge.
 * @param body The Bridge object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnBridgeResponse
 */
func (a MsgVpnApi) ReplaceMsgVpnBridge(msgVpnName string, bridgeName string, bridgeVirtualRouter string, body MsgVpnBridge, select_ []string) (*MsgVpnBridgeResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Put")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeName"+"}", fmt.Sprintf("%v", bridgeName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeVirtualRouter"+"}", fmt.Sprintf("%v", bridgeVirtualRouter), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnBridgeResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ReplaceMsgVpnBridge", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Replaces a Remote Message VPN object.
 * Replaces a Remote Message VPN object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: bridgeName|x|x||| bridgeVirtualRouter|x|x||| clientUsername||||x| compressedDataEnabled||||x| egressFlowWindowSize||||x| msgVpnName|x|x||| password|||x|x| remoteMsgVpnInterface|x|x||| remoteMsgVpnLocation|x|x||| remoteMsgVpnName|x|x||| tlsEnabled||||x|    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnBridgeRemoteMsgVpn|clientUsername|password| MsgVpnBridgeRemoteMsgVpn|password|clientUsername|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param bridgeName The bridgeName of the Bridge.
 * @param bridgeVirtualRouter The bridgeVirtualRouter of the Bridge.
 * @param remoteMsgVpnName The remoteMsgVpnName of the Remote Message VPN.
 * @param remoteMsgVpnLocation The remoteMsgVpnLocation of the Remote Message VPN.
 * @param remoteMsgVpnInterface The remoteMsgVpnInterface of the Remote Message VPN.
 * @param body The Remote Message VPN object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnBridgeRemoteMsgVpnResponse
 */
func (a MsgVpnApi) ReplaceMsgVpnBridgeRemoteMsgVpn(msgVpnName string, bridgeName string, bridgeVirtualRouter string, remoteMsgVpnName string, remoteMsgVpnLocation string, remoteMsgVpnInterface string, body MsgVpnBridgeRemoteMsgVpn, select_ []string) (*MsgVpnBridgeRemoteMsgVpnResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Put")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeName"+"}", fmt.Sprintf("%v", bridgeName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeVirtualRouter"+"}", fmt.Sprintf("%v", bridgeVirtualRouter), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"remoteMsgVpnName"+"}", fmt.Sprintf("%v", remoteMsgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"remoteMsgVpnLocation"+"}", fmt.Sprintf("%v", remoteMsgVpnLocation), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"remoteMsgVpnInterface"+"}", fmt.Sprintf("%v", remoteMsgVpnInterface), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnBridgeRemoteMsgVpnResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ReplaceMsgVpnBridgeRemoteMsgVpn", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Replaces a Client Profile object.
 * Replaces a Client Profile object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: clientProfileName|x|x||| msgVpnName|x|x|||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent EventThresholdByPercent|clearPercent|setPercent| EventThresholdByPercent|setPercent|clearPercent|    A SEMP client authorized with a minimum access scope/level of \&quot;global/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param clientProfileName The clientProfileName of the Client Profile.
 * @param body The Client Profile object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnClientProfileResponse
 */
func (a MsgVpnApi) ReplaceMsgVpnClientProfile(msgVpnName string, clientProfileName string, body MsgVpnClientProfile, select_ []string) (*MsgVpnClientProfileResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Put")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/clientProfiles/{clientProfileName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientProfileName"+"}", fmt.Sprintf("%v", clientProfileName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnClientProfileResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ReplaceMsgVpnClientProfile", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Replaces a Client Username object.
 * Replaces a Client Username object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: aclProfileName||||x| clientProfileName||||x| clientUsername|x|x||| msgVpnName|x|x||| password|||x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param clientUsername The clientUsername of the Client Username.
 * @param body The Client Username object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnClientUsernameResponse
 */
func (a MsgVpnApi) ReplaceMsgVpnClientUsername(msgVpnName string, clientUsername string, body MsgVpnClientUsername, select_ []string) (*MsgVpnClientUsernameResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Put")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/clientUsernames/{clientUsername}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientUsername"+"}", fmt.Sprintf("%v", clientUsername), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnClientUsernameResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ReplaceMsgVpnClientUsername", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Replaces a JNDI Connection Factory object.
 * Replaces a JNDI Connection Factory object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: connectionFactoryName|x|x||| msgVpnName|x|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.2.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param connectionFactoryName The connectionFactoryName of the JNDI Connection Factory.
 * @param body The JNDI Connection Factory object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnJndiConnectionFactoryResponse
 */
func (a MsgVpnApi) ReplaceMsgVpnJndiConnectionFactory(msgVpnName string, connectionFactoryName string, body MsgVpnJndiConnectionFactory, select_ []string) (*MsgVpnJndiConnectionFactoryResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Put")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"connectionFactoryName"+"}", fmt.Sprintf("%v", connectionFactoryName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnJndiConnectionFactoryResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ReplaceMsgVpnJndiConnectionFactory", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Replaces a JNDI Queue object.
 * Replaces a JNDI Queue object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: msgVpnName|x|x||| queueName|x|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.2.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param queueName The queueName of the JNDI Queue.
 * @param body The JNDI Queue object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnJndiQueueResponse
 */
func (a MsgVpnApi) ReplaceMsgVpnJndiQueue(msgVpnName string, queueName string, body MsgVpnJndiQueue, select_ []string) (*MsgVpnJndiQueueResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Put")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/jndiQueues/{queueName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueName"+"}", fmt.Sprintf("%v", queueName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnJndiQueueResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ReplaceMsgVpnJndiQueue", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Replaces a JNDI Topic object.
 * Replaces a JNDI Topic object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: msgVpnName|x|x||| topicName|x|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.2.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param topicName The topicName of the JNDI Topic.
 * @param body The JNDI Topic object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnJndiTopicResponse
 */
func (a MsgVpnApi) ReplaceMsgVpnJndiTopic(msgVpnName string, topicName string, body MsgVpnJndiTopic, select_ []string) (*MsgVpnJndiTopicResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Put")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/jndiTopics/{topicName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topicName"+"}", fmt.Sprintf("%v", topicName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnJndiTopicResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ReplaceMsgVpnJndiTopic", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Replaces an MQTT Session object.
 * Replaces an MQTT Session object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: mqttSessionClientId|x|x||| mqttSessionVirtualRouter|x|x||| msgVpnName|x|x||| owner||||x|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.1.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param mqttSessionClientId The mqttSessionClientId of the MQTT Session.
 * @param mqttSessionVirtualRouter The mqttSessionVirtualRouter of the MQTT Session.
 * @param body The MQTT Session object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnMqttSessionResponse
 */
func (a MsgVpnApi) ReplaceMsgVpnMqttSession(msgVpnName string, mqttSessionClientId string, mqttSessionVirtualRouter string, body MsgVpnMqttSession, select_ []string) (*MsgVpnMqttSessionResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Put")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mqttSessionClientId"+"}", fmt.Sprintf("%v", mqttSessionClientId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mqttSessionVirtualRouter"+"}", fmt.Sprintf("%v", mqttSessionVirtualRouter), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnMqttSessionResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ReplaceMsgVpnMqttSession", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Replaces an MQTT Session Subscription object.
 * Replaces an MQTT Session Subscription object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: mqttSessionClientId|x|x||| mqttSessionVirtualRouter|x|x||| msgVpnName|x|x||| subscriptionTopic|x|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.1.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param mqttSessionClientId The mqttSessionClientId of the MQTT Session.
 * @param mqttSessionVirtualRouter The mqttSessionVirtualRouter of the MQTT Session.
 * @param subscriptionTopic The subscriptionTopic of the MQTT Session Subscription.
 * @param body The MQTT Session Subscription object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnMqttSessionSubscriptionResponse
 */
func (a MsgVpnApi) ReplaceMsgVpnMqttSessionSubscription(msgVpnName string, mqttSessionClientId string, mqttSessionVirtualRouter string, subscriptionTopic string, body MsgVpnMqttSessionSubscription, select_ []string) (*MsgVpnMqttSessionSubscriptionResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Put")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mqttSessionClientId"+"}", fmt.Sprintf("%v", mqttSessionClientId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mqttSessionVirtualRouter"+"}", fmt.Sprintf("%v", mqttSessionVirtualRouter), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionTopic"+"}", fmt.Sprintf("%v", subscriptionTopic), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnMqttSessionSubscriptionResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ReplaceMsgVpnMqttSessionSubscription", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Replaces a Queue object.
 * Replaces a Queue object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: accessType||||x| msgVpnName|x|x||| owner||||x| permission||||x| queueName|x|x|||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param queueName The queueName of the Queue.
 * @param body The Queue object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnQueueResponse
 */
func (a MsgVpnApi) ReplaceMsgVpnQueue(msgVpnName string, queueName string, body MsgVpnQueue, select_ []string) (*MsgVpnQueueResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Put")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/queues/{queueName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueName"+"}", fmt.Sprintf("%v", queueName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnQueueResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ReplaceMsgVpnQueue", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Replaces a Replicated Topic object.
 * Replaces a Replicated Topic object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: msgVpnName|x|x||| replicatedTopic|x|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.1.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param replicatedTopic The replicatedTopic of the Replicated Topic.
 * @param body The Replicated Topic object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnReplicatedTopicResponse
 */
func (a MsgVpnApi) ReplaceMsgVpnReplicatedTopic(msgVpnName string, replicatedTopic string, body MsgVpnReplicatedTopic, select_ []string) (*MsgVpnReplicatedTopicResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Put")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"replicatedTopic"+"}", fmt.Sprintf("%v", replicatedTopic), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnReplicatedTopicResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ReplaceMsgVpnReplicatedTopic", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Replaces a REST Delivery Point object.
 * Replaces a REST Delivery Point object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: clientProfileName||||x| msgVpnName|x|x||| restDeliveryPointName|x|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param restDeliveryPointName The restDeliveryPointName of the REST Delivery Point.
 * @param body The REST Delivery Point object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnRestDeliveryPointResponse
 */
func (a MsgVpnApi) ReplaceMsgVpnRestDeliveryPoint(msgVpnName string, restDeliveryPointName string, body MsgVpnRestDeliveryPoint, select_ []string) (*MsgVpnRestDeliveryPointResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Put")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnRestDeliveryPointResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ReplaceMsgVpnRestDeliveryPoint", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Replaces a Queue Binding object.
 * Replaces a Queue Binding object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: msgVpnName|x|x||| queueBindingName|x|x||| restDeliveryPointName|x|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param restDeliveryPointName The restDeliveryPointName of the REST Delivery Point.
 * @param queueBindingName The queueBindingName of the Queue Binding.
 * @param body The Queue Binding object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnRestDeliveryPointQueueBindingResponse
 */
func (a MsgVpnApi) ReplaceMsgVpnRestDeliveryPointQueueBinding(msgVpnName string, restDeliveryPointName string, queueBindingName string, body MsgVpnRestDeliveryPointQueueBinding, select_ []string) (*MsgVpnRestDeliveryPointQueueBindingResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Put")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueBindingName"+"}", fmt.Sprintf("%v", queueBindingName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnRestDeliveryPointQueueBindingResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ReplaceMsgVpnRestDeliveryPointQueueBinding", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Replaces a REST Consumer object.
 * Replaces a REST Consumer object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: authenticationHttpBasicPassword|||x|x| authenticationHttpBasicUsername||||x| authenticationScheme||||x| msgVpnName|x|x||| outgoingConnectionCount||||x| remoteHost||||x| remotePort||||x| restConsumerName|x|x||| restDeliveryPointName|x|x||| tlsCipherSuiteList||||x| tlsEnabled||||x|    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnRestDeliveryPointRestConsumer|authenticationHttpBasicPassword|authenticationHttpBasicUsername| MsgVpnRestDeliveryPointRestConsumer|authenticationHttpBasicUsername|authenticationHttpBasicPassword| MsgVpnRestDeliveryPointRestConsumer|remotePort|tlsEnabled| MsgVpnRestDeliveryPointRestConsumer|tlsEnabled|remotePort|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param restDeliveryPointName The restDeliveryPointName of the REST Delivery Point.
 * @param restConsumerName The restConsumerName of the REST Consumer.
 * @param body The REST Consumer object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnRestDeliveryPointRestConsumerResponse
 */
func (a MsgVpnApi) ReplaceMsgVpnRestDeliveryPointRestConsumer(msgVpnName string, restDeliveryPointName string, restConsumerName string, body MsgVpnRestDeliveryPointRestConsumer, select_ []string) (*MsgVpnRestDeliveryPointRestConsumerResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Put")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restConsumerName"+"}", fmt.Sprintf("%v", restConsumerName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnRestDeliveryPointRestConsumerResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ReplaceMsgVpnRestDeliveryPointRestConsumer", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Replaces a Topic Endpoint object.
 * Replaces a Topic Endpoint object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: msgVpnName|x|x||| owner||||x| permission||||x| topicEndpointName|x|x|||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.1.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param topicEndpointName The topicEndpointName of the Topic Endpoint.
 * @param body The Topic Endpoint object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnTopicEndpointResponse
 */
func (a MsgVpnApi) ReplaceMsgVpnTopicEndpoint(msgVpnName string, topicEndpointName string, body MsgVpnTopicEndpoint, select_ []string) (*MsgVpnTopicEndpointResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Put")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topicEndpointName"+"}", fmt.Sprintf("%v", topicEndpointName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnTopicEndpointResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ReplaceMsgVpnTopicEndpoint", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Updates a Message VPN object.
 * Updates a Message VPN object. Any attribute missing from the request will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: msgVpnName|x|x||| replicationBridgeAuthenticationBasicPassword|||x|| replicationEnabledQueueBehavior|||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent EventThresholdByValue|clearValue|setValue| EventThresholdByValue|setValue|clearValue| MsgVpn|authenticationBasicProfileName|authenticationBasicType| MsgVpn|authorizationProfileName|authorizationType| MsgVpn|eventPublishTopicFormatMqttEnabled|eventPublishTopicFormatSmfEnabled| MsgVpn|eventPublishTopicFormatSmfEnabled|eventPublishTopicFormatMqttEnabled| MsgVpn|replicationBridgeAuthenticationBasicClientUsername|replicationBridgeAuthenticationBasicPassword| MsgVpn|replicationBridgeAuthenticationBasicPassword|replicationBridgeAuthenticationBasicClientUsername| MsgVpn|replicationEnabledQueueBehavior|replicationEnabled|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation. Requests which include the following attributes require greater access scope/level:   Attribute|Access Scope/Level :---|:---: authenticationBasicEnabled|global/readwrite authenticationBasicProfileName|global/readwrite authenticationBasicRadiusDomain|global/readwrite authenticationBasicType|global/readwrite authenticationClientCertAllowApiProvidedUsernameEnabled|global/readwrite authenticationClientCertEnabled|global/readwrite authenticationClientCertMaxChainDepth|global/readwrite authenticationClientCertValidateDateEnabled|global/readwrite authenticationKerberosAllowApiProvidedUsernameEnabled|global/readwrite authenticationKerberosEnabled|global/readwrite bridgingTlsServerCertEnforceTrustedCommonNameEnabled|global/readwrite bridgingTlsServerCertMaxChainDepth|global/readwrite bridgingTlsServerCertValidateDateEnabled|global/readwrite exportSubscriptionsEnabled|global/readwrite maxConnectionCount|global/readwrite maxEgressFlowCount|global/readwrite maxEndpointCount|global/readwrite maxIngressFlowCount|global/readwrite maxMsgSpoolUsage|global/readwrite maxSubscriptionCount|global/readwrite maxTransactedSessionCount|global/readwrite maxTransactionCount|global/readwrite replicationBridgeAuthenticationBasicClientUsername|global/readwrite replicationBridgeAuthenticationBasicPassword|global/readwrite replicationBridgeAuthenticationScheme|global/readwrite replicationBridgeCompressedDataEnabled|global/readwrite replicationBridgeEgressFlowWindowSize|global/readwrite replicationBridgeRetryDelay|global/readwrite replicationBridgeTlsEnabled|global/readwrite replicationBridgeUnidirectionalClientProfileName|global/readwrite replicationEnabled|global/readwrite replicationEnabledQueueBehavior|global/readwrite replicationQueueMaxMsgSpoolUsage|global/readwrite replicationRole|global/readwrite restTlsServerCertEnforceTrustedCommonNameEnabled|global/readwrite restTlsServerCertMaxChainDepth|global/readwrite restTlsServerCertValidateDateEnabled|global/readwrite sempOverMsgBusAdminClientEnabled|global/readwrite sempOverMsgBusAdminDistributedCacheEnabled|global/readwrite sempOverMsgBusAdminEnabled|global/readwrite sempOverMsgBusEnabled|global/readwrite sempOverMsgBusShowEnabled|global/readwrite serviceRestIncomingMaxConnectionCount|global/readwrite serviceRestIncomingPlainTextListenPort|global/readwrite serviceRestIncomingTlsListenPort|global/readwrite serviceRestOutgoingMaxConnectionCount|global/readwrite serviceSmfMaxConnectionCount|global/readwrite serviceWebMaxConnectionCount|global/readwrite    This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param body The Message VPN object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnResponse
 */
func (a MsgVpnApi) UpdateMsgVpn(msgVpnName string, body MsgVpn, select_ []string) (*MsgVpnResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Patch")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "UpdateMsgVpn", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Updates an ACL Profile object.
 * Updates an ACL Profile object. Any attribute missing from the request will be left unchanged.  ACL profiles are a collection of user-defined profiles for specifying Access Control List properties.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: aclProfileName|x|x||| msgVpnName|x|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param aclProfileName The aclProfileName of the ACL Profile.
 * @param body The ACL Profile object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnAclProfileResponse
 */
func (a MsgVpnApi) UpdateMsgVpnAclProfile(msgVpnName string, aclProfileName string, body MsgVpnAclProfile, select_ []string) (*MsgVpnAclProfileResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Patch")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"aclProfileName"+"}", fmt.Sprintf("%v", aclProfileName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnAclProfileResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "UpdateMsgVpnAclProfile", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Updates an LDAP Authorization Group object.
 * Updates an LDAP Authorization Group object. Any attribute missing from the request will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: aclProfileName||||x| authorizationGroupName|x|x||| clientProfileName||||x| msgVpnName|x|x||| orderAfterAuthorizationGroupName|||x|| orderBeforeAuthorizationGroupName|||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnAuthorizationGroup|orderAfterAuthorizationGroupName||orderBeforeAuthorizationGroupName MsgVpnAuthorizationGroup|orderBeforeAuthorizationGroupName||orderAfterAuthorizationGroupName    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param authorizationGroupName The authorizationGroupName of the LDAP Authorization Group.
 * @param body The LDAP Authorization Group object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnAuthorizationGroupResponse
 */
func (a MsgVpnApi) UpdateMsgVpnAuthorizationGroup(msgVpnName string, authorizationGroupName string, body MsgVpnAuthorizationGroup, select_ []string) (*MsgVpnAuthorizationGroupResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Patch")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"authorizationGroupName"+"}", fmt.Sprintf("%v", authorizationGroupName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnAuthorizationGroupResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "UpdateMsgVpnAuthorizationGroup", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Updates a Bridge object.
 * Updates a Bridge object. Any attribute missing from the request will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: bridgeName|x|x||| bridgeVirtualRouter|x|x||| maxTtl||||x| msgVpnName|x|x||| remoteAuthenticationBasicClientUsername||||x| remoteAuthenticationBasicPassword|||x|x| remoteAuthenticationScheme||||x| remoteDeliverToOnePriority||||x|    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnBridge|remoteAuthenticationBasicClientUsername|remoteAuthenticationBasicPassword| MsgVpnBridge|remoteAuthenticationBasicPassword|remoteAuthenticationBasicClientUsername|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation. Requests which include the following attributes require greater access scope/level:   Attribute|Access Scope/Level :---|:---: tlsCipherSuiteList|global/readwrite    This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param bridgeName The bridgeName of the Bridge.
 * @param bridgeVirtualRouter The bridgeVirtualRouter of the Bridge.
 * @param body The Bridge object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnBridgeResponse
 */
func (a MsgVpnApi) UpdateMsgVpnBridge(msgVpnName string, bridgeName string, bridgeVirtualRouter string, body MsgVpnBridge, select_ []string) (*MsgVpnBridgeResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Patch")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeName"+"}", fmt.Sprintf("%v", bridgeName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeVirtualRouter"+"}", fmt.Sprintf("%v", bridgeVirtualRouter), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnBridgeResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "UpdateMsgVpnBridge", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Updates a Remote Message VPN object.
 * Updates a Remote Message VPN object. Any attribute missing from the request will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: bridgeName|x|x||| bridgeVirtualRouter|x|x||| clientUsername||||x| compressedDataEnabled||||x| egressFlowWindowSize||||x| msgVpnName|x|x||| password|||x|x| remoteMsgVpnInterface|x|x||| remoteMsgVpnLocation|x|x||| remoteMsgVpnName|x|x||| tlsEnabled||||x|    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnBridgeRemoteMsgVpn|clientUsername|password| MsgVpnBridgeRemoteMsgVpn|password|clientUsername|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param bridgeName The bridgeName of the Bridge.
 * @param bridgeVirtualRouter The bridgeVirtualRouter of the Bridge.
 * @param remoteMsgVpnName The remoteMsgVpnName of the Remote Message VPN.
 * @param remoteMsgVpnLocation The remoteMsgVpnLocation of the Remote Message VPN.
 * @param remoteMsgVpnInterface The remoteMsgVpnInterface of the Remote Message VPN.
 * @param body The Remote Message VPN object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnBridgeRemoteMsgVpnResponse
 */
func (a MsgVpnApi) UpdateMsgVpnBridgeRemoteMsgVpn(msgVpnName string, bridgeName string, bridgeVirtualRouter string, remoteMsgVpnName string, remoteMsgVpnLocation string, remoteMsgVpnInterface string, body MsgVpnBridgeRemoteMsgVpn, select_ []string) (*MsgVpnBridgeRemoteMsgVpnResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Patch")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeName"+"}", fmt.Sprintf("%v", bridgeName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeVirtualRouter"+"}", fmt.Sprintf("%v", bridgeVirtualRouter), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"remoteMsgVpnName"+"}", fmt.Sprintf("%v", remoteMsgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"remoteMsgVpnLocation"+"}", fmt.Sprintf("%v", remoteMsgVpnLocation), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"remoteMsgVpnInterface"+"}", fmt.Sprintf("%v", remoteMsgVpnInterface), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnBridgeRemoteMsgVpnResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "UpdateMsgVpnBridgeRemoteMsgVpn", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Updates a Client Profile object.
 * Updates a Client Profile object. Any attribute missing from the request will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: clientProfileName|x|x||| msgVpnName|x|x|||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent EventThresholdByPercent|clearPercent|setPercent| EventThresholdByPercent|setPercent|clearPercent|    A SEMP client authorized with a minimum access scope/level of \&quot;global/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param clientProfileName The clientProfileName of the Client Profile.
 * @param body The Client Profile object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnClientProfileResponse
 */
func (a MsgVpnApi) UpdateMsgVpnClientProfile(msgVpnName string, clientProfileName string, body MsgVpnClientProfile, select_ []string) (*MsgVpnClientProfileResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Patch")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/clientProfiles/{clientProfileName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientProfileName"+"}", fmt.Sprintf("%v", clientProfileName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnClientProfileResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "UpdateMsgVpnClientProfile", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Updates a Client Username object.
 * Updates a Client Username object. Any attribute missing from the request will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: aclProfileName||||x| clientProfileName||||x| clientUsername|x|x||| msgVpnName|x|x||| password|||x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param clientUsername The clientUsername of the Client Username.
 * @param body The Client Username object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnClientUsernameResponse
 */
func (a MsgVpnApi) UpdateMsgVpnClientUsername(msgVpnName string, clientUsername string, body MsgVpnClientUsername, select_ []string) (*MsgVpnClientUsernameResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Patch")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/clientUsernames/{clientUsername}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientUsername"+"}", fmt.Sprintf("%v", clientUsername), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnClientUsernameResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "UpdateMsgVpnClientUsername", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Updates a JNDI Connection Factory object.
 * Updates a JNDI Connection Factory object. Any attribute missing from the request will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: connectionFactoryName|x|x||| msgVpnName|x|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.2.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param connectionFactoryName The connectionFactoryName of the JNDI Connection Factory.
 * @param body The JNDI Connection Factory object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnJndiConnectionFactoryResponse
 */
func (a MsgVpnApi) UpdateMsgVpnJndiConnectionFactory(msgVpnName string, connectionFactoryName string, body MsgVpnJndiConnectionFactory, select_ []string) (*MsgVpnJndiConnectionFactoryResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Patch")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"connectionFactoryName"+"}", fmt.Sprintf("%v", connectionFactoryName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnJndiConnectionFactoryResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "UpdateMsgVpnJndiConnectionFactory", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Updates a JNDI Queue object.
 * Updates a JNDI Queue object. Any attribute missing from the request will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: msgVpnName|x|x||| queueName|x|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.2.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param queueName The queueName of the JNDI Queue.
 * @param body The JNDI Queue object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnJndiQueueResponse
 */
func (a MsgVpnApi) UpdateMsgVpnJndiQueue(msgVpnName string, queueName string, body MsgVpnJndiQueue, select_ []string) (*MsgVpnJndiQueueResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Patch")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/jndiQueues/{queueName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueName"+"}", fmt.Sprintf("%v", queueName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnJndiQueueResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "UpdateMsgVpnJndiQueue", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Updates a JNDI Topic object.
 * Updates a JNDI Topic object. Any attribute missing from the request will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: msgVpnName|x|x||| topicName|x|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.2.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param topicName The topicName of the JNDI Topic.
 * @param body The JNDI Topic object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnJndiTopicResponse
 */
func (a MsgVpnApi) UpdateMsgVpnJndiTopic(msgVpnName string, topicName string, body MsgVpnJndiTopic, select_ []string) (*MsgVpnJndiTopicResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Patch")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/jndiTopics/{topicName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topicName"+"}", fmt.Sprintf("%v", topicName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnJndiTopicResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "UpdateMsgVpnJndiTopic", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Updates an MQTT Session object.
 * Updates an MQTT Session object. Any attribute missing from the request will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: mqttSessionClientId|x|x||| mqttSessionVirtualRouter|x|x||| msgVpnName|x|x||| owner||||x|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.1.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param mqttSessionClientId The mqttSessionClientId of the MQTT Session.
 * @param mqttSessionVirtualRouter The mqttSessionVirtualRouter of the MQTT Session.
 * @param body The MQTT Session object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnMqttSessionResponse
 */
func (a MsgVpnApi) UpdateMsgVpnMqttSession(msgVpnName string, mqttSessionClientId string, mqttSessionVirtualRouter string, body MsgVpnMqttSession, select_ []string) (*MsgVpnMqttSessionResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Patch")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mqttSessionClientId"+"}", fmt.Sprintf("%v", mqttSessionClientId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mqttSessionVirtualRouter"+"}", fmt.Sprintf("%v", mqttSessionVirtualRouter), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnMqttSessionResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "UpdateMsgVpnMqttSession", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Updates an MQTT Session Subscription object.
 * Updates an MQTT Session Subscription object. Any attribute missing from the request will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: mqttSessionClientId|x|x||| mqttSessionVirtualRouter|x|x||| msgVpnName|x|x||| subscriptionTopic|x|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.1.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param mqttSessionClientId The mqttSessionClientId of the MQTT Session.
 * @param mqttSessionVirtualRouter The mqttSessionVirtualRouter of the MQTT Session.
 * @param subscriptionTopic The subscriptionTopic of the MQTT Session Subscription.
 * @param body The MQTT Session Subscription object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnMqttSessionSubscriptionResponse
 */
func (a MsgVpnApi) UpdateMsgVpnMqttSessionSubscription(msgVpnName string, mqttSessionClientId string, mqttSessionVirtualRouter string, subscriptionTopic string, body MsgVpnMqttSessionSubscription, select_ []string) (*MsgVpnMqttSessionSubscriptionResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Patch")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mqttSessionClientId"+"}", fmt.Sprintf("%v", mqttSessionClientId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mqttSessionVirtualRouter"+"}", fmt.Sprintf("%v", mqttSessionVirtualRouter), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionTopic"+"}", fmt.Sprintf("%v", subscriptionTopic), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnMqttSessionSubscriptionResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "UpdateMsgVpnMqttSessionSubscription", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Updates a Queue object.
 * Updates a Queue object. Any attribute missing from the request will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: accessType||||x| msgVpnName|x|x||| owner||||x| permission||||x| queueName|x|x|||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param queueName The queueName of the Queue.
 * @param body The Queue object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnQueueResponse
 */
func (a MsgVpnApi) UpdateMsgVpnQueue(msgVpnName string, queueName string, body MsgVpnQueue, select_ []string) (*MsgVpnQueueResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Patch")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/queues/{queueName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueName"+"}", fmt.Sprintf("%v", queueName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnQueueResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "UpdateMsgVpnQueue", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Updates a Replicated Topic object.
 * Updates a Replicated Topic object. Any attribute missing from the request will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: msgVpnName|x|x||| replicatedTopic|x|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.1.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param replicatedTopic The replicatedTopic of the Replicated Topic.
 * @param body The Replicated Topic object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnReplicatedTopicResponse
 */
func (a MsgVpnApi) UpdateMsgVpnReplicatedTopic(msgVpnName string, replicatedTopic string, body MsgVpnReplicatedTopic, select_ []string) (*MsgVpnReplicatedTopicResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Patch")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"replicatedTopic"+"}", fmt.Sprintf("%v", replicatedTopic), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnReplicatedTopicResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "UpdateMsgVpnReplicatedTopic", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Updates a REST Delivery Point object.
 * Updates a REST Delivery Point object. Any attribute missing from the request will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: clientProfileName||||x| msgVpnName|x|x||| restDeliveryPointName|x|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param restDeliveryPointName The restDeliveryPointName of the REST Delivery Point.
 * @param body The REST Delivery Point object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnRestDeliveryPointResponse
 */
func (a MsgVpnApi) UpdateMsgVpnRestDeliveryPoint(msgVpnName string, restDeliveryPointName string, body MsgVpnRestDeliveryPoint, select_ []string) (*MsgVpnRestDeliveryPointResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Patch")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnRestDeliveryPointResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "UpdateMsgVpnRestDeliveryPoint", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Updates a Queue Binding object.
 * Updates a Queue Binding object. Any attribute missing from the request will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: msgVpnName|x|x||| queueBindingName|x|x||| restDeliveryPointName|x|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param restDeliveryPointName The restDeliveryPointName of the REST Delivery Point.
 * @param queueBindingName The queueBindingName of the Queue Binding.
 * @param body The Queue Binding object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnRestDeliveryPointQueueBindingResponse
 */
func (a MsgVpnApi) UpdateMsgVpnRestDeliveryPointQueueBinding(msgVpnName string, restDeliveryPointName string, queueBindingName string, body MsgVpnRestDeliveryPointQueueBinding, select_ []string) (*MsgVpnRestDeliveryPointQueueBindingResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Patch")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueBindingName"+"}", fmt.Sprintf("%v", queueBindingName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnRestDeliveryPointQueueBindingResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "UpdateMsgVpnRestDeliveryPointQueueBinding", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Updates a REST Consumer object.
 * Updates a REST Consumer object. Any attribute missing from the request will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: authenticationHttpBasicPassword|||x|x| authenticationHttpBasicUsername||||x| authenticationScheme||||x| msgVpnName|x|x||| outgoingConnectionCount||||x| remoteHost||||x| remotePort||||x| restConsumerName|x|x||| restDeliveryPointName|x|x||| tlsCipherSuiteList||||x| tlsEnabled||||x|    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnRestDeliveryPointRestConsumer|authenticationHttpBasicPassword|authenticationHttpBasicUsername| MsgVpnRestDeliveryPointRestConsumer|authenticationHttpBasicUsername|authenticationHttpBasicPassword| MsgVpnRestDeliveryPointRestConsumer|remotePort|tlsEnabled| MsgVpnRestDeliveryPointRestConsumer|tlsEnabled|remotePort|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param restDeliveryPointName The restDeliveryPointName of the REST Delivery Point.
 * @param restConsumerName The restConsumerName of the REST Consumer.
 * @param body The REST Consumer object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnRestDeliveryPointRestConsumerResponse
 */
func (a MsgVpnApi) UpdateMsgVpnRestDeliveryPointRestConsumer(msgVpnName string, restDeliveryPointName string, restConsumerName string, body MsgVpnRestDeliveryPointRestConsumer, select_ []string) (*MsgVpnRestDeliveryPointRestConsumerResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Patch")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restConsumerName"+"}", fmt.Sprintf("%v", restConsumerName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnRestDeliveryPointRestConsumerResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "UpdateMsgVpnRestDeliveryPointRestConsumer", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Updates a Topic Endpoint object.
 * Updates a Topic Endpoint object. Any attribute missing from the request will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: msgVpnName|x|x||| owner||||x| permission||||x| topicEndpointName|x|x|||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.1.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param topicEndpointName The topicEndpointName of the Topic Endpoint.
 * @param body The Topic Endpoint object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnTopicEndpointResponse
 */
func (a MsgVpnApi) UpdateMsgVpnTopicEndpoint(msgVpnName string, topicEndpointName string, body MsgVpnTopicEndpoint, select_ []string) (*MsgVpnTopicEndpointResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Patch")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topicEndpointName"+"}", fmt.Sprintf("%v", topicEndpointName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var select_CollectionFormat = "csv"
	localVarQueryParams.Add("select", a.Configuration.APIClient.ParameterToString(select_, select_CollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(MsgVpnTopicEndpointResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "UpdateMsgVpnTopicEndpoint", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

