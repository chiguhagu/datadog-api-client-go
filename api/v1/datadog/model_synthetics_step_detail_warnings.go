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

// SyntheticsStepDetailWarnings struct for SyntheticsStepDetailWarnings
type SyntheticsStepDetailWarnings struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

// GetMessage returns the Message field value
func (o *SyntheticsStepDetailWarnings) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// SetMessage sets field value
func (o *SyntheticsStepDetailWarnings) SetMessage(v string) {
	o.Message = v
}

// GetType returns the Type field value
func (o *SyntheticsStepDetailWarnings) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// SetType sets field value
func (o *SyntheticsStepDetailWarnings) SetType(v string) {
	o.Type = v
}

type NullableSyntheticsStepDetailWarnings struct {
	Value        SyntheticsStepDetailWarnings
	ExplicitNull bool
}

func (v NullableSyntheticsStepDetailWarnings) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableSyntheticsStepDetailWarnings) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}