package main

import (
	"fmt"
	"sort"
	"strings"

	"sigs.k8s.io/kustomize/api/types"
	"sigs.k8s.io/kustomize/kyaml/resid"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

const (
	buildAnnotationPreviousKinds      = "internal.config.kubernetes.io/previousKinds"
	buildAnnotationPreviousNames      = "internal.config.kubernetes.io/previousNames"
	buildAnnotationPreviousNamespaces = "internal.config.kubernetes.io/previousNamespaces"
)

func matchesTarget(id resid.ResId, target resid.ResId, namespaceSpecified bool) bool {
	if namespaceSpecified {
		return id.Equals(target)
	}
	return id.Gvk.Equals(target.Gvk) && id.Name == target.Name
}

func previousIds(item *yaml.RNode) ([]resid.ResId, error) {
	annotations := item.GetAnnotations(
		buildAnnotationPreviousNames,
		buildAnnotationPreviousNamespaces,
		buildAnnotationPreviousKinds,
	)
	if _, ok := annotations[buildAnnotationPreviousNames]; !ok {
		return nil, nil
	}

	names := strings.Split(annotations[buildAnnotationPreviousNames], ",")
	namespaces := strings.Split(annotations[buildAnnotationPreviousNamespaces], ",")
	kinds := strings.Split(annotations[buildAnnotationPreviousKinds], ",")
	if len(names) != len(namespaces) || len(names) != len(kinds) {
		return nil, fmt.Errorf(
			"number of previous names, number of previous namespaces, number of previous kinds not equal",
		)
	}

	group, version := resid.ParseGroupVersion(item.GetApiVersion())
	ids := make([]resid.ResId, 0, len(names))
	for i := range names {
		ids = append(ids, resid.NewResIdWithNamespace(
			resid.Gvk{
				Group:   group,
				Version: version,
				Kind:    kinds[i],
			},
			names[i],
			namespaces[i],
		))
	}
	return ids, nil
}

func allIds(item *yaml.RNode) ([]resid.ResId, error) {
	ids := []resid.ResId{resid.FromRNode(item)}
	prevIDs, err := previousIds(item)
	if err != nil {
		return nil, err
	}
	return append(ids, prevIDs...), nil
}

func FindObject(items []*yaml.RNode, t types.Target) (*yaml.RNode, error) {
	targetID := resid.NewResIdWithNamespace(
		t.GVK(),
		t.Name,
		t.Namespace,
	)
	namespaceSpecified := t.Namespace != ""
	var matches []*yaml.RNode
	for _, item := range items {
		ids, err := allIds(item)
		if err != nil {
			return nil, fmt.Errorf("build resource ids for %s: %w", formatResourceRef(resid.FromRNode(item)), err)
		}
		for _, id := range ids {
			if matchesTarget(id, targetID, namespaceSpecified) {
				matches = append(matches, item)
				break
			}
		}
	}
	if len(matches) == 0 {
		return nil, nil
	}
	if len(matches) > 1 {
		return nil, fmt.Errorf(
			"ambiguous object reference: %s, matched=%s",
			formatTargetRef(t),
			formatMatchedResourceRefs(matches),
		)
	}
	return matches[0], nil
}

func formatResourceRef(id resid.ResId) string {
	if id.Namespace == "" {
		return fmt.Sprintf("%s", id.Name)
	}
	return fmt.Sprintf("%s/%s", id.Namespace, id.Name)
}

func formatTargetRef(t types.Target) string {
	parts := []string{fmt.Sprintf("gvk=%s", t.GVK().String())}
	if t.Namespace != "" {
		parts = append(parts, fmt.Sprintf("namespace=%s", t.Namespace))
	}
	parts = append(parts, fmt.Sprintf("name=%s", t.Name))
	return strings.Join(parts, ", ")
}

func formatMatchedResourceRefs(matches []*yaml.RNode) string {
	resources := make([]string, 0, len(matches))
	for _, match := range matches {
		resources = append(resources, formatResourceRef(resid.FromRNode(match)))
	}
	sort.Strings(resources)
	return strings.Join(resources, ", ")
}
