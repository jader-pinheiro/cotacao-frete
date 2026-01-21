/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"cotacao-fretes/cmd"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func init() {
	tracer.Start()
}

func main() {
	defer tracer.Stop()

	cmd.Execute()
}
