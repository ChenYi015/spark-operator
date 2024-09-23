# ack-spark-operator

![Version: 2.0.0](https://img.shields.io/badge/Version-2.0.0-informational?style=flat-square) ![AppVersion: 2.0.0](https://img.shields.io/badge/AppVersion-2.0.0-informational?style=flat-square)

This chart bootstraps a [Kubernetes Operator for Apache Spark](https://github.com/kubeflow/spark-operator) deployment using the [Helm](https://helm.sh) package manager.

**Homepage:** <https://github.com/kubeflow/spark-operator>

## Introduction

Apache Spark is a fast and general-purpose computing engine specifically designed for large-scale data processing, and it has been widely applied in various big data processing scenarios across industries.

The Spark Operator is specifically designed for running Spark workloads in Kubernetes clusters. It defines two custom resources and implements corresponding controllers, allowing users to submit one-time Spark jobs by writing a SparkApplication manifest file, or to submit scheduled Spark jobs by writing a ScheduledSparkApplication manifest file. Running Spark jobs with the Spark Operator has the following advantages:

- Automated management of the lifecycle of Spark applications.
- Support for certain Kubernetes features not implemented in Spark Core.

## Prerequisites

- [Kubernetes](https://kubernetes.io/) >= 1.24
- [Helm](https://helm.sh) >= 3

## Install

Install the chart with following steps:

1. Log in to the [Container Service Console](https://cs.console.aliyun.com/), and in the left navigation bar, select **Marketplace** > **Marketplace**.
2. On the **Marketplace** page, search for **ack-spark-operator** and then click on the application.
3. On the **Application Details** page, click on the **Deploy** button in the top right corner, and follow the on-screen instructions to proceed with the deployment.
4. On the **Basic Information** page, after filling in the **Cluster**, **Namespace**, and **Release Name**, click on **Next**.
5. On the **Parameter Configuration** page, select the Chart version, configure the parameters, and then click on **OK**.

## Configurations

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| nameOverride | string | `"spark-operator"` | String to partially override release name. |
| fullnameOverride | string | `""` | String to fully override release name. |
| commonLabels | object | `{}` | Common labels to add to the resources. |
| image.registry | string | `"docker.io"` | Image registry. |
| image.repository | string | `"kubeflow/spark-operator"` | Image repository. |
| image.tag | string | If not set, the chart appVersion will be used. | Image tag. |
| image.pullPolicy | string | `"IfNotPresent"` | Image pull policy. |
| image.pullSecrets | list | `[]` | Image pull secrets for private image registry. |
| controller.replicas | int | `1` | Number of replicas of controller. |
| controller.workers | int | `10` | Reconcile concurrency, higher values might increase memory usage. |
| controller.logLevel | string | `"info"` | Configure the verbosity of logging, can be one of `debug`, `info`, `error`. |
| controller.uiService.enable | bool | `true` | Specifies whether to create service for Spark web UI. |
| controller.uiIngress.enable | bool | `false` | Specifies whether to create ingress for Spark web UI. `controller.uiService.enable` must be `true` to enable ingress. |
| controller.uiIngress.urlFormat | string | `""` | Ingress URL format. Required if `controller.uiIngress.enable` is true. |
| controller.batchScheduler.enable | bool | `false` | Specifies whether to enable batch scheduler for spark jobs scheduling. If enabled, users can specify batch scheduler name in spark application. |
| controller.batchScheduler.kubeSchedulerNames | list | `[]` | Specifies a list of kube-scheduler names for scheduling Spark pods. |
| controller.batchScheduler.default | string | `""` | Default batch scheduler to be used if not specified by the user. If specified, this value must be either "volcano" or "yunikorn". Specifying any other value will cause the controller to error on startup. |
| controller.serviceAccount.create | bool | `true` | Specifies whether to create a service account for the controller. |
| controller.serviceAccount.name | string | `""` | Optional name for the controller service account. |
| controller.serviceAccount.annotations | object | `{}` | Extra annotations for the controller service account. |
| controller.rbac.create | bool | `true` | Specifies whether to create RBAC resources for the controller. |
| controller.rbac.annotations | object | `{}` | Extra annotations for the controller RBAC resources. |
| controller.labels | object | `{}` | Extra labels for controller pods. |
| controller.annotations | object | `{}` | Extra annotations for controller pods. |
| controller.volumes | list | `[]` | Volumes for controller pods. |
| controller.nodeSelector | object | `{}` | Node selector for controller pods. |
| controller.affinity | object | `{}` | Affinity for controller pods. |
| controller.tolerations | list | `[]` | List of node taints to tolerate for controller pods. |
| controller.priorityClassName | string | `""` | Priority class for controller pods. |
| controller.podSecurityContext | object | `{}` | Security context for controller pods. |
| controller.topologySpreadConstraints | list | `[]` | Topology spread constraints rely on node labels to identify the topology domain(s) that each Node is in. Ref: [Pod Topology Spread Constraints](https://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/). The labelSelector field in topology spread constraint will be set to the selector labels for controller pods if not specified. |
| controller.env | list | `[]` | Environment variables for controller containers. |
| controller.envFrom | list | `[]` | Environment variable sources for controller containers. |
| controller.volumeMounts | list | `[]` | Volume mounts for controller containers. |
| controller.resources | object | `{}` | Pod resource requests and limits for controller containers. Note, that each job submission will spawn a JVM within the controller pods using "/usr/local/openjdk-11/bin/java -Xmx128m". Kubernetes may kill these Java processes at will to enforce resource limits. When that happens, you will see the following error: 'failed to run spark-submit for SparkApplication [...]: signal: killed' - when this happens, you may want to increase memory limits. |
| controller.securityContext | object | `{}` | Security context for controller containers. |
| controller.sidecars | list | `[]` | Sidecar containers for controller pods. |
| controller.podDisruptionBudget.enable | bool | `false` | Specifies whether to create pod disruption budget for controller. Ref: [Specifying a Disruption Budget for your Application](https://kubernetes.io/docs/tasks/run-application/configure-pdb/) |
| controller.podDisruptionBudget.minAvailable | int | `1` | The number of pods that must be available. Require `controller.replicas` to be greater than 1 |
| controller.pprof.enable | bool | `false` | Specifies whether to enable pprof. |
| controller.pprof.port | int | `6060` | Specifies pprof port. |
| controller.pprof.portName | string | `"pprof"` | Specifies pprof service port name. |
| webhook.enable | bool | `true` | Specifies whether to enable webhook. |
| webhook.replicas | int | `1` | Number of replicas of webhook server. |
| webhook.logLevel | string | `"info"` | Configure the verbosity of logging, can be one of `debug`, `info`, `error`. |
| webhook.port | int | `9443` | Specifies webhook port. |
| webhook.portName | string | `"webhook"` | Specifies webhook service port name. |
| webhook.failurePolicy | string | `"Fail"` | Specifies how unrecognized errors are handled. Available options are `Ignore` or `Fail`. |
| webhook.timeoutSeconds | int | `10` | Specifies the timeout seconds of the webhook, the value must be between 1 and 30. |
| webhook.resourceQuotaEnforcement.enable | bool | `false` | Specifies whether to enable the ResourceQuota enforcement for SparkApplication resources. |
| webhook.serviceAccount.create | bool | `true` | Specifies whether to create a service account for the webhook. |
| webhook.serviceAccount.name | string | `""` | Optional name for the webhook service account. |
| webhook.serviceAccount.annotations | object | `{}` | Extra annotations for the webhook service account. |
| webhook.rbac.create | bool | `true` | Specifies whether to create RBAC resources for the webhook. |
| webhook.rbac.annotations | object | `{}` | Extra annotations for the webhook RBAC resources. |
| webhook.labels | object | `{}` | Extra labels for webhook pods. |
| webhook.annotations | object | `{}` | Extra annotations for webhook pods. |
| webhook.sidecars | list | `[]` | Sidecar containers for webhook pods. |
| webhook.volumes | list | `[]` | Volumes for webhook pods. |
| webhook.nodeSelector | object | `{}` | Node selector for webhook pods. |
| webhook.affinity | object | `{}` | Affinity for webhook pods. |
| webhook.tolerations | list | `[]` | List of node taints to tolerate for webhook pods. |
| webhook.priorityClassName | string | `""` | Priority class for webhook pods. |
| webhook.podSecurityContext | object | `{}` | Security context for webhook pods. |
| webhook.topologySpreadConstraints | list | `[]` | Topology spread constraints rely on node labels to identify the topology domain(s) that each Node is in. Ref: [Pod Topology Spread Constraints](https://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/). The labelSelector field in topology spread constraint will be set to the selector labels for webhook pods if not specified. |
| webhook.env | list | `[]` | Environment variables for webhook containers. |
| webhook.envFrom | list | `[]` | Environment variable sources for webhook containers. |
| webhook.volumeMounts | list | `[]` | Volume mounts for webhook containers. |
| webhook.resources | object | `{}` | Pod resource requests and limits for webhook pods. |
| webhook.securityContext | object | `{}` | Security context for webhook containers. |
| webhook.podDisruptionBudget.enable | bool | `false` | Specifies whether to create pod disruption budget for webhook. Ref: [Specifying a Disruption Budget for your Application](https://kubernetes.io/docs/tasks/run-application/configure-pdb/) |
| webhook.podDisruptionBudget.minAvailable | int | `1` | The number of pods that must be available. Require `webhook.replicas` to be greater than 1 |
| spark.jobNamespaces | list | `["default"]` | List of namespaces where to run spark jobs. If empty string is included, all namespaces will be allowed. Make sure the namespaces have already existed. |
| spark.serviceAccount.create | bool | `true` | Specifies whether to create a service account for spark applications. |
| spark.serviceAccount.name | string | `""` | Optional name for the spark service account. |
| spark.serviceAccount.annotations | object | `{}` | Optional annotations for the spark service account. |
| spark.rbac.create | bool | `true` | Specifies whether to create RBAC resources for spark applications. |
| spark.rbac.annotations | object | `{}` | Optional annotations for the spark application RBAC resources. |
| prometheus.metrics.enable | bool | `true` | Specifies whether to enable prometheus metrics scraping. |
| prometheus.metrics.port | int | `8080` | Metrics port. |
| prometheus.metrics.portName | string | `"metrics"` | Metrics port name. |
| prometheus.metrics.endpoint | string | `"/metrics"` | Metrics serving endpoint. |
| prometheus.metrics.prefix | string | `""` | Metrics prefix, will be added to all exported metrics. |
| prometheus.podMonitor.create | bool | `false` | Specifies whether to create pod monitor. Note that prometheus metrics should be enabled as well. |
| prometheus.podMonitor.labels | object | `{}` | Pod monitor labels |
| prometheus.podMonitor.jobLabel | string | `"spark-operator-podmonitor"` | The label to use to retrieve the job name from |
| prometheus.podMonitor.podMetricsEndpoint | object | `{"interval":"5s","scheme":"http"}` | Prometheus metrics endpoint properties. `metrics.portName` will be used as a port |

## Maintainers

| Name | Email |
| ---- | ------ |
| 听泉 | <tingquan.cy@alibaba-inc.com> |

## Release Note

| version | date | changes |
| ------- | ---- | ------- |
| 0.1.17 | Dec 14 ,2022 |（1）Add new parameter `ingress-url-format`<br> |
| 0.1.18 | Jan 3, 2024 |（1）Narrow RBAC rules<br>（2）Update templates and parameters<br>（3）Add Chinese README<br> |
| 1.1.28 | Jan 8, 2024 |（1）The same with `v0.1.18`, override the old `v1.1.28` release<br> |
| 1.1.29 | Feb 1, 2024 |（1）Support adding `scheduling.x-k8s.io/suspend: "true"` annotation in SparkApplication to implement suspend operation, thereby allowing kube-queue to perform queuing processing<br> |
| 1.1.30 | Apr 8, 2024 |（1）Update RBAC resources<br> |
| 1.1.31 | June 31, 2024 |（1）map[Fix:webhook nodeSelector did not match operator pod labels]<br> |
| 2.0.0 | Sep 23, 2024 |（1）Sync [kubeflow/spark-operator](https://github.com/kubeflow/spark-operator) 2.0.0 version<br> |
