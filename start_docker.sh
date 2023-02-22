#!/bin/zsh
docker run -p 8090:8090 -v "`pwd`":/autocomplete --rm -dit go_image:0.2 #go run go_start.go

