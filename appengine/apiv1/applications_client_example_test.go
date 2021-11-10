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

package appengine_test

import (
	"context"

	appengine "cloud.google.com/go/appengine/apiv1"
	appenginepb "google.golang.org/genproto/googleapis/appengine/v1"
)

func ExampleNewApplicationsClient() {
	ctx := context.Background()
	c, err := appengine.NewApplicationsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleApplicationsClient_GetApplication() {
	ctx := context.Background()
	c, err := appengine.NewApplicationsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &appenginepb.GetApplicationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/appengine/v1#GetApplicationRequest.
	}
	resp, err := c.GetApplication(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleApplicationsClient_CreateApplication() {
	ctx := context.Background()
	c, err := appengine.NewApplicationsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &appenginepb.CreateApplicationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/appengine/v1#CreateApplicationRequest.
	}
	op, err := c.CreateApplication(ctx, req)
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

func ExampleApplicationsClient_UpdateApplication() {
	ctx := context.Background()
	c, err := appengine.NewApplicationsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &appenginepb.UpdateApplicationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/appengine/v1#UpdateApplicationRequest.
	}
	op, err := c.UpdateApplication(ctx, req)
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

func ExampleApplicationsClient_RepairApplication() {
	ctx := context.Background()
	c, err := appengine.NewApplicationsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &appenginepb.RepairApplicationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/appengine/v1#RepairApplicationRequest.
	}
	op, err := c.RepairApplication(ctx, req)
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
