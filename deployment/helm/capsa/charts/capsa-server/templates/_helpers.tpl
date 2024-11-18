{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "capsa-server.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | quote }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "capsa-server.labels" -}}
helm.sh/chart: {{ include "capsa-server.chart" . }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
