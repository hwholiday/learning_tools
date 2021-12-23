package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/hwholiday/learning_tools/grpc/protoc-gen-hw/generator"
	"github.com/hwholiday/learning_tools/grpc/protoc-gen-hw/plugin/hw"
	"io/ioutil"
	"os"
)

func main() {
	generator.RegisterPlugin(new(hw.Hw))
	g := generator.New()
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		g.Error(err, "reading input")
	}
	if err := proto.Unmarshal(data, g.Request); err != nil {
		g.Error(err, "parsing input proto")
	}
	if len(g.Request.FileToGenerate) == 0 {
		g.Fail("no files to generate")
	}
	g.CommandLineParameters(g.Request.GetParameter())
	g.WrapTypes()
	g.SetPackageNames()
	g.BuildTypeNameMap()
	g.GenerateAllFiles()
	data, err = proto.Marshal(g.Response)
	if err != nil {
		g.Error(err, "failed to marshal output proto")
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		g.Error(err, "failed to write output proto")
	}
}
