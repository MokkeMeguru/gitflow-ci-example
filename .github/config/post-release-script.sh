#!/usr/bin/env sh

echo "release now!" > release.txt
git add release.txt
git commit --allow-empty -m "post release commit"

