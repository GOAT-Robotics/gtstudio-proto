#!/bin/bash

# This script will do the following.
# - Generate proto files
# - Add to git, commit, create tag, push both tag and commit.

# Derive version from remote tags; no hardcoded version here
version=""

set -euo pipefail

echo -e '\033[4;32;1mGenerating Proto Code\033[m'
./generate.sh

echo -e '\033[4;32;1mPreparing Version Bump\033[m'

# Compute next semantic version from remote tags (pattern vMAJOR.MINOR.PATCH)
latest=$(git ls-remote --tags --refs origin | awk -F/ '/v[0-9]+\.[0-9]+\.[0-9]+$/ {print $NF}' | sed 's/^v//' | sort -t. -k1,1n -k2,2n -k3,3n | tail -1)
if [ -n "${latest:-}" ]; then
  IFS=. read -r major minor patch <<< "$latest"
else
  major=1; minor=0; patch=0
fi

# Default bump: patch++ ; if patch >= 999 then minor++ and patch=0
patch=$((patch+1))
if [ "$patch" -ge 999 ]; then
  patch=0
  minor=$((minor+1))
fi

version="v${major}.${minor}.${patch}"

# Ensure no local/remote collision; keep bumping patch with rollover
while git rev-parse -q --verify "refs/tags/$version" >/dev/null 2>&1 || \
      git ls-remote --exit-code --tags origin "$version" >/dev/null 2>&1; do
  patch=$((patch+1))
  if [ "$patch" -ge 999 ]; then
    patch=0
    minor=$((minor+1))
  fi
  version="v${major}.${minor}.${patch}"
done

echo -e '\033[4;32;1mAdding To Git\033[m'
git add .

# Guard: exit if nothing to commit
if git diff --cached --quiet; then
  echo -e '\033[4;33;1mNo staged changes; skipping commit and tag.\033[m'
  exit 0
fi

echo -e '\033[4;32;1mGive a message: \033[m'
read message
git commit -m "$message"

echo "Creating a tag - $version"
git tag "$version"
echo -e '\033[4;32;1mPushing Code To Github\033[m'
git push

# Guard: only push the new tag explicitly; fail if remote tag exists
if git ls-remote --exit-code --tags origin "$version" >/dev/null 2>&1; then
  echo -e '\033[4;31;1mRemote tag $version already exists. Abort.\033[m'
  exit 1
fi
git push origin "$version"
echo -e '\033[4;32;1mUpdating Go Repository\033[m'
echo -e '\033[4;32;1mDone!\033[m'
