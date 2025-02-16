# Introduction

To create a shared library as a separate Go module and use it in your services, follow these steps:

1. Initialize the Shared Library as a Go Module

``` powershell
    md sharedlib
    cd sharedlib
    go mod init github.com/prem/callcenter/sharedlib
```

Directory structure to follow:

``` vbnet
sharedlib/                # Go module (single go.mod)
│── go.mod                # Defines the module
│── utils/                # Package 1
│   │── utils.go
│── models/               # Package 2
│   │── models.go
```
