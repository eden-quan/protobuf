// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jsonfuzz

import (
	"testing"

	"gitlab.lainuoniao.cn/rhinobird/backend/protobuf.git/internal/fuzztest"
)

func Test(t *testing.T) {
	fuzztest.Test(t, Fuzz)
}
