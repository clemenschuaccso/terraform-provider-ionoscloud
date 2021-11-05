/*
 * CLOUD API
 *
 * IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0-SDK.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// TemplateProperties struct for TemplateProperties
type TemplateProperties struct {
	// A name of that resource
	Name *string `json:"name"`
	// The CPU cores count
	Cores *float32 `json:"cores"`
	// The RAM size in MB
	Ram *float32 `json:"ram"`
	// The storage size in GB
	StorageSize *float32 `json:"storageSize"`
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TemplateProperties) GetName() *string {
	if o == nil {
		return nil
	}

	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Name, true
}

// SetName sets field value
func (o *TemplateProperties) SetName(v string) {

	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *TemplateProperties) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// GetCores returns the Cores field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *TemplateProperties) GetCores() *float32 {
	if o == nil {
		return nil
	}

	return o.Cores

}

// GetCoresOk returns a tuple with the Cores field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateProperties) GetCoresOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Cores, true
}

// SetCores sets field value
func (o *TemplateProperties) SetCores(v float32) {

	o.Cores = &v

}

// HasCores returns a boolean if a field has been set.
func (o *TemplateProperties) HasCores() bool {
	if o != nil && o.Cores != nil {
		return true
	}

	return false
}

// GetRam returns the Ram field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *TemplateProperties) GetRam() *float32 {
	if o == nil {
		return nil
	}

	return o.Ram

}

// GetRamOk returns a tuple with the Ram field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateProperties) GetRamOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Ram, true
}

// SetRam sets field value
func (o *TemplateProperties) SetRam(v float32) {

	o.Ram = &v

}

// HasRam returns a boolean if a field has been set.
func (o *TemplateProperties) HasRam() bool {
	if o != nil && o.Ram != nil {
		return true
	}

	return false
}

// GetStorageSize returns the StorageSize field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *TemplateProperties) GetStorageSize() *float32 {
	if o == nil {
		return nil
	}

	return o.StorageSize

}

// GetStorageSizeOk returns a tuple with the StorageSize field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateProperties) GetStorageSizeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}

	return o.StorageSize, true
}

// SetStorageSize sets field value
func (o *TemplateProperties) SetStorageSize(v float32) {

	o.StorageSize = &v

}

// HasStorageSize returns a boolean if a field has been set.
func (o *TemplateProperties) HasStorageSize() bool {
	if o != nil && o.StorageSize != nil {
		return true
	}

	return false
}

func (o TemplateProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	if o.Cores != nil {
		toSerialize["cores"] = o.Cores
	}

	if o.Ram != nil {
		toSerialize["ram"] = o.Ram
	}

	if o.StorageSize != nil {
		toSerialize["storageSize"] = o.StorageSize
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateProperties struct {
	value *TemplateProperties
	isSet bool
}

func (v NullableTemplateProperties) Get() *TemplateProperties {
	return v.value
}

func (v *NullableTemplateProperties) Set(val *TemplateProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateProperties(val *TemplateProperties) *NullableTemplateProperties {
	return &NullableTemplateProperties{value: val, isSet: true}
}

func (v NullableTemplateProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
