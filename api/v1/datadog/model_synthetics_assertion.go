/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"bytes"
	"encoding/json"
)

// SyntheticsAssertion struct for SyntheticsAssertion
type SyntheticsAssertion struct {
	Operator SyntheticsAssertionOperator `json:"operator"`
	Property *string                     `json:"property,omitempty"`
	Target   *interface{}                `json:"target,omitempty"`
	Type     SyntheticsAssertionType     `json:"type"`
}

// GetOperator returns the Operator field value
func (o *SyntheticsAssertion) GetOperator() SyntheticsAssertionOperator {
	if o == nil {
		var ret SyntheticsAssertionOperator
		return ret
	}

	return o.Operator
}

// SetOperator sets field value
func (o *SyntheticsAssertion) SetOperator(v SyntheticsAssertionOperator) {
	o.Operator = v
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *SyntheticsAssertion) GetProperty() string {
	if o == nil || o.Property == nil {
		var ret string
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAssertion) GetPropertyOk() (string, bool) {
	if o == nil || o.Property == nil {
		var ret string
		return ret, false
	}
	return *o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *SyntheticsAssertion) HasProperty() bool {
	if o != nil && o.Property != nil {
		return true
	}

	return false
}

// SetProperty gets a reference to the given string and assigns it to the Property field.
func (o *SyntheticsAssertion) SetProperty(v string) {
	o.Property = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *SyntheticsAssertion) GetTarget() interface{} {
	if o == nil || o.Target == nil {
		var ret interface{}
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAssertion) GetTargetOk() (interface{}, bool) {
	if o == nil || o.Target == nil {
		var ret interface{}
		return ret, false
	}
	return *o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *SyntheticsAssertion) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given interface{} and assigns it to the Target field.
func (o *SyntheticsAssertion) SetTarget(v interface{}) {
	o.Target = &v
}

// GetType returns the Type field value
func (o *SyntheticsAssertion) GetType() SyntheticsAssertionType {
	if o == nil {
		var ret SyntheticsAssertionType
		return ret
	}

	return o.Type
}

// SetType sets field value
func (o *SyntheticsAssertion) SetType(v SyntheticsAssertionType) {
	o.Type = v
}

type NullableSyntheticsAssertion struct {
	Value        SyntheticsAssertion
	ExplicitNull bool
}

func (v NullableSyntheticsAssertion) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableSyntheticsAssertion) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
