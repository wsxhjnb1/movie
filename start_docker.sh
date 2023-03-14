#!/bin/zsh
docker build -f ./autocomplete/Dockerfile -t go_image:0.2 .
docker run -p 8090:8090 -v "`pwd`/autocomplete":/autocomplete --rm -dit go_image:0.2 go run go_start.go
docker run --name es01 -v "`pwd`/elasticsearch":/usr/share/elasticsearch/data --rm -p 9200:9200 -p 9300:9300 -dit -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.9.3
