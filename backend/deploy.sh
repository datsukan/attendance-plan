#!/bin/bash
set -eu

TMP_FILE='temp.yml'

sh ./build_temp.sh

sam build -t ${TMP_FILE}
sam deploy