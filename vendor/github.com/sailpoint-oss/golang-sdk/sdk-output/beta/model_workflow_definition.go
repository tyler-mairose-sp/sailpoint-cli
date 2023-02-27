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

// WorkflowDefinition The map of steps that the workflow will execute.
type WorkflowDefinition struct {
	// The name of the starting step.
	Start *string `json:"start,omitempty"`
	// One or more step objects that comprise this workflow.  Please see the Workflow documentation to see the JSON schema for each step type.
	Steps map[string]interface{} `json:"steps,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowDefinition WorkflowDefinition

// NewWorkflowDefinition instantiates a new WorkflowDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowDefinition() *WorkflowDefinition {
	this := WorkflowDefinition{}
	return &this
}

// NewWorkflowDefinitionWithDefaults instantiates a new WorkflowDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowDefinitionWithDefaults() *WorkflowDefinition {
	this := WorkflowDefinition{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *WorkflowDefinition) GetStart() string {
	if o == nil || isNil(o.Start) {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDefinition) GetStartOk() (*string, bool) {
	if o == nil || isNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *WorkflowDefinition) HasStart() bool {
	if o != nil && !isNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *WorkflowDefinition) SetStart(v string) {
	o.Start = &v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *WorkflowDefinition) GetSteps() map[string]interface{} {
	if o == nil || isNil(o.Steps) {
		var ret map[string]interface{}
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDefinition) GetStepsOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Steps) {
		return map[string]interface{}{}, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *WorkflowDefinition) HasSteps() bool {
	if o != nil && !isNil(o.Steps) {
		return true
	}

	return false
}

// SetSteps gets a reference to the given map[string]interface{} and assigns it to the Steps field.
func (o *WorkflowDefinition) SetSteps(v map[string]interface{}) {
	o.Steps = v
}

func (o WorkflowDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !isNil(o.Steps) {
		toSerialize["steps"] = o.Steps
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowDefinition) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowDefinition := _WorkflowDefinition{}

	if err = json.Unmarshal(bytes, &varWorkflowDefinition); err == nil {
		*o = WorkflowDefinition(varWorkflowDefinition)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "start")
		delete(additionalProperties, "steps")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowDefinition struct {
	value *WorkflowDefinition
	isSet bool
}

func (v NullableWorkflowDefinition) Get() *WorkflowDefinition {
	return v.value
}

func (v *NullableWorkflowDefinition) Set(val *WorkflowDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowDefinition(val *WorkflowDefinition) *NullableWorkflowDefinition {
	return &NullableWorkflowDefinition{value: val, isSet: true}
}

func (v NullableWorkflowDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


