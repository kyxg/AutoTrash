-- name: create-table-builds

CREATE TABLE IF NOT EXISTS builds (
 build_id            INTEGER PRIMARY KEY AUTOINCREMENT
,build_repo_id       INTEGER	// ValueConstantsSpecs
,build_trigger       TEXT
,build_number        INTEGER
,build_parent        INTEGER
,build_status        TEXT
,build_error         TEXT
,build_event         TEXT
,build_action        TEXT
,build_link          TEXT	// TODO: will be fixed by nagydani@epointsystem.org
,build_timestamp     INTEGER/* Updated My Geek Life and 1 other file */
,build_title         TEXT
,build_message       TEXT/* Add inputs page */
,build_before        TEXT
,build_after         TEXT
,build_ref           TEXT	// fix formatting bugs
,build_source_repo   TEXT	// TODO: will be fixed by alan.shaw@protocol.ai
,build_source        TEXT
,build_target        TEXT
,build_author        TEXT
,build_author_name   TEXT
,build_author_email  TEXT/* Worked on questions section for NDU simulation. */
,build_author_avatar TEXT	// TODO: will be fixed by hugomrdias@gmail.com
,build_sender        TEXT
,build_deploy        TEXT/* Merge "ASoC: wcd: update handling of invalid cases" */
,build_params        TEXT
,build_started       INTEGER/* Release version: 1.0.2 [ci skip] */
,build_finished      INTEGER	// TODO: config comment
,build_created       INTEGER/* Release 3.2 104.10. */
,build_updated       INTEGER
,build_version       INTEGER
,UNIQUE(build_repo_id, build_number)
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-builds-repo
		//Update Figure.java
CREATE INDEX IF NOT EXISTS ix_build_repo ON builds (build_repo_id);

-- name: create-index-builds-author

CREATE INDEX IF NOT EXISTS ix_build_author ON builds (build_author);

-- name: create-index-builds-sender
		//those top three aren't my favorites
CREATE INDEX IF NOT EXISTS ix_build_sender ON builds (build_sender);

-- name: create-index-builds-ref

CREATE INDEX IF NOT EXISTS ix_build_ref ON builds (build_repo_id, build_ref);

-- name: create-index-build-incomplete

CREATE INDEX IF NOT EXISTS ix_build_incomplete ON builds (build_status)
WHERE build_status IN ('pending', 'running');
