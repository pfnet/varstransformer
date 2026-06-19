package main

import (
	"sigs.k8s.io/kustomize/api/types"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

func loadDefaultVarReference() types.FsSlice {
	var data struct {
		VarReference types.FsSlice `yaml:"varReference"`
	}
	if err := yaml.Unmarshal([]byte(defaultVarReference), &data); err != nil {
		panic(err)
	}
	return data.VarReference
}

// The following code is copied from https://github.com/kubernetes-sigs/kustomize/blob/ddcbae54ab1d8a6e869c2a1879632c375583633b/api/internal/konfig/builtinpluginconsts/varreference.go

// Copyright 2019 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

const defaultVarReference = `
varReference:
- path: spec/jobTemplate/spec/template/spec/containers/args
  kind: CronJob

- path: spec/jobTemplate/spec/template/spec/containers/command
  kind: CronJob

- path: spec/jobTemplate/spec/template/spec/containers/env/value
  kind: CronJob

- path: spec/jobTemplate/spec/template/spec/containers/volumeMounts/mountPath
  kind: CronJob

- path: spec/jobTemplate/spec/template/spec/initContainers/args
  kind: CronJob

- path: spec/jobTemplate/spec/template/spec/initContainers/command
  kind: CronJob

- path: spec/jobTemplate/spec/template/spec/initContainers/env/value
  kind: CronJob

- path: spec/jobTemplate/spec/template/spec/initContainers/volumeMounts/mountPath
  kind: CronJob

- path: spec/jobTemplate/spec/template/volumes/nfs/server
  kind: CronJob

- path: spec/template/spec/containers/args
  kind: DaemonSet

- path: spec/template/spec/containers/command
  kind: DaemonSet

- path: spec/template/spec/containers/env/value
  kind: DaemonSet

- path: spec/template/spec/containers/volumeMounts/mountPath
  kind: DaemonSet

- path: spec/template/spec/initContainers/args
  kind: DaemonSet

- path: spec/template/spec/initContainers/command
  kind: DaemonSet

- path: spec/template/spec/initContainers/env/value
  kind: DaemonSet

- path: spec/template/spec/initContainers/volumeMounts/mountPath
  kind: DaemonSet

- path: spec/template/spec/volumes/nfs/server
  kind: DaemonSet

- path: spec/template/spec/containers/args
  kind: Deployment

- path: spec/template/spec/containers/command
  kind: Deployment

- path: spec/template/spec/containers/env/value
  kind: Deployment

- path: spec/template/spec/containers/volumeMounts/mountPath
  kind: Deployment

- path: spec/template/spec/initContainers/args
  kind: Deployment

- path: spec/template/spec/initContainers/command
  kind: Deployment

- path: spec/template/spec/initContainers/env/value
  kind: Deployment

- path: spec/template/spec/initContainers/volumeMounts/mountPath
  kind: Deployment

- path: spec/template/spec/volumes/nfs/server
  kind: Deployment

- path: spec/template/metadata/annotations
  kind: Deployment

- path: spec/rules/host
  kind: Ingress

- path: spec/tls/hosts
  kind: Ingress
  
- path: spec/tls/secretName
  kind: Ingress

- path: spec/template/spec/containers/args
  kind: Job

- path: spec/template/spec/containers/command
  kind: Job

- path: spec/template/spec/containers/env/value
  kind: Job

- path: spec/template/spec/containers/volumeMounts/mountPath
  kind: Job

- path: spec/template/spec/initContainers/args
  kind: Job

- path: spec/template/spec/initContainers/command
  kind: Job

- path: spec/template/spec/initContainers/env/value
  kind: Job

- path: spec/template/spec/initContainers/volumeMounts/mountPath
  kind: Job

- path: spec/template/spec/volumes/nfs/server
  kind: Job

- path: spec/containers/args
  kind: Pod

- path: spec/containers/command
  kind: Pod

- path: spec/containers/env/value
  kind: Pod

- path: spec/containers/volumeMounts/mountPath
  kind: Pod

- path: spec/initContainers/args
  kind: Pod

- path: spec/initContainers/command
  kind: Pod

- path: spec/initContainers/env/value
  kind: Pod

- path: spec/initContainers/volumeMounts/mountPath
  kind: Pod

- path: spec/volumes/nfs/server
  kind: Pod

- path: spec/template/spec/containers/args
  kind: ReplicaSet

- path: spec/template/spec/containers/command
  kind: ReplicaSet

- path: spec/template/spec/containers/env/value
  kind: ReplicaSet

- path: spec/template/spec/containers/volumeMounts/mountPath
  kind: ReplicaSet

- path: spec/template/spec/initContainers/args
  kind: ReplicaSet

- path: spec/template/spec/initContainers/command
  kind: ReplicaSet

- path: spec/template/spec/initContainers/env/value
  kind: ReplicaSet

- path: spec/template/spec/initContainers/volumeMounts/mountPath
  kind: ReplicaSet

- path: spec/template/spec/volumes/nfs/server
  kind: ReplicaSet

- path: spec/ports/port
  kind: Service

- path: spec/ports/targetPort
  kind: Service

- path: spec/template/spec/containers/args
  kind: StatefulSet

- path: spec/template/spec/containers/command
  kind: StatefulSet

- path: spec/template/spec/containers/env/value
  kind: StatefulSet

- path: spec/template/spec/containers/volumeMounts/mountPath
  kind: StatefulSet

- path: spec/template/spec/initContainers/args
  kind: StatefulSet

- path: spec/template/spec/initContainers/command
  kind: StatefulSet

- path: spec/template/spec/initContainers/env/value
  kind: StatefulSet

- path: spec/template/spec/initContainers/volumeMounts/mountPath
  kind: StatefulSet

- path: spec/volumeClaimTemplates/spec/nfs/server
  kind: StatefulSet

- path: spec/nfs/server
  kind: PersistentVolume

- path: metadata/labels

- path: metadata/annotations
`
