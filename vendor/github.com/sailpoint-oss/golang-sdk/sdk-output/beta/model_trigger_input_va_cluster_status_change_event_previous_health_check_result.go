/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
)

// TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult The results of the last health check.
type TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult struct {
	// Detailed message of the result of the health check.
	Message string `json:"message"`
	// The type of the health check result.
	ResultType string `json:"resultType"`
	// The status of the health check.
	Status map[string]interface{} `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult

// NewTriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult instantiates a new TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult(message string, resultType string, status map[string]interface{}) *TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult {
	this := TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult{}
	this.Message = message
	this.ResultType = resultType
	this.Status = status
	return &this
}

// NewTriggerInputVAClusterStatusChangeEventPreviousHealthCheckResultWithDefaults instantiates a new TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerInputVAClusterStatusChangeEventPreviousHealthCheckResultWithDefaults() *TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult {
	this := TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult{}
	return &this
}

// GetMessage returns the Message field value
func (o *TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult) SetMessage(v string) {
	o.Message = v
}

// GetResultType returns the ResultType field value
func (o *TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult) GetResultType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResultType
}

// GetResultTypeOk returns a tuple with the ResultType field value
// and a boolean to check if the value has been set.
func (o *TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult) GetResultTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultType, true
}

// SetResultType sets field value
func (o *TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult) SetResultType(v string) {
	o.ResultType = v
}

// GetStatus returns the Status field value
func (o *TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult) GetStatus() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult) GetStatusOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Status, true
}

// SetStatus sets field value
func (o *TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult) SetStatus(v map[string]interface{}) {
	o.Status = v
}

func (o TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["resultType"] = o.ResultType
	}
	if true {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult) UnmarshalJSON(bytes []byte) (err error) {
	varTriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult := _TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult{}

	if err = json.Unmarshal(bytes, &varTriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult); err == nil {
		*o = TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult(varTriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "message")
		delete(additionalProperties, "resultType")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult struct {
	value *TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult
	isSet bool
}

func (v NullableTriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult) Get() *TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult {
	return v.value
}

func (v *NullableTriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult) Set(val *TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult(val *TriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult) *NullableTriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult {
	return &NullableTriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult{value: val, isSet: true}
}

func (v NullableTriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerInputVAClusterStatusChangeEventPreviousHealthCheckResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


