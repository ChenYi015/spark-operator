# CHANGELOG

## 1.1

### 1.1.26

- 同步社区版 [spark-on-k8s-operator](https://github.com/GoogleCloudPlatform/spark-on-k8s-operator) v1.1.26

### 1.1.27

- 更新 webhook init job 和 cleanup job 的模版，解决不启用 webhook 但配置了 `tolerations` 和 `nodeSelector` 字段时部署失败的问题
- 增加 `nameOverride` 和 `fullnameOverride` 配置项，并为 `fullnameOverride` 字段提供默认参数 `ack-spark-operator3` 以解决 webhook Service 名称中包含 `.` 从而导致部署失败的问题
- 更新开启高可用时需要的 RBAC 权限
- 新增中文 README 文档

### 1.1.28

- 支持为 webhook 初始化和清理作业 pod 添加额外标签
- Spark master URL 支持 IPv6（Spark 从 v3.4 开始支持 IPv6）
- SparkApplication CRD 新增 `seccompProfile` 字段， 参见 https://kubernetes.io/docs/concepts/security/pod-security-standards/#restricted
- 支持为 RBAC 资源添加额外注解信息
- 修复 [SparkApplication with volcano scheduler not working on OCP #1723](https://github.com/GoogleCloudPlatform/spark-on-k8s-operator/issues/1723)
- 镜像仓库迁移至 ACR EE

### 1.1.29

- 支持在 SparkApplication 中添加 `scheduling.x-k8s.io/suspend: "true"` 注解实现 suspend 操作，从而允许 kube-queue 进行排队处理
