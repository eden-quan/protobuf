// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The protoc-gen-go-fx binary is a protoc plugin to generate Go code for
// both proto2 and proto3 versions of the protocol buffer language.
//
// For more information about the usage of this plugin, see:
// https://protobuf.dev/reference/go/go-generated.
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	gengo "gitlab.lainuoniao.cn/rhinobird/backend/protobuf.git/cmd/protoc-gen-go-fx/internal_gengo"
	"gitlab.lainuoniao.cn/rhinobird/backend/protobuf.git/compiler/protogen"
	"gitlab.lainuoniao.cn/rhinobird/backend/protobuf.git/internal/version"
)

const genGoDocURL = "https://protobuf.dev/reference/go/go-generated"
const grpcDocURL = "https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code"

func main() {
	if len(os.Args) == 2 && os.Args[1] == "--version" {
		fmt.Fprintf(os.Stdout, "%v %v\n", filepath.Base(os.Args[0]), version.String())
		os.Exit(0)
	}
	if len(os.Args) == 2 && os.Args[1] == "--help" {
		fmt.Fprintf(os.Stdout, "See "+genGoDocURL+" for usage information.\n")
		os.Exit(0)
	}

	//in, _ := io.ReadAll(os.Stdin)
	//
	//f, _ := os.OpenFile("./test.data", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0766)
	//_, _ = f.Write(in)
	//_ = f.Close()
	//return

	//// load test file
	//tempFile, err := os.OpenFile("./testdata/test.data", os.O_RDONLY, 0766)
	//if err != nil {
	//	fmt.Println("load test file failed:", err)
	//	return
	//}
	//defer func() { _ = tempFile.Close() }()
	//
	////replace test file as stdin
	//os.Stdin = tempFile

	var (
		flags   flag.FlagSet
		plugins = flags.String("plugins", "", "deprecated option")
	)
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		if *plugins != "" {
			return errors.New("protoc-gen-go-fx: plugins are not supported; use 'protoc --go-grpc_out=...' to generate gRPC\n\n" +
				"See " + grpcDocURL + " for more information.")
		}
		for _, f := range gen.Files {
			if f.Generate {
				gengo.GenerateFile(gen, f)
			}
		}
		gen.SupportedFeatures = gengo.SupportedFeatures
		return nil
	})
}
