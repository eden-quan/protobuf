// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package prototest_test

import (
	"fmt"
	"testing"

	"gitlab.lainuoniao.cn/rhinobird/backend/protobuf.git/internal/flags"
	"gitlab.lainuoniao.cn/rhinobird/backend/protobuf.git/proto"
	"gitlab.lainuoniao.cn/rhinobird/backend/protobuf.git/runtime/protoimpl"
	"gitlab.lainuoniao.cn/rhinobird/backend/protobuf.git/testing/prototest"

	irregularpb "gitlab.lainuoniao.cn/rhinobird/backend/protobuf.git/internal/testprotos/irregular"
	legacypb "gitlab.lainuoniao.cn/rhinobird/backend/protobuf.git/internal/testprotos/legacy"
	legacy1pb "gitlab.lainuoniao.cn/rhinobird/backend/protobuf.git/internal/testprotos/legacy/proto2_20160225_2fc053c5"
	testpb "gitlab.lainuoniao.cn/rhinobird/backend/protobuf.git/internal/testprotos/test"
	_ "gitlab.lainuoniao.cn/rhinobird/backend/protobuf.git/internal/testprotos/test/weak1"
	_ "gitlab.lainuoniao.cn/rhinobird/backend/protobuf.git/internal/testprotos/test/weak2"
	test3pb "gitlab.lainuoniao.cn/rhinobird/backend/protobuf.git/internal/testprotos/test3"
)

func Test(t *testing.T) {
	ms := []proto.Message{
		(*testpb.TestAllTypes)(nil),
		(*test3pb.TestAllTypes)(nil),
		(*testpb.TestRequired)(nil),
		(*irregularpb.Message)(nil),
		(*testpb.TestAllExtensions)(nil),
		(*legacypb.Legacy)(nil),
		protoimpl.X.MessageOf((*legacy1pb.Message)(nil)).Interface(),
	}
	if flags.ProtoLegacy {
		ms = append(ms, (*testpb.TestWeak)(nil))
	}

	for _, m := range ms {
		t.Run(fmt.Sprintf("%T", m), func(t *testing.T) {
			prototest.Message{}.Test(t, m.ProtoReflect().Type())
		})
	}
}
