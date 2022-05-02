# UrlInfo
## Version

### /urlinfo/{apiversion}/batchupdate

#### POST
##### Summary

Batch add new Url

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| apiversion | path |  | Yes | string |
| body | body |  | Yes | [BatchUpdateRequest](#batchupdaterequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [Response](#response) |

### /urlinfo/{apiversion}/update

#### POST
##### Summary

Add new Url

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| apiversion | path |  | Yes | string |
| body | body |  | Yes | [UpdateRequest](#updaterequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [Response](#response) |

### /urlinfo/{apiversion}/{hostnameport}/{queryparamter}

#### GET
##### Summary

Look up url

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| apiversion | path |  | Yes | string |
| hostnameport | path |  | Yes | string |
| queryparamter | path |  | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [LookupResponse](#lookupresponse) |

### Models

#### BatchUpdateRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| apiversion | string |  | Yes |
| requests | [ [UpdateRequestWithoutVersion](#updaterequestwithoutversion) ] |  | Yes |

#### LookupResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string |  | Yes |
| allow | boolean (boolean) |  | Yes |

#### Request

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| apiversion | string |  | Yes |
| hostnameport | string |  | Yes |
| queryparamter | string |  | Yes |

#### Response

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string |  | Yes |

#### UpdateRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| apiversion | string |  | Yes |
| hostnameport | string |  | Yes |
| queryparamter | string |  | Yes |

#### UpdateRequestWithoutVersion

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| hostnameport | string |  | Yes |
| queryparamter | string |  | Yes |
