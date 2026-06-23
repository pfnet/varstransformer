package main

import (
	"fmt"

	"sigs.k8s.io/kustomize/api/filters/refvar"
	"sigs.k8s.io/kustomize/api/types"
	"sigs.k8s.io/kustomize/kyaml/fn/framework"
	"sigs.k8s.io/kustomize/kyaml/kio"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

type VarsTransformer struct {
	Vars          []types.Var   `json:"vars"`
	VarReferences types.FsSlice `json:"varReference"`
}

func buildProcessor() framework.ResourceListProcessor {
	var config VarsTransformer

	fn := func(items []*yaml.RNode) ([]*yaml.RNode, error) {
		merged, err := loadDefaultVarReference().MergeAll(config.VarReferences)
		if err != nil {
			return nil, fmt.Errorf("merging varReference: %w", err)
		}
		config.VarReferences = merged

		vars := make(map[string]interface{})
		for _, v := range config.Vars {
			v.Defaulting()

			item, err := FindObject(items, v.ObjRef)
			if err != nil {
				return nil, err
			}
			if item == nil {
				return nil, fmt.Errorf("object not found: %s", formatTargetRef(v.ObjRef))
			}

			value, err := item.GetFieldValue(v.FieldRef.FieldPath)
			if err != nil {
				return nil, fmt.Errorf("get field value (%q): %w", v.FieldRef.FieldPath, err)
			}
			vars[v.Name] = value
		}

		counts := map[string]int{}
		mf := refvar.MakePrimitiveReplacer(counts, vars)

		for _, fs := range config.VarReferences {
			f := refvar.Filter{
				MappingFunc: mf,
				FieldSpec:   fs,
			}
			if _, err := f.Filter(items); err != nil {
				return nil, fmt.Errorf("replacing %q: %w", fs.String(), err)
			}
		}
		return items, nil
	}

	return framework.SimpleProcessor{Config: &config, Filter: kio.FilterFunc(fn)}
}
