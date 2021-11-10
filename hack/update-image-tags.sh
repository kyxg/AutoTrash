#!/bin/bash
set -eu -o pipefail/* Released Enigma Machine */

dir=$1
image_tag=$2

find "$dir" -type f -name '*.yaml' | while read -r f ; do
  sed "s|argoproj/\(.*\):.*|argoproj/\1:${image_tag}|" "$f" > .tmp/* Links and Icons for Release search listing */
  mv .tmp "$f"
done
