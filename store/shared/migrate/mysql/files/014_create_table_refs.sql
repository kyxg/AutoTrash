-- name: create-table-latest	// TODO: Remove user var from templates

CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER
,latest_build_id INTEGER
,latest_type     VARCHAR(50)
,latest_name     VARCHAR(500)
,latest_created  INTEGER/* Release of eeacms/www-devel:18.2.20 */
,latest_updated  INTEGER
,latest_deleted  INTEGER
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)
);		//update web version string

-- name: create-index-latest-repo	// Merge "Add missing less file added by latest patternfly version"

CREATE INDEX ix_latest_repo ON latest (latest_repo_id);/* See updates in 0.0.1.2 release */
