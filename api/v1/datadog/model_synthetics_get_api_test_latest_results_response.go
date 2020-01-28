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

// SyntheticsGetApiTestLatestResultsResponse struct for SyntheticsGetApiTestLatestResultsResponse
type SyntheticsGetApiTestLatestResultsResponse struct {
	LastTimestampFetched *int64                          `json:"last_timestamp_fetched,omitempty"`
	Results              *[]SyntheticsApiTestResultShort `json:"results,omitempty"`
}

// GetLastTimestampFetched returns the LastTimestampFetched field value if set, zero value otherwise.
func (o *SyntheticsGetApiTestLatestResultsResponse) GetLastTimestampFetched() int64 {
	if o == nil || o.LastTimestampFetched == nil {
		var ret int64
		return ret
	}
	return *o.LastTimestampFetched
}

// GetLastTimestampFetchedOk returns a tuple with the LastTimestampFetched field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsGetApiTestLatestResultsResponse) GetLastTimestampFetchedOk() (int64, bool) {
	if o == nil || o.LastTimestampFetched == nil {
		var ret int64
		return ret, false
	}
	return *o.LastTimestampFetched, true
}

// HasLastTimestampFetched returns a boolean if a field has been set.
func (o *SyntheticsGetApiTestLatestResultsResponse) HasLastTimestampFetched() bool {
	if o != nil && o.LastTimestampFetched != nil {
		return true
	}

	return false
}

// SetLastTimestampFetched gets a reference to the given int64 and assigns it to the LastTimestampFetched field.
func (o *SyntheticsGetApiTestLatestResultsResponse) SetLastTimestampFetched(v int64) {
	o.LastTimestampFetched = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *SyntheticsGetApiTestLatestResultsResponse) GetResults() []SyntheticsApiTestResultShort {
	if o == nil || o.Results == nil {
		var ret []SyntheticsApiTestResultShort
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsGetApiTestLatestResultsResponse) GetResultsOk() ([]SyntheticsApiTestResultShort, bool) {
	if o == nil || o.Results == nil {
		var ret []SyntheticsApiTestResultShort
		return ret, false
	}
	return *o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *SyntheticsGetApiTestLatestResultsResponse) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []SyntheticsApiTestResultShort and assigns it to the Results field.
func (o *SyntheticsGetApiTestLatestResultsResponse) SetResults(v []SyntheticsApiTestResultShort) {
	o.Results = &v
}

type NullableSyntheticsGetApiTestLatestResultsResponse struct {
	Value        SyntheticsGetApiTestLatestResultsResponse
	ExplicitNull bool
}

func (v NullableSyntheticsGetApiTestLatestResultsResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableSyntheticsGetApiTestLatestResultsResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
