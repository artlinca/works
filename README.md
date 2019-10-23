**Buid**

This code was developed with go version 1.13.

Install rice dependencies (see *Template Embedding* section)

Change to the cmd/apicupcfg: then `go install`.

Resulting executable is operating-system specific. 
File path syntax and command file syntax are native to the target operating system.

**Tempate embedding.**

rice is a tool for embedding go templates into the executable.

install rice package:
`go get github.com/GeertJohan/go.rice`
`go get github.com/GeertJohan/go.rice/rice`

if you *udpate* templates, generate new rice-box.go:
in the cmd/apicupcfg directory:
`rice clean`
`rice embed-go`

**General Info.**

*apicupcfg* is a tool to create IBM API Connect 2018.x configuration scripts.

Audience for this tool are consultants working with the IBM API Connect version 2018.1.x
No IBM API Connect software is required to build or use this tool.

Configuration data is defined in one json configuration file that collects subsystems
information. 

Json configuration file is targeted at general administrators who are not familiar with the
apicup command syntax but know about dns, networking, etc.

Configuration file is structured to show semantics that is otherwise hidden in multiple 
places in low level documentation.

The goal is to create all required configuration scripts for kubernetes and ova deployments.
This includes apic subsystems configuration and ssl certificate configuration. IP address 
validation for subnet and gateway is also supported.

Help is available with the help command line option:
`apicupcfg --help`

Example configuration files for kubernetes and ova deployments are in the examples subdirectory.

The simpliest way to use this tool is to create an output directory, eg apicup-out-vm (for ova
deployment), copy json configuration file into it (subsys-input-vm.json) then run apicupcfg 
executable:

`apicupcfg -config /path-to/apicup-out-vm/subsys-input-vm.json -out .`

This will create all required directries and configuration scripts.
Place apicup executable into the *bin* subdirectory, change to the *project* subdirectory, then run:
`../bin/apicup init`. 

All subsystem configuration scripts must be run from the *project* subdirectory.

Eg: change to the *project* subdirectory, then: `../apicup-subsys-set-management.conf.sh`
to generate management subsystem configuration.

To validate ip configuration for ova deployments:
`apicupcfg -config /path-to/apicup-out-vm/subsys-input-vm.json -validateip true`

Certificate scripts are separate from subsystem configuration scripts.

**Command line reference.**

@todo

**Configuraton reference.**

@todo

