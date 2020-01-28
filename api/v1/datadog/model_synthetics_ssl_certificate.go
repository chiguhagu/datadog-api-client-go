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
	"time"
)

// SyntheticsSslCertificate struct for SyntheticsSslCertificate
type SyntheticsSslCertificate struct {
	Cipher         *string                          `json:"cipher,omitempty"`
	Exponent       *float64                         `json:"exponent,omitempty"`
	ExtKeyUsage    *[]string                        `json:"extKeyUsage,omitempty"`
	Fingerprint    *string                          `json:"fingerprint,omitempty"`
	Fingerprint256 *string                          `json:"fingerprint256,omitempty"`
	Issuer         *SyntheticsSslCertificateIssuer  `json:"issuer,omitempty"`
	Modulus        *string                          `json:"modulus,omitempty"`
	Protocol       *string                          `json:"protocol,omitempty"`
	SerialNumber   *string                          `json:"serialNumber,omitempty"`
	Subject        *SyntheticsSslCertificateSubject `json:"subject,omitempty"`
	ValidFrom      *time.Time                       `json:"validFrom,omitempty"`
	ValidTo        *time.Time                       `json:"validTo,omitempty"`
}

// GetCipher returns the Cipher field value if set, zero value otherwise.
func (o *SyntheticsSslCertificate) GetCipher() string {
	if o == nil || o.Cipher == nil {
		var ret string
		return ret
	}
	return *o.Cipher
}

// GetCipherOk returns a tuple with the Cipher field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSslCertificate) GetCipherOk() (string, bool) {
	if o == nil || o.Cipher == nil {
		var ret string
		return ret, false
	}
	return *o.Cipher, true
}

// HasCipher returns a boolean if a field has been set.
func (o *SyntheticsSslCertificate) HasCipher() bool {
	if o != nil && o.Cipher != nil {
		return true
	}

	return false
}

// SetCipher gets a reference to the given string and assigns it to the Cipher field.
func (o *SyntheticsSslCertificate) SetCipher(v string) {
	o.Cipher = &v
}

// GetExponent returns the Exponent field value if set, zero value otherwise.
func (o *SyntheticsSslCertificate) GetExponent() float64 {
	if o == nil || o.Exponent == nil {
		var ret float64
		return ret
	}
	return *o.Exponent
}

// GetExponentOk returns a tuple with the Exponent field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSslCertificate) GetExponentOk() (float64, bool) {
	if o == nil || o.Exponent == nil {
		var ret float64
		return ret, false
	}
	return *o.Exponent, true
}

// HasExponent returns a boolean if a field has been set.
func (o *SyntheticsSslCertificate) HasExponent() bool {
	if o != nil && o.Exponent != nil {
		return true
	}

	return false
}

// SetExponent gets a reference to the given float64 and assigns it to the Exponent field.
func (o *SyntheticsSslCertificate) SetExponent(v float64) {
	o.Exponent = &v
}

// GetExtKeyUsage returns the ExtKeyUsage field value if set, zero value otherwise.
func (o *SyntheticsSslCertificate) GetExtKeyUsage() []string {
	if o == nil || o.ExtKeyUsage == nil {
		var ret []string
		return ret
	}
	return *o.ExtKeyUsage
}

// GetExtKeyUsageOk returns a tuple with the ExtKeyUsage field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSslCertificate) GetExtKeyUsageOk() ([]string, bool) {
	if o == nil || o.ExtKeyUsage == nil {
		var ret []string
		return ret, false
	}
	return *o.ExtKeyUsage, true
}

// HasExtKeyUsage returns a boolean if a field has been set.
func (o *SyntheticsSslCertificate) HasExtKeyUsage() bool {
	if o != nil && o.ExtKeyUsage != nil {
		return true
	}

	return false
}

// SetExtKeyUsage gets a reference to the given []string and assigns it to the ExtKeyUsage field.
func (o *SyntheticsSslCertificate) SetExtKeyUsage(v []string) {
	o.ExtKeyUsage = &v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *SyntheticsSslCertificate) GetFingerprint() string {
	if o == nil || o.Fingerprint == nil {
		var ret string
		return ret
	}
	return *o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSslCertificate) GetFingerprintOk() (string, bool) {
	if o == nil || o.Fingerprint == nil {
		var ret string
		return ret, false
	}
	return *o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *SyntheticsSslCertificate) HasFingerprint() bool {
	if o != nil && o.Fingerprint != nil {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *SyntheticsSslCertificate) SetFingerprint(v string) {
	o.Fingerprint = &v
}

// GetFingerprint256 returns the Fingerprint256 field value if set, zero value otherwise.
func (o *SyntheticsSslCertificate) GetFingerprint256() string {
	if o == nil || o.Fingerprint256 == nil {
		var ret string
		return ret
	}
	return *o.Fingerprint256
}

// GetFingerprint256Ok returns a tuple with the Fingerprint256 field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSslCertificate) GetFingerprint256Ok() (string, bool) {
	if o == nil || o.Fingerprint256 == nil {
		var ret string
		return ret, false
	}
	return *o.Fingerprint256, true
}

