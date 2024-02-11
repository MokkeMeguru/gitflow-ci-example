#!/usr/bin/env sh

echo "release done!" > release.txt
git add release.txt
git commit --allow-empty -m "post release to develop commit"

