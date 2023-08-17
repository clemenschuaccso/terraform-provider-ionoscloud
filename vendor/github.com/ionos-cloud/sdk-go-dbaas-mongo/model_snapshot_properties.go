/*
 * IONOS DBaaS MongoDB REST API
 *
 * With IONOS Cloud Database as a Service, you have the ability to quickly set up and manage a MongoDB database. You can also delete clusters, manage backups and users via the API.  MongoDB is an open source, cross-platform, document-oriented database program. Classified as a NoSQL database program, it uses JSON-like documents with optional schemas.  The MongoDB API allows you to create additional database clusters or modify existing ones. Both tools, the Data Center Designer (DCD) and the API use the same concepts consistently and are well suited for smooth and intuitive use.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
	"time"
)

// SnapshotProperties Properties of a snapshot.
type SnapshotProperties struct {
	// The MongoDB version this backup was created from.
	Version *string `json:"version,omitempty"`
	// The size of the snapshot in Mebibytes.
	Size *int32 `json:"size,omitempty"`
	// The date the resource was created.
	CreationTime *IonosTime `json:"creationTime,omitempty"`
}

// NewSnapshotProperties instantiates a new SnapshotProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotProperties() *SnapshotProperties {
	this := SnapshotProperties{}

	return &this
}

// NewSnapshotPropertiesWithDefaults instantiates a new SnapshotProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotPropertiesWithDefaults() *SnapshotProperties {
	this := SnapshotProperties{}
	return &this
}

// GetVersion returns the Version field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SnapshotProperties) GetVersion() *string {
	if o == nil {
		return nil
	}

	return o.Version

}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotProperties) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Version, true
}

// SetVersion sets field value
func (o *SnapshotProperties) SetVersion(v string) {

	o.Version = &v

}

// HasVersion returns a boolean if a field has been set.
func (o *SnapshotProperties) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// GetSize returns the Size field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *SnapshotProperties) GetSize() *int32 {
	if o == nil {
		return nil
	}

	return o.Size

}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotProperties) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Size, true
}

// SetSize sets field value
func (o *SnapshotProperties) SetSize(v int32) {

	o.Size = &v

}

// HasSize returns a boolean if a field has been set.
func (o *SnapshotProperties) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// GetCreationTime returns the CreationTime field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *SnapshotProperties) GetCreationTime() *time.Time {
	if o == nil {
		return nil
	}

	if o.CreationTime == nil {
		return nil
	}
	return &o.CreationTime.Time

}

// GetCreationTimeOk returns a tuple with the CreationTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotProperties) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}

	if o.CreationTime == nil {
		return nil, false
	}
	return &o.CreationTime.Time, true

}

// SetCreationTime sets field value
func (o *SnapshotProperties) SetCreationTime(v time.Time) {

	o.CreationTime = &IonosTime{v}

}

// HasCreationTime returns a boolean if a field has been set.
func (o *SnapshotProperties) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

func (o SnapshotProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.CreationTime != nil {
		toSerialize["creationTime"] = o.CreationTime
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshotProperties struct {
	value *SnapshotProperties
	isSet bool
}

func (v NullableSnapshotProperties) Get() *SnapshotProperties {
	return v.value
}

func (v *NullableSnapshotProperties) Set(val *SnapshotProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotProperties(val *SnapshotProperties) *NullableSnapshotProperties {
	return &NullableSnapshotProperties{value: val, isSet: true}
}

func (v NullableSnapshotProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
