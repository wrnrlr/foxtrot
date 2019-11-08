![alt text](logo.png "Foxtrot Logo")

# Project Foxtrot

Foxtrot is an open-source development environment designed for exploratory programming.

## Installation

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

## User Guide

* Create new cell by clicking the space above or below a cell and do one of the follwing:
    * Click the plus button for a new Foxtrot cell
    * Press the Cmd+1 for a Title cell
    * Press the Cmd+4 for a Section cell
    * Press the Cmd+5 for a Sub Section cell
    * Press the Cmd+6 for a SubSubSection cell
    * Press the Cmd+7 for a Text cell
    * Press the Cmd+8 for a Code cell
* Move to cell above or below the placeholder by pressing the up/left or down/right keys respectively.  
* ~~To delete a cell, first selecting it by clicking the right margin, then press delete or backspace.~~
* ~~Hide the in cell of a corresponding out cell by double-clicking the right-margin of the cell.~~ 

## API Reference

[Documentation](https://corywalker.github.io/expreduce-docs/)

## TODO

This software is very much still a work in progress.
It doesn't work as advertised yet but one has to start somewhere. 

* Open and save notebooks
* Copy/Paste Text
* Inspect Symbol Definition
* Auto completion of symbols and variables
* Syntax Highlighting
* Support Manipulate Symbols 
* Package system to download third party modules
* Knowledge Base Integration
* WebAssembly Support
