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

### Install Software

#### OSX

#### Windows

#### Linux

### Open Notebook

Open an existing notebook, or create a new notebook by calling the `foxtrot` command with as argument the path of the notebook file. 

```bash
$ foxtrot notebook.xnb
```

Start a temporary notebook by calling `foxtrot` without any argument. After closing the editor, all changes to the temporary notebook will be lost.


### Create Cells

Create new cell by clicking the space above or below a cell, and either click the plus button or use on of the following keyboard shortcuts:

* Press the Cmd+1 for a Title cell
* Press the Cmd+4 for a Section cell
* Press the Cmd+5 for a Sub Section cell
* Press the Cmd+6 for a SubSubSection cell
* Press the Cmd+7 for a Text cell
* Press the Cmd+8 for a Code cell

### Move between cells

Move to cell above or below the slot by pressing the up/left or down/right keys respectively.

### Select Cells

Cells can be selected with either the mouse or the keyboard.
Click or tab on a cell's margin to select it, hold down the shift key to select a range of cells.
When a cell is selected its margin is highlighted with a blue background.


#### Delete Cells

Selected cells are removed with the the delete or backspace key.

#### ~~Copy, Cut and Paste cells~~

Not yet implemented.

### Plot Function

```
Plot[Sin[x],{x,0,7}]
```

### Draw Graphics

## API Reference

[Documentation](https://corywalker.github.io/expreduce-docs/)

## TODO

This software is very much still a work in progress.
It doesn't work as advertised yet but one has to start somewhere. 
