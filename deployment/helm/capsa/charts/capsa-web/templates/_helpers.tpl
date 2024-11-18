{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "capsa-web.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | quote }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "capsa-web.labels" -}}
helm.sh/chart: {{ include "capsa-web.chart" . }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
