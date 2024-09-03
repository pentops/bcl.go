package j5convert

import (
	"fmt"
	"log"
	"strings"

	"github.com/pentops/j5/gen/j5/sourcedef/v1/sourcedef_j5pb"
)

type TypeNotFoundError struct {
	Package string
	Name    string
}

func (e *TypeNotFoundError) Error() string {
	return fmt.Sprintf("type %s not found in package %s", e.Name, e.Package)
}

type PackageNotFoundError struct {
	Package string
	Name    string
}

func (e *PackageNotFoundError) Error() string {
	return fmt.Sprintf("namespace %s not found (looking for %s.%s), missing import?", e.Package, e.Package, e.Name)
}

func (fb *Root) AddImports(spec ...*sourcedef_j5pb.Import) error {
	for _, imp := range spec {
		if strings.Contains(imp.Path, "/") {
			importPath := imp.Path
			if strings.HasSuffix(importPath, ".j5s") {
				importPath = importPath + ".proto"
			}
			fb.ensureImport(importPath)
			pkg := PackageFromFilename(imp.Path)
			fb.importAliases[pkg] = pkg
			continue
		}

		pkg := imp.Path
		if imp.Alias != "" {
			fb.importAliases[imp.Alias] = pkg
			continue
		}
		parts := strings.Split(pkg, ".")
		if len(parts) > 2 {
			return fmt.Errorf("AddImports: invalid package %q", pkg)
		}
		withoutVersion := parts[len(parts)-2]
		fb.importAliases[withoutVersion] = pkg
		fb.importAliases[pkg] = pkg
	}
	return nil
}

var implicitImports = map[string]*PackageSummary{
	"j5.state.v1": {
		Exports: map[string]*TypeRef{
			"StateMetadata": {
				Package:    "j5.state.v1",
				Name:       "StateMetadata",
				File:       "j5/state/v1/metadata.proto",
				MessageRef: &MessageRef{},
			},
			"EventMetadata": {
				Package:    "j5.state.v1",
				Name:       "EventMetadata",
				File:       "j5/state/v1/metadata.proto",
				MessageRef: &MessageRef{},
			},
		},
	},
}

func (fb *Root) resolveTypeNoImport(pkg string, name string) (*TypeRef, error) {
	thisPackage := fb.packageName
	if pkg == "" || pkg == thisPackage {
		pkg = thisPackage
		typeRef, err := fb.deps.ResolveType(pkg, name)
		if err != nil {
			return nil, fmt.Errorf("self import: %w", err)
		}

		return typeRef, nil
	}

	if implicit, ok := implicitImports[pkg]; ok {
		typeRef, ok := implicit.Exports[name]
		if ok {
			fb.ensureImport(typeRef.File)
			return typeRef, nil
		}
	}

	alias, ok := fb.importAliases[pkg]
	if !ok {
		log.Printf("resolveType: %q not found in %v", pkg, fb.importAliases)
		return nil, &PackageNotFoundError{
			Package: pkg,
			Name:    name,
		}
	}

	typeRef, err := fb.deps.ResolveType(alias, name)
	if err != nil {
		return nil, err
	}
	return typeRef, nil

}