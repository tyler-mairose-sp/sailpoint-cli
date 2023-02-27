/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"time"
	"encoding/json"
)

// EntitlementDocument Entitlement
type EntitlementDocument struct {
	// The unique ID of the referenced object.
	Id string `json:"id"`
	// The human readable name of the referenced object.
	Name string `json:"name"`
	Type DocumentType `json:"_type"`
	// A description of the entitlement
	Description *string `json:"description,omitempty"`
	// The name of the entitlement attribute
	Attribute *string `json:"attribute,omitempty"`
	// The value of the entitlement
	Value *string `json:"value,omitempty"`
	// A date-time in ISO-8601 format
	Modified NullableTime `json:"modified,omitempty"`
	// A date-time in ISO-8601 format
	Synced NullableTime `json:"synced,omitempty"`
	// The display name of the entitlement
	DisplayName *string `json:"displayName,omitempty"`
	Source *Reference `json:"source,omitempty"`
	Privileged *bool `json:"privileged,omitempty"`
	IdentityCount *int32 `json:"identityCount,omitempty"`
	Tags []string `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementDocument EntitlementDocument

// NewEntitlementDocument instantiates a new EntitlementDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementDocument(id string, name string, type_ DocumentType) *EntitlementDocument {
	this := EntitlementDocument{}
	this.Id = id
	this.Name = name
	this.Type = type_
	return &this
}

// NewEntitlementDocumentWithDefaults instantiates a new EntitlementDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementDocumentWithDefaults() *EntitlementDocument {
	this := EntitlementDocument{}
	return &this
}

// GetId returns the Id field value
func (o *EntitlementDocument) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EntitlementDocument) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EntitlementDocument) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *EntitlementDocument) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EntitlementDocument) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EntitlementDocument) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *EntitlementDocument) GetType() DocumentType {
	if o == nil {
		var ret DocumentType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EntitlementDocument) GetTypeOk() (*DocumentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EntitlementDocument) SetType(v DocumentType) {
	o.Type = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EntitlementDocument) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocument) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EntitlementDocument) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EntitlementDocument) SetDescription(v string) {
	o.Description = &v
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *EntitlementDocument) GetAttribute() string {
	if o == nil || isNil(o.Attribute) {
		var ret string
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocument) GetAttributeOk() (*string, bool) {
	if o == nil || isNil(o.Attribute) {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *EntitlementDocument) HasAttribute() bool {
	if o != nil && !isNil(o.Attribute) {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given string and assigns it to the Attribute field.
func (o *EntitlementDocument) SetAttribute(v string) {
	o.Attribute = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EntitlementDocument) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocument) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EntitlementDocument) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *EntitlementDocument) SetValue(v string) {
	o.Value = &v
}

// GetModified returns the Modified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntitlementDocument) GetModified() time.Time {
	if o == nil || isNil(o.Modified.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Modified.Get()
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntitlementDocument) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modified.Get(), o.Modified.IsSet()
}

// HasModified returns a boolean if a field has been set.
func (o *EntitlementDocument) HasModified() bool {
	if o != nil && o.Modified.IsSet() {
		return true
	}

	return false
}

// SetModified gets a reference to the given NullableTime and assigns it to the Modified field.
func (o *EntitlementDocument) SetModified(v time.Time) {
	o.Modified.Set(&v)
}
// SetModifiedNil sets the value for Modified to be an explicit nil
func (o *EntitlementDocument) SetModifiedNil() {
	o.Modified.Set(nil)
}

// UnsetModified ensures that no value is present for Modified, not even an explicit nil
func (o *EntitlementDocument) UnsetModified() {
	o.Modified.Unset()
}

// GetSynced returns the Synced field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntitlementDocument) GetSynced() time.Time {
	if o == nil || isNil(o.Synced.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Synced.Get()
}

// GetSyncedOk returns a tuple with the Synced field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntitlementDocument) GetSyncedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Synced.Get(), o.Synced.IsSet()
}

// HasSynced returns a boolean if a field has been set.
func (o *EntitlementDocument) HasSynced() bool {
	if o != nil && o.Synced.IsSet() {
		return true
	}

	return false
}

// SetSynced gets a reference to the given NullableTime and assigns it to the Synced field.
func (o *EntitlementDocument) SetSynced(v time.Time) {
	o.Synced.Set(&v)
}
// SetSyncedNil sets the value for Synced to be an explicit nil
func (o *EntitlementDocument) SetSyncedNil() {
	o.Synced.Set(nil)
}

// UnsetSynced ensures that no value is present for Synced, not even an explicit nil
func (o *EntitlementDocument) UnsetSynced() {
	o.Synced.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *EntitlementDocument) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocument) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *EntitlementDocument) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *EntitlementDocument) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *EntitlementDocument) GetSource() Reference {
	if o == nil || isNil(o.Source) {
		var ret Reference
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocument) GetSourceOk() (*Reference, bool) {
	if o == nil || isNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *EntitlementDocument) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given Reference and assigns it to the Source field.
func (o *EntitlementDocument) SetSource(v Reference) {
	o.Source = &v
}

// GetPrivileged returns the Privileged field value if set, zero value otherwise.
func (o *EntitlementDocument) GetPrivileged() bool {
	if o == nil || isNil(o.Privileged) {
		var ret bool
		return ret
	}
	return *o.Privileged
}

// GetPrivilegedOk returns a tuple with the Privileged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocument) GetPrivilegedOk() (*bool, bool) {
	if o == nil || isNil(o.Privileged) {
		return nil, false
	}
	return o.Privileged, true
}

// HasPrivileged returns a boolean if a field has been set.
func (o *EntitlementDocument) HasPrivileged() bool {
	if o != nil && !isNil(o.Privileged) {
		return true
	}

	return false
}

// SetPrivileged gets a reference to the given bool and assigns it to the Privileged field.
func (o *EntitlementDocument) SetPrivileged(v bool) {
	o.Privileged = &v
}

// GetIdentityCount returns the IdentityCount field value if set, zero value otherwise.
func (o *EntitlementDocument) GetIdentityCount() int32 {
	if o == nil || isNil(o.IdentityCount) {
		var ret int32
		return ret
	}
	return *o.IdentityCount
}

// GetIdentityCountOk returns a tuple with the IdentityCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocument) GetIdentityCountOk() (*int32, bool) {
	if o == nil || isNil(o.IdentityCount) {
		return nil, false
	}
	return o.IdentityCount, true
}

// HasIdentityCount returns a boolean if a field has been set.
func (o *EntitlementDocument) HasIdentityCount() bool {
	if o != nil && !isNil(o.IdentityCount) {
		return true
	}

	return false
}

// SetIdentityCount gets a reference to the given int32 and assigns it to the IdentityCount field.
func (o *EntitlementDocument) SetIdentityCount(v int32) {
	o.IdentityCount = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *EntitlementDocument) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocument) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *EntitlementDocument) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *EntitlementDocument) SetTags(v []string) {
	o.Tags = v
}

func (o EntitlementDocument) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["_type"] = o.Type
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Attribute) {
		toSerialize["attribute"] = o.Attribute
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if o.Modified.IsSet() {
		toSerialize["modified"] = o.Modified.Get()
	}
	if o.Synced.IsSet() {
		toSerialize["synced"] = o.Synced.Get()
	}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !isNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !isNil(o.Privileged) {
		toSerialize["privileged"] = o.Privileged
	}
	if !isNil(o.IdentityCount) {
		toSerialize["identityCount"] = o.IdentityCount
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntitlementDocument) UnmarshalJSON(bytes []byte) (err error) {
	varEntitlementDocument := _EntitlementDocument{}

	if err = json.Unmarshal(bytes, &varEntitlementDocument); err == nil {
		*o = EntitlementDocument(varEntitlementDocument)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "_type")
		delete(additionalProperties, "description")
		delete(additionalProperties, "attribute")
		delete(additionalProperties, "value")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "synced")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "source")
		delete(additionalProperties, "privileged")
		delete(additionalProperties, "identityCount")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementDocument struct {
	value *EntitlementDocument
	isSet bool
}

func (v NullableEntitlementDocument) Get() *EntitlementDocument {
	return v.value
}

func (v *NullableEntitlementDocument) Set(val *EntitlementDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementDocument(val *EntitlementDocument) *NullableEntitlementDocument {
	return &NullableEntitlementDocument{value: val, isSet: true}
}

func (v NullableEntitlementDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


