// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package eventgrid

import original "github.com/Azure/azure-sdk-for-go/services/eventgrid/2018-01-01/eventgrid"

type BaseClient = original.BaseClient
type JobState = original.JobState

const (
	Canceled   JobState = original.Canceled
	Canceling  JobState = original.Canceling
	Error      JobState = original.Error
	Finished   JobState = original.Finished
	Processing JobState = original.Processing
	Queued     JobState = original.Queued
	Scheduled  JobState = original.Scheduled
)

type ContainerRegistryEventActor = original.ContainerRegistryEventActor
type ContainerRegistryEventData = original.ContainerRegistryEventData
type ContainerRegistryEventRequest = original.ContainerRegistryEventRequest
type ContainerRegistryEventSource = original.ContainerRegistryEventSource
type ContainerRegistryEventTarget = original.ContainerRegistryEventTarget
type ContainerRegistryImageDeletedEventData = original.ContainerRegistryImageDeletedEventData
type ContainerRegistryImagePushedEventData = original.ContainerRegistryImagePushedEventData
type DeviceConnectionStateEventInfo = original.DeviceConnectionStateEventInfo
type DeviceConnectionStateEventProperties = original.DeviceConnectionStateEventProperties
type DeviceLifeCycleEventProperties = original.DeviceLifeCycleEventProperties
type DeviceTwinInfo = original.DeviceTwinInfo
type DeviceTwinInfoProperties = original.DeviceTwinInfoProperties
type DeviceTwinInfoX509Thumbprint = original.DeviceTwinInfoX509Thumbprint
type DeviceTwinMetadata = original.DeviceTwinMetadata
type DeviceTwinProperties = original.DeviceTwinProperties
type Event = original.Event
type EventHubCaptureFileCreatedEventData = original.EventHubCaptureFileCreatedEventData
type IotHubDeviceConnectedEventData = original.IotHubDeviceConnectedEventData
type IotHubDeviceCreatedEventData = original.IotHubDeviceCreatedEventData
type IotHubDeviceDeletedEventData = original.IotHubDeviceDeletedEventData
type IotHubDeviceDisconnectedEventData = original.IotHubDeviceDisconnectedEventData
type MediaJobStateChangeEventData = original.MediaJobStateChangeEventData
type ResourceDeleteCancelData = original.ResourceDeleteCancelData
type ResourceDeleteFailureData = original.ResourceDeleteFailureData
type ResourceDeleteSuccessData = original.ResourceDeleteSuccessData
type ResourceWriteCancelData = original.ResourceWriteCancelData
type ResourceWriteFailureData = original.ResourceWriteFailureData
type ResourceWriteSuccessData = original.ResourceWriteSuccessData
type ServiceBusActiveMessagesAvailableWithNoListenersEventData = original.ServiceBusActiveMessagesAvailableWithNoListenersEventData
type ServiceBusDeadletterMessagesAvailableWithNoListenersEventData = original.ServiceBusDeadletterMessagesAvailableWithNoListenersEventData
type StorageBlobCreatedEventData = original.StorageBlobCreatedEventData
type StorageBlobDeletedEventData = original.StorageBlobDeletedEventData
type SubscriptionDeletedEventData = original.SubscriptionDeletedEventData
type SubscriptionValidationEventData = original.SubscriptionValidationEventData
type SubscriptionValidationResponse = original.SubscriptionValidationResponse

func New() BaseClient {
	return original.New()
}
func NewWithoutDefaults() BaseClient {
	return original.NewWithoutDefaults()
}
func PossibleJobStateValues() []JobState {
	return original.PossibleJobStateValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
