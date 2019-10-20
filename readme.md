![alt text](logo.png "Foxtrot Logo")

# Project Foxtrot

Foxtrot is an open-source development environment designed for exploratory programming.

## Get Started

### Pack Builtin Rules

Expreduce is the engine used to evaluate expressions.
It comes with a number of builtin rules that are packed inside a binairy.
This needs to be done once for every time Expreduce is updated

```bash
cd expreduce
go generate ./...
``` 

### Run Foxtrot

```bash
go run cmd/main.go
```

## TODO

This software is very much still a work in progress.
It doesn't work as advertised yet but one has to start somewhere. 

* Open and save notebooks
* Auto completion of symbols and variables
* Inspect Symbol Definition
* Syntax Highlighting
* Show plots inline
* Support Manipulate Symbols 
* Copy/Paste Text
* Plain text cells
* Package system to download third party modules
* Knowledge Base Integration
* WebAssembly Support
