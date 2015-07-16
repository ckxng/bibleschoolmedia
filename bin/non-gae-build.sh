#!/bin/bash

ls main.go || exit 1

go build -v main.go

