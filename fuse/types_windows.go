// Copyright 2016 the Go-FUSE Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fuse

import "syscall"

const (
	ENODATA = Status(syscall.ENODATA)
	// Windows has no ENOATTR; alias it as Linux does.
	ENOATTR = ENODATA
)
