/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
)

// RoleSummaryAllOf struct for RoleSummaryAllOf
type RoleSummaryAllOf struct {
	Owner *DisplayReference `json:"owner,omitempty"`
	Disabled *bool `json:"disabled,omitempty"`
	Revocable *bool `json:"revocable,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleSummaryAllOf RoleSummaryAllOf

// NewRoleSummaryAllOf instantiates a new RoleSummaryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleSummaryAllOf() *RoleSummaryAllOf {
	this := RoleSummaryAllOf{}
	return &this
}

// NewRoleSummaryAllOfWithDefaults instantiates a new RoleSummaryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleSummaryAllOfWithDefaults() *RoleSummaryAllOf {
	this := RoleSummaryAllOf{}
	return &this
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *RoleSummaryAllOf) GetOwner() DisplayReference {
	if o == nil || isNil(o.Owner) {
		var ret DisplayReference
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleSummaryAllOf) GetOwnerOk() (*DisplayReference, bool) {
	if o == nil || isNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *RoleSummaryAllOf) HasOwner() bool {
	if o != nil && !isNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given DisplayReference and assigns it to the Owner field.
func (o *RoleSummaryAllOf) SetOwner(v DisplayReference) {
	o.Owner = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *RoleSummaryAllOf) GetDisabled() bool {
	if o == nil || isNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleSummaryAllOf) GetDisabledOk() (*bool, bool) {
	if o == nil || isNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *RoleSummaryAllOf) HasDisabled() bool {
	if o != nil && !isNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *RoleSummaryAllOf) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetRevocable returns the Revocable field value if set, zero value otherwise.
func (o *RoleSummaryAllOf) GetRevocable() bool {
	if o == nil || isNil(o.Revocable) {
		var ret bool
		return ret
	}
	return *o.Revocable
}

// GetRevocableOk returns a tuple with the Revocable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleSummaryAllOf) GetRevocableOk() (*bool, bool) {
	if o == nil || isNil(o.Revocable) {
		return nil, false
	}
	return o.Revocable, true
}

// HasRevocable returns a boolean if a field has been set.
func (o *RoleSummaryAllOf) HasRevocable() bool {
	if o != nil && !isNil(o.Revocable) {
		return true
	}

	return false
}

// SetRevocable gets a reference to the given bool and assigns it to the Revocable field.
func (o *RoleSummaryAllOf) SetRevocable(v bool) {
	o.Revocable = &v
}

func (o RoleSummaryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !isNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !isNil(o.Revocable) {
		toSerialize["revocable"] = o.Revocable
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RoleSummaryAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varRoleSummaryAllOf := _RoleSummaryAllOf{}

	if err = json.Unmarshal(bytes, &varRoleSummaryAllOf); err == nil {
		*o = RoleSummaryAllOf(varRoleSummaryAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "owner")
		delete(additionalProperties, "disabled")
		delete(additionalProperties, "revocable")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleSummaryAllOf struct {
	value *RoleSummaryAllOf
	isSet bool
}

func (v NullableRoleSummaryAllOf) Get() *RoleSummaryAllOf {
	return v.value
}

func (v *NullableRoleSummaryAllOf) Set(val *RoleSummaryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleSummaryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleSummaryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleSummaryAllOf(val *RoleSummaryAllOf) *NullableRoleSummaryAllOf {
	return &NullableRoleSummaryAllOf{value: val, isSet: true}
}

func (v NullableRoleSummaryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleSummaryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


