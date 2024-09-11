// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

//go:build go1.23

package vmwareengine

import (
	"iter"

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	vmwareenginepb "cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb"
	"github.com/googleapis/gax-go/v2/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
)

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ClusterIterator) All() iter.Seq2[*vmwareenginepb.Cluster, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ExternalAccessRuleIterator) All() iter.Seq2[*vmwareenginepb.ExternalAccessRule, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ExternalAddressIterator) All() iter.Seq2[*vmwareenginepb.ExternalAddress, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *HcxActivationKeyIterator) All() iter.Seq2[*vmwareenginepb.HcxActivationKey, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *LocationIterator) All() iter.Seq2[*locationpb.Location, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *LoggingServerIterator) All() iter.Seq2[*vmwareenginepb.LoggingServer, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ManagementDnsZoneBindingIterator) All() iter.Seq2[*vmwareenginepb.ManagementDnsZoneBinding, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *NetworkPeeringIterator) All() iter.Seq2[*vmwareenginepb.NetworkPeering, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *NetworkPolicyIterator) All() iter.Seq2[*vmwareenginepb.NetworkPolicy, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *NodeIterator) All() iter.Seq2[*vmwareenginepb.Node, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *NodeTypeIterator) All() iter.Seq2[*vmwareenginepb.NodeType, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *OperationIterator) All() iter.Seq2[*longrunningpb.Operation, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *PeeringRouteIterator) All() iter.Seq2[*vmwareenginepb.PeeringRoute, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *PrivateCloudIterator) All() iter.Seq2[*vmwareenginepb.PrivateCloud, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *PrivateConnectionIterator) All() iter.Seq2[*vmwareenginepb.PrivateConnection, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *SubnetIterator) All() iter.Seq2[*vmwareenginepb.Subnet, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *VmwareEngineNetworkIterator) All() iter.Seq2[*vmwareenginepb.VmwareEngineNetwork, error] {
	return iterator.RangeAdapter(it.Next)
}
