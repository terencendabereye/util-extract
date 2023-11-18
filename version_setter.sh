#! /usr/local/bin/fish

# sets the current verison number based on the current git tag

set VER $(git describe --tags)
sed -i "" 's/v[0-9]\.[0-9]\.[0.9]/'$VER'/g' cmd/root.go