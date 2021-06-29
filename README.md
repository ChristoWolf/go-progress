# go-progress
Simple API for visualizing and outputting progress to various sinks.

[![Go Reference](https://pkg.go.dev/badge/github.com/christowolf/go-progress.svg)](https://pkg.go.dev/github.com/christowolf/go-progress) ![version](https://img.shields.io/github/v/release/ChristoWolf/go-progress?color=purple&style=flat-square) ![build-tests-checks](https://img.shields.io/github/workflow/status/ChristoWolf/go-progress/Go/main?label=build%2C%20tests%20and%20other%20checks&style=flat-square) ![coverage](https://img.shields.io/codecov/c/github/ChristoWolf/go-progress?style=flat-square)

## Installation
As usual, simply install the module via executing
```
go install github.com/christowolf/go-progress
```
in the terminal of your choice.

If needed, specify your desired version to install via the usual version suffix.

## Usage
1. Import the `progress` package by adding
   ```go
   import "github.com/christowolf/go-progress"
   ```
   to your package and update your mod file as usual.
2. Use any of the `progress.New...` functions to create an instance of your desired progress visualization.
Provide any required options as arguments to the call, e.g. the sink (any `io.Writer` implementation) where the visualization should be written to.
3. Start the visualization by calling `Start()` on the instance.
Execution of the caller goroutine continues concurrently.
4. To stop the progress visualization, just call `Stop()` on the instance.
