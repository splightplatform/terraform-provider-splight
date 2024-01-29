#!/bin/bash

VERSION=$(cat version)
echo "building terraform-provider-spl_${VERSION}"
go build -o terraform-provider-spl_${VERSION}