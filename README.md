# cmt-controller

This repository implements a simple controller for watching ConfigMapTemplate
resources as defined with a CustomResourceDefinition (CRD).

It is largely derived from the sample at https://github.com/kubernetes/sample-controller

It makes use of the generators in [k8s.io/code-generator](https://github.com/kubernetes/code-generator)
to generate a typed client, informers, listers and deep-copy functions. You can
do this yourself using the `./hack/update-codegen.sh` script.

The `update-codegen` script will automatically generate the following files &
directories:

* `pkg/apis/cmt/v1alpha1/zz_generated.deepcopy.go`
* `pkg/client/`

Changes should not be made to these files manually, and when modifying this implementation you should run the `update-codegen` script to generate new versions.

## Running

**Prerequisite**: Since the sample-controller uses `apps/v1` deployments, the Kubernetes cluster version should be greater than 1.9.

```sh
# create a CustomResourceDefinition, ClusterRole, ServiceAccount, and
# Deployment that set up in Kubernetes the customresource and the
# controller that makes it work
$ kubectl create -f artifacts/configmaptemplate.yaml

# create a custom resource of type ConfigMapTemplate
$ kubectl create -f artifacts/examples/example-configmap-template.yaml

# check configmaps created through the custom resource
$ kubectl get configmaps
```

## Cleanup

You can clean up the created CustomResourceDefinition with:

    $ kubectl delete -f artifacts/configmaptemplate.yaml