// HasFingerprint256 returns a boolean if a field has been set.
func (o *SyntheticsSslCertificate) HasFingerprint256() bool {
	if o != nil && o.Fingerprint256 != nil {
		return true
	}

	return false
}

// SetFingerprint256 gets a reference to the given string and assigns it to the Fingerprint256 field.
func (o *SyntheticsSslCertificate) SetFingerprint256(v string) {
	o.Fingerprint256 = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *SyntheticsSslCertificate) GetIssuer() SyntheticsSslCertificateIssuer {
	if o == nil || o.Issuer == nil {
		var ret SyntheticsSslCertificateIssuer
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSslCertificate) GetIssuerOk() (SyntheticsSslCertificateIssuer, bool) {
	if o == nil || o.Issuer == nil {
		var ret SyntheticsSslCertificateIssuer
		return ret, false
	}
	return *o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *SyntheticsSslCertificate) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given SyntheticsSslCertificateIssuer and assigns it to the Issuer field.
func (o *SyntheticsSslCertificate) SetIssuer(v SyntheticsSslCertificateIssuer) {
	o.Issuer = &v
}

// GetModulus returns the Modulus field value if set, zero value otherwise.
func (o *SyntheticsSslCertificate) GetModulus() string {
	if o == nil || o.Modulus == nil {
		var ret string
		return ret
	}
	return *o.Modulus
}

// GetModulusOk returns a tuple with the Modulus field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSslCertificate) GetModulusOk() (string, bool) {
	if o == nil || o.Modulus == nil {
		var ret string
		return ret, false
	}
	return *o.Modulus, true
}

// HasModulus returns a boolean if a field has been set.
func (o *SyntheticsSslCertificate) HasModulus() bool {
	if o != nil && o.Modulus != nil {
		return true
	}

	return false
}

// SetModulus gets a reference to the given string and assigns it to the Modulus field.
func (o *SyntheticsSslCertificate) SetModulus(v string) {
	o.Modulus = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *SyntheticsSslCertificate) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSslCertificate) GetProtocolOk() (string, bool) {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret, false
	}
	return *o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *SyntheticsSslCertificate) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *SyntheticsSslCertificate) SetProtocol(v string) {
	o.Protocol = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *SyntheticsSslCertificate) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSslCertificate) GetSerialNumberOk() (string, bool) {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret, false
	}
	return *o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *SyntheticsSslCertificate) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *SyntheticsSslCertificate) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *SyntheticsSslCertificate) GetSubject() SyntheticsSslCertificateSubject {
	if o == nil || o.Subject == nil {
		var ret SyntheticsSslCertificateSubject
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSslCertificate) GetSubjectOk() (SyntheticsSslCertificateSubject, bool) {
	if o == nil || o.Subject == nil {
		var ret SyntheticsSslCertificateSubject
		return ret, false
	}
	return *o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *SyntheticsSslCertificate) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given SyntheticsSslCertificateSubject and assigns it to the Subject field.
func (o *SyntheticsSslCertificate) SetSubject(v SyntheticsSslCertificateSubject) {
	o.Subject = &v
}

// GetValidFrom returns the ValidFrom field value if set, zero value otherwise.
func (o *SyntheticsSslCertificate) GetValidFrom() time.Time {
	if o == nil || o.ValidFrom == nil {
		var ret time.Time
		return ret
	}
	return *o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSslCertificate) GetValidFromOk() (time.Time, bool) {
	if o == nil || o.ValidFrom == nil {
		var ret time.Time
		return ret, false
	}
	return *o.ValidFrom, true
}

// HasValidFrom returns a boolean if a field has been set.
func (o *SyntheticsSslCertificate) HasValidFrom() bool {
	if o != nil && o.ValidFrom != nil {
		return true
	}

	return false
}

// SetValidFrom gets a reference to the given time.Time and assigns it to the ValidFrom field.
func (o *SyntheticsSslCertificate) SetValidFrom(v time.Time) {
	o.ValidFrom = &v
}

// GetValidTo returns the ValidTo field value if set, zero value otherwise.
func (o *SyntheticsSslCertificate) GetValidTo() time.Time {
	if o == nil || o.ValidTo == nil {
		var ret time.Time
		return ret
	}
	return *o.ValidTo
}

// GetValidToOk returns a tuple with the ValidTo field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSslCertificate) GetValidToOk() (time.Time, bool) {
	if o == nil || o.ValidTo == nil {
		var ret time.Time
		return ret, false
	}
	return *o.ValidTo, true
}

// HasValidTo returns a boolean if a field has been set.
func (o *SyntheticsSslCertificate) HasValidTo() bool {
	if o != nil && o.ValidTo != nil {
		return true
	}

	return false
}

// SetValidTo gets a reference to the given time.Time and assigns it to the ValidTo field.
func (o *SyntheticsSslCertificate) SetValidTo(v time.Time) {
	o.ValidTo = &v
}

type NullableSyntheticsSslCertificate struct {
	Value        SyntheticsSslCertificate
	ExplicitNull bool
}

func (v NullableSyntheticsSslCertificate) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableSyntheticsSslCertificate) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
