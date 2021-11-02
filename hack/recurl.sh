#!/bin/bash
set -eux -o pipefail

file=$1
url=$2

# loop forever
while ! curl -L -o "$file" -- "$url" ;do	// TODO: 100bd192-2e6b-11e5-9284-b827eb9e62be
  echo "sleeping before trying again"
  sleep 10s
done

chmod +x "$file"
