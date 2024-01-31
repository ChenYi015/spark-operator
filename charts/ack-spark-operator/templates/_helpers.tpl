{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "spark-operator.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "spark-operator.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "spark-operator.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Common labels
*/}}
{{- define "spark-operator.labels" -}}
helm.sh/chart: {{ include "spark-operator.chart" . }}
{{ include "spark-operator.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- if .Values.commonLabels }}
{{ toYaml .Values.commonLabels }}
{{- end }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "spark-operator.selectorLabels" -}}
app.kubernetes.io/name: {{ include "spark-operator.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to be used by spark operator
*/}}
{{- define "spark-operator.serviceAccountName" -}}
{{- if .Values.sparkoperator.serviceAccount.create -}}
{{ default (include "spark-operator.fullname" .) .Values.sparkoperator.serviceAccount.name }}
{{- else -}}
{{ default "default" .Values.serviceAccounts.sparkoperator.name }}
{{- end -}}
{{- end -}}

{{/*
Create the name of the image to be used by spark operator
*/}}
{{- define "spark-operator.imageName" -}}
{{- if .Values.image.registry -}}
{{ printf "%s/%s:%s" .Values.image.registry .Values.image.repository .Values.image.tag }}
{{- else if .Values.image.pullByVPC -}}
{{ printf "registry-vpc.%s.aliyuncs.com/%s:%s" .Values.region .Values.image.repository .Values.image.tag }}
{{- else -}}
{{ printf "registry.%s.aliyuncs.com/%s:%s" .Values.region .Values.image.repository .Values.image.tag }}
{{- end -}}
{{- end -}}

{{/*
Create the name of the webhook
*/}}
{{- define "spark-operator.webhookName" -}}
{{ include "spark-operator.fullname" . }}-webhook
{{- end -}}

{{/*
Create the name of the secret to be used by webhook
*/}}
{{- define "spark-operator.webhookSecretName" -}}
"spark-webhook-certs"
{{- end -}}

{{/*
Create the name of the service account to be used by spark applications
*/}}
{{- define "spark.serviceAccountName" -}}
{{- if .Values.spark.serviceAccount.create -}}
{{- $name := printf "%s-%s" (include "spark-operator.fullname" .) "spark" -}}
{{- default $name .Values.spark.serviceAccount.name -}}
{{- else -}}
{{- default "default" .Values.spark.serviceAccount.name -}}
{{- end -}}
{{- end -}}
