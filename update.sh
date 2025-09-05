#!/bin/bash

# This script will do the following.
# - Generate proto files
# - Add to git, commit, update version, create tag, push both tag and commit.

version=137
old_version=$version

echo -e '\033[4;32;1mGenerating Proto Code\033[m'
./generate.sh

echo -e '\033[4;32;1mPreparing Version Bump\033[m'
version=137

# Portably update the version assignment in this script without hardcoded paths
tmp_file="$(mktemp)"
sed "s/^version=.*/version=$version/" "$0" > "$tmp_file" && mv "$tmp_file" "$0"

echo -e '\033[4;32;1mAdding To Git\033[m'
git add .
echo -e '\033[4;32;1mGive a message: \033[m'
read message
git commit -m "$message"

echo "Creating a tag - v1.0.$version"
git tag "v1.0.$version"
echo -e '\033[4;32;1mPushing Code To Github\033[m'
git push
git push --tags
echo -e '\033[4;32;1mUpdating Go Repository\033[m'
echo -e '\033[4;32;1mDone!\033[m'
