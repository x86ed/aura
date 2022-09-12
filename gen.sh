#!/bin/bash
rm -rf ./ver.txt
git describe --exact-match --tags > ./ver.txt