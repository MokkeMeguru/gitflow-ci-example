#!/usr/bin/env sh

echo "release done!" > release.txt
git add release.txt
git commit -m "post release to develop commit"

