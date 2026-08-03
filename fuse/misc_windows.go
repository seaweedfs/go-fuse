// Copyright 2016 the Go-FUSE Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fuse

// Value borrowed from Linux. Windows never runs the request path that
// consults it; it exists so the shared code compiles.
const _UTIME_OMIT = (1 << 30) - 2
