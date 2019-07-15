# struct2oas ![](https://img.shields.io/badge/version-1.1.0-blue.svg) [![Build Status](https://travis-ci.org/paulvollmer/struct2oas.svg?branch=master)](https://travis-ci.org/paulvollmer/struct2oas)

struct2oas is a tool to automate the creation of openapi schemas.
It will parse the go sourcecode and generate a schema yaml file.

## Installation

```sh
go get github.com/paulvollmer/struct2oas
```

## Usage

```sh
struct2oas -source file.go
```
