package protobuild

import (
	"bytes"
	"fmt"
	"strings"

	proto_parser "github.com/bufbuild/protocompile/parser"
	"github.com/bufbuild/protocompile/reporter"
	"github.com/pentops/bcl.go/internal/j5convert"
	"google.golang.org/protobuf/types/descriptorpb"
)

type ProtoParser struct {
	handler *reporter.Handler
}

func NewProtoParser(rep reporter.Reporter) *ProtoParser {
	reportHandler := reporter.NewHandler(rep)
	return &ProtoParser{handler: reportHandler}
}

func (pp ProtoParser) protoToDescriptor(filename string, data []byte) (*descriptorpb.FileDescriptorProto, error) {
	fileNode, err := proto_parser.Parse(filename, bytes.NewReader(data), pp.handler)
	if err != nil {
		return nil, err
	}
	result, err := proto_parser.ResultFromAST(fileNode, true, pp.handler)
	if err != nil {
		return nil, err
	}
	return result.FileDescriptorProto(), nil
}

func (pp ProtoParser) buildSummary(filename string, data []byte) (*j5convert.FileSummary, error) {

	res, err := pp.protoToDescriptor(filename, data)
	if err != nil {
		return nil, err
	}

	return pp.buildSummaryFromDescriptor(res)
}

func (pp ProtoParser) buildSummaryFromDescriptor(res *descriptorpb.FileDescriptorProto) (*j5convert.FileSummary, error) {
	filename := res.GetName()
	exports := map[string]*j5convert.TypeRef{}

	for _, msg := range res.MessageType {
		exports[msg.GetName()] = &j5convert.TypeRef{
			Name:       msg.GetName(),
			File:       filename,
			MessageRef: &j5convert.MessageRef{},
		}
	}
	for _, en := range res.EnumType {
		built, err := buildEnumRef(en)
		if err != nil {
			return nil, err
		}
		exports[en.GetName()] = &j5convert.TypeRef{
			Name:    en.GetName(),
			File:    filename,
			EnumRef: built,
		}
	}

	return &j5convert.FileSummary{
		Exports:          exports,
		FileDependencies: res.Dependency,
		// No type dependencies for proto files, all deps come from the files.
	}, nil
}

func buildEnumRef(enumDescriptor *descriptorpb.EnumDescriptorProto) (*j5convert.EnumRef, error) {
	for idx, value := range enumDescriptor.Value {
		if value.Number == nil {
			return nil, fmt.Errorf("enum value[%d] does not have a number", idx)
		}
		if value.Name == nil {
			return nil, fmt.Errorf("enum value[%d] does not have a name", idx)
		}
	}

	if len(enumDescriptor.Value) < 1 {
		return nil, fmt.Errorf("enum has no values")
	}
	if *enumDescriptor.Value[0].Number != 0 {
		return nil, fmt.Errorf("enum does not have a value 0")
	}

	suffix := "UNSPECIFIED"
	unspecifiedVal := *enumDescriptor.Value[0].Name
	if !strings.HasSuffix(unspecifiedVal, suffix) {
		return nil, fmt.Errorf("enum value 0 should have suffix %s", suffix)
	}
	trimPrefix := strings.TrimSuffix(unspecifiedVal, suffix)
	ref := &j5convert.EnumRef{
		Prefix: trimPrefix,
		ValMap: map[string]int32{},
	}

	for _, value := range enumDescriptor.Value {
		ref.ValMap[value.GetName()] = value.GetNumber()
	}
	return ref, nil
}