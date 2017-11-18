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

type AclProfileApi struct {
	Configuration *Configuration
}

func NewAclProfileApi() *AclProfileApi {
	configuration := NewConfiguration()
	return &AclProfileApi{
		Configuration: configuration,
	}
}

func NewAclProfileApiWithBasePath(basePath string) *AclProfileApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &AclProfileApi{
		Configuration: configuration,
	}
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
func (a AclProfileApi) CreateMsgVpnAclProfile(msgVpnName string, body MsgVpnAclProfile, select_ []string) (*MsgVpnAclProfileResponse, *APIResponse, error) {

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
func (a AclProfileApi) CreateMsgVpnAclProfileClientConnectException(msgVpnName string, aclProfileName string, body MsgVpnAclProfileClientConnectException, select_ []string) (*MsgVpnAclProfileClientConnectExceptionResponse, *APIResponse, error) {

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
func (a AclProfileApi) CreateMsgVpnAclProfilePublishException(msgVpnName string, aclProfileName string, body MsgVpnAclProfilePublishException, select_ []string) (*MsgVpnAclProfilePublishExceptionResponse, *APIResponse, error) {

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
func (a AclProfileApi) CreateMsgVpnAclProfileSubscribeException(msgVpnName string, aclProfileName string, body MsgVpnAclProfileSubscribeException, select_ []string) (*MsgVpnAclProfileSubscribeExceptionResponse, *APIResponse, error) {

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
 * Deletes an ACL Profile object.
 * Deletes an ACL Profile object.  ACL profiles are a collection of user-defined profiles for specifying Access Control List properties.  A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param aclProfileName The aclProfileName of the ACL Profile.
 * @return *SempMetaOnlyResponse
 */
func (a AclProfileApi) DeleteMsgVpnAclProfile(msgVpnName string, aclProfileName string) (*SempMetaOnlyResponse, *APIResponse, error) {

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
func (a AclProfileApi) DeleteMsgVpnAclProfileClientConnectException(msgVpnName string, aclProfileName string, clientConnectExceptionAddress string) (*SempMetaOnlyResponse, *APIResponse, error) {

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
func (a AclProfileApi) DeleteMsgVpnAclProfilePublishException(msgVpnName string, aclProfileName string, topicSyntax string, publishExceptionTopic string) (*SempMetaOnlyResponse, *APIResponse, error) {

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
func (a AclProfileApi) DeleteMsgVpnAclProfileSubscribeException(msgVpnName string, aclProfileName string, topicSyntax string, subscribeExceptionTopic string) (*SempMetaOnlyResponse, *APIResponse, error) {

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
 * Gets an ACL Profile object.
 * Gets an ACL Profile object.  ACL profiles are a collection of user-defined profiles for specifying Access Control List properties.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: aclProfileName|x|| msgVpnName|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param aclProfileName The aclProfileName of the ACL Profile.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnAclProfileResponse
 */
func (a AclProfileApi) GetMsgVpnAclProfile(msgVpnName string, aclProfileName string, select_ []string) (*MsgVpnAclProfileResponse, *APIResponse, error) {

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
func (a AclProfileApi) GetMsgVpnAclProfileClientConnectException(msgVpnName string, aclProfileName string, clientConnectExceptionAddress string, select_ []string) (*MsgVpnAclProfileClientConnectExceptionResponse, *APIResponse, error) {

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
func (a AclProfileApi) GetMsgVpnAclProfileClientConnectExceptions(msgVpnName string, aclProfileName string, count int32, cursor string, where []string, select_ []string) (*MsgVpnAclProfileClientConnectExceptionsResponse, *APIResponse, error) {

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
func (a AclProfileApi) GetMsgVpnAclProfilePublishException(msgVpnName string, aclProfileName string, topicSyntax string, publishExceptionTopic string, select_ []string) (*MsgVpnAclProfilePublishExceptionResponse, *APIResponse, error) {

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
func (a AclProfileApi) GetMsgVpnAclProfilePublishExceptions(msgVpnName string, aclProfileName string, count int32, cursor string, where []string, select_ []string) (*MsgVpnAclProfilePublishExceptionsResponse, *APIResponse, error) {

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
func (a AclProfileApi) GetMsgVpnAclProfileSubscribeException(msgVpnName string, aclProfileName string, topicSyntax string, subscribeExceptionTopic string, select_ []string) (*MsgVpnAclProfileSubscribeExceptionResponse, *APIResponse, error) {

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
func (a AclProfileApi) GetMsgVpnAclProfileSubscribeExceptions(msgVpnName string, aclProfileName string, count int32, cursor string, where []string, select_ []string) (*MsgVpnAclProfileSubscribeExceptionsResponse, *APIResponse, error) {

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
func (a AclProfileApi) GetMsgVpnAclProfiles(msgVpnName string, count int32, cursor string, where []string, select_ []string) (*MsgVpnAclProfilesResponse, *APIResponse, error) {

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
 * Replaces an ACL Profile object.
 * Replaces an ACL Profile object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.  ACL profiles are a collection of user-defined profiles for specifying Access Control List properties.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: aclProfileName|x|x||| msgVpnName|x|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param aclProfileName The aclProfileName of the ACL Profile.
 * @param body The ACL Profile object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnAclProfileResponse
 */
func (a AclProfileApi) ReplaceMsgVpnAclProfile(msgVpnName string, aclProfileName string, body MsgVpnAclProfile, select_ []string) (*MsgVpnAclProfileResponse, *APIResponse, error) {

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
 * Updates an ACL Profile object.
 * Updates an ACL Profile object. Any attribute missing from the request will be left unchanged.  ACL profiles are a collection of user-defined profiles for specifying Access Control List properties.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: aclProfileName|x|x||| msgVpnName|x|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param aclProfileName The aclProfileName of the ACL Profile.
 * @param body The ACL Profile object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnAclProfileResponse
 */
func (a AclProfileApi) UpdateMsgVpnAclProfile(msgVpnName string, aclProfileName string, body MsgVpnAclProfile, select_ []string) (*MsgVpnAclProfileResponse, *APIResponse, error) {

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

