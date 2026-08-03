//go:build !windows

// Copyright 2016 the Go-FUSE Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fuse

import "syscall"

const o_DIRECTORY = syscall.O_DIRECTORY
