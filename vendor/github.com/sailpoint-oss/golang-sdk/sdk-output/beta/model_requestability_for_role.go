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

// RequestabilityForRole struct for RequestabilityForRole
type RequestabilityForRole struct {
	// Whether the requester of the containing object must provide comments justifying the request
	CommentsRequired *bool `json:"commentsRequired,omitempty"`
	// Whether an approver must provide comments when denying the request
	DenialCommentsRequired *bool `json:"denialCommentsRequired,omitempty"`
	// List describing the steps in approving the request
	ApprovalSchemes []ApprovalSchemeForRole `json:"approvalSchemes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestabilityForRole RequestabilityForRole

// NewRequestabilityForRole instantiates a new RequestabilityForRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestabilityForRole() *RequestabilityForRole {
	this := RequestabilityForRole{}
	return &this
}

// NewRequestabilityForRoleWithDefaults instantiates a new RequestabilityForRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestabilityForRoleWithDefaults() *RequestabilityForRole {
	this := RequestabilityForRole{}
	return &this
}

// GetCommentsRequired returns the CommentsRequired field value if set, zero value otherwise.
func (o *RequestabilityForRole) GetCommentsRequired() bool {
	if o == nil || isNil(o.CommentsRequired) {
		var ret bool
		return ret
	}
	return *o.CommentsRequired
}

// GetCommentsRequiredOk returns a tuple with the CommentsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestabilityForRole) GetCommentsRequiredOk() (*bool, bool) {
	if o == nil || isNil(o.CommentsRequired) {
		return nil, false
	}
	return o.CommentsRequired, true
}

// HasCommentsRequired returns a boolean if a field has been set.
func (o *RequestabilityForRole) HasCommentsRequired() bool {
	if o != nil && !isNil(o.CommentsRequired) {
		return true
	}

	return false
}

// SetCommentsRequired gets a reference to the given bool and assigns it to the CommentsRequired field.
func (o *RequestabilityForRole) SetCommentsRequired(v bool) {
	o.CommentsRequired = &v
}

// GetDenialCommentsRequired returns the DenialCommentsRequired field value if set, zero value otherwise.
func (o *RequestabilityForRole) GetDenialCommentsRequired() bool {
	if o == nil || isNil(o.DenialCommentsRequired) {
		var ret bool
		return ret
	}
	return *o.DenialCommentsRequired
}

// GetDenialCommentsRequiredOk returns a tuple with the DenialCommentsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestabilityForRole) GetDenialCommentsRequiredOk() (*bool, bool) {
	if o == nil || isNil(o.DenialCommentsRequired) {
		return nil, false
	}
	return o.DenialCommentsRequired, true
}

// HasDenialCommentsRequired returns a boolean if a field has been set.
func (o *RequestabilityForRole) HasDenialCommentsRequired() bool {
	if o != nil && !isNil(o.DenialCommentsRequired) {
		return true
	}

	return false
}

// SetDenialCommentsRequired gets a reference to the given bool and assigns it to the DenialCommentsRequired field.
func (o *RequestabilityForRole) SetDenialCommentsRequired(v bool) {
	o.DenialCommentsRequired = &v
}

// GetApprovalSchemes returns the ApprovalSchemes field value if set, zero value otherwise.
func (o *RequestabilityForRole) GetApprovalSchemes() []ApprovalSchemeForRole {
	if o == nil || isNil(o.ApprovalSchemes) {
		var ret []ApprovalSchemeForRole
		return ret
	}
	return o.ApprovalSchemes
}

// GetApprovalSchemesOk returns a tuple with the ApprovalSchemes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestabilityForRole) GetApprovalSchemesOk() ([]ApprovalSchemeForRole, bool) {
	if o == nil || isNil(o.ApprovalSchemes) {
		return nil, false
	}
	return o.ApprovalSchemes, true
}

// HasApprovalSchemes returns a boolean if a field has been set.
func (o *RequestabilityForRole) HasApprovalSchemes() bool {
	if o != nil && !isNil(o.ApprovalSchemes) {
		return true
	}

	return false
}

// SetApprovalSchemes gets a reference to the given []ApprovalSchemeForRole and assigns it to the ApprovalSchemes field.
func (o *RequestabilityForRole) SetApprovalSchemes(v []ApprovalSchemeForRole) {
	o.ApprovalSchemes = v
}

func (o RequestabilityForRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CommentsRequired) {
		toSerialize["commentsRequired"] = o.CommentsRequired
	}
	if !isNil(o.DenialCommentsRequired) {
		toSerialize["denialCommentsRequired"] = o.DenialCommentsRequired
	}
	if !isNil(o.ApprovalSchemes) {
		toSerialize["approvalSchemes"] = o.ApprovalSchemes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestabilityForRole) UnmarshalJSON(bytes []byte) (err error) {
	varRequestabilityForRole := _RequestabilityForRole{}

	if err = json.Unmarshal(bytes, &varRequestabilityForRole); err == nil {
		*o = RequestabilityForRole(varRequestabilityForRole)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "commentsRequired")
		delete(additionalProperties, "denialCommentsRequired")
		delete(additionalProperties, "approvalSchemes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestabilityForRole struct {
	value *RequestabilityForRole
	isSet bool
}

func (v NullableRequestabilityForRole) Get() *RequestabilityForRole {
	return v.value
}

func (v *NullableRequestabilityForRole) Set(val *RequestabilityForRole) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestabilityForRole) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestabilityForRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestabilityForRole(val *RequestabilityForRole) *NullableRequestabilityForRole {
	return &NullableRequestabilityForRole{value: val, isSet: true}
}

func (v NullableRequestabilityForRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestabilityForRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


