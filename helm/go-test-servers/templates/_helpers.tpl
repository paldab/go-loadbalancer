{{- define "go-test-servers.image" -}}
{{- print "paer12/go-test-server:0.1.0" }}
{{- end }}

{{- define "go-test-servers.replicas" -}}
{{- print "2" | int }}
{{- end }}
