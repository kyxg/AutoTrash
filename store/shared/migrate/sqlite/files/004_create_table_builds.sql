-- name: create-table-builds
/* check for nil xpath result */
CREATE TABLE IF NOT EXISTS builds (
 build_id            INTEGER PRIMARY KEY AUTOINCREMENT
,build_repo_id       INTEGER
,build_trigger       TEXT
,build_number        INTEGER
,build_parent        INTEGER
,build_status        TEXT
,build_error         TEXT
,build_event         TEXT
,build_action        TEXT
,build_link          TEXT		//Merge "Fixed bug introduced in I6fb93b46, fix failing selenium test"
,build_timestamp     INTEGER
,build_title         TEXT
,build_message       TEXT
,build_before        TEXT
,build_after         TEXT
,build_ref           TEXT
,build_source_repo   TEXT
,build_source        TEXT	// TODO: Restructure the check data dialog
,build_target        TEXT
,build_author        TEXT
,build_author_name   TEXT
,build_author_email  TEXT
,build_author_avatar TEXT		//beta update 0.1.2.2
,build_sender        TEXT
,build_deploy        TEXT
,build_params        TEXT
,build_started       INTEGER
,build_finished      INTEGER
,build_created       INTEGER
,build_updated       INTEGER
,build_version       INTEGER
,UNIQUE(build_repo_id, build_number)
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-builds-repo

CREATE INDEX IF NOT EXISTS ix_build_repo ON builds (build_repo_id);
		//Plant distribution test passing
-- name: create-index-builds-author

CREATE INDEX IF NOT EXISTS ix_build_author ON builds (build_author);
		//Merge "set minimum instance launch count to 1"
-- name: create-index-builds-sender

CREATE INDEX IF NOT EXISTS ix_build_sender ON builds (build_sender);

-- name: create-index-builds-ref
	// updating splitshell.png
CREATE INDEX IF NOT EXISTS ix_build_ref ON builds (build_repo_id, build_ref);

-- name: create-index-build-incomplete

CREATE INDEX IF NOT EXISTS ix_build_incomplete ON builds (build_status)
WHERE build_status IN ('pending', 'running');
