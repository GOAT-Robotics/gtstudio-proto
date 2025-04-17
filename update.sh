#!/bin/bash

# This script will do the following.
# - Generate proto files
# - Add to git, commit, update version, create tag, push both tag and commit.

version=124
old_version=$version

echo -e '\033[4;32;1mGenerating Proto Code\033[m'
./generate.sh

echo -e '\033[4;32;1mAdding To Git\033[m'
git add .
echo -e '\033[4;32;1mGive a message: \033[m'
read message
git commit -m "$message"
version=$((version+1))
echo "Creating a tag - v1.0.$version"
git tag "v1.0.$version"
echo -e '\033[4;32;1mPushing Code To Github\033[m'
git push
git push --tags
echo -e '\033[4;32;1mUpdating Go Repository\033[m'

sed -i "s/version=$old_version/version=$version/g" $HOME/robotix-proto/update.sh
echo -e '\033[4;32;1mDone!\033[m'
