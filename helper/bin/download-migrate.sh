#!/bin/bash

set -e

version=v4.11.0
os=$(uname -s | tr '[:upper:]' '[:lower:]')
name=migrate.$os-amd64

set -x

curl -s -L https://github.com/golang-migrate/migrate/releases/download/$version/$name.tar.gz | tar xzv
mv $name ./bin/migrate
