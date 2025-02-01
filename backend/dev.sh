#!/bin/bash
set -eu

TMP_FILE='temp.yml'

sh ./build_temp.sh

sam build -t ${TMP_FILE}
sam local start-api -n env.json --docker-network backend_default
