// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// init() is copied from filepath/path_windows_test.go

package globule_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func init() {
	tmpdir, err := ioutil.TempDir("", "symtest")
	if err != nil {
		panic("failed to create temp directory: " + err.Error())
	}
	defer os.RemoveAll(tmpdir)

	err = os.Symlink("target", filepath.Join(tmpdir, "symlink"))
	if err == nil {
		return
	}

	err = err.(*os.LinkError).Err
	switch err {
	case syscall.EWINDOWS, syscall.ERROR_PRIVILEGE_NOT_HELD:
		supportsSymlinks = false
	}
}
