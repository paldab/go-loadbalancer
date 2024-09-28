{{- define "paldab-loadbalancer.name" -}}
{{-  .Chart.Name | trunc 63 | trimSuffix "-" }}
{{- end }}

{{- define "paldab-loadbalancer.fullname" -}}
{{- $name := .Chart.Name }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }} 

{{- define "paldab-loadbalancer.namespace" -}}
{{ print .Release.Namespace }}
{{- end }}

{{- define "paldab-loadbalancer.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "paldab-loadbalancer.labels" -}}
helm.sh/chart: {{ include "paldab-loadbalancer.chart" . }}
{{ include "paldab-loadbalancer.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "paldab-loadbalancer.selectorLabels" -}}
app.kubernetes.io/name: {{ include "paldab-loadbalancer.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

# Service account
{{- define "paldab-loadbalancer.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- print "paldab-loadbalancer" }}
{{- else }}
{{- print "default" }}
{{- end }}
{{- end }}

{{- define "paldab-loadbalancer.saLabels" -}}
helm.sh/chart: {{ include "paldab-loadbalancer.chart" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

# Rbac
{{- define "paldab-loadbalancer.serverMonitorRoleName" -}}
{{- print "loadbalancer-server-monitor" -}}
{{- end }}

{{- define "paldab-loadbalancer.serverMonitorRolebindingName" -}}
{{- print "loadbalancer-server-monitor-binding" -}}
{{- end }}

# Crds
{{- define "paldab-loadbalancer.crdGroup" -}}
{{ print "loadbalancer.paldab.io" }}
{{- end }}

{{- define "paldab-loadbalancer.serverCrdName" -}}
{{- printf "%s.%s" ("servers") (include "paldab-loadbalancer.crdGroup" . )}}
{{- end }}
