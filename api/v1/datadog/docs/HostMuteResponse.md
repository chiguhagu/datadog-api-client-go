# HostMuteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**End** | Pointer to **int64** | POSIX timestamp in seconds when the host is unmuted. | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### GetAction

`func (o *HostMuteResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *HostMuteResponse) GetActionOk() (string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAction

`func (o *HostMuteResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetAction

`func (o *HostMuteResponse) SetAction(v string)`

SetAction gets a reference to the given string and assigns it to the Action field.

### GetEnd

`func (o *HostMuteResponse) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *HostMuteResponse) GetEndOk() (int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnd

`func (o *HostMuteResponse) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEnd

`func (o *HostMuteResponse) SetEnd(v int64)`

SetEnd gets a reference to the given int64 and assigns it to the End field.

### GetHostname

`func (o *HostMuteResponse) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *HostMuteResponse) GetHostnameOk() (string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHostname

`func (o *HostMuteResponse) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### SetHostname

`func (o *HostMuteResponse) SetHostname(v string)`

SetHostname gets a reference to the given string and assigns it to the Hostname field.

### GetMessage

`func (o *HostMuteResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HostMuteResponse) GetMessageOk() (string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessage

`func (o *HostMuteResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessage

`func (o *HostMuteResponse) SetMessage(v string)`

SetMessage gets a reference to the given string and assigns it to the Message field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

