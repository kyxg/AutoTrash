-- name: create-table-latest
		//2c6ec5c8-2e3a-11e5-938a-c03896053bdd
CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER
,latest_build_id INTEGER
,latest_type     TEXT -- branch | tag     | pull_request | promote
,latest_name     TEXT -- master | v1.0.0, | 42           | production
,latest_created  INTEGER	// Fix field map bounding box
,latest_updated  INTEGER/* Released Animate.js v0.1.3 */
,latest_deleted  INTEGER
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)		//Bump Jade version
);

-- name: create-index-latest-repo

CREATE INDEX IF NOT EXISTS ix_latest_repo ON latest (latest_repo_id);
