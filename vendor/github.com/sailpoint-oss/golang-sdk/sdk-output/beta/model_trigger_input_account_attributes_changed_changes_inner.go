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

// TriggerInputAccountAttributesChangedChangesInner struct for TriggerInputAccountAttributesChangedChangesInner
type TriggerInputAccountAttributesChangedChangesInner struct {
	// The name of the attribute.
	Attribute string `json:"attribute"`
	OldValue NullableTriggerInputAccountAttributesChangedChangesInnerOldValue `json:"oldValue"`
	NewValue NullableTriggerInputAccountAttributesChangedChangesInnerNewValue `json:"newValue"`
	AdditionalProperties map[string]interface{}
}

type _TriggerInputAccountAttributesChangedChangesInner TriggerInputAccountAttributesChangedChangesInner

// NewTriggerInputAccountAttributesChangedChangesInner instantiates a new TriggerInputAccountAttributesChangedChangesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerInputAccountAttributesChangedChangesInner(attribute string, oldValue NullableTriggerInputAccountAttributesChangedChangesInnerOldValue, newValue NullableTriggerInputAccountAttributesChangedChangesInnerNewValue) *TriggerInputAccountAttributesChangedChangesInner {
	this := TriggerInputAccountAttributesChangedChangesInner{}
	this.Attribute = attribute
	this.OldValue = oldValue
	this.NewValue = newValue
	return &this
}

// NewTriggerInputAccountAttributesChangedChangesInnerWithDefaults instantiates a new TriggerInputAccountAttributesChangedChangesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerInputAccountAttributesChangedChangesInnerWithDefaults() *TriggerInputAccountAttributesChangedChangesInner {
	this := TriggerInputAccountAttributesChangedChangesInner{}
	return &this
}

// GetAttribute returns the Attribute field value
func (o *TriggerInputAccountAttributesChangedChangesInner) GetAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccountAttributesChangedChangesInner) GetAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attribute, true
}

// SetAttribute sets field value
func (o *TriggerInputAccountAttributesChangedChangesInner) SetAttribute(v string) {
	o.Attribute = v
}

// GetOldValue returns the OldValue field value
// If the value is explicit nil, the zero value for TriggerInputAccountAttributesChangedChangesInnerOldValue will be returned
func (o *TriggerInputAccountAttributesChangedChangesInner) GetOldValue() TriggerInputAccountAttributesChangedChangesInnerOldValue {
	if o == nil || o.OldValue.Get() == nil {
		var ret TriggerInputAccountAttributesChangedChangesInnerOldValue
		return ret
	}

	return *o.OldValue.Get()
}

// GetOldValueOk returns a tuple with the OldValue field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TriggerInputAccountAttributesChangedChangesInner) GetOldValueOk() (*TriggerInputAccountAttributesChangedChangesInnerOldValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.OldValue.Get(), o.OldValue.IsSet()
}

// SetOldValue sets field value
func (o *TriggerInputAccountAttributesChangedChangesInner) SetOldValue(v TriggerInputAccountAttributesChangedChangesInnerOldValue) {
	o.OldValue.Set(&v)
}

// GetNewValue returns the NewValue field value
// If the value is explicit nil, the zero value for TriggerInputAccountAttributesChangedChangesInnerNewValue will be returned
func (o *TriggerInputAccountAttributesChangedChangesInner) GetNewValue() TriggerInputAccountAttributesChangedChangesInnerNewValue {
	if o == nil || o.NewValue.Get() == nil {
		var ret TriggerInputAccountAttributesChangedChangesInnerNewValue
		return ret
	}

	return *o.NewValue.Get()
}

// GetNewValueOk returns a tuple with the NewValue field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TriggerInputAccountAttributesChangedChangesInner) GetNewValueOk() (*TriggerInputAccountAttributesChangedChangesInnerNewValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewValue.Get(), o.NewValue.IsSet()
}

// SetNewValue sets field value
func (o *TriggerInputAccountAttributesChangedChangesInner) SetNewValue(v TriggerInputAccountAttributesChangedChangesInnerNewValue) {
	o.NewValue.Set(&v)
}

func (o TriggerInputAccountAttributesChangedChangesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["attribute"] = o.Attribute
	}
	if true {
		toSerialize["oldValue"] = o.OldValue.Get()
	}
	if true {
		toSerialize["newValue"] = o.NewValue.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TriggerInputAccountAttributesChangedChangesInner) UnmarshalJSON(bytes []byte) (err error) {
	varTriggerInputAccountAttributesChangedChangesInner := _TriggerInputAccountAttributesChangedChangesInner{}

	if err = json.Unmarshal(bytes, &varTriggerInputAccountAttributesChangedChangesInner); err == nil {
		*o = TriggerInputAccountAttributesChangedChangesInner(varTriggerInputAccountAttributesChangedChangesInner)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "attribute")
		delete(additionalProperties, "oldValue")
		delete(additionalProperties, "newValue")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTriggerInputAccountAttributesChangedChangesInner struct {
	value *TriggerInputAccountAttributesChangedChangesInner
	isSet bool
}

func (v NullableTriggerInputAccountAttributesChangedChangesInner) Get() *TriggerInputAccountAttributesChangedChangesInner {
	return v.value
}

func (v *NullableTriggerInputAccountAttributesChangedChangesInner) Set(val *TriggerInputAccountAttributesChangedChangesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerInputAccountAttributesChangedChangesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerInputAccountAttributesChangedChangesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerInputAccountAttributesChangedChangesInner(val *TriggerInputAccountAttributesChangedChangesInner) *NullableTriggerInputAccountAttributesChangedChangesInner {
	return &NullableTriggerInputAccountAttributesChangedChangesInner{value: val, isSet: true}
}

func (v NullableTriggerInputAccountAttributesChangedChangesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerInputAccountAttributesChangedChangesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


