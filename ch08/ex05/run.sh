#!/bin/sh
time go run nogoroutine.go  > test.png
time go run main.go  > test.png
