// Code generated by rice embed-go; DO NOT EDIT.
package main

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "analytics-k8s.tmpl",
		FileModTime: time.Unix(1571805848, 0),

		Content: string("{{$subsys := .SubsysName}}\n\n{{$pathsep := .OsEnv.PathSeparator}}\n{{$apicup := .OsEnv.BinApicup}}\n\n{{ template \"scriptheader\" .OsEnv }}\n\n{{$apicup}} subsys create {{$subsys}} analytics --k8s\n\n{{ if len .ExtraValuesFile -}}\n{{$apicup}} subsys set {{$subsys}} extra-values-file=..{{$pathsep}}{{ .ExtraValuesFile }}\n{{ end -}}\n\n{{$apicup}} subsys set {{$subsys}} analytics-ingestion={{.AnalyticsIngestionEndpoint}}\n{{$apicup}} subsys set {{$subsys}} analytics-client={{.AnalyticsClientEndpoint}}\n{{$apicup}} subsys set {{$subsys}} namespace={{.Namespace}}\n{{$apicup}} subsys set {{$subsys}} registry={{.RegistryUrl}}\n{{$apicup}} subsys set {{$subsys}} registry-secret={{.RegistrySecret}}\n\n{{/* default 6, recommended minimum 12 */}}\n{{$apicup}} subsys set {{$subsys}} coordinating-max-memory-gb={{.CoordinatingMaxMemoryGb | default 12}}\n{{/* default 6, recommended minimum 12 */}}\n{{$apicup}} subsys set {{$subsys}} data-max-memory-gb={{.DataMaxMemoryGb | default 12}}\n{{/* default 200, minimum 200, this is per pod, 3 pods: 600gb */}}\n{{$apicup}} subsys set {{$subsys}} data-storage-size-gb={{.DataStorageSizeGb | default 200}}\n{{/* miminum 12 */}}\n{{$apicup}} subsys set {{$subsys}} master-max-memory-gb={{.MasterMaxMemoryGb | default 12}}\n{{/* minimum 5 */}}\n{{$apicup}} subsys set {{$subsys}} master-storage-size-gb={{.MasterStorageSizeGb | default 5}}\n\n{{$apicup}} subsys set {{$subsys}} enable-message-queue={{.EnableMessageQueue | default false}}\n\n{{$apicup}} subsys set {{$subsys}} storage-class={{.StorageClass}}\n{{ if len .EsStorageClass }}\n{{$apicup}} subsys set {{$subsys}} es-storage-class={{.EsStorageClass}}\n{{end}}\n{{ if len .MqStorageClass }}\n{{$apicup}} subsys set {{$subsys}} mq-storage-class={{.MqStorageClass}}\n{{end}}\n\n{{$apicup}} subsys set {{$subsys}} mode={{.Mode}}\n{{$apicup}} subsys set {{$subsys}} ingress-type={{.IngressType}}\n"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "analytics-vm.tmpl",
		FileModTime: time.Unix(1571805848, 0),

		Content: string("{{$subsys := .SubsysName}}\n\n{{$pathsep := .OsEnv.PathSeparator}}\n{{$apicup := .OsEnv.BinApicup}}\n{{$islinux := .OsEnv.IsLinux}}\n\n{{ template \"scriptheader\" .OsEnv }}\n\n{{ $apicup }} subsys create {{ $subsys }} analytics\n{{ $apicup }} subsys set {{ $subsys }} mode={{ .Mode }}\n\n{{ with .CloudInit }}\n{{ if len .CloudInitFile }}\n    {{ $apicup }} subsys set {{ $subsys }} additional-cloud-init-file=..{{$pathsep}}{{ .CloudInitFile }}\n{{ end }}\n{{- end }}\n\n{{ $apicup }} subsys set {{ $subsys }} analytics-ingestion={{ .AnalyticsIngestion }}\n{{ $apicup }} subsys set {{ $subsys }} analytics-client={{ .AnalyticsClient }}\n\n{{ $apicup }} subsys set {{ $subsys }} search-domain={{ join \",\" .SearchDomains | trim | quote }}\n\n{{ $apicup }} subsys set {{ $subsys }} ssh-keyfiles=..{{$pathsep}}{{ .SshPublicKeyFile }}\n\n{{ with .VmFirstBoot }}\n{{ $apicup }} subsys set {{ $subsys }} dns-servers={{ join \",\" .DnsServers | trim | quote }}\n\n{{ if $islinux }}\n    {{ $apicup }} subsys set {{ $subsys }} default-password={{ .VmwareConsolePasswordHash | squote }}\n{{ else}}\n    {{ $apicup }} subsys set {{ $subsys }} default-password={{ .VmwareConsolePasswordHash | quote }}\n{{end}}\n\n{{ with .IpRanges }}\n{{ if len .PodNetwork }}\n    {{ $apicup }} subsys set {{ $subsys }} k8s-pod-network={{ .PodNetwork | quote }}\n{{- end}}\n{{ if len .ServiceNetwork }}\n    {{ $apicup }} subsys set {{ $subsys }} k8s-service-network={{ .ServiceNetwork | quote }}\n{{- end}}\n{{- end}}\n\n{{- range .Hosts}}\n{{$h := .}}\n{{ $apicup }} hosts create {{$subsys}} {{$h.Name}} {{$h.HardDiskPassword}}\n{{ $apicup }} iface create {{$subsys}} {{$h.Name}} {{$h.Device}} {{$h.IpAddress}}/{{$h.SubnetMask}} {{$h.Gateway}}\n{{- end}}\n\n{{- end}}\n\n{{ $apicup }} subsys set {{ $subsys }} enable-message-queue={{ .EnableMessageQueue }}\n"),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    "cloud-init-vm.tmpl",
		FileModTime: time.Unix(1573371241, 0),

		Content: string("{{- $a := list 0 . }}\n{{- template \"map2yml\" $a }}\n"),
	}
	file5 := &embedded.EmbeddedFile{
		Filename:    "combined-csr.tmpl",
		FileModTime: time.Unix(1571805848, 0),

		Content: string("{{$pathsep := .OsEnv.PathSeparator}}\n{{$shellext := .OsEnv.ShellExt}}\n{{$scriptinvoke := .OsEnv.ScriptInvoke}}\n\n{{ template \"scriptheader1\" .OsEnv }}\n\n{{ range $k, $cs := .CertSpecs }}\n    {{$scriptinvoke}} .{{$pathsep}}{{$cs.CsrConf}}{{$shellext}}\n{{ end }}\n"),
	}
	file6 := &embedded.EmbeddedFile{
		Filename:    "csr-client-auth.tmpl",
		FileModTime: time.Unix(1571805848, 0),

		Content: string("[req]\n# default key length for rsa key\ndefault_bits = 2048\n\n# do not encrypt private key\nencrypt_key = no\nencrypt_rsa_key = no\n\n# default message digest alg for signing certs and cert reqs\ndefault_md = sha256\n\n# cert request extensions section\nreq_extensions = req_ext\n\n# self-signed cert extensions section\n#x509_extensions = self_signed_extensions\n\n# do not prompt for the dn\nprompt = no\n\n# section name for dn fields\ndistinguished_name = dn\n\n# make sure dn components match ca policy\n[dn]\n{{- range .DnFields | reverse }}\n    {{- nindent 0 . }}\n{{- end }}\nCN = {{ .Cn }}\n\n[req_ext]\nextendedKeyUsage = clientAuth\n"),
	}
	file7 := &embedded.EmbeddedFile{
		Filename:    "csr-server-auth.tmpl",
		FileModTime: time.Unix(1573284287, 0),

		Content: string("[req]\n# default key length for rsa key\ndefault_bits = 2048\n\n# do not encrypt private key\nencrypt_key = no\nencrypt_rsa_key = no\n\n# default message digest alg for signing certs and cert reqs\ndefault_md = sha256\n\n# cert request extensions section\nreq_extensions = req_ext\n\n# self-signed cert extensions section\n#x509_extensions = self_signed_extensions\n\n# do not prompt for the dn\nprompt = no\n\n# section name for dn fields\ndistinguished_name = dn\n\n# make sure dn components match ca policy\n[dn]\n{{- range .DnFields | reverse }}\n    {{- nindent 0 . }}\n{{- end }}\nCN = {{ .Cn }}\n\n[req_ext]\nextendedKeyUsage = serverAuth\n# update subject alt name\n{{- if len .AltCns }}\nsubjectAltName = DNS:{{.Cn}},DNS:{{ join \",DNS:\" .AltCns }}\n{{- else }}\nsubjectAltName = DNS:{{.Cn}}\n{{- end }}"),
	}
	file8 := &embedded.EmbeddedFile{
		Filename:    "extra-values.tmpl",
		FileModTime: time.Unix(1572973028, 0),

		Content: string("{{- $a := list 0 . }}\n{{- template \"map2yml\" $a }}"),
	}
	file9 := &embedded.EmbeddedFile{
		Filename:    "gateway-k8s.tmpl",
		FileModTime: time.Unix(1571805848, 0),

		Content: string("{{$subsys := .SubsysName}}\n\n{{$pathsep := .OsEnv.PathSeparator}}\n{{$apicup := .OsEnv.BinApicup}}\n\n{{ template \"scriptheader\" .OsEnv }}\n\n{{$apicup}} subsys create {{$subsys}} gateway --k8s\n\n{{if len .ExtraValuesFile}}\n{{$apicup}} subsys set {{$subsys}} extra-values-file=..{{$pathsep}}{{.ExtraValuesFile}}\n{{end}}\n\n{{$apicup}} subsys set {{$subsys}} api-gateway={{.ApiGateway}}\n{{$apicup}} subsys set {{$subsys}} apic-gw-service={{.ApicGwService}}\n\n{{$apicup}} subsys set {{$subsys}} namespace={{.Namespace}}\n\n{{if len .RegistryUrl}}\n{{$apicup}} subsys set {{$subsys}} registry={{.RegistryUrl}}\n{{end}}\n{{if len .RegistrySecret}}\n{{$apicup}} subsys set {{$subsys}} registry-secret={{.RegistrySecret}}\n{{end}}\n\n{{$apicup}} subsys set {{$subsys}} image-pull-policy={{.ImagePullPolicy | default \"IfNotPresent\"}}\n\n{{$apicup}} subsys set {{$subsys}} replica-count={{.ReplicaCount | default 3}}\n{{$apicup}} subsys set {{$subsys}} max-cpu={{.MaxCpu | default 4}}\n{{$apicup}} subsys set gwy max-memory-gb={{.MaxMemoryGb | default 6}}\n\n{{$apicup}} subsys set {{$subsys}} storage-class={{.StorageClass}}\n\n{{$apicup}} subsys set {{$subsys}} v5-compatibility-mode={{.V5CompatibilityMode | default false}}\n{{$apicup}} subsys set {{$subsys}} enable-tms={{.EnableTms}}\n{{$apicup}} subsys set {{$subsys}} tms-peering-storage-size-gb={{.TmsPeeringStorageSizeGb | default 10}}\n{{$apicup}} subsys set {{$subsys}} enable-high-performance-peering={{.EnableHighPerformancePeering | quote}}\n\n{{$apicup}} subsys set {{$subsys}} license-version={{.LicenseVersion}}\n{{$apicup}} subsys set {{$subsys}} mode={{.Mode}}\n{{$apicup}} subsys set {{$subsys}} ingress-type={{.IngressType | default \"ingress\"}}\n"),
	}
	filea := &embedded.EmbeddedFile{
		Filename:    "helpers.tmpl",
		FileModTime: time.Unix(1571805848, 0),

		Content: string("{{/* write out script header. pass OsEnv as scope */}}\n{{ define \"scriptheader\"}}\n    {{ if .IsLinux }}\n        #!/bin/bash\n        # run this script from project directory\n        set -x\n    {{else}}\n        rem run this script from project directory\n    {{end}}\n{{ end }}\n\n{{ define \"scriptheader1\"}}\n    {{ if .IsLinux }}\n        #!/bin/bash\n        # run this script from current directory\n        set -x\n    {{else}}\n        rem run this script from current directory\n    {{end}}\n{{ end }}\n\n{{- /* print out map as yml... pass list as scope: [depth, map] */}}\n{{- define \"map2yml\" }}\n    {{- $d := first . }}\n    {{- $m := last . }}\n    {{- range $k, $v := $m }}\n        {{- $vt := typeOf $v }}\n        {{- if eq $vt \"map[string]interface {}\" }}\n            {{- /*(key-depth -- {{$d}})*/}}\n            {{- $s := int $d}}\n            {{- nindent $s $k}}:\n            {{- $d1 := add1 $d}}\n            {{- $a := list $d1 $v}}\n            {{- template \"map2yml\" $a }}\n        {{- else}}\n            {{- /*(key-val-depth -- {{$d}})*/}}\n            {{- $s := int $d }}\n            {{- nindent $s $k }}: {{$v}}\n        {{- end}}\n    {{- end }}\n{{- end }}\n"),
	}
	fileb := &embedded.EmbeddedFile{
		Filename:    "keypair.tmpl",
		FileModTime: time.Unix(1571805848, 0),

		Content: string("{{$pathsep := .OsEnv.PathSeparator}}\n{{$shellext := .OsEnv.ShellExt}}\n\n{{ template \"scriptheader1\" .OsEnv }}\n\nopenssl req -config {{.CertSpec.CsrConf}} -out {{.CertSpec.CertFile}}.csr -outform PEM -new -keyout {{.CertSpec.KeyFile}}\n"),
	}
	filec := &embedded.EmbeddedFile{
		Filename:    "management-k8s.tmpl",
		FileModTime: time.Unix(1572486449, 0),

		Content: string("{{ $subsys := .SubsysName }}\n\n{{$pathsep := .OsEnv.PathSeparator}}\n{{$apicup := .OsEnv.BinApicup}}\n\n{{ template \"scriptheader\" .OsEnv }}\n\n{{$apicup}} subsys create {{ $subsys }} management --k8s\n\n{{$apicup}} subsys set {{ $subsys }} mode={{ .Mode }}\n\n{{$apicup}} subsys set {{ $subsys }} ingress-type={{ .IngressType }}\n{{$apicup}} subsys set {{ $subsys }} namespace={{ .Namespace }}\n{{$apicup}} subsys set {{ $subsys }} registry={{ .RegistryUrl }}\n{{$apicup}} subsys set {{ $subsys }} registry-secret={{ .RegistrySecret }}\n{{$apicup}} subsys set {{ $subsys }} storage-class={{ .StorageClass }}\n\n{{ if len .ExtraValuesFile -}}\n    {{$apicup}} subsys set {{ $subsys }} extra-values-file=..{{$pathsep}}{{ .ExtraValuesFile | quote }}\n{{ end -}}\n\n{{ with .CassandraBackup }}\n    {{ if .BackupProtocol | lower | eq \"sftp\" -}}\n        {{$apicup}} subsys set {{ $subsys }} cassandra-backup-protocol=sftp\n        {{$apicup}} subsys set {{ $subsys }} cassandra-backup-host={{ .BackupHost }}\n        {{$apicup}} subsys set {{ $subsys }} cassandra-backup-port={{ .BackupPort | default 22 }}\n        {{$apicup}} subsys set {{ $subsys }} cassandra-backup-auth-user={{ .BackupAuthUser }}\n        {{$apicup}} subsys set {{ $subsys }} cassandra-backup-auth-pass={{ .BackupAuthPass }}\n        {{$apicup}} subsys set {{ $subsys }} cassandra-backup-path={{ .BackupPath }}\n        {{$apicup}} subsys set {{ $subsys }} cassandra-backup-schedule={{ .BackupSchedule | quote }}\n    {{ else if .BackupProtocol | lower | eq \"objstore\" -}}\n        {{$apicup}} subsys set {{ $subsys }} cassandra-backup-protocol=objstore\n        {{$apicup}} subsys set {{ $subsys }} cassandra-backup-host={{ .ObjstoreEndpointRegion }}\n        {{$apicup}} subsys set {{ $subsys }} cassandra-backup-auth-user={{ .ObjstoreS3SecretKeyId }}\n        {{$apicup}} subsys set {{ $subsys }} cassandra-backup-auth-pass={{ .ObjstoreS3SecretAccessKey }}\n        {{$apicup}} subsys set {{ $subsys }} cassandra-backup-path={{ .ObjstoreBucketSubfolder }}\n        {{$apicup}} subsys set {{ $subsys }} cassandra-backup-schedule={{ .BackupSchedule | quote }}\n    {{- end }}\n{{- end }}\n\n{{$apicup}} subsys set {{ $subsys }} cassandra-max-memory-gb={{ .CassandraMaxMemoryGb | default 9 }}\n{{$apicup}} subsys set {{ $subsys }} cassandra-cluster-size={{ .CassandraClusterSize }}\n{{$apicup}} subsys set {{ $subsys }} cassandra-volume-size-gb={{ .CassandraVolumeSizeGb }}\n{{ if .ExternalCassandraHost }}\n    {{$apicup}} subsys set {{ $subsys }} external-cassandra-host={{ .ExternalCassandraHost }}\n{{ end }}\n\n{{$apicup}} subsys set {{ $subsys }} create-crd={{ .CreateCrd | default true }}\n\n{{$apicup}} subsys set {{ $subsys }} platform-api={{ .PlatformApi }}\n{{$apicup}} subsys set {{ $subsys }} api-manager-ui={{ .ApiManagerUi }}\n{{$apicup}} subsys set {{ $subsys }} cloud-admin-ui={{ .CloudAdminUi }}\n{{$apicup}} subsys set {{ $subsys }} consumer-api={{ .ConsumerApi }}\n\n{{ if .CassandraEncryptionKeyFile }}\n    {{$apicup}} certs set {{ $subsys }} encryption-secret ..{{$pathsep}}{{ .CassandraEncryptionKeyFile }}\n{{ end }}"),
	}
	filed := &embedded.EmbeddedFile{
		Filename:    "management-vm.tmpl",
		FileModTime: time.Unix(1572486425, 0),

		Content: string("{{- $subsys := .SubsysName}}\n\n{{$pathsep := .OsEnv.PathSeparator}}\n{{$apicup := .OsEnv.BinApicup}}\n{{$islinux := .OsEnv.IsLinux}}\n\n{{ template \"scriptheader\" .OsEnv }}\n\n{{$apicup}} subsys create {{ $subsys }} management\n\n{{$apicup}} subsys set {{ $subsys }} mode={{ .Mode }}\n\n{{ with .CloudInit }}\n    {{- if len .CloudInitFile }}\n        {{$apicup}} subsys set {{ $subsys }} additional-cloud-init-file=..{{$pathsep}}{{ .CloudInitFile }}\n    {{- end }}\n{{ end }}\n\n{{$apicup}} subsys set {{ $subsys }} search-domain={{ join \",\" .SearchDomains | trim | quote }}\n\n{{ with .VmFirstBoot }}\n    {{$apicup}} subsys set {{ $subsys }} dns-servers={{ join \",\" .DnsServers | trim | quote }}\n\n    {{- if $islinux }}\n        {{$apicup}} subsys set {{ $subsys }} default-password={{ .VmwareConsolePasswordHash | squote }}\n    {{- else }}\n        {{$apicup}} subsys set {{ $subsys }} default-password={{ .VmwareConsolePasswordHash | quote }}\n    {{- end }}\n\n    {{ with .IpRanges }}\n        {{- if len .PodNetwork }}\n            {{$apicup}} subsys set {{ $subsys }} k8s-pod-network={{ .PodNetwork | quote }}\n        {{- end}}\n        {{- if len .ServiceNetwork }}\n            {{$apicup}} subsys set {{ $subsys }} k8s-service-network={{ .ServiceNetwork | quote }}\n        {{- end }}\n    {{ end }}\n\n    {{ range .Hosts}}\n        {{- $h := .}}\n        {{$apicup}} hosts create {{$subsys}} {{$h.Name}} {{$h.HardDiskPassword}}\n        {{$apicup}} iface create {{$subsys}} {{$h.Name}} {{$h.Device}} {{$h.IpAddress}}/{{$h.SubnetMask}} {{$h.Gateway}}\n    {{ end}}\n\n{{ end}}\n\n{{$apicup}} subsys set {{ $subsys }} ssh-keyfiles=..{{$pathsep}}{{ .SshPublicKeyFile }}\n\n{{ with .CassandraBackup }}\n    {{- if .BackupProtocol | lower | eq \"sftp\" }}\n        {{$apicup}} subsys set {{ $subsys }} cassandra-backup-protocol=sftp\n        {{$apicup}} subsys set {{ $subsys }} cassandra-backup-host={{ .BackupHost }}\n        {{$apicup}} subsys set {{ $subsys }} cassandra-backup-port={{ .BackupPort | default 22 }}\n        {{$apicup}} subsys set {{ $subsys }} cassandra-backup-auth-user={{ .BackupAuthUser }}\n        {{$apicup}} subsys set {{ $subsys }} cassandra-backup-auth-pass={{ .BackupAuthPass }}\n        {{$apicup}} subsys set {{ $subsys }} cassandra-backup-path={{ .BackupPath }}\n        {{$apicup}} subsys set {{ $subsys }} cassandra-backup-schedule={{ .BackupSchedule | quote }}\n    {{- else if .BackupProtocol | lower | eq \"objstore\" }}\n        {{$apicup}} subsys set {{ $subsys }} cassandra-backup-protocol=objstore\n        {{$apicup}} subsys set {{ $subsys }} cassandra-backup-host={{ .ObjstoreEndpointRegion }}\n        {{$apicup}} subsys set {{ $subsys }} cassandra-backup-auth-user={{ .ObjstoreS3SecretKeyId }}\n        {{$apicup}} subsys set {{ $subsys }} cassandra-backup-auth-pass={{ .ObjstoreS3SecretAccessKey }}\n        {{$apicup}} subsys set {{ $subsys }} cassandra-backup-path={{ .ObjstoreBucketSubfolder }}\n        {{$apicup}} subsys set {{ $subsys }} cassandra-backup-schedule={{ .BackupSchedule | quote }}\n    {{- end }}\n{{ end }}\n\n{{$apicup}} subsys set {{ $subsys }} platform-api={{ .PlatformApi }}\n{{$apicup}} subsys set {{ $subsys }} api-manager-ui={{ .ApiManagerUi }}\n{{$apicup}} subsys set {{ $subsys }} cloud-admin-ui={{ .CloudAdminUi }}\n{{$apicup}} subsys set {{ $subsys }} consumer-api={{ .ConsumerApi }}\n\n{{ if .CassandraEncryptionKeyFile }}\n    {{$apicup}} certs set {{ $subsys }} encryption-secret ..{{$pathsep}}{{ .CassandraEncryptionKeyFile }}\n{{ end }}"),
	}
	filee := &embedded.EmbeddedFile{
		Filename:    "portal-k8s.tmpl",
		FileModTime: time.Unix(1571805848, 0),

		Content: string("{{ $subsys := .SubsysName}}\n\n{{$pathsep := .OsEnv.PathSeparator}}\n{{$apicup := .OsEnv.BinApicup}}\n\n{{ template \"scriptheader\" .OsEnv }}\n\n{{$apicup}} subsys create {{$subsys}} portal --k8s\n\n{{$apicup}} subsys set {{$subsys}} mode={{.Mode}}\n{{$apicup}} subsys set {{$subsys}} extra-values-file=..{{$pathsep}}{{.ExtraValuesFile | quote}}\n\n{{$apicup}} subsys set {{$subsys}} ingress-type={{.IngressType}}\n{{$apicup}} subsys set {{$subsys}} namespace={{ .Namespace }}\n{{$apicup}} subsys set {{$subsys}} storage-class={{ .StorageClass }}\n\n{{$apicup}} subsys set {{$subsys}} registry={{ .RegistryUrl }}\n{{$apicup}} subsys set {{$subsys}} registry-secret={{ .RegistrySecret }}\n\n{{$apicup}} subsys set {{$subsys}} portal-admin={{ .PortalAdmin }}\n{{$apicup}} subsys set {{$subsys}} portal-www={{ .PortalWWW }}\n\n{{$apicup}} subsys set {{$subsys}} www-storage-size-gb={{ .WwwStorageSizeGb | default 5 }}\n{{$apicup}} subsys set {{$subsys}} backup-storage-size-gb={{ .BackupStorageSizeGb | default 5 }}\n{{$apicup}} subsys set {{$subsys}} db-storage-size-gb={{ .DbStorageSizeGb | default 12 }}\n{{$apicup}} subsys set {{$subsys}} db-logs-storage-size-gb={{ .DbLogsStorageSizeGb | default 12 }}\n{{/*apicup subsys set {{$subsys}} admin-storage-size-gb={{ .AdminStorageSizeGb | default 1 }}*/}}\n\n{{ with .SiteBackup }}\n{{ if .BackupProtocol | lower | eq \"sftp\" }}\n{{$apicup}} subsys set {{$subsys}} site-backup-protocol=sftp\n{{$apicup}} subsys set {{$subsys}} site-backup-host={{.BackupHost}}\n{{$apicup}} subsys set {{$subsys}} site-backup-port={{.BackupPort | default 22}}\n{{$apicup}} subsys set {{$subsys}} site-backup-auth-user={{.BackupAuthUser}}\n{{$apicup}} subsys set {{$subsys}} site-backup-auth-pass={{.BackupAuthPass}}\n{{$apicup}} subsys set {{$subsys}} site-backup-path={{.BackupPath }}\n{{$apicup}} subsys set {{$subsys}} site-backup-schedule={{.BackupSchedule | quote}}\n{{ else if .BackupProtocol | lower | eq \"objstore\" }}\n{{$apicup}} subsys set {{$subsys}} site-backup-protocol=objstore\n{{$apicup}} subsys set {{$subsys}} site-backup-host={{ .ObjstoreEndpointRegion }}\n{{$apicup}} subsys set {{$subsys}} site-backup-auth-user={{ .ObjstoreS3SecretKeyId }}\n{{$apicup}} subsys set {{$subsys}} site-backup-auth-pass={{ .ObjstoreS3SecretAccessKey }}\n{{$apicup}} subsys set {{$subsys}} site-backup-path={{ .ObjstoreBucketSubfolder }}\n{{$apicup}} subsys set {{$subsys}} site-backup-schedule={{ .BackupSchedule | quote }}\n{{ end }}\n{{ end }}"),
	}
	filef := &embedded.EmbeddedFile{
		Filename:    "portal-vm.tmpl",
		FileModTime: time.Unix(1571805848, 0),

		Content: string("{{$subsys := .SubsysName}}\n\n{{$pathsep := .OsEnv.PathSeparator}}\n{{$apicup := .OsEnv.BinApicup}}\n{{$islinux := .OsEnv.IsLinux}}\n\n{{ template \"scriptheader\" .OsEnv }}\n\n{{$apicup}} subsys create {{ $subsys }} portal\n{{$apicup}} subsys set {{ $subsys }} mode={{ .Mode }}\n\n{{ with .CloudInit }}\n{{ if len .CloudInitFile }}\n{{$apicup}} subsys set {{ $subsys }} additional-cloud-init-file=..{{$pathsep}}{{ .CloudInitFile }}\n{{ end }}\n{{ end }}\n\n{{$apicup}} subsys set {{ $subsys }} search-domain={{ join \",\" .SearchDomains | trim | quote }}\n\n{{$apicup}} subsys set {{ $subsys }} ssh-keyfiles=..{{$pathsep}}{{ .SshPublicKeyFile }}\n\n{{ with .VmFirstBoot }}\n{{$apicup}} subsys set {{ $subsys }} dns-servers={{ join \",\" .DnsServers | trim | quote }}\n\n{{ if $islinux }}\n{{$apicup}} subsys set {{ $subsys }} default-password={{ .VmwareConsolePasswordHash | squote }}\n{{ else}}\n{{$apicup}} subsys set {{ $subsys }} default-password={{ .VmwareConsolePasswordHash | quote }}\n{{end}}\n\n{{ with .IpRanges }}\n{{ if len .PodNetwork }}\n{{$apicup}} subsys set {{ $subsys }} k8s-pod-network={{ .PodNetwork | quote }}\n{{ end}}\n{{ if len .ServiceNetwork }}\n{{$apicup}} subsys set {{ $subsys }} k8s-service-network={{ .ServiceNetwork | quote }}\n{{ end}}\n{{ end}}\n\n{{range .Hosts}}\n{{$h := .}}\n{{$apicup}} hosts create {{$subsys}} {{$h.Name}} {{$h.HardDiskPassword}}\n{{$apicup}} iface create {{$subsys}} {{$h.Name}} {{$h.Device}} {{$h.IpAddress}}/{{$h.SubnetMask}} {{$h.Gateway}}\n{{ end}}\n\n{{ end}}\n\n{{ with .SiteBackup }}\n    {{ if .BackupProtocol | lower | eq \"sftp\" }}\n        {{$apicup}} subsys set {{ $subsys }} site-backup-protocol=sftp\n        {{$apicup}} subsys set {{ $subsys }} site-backup-host={{ .BackupHost }}\n        {{$apicup}} subsys set {{ $subsys }} site-backup-port={{ .BackupPort | default 22 }}\n        {{$apicup}} subsys set {{ $subsys }} site-backup-auth-user={{ .BackupAuthUser }}\n        {{$apicup}} subsys set {{ $subsys }} site-backup-auth-pass={{ .BackupAuthPass }}\n        {{$apicup}} subsys set {{ $subsys }} site-backup-path={{ .BackupPath }}\n        {{$apicup}} subsys set {{ $subsys }} site-backup-schedule={{ .BackupSchedule | quote }}\n    {{ else if .BackupProtocol | lower | eq \"objstore\" }}\n        {{$apicup}} subsys set {{ $subsys }} site-backup-protocol=objstore\n        {{$apicup}} subsys set {{ $subsys }} site-backup-host={{ .ObjstoreEndpointRegion }}\n        {{$apicup}} subsys set {{ $subsys }} site-backup-auth-user={{ .ObjstoreS3SecretKeyId }}\n        {{$apicup}} subsys set {{ $subsys }} site-backup-auth-pass={{ .ObjstoreS3SecretAccessKey }}\n        {{$apicup}} subsys set {{ $subsys }} site-backup-path={{ .ObjstoreBucketSubfolder }}\n        {{$apicup}} subsys set {{ $subsys }} site-backup-schedule={{ .BackupSchedule | quote }}\n    {{ end }}\n{{ end }}\n\n{{$apicup}} subsys set {{ $subsys }} portal-admin={{ .PortalAdmin }}\n{{$apicup}} subsys set {{ $subsys }} portal-www={{ .PortalWww }}\n"),
	}
	fileg := &embedded.EmbeddedFile{
		Filename:    "subsys-certs.tmpl",
		FileModTime: time.Unix(1573365115, 0),

		Content: string("{{$pathsep := .OsEnv.PathSeparator}}\n{{$apicup := .OsEnv.BinApicup}}\n\n{{ template \"scriptheader\" .OsEnv }}\n\n{{ range .CertSpecs}}\n    {{ $certSpec := .}}\n    {{$apicup}} certs set {{ $certSpec.SubsysName}} {{$certSpec.CertName}} ..{{$pathsep}}{{$certSpec.CertSubdir}}{{$pathsep}}{{$certSpec.CertFile}} ..{{$pathsep}}{{$certSpec.KeySubdir}}{{$pathsep}}{{$certSpec.KeyFile}} ..{{$pathsep}}{{$certSpec.CaSubdir}}{{$pathsep}}{{$certSpec.CaFile}}\n{{ end }}\n"),
	}
	fileh := &embedded.EmbeddedFile{
		Filename:    "subsys-config-k8s.tmpl",
		FileModTime: time.Unix(1572486003, 0),

		Content: string("{\n    \"InstallType\": \"k8s\",\n    \"Version\": {{ .Version | default \"2018.1.x\" | quote }},\n    \"Tag\": \"tag\",\n\n    \"Mode\": {{ .Mode | default \"dev|standard\" | quote }},\n\n    \"Namespace\": {{ .Namespace | default \"apic\" | quote }},\n    \"RegistryUrl\": \"container-image-registry-url\",\n    \"RegistrySecret\": \"container-image-registry-secret\",\n    \"IngressType\": \"ingress|route\",\n    \"StorageClass\": \"gp2|etc\",\n\n    \"Certs\": {\n        \"DnFields\": [\"O=APIC|match ca reqs\",\"C=US|match ca reqs\"],\n        \"K8sNamespace\": {{ .Namespace | default \"apic\" | quote }},\n        \"CaFile\": \"ca-chain-root-last.crt\",\n\n        \"Certbot\": {\n            \"CertDir\": \"letsencrypt/live/my.domain.com\",\n            \"Cert\": \"cert.pem\",\n            \"Key\": \"privkey.pem\",\n            \"CaChain\": \"chain.pem\"\n        },\n\n        \"PublicUserFacingCerts\": true,\n        \"PublicCerts\": false,\n        \"CommonCerts\": false\n    },\n\n    \"Management\": {\n        \"SubsysName\": \"mgmt\",\n        \"ExtraValuesFile\": \"mgmt-values.yaml\",\n        \"ExtraValues\": {},\n\n        \"CassandraBackup\": {\n            \"BackupProtocol\": \"sftp|objstore\",\n            \"BackupAuthUser\": \"admin\",\n            \"BackupAuthPass\": \"secret\",\n            \"BackupHost\": \"backup.my.domain.com\",\n            \"BackupPort\": 1022,\n            \"BackupPath\": \"/backup\",\n            \"ObjstoreS3SecretKeyId\": \"\",\n            \"ObjstoreS3SecretAccessKey\": \"\",\n            \"ObjstoreEndpointRegion\": \"\",\n            \"ObjstoreBucketSubfolder\": \"\",\n            \"BackupEncoding\": \"min(0-59) hour(0-23) dayofmonth(1-31) month(1-12) dayofweek(0-6)\",\n            \"BackupSchedule\": \"0 0 * * 0\"\n        },\n\n        \"CassandraMaxMemoryGb\": 9,\n        \"CassandraVolumeSizeGb\": 50,\n        \"CassandraClusterSize\": 3,\n        \"ExternalCassandraHost\": \"ext.my.domain.com\",\n        \"CreateCrd\": true,\n\n        \"CassandraEncryptionKeyFile\": \"encryption-secret.bin\",\n\n        \"PlatformApi\": \"api.my.domain.com\",\n        \"ApiManagerUi\": \"apim.my.domain.com\",\n        \"CloudAdminUi\": \"cm.my.domain.com\",\n        \"ConsumerApi\": \"consumer.my.domain.com\"\n    },\n\n    \"Analytics\": {\n        \"SubsysName\": \"analyt\",\n        \"ExtraValuesFile\": \"analyt-values.yaml\",\n        \"ExtraValues\": {},\n\n        \"CoordinatingMaxMemoryGb\": 12,\n        \"DataMaxMemoryGb\": 12,\n        \"DataStorageSizeGb\": 200,\n        \"MasterMaxMemoryGb\": 12,\n        \"MasterStorageSizeGb\": 5,\n\n        \"EnableMessageQueue\": false,\n\n        \"EsStorageClass\": \"\",\n        \"MqStorageClass\": \"\",\n\n        \"AnalyticsIngestionEndpoint\": \"ai.my.domain.com\",\n        \"AnalyticsClientEndpoint\": \"ac.my.domain.com\"\n    },\n\n    \"Portal\": {\n        \"SubsysName\": \"ptl\",\n\n        \"ExtraValuesFile\": \"ptl-values.yaml\",\n        \"ExtraValues\": {},\n\n        \"SiteBackup\": {\n            \"BackupProtocol\": \"sftp|objstore\",\n            \"BackupAuthUser\": \"admin\",\n            \"BackupAuthPass\": \"secret\",\n            \"BackupHost\": \"backup.my.domain.com\",\n            \"BackupPort\": 1022,\n            \"BackupPath\": \"/backup\",\n            \"ObjstoreS3SecretKeyId\": \"\",\n            \"ObjstoreS3SecretAccessKey\": \"\",\n            \"ObjstoreEndpointRegion\": \"\",\n            \"ObjstoreBucketSubfolder\": \"\",\n            \"BackupEncoding\": \"min(0-59) hour(0-23) dayofmonth(1-31) month(1-12) dayofweek(0-6)\",\n            \"BackupSchedule\": \"0 0 * * 0\"\n        },\n\n        \"WwwStorageSizeGb\": 5,\n        \"BackupStorageSizeGb\": 5,\n        \"DbStorageSizeGb\": 12,\n\n        \"Fixed\": {\n            \"DbLogsStorageSizeGb\": 2,\n            \"AdminStorageSizeGb\": 1\n        },\n\n        \"PortalAdmin\": \"padmin.my.domain.com\",\n        \"PortalWWW\": \"portal.my.domain.com\"\n    },\n\n    \"Gateway\": {\n        \"SubsysName\": \"gwy\",\n        \"Mode\": \"dev\",\n\n        \"ExtraValuesFile\": \"gwy-values.yaml\",\n        \"ExtraValues\": {\n            \"datapower\": {\n                \"webGuiManagementState\": \"enabled\",\n                \"apiDebugProbe\": \"enabled\"\n            }\n        },\n\n        \"LicenseVersion\": \"Production|Development\",\n        \"ImagePullPolicy\": \"IfNotPresent\",\n\n        \"ReplicaCount\": 3,\n        \"MaxCpu\": 4,\n        \"MaxMemoryGb\": 6,\n\n        \"V5ComatabilityMode\": false,\n        \"EnableTms\": true,\n        \"TmsPeeringStorageSizeGb\": 10,\n        \"EnableHighPerformancePeering\": \"true\",\n\n        \"ApiGateway\": \"gw.my.domain.com\",\n        \"ApicGwService\": \"gwd.my.domain.com\"\n    }\n}"),
	}
	filei := &embedded.EmbeddedFile{
		Filename:    "subsys-config-ova.tmpl",
		FileModTime: time.Unix(1572485996, 0),

		Content: string("{\n    \"InstallType\": \"ova\",\n    \"Version\": {{ .Version | default \"2018.1.x\" | quote }},\n    \"Tag\": \"tag\",\n\n    \"Mode\": {{ .Mode | default \"dev|standard\" | quote }},\n\n    \"SshPublicKeyFile\": \"/path/to/public/key/file\",\n\n    \"SearchDomains\": [\"my.domain.com\", \"domain.com\"],\n\n    \"VmFirstBoot\": {\n        \"DnsServers\": [\"dns-ip-1\",\"dns-ip-2\"],\n        \"VmwareConsolePasswordHash\": \"hash-output-b64\",\n\n        \"IpRanges\": {\n            \"PodNetwork\": \"172.16.0.0/16\",\n            \"ServiceNetwork\": \"172.17.0.0/16\"\n        }\n    },\n\n    \"CloudInit\": {\n        \"CloudInitFile\": \"cloud-init-file.yml\",\n        \"InitValues\": {\n            \"rsyslog\": {\n                \"remotes\": {\n                    \"syslog_server1\": \"syslog-collector-ip-1:514|601\",\n                    \"syslog_server2\": \"syslog-collector-ip-2:514|601\"\n                }\n            }\n        }\n    },\n\n    \"Certs\": {\n        \"DnFields\": [\"O=APIC|match ca reqs\",\"C=US|match ca reqs\"],\n        \"K8sNamespace\": \"default\",\n        \"CaFile\": \"ca-chain-root-last.crt\",\n\n        \"Certbot\": {\n            \"CertDir\": \"letsencrypt/live/my.domain.com\",\n            \"Cert\": \"cert.pem\",\n            \"Key\": \"privkey.pem\",\n            \"CaChain\": \"chain.pem\"\n        },\n\n        \"PublicUserFacingCerts\": true,\n        \"PublicCerts\": false,\n        \"CommonCerts\": false\n    },\n\n    \"Management\": {\n        \"SubsysName\": \"mgmt\",\n\n        \"VmFirstBoot\": {\n            \"Hosts\": [\n                {\"Name\": \"m1.my.domain.com\", \"HardDiskPassword\": \"password\", \"Device\": \"eth0\",\n                    \"IpAddress\": \"ip-address\", \"SubnetMask\": \"dot.subnet.mask\",\"Gateway\": \"gw-ip-address\"},\n                {\"Name\": \"m2.my.domain.com\", \"HardDiskPassword\": \"password\", \"Device\": \"eth0\",\n                    \"IpAddress\": \"ip-address\", \"SubnetMask\": \"dot.subnet.mask\", \"Gateway\": \"gw-ip-address\"},\n                {\"Name\": \"m3.my.domain.com\", \"HardDiskPassword\": \"password\", \"Device\": \"eth0\",\n                    \"IpAddress\": \"ip-address\", \"SubnetMask\": \"dot.subnet.mask\", \"Gateway\": \"gw-ip-address\"}\n            ]\n        },\n\n        \"CassandraBackup\": {\n            \"BackupProtocol\": \"sftp|objstore\",\n            \"BackupAuthUser\": \"admin\",\n            \"BackupAuthPass\": \"secret\",\n            \"BackupHost\": \"backup.my.domain.com\",\n            \"BackupPort\": 1022,\n            \"BackupPath\": \"/backup\",\n            \"ObjstoreS3SecretKeyId\": \"\",\n            \"ObjstoreS3SecretAccessKey\": \"\",\n            \"ObjstoreEndpointRegion\": \"\",\n            \"ObjstoreBucketSubfolder\": \"\",\n            \"BackupEncoding\": \"min(0-59) hour(0-23) dayofmonth(1-31) month(1-12) dayofweek(0-6)\",\n            \"BackupSchedule\": \"0 0 * * 0\"\n        },\n\n        \"CassandraEncryptionKeyFile\": \"encryption-secret.bin\",\n\n        \"PlatformApi\": \"platform.my.domain.com\",\n        \"ApiManagerUi\": \"ui.my.domain.com\",\n        \"CloudAdminUi\": \"cm.my.domain.com\",\n        \"ConsumerApi\": \"consumer.my.domain.com\"\n    },\n\n    \"Analytics\": {\n        \"SubsysName\": \"alt\",\n\n        \"VmFirstBoot\": {\n            \"Hosts\": [\n                {\"Name\": \"a1.my.domain.com\", \"HardDiskPassword\": \"password\", \"Device\": \"eth0\",\n                    \"IpAddress\": \"ip-address\", \"SubnetMask\": \"dot.subnet.mask\",\"Gateway\": \"gw-ip-address\"},\n                {\"Name\": \"a2.my.domain.com\", \"HardDiskPassword\": \"password\", \"Device\": \"eth0\",\n                    \"IpAddress\": \"ip-address\", \"SubnetMask\": \"dot.subnet.mask\", \"Gateway\": \"gw-ip-address\"},\n                {\"Name\": \"a3.my.domain.com\", \"HardDiskPassword\": \"password\", \"Device\": \"eth0\",\n                    \"IpAddress\": \"ip-address\", \"SubnetMask\": \"dot.subnet.mask\", \"Gateway\": \"gw-ip-address\"}\n            ]\n        },\n\n        \"EnableMessageQueue\": false,\n\n        \"AnalyticsIngestion\": \"ai.my.domain.com\",\n        \"AnalyticsClient\": \"ac.my.domain.com\"\n    },\n\n    \"Portal\": {\n        \"SubsysName\": \"ptl\",\n\n        \"VmFirstBoot\": {\n            \"Hosts\": [\n                {\"Name\": \"p1.my.domain.com\", \"HardDiskPassword\": \"password\", \"Device\": \"eth0\",\n                    \"IpAddress\": \"ip-address\", \"SubnetMask\": \"dot.subnet.mask\",\"Gateway\": \"gw-ip-address\"},\n                {\"Name\": \"p2.my.domain.com\", \"HardDiskPassword\": \"password\", \"Device\": \"eth0\",\n                    \"IpAddress\": \"ip-address\", \"SubnetMask\": \"dot.subnet.mask\", \"Gateway\": \"gw-ip-address\"},\n                {\"Name\": \"p3.my.domain.com\", \"HardDiskPassword\": \"password\", \"Device\": \"eth0\",\n                    \"IpAddress\": \"ip-address\", \"SubnetMask\": \"dot.subnet.mask\", \"Gateway\": \"gw-ip-address\"}\n            ]\n        },\n\n        \"SiteBackup\": {\n            \"BackupProtocol\": \"sftp|objstore\",\n            \"BackupAuthUser\": \"admin\",\n            \"BackupAuthPass\": \"secret\",\n            \"BackupHost\": \"backup.my.domain.com\",\n            \"BackupPort\": 1022,\n            \"BackupPath\": \"/backup\",\n            \"ObjstoreS3SecretKeyId\": \"\",\n            \"ObjstoreS3SecretAccessKey\": \"\",\n            \"ObjstoreEndpointRegion\": \"\",\n            \"ObjstoreBucketSubfolder\": \"\",\n            \"BackupEncoding\": \"min(0-59) hour(0-23) dayofmonth(1-31) month(1-12) dayofweek(0-6)\",\n            \"BackupSchedule\": \"0 2 * * *\"\n        },\n\n        \"PortalAdmin\": \"padmin.my.domain.com\",\n        \"PortalWww\": \"portal.my.domain.com\"\n    },\n\n    \"Gateway\": {\n        \"SubsysName\": \"gwy\",\n        \"Mode\": \"standard\",\n\n        \"SearchDomains\": [],\n        \"DnsServers\": [],\n\n        \"Hosts\": [\n            {\"Name\": \"dp1.my.domain.com\", \"HardDiskPassword\": \"password\", \"Device\": \"eth0\",\n                \"IpAddress\": \"ip-address\", \"SubnetMask\": \"dot.subnet.mask\",\"Gateway\": \"gw-ip-address\"},\n            {\"Name\": \"dp2.my.domain.com\", \"HardDiskPassword\": \"password\", \"Device\": \"eth0\",\n                \"IpAddress\": \"ip-address\", \"SubnetMask\": \"dot.subnet.mask\", \"Gateway\": \"gw-ip-address\"},\n            {\"Name\": \"dp3.my.domain.com\", \"HardDiskPassword\": \"password\", \"Device\": \"eth0\",\n                \"IpAddress\": \"ip-address\", \"SubnetMask\": \"dot.subnet.mask\", \"Gateway\": \"gw-ip-address\"}\n        ],\n\n        \"ApiGateway\": \"gw.my.domain.com\",\n        \"ApicGwService\": \"gwd.my.domain.com\"\n    }\n}"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1573371241, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "analytics-k8s.tmpl"
			file3, // "analytics-vm.tmpl"
			file4, // "cloud-init-vm.tmpl"
			file5, // "combined-csr.tmpl"
			file6, // "csr-client-auth.tmpl"
			file7, // "csr-server-auth.tmpl"
			file8, // "extra-values.tmpl"
			file9, // "gateway-k8s.tmpl"
			filea, // "helpers.tmpl"
			fileb, // "keypair.tmpl"
			filec, // "management-k8s.tmpl"
			filed, // "management-vm.tmpl"
			filee, // "portal-k8s.tmpl"
			filef, // "portal-vm.tmpl"
			fileg, // "subsys-certs.tmpl"
			fileh, // "subsys-config-k8s.tmpl"
			filei, // "subsys-config-ova.tmpl"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`../../templates`, &embedded.EmbeddedBox{
		Name: `../../templates`,
		Time: time.Unix(1573371241, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"analytics-k8s.tmpl":     file2,
			"analytics-vm.tmpl":      file3,
			"cloud-init-vm.tmpl":     file4,
			"combined-csr.tmpl":      file5,
			"csr-client-auth.tmpl":   file6,
			"csr-server-auth.tmpl":   file7,
			"extra-values.tmpl":      file8,
			"gateway-k8s.tmpl":       file9,
			"helpers.tmpl":           filea,
			"keypair.tmpl":           fileb,
			"management-k8s.tmpl":    filec,
			"management-vm.tmpl":     filed,
			"portal-k8s.tmpl":        filee,
			"portal-vm.tmpl":         filef,
			"subsys-certs.tmpl":      fileg,
			"subsys-config-k8s.tmpl": fileh,
			"subsys-config-ova.tmpl": filei,
		},
	})
}
