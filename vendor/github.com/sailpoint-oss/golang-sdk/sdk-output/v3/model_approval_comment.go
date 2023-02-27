/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
	"time"
)

// ApprovalComment struct for ApprovalComment
type ApprovalComment struct {
	// The comment text
	Comment *string `json:"comment,omitempty"`
	// The name of the commenter
	Commenter *string `json:"commenter,omitempty"`
	// A date-time in ISO-8601 format
	Date NullableTime `json:"date,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApprovalComment ApprovalComment

// NewApprovalComment instantiates a new ApprovalComment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApprovalComment() *ApprovalComment {
	this := ApprovalComment{}
	return &this
}

// NewApprovalCommentWithDefaults instantiates a new ApprovalComment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApprovalCommentWithDefaults() *ApprovalComment {
	this := ApprovalComment{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ApprovalComment) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalComment) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ApprovalComment) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ApprovalComment) SetComment(v string) {
	o.Comment = &v
}

// GetCommenter returns the Commenter field value if set, zero value otherwise.
func (o *ApprovalComment) GetCommenter() string {
	if o == nil || isNil(o.Commenter) {
		var ret string
		return ret
	}
	return *o.Commenter
}

// GetCommenterOk returns a tuple with the Commenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalComment) GetCommenterOk() (*string, bool) {
	if o == nil || isNil(o.Commenter) {
		return nil, false
	}
	return o.Commenter, true
}

// HasCommenter returns a boolean if a field has been set.
func (o *ApprovalComment) HasCommenter() bool {
	if o != nil && !isNil(o.Commenter) {
		return true
	}

	return false
}

// SetCommenter gets a reference to the given string and assigns it to the Commenter field.
func (o *ApprovalComment) SetCommenter(v string) {
	o.Commenter = &v
}

// GetDate returns the Date field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApprovalComment) GetDate() time.Time {
	if o == nil || isNil(o.Date.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Date.Get()
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApprovalComment) GetDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Date.Get(), o.Date.IsSet()
}

// HasDate returns a boolean if a field has been set.
func (o *ApprovalComment) HasDate() bool {
	if o != nil && o.Date.IsSet() {
		return true
	}

	return false
}

// SetDate gets a reference to the given NullableTime and assigns it to the Date field.
func (o *ApprovalComment) SetDate(v time.Time) {
	o.Date.Set(&v)
}
// SetDateNil sets the value for Date to be an explicit nil
func (o *ApprovalComment) SetDateNil() {
	o.Date.Set(nil)
}

// UnsetDate ensures that no value is present for Date, not even an explicit nil
func (o *ApprovalComment) UnsetDate() {
	o.Date.Unset()
}

func (o ApprovalComment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !isNil(o.Commenter) {
		toSerialize["commenter"] = o.Commenter
	}
	if o.Date.IsSet() {
		toSerialize["date"] = o.Date.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApprovalComment) UnmarshalJSON(bytes []byte) (err error) {
	varApprovalComment := _ApprovalComment{}

	if err = json.Unmarshal(bytes, &varApprovalComment); err == nil {
		*o = ApprovalComment(varApprovalComment)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "comment")
		delete(additionalProperties, "commenter")
		delete(additionalProperties, "date")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApprovalComment struct {
	value *ApprovalComment
	isSet bool
}

func (v NullableApprovalComment) Get() *ApprovalComment {
	return v.value
}

func (v *NullableApprovalComment) Set(val *ApprovalComment) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalComment) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalComment(val *ApprovalComment) *NullableApprovalComment {
	return &NullableApprovalComment{value: val, isSet: true}
}

func (v NullableApprovalComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


