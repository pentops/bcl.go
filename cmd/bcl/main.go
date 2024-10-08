package main

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"

	"github.com/pentops/bcl.go/bcl"
	"github.com/pentops/bcl.go/bcl/bclsp"
	"github.com/pentops/bcl.go/bcl/errpos"
	"github.com/pentops/bcl.go/gen/j5/bcl/v1/bcl_j5pb"
	"github.com/pentops/bcl.go/internal/parser"
	"github.com/pentops/runner/commander"
	"google.golang.org/protobuf/encoding/prototext"
)

var Version = "dev"

func main() {
	cmdGroup := commander.NewCommandSet()
	cmdGroup.Add("lint", commander.NewCommand(runLint))
	cmdGroup.Add("fmt", commander.NewCommand(runFmt))
	cmdGroup.Add("lsp", commander.NewCommand(runLSP))
	cmdGroup.RunMain("bcl", Version)
}

type RootConfig struct {
	ProjectRoot string `flag:"project-root" default:"" desc:"Project root directory"`
	Verbose     bool   `flag:"verbose" env:"BCL_VERBOSE" default:"false" desc:"Verbose output"`
}

func runLint(ctx context.Context, cfg struct {
	RootConfig
	Filename string `flag:"filename" desc:"Filename to lint"`
}) error {

	schemaSpec := &bcl_j5pb.Schema{
		Blocks: []*bcl_j5pb.Block{{
			SchemaName: "j5.bcl.v1.Block",
			Name: &bcl_j5pb.Tag{
				FieldName: "schemaName",
			},
		}, {
			SchemaName: "j5.bcl.v1.ScalarSplit",
			Alias: []*bcl_j5pb.Alias{{
				Name: "required",
				Path: &bcl_j5pb.Path{Path: []string{"requiredFields", "path"}},
			}, {
				Name: "optional",
				Path: &bcl_j5pb.Path{Path: []string{"optionalFields", "path"}},
			}, {
				Name: "remainder",
				Path: &bcl_j5pb.Path{Path: []string{"remainderField", "path"}},
			}},
		}, {
			SchemaName: "j5.bcl.v1.Tag",
			Alias: []*bcl_j5pb.Alias{{
				Name: "path",
				Path: &bcl_j5pb.Path{
					Path: []string{"path", "path"},
				},
			}},
		}},
	}

	msg := &bcl_j5pb.SchemaFile{}
	/*
		//sc := j5schema.NewSchemaCache()
			rootSchema, err := sc.Schema(msg.ProtoReflect().Descriptor())
			if err != nil {
				return err
			}*/

	parser, err := bcl.NewParser(schemaSpec) //rootSchema.(*j5schema.ObjectSchema))
	if err != nil {
		return err
	}
	parser.Verbose = cfg.Verbose

	content, err := os.ReadFile(cfg.Filename)
	if err != nil {
		return err
	}

	_, mainError := parser.ParseFile(cfg.Filename, string(content), msg.ProtoReflect())
	if mainError == nil {
		fmt.Println(prototext.Format(msg))
		return nil
	}

	locErr, ok := errpos.AsErrorsWithSource(mainError)
	if !ok {
		return mainError
	}

	log.Println(locErr.HumanString(2))

	os.Exit(100)
	return nil
}

func runFmt(ctx context.Context, cfg struct {
	Dir   string `flag:"dir" default:"." desc:"Root schema directory, or single file"`
	Write bool   `flag:"write" default:"false" desc:"Write fixes to files"`
}) error {

	doFile := func(data []byte) (string, error) {
		fixed, err := parser.Fmt(string(data))
		if err != nil {
			return "", err
		}
		return fixed, nil
	}

	stat, err := os.Lstat(cfg.Dir)
	if err != nil {
		return err
	}
	if !stat.IsDir() {
		data, err := os.ReadFile(cfg.Dir)
		if err != nil {
			return err
		}
		out, err := doFile(data)
		if err != nil {
			return err
		}
		if !cfg.Write {
			fmt.Printf("Fixed: %s\n", cfg.Dir)
			fmt.Println(out)
		} else {
			return os.WriteFile(cfg.Dir, []byte(out), 0644)
		}
		return nil
	}

	outWriter := &fileWriter{dir: cfg.Dir}
	root := os.DirFS(cfg.Dir)
	err = fs.WalkDir(root, ".", func(pathname string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		if path.Ext(pathname) != ".j5s" {
			return nil
		}

		data, err := fs.ReadFile(root, pathname)
		if err != nil {
			return err
		}

		out, err := doFile(data)
		if err != nil {
			return err
		}
		if !cfg.Write {
			fmt.Printf("Fixed: %s\n", pathname)
			fmt.Println(out)
			return nil
		} else {
			return outWriter.PutFile(ctx, pathname, []byte(out))
		}
	})
	if err != nil {
		return err
	}
	return nil
}

type fileWriter struct {
	dir string
}

func (f *fileWriter) PutFile(ctx context.Context, filename string, data []byte) error {
	dir := path.Join(f.dir, path.Dir(filename))
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return os.WriteFile(path.Join(f.dir, filename), data, 0644)
}

func runLSP(ctx context.Context, cfg struct {
	Dir string `flag:"project-root" default:"" desc:"Root schema directory"`
}) error {

	if cfg.Dir == "" {
		pwd, err := os.Getwd()
		if err != nil {
			return err
		}
		cfg.Dir = pwd
	}

	log.Printf("ARGS: %+v", os.Args)

	return bclsp.RunLSP(ctx, bclsp.Config{
		ProjectRoot: cfg.Dir,
		Schema:      nil,
		FileFactory: nil,
	})
}
