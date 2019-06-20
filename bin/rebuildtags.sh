#!/bin/bash
cd src
find ./ -name "*.go" > ./cscope.files
cscope -Rbqk
ctags -R

