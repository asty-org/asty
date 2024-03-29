# asty

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/asty-org/asty)
[![Go Reference](https://pkg.go.dev/badge/github.com/asty-org/asty.svg)](https://pkg.go.dev/github.com/asty-org/asty)
[![Go Report Card](https://goreportcard.com/badge/github.com/asty-org/asty)](https://goreportcard.com/report/github.com/asty-org/asty)
![Codecov](https://img.shields.io/codecov/c/github/asty-org/asty)
[![GitHub last commit](https://img.shields.io/github/last-commit/asty-org/asty)](https://github.com/asty-org/asty)
[![Docker Pulls](https://img.shields.io/docker/pulls/astyorg/asty)](https://hub.docker.com/r/astyorg/asty)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go#code-analysis)

_Not another JSON parser!_

**AST &#8594; JSON** | **JSON &#8594; AST**

Marshals golang [AST](https://pkg.go.dev/go/ast) into JSON and unmarshals it back from JSON.

It allows building pattern matching, statistical analysis, language transformation, search/data-mine/anything algorithms
for golang with any other language (I like to do it with python. Check out [asty-python](https://github.com/asty-org/asty-python))

## Example

[Try it!](https://asty-org.github.io/)

Input golang source

```golang
package main

import "fmt"

func main() {
    fmt.Println("hello world")
}
```

Ouput AST in JSON

```json
{
  "NodeType": "File",
  "Name": {
    "NodeType": "Ident",
    "Name": "main"
  },
  "Decls": [
    {
      "NodeType": "GenDecl",
      "Tok": "import",
      "Specs": [
        {
          "NodeType": "ImportSpec",
          "Name": null,
          "Path": {
            "NodeType": "BasicLit",
            "Kind": "STRING",
            "Value": "\"fmt\""
          }
        }
      ]
    },
    {
      "NodeType": "FuncDecl",
      "Recv": null,
      "Name": {
        "NodeType": "Ident",
        "Name": "main"
      },
      "Type": {
        "NodeType": "FuncType",
        "TypeParams": null,
        "Params": {
          "NodeType": "FieldList",
          "List": null
        },
        "Results": null
      },
      "Body": {
        "NodeType": "BlockStmt",
        "List": [
          {
            "NodeType": "ExprStmt",
            "X": {
              "NodeType": "CallExpr",
              "Fun": {
                "NodeType": "SelectorExpr",
                "X": {
                  "NodeType": "Ident",
                  "Name": "fmt"
                },
                "Sel": {
                  "NodeType": "Ident",
                  "Name": "Println"
                }
              },
              "Args": [
                {
                  "NodeType": "BasicLit",
                  "Kind": "STRING",
                  "Value": "\"hello world\""
                }
              ]
            }
          }
        ]
      }
    }
  ]
}
```

## Install

Install `asty` under `$GOPATH/bin`

```bash
go install github.com/asty-org/asty

asty -h
```

## Building

Just `make`
If you want to do it differently use `go build`

## Usage

Convert AST to JSON

```bash
asty go2json -input <input.go> -output <output.json>
```

Convert JSON to AST

```bash
asty json2go -input <input.json> -output <output.go>
```

Use `asty help` for more information

Using with docker

```bash
docker run astyorg/asty go2json -input <input.go> -output <output.json>
```

## Development principles

- Make json output as close to real golang structures as possible. There is no additional logic introduced.
  No normalization. No reinterpretation. The only things that were introduced are the names of some enum values.
- Make it very explicit. No reflection. No listing of fields. This is done to facilitate future maintenance.
  If something will be changed in future versions of golang this code will probably break compile-time.
- Keep polymorphism in JSON structure. If some field references _expression_ then particular type will be
  discriminated from object type name stored in separate field.

## Other solutions

- https://github.com/ReconfigureIO/goblin reinterpret some structures (trying to simplify).
  Out of maintenance for a long time. Still works in some forks.
- https://github.com/CreativeInquiry/go2json tries to parse golang code with parser written in javascript.
  Also lacks maintenance. Developed for particular use case of HaXe traspiler.

## Article

I wrote [an article](https://dev.to/evgenus/analyzing-ast-in-go-with-json-tools-36dg) about this project.
