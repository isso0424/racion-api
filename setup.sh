#!/bin/sh
package=$(echo $1 | sed -e "s/\//\\\\\//g")
find . -type f -exec sed -i "s/isso0424\/gorilla-template/${package}/g" {} \;
go mod download
