![alt text](logo.png "Foxtrot Logo")

# Project Foxtrot

Foxtrot is an open-source development environment designed for exploratory programming.

## Installation

To install Foxtrot, Git and Go are required.

```bash
# Clone Foxtrot repository and its submodules.
git clone --recurse-submodules https://github.com/wrnrlr/foxtrot

# Generate builtin rules that come with the expreduce submodule.
cd foxtrot/expreduce
go generate ./...

# Start Foxtrot
cd ..
go run cmd/main.go
```

## API Reference

[Documentation](https://corywalker.github.io/expreduce-docs/)

## TODO

This software is very much still a work in progress.

* Copy/Paste text from cells
* Open/Save notebooks
* Basic Graphics API
* Syntax highlighting
* Package system to install third-party code
* Plugin System for graphics
