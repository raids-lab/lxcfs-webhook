# LXCFS Admission Webhook Documentation

## 1. Introduction

The LXCFS Admission Webhook is a crucial component designed to streamline and automate the integration of LXCFS (Linux Containers Filesystem) within Kubernetes environments. 

LXCFS provides a set of virtual filesystems that expose container-specific information, such as CPU and memory usage, in a format that is consistent with the Linux kernel's interfaces. 

This webhook acts as an admission controller, intercepting Kubernetes API requests related to namespaces and pods, and applying necessary configurations to enable LXCFS for the relevant resources.

## 2. Functionality

### 2.1 Namespace-Level Configuration

The webhook uses Kubernetes labels to determine which namespaces should have LXCFS enabled. When a namespace is labeled with `lxcfs-admission-webhook: enabled`, the webhook will automatically configure the namespace to use LXCFS.

### 2.2 Pod-Level Exemption

In some cases, there may be specific pods within a namespace that should not be affected by the LXCFS configuration. 

To handle such scenarios, the webhook supports an annotation at the pod level. By adding the annotation `lxcfs-admission-webhook.raids-lab.github.io/mutate=false` to a pod, the webhook will skip any LXCFS-related modifications for that pod. This provides flexibility in managing LXCFS integration on a per-pod basis within a namespace.

## 3. Installation

Make sure you have the following prerequisites before proceeding with the installation:

- A Kubernetes cluster
- Helm installed
- Cert Manager installed

The LXCFS Admission Webhook can be installed using Helm, a popular package manager for Kubernetes. Follow the steps below to install the webhook:

1. Navigate to the directory containing the Helm chart for the LXCFS Admission Webhook. In this case, the chart is located in the `./dist/chart` directory.

2. Execute the following Helm command to install or upgrade the webhook in the lxcfs namespace:

```bash
helm upgrade --install lxcfs-webhook ./dist/chart -n lxcfs
```

This command will deploy all the necessary components of the webhook, including the webhook server, service, and related Kubernetes resources.

## 4. Configuration

### 4.1 Enabling LXCFS for a Namespace

To enable LXCFS for a namespace, add the following label to the namespace:

```bash
kubectl label namespace <namespace-name> lxcfs-admission-webhook:enabled
```

Replace <namespace-name> with the actual name of the namespace you want to configure. After adding the label, the webhook will detect the change and automatically configure the namespace to use LXCFS.

### 4.2 Exempting a Pod from LXCFS Configuration

To exempt a pod from LXCFS-related modifications, add the following annotation to the pod's manifest:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: <pod-name>
  annotations:
    lxcfs-admission-webhook.raids-lab.github.io/mutate: false
spec:
  containers:
    - name: <container-name>
      image: <image-name>
```

Replace <pod-name>, <container-name>, and <image-name> with the appropriate values for your pod. When creating or updating the pod, the webhook will respect this annotation and skip any LXCFS-related changes for the pod.

## 5. Conclusion

The LXCFS Admission Webhook simplifies the process of integrating LXCFS into Kubernetes environments by automating namespace and pod-level configurations. By following the installation and configuration steps outlined in this document, users can easily enable LXCFS for specific namespaces while having the flexibility to exempt individual pods as needed. Regular verification, testing, and troubleshooting will ensure the smooth operation of the webhook and the effective use of LXCFS within the Kubernetes cluster.

LXCFS Webhook is a Kubernetes admission controller that provides
a way to manage and enforce LXCFS configurations for containers
running in a Kubernetes cluster.

---
## Description
// TODO(user): An in-depth paragraph about your project and overview of use

## Getting Started

### Prerequisites
- go version v1.23.0+
- docker version 17.03+.
- kubectl version v1.11.3+.
- Access to a Kubernetes v1.11.3+ cluster.

### To Deploy on the cluster
**Build and push your image to the location specified by `IMG`:**

```sh
make docker-build docker-push IMG=<some-registry>/lxcfs-webhook:tag
```

**NOTE:** This image ought to be published in the personal registry you specified.
And it is required to have access to pull the image from the working environment.
Make sure you have the proper permission to the registry if the above commands don’t work.

**Install the CRDs into the cluster:**

```sh
make install
```

**Deploy the Manager to the cluster with the image specified by `IMG`:**

```sh
make deploy IMG=<some-registry>/lxcfs-webhook:tag
```

> **NOTE**: If you encounter RBAC errors, you may need to grant yourself cluster-admin
privileges or be logged in as admin.

**Create instances of your solution**
You can apply the samples (examples) from the config/sample:

```sh
kubectl apply -k config/samples/
```

>**NOTE**: Ensure that the samples has default values to test it out.

### To Uninstall
**Delete the instances (CRs) from the cluster:**

```sh
kubectl delete -k config/samples/
```

**Delete the APIs(CRDs) from the cluster:**

```sh
make uninstall
```

**UnDeploy the controller from the cluster:**

```sh
make undeploy
```

## Project Distribution

Following the options to release and provide this solution to the users.

### By providing a bundle with all YAML files

1. Build the installer for the image built and published in the registry:

```sh
make build-installer IMG=<some-registry>/lxcfs-webhook:tag
```

**NOTE:** The makefile target mentioned above generates an 'install.yaml'
file in the dist directory. This file contains all the resources built
with Kustomize, which are necessary to install this project without its
dependencies.

2. Using the installer

Users can just run 'kubectl apply -f <URL for YAML BUNDLE>' to install
the project, i.e.:

```sh
kubectl apply -f https://raw.githubusercontent.com/<org>/lxcfs-webhook/<tag or branch>/dist/install.yaml
```

### By providing a Helm Chart

1. Build the chart using the optional helm plugin

```sh
kubebuilder edit --plugins=helm/v1-alpha
```

2. See that a chart was generated under 'dist/chart', and users
can obtain this solution from there.

**NOTE:** If you change the project, you need to update the Helm Chart
using the same command above to sync the latest changes. Furthermore,
if you create webhooks, you need to use the above command with
the '--force' flag and manually ensure that any custom configuration
previously added to 'dist/chart/values.yaml' or 'dist/chart/manager/manager.yaml'
is manually re-applied afterwards.

## Contributing
// TODO(user): Add detailed information on how you would like others to contribute to this project

**NOTE:** Run `make help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

