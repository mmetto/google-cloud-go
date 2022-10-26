// Copyright 2022 Google LLC
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

package gkebackup_test

import (
	"context"

	gkebackup "cloud.google.com/go/gkebackup/apiv1"
	"google.golang.org/api/iterator"
	gkebackuppb "google.golang.org/genproto/googleapis/cloud/gkebackup/v1"
)

func ExampleNewBackupForGKEClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkebackup.NewBackupForGKEClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleBackupForGKEClient_CreateBackupPlan() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkebackup.NewBackupForGKEClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gkebackuppb.CreateBackupPlanRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/gkebackup/v1#CreateBackupPlanRequest.
	}
	op, err := c.CreateBackupPlan(ctx, req)
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

func ExampleBackupForGKEClient_ListBackupPlans() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkebackup.NewBackupForGKEClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gkebackuppb.ListBackupPlansRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/gkebackup/v1#ListBackupPlansRequest.
	}
	it := c.ListBackupPlans(ctx, req)
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

func ExampleBackupForGKEClient_GetBackupPlan() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkebackup.NewBackupForGKEClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gkebackuppb.GetBackupPlanRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/gkebackup/v1#GetBackupPlanRequest.
	}
	resp, err := c.GetBackupPlan(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleBackupForGKEClient_UpdateBackupPlan() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkebackup.NewBackupForGKEClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gkebackuppb.UpdateBackupPlanRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/gkebackup/v1#UpdateBackupPlanRequest.
	}
	op, err := c.UpdateBackupPlan(ctx, req)
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

func ExampleBackupForGKEClient_DeleteBackupPlan() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkebackup.NewBackupForGKEClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gkebackuppb.DeleteBackupPlanRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/gkebackup/v1#DeleteBackupPlanRequest.
	}
	op, err := c.DeleteBackupPlan(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleBackupForGKEClient_CreateBackup() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkebackup.NewBackupForGKEClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gkebackuppb.CreateBackupRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/gkebackup/v1#CreateBackupRequest.
	}
	op, err := c.CreateBackup(ctx, req)
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

func ExampleBackupForGKEClient_ListBackups() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkebackup.NewBackupForGKEClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gkebackuppb.ListBackupsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/gkebackup/v1#ListBackupsRequest.
	}
	it := c.ListBackups(ctx, req)
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

func ExampleBackupForGKEClient_GetBackup() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkebackup.NewBackupForGKEClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gkebackuppb.GetBackupRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/gkebackup/v1#GetBackupRequest.
	}
	resp, err := c.GetBackup(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleBackupForGKEClient_UpdateBackup() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkebackup.NewBackupForGKEClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gkebackuppb.UpdateBackupRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/gkebackup/v1#UpdateBackupRequest.
	}
	op, err := c.UpdateBackup(ctx, req)
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

func ExampleBackupForGKEClient_DeleteBackup() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkebackup.NewBackupForGKEClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gkebackuppb.DeleteBackupRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/gkebackup/v1#DeleteBackupRequest.
	}
	op, err := c.DeleteBackup(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleBackupForGKEClient_ListVolumeBackups() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkebackup.NewBackupForGKEClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gkebackuppb.ListVolumeBackupsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/gkebackup/v1#ListVolumeBackupsRequest.
	}
	it := c.ListVolumeBackups(ctx, req)
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

func ExampleBackupForGKEClient_GetVolumeBackup() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkebackup.NewBackupForGKEClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gkebackuppb.GetVolumeBackupRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/gkebackup/v1#GetVolumeBackupRequest.
	}
	resp, err := c.GetVolumeBackup(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleBackupForGKEClient_CreateRestorePlan() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkebackup.NewBackupForGKEClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gkebackuppb.CreateRestorePlanRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/gkebackup/v1#CreateRestorePlanRequest.
	}
	op, err := c.CreateRestorePlan(ctx, req)
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

func ExampleBackupForGKEClient_ListRestorePlans() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkebackup.NewBackupForGKEClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gkebackuppb.ListRestorePlansRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/gkebackup/v1#ListRestorePlansRequest.
	}
	it := c.ListRestorePlans(ctx, req)
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

func ExampleBackupForGKEClient_GetRestorePlan() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkebackup.NewBackupForGKEClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gkebackuppb.GetRestorePlanRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/gkebackup/v1#GetRestorePlanRequest.
	}
	resp, err := c.GetRestorePlan(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleBackupForGKEClient_UpdateRestorePlan() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkebackup.NewBackupForGKEClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gkebackuppb.UpdateRestorePlanRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/gkebackup/v1#UpdateRestorePlanRequest.
	}
	op, err := c.UpdateRestorePlan(ctx, req)
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

func ExampleBackupForGKEClient_DeleteRestorePlan() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkebackup.NewBackupForGKEClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gkebackuppb.DeleteRestorePlanRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/gkebackup/v1#DeleteRestorePlanRequest.
	}
	op, err := c.DeleteRestorePlan(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleBackupForGKEClient_CreateRestore() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkebackup.NewBackupForGKEClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gkebackuppb.CreateRestoreRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/gkebackup/v1#CreateRestoreRequest.
	}
	op, err := c.CreateRestore(ctx, req)
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

func ExampleBackupForGKEClient_ListRestores() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkebackup.NewBackupForGKEClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gkebackuppb.ListRestoresRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/gkebackup/v1#ListRestoresRequest.
	}
	it := c.ListRestores(ctx, req)
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

func ExampleBackupForGKEClient_GetRestore() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkebackup.NewBackupForGKEClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gkebackuppb.GetRestoreRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/gkebackup/v1#GetRestoreRequest.
	}
	resp, err := c.GetRestore(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleBackupForGKEClient_UpdateRestore() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkebackup.NewBackupForGKEClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gkebackuppb.UpdateRestoreRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/gkebackup/v1#UpdateRestoreRequest.
	}
	op, err := c.UpdateRestore(ctx, req)
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

func ExampleBackupForGKEClient_DeleteRestore() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkebackup.NewBackupForGKEClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gkebackuppb.DeleteRestoreRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/gkebackup/v1#DeleteRestoreRequest.
	}
	op, err := c.DeleteRestore(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleBackupForGKEClient_ListVolumeRestores() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkebackup.NewBackupForGKEClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gkebackuppb.ListVolumeRestoresRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/gkebackup/v1#ListVolumeRestoresRequest.
	}
	it := c.ListVolumeRestores(ctx, req)
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

func ExampleBackupForGKEClient_GetVolumeRestore() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkebackup.NewBackupForGKEClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gkebackuppb.GetVolumeRestoreRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/gkebackup/v1#GetVolumeRestoreRequest.
	}
	resp, err := c.GetVolumeRestore(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
