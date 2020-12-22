/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// SecurityMonitoringRuleQueryAggregation The aggregation type.
type SecurityMonitoringRuleQueryAggregation string

// List of SecurityMonitoringRuleQueryAggregation
const (
	SECURITYMONITORINGRULEQUERYAGGREGATION_COUNT       SecurityMonitoringRuleQueryAggregation = "count"
	SECURITYMONITORINGRULEQUERYAGGREGATION_CARDINALITY SecurityMonitoringRuleQueryAggregation = "cardinality"
	SECURITYMONITORINGRULEQUERYAGGREGATION_SUM         SecurityMonitoringRuleQueryAggregation = "sum"
	SECURITYMONITORINGRULEQUERYAGGREGATION_MAX         SecurityMonitoringRuleQueryAggregation = "max"
)

var allowedSecurityMonitoringRuleQueryAggregationEnumValues = []SecurityMonitoringRuleQueryAggregation{
	"count",
	"cardinality",
	"sum",
	"max",
}

func (v *SecurityMonitoringRuleQueryAggregation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SecurityMonitoringRuleQueryAggregation(value)
	for _, existing := range allowedSecurityMonitoringRuleQueryAggregationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SecurityMonitoringRuleQueryAggregation", value)
}

// NewSecurityMonitoringRuleQueryAggregationFromValue returns a pointer to a valid SecurityMonitoringRuleQueryAggregation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSecurityMonitoringRuleQueryAggregationFromValue(v string) (*SecurityMonitoringRuleQueryAggregation, error) {
	ev := SecurityMonitoringRuleQueryAggregation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringRuleQueryAggregation: valid values are %v", v, allowedSecurityMonitoringRuleQueryAggregationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SecurityMonitoringRuleQueryAggregation) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringRuleQueryAggregationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringRuleQueryAggregation value
func (v SecurityMonitoringRuleQueryAggregation) Ptr() *SecurityMonitoringRuleQueryAggregation {
	return &v
}

type NullableSecurityMonitoringRuleQueryAggregation struct {
	value *SecurityMonitoringRuleQueryAggregation
	isSet bool
}

func (v NullableSecurityMonitoringRuleQueryAggregation) Get() *SecurityMonitoringRuleQueryAggregation {
	return v.value
}

func (v *NullableSecurityMonitoringRuleQueryAggregation) Set(val *SecurityMonitoringRuleQueryAggregation) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityMonitoringRuleQueryAggregation) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityMonitoringRuleQueryAggregation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityMonitoringRuleQueryAggregation(val *SecurityMonitoringRuleQueryAggregation) *NullableSecurityMonitoringRuleQueryAggregation {
	return &NullableSecurityMonitoringRuleQueryAggregation{value: val, isSet: true}
}

func (v NullableSecurityMonitoringRuleQueryAggregation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityMonitoringRuleQueryAggregation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
