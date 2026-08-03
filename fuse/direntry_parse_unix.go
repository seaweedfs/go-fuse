//go:build !windows

// Copyright 2016 the Go-FUSE Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fuse

import (
	"bytes"
	"unsafe"
)

// Parse reads an entry from getdents(2) buffer. It returns the number
// of bytes consumed.
func (d *DirEntry) Parse(buf []byte) int {
	// We can't use syscall.Dirent here, because it declares a
	// [256]byte name, which may run beyond the end of ds.todo.
	// when that happens in the race detector, it causes a panic
	// "converted pointer straddles multiple allocations"
	de := (*dirent)(unsafe.Pointer(&buf[0]))
	off := unsafe.Offsetof(dirent{}.Name)
	nameBytes := buf[off : off+uintptr(de.nameLength())]
	n := de.Reclen

	l := bytes.IndexByte(nameBytes, 0)
	if l >= 0 {
		nameBytes = nameBytes[:l]
	}
	*d = DirEntry{
		Ino:  de.Ino,
		Mode: (uint32(de.Type) << 12),
		Name: string(nameBytes),
		Off:  uint64(de.Off),
	}
	return int(n)
}
