#!/bin/bash
cd ../
rm cscope*
rm tags
find ./ -name "*.go" > ./cscope.files
cscope -Rbqk
ctags -R

