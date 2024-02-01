# ack-spark-operator

一个用于部署 Apache Spark on Kubernetes operator 的 Helm chart。

## 概述

本 Chart 通过使用 [Helm](https://helm.sh/) 包管理器部署一个 deployment 资源用于运行 Apache  Spark on Kubernetes Operator。

ack-spark-operator 中定义了 `SparkApplication` 和 `ScheduledSparkApplication` 两种自定义资源，用户可以以更加符合 Kubernetes 原生的方式在 Kubernetes 集群中运行 Spark 作业或按照指定的时间表定时执行作业。

ack-spark-operator 基于 spark 2.4.5 版本，如果需要运行 Spark 3.x 版本的应用，请使用 `ack-spark-operator3.0`。

## 先决条件

- `Helm >= 3`
- `Kubernetes >= 1.14`

## 开始之前

Spark operator 和 Spark 作业可以运行在相同的命名空间中，也可以运行在不同的命名空间中：

- Spark operator 默认运行在命名空间 `spark-operator` 中，可以在执行 `helm install` 安装命令时使用 `--namespace` 参数执指定需要安装的命名空间。
- Spark 作业默认运行在命名空间 `default` 中，可以通过在 `values.yaml` 中修改 `spark.jobNamespace` 参数指定需要运行 Spark 作业的命名空间。

 `spark.jobNamespace` 取值默认为 `default`，这意味着 ack-spark-operator 只会监听命名空间 `default` 中提交的 Spark 应用，同时，会在命名空间 `default` 中创建名为 `spark` 的 ServiceAccount 并赋予了合适的权限。

如果你需要监听所有命名空间中提交的 Spark 应用，你需要将字段  `spark.jobNamespace` 设置为 `""`，此时只会在 Spark operator 所在的命名空间中创建 Spark 应用所需的 ServiceAccount。
因此，你需要手动在每个你想要提交 Spark 作业的命名空间中创建好相关 RBAC 资源并授予 Spark operator 和 Spark 应用合适的权限。
例如，你想要在 `default` 这个命名空间中运行 Spark 作业，则可以创建如下清单文件：

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

将上述资源清单保存为文件，例如 `spark-rbac-default.yaml`，然后执行如下命令创建资源：

```bash
kubectl apply -f spark-rbac-default.yaml
```

## 什么是 enableAlibabaCloudFeatureGates

`enableAlibabaCloudFeatureGates` 和 `alibabaCloudFeatureGates` 字段用于控制是否开启某些特性，如果你需要以更好的性能运行 Spark 作业，则可以将 `enableAlibabaCloudFeatureGates` 设置为 `true`，同时你的作业必须使用 `registry.aliyuncs.com/acs/spark:ack-2.4.5-latest` 作为 Spark 基础镜像。

## 安装

### 控制台安装

1. 登录容器服务 ACK 控制台；
2. 在应用市场中搜索并点击 `ack-spark-operator`；
3. 点击一键部署；
4. 选择集群、修改命名空间和发布名称，点击下一步；
5. 选择 Chart 版本并修改配置参数，点击确定。

### 命令行安装

添加阿里云 Helm 仓库：

```bash
helm repo add aliyunhub https://aliacs-app-catalog.oss-cn-hangzhou.aliyuncs.com/charts-incubator/
```

执行下述命令在命名空间 spark-operator 中安装最新版 ack-spark-operator，如果该命名空间不存在则创建，并从 `cn-bejing` 地域的镜像仓库中拉取镜像：

```bash
helm install ack-spark-operator aliyunhub/ack-spark-operator \
    --namespace spark-operator \
    --create-namespace \
    --set region=cn-beijing
```

注：

- 如需修改配置项，可以添加多个格式形如 `--set key1=value1,key2=value2` 的参数

## 配置项

下表列出了本 chart 中可以配置的参数以及它们的默认值：

| 参数                                  | 类型     | 描述                                                   | 默认值               |
| ------------------------------------- | -------- | ------------------------------------------------------ | -------------------- |
| `region`                              | string   | ACK 集群所在地域                                       | `__ACK__REGION_ID__` |
| `version`                             | string   | Operator 版本                                          | `v2.4.5-v1beta2`     |
| `image.registry`                      | string   | 镜像仓库名称，如果为空，则使用 ACK 镜像仓库            | `""`                 |
| `image.pullByVPC`                     | bool     | 当使用默认的 ACK 镜像仓库时，是否使用 VPC 网络拉取镜像 | `true`               |
| `image.repository`                    | string   | 镜像名称                                               | `acs/spark-operator` |
| `image.tag`                           | string   | 镜像标签                                               | `ack-2.4.5-ce32621`  |
| `image.pullPolicy`                    | string   | 镜像拉取策略                                           | `IfNotPresent`       |
| `image.pullSecrets`                   | []string | 镜像拉取秘钥                                           | `[]`                 |
| `sparkoperator.serviceAccount.create` | bool     | 是否为 operator 创建新的 ServiceAccount                | `true`               |
| `sparkoperator.serviceAccount.name`   | string   | operator 使用的 ServiceAccount 的名称                  | `ack-spark-operator` |
| `sparkoperator.rbac.create`           | bool     | 是否为 operator 使用的 ServiceAccount 创建 RBAC 资源   | `true`               |
| `spark.jobNamespace`                  | string   | Spark 作业运行的命名空间                               | `default`            |
| `spark.serviceAccount.create`         | bool     | 是否为 Spark 应用创建新的 ServiceAccount               | `true`               |
| `spark.serviceAccount.name`           | string   | Spark 应用使用的 ServiceAccount 的名称                 | `spark`              |
| `spark.rbac.create`                   | bool     | 是否为 Spark 使用的 ServiceAccount 创建 RBAC 资源      | `true`               |
| `webhook.enable`                      | bool     | 是否创建 webhook                                       | `true`               |
| `webhook.port`                        | int      | webhook service 端口号                                 | `8080`               |
| `metrics.enable`                      | bool     | 是否启用 Prometheus 监控指标抓取                       | `true`               |
| `metrics.port`                        | int      | 监控指标端口号                                         | `10254`              |
| `metrics.portName`                    | string   | 监控指标端口名                                         | `metrics`            |
| `metrics.endpoint`                    | string   | 监控指标服务端点                                       | `/metrics`           |
| `metrics.prefix`                      | string   | 监控指标前缀                                           | `""`                 |
| `logLevel`                            | int      | 日志级别                                               | `2`                  |
| `controllerThreads`                   | int      | Operator 并发度，内存使用量会随着并发度增加而增加      | `10`                 |
| `resyncInterval`                      | int      | Operator resync 时间间隔                               | `30`                 |
| `ingressUrlFormat`                    | string   | Ingress URL 格式                                       | `""`                 |
| `nodeSelector`                        | object   | Pod 调度时需要匹配的节点选择器                         | `{}`                 |
| `tolerations`                         | []       | Pod 容忍度                                             | `[]`                 |
| `enableAlibabaCloudFeatureGates`      | bool     | 是否启用阿里云特性                                     | `false`              |
| `alibabaCloudFeatureGates`            | string   | 启用的阿里云特性                                       | `""`                 |

## 发布说明

| 版本号                                 | 变更时间                                     | 变更内容         |
| ----------------------------------------- | ----------------------------------------------- | --------------- |
| `v0.1.17`                       |   2022年12月14日                 | 新增配置项 `ingress-url-format` |
| `v0.1.18` | 2024年01月02日 | （1）RBAC 权限收敛（2）更新模版和配置参数（3）增加中文 README |
| `v1.1.28` | 2024年01月08日 | 变更同 `v0.1.18`, 覆盖旧的 `v1.1.28` 版本 |
| `v1.1.29` | 2024年02月01日 | （1）支持在 SparkApplication 中添加 `scheduling.x-k8s.io/suspend: "true"` 注解实现 suspend 操作，从而允许 kube-queue 进行排队处理 |
