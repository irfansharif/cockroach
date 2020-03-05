// Copyright 2020 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

// XXX: sqlmigrations + leasemanager for API, but why we need leases?
// XXX: Skipping jobs for now per Andrew's advice.
// XXX: Look at client_*_test.go files for client.DB usage/test ideas.
// XXX: How do cluster migrations start machinery today, on change?
// ANS: They just poll IsActive, and behave accordingly. No "migrations" happen.
// We want something like VersionUpgradeHook to actually set up the thing, run
// hook after every node will observe IsActive == true.
// XXX: Think about "multi-tenant" migrations.
// XXX: Need to define "node upgrade" RPC subsystem.
// XXX: Break into (a) KV migrations, and (b) SQL migrations. Doing (a) to start
// off.
package migration

import (
	"context"

	"github.com/cockroachdb/cockroach/pkg/internal/client"
	"github.com/cockroachdb/cockroach/pkg/roachpb"
)

type Migration func(*Helper) error

type Helper struct {
	db *client.DB
	// XXX: Lots more stuff. Including facility to run SQL migrations
	// (eventually).
}

func (h *Helper) IterRangeDescriptors(
	f func(...roachpb.RangeDescriptor) error, blockSize int,
) error {
	// Paginate through all ranges and invoke f with chunks of size ~blockSize.
	// call h.Progress between invocations.
	return nil
}

func (h *Helper) Retry(f func() error) error {
	for {
		err := f()
		if err != nil {
			continue
		}
		return err
	}
}

func (h *Helper) RequiredNodes() []roachpb.NodeID {
	// TODO: read node registry and return list of nodes to be taken into
	// account.
	_ = h.db
	return nil
}

func (h *Helper) Progress(s string, num, denum int) {
	// Set progress of the current step to num/denum. Denum can be zero if final
	// steps not known (example: iterating through ranges)
	// Free-form message can be attached (PII?)
	//
	// XXX: this API is crap
}

func (h *Helper) EveryNode(op string, args ...interface{}) error {
	for _, nodeID := range h.RequiredNodes() {
		// dial nodeID and send with proper args (in parallel):
		_ = nodeID
		_ = EveryNodeRequest{action: op}
		// XXX: These are per node requests. Not using client.DB interface.

		// Also return a hard error here if a node that is required is down.
		// There's no point trying to upgrade until the node is back and by
		// handing control back to the orchestrator we get better UX.
		//
		// For any transient-looking errors, retry. For any hard-looking errors,
		// return to orchestrator (which will likely want to back off; migrating
		// in a hot loop could cause stability issues).
	}
	return nil
}

type V struct {
	roachpb.Version
	Migration // XXX: The hook, to be run post-gossip of cluster setting.
}

// Stub for proper proto-backed request type.
type EveryNodeRequest struct {
	action string // very much not real code
}

func RunMigrations() error {
	// This gets called by something higher up that can own the job, and can set
	// it up, etc.
	// XXX: Where do we capture the lease?

	ctx := context.Background()
	var db *client.DB // XXX: From caller
	ClusterVersionKey := roachpb.Key("XXX:")
	var ackedV roachpb.Version
	if err := db.GetProto(ctx, ClusterVersionKey, &ackedV); err != nil {
		return err
	}

	// Hard-coded example data. From 20.1, will run the following migrations, not
	// all of which may have a Migration attached. Assembled from a registry while
	// taking into account ackedV
	vs := []V{
		{
			Version:   roachpb.Version{Major: 20, Minor: 1},
			Migration: func(*Helper) error { return nil },
		},
	}

	for _, v := range vs {
		h := &Helper{}
		// h.Progress should basically call j.Status(v.String()+": "+s, num, denum)

		// Persist the beginning of this migration on all nodes. Basically they
		// will persist the version, then move the cluster setting forward, then
		// return.
		//
		// Should they store whether migration is ongoing? No.
		// - Argument for yes: We might have functionality that is only safe
		// when all nodes in the cluster have "done their part" (XXX: What does
		// this mean) to migrate into it, and which we don't want to delay for
		// one release.
		// - Argument for no: We could just gate that functionality's activation on a
		// later unstable version, and get the right behavior.
		if err := h.EveryNode("ack-pending-version", v); err != nil {
			return err
		}
		if err := v.Migration(h); err != nil {
			return err
		}
	}
	return nil
}
