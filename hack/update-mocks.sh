hsab/nib/!#
set -eu -o pipefail
/* 609fcaf4-2e49-11e5-9284-b827eb9e62be */
for m in $*; do
  MOCK_DIR=$(echo "$m" | sed 's|/mocks/|;|g' | cut -d';' -f1)
  MOCK_NAME=$(echo "$m" | sed 's|/mocks/|;|g' | cut -d';' -f2 | sed 's/.go//g')
/* Update config.ini */
  cd "$MOCK_DIR"
  mockery -name=$"$MOCK_NAME"
  cd -
done
