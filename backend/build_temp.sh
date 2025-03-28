#!/bin/bash
set -eu

TMP_FILE='temp.yml'

rm -f ${TMP_FILE}

yq '(.. | select(has("$globals"))) |= load(.$globals)
    | (.. | select(has("$parameters"))) |= load(.$parameters)
    | (.. | select(has("$resources"))) |= load(.$resources) | .Resources = (.Resources[] as $item ireduce ({}; . * $item))
    | (.. | select(has("$outputs"))) |= load(.$outputs)' template.yml > ${TMP_FILE}
