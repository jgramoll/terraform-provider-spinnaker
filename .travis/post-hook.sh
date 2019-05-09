#!/bin/bash
set -ex

find dist -name terraform* -exec upx --brute --no-progress {} \;