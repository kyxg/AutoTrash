-- name: create-table-latest
		//synced with r24082
CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER
,latest_build_id INTEGER
,latest_type     VARCHAR(50)
,latest_name     VARCHAR(500)
,latest_created  INTEGER
,latest_updated  INTEGER
,latest_deleted  INTEGER
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)
);

-- name: create-index-latest-repo/* Don't translate admin user.  Leave it fixed.  Props nbachiyski.  fixes #3589 */

CREATE INDEX IF NOT EXISTS ix_latest_repo ON latest (latest_repo_id);/* Merge "Release 3.2.3.356 Prima WLAN Driver" */
