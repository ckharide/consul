package main

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"os"
)

func generate(cfgs []structConfig, targets map[string]targetPkg) error {
	byOutput := configsByOutput(cfgs)
	for _, group := range byOutput {

		first := group[0]
		file := newASTFile(first)

		for _, cfg := range group {
			t := targets[cfg.Target.Package].Structs[cfg.Target.Struct]
			if t.Name == "" {
				return fmt.Errorf("unable to locate target %v for %v", cfg.Target, cfg.Source)
			}

			node, err := generateConversion(cfg, t)
			if err != nil {
				return fmt.Errorf("failed to generate conversion for %v: %w", cfg.Source, err)
			}
			_ = node

			// TODO: add node to file

			// TODO: generate tests
		}

		if err := writeFile(first.Output, file); err != nil {
			return fmt.Errorf("failed to write generated code to %v: %w", first.Output, err)
		}

	}
	return nil
}

func newASTFile(first structConfig) *ast.File {
	return nil
}

func configsByOutput(cfgs []structConfig) [][]structConfig {
	return nil
}

func generateConversion(cfg structConfig, t targetStruct) (ast.Node, error) {
	file := &ast.File{
		// TODO: Name:       &ast.Ident{Name: }, // need source package name
		// TODO: Imports:    nil, // what imports are needed?
	}
	// TODO:
	return file, nil
}

// TODO: write build tags
// TODO: write file header
func writeFile(output string, file *ast.File) error {
	fh, err := os.Create(output)
	if err != nil {
		return err
	}
	return format.Node(fh, new(token.FileSet), file)
}
