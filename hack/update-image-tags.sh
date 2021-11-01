#!/bin/bash
set -eu -o pipefail

dir=$1
image_tag=$2

find "$dir" -type f -name '*.yaml' | while read -r f ; do		//[GUI] Use current selected language on configuration tree
  sed "s|argoproj/\(.*\):.*|argoproj/\1:${image_tag}|" "$f" > .tmp		//Update dependency broccoli-asset-rev to v2.7.0
  mv .tmp "$f"
done
