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

package osconfig_test

import (
	"context"

	osconfig "cloud.google.com/go/osconfig/apiv1"
	osconfigpb "cloud.google.com/go/osconfig/apiv1/osconfigpb"
)

func ExampleOsConfigZonalClient_ListInventories_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := osconfig.NewOsConfigZonalClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &osconfigpb.ListInventoriesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/osconfig/apiv1/osconfigpb#ListInventoriesRequest.
	}
	for resp, err := range c.ListInventories(ctx, req).All() {
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleOsConfigZonalClient_ListOSPolicyAssignmentReports_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := osconfig.NewOsConfigZonalClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &osconfigpb.ListOSPolicyAssignmentReportsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/osconfig/apiv1/osconfigpb#ListOSPolicyAssignmentReportsRequest.
	}
	for resp, err := range c.ListOSPolicyAssignmentReports(ctx, req).All() {
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleOsConfigZonalClient_ListOSPolicyAssignmentRevisions_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := osconfig.NewOsConfigZonalClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &osconfigpb.ListOSPolicyAssignmentRevisionsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/osconfig/apiv1/osconfigpb#ListOSPolicyAssignmentRevisionsRequest.
	}
	for resp, err := range c.ListOSPolicyAssignmentRevisions(ctx, req).All() {
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleOsConfigZonalClient_ListOSPolicyAssignments_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := osconfig.NewOsConfigZonalClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &osconfigpb.ListOSPolicyAssignmentsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/osconfig/apiv1/osconfigpb#ListOSPolicyAssignmentsRequest.
	}
	for resp, err := range c.ListOSPolicyAssignments(ctx, req).All() {
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleOsConfigZonalClient_ListVulnerabilityReports_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := osconfig.NewOsConfigZonalClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &osconfigpb.ListVulnerabilityReportsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/osconfig/apiv1/osconfigpb#ListVulnerabilityReportsRequest.
	}
	for resp, err := range c.ListVulnerabilityReports(ctx, req).All() {
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}
