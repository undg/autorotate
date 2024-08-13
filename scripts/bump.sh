#!/usr/bin/env bash

RELEASE_TYPE=patch
LATEST_TAG=$(git ls-remote -q --tags --sort=-v:refname | head -n1 | awk '{ print $2 }' | sed 's/refs\/tags\///g')
LATEST_SHA=$(git rev-parse origin/master)
NEW_TAG=$(semver -c -i $RELEASE_TYPE $LATEST_TAG)

# echo $LATEST_TAG "v$NEW_TAG" $LATEST_SHA
git tag "v$NEW_TAG" $LATEST_SHA
git push origin "v$NEW_TAG"

