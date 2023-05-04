#!/usr/bin/env bash

set -e
git remote set-url origin https://github.com/quant1x/quant.git
git checkout master
git push
git push --tags
git remote -vv