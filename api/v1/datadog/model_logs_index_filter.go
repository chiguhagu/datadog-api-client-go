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

// LogsIndexFilter Only logs matching the filter criteria are considered for this index.
type LogsIndexFilter struct {
	// The search query following the Log search syntax.
	Query *string `json:"query,omitempty"`
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *LogsIndexFilter) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LogsIndexFilter) GetQueryOk() (string, bool) {
	if o == nil || o.Query == nil {
		var ret string
		return ret, false
	}
	return *o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *LogsIndexFilter) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *LogsIndexFilter) SetQuery(v string) {
	o.Query = &v
}

type NullableLogsIndexFilter struct {
	Value        LogsIndexFilter
	ExplicitNull bool
}

func (v NullableLogsIndexFilter) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableLogsIndexFilter) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}