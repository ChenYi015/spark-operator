# ack-spark-operator3.0

一个用于部署 Apache Spark on Kubernetes operator 的 Helm chart。

## 概述

本 Chart 通过使用 [Helm](https://helm.sh/) 包管理器部署一个 deployment 资源用于运行 Apache  Spark on Kubernetes Operator。

3.0 版本表示该 operator 基于 Spark 3.x 版本实现，由于 Spark 2.x 版本和 3.x 版本差异较大，为避免兼容性问题，如果需要运行 Spark 2.x 版本作业，请使用 ack-spark-operator。

## 先决条件

- Helm >= 3
- Kubernetes >= 1.16

## 安装

你可以通过阿里云容器服务 ACK 控制台或 Helm 命令行方式安装本 Chart。

### 容器服务 ACK 控制台

安装方式如下：

1. 登录容器服务 ACK 控制台；
2. 在应用市场中搜索并点击 `ack-spark-operator3.0`；
3. 点击一键部署；
4. 选择集群、修改命名空间和发布名称，点击下一步；
5. 选择 Chart 版本并修改配置参数，点击确定。

### Helm 命令行

添加阿里云 Helm 仓库：

```bash
helm repo add aliyunhub https://aliacs-app-catalog.oss-cn-hangzhou.aliyuncs.com/charts-incubator/
```

执行下述命令在命名空间 spark-operator 中安装最新版 ack-spark-operator3.0，如果该命名空间不存在则创建，并从 `cn-beijing` 地域的镜像仓库中拉取镜像：

```bash
helm install ack-spark-operator3.0 aliyunhub/ack-spark-operator3.0 \
    --namespace spark-operator \
    --create-namespace \
    --set image.repository=registry.cn-beijing.aliyuncs.com/acs/spark-operator
```

注：

- 如需修改配置项，可以添加多个格式形如 `--set key1=value1,key2=value2` 的参数，或者修改 values.yaml 配置文件，并指定 `--values values.yaml` 参数。

执行如下命令可以卸载本 chart：

```bash
helm uninstall -n spark-operator ack-spark-operator3.0
```

上述命令会卸载除了 CRD 资源之外所有与本 chart 关联的 Kubernetes 资源，CRD 资源需要你手动删除，需要注意的是，一旦删除了 CRD，则所有已创建的 CR 都会删除。

## 配置

下表列出了本 chart 中可以配置的参数以及它们的默认值。

| 参数                                        | 类型   | 默认值                                                       | 描述                                                         |
| ------------------------------------------- | ------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `commonLabels`                              | object | `{}`                                                         | 自定义的资源标签                                             |
| `replicaCount`                              | int    | `1`                                                          | 期望的 Pod 副本数量，如果副本数大于 1，则进行领导人选举。    |
| `image.repository`                          | string | `registry-vpc.__ACK_REGION_ID__.aliyuncs.com/acs/spark-operator` | 镜像名称                                                     |
| `image.pullPolicy`                          | string | `IfNotPresent`                                               | 镜像拉取策略                                                 |
| `image.tag`                                 | string | `v1beta2-1.3.8-3.1.1`                                        | 镜像标签                                                     |
| `imagePullSecrets`                          | list   | `[]`                                                         | 镜像拉取秘钥                                                 |
| `nameOverride`                              | string | `""`                                                         | 用于部分重写 `spark-operator.fullname` 模版                  |
| `fullnameOverride`                          | string | `ack-spark-operator3`                                        | 用于重写发布名称                                             |
| `rbac.create`                               | bool   | `false`                                                      | **已废弃**，请使用 `createRole` 和 `createClusterRole`       |
| `rbac.createRole`                           | bool   | `true`                                                       | 创建并使用 RBAC `Role` 资源                                  |
| `rbac.createClusterRole`                    | bool   | `true`                                                       | 创建并使用 RBAC `ClusterRole` 资源                           |
| `serviceAccounts.spark.create`              | bool   | `true`                                                       | 是否为 spark 应用创建一个 service account                    |
| `serviceAccounts.spark.name`                | string | `""`                                                         | 可选，指定 spark 应用使用的 service account 名称             |
| `serviceAccounts.spark.annotations`         | object | `{}`                                                         | 可选，为 spark 应用使用的 service account 添加额外注解       |
| `serviceAccounts.sparkoperator.create`      | bool   | `true`                                                       | 是否为operator 创建一个 service account                      |
| `serviceAccounts.sparkoperator.name`        | string | `""`                                                         | 可选，指定 operator 使用的 service account 名称              |
| `serviceAccounts.sparkoperator.annotations` | object | `{}`                                                         | 可选，为 operator 使用的 service account 添加额外注解        |
| `sparkJobNamespace`                         | string | `""`                                                         | 可选， 如果需要在不同于 operator 所在的命名空间中运行 Spark 作业，可以修改此项 |
| `controllerThreads`                         | int    | `10`                                                         | Operator 并发度，内存使用量会随着并发度增加而增加            |
| `resyncInterval`                            | int    | `30`                                                         | Operator resync 时间间隔                                     |
| `uiService.enable`                          | bool   | `true`                                                       | 是否为 Spark Web UI 创建 Service 资源                        |
| `ingressUrlFormat`                          | string | `""`                                                         | Ingress URL 格式                                             |
| `logLevel`                                  | int    | `2`                                                          | 日志级别，级别越高日志越详细                                 |
| `podSecurityContext`                        | object | `{}`                                                         | Pod 安全上下文                                               |
| `securityContext`                           | object | `{}`                                                         | Operator 容器安全上下文                                      |
| `sidecars`                                  | list   | `[]`                                                         | Sidecar 容器                                                 |
| `volumes`                                   | list   | `[]`                                                         | Operator 数据卷                                              |
| `volumeMounts`                              | list   | `[]`                                                         | Operator 数据卷挂载                                          |
| `webhook.enable`                            | bool   | `false`                                                      | 启用 webhook 服务器                                          |
| `webhook.port`                              | int    | `8080`                                                       | webhook 服务端口号                                           |
| `webhook.namespaceSelector`                 | string | `""`                                                         | webhook 服务器监听的命名空间选择器，格式为 `key1=value1,key2=value2`，空字符串表示监听所有命名空间 |
| `webhook.initAnnotations`                   | object | `{"helm.sh/hook":"pre-install, pre-upgrade","helm.sh/hook-weight":"50"}` | 应用于 webhook 初始化 job 的注解信息，Helm 生命周期钩子会根据注解信息触发 webhook 初始化作业 |
| `webhook.cleanupAnnotations`                | object | `{"helm.sh/hook":"pre-delete, pre-upgrade","helm.sh/hook-delete-policy":"hook-succeeded"}` | 应用于 webhook 清理 job 的注解信息，Helm 生命周期钩子会根据注解信息触发 webhook 清理作业 |
| `webhook.timeout`                           | int    | `30`                                                         | webhook 超时时间，单位为秒                                   |
| `metrics.enable`                            | bool   | `true`                                                       | 是否启用 Prometheus 监控指标抓取                             |
| `metrics.port`                              | int    | `10254`                                                      | 监控指标端口号                                               |
| `metrics.portName`                          | string | `metrics`                                                    | 监控指标端口名                                               |
| `metrics.endpoint`                          | string | `/metrics`                                                   | 监控指标服务端点                                             |
| `metrics.prefix`                            | string | `""`                                                         | 监控指标前缀                                                 |
| `podMonitor.enable`                         | bool   | `false`                                                      | 是否会 operator pod 创建 pod monitor，需要同时设置 `metrics.enable` 为 `true` |
| `podMonitor.labels`                         | object | `{}`                                                         | Pod monitor 的标签信息                                       |
| `podMonitor.jobLabel`                       | string | `spark-operator-podmonitor`                                  | 从标签中获取作业名称                                         |
| `podMonitor.podMetricsEndpoint`             | object | `{"interval":"5s","scheme":"http"}`                          | Prometheus 指标端点属性，将会使用 `metrics.portName` 作为端口 |
| `nodeSelector`                              | object | `{}`                                                         | Pod 调度时需要匹配的节点选择器                               |
| `tolerations`                               | list   | `[]`                                                         | Pod 容忍度                                                   |
| `affinity`                                  | object | `{}`                                                         | Pod 亲和性                                                   |
| `podAnnotations`                            | object | `{}`                                                         | Pod 注解信息                                                 |
| `podLabels`                                 | object | `{}`                                                         | Pod 标签信息                                                 |
| `resources`                                 | object | `{}`                                                         | Pod 资源 requests 和 limits，每次提交 Spark 作业时都会执行 `/usr/local/openjdk-11/bin/java -Xmx128m` 命令创建一个 JVM，如果资源不足，Kubernetes 可能会杀死这些 Java 进程，因此如果看到如下报错信息： `"failed to run spark-submit for SparkApplication [...]: signal: killed"`，你需要增大内存的 limits。 |
| `batchScheduler.enable`                     | bool   | `false`                                                      | 是否启用 batch 调度                                          |
| `resourceQuotaEnforcement.enable`           | bool   | `false`                                                      | 是否启用 ResourceQuota 特性                                  |
| `leaderElection.lockName`                   | string | `spark-operator-lock`                                        | 领导人选举中使用的锁的名称                                   |
| `leaderElection.lockNamespace`              | string | `""`                                                         | 领导人选举中使用的锁所在的命名空间                           |
| `istio.enabled`                             | bool   | `false`                                                      | 是否启用 Istio                                               |
| `labelSelectorFilter`                       | string | ""                                                           | 基于指定的标签过滤需要 watch 和 list 的资源，格式为 `key1=value1,key2=value2,key3,key4` |

## 发布说明

| 版本号    | 变更时间       | 变更内容                                                     |
| --------- | -------------- | ------------------------------------------------------------ |
| `v1.1.26` | 2023年03月14日 | （1）新增 ack-spark-operator3.0  chart                       |
| `v1.1.27` | 2024年01月04日 | （1）更新 webhook init 和 cleanup job 模版（2）新增配置项 `nameOverride` 和 `fullnameOverride` （3）更新开启高可用时需要的 RBAC 权限（4）新增中文 README 文档 |
