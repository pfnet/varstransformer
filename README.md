# VarsTransformer

VarsTransformer is a Kustomize plugin implemented as a KRM function, providing `vars`-style substitution without relying on Kustomize's deprecated built-in `vars` feature.

This project keeps the familiar `vars` and `varReference` configuration format, so existing users of Kustomize `vars` can migrate from the deprecated built-in feature easily.

## How to use

Create a function config like the following:

```yaml
apiVersion: preferred.jp/v1alpha1
kind: VarsTransformer
metadata:
  name: vars
  annotations:
    config.kubernetes.io/function: |
      container:
        image: ghcr.io/pfnet/varstransformer:latest
vars:
  - name: SERVICE_NAME
    objref:
      apiVersion: v1
      kind: Service
      name: my-app
    fieldref:
      fieldPath: metadata.name
```

Then reference it from `transformers` in your `kustomization.yaml`:

```yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

resources:
  - service.yaml
  - deployment.yaml

transformers:
  - vars-transformer.yaml
```

When running `kustomize build`, `--enable-alpha-plugins` is required to execute plugin:

```bash
kustomize build --enable-alpha-plugins .
```

## Differences from Kustomize built-in `vars`

VarsTransformer uses the same `vars` and `varReference` format as Kustomize built-in `vars`, but it behaves like a normal transformer.
In other words, VarsTransformer only sees the resources available where it runs, and it resolves references and performs substitution at that point in the build.

Kustomize built-in `vars` behave differently, so some cases do not work in exactly the same way.

### Scope is different

Because VarsTransformer behaves like a normal transformer, its effect is limited to the resources visible where it runs.

Built-in `vars`, by contrast, are collected across the full accumulated build. As a result, a build that combines several related overlays or variants which all reuse the same base can run into conflicts when the same var is defined more than once.

VarsTransformer does not provide that kind of build-wide behavior.

### Execution order is different

VarsTransformer resolves references and performs substitution when the transformer itself runs.

Built-in `vars`, by contrast, are handled later in Kustomize's build process as a special step.

This also matters with overlays. If VarsTransformer runs in a base, and a later overlay adds a new `$(VAR)` reference, that new reference will not be replaced. In such cases, VarsTransformer should run at the overlay level or after the relevant patches and composition steps.

VarsTransformer can match `objref` against both the current resource name and the previous name recorded during Kustomize transformations. This allows references written with the original name to keep working after transforms such as `namePrefix` or `nameSuffix`.
This fallback only works when Kustomize has recorded the previous identity correctly. If multiple resources share the same previous identity, VarsTransformer reports the reference as ambiguous. Substitution still happens only at the point where VarsTransformer runs, so later overlays cannot introduce new `$(VAR)` references and expect them to be resolved automatically.
