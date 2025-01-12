// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin
// +build darwin

package fsnotifier

import "golang.org/x/sys/unix"

// note: this constant is not defined on BSD
const openMode = unix.O_EVTONLY | unix.O_CLOEXEC
