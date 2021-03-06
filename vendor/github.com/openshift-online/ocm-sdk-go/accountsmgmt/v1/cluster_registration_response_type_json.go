/*
Copyright (c) 2019 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// MarshalClusterRegistrationResponse writes a value of the 'cluster_registration_response' type to the given writer.
func MarshalClusterRegistrationResponse(object *ClusterRegistrationResponse, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeClusterRegistrationResponse(object, stream)
	stream.Flush()
	return stream.Error
}

// writeClusterRegistrationResponse writes a value of the 'cluster_registration_response' type to the given stream.
func writeClusterRegistrationResponse(object *ClusterRegistrationResponse, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	if object.accountID != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("account_id")
		stream.WriteString(*object.accountID)
		count++
	}
	if object.authorizationToken != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("authorization_token")
		stream.WriteString(*object.authorizationToken)
		count++
	}
	if object.clusterID != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("cluster_id")
		stream.WriteString(*object.clusterID)
		count++
	}
	if object.expiresAt != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("expires_at")
		stream.WriteString(*object.expiresAt)
		count++
	}
	stream.WriteObjectEnd()
}

// UnmarshalClusterRegistrationResponse reads a value of the 'cluster_registration_response' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalClusterRegistrationResponse(source interface{}) (object *ClusterRegistrationResponse, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = readClusterRegistrationResponse(iterator)
	err = iterator.Error
	return
}

// readClusterRegistrationResponse reads a value of the 'cluster_registration_response' type from the given iterator.
func readClusterRegistrationResponse(iterator *jsoniter.Iterator) *ClusterRegistrationResponse {
	object := &ClusterRegistrationResponse{}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "account_id":
			value := iterator.ReadString()
			object.accountID = &value
		case "authorization_token":
			value := iterator.ReadString()
			object.authorizationToken = &value
		case "cluster_id":
			value := iterator.ReadString()
			object.clusterID = &value
		case "expires_at":
			value := iterator.ReadString()
			object.expiresAt = &value
		default:
			iterator.ReadAny()
		}
	}
	return object
}
