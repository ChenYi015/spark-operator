# ack-spark-operator3.0

A Helm chart for Apache Spark on Kubernetes operator.

## Introduction

This chart bootstraps a Apache Spark on Kubernetes deployment using the [Helm](https://helm.sh) package manager.

3.0 means the operator is implemented based on Spark 3.x version, since there is a big difference between Spark 2.x version and 3.x version, in order to avoid compatibility issues, if you need to run Spark 2.x version applications, please use ack-spark-operator instead.

## Prerequisites

- Helm >= 3
- Kubernetes >= 1.16

## Installing the chart

You can install the chart by Alibaba cloud container service ACK console or by helm cli.

### Container Service ACK Console

Install the chart with following steps:

1. Log into  Alibaba cloud container service ACK console
2. Search `ack-spark-operator3.0` in marketplace
3. Click `Deploy` button
4. Choose  ACK cluster and modify  namespace and release name, then click `Next` button
5. Choose chart version and modify configurations, then click `OK` button

### Helm CLI

Add Alibaba cloud helm registry:

```bash
helm repo add aliyunhub https://aliacs-app-catalog.oss-cn-hangzhou.aliyuncs.com/charts-incubator/
```

Then install the chart:

```bash
helm install ack-spark-operator3.0 aliyunhub/ack-spark-operator3.0 \
    --namespace spark-operator \
    --create-namespace \
    --set image.repository=registry.cn-beijing.aliyuncs.com/acs/spark-operator
```

This will create a release of `ack-spark-operator3.0` in the `spark-operator` namespace.

Note that `helm` will fail to install if the namespace doesn't exist. Either create the namespace beforehand or pass the `--create-namespace` flag to the `helm install` command.

## Uninstalling the chart

To uninstall release `ack-spark-operator3.0`:

```bash
helm uninstall -n spark-operator ack-spark-operator3.9
```

The command removes all the Kubernetes components associated with the chart and deletes the release, except for the `crds`, those will have to be removed manually.

## Configurations

| 参数                                        | 类型   | 默认值                                                       | 描述                                                         |
| ------------------------------------------- | ------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `replicaCount`                              | int    | `1`                                                          | Desired number of pods, leaderElection will be enabled if this is greater than 1 |
| `image.repository`                          | string | `registry-vpc.__ACK_REGION_ID__.aliyuncs.com/acs/spark-operator` | Image repository                                             |
| `image.pullPolicy`                          | string | `IfNotPresent`                                               | Image pull policy                                            |
| `image.tag`                                 | string | `v1beta2-1.3.8-3.1.1`                                        | if set, override the image tag whose default is the chart appVersion. |
| `imagePullSecrets`                          | list   | `[]`                                                         | Image pull secrets                                           |
| `nameOverride`                              | string | `""`                                                         | String to partially override `spark-operator.fullname` template (will maintain the release name) |
| `fullnameOverride`                          | string | `ack-spark-operator3`                                        | String to override release name                              |
| `rbac.create`                               | bool   | `false`                                                      | **DEPRECATED** use `createRole` and `createClusterRole`      |
| `rbac.createRole`                           | bool   | `true`                                                       | Create and use RBAC `Role` resources                         |
| `rbac.createClusterRole`                    | bool   | `true`                                                       | Create and use RBAC `ClusterRole` resources                  |
| `serviceAccounts.spark.create`              | bool   | `true`                                                       | Create a service account for spark apps                      |
| `serviceAccounts.spark.name`                | string | `""`                                                         | Optional name for the spark service account                  |
| `serviceAccounts.spark.annotations`         | object | `{}`                                                         | Optional annotations for the spark service account           |
| `serviceAccounts.sparkoperator.create`      | bool   | `true`                                                       | Create a service account for the operator                    |
| `serviceAccounts.sparkoperator.name`        | string | `""`                                                         | Optional name for the operator service account               |
| `serviceAccounts.sparkoperator.annotations` | object | `{}`                                                         | Optional annotations for the operator service account        |
| `sparkJobNamespace`                         | string | `""`                                                         | Set this if running spark jobs in a different namespace than the operator |
| `controllerThreads`                         | int    | `10`                                                         | Operator concurrency, higher values might increase memory usage |
| `resyncInterval`                            | int    | `30`                                                         | Operator resync interval. Note that the operator will respond to events (e.g. create, update) unrelated to this setting |
| `uiService.enable`                          | bool   | `true`                                                       | Enable UI service creation for Spark application             |
| `ingressUrlFormat`                          | string | `""`                                                         | Ingress URL format. Requires the UI service to be enabled by setting `uiService.enable` to true. |
| `logLevel`                                  | int    | `2`                                                          | Set higher levels for more verbose logging                   |
| `podSecurityContext`                        | object | `{}`                                                         | Pod security context                                         |
| `securityContext`                           | object | `{}`                                                         | Operator container security context                          |
| `volumes`                                   | list   | `[]`                                                         | Operator volumes                                             |
| `volumeMounts`                              | list   | `[]`                                                         | Operator volume mounts                                       |
| `webhook.enable`                            | bool   | `false`                                                      | Enable webhook server                                        |
| `webhook.port`                              | int    | `8080`                                                       | Webhook service port                                         |
| `webhook.namespaceSelector`                 | string | `""`                                                         | The webhook server will only operate on namespaces with this label, specified in the form key1=value1,key2=value2. Empty string (default) will operate on all namespaces |
| `webhook.initAnnotations`                   | object | `{"helm.sh/hook":"pre-install, pre-upgrade","helm.sh/hook-weight":"50"}` | The annotations applied to init job, required to restore certs deleted by the cleanup job during upgrade |
| `webhook.cleanupAnnotations`                | object | `{"helm.sh/hook":"pre-delete, pre-upgrade","helm.sh/hook-delete-policy":"hook-succeeded"}` | The annotations applied to the cleanup job, required for helm lifecycle hooks |
| `webhook.timeout`                           | int    | `30`                                                         | Webhook timeout in seconds                                   |
| `metrics.enable`                            | bool   | `true`                                                       | Enable prometheus metric scraping                            |
| `metrics.port`                              | int    | `10254`                                                      | Metrics port                                                 |
| `metrics.portName`                          | string | `metrics`                                                    | Metrics port name                                            |
| `metrics.endpoint`                          | string | `/metrics`                                                   | Metrics serving endpoint                                     |
| `metrics.prefix`                            | string | `""`                                                         | Metric prefix, will be added to all exported metrics         |
| `podMonitor.enable`                         | bool   | `false`                                                      | If enabled, a pod monitor for operator's pod will be submitted. Note that prometheus metrics should be enabled as well. |
| `podMonitor.labels`                         | object | `{}`                                                         | Pod monitor labels                                           |
| `podMonitor.jobLabel`                       | string | `spark-operator-podmonitor`                                  | The label to use to retrieve the job name from               |
| `podMonitor.podMetricsEndpoint`             | object | `{"interval":"5s","scheme":"http"}`                          | Prometheus metrics endpoint properties. `metrics.portName` will be used as a port |
| `nodeSelector`                              | object | `{}`                                                         | Node labels for pod assignment                               |
| `tolerations`                               | list   | `[]`                                                         | List of node taints to tolerate                              |
| `affinity`                                  | object | `{}`                                                         | Affinity for pod assignment                                  |
| `podAnnotations`                            | object | `{}`                                                         | Additional annotations to add to the pod                     |
| `podLabels`                                 | object | `{}`                                                         | Additional labels to add to the pod                          |
| `resources`                                 | object | `{}`                                                         | Pod resource requests and limits Note, that each job submission will spawn a JVM within the Spark Operator Pod using "/usr/local/openjdk-11/bin/java -Xmx128m". Kubernetes may kill these Java processes at will to enforce resource limits. When that happens, you will see the following error: 'failed to run spark-submit for SparkApplication [...]: signal: killed' - when this happens, you may want to increase memory limits. |
| `batchScheduler.enable`                     | bool   | `false`                                                      | Enable batch scheduler for spark jobs scheduling. If enabled, users can specify batch scheduler name in spark application |
| `resourceQuotaEnforcement.enable`           | bool   | `false`                                                      | Whether to enable the ResourceQuota enforcement for SparkApplication resources. Requires the webhook to be enabled by setting `webhook.enable` to true. |
| `leaderElection.lockName`                   | string | `spark-operator-lock`                                        | Leader election lock name.                                   |
| `leaderElection.lockNamespace`              | string | `""`                                                         | Optionally store the lock in another namespace. Defaults to operator's namespace |
| `istio.enabled`                             | bool   | `false`                                                      | When using `istio`, spark jobs need to run without a sidecar to properly terminate |
| `labelSelectorFilter`                       | string | ""                                                           | A comma-separated list of key=value, or key labels to filter resources during watch and list based on the specified labels. |

## Release Note

| Version   | Release Date | Changes                                                      |
| --------- | ------------ | ------------------------------------------------------------ |
| `v1.1.26` | Mar 14, 2023 | (1) Add ack-spark-operator3.0 chart                          |
| `v1.1.27` | Jan 4, 2024  | (1) Update webhook init and cleanup job templates (2) Add `nameOverride` and `fullnameOverride` configurations (3) Update RBAC rules for HA (4) Add Chinese README |
