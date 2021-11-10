// Copyright 2021 Google LLC
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

package notebooks_test

import (
	"context"

	notebooks "cloud.google.com/go/notebooks/apiv1beta1"
	"google.golang.org/api/iterator"
	notebookspb "google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1"
)

func ExampleNewNotebookClient() {
	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleNotebookClient_ListInstances() {
	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &notebookspb.ListInstancesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1#ListInstancesRequest.
	}
	it := c.ListInstances(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleNotebookClient_GetInstance() {
	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &notebookspb.GetInstanceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1#GetInstanceRequest.
	}
	resp, err := c.GetInstance(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNotebookClient_CreateInstance() {
	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &notebookspb.CreateInstanceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1#CreateInstanceRequest.
	}
	op, err := c.CreateInstance(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNotebookClient_RegisterInstance() {
	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &notebookspb.RegisterInstanceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1#RegisterInstanceRequest.
	}
	op, err := c.RegisterInstance(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNotebookClient_SetInstanceAccelerator() {
	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &notebookspb.SetInstanceAcceleratorRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1#SetInstanceAcceleratorRequest.
	}
	op, err := c.SetInstanceAccelerator(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNotebookClient_SetInstanceMachineType() {
	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &notebookspb.SetInstanceMachineTypeRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1#SetInstanceMachineTypeRequest.
	}
	op, err := c.SetInstanceMachineType(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNotebookClient_SetInstanceLabels() {
	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &notebookspb.SetInstanceLabelsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1#SetInstanceLabelsRequest.
	}
	op, err := c.SetInstanceLabels(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNotebookClient_DeleteInstance() {
	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &notebookspb.DeleteInstanceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1#DeleteInstanceRequest.
	}
	op, err := c.DeleteInstance(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleNotebookClient_StartInstance() {
	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &notebookspb.StartInstanceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1#StartInstanceRequest.
	}
	op, err := c.StartInstance(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNotebookClient_StopInstance() {
	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &notebookspb.StopInstanceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1#StopInstanceRequest.
	}
	op, err := c.StopInstance(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNotebookClient_ResetInstance() {
	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &notebookspb.ResetInstanceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1#ResetInstanceRequest.
	}
	op, err := c.ResetInstance(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNotebookClient_ReportInstanceInfo() {
	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &notebookspb.ReportInstanceInfoRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1#ReportInstanceInfoRequest.
	}
	op, err := c.ReportInstanceInfo(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNotebookClient_IsInstanceUpgradeable() {
	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &notebookspb.IsInstanceUpgradeableRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1#IsInstanceUpgradeableRequest.
	}
	resp, err := c.IsInstanceUpgradeable(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNotebookClient_UpgradeInstance() {
	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &notebookspb.UpgradeInstanceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1#UpgradeInstanceRequest.
	}
	op, err := c.UpgradeInstance(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNotebookClient_UpgradeInstanceInternal() {
	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &notebookspb.UpgradeInstanceInternalRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1#UpgradeInstanceInternalRequest.
	}
	op, err := c.UpgradeInstanceInternal(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNotebookClient_ListEnvironments() {
	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &notebookspb.ListEnvironmentsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1#ListEnvironmentsRequest.
	}
	it := c.ListEnvironments(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleNotebookClient_GetEnvironment() {
	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &notebookspb.GetEnvironmentRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1#GetEnvironmentRequest.
	}
	resp, err := c.GetEnvironment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNotebookClient_CreateEnvironment() {
	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &notebookspb.CreateEnvironmentRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1#CreateEnvironmentRequest.
	}
	op, err := c.CreateEnvironment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNotebookClient_DeleteEnvironment() {
	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &notebookspb.DeleteEnvironmentRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1#DeleteEnvironmentRequest.
	}
	op, err := c.DeleteEnvironment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}
