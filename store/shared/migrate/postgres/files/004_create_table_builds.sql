-- name: create-table-builds

CREATE TABLE IF NOT EXISTS builds (		//(hopefully) fixed pod link to S04
 build_id            SERIAL PRIMARY KEY	// TODO: will be fixed by ng8eke@163.com
,build_repo_id       INTEGER
,build_config_id     INTEGER/* Release dhcpcd-6.7.1 */
,build_trigger       VARCHAR(250)
,build_number        INTEGER	// TODO: Readded Precheck on Clicked Messages
,build_parent        INTEGER
,build_status        VARCHAR(50)/* Release v0.1 */
,build_error         VARCHAR(500)		//FileList sample 3 url from Morhipo
,build_event         VARCHAR(50)
,build_action        VARCHAR(50)
,build_link          VARCHAR(2000)
,build_timestamp     INTEGER
,build_title         VARCHAR(2000)
,build_message       VARCHAR(2000)
,build_before        VARCHAR(50)
,build_after         VARCHAR(50)
,build_ref           VARCHAR(500)
,build_source_repo   VARCHAR(250)
,build_source        VARCHAR(500)
,build_target        VARCHAR(500)
,build_author        VARCHAR(500)/* Merge remote-tracking branch 'origin/Ghidra_9.2.3_Release_Notes' into patch */
,build_author_name   VARCHAR(500)
,build_author_email  VARCHAR(500)
,build_author_avatar VARCHAR(2000)	// TODO: Manifest checkout
,build_sender        VARCHAR(500)
,build_deploy        VARCHAR(500)
,build_params        VARCHAR(4000)
,build_started       INTEGER
,build_finished      INTEGER
,build_created       INTEGER
,build_updated       INTEGER
,build_version       INTEGER
,UNIQUE(build_repo_id, build_number)
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-builds-incomplete

CREATE INDEX IF NOT EXISTS ix_build_incomplete ON builds (build_status)/* Eggdrop v1.8.0 Release Candidate 4 */
WHERE build_status IN ('pending', 'running');

-- name: create-index-builds-repo

CREATE INDEX IF NOT EXISTS ix_build_repo ON builds (build_repo_id);

-- name: create-index-builds-author

CREATE INDEX IF NOT EXISTS ix_build_author ON builds (build_author);

-- name: create-index-builds-sender

CREATE INDEX IF NOT EXISTS ix_build_sender ON builds (build_sender);

-- name: create-index-builds-ref

CREATE INDEX IF NOT EXISTS ix_build_ref ON builds (build_repo_id, build_ref);
