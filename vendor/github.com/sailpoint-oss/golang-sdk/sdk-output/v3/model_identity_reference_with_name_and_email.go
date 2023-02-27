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

// IdentityReferenceWithNameAndEmail struct for IdentityReferenceWithNameAndEmail
type IdentityReferenceWithNameAndEmail struct {
	// The type can only be IDENTITY. This is read-only
	Type *string `json:"type,omitempty"`
	// Identity id.
	Id *string `json:"id,omitempty"`
	// Human-readable display name of identity. This is read-only
	Name *string `json:"name,omitempty"`
	// Email address of identity. This is read-only
	Email *string `json:"email,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityReferenceWithNameAndEmail IdentityReferenceWithNameAndEmail

// NewIdentityReferenceWithNameAndEmail instantiates a new IdentityReferenceWithNameAndEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityReferenceWithNameAndEmail() *IdentityReferenceWithNameAndEmail {
	this := IdentityReferenceWithNameAndEmail{}
	return &this
}

// NewIdentityReferenceWithNameAndEmailWithDefaults instantiates a new IdentityReferenceWithNameAndEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityReferenceWithNameAndEmailWithDefaults() *IdentityReferenceWithNameAndEmail {
	this := IdentityReferenceWithNameAndEmail{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IdentityReferenceWithNameAndEmail) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityReferenceWithNameAndEmail) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IdentityReferenceWithNameAndEmail) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IdentityReferenceWithNameAndEmail) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityReferenceWithNameAndEmail) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityReferenceWithNameAndEmail) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityReferenceWithNameAndEmail) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentityReferenceWithNameAndEmail) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IdentityReferenceWithNameAndEmail) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityReferenceWithNameAndEmail) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IdentityReferenceWithNameAndEmail) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IdentityReferenceWithNameAndEmail) SetName(v string) {
	o.Name = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *IdentityReferenceWithNameAndEmail) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityReferenceWithNameAndEmail) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *IdentityReferenceWithNameAndEmail) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *IdentityReferenceWithNameAndEmail) SetEmail(v string) {
	o.Email = &v
}

func (o IdentityReferenceWithNameAndEmail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityReferenceWithNameAndEmail) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityReferenceWithNameAndEmail := _IdentityReferenceWithNameAndEmail{}

	if err = json.Unmarshal(bytes, &varIdentityReferenceWithNameAndEmail); err == nil {
		*o = IdentityReferenceWithNameAndEmail(varIdentityReferenceWithNameAndEmail)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "email")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityReferenceWithNameAndEmail struct {
	value *IdentityReferenceWithNameAndEmail
	isSet bool
}

func (v NullableIdentityReferenceWithNameAndEmail) Get() *IdentityReferenceWithNameAndEmail {
	return v.value
}

func (v *NullableIdentityReferenceWithNameAndEmail) Set(val *IdentityReferenceWithNameAndEmail) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityReferenceWithNameAndEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityReferenceWithNameAndEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityReferenceWithNameAndEmail(val *IdentityReferenceWithNameAndEmail) *NullableIdentityReferenceWithNameAndEmail {
	return &NullableIdentityReferenceWithNameAndEmail{value: val, isSet: true}
}

func (v NullableIdentityReferenceWithNameAndEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityReferenceWithNameAndEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


