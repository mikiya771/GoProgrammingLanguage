#!/bin/sh
cat main.txt|go run bzip2.go |bunzip2
