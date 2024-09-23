# ack-spark-operator

![Version: 2.0.0](https://img.shields.io/badge/Version-2.0.0-informational?style=flat-square) ![AppVersion: 2.0.0](https://img.shields.io/badge/AppVersion-2.0.0-informational?style=flat-square)

一个用于在 ACK 集群中部署 Spark operator 的 Helm chart。

**Homepage:** <https://github.com/kubeflow/spark-operator>

## 概述

Apache Spark 是一种专门为大规模数据处理而设计的快速且通用的计算引擎，并已经广泛应用于各行各业的大数据处理场景中。
Spark Operator 是专门为在 Kubernetes 集群中运行 Spark 工作负载而设计的 Operator，它定义了两种自定义资源并实现了相应的控制器，用户可以通过编写 `SparkApplication` 清单文件提交一次性 Spark 作业，或者通过编写 `ScheduledSparkApplication` 清单文件提交定时 Spark 作业。使用 Spark Operator 运行 Spark 作业有如下优势：

- 自动化管理 Spark 作业的生命周期。
- 支持部分在 Spark Core 中没有实现的 Kubernetes 特性。

## 先决条件

- [Kubernetes](https://kubernetes.io/) >= 1.24
- [Helm](https://helm.sh) >= 3

## 安装

部署 ack-spark-operator 组件请按照如下步骤进行操作：

1. 登录[容器服务管理控制台](https://cs.console.aliyun.com/)，在左侧导航栏**选择市场** > **应用市场**。
2. 在**应用市场**页面，搜索 `ack-spark-operator`，然后单击该应用。
3. 在**应用详情**页面，单击右上角的**一键部署**，然后按照页面提示进行部署。
4. 在**基本信息**页面，填写**目标集群**、**命名空间**和**发布名称**后，单击**下一步**。
5. 在**参数配置**页面，选择Chart版本，将参数配置完成后，单击**确定**。

## 配置项

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| nameOverride | string | `"spark-operator"` | 部分覆盖发布名称的名字 |
| fullnameOverride | string | `""` | 完全覆盖发布名称的名字 |
| commonLabels | object | `{}` | 将会添加标签至所创建的资源 |
| image.registry | string | `"registry-__ACK_REGION_ID__-vpc.ack.aliyuncs.com"` | 镜像仓库 |
| image.repository | string | `"acs/spark-operator"` | 镜像名称 |
| image.tag | string | `"2.0.0-ack"` | 镜像标签 |
| image.pullPolicy | string | `"IfNotPresent"` | 镜像拉取策略 |
| image.pullSecrets | list | `[]` | 镜像拉取秘钥 |
| controller.replicas | int | `1` | 控制器副本数量 |
| controller.workers | int | `10` | 调谐并发度, 值越大内存消耗越高 |
| controller.logLevel | string | `"info"` | 日志级别，可以是 `debug`, `info`, `error`。 |
| controller.uiService.enable | bool | `true` | 指定是否需要为 Spark web UI 创建 Service 资源。 |
| controller.uiIngress.enable | bool | `false` | 指定是否需要为 Spark web UI 创建 Ingress 资源。 `controller.uiService.enable` 必须设置为 `true` 以启用创建 ingress 功能。 |
| controller.uiIngress.urlFormat | string | `""` | Ingress URL 格式。 当 `controller.uiIngress.enable` 设置为 `true` 时必填. |
| controller.batchScheduler.enable | bool | `false` | 指定是否为 Spark 作业启用 batch 调度功能。 如果启用，用户可以在 Spark 作业中指定 batch 调度器名称。 |
| controller.batchScheduler.kubeSchedulerNames | list | `[]` | 指定 kube-scheduler 调度器名称列表。 |
| controller.batchScheduler.default | string | `""` | (可选）为 Spark 作业指定默认 batch 调度器名称，当提交 Spark 作业时，如果未指定 batch 调度器，则会使用该默认值。 支持的值可以是 `volcano`、`yunikorn` 和注册的 kube-scheduler 名称，使用其他值将会导致控制器启动失败。 |
| controller.serviceAccount.create | bool | `true` | 是否为控制器创建服务账号 |
| controller.serviceAccount.name | 可选 | `""` | 控制器使用的服务账号名称 |
| controller.serviceAccount.annotations | object | `{}` | 为控制器的服务账号指定额外注解。 |
| controller.rbac.create | bool | `true` | 是否为控制器创建 RBAC 资源。 |
| controller.rbac.annotations | object | `{}` | 为控制器的 RBAC 资源指定额外注解。 |
| controller.labels | object | `{}` | 控制器 pod 额外标签 |
| controller.annotations | object | `{}` | 控制器 pod 额外注解 |
| controller.volumes | list | `[]` | 控制器 pod 额外数据卷 |
| controller.nodeSelector | object | `{}` | 控制器 pod 节点选择器 |
| controller.affinity | object | `{}` | 控制器 pod 亲和性 |
| controller.tolerations | list | `[]` | 控制器 pod 污点容忍 |
| controller.priorityClassName | string | `""` | 控制器 pod 优先级类 |
| controller.podSecurityContext | object | `{}` | 控制器 pod 安全上下文 |
| controller.topologySpreadConstraints | list | `[]` | 控制器 pod 拓扑分布约束 更多信息请参考: [Pod Topology Spread Constraints](https://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/). |
| controller.env | list | `[]` | 控制器容器环境变量 |
| controller.envFrom | list | `[]` | 控制器容器环境变量源 |
| controller.volumeMounts | list | `[]` | 控制器容器数据卷挂载 |
| controller.resources | object | `{}` | 控制器容器资源需求 |
| controller.securityContext | object | `{}` | 控制器容器安全上下文 |
| controller.sidecars | list | `[]` | Sidecar 容器 |
| controller.podDisruptionBudget.enable | bool | `false` | 指定是否为控制器创建 Pod disruption budget。 更多信息请参见: [Specifying a Disruption Budget for your Application](https://kubernetes.io/docs/tasks/run-application/configure-pdb/) |
| controller.podDisruptionBudget.minAvailable | int | `1` | 最小可用的 Pod 数量，需要 `controller.replicas` 大于 `1`。 |
| controller.pprof.enable | bool | `false` | 是否启用 pprof |
| controller.pprof.port | int | `6060` | 指定 pprof 端口号 |
| controller.pprof.portName | string | `"pprof"` | 指定 pprof 端口号名称 |
| webhook.enable | bool | `true` | 是否启用 webhook |
| webhook.replicas | int | `1` | Webhook 副本数量 |
| webhook.logLevel | string | `"info"` | Webhook 日志级别，可以是 `debug`、`info` 或 `error`。 |
| webhook.port | int | `9443` | Webhook 端口号 |
| webhook.portName | string | `"webhook"` | Webhook 端口名称 |
| webhook.failurePolicy | string | `"Fail"` | Webhook 调用失败处理策略，支持的选项包括 `Ignore` 和 `Fail`。 |
| webhook.timeoutSeconds | int | `10` | 指定 Webhook 调用超时时间（单位：秒），值必须为 1 到 30 之间。 |
| webhook.resourceQuotaEnforcement.enable | bool | `false` | 指定是否需要启用资源配额功能 |
| webhook.serviceAccount.create | bool | `true` | 是否为 webhook 创建服务账号 |
| webhook.serviceAccount.name | string | `""` | （可选）Webhook 服务账号名称 |
| webhook.serviceAccount.annotations | object | `{}` | Webhook 服务账号额外注解 |
| webhook.rbac.create | bool | `true` | 是否为 Webhook 创建 RBAC 资源。 |
| webhook.rbac.annotations | object | `{}` | 为 Webhook 的 RBAC 资源指定额外注解。 |
| webhook.labels | object | `{}` | Webhook pod 额外标签 |
| webhook.annotations | object | `{}` | Webhook pod 额外注解 |
| webhook.sidecars | list | `[]` | Webhook pod sidecar 容器 |
| webhook.volumes | list | `[]` | Webhook pod 额外数据卷 |
| webhook.nodeSelector | object | `{}` | Webhook pod 节点选择器 |
| webhook.affinity | object | `{}` | Webhook pod 亲和性 |
| webhook.tolerations | list | `[]` | Webhook pod 污点容忍 |
| webhook.priorityClassName | string | `""` | Webhook pod 优先级类 |
| webhook.podSecurityContext | object | `{}` | Webhook pod 安全上下文 |
| webhook.topologySpreadConstraints | list | `[]` | Webhook pod 拓扑分布约束 更多信息请参考: [Pod Topology Spread Constraints](https://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/). |
| webhook.env | list | `[]` | Webhook 容器环境变量 |
| webhook.envFrom | list | `[]` | Webhook 容器环境变量源 |
| webhook.volumeMounts | list | `[]` | Webhook 容器数据卷挂载 |
| webhook.resources | object | `{}` | Webhook 容器资源需求 |
| webhook.securityContext | object | `{}` | Webhook 容器安全上下文 |
| webhook.podDisruptionBudget.enable | bool | `false` | 是否为 Webhook 创建 Pod disruption budget， 更多信息请参见: [Specifying a Disruption Budget for your Application](https://kubernetes.io/docs/tasks/run-application/configure-pdb/) |
| webhook.podDisruptionBudget.minAvailable | int | `1` | 最小可用的 Pod 副本数量，需要 `webhook.replicas` 大于 `1`。` |
| spark.jobNamespaces | list | `["default"]` | 可以运行 spark 任务的命名空间列表，如果包含空字符串，则允许所有命名空间。 注意：需要确保命名空间已经存在。 |
| spark.serviceAccount.create | bool | `true` | 是否为 Spark 创建服务账号 |
| spark.serviceAccount.name | string | `""` | （可选）Spark 服务账号名称 |
| spark.serviceAccount.annotations | object | `{}` | （可选）Spark 服务账号额外注解 |
| spark.rbac.create | bool | `true` | 指定是否为 Spark 创建 RBAC 资源 |
| spark.rbac.annotations | object | `{}` | （可选）为 Spark 的 RBAC 资源指定额外注解 |
| prometheus.metrics.enable | bool | `true` | 是否启用 Prometheus 监控指标抓取 |
| prometheus.metrics.port | int | `8080` | Metric 端口号 |
| prometheus.metrics.portName | string | `"metrics"` | Metrics 端口名称 |
| prometheus.metrics.endpoint | string | `"/metrics"` | Metrics 服务端点 |
| prometheus.metrics.prefix | string | `""` | Metrics 名称前缀，将会添加至所有导出的指标 |
| prometheus.podMonitor.create | bool | `false` | 是否创建 Pod monitor 资源，需要将 `prometheus.metrics.enable` 设置为 `true` 以启用 Prometheus。 |
| prometheus.podMonitor.labels | object | `{}` | Pod monitor 标签 |
| prometheus.podMonitor.jobLabel | string | `"spark-operator-podmonitor"` | 从作业标签中检索出作业名称 |
| prometheus.podMonitor.podMetricsEndpoint | object | `{"interval":"5s","scheme":"http"}` | Prometheus 监控指标 endpoint 属性，`metrics.portName` 将被用作端口 |

## 维护者

| 姓名 | 邮件 |
| ---- | ------ |
| 听泉 | <tingquan.cy@alibaba-inc.com> |

## 发布说明

| 版本号 | 变更日期 | 变更内容 |
| --- | --- | --- |
| 0.1.17 | 2022年12月14日 |（1）新增配置项 `ingress-url-format`<br> |
| 0.1.18 | 2024年01月02日 |（1）RBAC 权限收敛<br>（2）更新模版和配置参数<br>（3）增加中文 README<br> |
| 1.1.28 | 2024年01月08日 |（1）变更同 `v0.1.18`, 覆盖旧的 `v1.1.28` 版本<br> |
| 1.1.29 | 2024年02月01日 |（1）支持在 SparkApplication 中添加 `scheduling.x-k8s.io/suspend: "true"` 注解实现 suspend 操作，从而允许 kube-queue 进行排队处理<br> |
| 1.1.30 | 2024年04月08日 |（1）修复 RBAC 权限导致作业提交失败的问题<br> |
| 1.1.31 | 2024年07月31日 |（1）修复 webhook 与 operator pod 标签不匹配的问题<br> |
| 2.0.0 | 2024年09月23日 |（1）同步 [kubeflow/spark-operator](https://github.com/kubeflow/spark-operator) 2.0.0 版本<br> |
