# aura

[![Go](https://img.shields.io/badge/--00ADD8?logo=go&logoColor=ffffff)](https://go.dev/) [![Go Report Card](https://goreportcard.com/badge/github.com/x86ed/aura)](https://goreportcard.com/report/github.com/x86ed/aura) [![PkgGoDev](https://pkg.go.dev/badge/github.com/x86ed/aura)](https://pkg.go.dev/github.com/x86ed/aura)

aura - terminal text effect generator/library in go. aura can be used as a comand line utility or a library in other go projects. It's meant to make working with the commeandline easier by producing canned strings/[]bytes so that you can incorporate these prompts in your projects and also as a library so you can quickly use escape commands.

## Library

### color

#### 16 & 32 color escape codes

This library can set any color settings you might want to use from Xterm. all commands can be used either with or without a string. if used without ie: `color.Red()` it just returns the string to open that color mode. To return your string to normal you must call `color.Reset()`. if used with a string for example `color.Red("Yikes!")` `color.Reset()` will be called automatically at the end of the string.

### cursor

This library allows you to set cursor commands. For example `cursor.Move2Coord(2,2)` moves the cursor to the position x:2 y:2.

### erase

This library is used to erase various pieces of content on the screen.

### screen

This library can set screen resolution/params and get info about the screen.

### style

This library styles text.

### image

Image is a library that can re-render images to ascii including gifs.

### modal

This library create modals which can hold any of the content listed above.

### select

This library allows you to make multi select options

### grid

this library can draw grids on the screen from tabular data.

### progbar

this library allows you to create custom progressbars and spinners for your application.

### graph

a package that allows you to render simple graphs.

## CLI

the cli allows you to generate tons of terminal strings in various languages. Here's how to use it:
