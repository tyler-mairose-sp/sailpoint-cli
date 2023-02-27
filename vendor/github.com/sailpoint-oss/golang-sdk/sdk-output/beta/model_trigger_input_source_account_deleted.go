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

// TriggerInputSourceAccountDeleted struct for TriggerInputSourceAccountDeleted
type TriggerInputSourceAccountDeleted struct {
	// Source unique identifier for the identity. UUID is generated by the source system.
	Uuid *string `json:"uuid,omitempty"`
	// SailPoint generated unique identifier.
	Id string `json:"id"`
	// Unique ID of the account on the source.
	NativeIdentifier string `json:"nativeIdentifier"`
	// The ID of the source.
	SourceId string `json:"sourceId"`
	// The name of the source.
	SourceName string `json:"sourceName"`
	// The ID of the identity that is corellated with this account.
	IdentityId string `json:"identityId"`
	// The name of the identity that is corellated with this account.
	IdentityName string `json:"identityName"`
	// The attributes of the account. The contents of attributes depends on the account schema for the source.
	Attributes map[string]interface{} `json:"attributes"`
	AdditionalProperties map[string]interface{}
}

type _TriggerInputSourceAccountDeleted TriggerInputSourceAccountDeleted

// NewTriggerInputSourceAccountDeleted instantiates a new TriggerInputSourceAccountDeleted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerInputSourceAccountDeleted(id string, nativeIdentifier string, sourceId string, sourceName string, identityId string, identityName string, attributes map[string]interface{}) *TriggerInputSourceAccountDeleted {
	this := TriggerInputSourceAccountDeleted{}
	this.Id = id
	this.NativeIdentifier = nativeIdentifier
	this.SourceId = sourceId
	this.SourceName = sourceName
	this.IdentityId = identityId
	this.IdentityName = identityName
	this.Attributes = attributes
	return &this
}

// NewTriggerInputSourceAccountDeletedWithDefaults instantiates a new TriggerInputSourceAccountDeleted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerInputSourceAccountDeletedWithDefaults() *TriggerInputSourceAccountDeleted {
	this := TriggerInputSourceAccountDeleted{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *TriggerInputSourceAccountDeleted) GetUuid() string {
	if o == nil || isNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerInputSourceAccountDeleted) GetUuidOk() (*string, bool) {
	if o == nil || isNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *TriggerInputSourceAccountDeleted) HasUuid() bool {
	if o != nil && !isNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *TriggerInputSourceAccountDeleted) SetUuid(v string) {
	o.Uuid = &v
}

// GetId returns the Id field value
func (o *TriggerInputSourceAccountDeleted) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TriggerInputSourceAccountDeleted) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TriggerInputSourceAccountDeleted) SetId(v string) {
	o.Id = v
}

// GetNativeIdentifier returns the NativeIdentifier field value
func (o *TriggerInputSourceAccountDeleted) GetNativeIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NativeIdentifier
}

// GetNativeIdentifierOk returns a tuple with the NativeIdentifier field value
// and a boolean to check if the value has been set.
func (o *TriggerInputSourceAccountDeleted) GetNativeIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NativeIdentifier, true
}

// SetNativeIdentifier sets field value
func (o *TriggerInputSourceAccountDeleted) SetNativeIdentifier(v string) {
	o.NativeIdentifier = v
}

// GetSourceId returns the SourceId field value
func (o *TriggerInputSourceAccountDeleted) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *TriggerInputSourceAccountDeleted) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *TriggerInputSourceAccountDeleted) SetSourceId(v string) {
	o.SourceId = v
}

// GetSourceName returns the SourceName field value
func (o *TriggerInputSourceAccountDeleted) GetSourceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value
// and a boolean to check if the value has been set.
func (o *TriggerInputSourceAccountDeleted) GetSourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceName, true
}

// SetSourceName sets field value
func (o *TriggerInputSourceAccountDeleted) SetSourceName(v string) {
	o.SourceName = v
}

// GetIdentityId returns the IdentityId field value
func (o *TriggerInputSourceAccountDeleted) GetIdentityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value
// and a boolean to check if the value has been set.
func (o *TriggerInputSourceAccountDeleted) GetIdentityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityId, true
}

// SetIdentityId sets field value
func (o *TriggerInputSourceAccountDeleted) SetIdentityId(v string) {
	o.IdentityId = v
}

// GetIdentityName returns the IdentityName field value
func (o *TriggerInputSourceAccountDeleted) GetIdentityName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityName
}

// GetIdentityNameOk returns a tuple with the IdentityName field value
// and a boolean to check if the value has been set.
func (o *TriggerInputSourceAccountDeleted) GetIdentityNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityName, true
}

// SetIdentityName sets field value
func (o *TriggerInputSourceAccountDeleted) SetIdentityName(v string) {
	o.IdentityName = v
}

// GetAttributes returns the Attributes field value
func (o *TriggerInputSourceAccountDeleted) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TriggerInputSourceAccountDeleted) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *TriggerInputSourceAccountDeleted) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

func (o TriggerInputSourceAccountDeleted) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["nativeIdentifier"] = o.NativeIdentifier
	}
	if true {
		toSerialize["sourceId"] = o.SourceId
	}
	if true {
		toSerialize["sourceName"] = o.SourceName
	}
	if true {
		toSerialize["identityId"] = o.IdentityId
	}
	if true {
		toSerialize["identityName"] = o.IdentityName
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TriggerInputSourceAccountDeleted) UnmarshalJSON(bytes []byte) (err error) {
	varTriggerInputSourceAccountDeleted := _TriggerInputSourceAccountDeleted{}

	if err = json.Unmarshal(bytes, &varTriggerInputSourceAccountDeleted); err == nil {
		*o = TriggerInputSourceAccountDeleted(varTriggerInputSourceAccountDeleted)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "id")
		delete(additionalProperties, "nativeIdentifier")
		delete(additionalProperties, "sourceId")
		delete(additionalProperties, "sourceName")
		delete(additionalProperties, "identityId")
		delete(additionalProperties, "identityName")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTriggerInputSourceAccountDeleted struct {
	value *TriggerInputSourceAccountDeleted
	isSet bool
}

func (v NullableTriggerInputSourceAccountDeleted) Get() *TriggerInputSourceAccountDeleted {
	return v.value
}

func (v *NullableTriggerInputSourceAccountDeleted) Set(val *TriggerInputSourceAccountDeleted) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerInputSourceAccountDeleted) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerInputSourceAccountDeleted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerInputSourceAccountDeleted(val *TriggerInputSourceAccountDeleted) *NullableTriggerInputSourceAccountDeleted {
	return &NullableTriggerInputSourceAccountDeleted{value: val, isSet: true}
}

func (v NullableTriggerInputSourceAccountDeleted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerInputSourceAccountDeleted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


