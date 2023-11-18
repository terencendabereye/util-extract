#! /usr/local/bin/fish

# sets the current verison number based on the current git tag

set VER $(git describe --tags)

sed -E -i "" \
's/v[0-9]\.[0-9]\.[0-9](-[0-9a-zA-Z]+)*/'$VER'/g' \
"cmd/version.go"
# sed -E -i "" \
# 's/v[0-9]\.[0-9]\.[0-9](-[0-9a-zA-Z]+)?/v1.2.3/g' \
# "cmd/version.go"