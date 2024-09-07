package bcl

import (
	"errors"
	"fmt"
	"strings"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"github.com/bufbuild/protovalidate-go"
	"github.com/pentops/bcl.go/bcl/errpos"
	"github.com/pentops/bcl.go/internal/ast"
	"github.com/pentops/bcl.go/internal/walker"
	"github.com/pentops/bcl.go/internal/walker/schema"
	"github.com/pentops/j5/gen/j5/bcl/v1/bcl_j5pb"
	"github.com/pentops/j5/lib/j5reflect"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Parser struct {
	refl     *j5reflect.Reflector
	Verbose  bool
	FailFast bool
	validate *protovalidate.Validator
	schema   *schema.SchemaSet
}

func NewParser(schemaSpec *bcl_j5pb.Schema) (*Parser, error) {
	pv, err := protovalidate.New()
	if err != nil {
		return nil, err
	}

	ss, err := schema.NewSchemaSet(schemaSpec)
	if err != nil {
		return nil, err
	}

	return &Parser{
		refl:     j5reflect.New(),
		FailFast: true,
		validate: pv,
		schema:   ss,
	}, nil
}

func (p *Parser) ParseFile(filename string, data string, msg protoreflect.Message) (*bcl_j5pb.SourceLocation, error) {

	obj, err := p.refl.NewObject(msg)
	if err != nil {
		return nil, err
	}

	tree, err := ast.ParseFile(data, p.FailFast)
	if err != nil {
		if err == ast.HadErrors {
			return nil, errpos.AddSourceFile(tree.Errors, filename, data)
		}
		return nil, fmt.Errorf("parse file not HadErrors - : %w", err)
	}

	source := &bcl_j5pb.SourceLocation{}
	scope, err := schema.NewRootSchemaWalker(p.schema, obj, source)
	if err != nil {
		return nil, err
	}

	err = walker.WalkSchema(scope, tree.Body, p.Verbose)
	if err != nil {
		err = errpos.AddSourceFile(err, filename, data)
		return source, fmt.Errorf("walkSchema: %w", err)
	}

	err = validateFile(p.validate, msg.Interface(), source)
	if err != nil {
		err = errpos.AddSourceFile(err, filename, data)
		return source, err
	}

	return source, nil
}

type baseSet struct {
	errors []*errpos.Err
}

type sourceSet struct {
	path []string
	loc  *bcl_j5pb.SourceLocation
	base *baseSet
}

func newSourceSet(locs *bcl_j5pb.SourceLocation) sourceSet {
	return sourceSet{
		loc:  locs,
		base: &baseSet{},
	}
}

func (s sourceSet) addViolation(violation *validate.Violation) {
	path := strings.Split(violation.FieldPath, ".")
	fullPath := make([]string, 0)
	for i, p := range path {
		parts := strings.Split(p, "[")
		if len(parts) > 1 {
			idx := parts[1]
			idx = strings.TrimSuffix(idx, "]")
			path[i] = parts[0]
			fullPath = append(fullPath, parts[0], idx)
		} else {
			fullPath = append(fullPath, p)
		}
	}

	ss := s
	for _, p := range fullPath {
		ss = ss.field(p)
	}
	ss.err(fmt.Errorf("%s: %s", violation.FieldPath, violation.Message))
}

func (s sourceSet) field(name string) sourceSet {
	childLoc := s.loc.Children[name]
	if childLoc == nil {
		childLoc = &bcl_j5pb.SourceLocation{
			StartLine:   s.loc.StartLine,
			StartColumn: s.loc.StartColumn,
		}
		if s.loc.Children == nil {
			s.loc.Children = make(map[string]*bcl_j5pb.SourceLocation)
		}
		s.loc.Children[name] = childLoc
	}
	if childLoc.StartLine == 0 {
		childLoc.StartLine = s.loc.StartLine
		childLoc.StartColumn = s.loc.StartColumn
	}

	child := sourceSet{
		path: append(s.path, name),
		base: s.base,
		loc:  childLoc,
	}
	return child
}

func (ss sourceSet) err(err error) {
	base, ok := errpos.AsError(err)
	if !ok {
		base = &errpos.Err{
			Err: err,
		}
	}

	if ss.loc != nil && base.Pos == nil {
		base.Pos = &errpos.Position{
			Start: errpos.Point{Line: int(ss.loc.StartLine), Column: int(ss.loc.StartColumn)},
			End:   errpos.Point{Line: int(ss.loc.EndLine), Column: int(ss.loc.EndColumn)},
		}
	}
	if len(base.Ctx) == 0 {
		base.Ctx = ss.path
	}
	ss.base.errors = append(ss.base.errors, base)
}

/*
func printSource(loc *sourcedef_j5pb.SourceLocation, prefix string) {
	fmt.Printf("%03d,%03d - %03d,%03d %s\n", loc.StartLine, loc.StartColumn, loc.EndLine, loc.EndColumn, prefix)
	for name, child := range loc.Children {
		printSource(child, prefix+"."+name)
	}
}
*/

func validateFile(pv *protovalidate.Validator, msg protoreflect.ProtoMessage, source *bcl_j5pb.SourceLocation) error {

	sources := newSourceSet(source)

	validationErr := pv.Validate(msg)
	if validationErr != nil {
		valErr := &protovalidate.ValidationError{}
		if errors.As(validationErr, &valErr) {
			//sources.err(valErr)
			for _, violation := range valErr.Violations {
				sources.addViolation(violation)
			}

		} else {
			return validationErr
		}
	}

	if len(sources.base.errors) == 0 {
		return nil
	}

	return errpos.Errors(sources.base.errors)
}