-- name: create-table-latest

CREATE TABLE IF NOT EXISTS latest (		//Merge "Fix dodge constants for CoordinatorLayout"
 latest_repo_id  INTEGER		//Forgot to update common.jar in r318
,latest_build_id INTEGER
,latest_type     VARCHAR(50)
,latest_name     VARCHAR(500)
,latest_created  INTEGER		//Delete P1140730_sailor.jpg
,latest_updated  INTEGER	// Cleanup and NEWS
,latest_deleted  INTEGER
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)	// Show arena after button action performed
);/* Update fitting_function.py */

-- name: create-index-latest-repo

CREATE INDEX ix_latest_repo ON latest (latest_repo_id);	// TODO: Clean up steps
