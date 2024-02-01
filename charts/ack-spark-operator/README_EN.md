# ack-spark-operator

A Helm chart for  [Apache Spark on Kubernetes Operator](https://github.com/AliyunContainerService/spark-on-k8s-operator/tree/alibabacloud-v1beta2-1.1.0-2.4.5).

## Prerequisites

- `Helm >= 3`
- `Kubernetes >= 1.14`

## Before you start

The spark operator would watch spark applications in namespaces which `spark.jobNamespace` specified. If you want to watch all namespaces, you can set `spark.jobNamespace` to `""`. And create RBAC resources in every namespace you want to run spark applications. For example, if you want to run spark applications  in namespace `default`, you can create the following manifest and save it as file `spark-rbac-default.yaml` :

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: spark
  namespace: default

---
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  namespace: default
  name: spark
rules:
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["*"]
- apiGroups: [""]
  resources: ["services"]
  verbs: ["*"]

---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: spark
  namespace: default
subjects:
- kind: ServiceAccount
  name: spark
  namespace: default
roleRef:
  kind: Role
  name: spark
  apiGroup: rbac.authorization.k8s.io

---
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  namespace: default
  name: ack-spark-operator
rules:
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["*"]
- apiGroups: [""]
  resources: ["configmaps"]
  verbs: ["*"]
- apiGroups: [""]
  resources: ["services"]
  verbs: ["create", "get", "delete"]
- apiGroups: [""]
  resources: ["persistentvolumeclaims"]
  verbs: ["create", "get", "list", "watch", "update", "patch"]
- apiGroups: [""]
  resources: ["events"]
  verbs: ["create", "update", "patch"]
- apiGroups: [""]
  resources: ["resourcequotas"]
  verbs: ["get", "list", "watch"]
- apiGroups: ["sparkoperator.k8s.io"]
  resources: ["sparkapplications", "scheduledsparkapplications", "sparkapplications/status", "scheduledsparkapplications/status"]
  verbs: ["*"]
- apiGroups: ["extensions"]
  resources: ["ingresses"]
  verbs: ["create", "get", "delete"]

---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: ack-spark-operator
  namespace: default
subjects:
- kind: ServiceAccount
  name: ack-spark-operator
  namespace: spark-operator
roleRef:
  kind: Role
  name: ack-spark-operator
  apiGroup: rbac.authorization.k8s.io
```

Then execute below command to create RBAC resources:

```bash
kubectl apply -f spark-rbac-default.yaml
```

## What is enableAlibabaCloudFeatureGates

`enableAlibabaCloudFeatureGates` is a featuregates option for spark operator. You can enable enableAlibabaCloudFeatureGates to run spark in high performance mode. If you want to use `enableAlibabaCloudFeatureGates`, You need to use `registry.aliyuncs.com/acs/spark:ack-2.4.5-latest` as base image.

## Installing the chart

Note that you need to use the `--namespace` flag during `helm install` to specify in which namespace you want to install the operator. The namespace can be existing or not. When it's not available, Helm would take care of creating the namespace. Note that this namespace has no relation to the namespace where you would like to deploy Spark jobs (i.e. the setting `sparkJobNamespace` shown in the table below). They can be the same namespace or different ones.

## Configuration

The following table lists the configurable parameters of the Spark operator chart and their default values.

| Parameter                             | Type     | Description                                                  | Default              |
| ------------------------------------- | -------- | ------------------------------------------------------------ | -------------------- |
| `region`                              | string   | ACK cluster region                                           | `__ACK_REGION_ID__`  |
| `version`                             | string   | The version of the operator to install                       | `v2.4.5-v1beta2`     |
| `image.registry`                      | string   | The name of the image registry                               | `""`                 |
| `image.pullByVPC`                     | bool     | Whether to pull images by VPC network when use default image registry. | `true`               |
| `image.repository`                    | string   | The name of the operator image                               | `acs/spark-operator` |
| `image.tag`                           | string   | The tag of the operator image                                | `ack-2.4.5-ce32621`  |
| `image.pullPolicy`                    | string   | Image pull policy                                            | `IfNotPresent`       |
| `image.pullSecrets`                   | []string | The names of secrets to use when pulling images from customized image registry | `[]`                 |
| `sparkoperator.serviceAccount.create` | bool     | Specifies whether to create a service account for spark operator | `true`               |
| `sparkoperator.serviceAccount.name`   | string   | The name of service account to be used by spark operator     | `ack-spark-operator` |
| `sparkoperator.rbac.create`           | bool     | Specifies whether to create RBAC resources for the service account used by spark operator | `true`               |
| `spark.jobNamespace`                  | string   | The namespace where spark applications deploy                | `default`            |
| `spark.serviceAccount.create`         | bool     | Specifies whether to create a service account for spark applications | `true`               |
| `spark.serviceAccount.name`           | string   | The name of service account to be used by spark applications | `spark`              |
| `spark.rbac.create`                   | bool     | Specifies whether to create RBAC resources for the service account used by spark applications. | `true`               |
| `webhook.enable`                      | bool     | Specifies whether to create webhooks                         | `true`               |
| `webhook.port`                        | int      | The port number of webhook service                           | `8080`               |
| `metrics.enable`                      | bool     | Specifies whether to enable Prometheus metrics scraping      | `true`               |
| `metrics.port`                        | int      | The metrics port number                                      | `10254`              |
| `metrics.portName`                    | string   | The metrics port name                                        | `metrics`            |
| `metrics.endpoint`                    | string   | The metrics serving endpoint                                 | `/metrics`           |
| `metrics.prefix`                      | string   | The metrics prefix to be added                               | `""`                 |
| `logLevel`                            | int      | Set higher level to collect more verbose logs.               | `2`                  |
| `controllerThreads`                   | int      | Number of worker threads used by the SparkApplication controller | `10`                 |
| `resyncInterval`                      | int      | Informer resync interval in seconds                          | `30`                 |
| `ingressUrlFormat`                    | string   | Ingress URL format                                           | `""`                 |
| `nodeSelector`                        | object   | Pod node selector                                            | `{}`                 |
| `tolerations`                         | []       | Pod tolerations                                              | `[]`                 |
| `enableAlibabaCloudFeatureGates`      | bool     | Specifies whether to enable Alibaba cloud feature gates      | `false`              |
| `alibabaCloudFeatureGates`            | string   | Specifies which features to enable                           | `""`                 |

## Release Note

| Version                                 | Release Date                                     | Changes        |
| ----------------------------------------- | ----------------------------------------------- | --------------- |
| `v0.1.17`                       |   Dec 14 ,2022                | (1) Add new parameter `ingress-url-format` |
| `v0.1.18`                       | Jan 3, 2024 |   (1) Narrow RBAC rules (2) Update templates and parameters (3) Add Chinese README |
| `v1.1.28` | Jan 8, 2024 | The same with `v0.1.18`, override the old `v1.1.28` release  |
| `v1.1.29` | Feb 1, 2024  | (1) Support adding `scheduling.x-k8s.io/suspend: "true"` annotation in SparkApplication to implement suspend operation, thereby allowing kube-queue to perform queuing processing |
