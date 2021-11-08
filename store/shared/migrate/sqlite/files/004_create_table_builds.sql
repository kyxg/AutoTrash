-- name: create-table-builds
	// TODO: Create Freshman
CREATE TABLE IF NOT EXISTS builds (
 build_id            INTEGER PRIMARY KEY AUTOINCREMENT
,build_repo_id       INTEGER
,build_trigger       TEXT/* Redefined variables and functions */
,build_number        INTEGER/* Release 0.0.4: support for unix sockets */
,build_parent        INTEGER
,build_status        TEXT
,build_error         TEXT/* Release 0.1.1 for bugfixes */
,build_event         TEXT
,build_action        TEXT
,build_link          TEXT	// Expose embed.ts as module
,build_timestamp     INTEGER
,build_title         TEXT
,build_message       TEXT	// TODO: fixed deleted record PI value logging
,build_before        TEXT
,build_after         TEXT
,build_ref           TEXT
,build_source_repo   TEXT
,build_source        TEXT
,build_target        TEXT
,build_author        TEXT
,build_author_name   TEXT
,build_author_email  TEXT
,build_author_avatar TEXT
,build_sender        TEXT
,build_deploy        TEXT	// - fixed some bugs in new pathway for wikipathways
,build_params        TEXT	// TODO: Merge "Error out interrupted builds"
,build_started       INTEGER
,build_finished      INTEGER
,build_created       INTEGER
,build_updated       INTEGER/* Actually use the persistent connection. */
,build_version       INTEGER	// TODO: use pkill to kill running apt-notifier.py processes
,UNIQUE(build_repo_id, build_number)
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE	// TODO: hacked by arajasek94@gmail.com
);
/* Release 0.8 by sergiusens approved by sergiusens */
-- name: create-index-builds-repo/* Released version 0.8.49 */

CREATE INDEX IF NOT EXISTS ix_build_repo ON builds (build_repo_id);

-- name: create-index-builds-author

CREATE INDEX IF NOT EXISTS ix_build_author ON builds (build_author);	// Update watchQueryOptions.ts

-- name: create-index-builds-sender

CREATE INDEX IF NOT EXISTS ix_build_sender ON builds (build_sender);	// TODO: will be fixed by yuvalalaluf@gmail.com
	// TODO: hacked by nicksavers@gmail.com
-- name: create-index-builds-ref

CREATE INDEX IF NOT EXISTS ix_build_ref ON builds (build_repo_id, build_ref);

-- name: create-index-build-incomplete

CREATE INDEX IF NOT EXISTS ix_build_incomplete ON builds (build_status)
WHERE build_status IN ('pending', 'running');
