// Copyright 2016 the Go-FUSE Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fuse

import "fmt"

// Server exists on Windows only so that RawFileSystem implementations, whose
// Init takes one, compile. There is no FUSE kernel transport here: mount
// through a Windows filesystem driver and drive the RawFileSystem directly.
type Server struct{}

func NewServer(fs RawFileSystem, mountPoint string, opts *MountOptions) (*Server, error) {
	return nil, fmt.Errorf("fuse: no kernel transport on windows")
}

func (ms *Server) Serve() {}

func (ms *Server) Wait() {}

func (ms *Server) WaitMount() error { return fmt.Errorf("fuse: no kernel transport on windows") }

func (ms *Server) Unmount() error { return fmt.Errorf("fuse: no kernel transport on windows") }

func (ms *Server) SetDebug(dbg bool) {}

// The kernel holds no cache to invalidate, so the notify calls are no-ops.

func (ms *Server) InodeNotify(node uint64, off int64, length int64) Status { return ENOSYS }

func (ms *Server) InodeNotifyStoreCache(node uint64, offset int64, data []byte) Status {
	return ENOSYS
}

func (ms *Server) InodeRetrieveCache(node uint64, offset int64, dest []byte) (int, Status) {
	return 0, ENOSYS
}

func (ms *Server) EntryNotify(parent uint64, name string) Status { return ENOSYS }

func (ms *Server) DeleteNotify(parent uint64, child uint64, name string) Status { return ENOSYS }
