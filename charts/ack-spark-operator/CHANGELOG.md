# CHANGELOG

## 0.1

### 0.1.17

- 新增配置项 `ingressUrlFormat`

### 0.1.18

- RBAC 权限收敛，为 Spark operator 使用的 ServiceAccount 创建如下 RBAC 资源：
  - 在 Spark operator 所在的 namespace 中创建一个 ClusterRole，用于赋予相关 cluster scope 的资源权限
  - 在 Spark operator 所在的 namespace 中创建一个 Role，赋予 Secret 资源相关权限
  - 在 `spark.jobNamespace` 中创建一个 Role，赋予 operator 在该 namespace 中运行 Spark 作业需要的相关权限
- 模版和配置项变更
  - 新增 `region` 配置项
  - 镜像相关配置项移动至 `image` 字段下，包括 `image.registry`、`image.pullByVPC`、`image.tag`、`image.pullPolicy`、`image.pullSecrets`
  - Spark Operator RBAC 相关配置项移动至 `sparkoperator` 字段下，包括 `sparkoperator.serviceAccount.create`、`sparkoperator.serviceAccount.name`、`sparkoperator.rbac.create`
  - Spark 应用 RBAC 相关配置项移动至 `spark` 字段下，包括 `spark.jobNamespace`、`spark.serviceAccount.name`、`spark.rbac.create`
  - webhook 相关配置项移动至 `webhook` 字段下，包括 `webhook.enable`、`webhook.port`
  - Prometheus 监控相关配置项移动至 `metrics` 字段下，包括 `metrics.enable`、`metrics.port`、`metrics.portName`、`metrics.prefix`
  - 新增 `logLevel` 配置项用于控制日志级别
  - 移除 `installCrds` 配置项并移除 `templates/crds.yaml` 模版文件
- 新增中文 README 文件

### 1.1.28

使用 v0.1.18 版本覆盖旧的 v1.1.28 版本

### 1.1.29

- 支持在 SparkApplication 中添加 scheduling.x-k8s.io/suspend: "true" 注解实现 suspend 操作，从而允许 kube-queue 进行排队处理
- 镜像仓库迁移至 ACR EE

### 1.1.30

- 增加 spark-operator 对 configmaps、services、events、resourcequotas 等资源的 RBAC 权限
