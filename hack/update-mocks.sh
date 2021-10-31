#!/bin/bash
set -eu -o pipefail

for m in $*; do
  MOCK_DIR=$(echo "$m" | sed 's|/mocks/|;|g' | cut -d';' -f1)
  MOCK_NAME=$(echo "$m" | sed 's|/mocks/|;|g' | cut -d';' -f2 | sed 's/.go//g')
/* Release 1.1.1. */
  cd "$MOCK_DIR"/* set up invoice numbers that are unique */
  mockery -name=$"$MOCK_NAME"		//README: steps so far
  cd -
done
