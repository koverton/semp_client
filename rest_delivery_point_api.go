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

type RestDeliveryPointApi struct {
	Configuration *Configuration
}

func NewRestDeliveryPointApi() *RestDeliveryPointApi {
	configuration := NewConfiguration()
	return &RestDeliveryPointApi{
		Configuration: configuration,
	}
}

func NewRestDeliveryPointApiWithBasePath(basePath string) *RestDeliveryPointApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &RestDeliveryPointApi{
		Configuration: configuration,
	}
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
func (a RestDeliveryPointApi) CreateMsgVpnRestDeliveryPoint(msgVpnName string, body MsgVpnRestDeliveryPoint, select_ []string) (*MsgVpnRestDeliveryPointResponse, *APIResponse, error) {

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
func (a RestDeliveryPointApi) CreateMsgVpnRestDeliveryPointQueueBinding(msgVpnName string, restDeliveryPointName string, body MsgVpnRestDeliveryPointQueueBinding, select_ []string) (*MsgVpnRestDeliveryPointQueueBindingResponse, *APIResponse, error) {

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
func (a RestDeliveryPointApi) CreateMsgVpnRestDeliveryPointRestConsumer(msgVpnName string, restDeliveryPointName string, body MsgVpnRestDeliveryPointRestConsumer, select_ []string) (*MsgVpnRestDeliveryPointRestConsumerResponse, *APIResponse, error) {

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
 * Deletes a REST Delivery Point object.
 * Deletes a REST Delivery Point object.  A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param restDeliveryPointName The restDeliveryPointName of the REST Delivery Point.
 * @return *SempMetaOnlyResponse
 */
func (a RestDeliveryPointApi) DeleteMsgVpnRestDeliveryPoint(msgVpnName string, restDeliveryPointName string) (*SempMetaOnlyResponse, *APIResponse, error) {

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
func (a RestDeliveryPointApi) DeleteMsgVpnRestDeliveryPointQueueBinding(msgVpnName string, restDeliveryPointName string, queueBindingName string) (*SempMetaOnlyResponse, *APIResponse, error) {

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
func (a RestDeliveryPointApi) DeleteMsgVpnRestDeliveryPointRestConsumer(msgVpnName string, restDeliveryPointName string, restConsumerName string) (*SempMetaOnlyResponse, *APIResponse, error) {

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
 * Gets a REST Delivery Point object.
 * Gets a REST Delivery Point object.   Attribute|Identifying|Write-Only|Deprecated :---|:---:|:---:|:---: msgVpnName|x|| restDeliveryPointName|x||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readonly\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param restDeliveryPointName The restDeliveryPointName of the REST Delivery Point.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnRestDeliveryPointResponse
 */
func (a RestDeliveryPointApi) GetMsgVpnRestDeliveryPoint(msgVpnName string, restDeliveryPointName string, select_ []string) (*MsgVpnRestDeliveryPointResponse, *APIResponse, error) {

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
func (a RestDeliveryPointApi) GetMsgVpnRestDeliveryPointQueueBinding(msgVpnName string, restDeliveryPointName string, queueBindingName string, select_ []string) (*MsgVpnRestDeliveryPointQueueBindingResponse, *APIResponse, error) {

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
func (a RestDeliveryPointApi) GetMsgVpnRestDeliveryPointQueueBindings(msgVpnName string, restDeliveryPointName string, count int32, cursor string, where []string, select_ []string) (*MsgVpnRestDeliveryPointQueueBindingsResponse, *APIResponse, error) {

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
func (a RestDeliveryPointApi) GetMsgVpnRestDeliveryPointRestConsumer(msgVpnName string, restDeliveryPointName string, restConsumerName string, select_ []string) (*MsgVpnRestDeliveryPointRestConsumerResponse, *APIResponse, error) {

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
func (a RestDeliveryPointApi) GetMsgVpnRestDeliveryPointRestConsumers(msgVpnName string, restDeliveryPointName string, count int32, cursor string, where []string, select_ []string) (*MsgVpnRestDeliveryPointRestConsumersResponse, *APIResponse, error) {

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
func (a RestDeliveryPointApi) GetMsgVpnRestDeliveryPoints(msgVpnName string, count int32, cursor string, where []string, select_ []string) (*MsgVpnRestDeliveryPointsResponse, *APIResponse, error) {

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
 * Replaces a REST Delivery Point object.
 * Replaces a REST Delivery Point object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: clientProfileName||||x| msgVpnName|x|x||| restDeliveryPointName|x|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param restDeliveryPointName The restDeliveryPointName of the REST Delivery Point.
 * @param body The REST Delivery Point object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnRestDeliveryPointResponse
 */
func (a RestDeliveryPointApi) ReplaceMsgVpnRestDeliveryPoint(msgVpnName string, restDeliveryPointName string, body MsgVpnRestDeliveryPoint, select_ []string) (*MsgVpnRestDeliveryPointResponse, *APIResponse, error) {

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
func (a RestDeliveryPointApi) ReplaceMsgVpnRestDeliveryPointQueueBinding(msgVpnName string, restDeliveryPointName string, queueBindingName string, body MsgVpnRestDeliveryPointQueueBinding, select_ []string) (*MsgVpnRestDeliveryPointQueueBindingResponse, *APIResponse, error) {

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
func (a RestDeliveryPointApi) ReplaceMsgVpnRestDeliveryPointRestConsumer(msgVpnName string, restDeliveryPointName string, restConsumerName string, body MsgVpnRestDeliveryPointRestConsumer, select_ []string) (*MsgVpnRestDeliveryPointRestConsumerResponse, *APIResponse, error) {

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
 * Updates a REST Delivery Point object.
 * Updates a REST Delivery Point object. Any attribute missing from the request will be left unchanged.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated :---|:---:|:---:|:---:|:---:|:---: clientProfileName||||x| msgVpnName|x|x||| restDeliveryPointName|x|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/readwrite\&quot; is required to perform this operation.  This has been available since 2.0.0.
 *
 * @param msgVpnName The msgVpnName of the Message VPN.
 * @param restDeliveryPointName The restDeliveryPointName of the REST Delivery Point.
 * @param body The REST Delivery Point object&#39;s attributes.
 * @param select_ Include in the response only selected attributes of the object. See [Select](#select \&quot;Description of the syntax of the &#x60;select&#x60; parameter\&quot;).
 * @return *MsgVpnRestDeliveryPointResponse
 */
func (a RestDeliveryPointApi) UpdateMsgVpnRestDeliveryPoint(msgVpnName string, restDeliveryPointName string, body MsgVpnRestDeliveryPoint, select_ []string) (*MsgVpnRestDeliveryPointResponse, *APIResponse, error) {

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
func (a RestDeliveryPointApi) UpdateMsgVpnRestDeliveryPointQueueBinding(msgVpnName string, restDeliveryPointName string, queueBindingName string, body MsgVpnRestDeliveryPointQueueBinding, select_ []string) (*MsgVpnRestDeliveryPointQueueBindingResponse, *APIResponse, error) {

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
func (a RestDeliveryPointApi) UpdateMsgVpnRestDeliveryPointRestConsumer(msgVpnName string, restDeliveryPointName string, restConsumerName string, body MsgVpnRestDeliveryPointRestConsumer, select_ []string) (*MsgVpnRestDeliveryPointRestConsumerResponse, *APIResponse, error) {

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

