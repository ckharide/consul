package main

import (
	"fmt"
	"go/ast"
)

type structConfig struct {
	Target       string
	Output       string
	Name         string
	IgnoreFields []string
	FuncFrom     string
	FuncTo       string
	Fields       []fieldConfig
}

type fieldConfig struct {
	Name string
	// TODO: Pointer pointerSettings
	FuncFrom string
	FuncTo   string
}

func configsFromAnnotations(sources pkg) ([]structConfig, error) {
	names := sources.Names()
	cfgs := make([]structConfig, 0, len(names))
	for _, name := range names {
		strct := sources.structs[name]
		cfg, err := parseStructAnnotation(strct)
		if err != nil {
			return nil, fmt.Errorf("from source %v: %w", name, err)
		}

		for _, field := range strct.Fields {
			f, err := parseFieldAnnotation(field)
			if err != nil {
				return nil, fmt.Errorf("from source %v.%v: %w", name, fieldNameFromAST(field.Names), err)
			}
			cfg.Fields = append(cfg.Fields, f)
		}

		cfgs = append(cfgs, cfg)
	}
	return cfgs, nil
}

func fieldNameFromAST(names []*ast.Ident) string {
	if len(names) == 0 {
		return "unknown"
	}
	return names[0].Name
}

func parseStructAnnotation(decl structDecl) (structConfig, error) {
	var c structConfig

	return c, nil
}

func parseFieldAnnotation(field *ast.Field) (fieldConfig, error) {
	var c fieldConfig
	return c, nil
}
