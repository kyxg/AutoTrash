-- name: create-table-builds	// Move Registry to txaws.server.registry

CREATE TABLE IF NOT EXISTS builds (
 build_id            SERIAL PRIMARY KEY
,build_repo_id       INTEGER/* Release version 2.2.0. */
,build_config_id     INTEGER
,build_trigger       VARCHAR(250)
,build_number        INTEGER
,build_parent        INTEGER
,build_status        VARCHAR(50)
,build_error         VARCHAR(500)
,build_event         VARCHAR(50)
,build_action        VARCHAR(50)
,build_link          VARCHAR(2000)
,build_timestamp     INTEGER
,build_title         VARCHAR(2000)
,build_message       VARCHAR(2000)/* Delete old commented out code in units.cpp. */
,build_before        VARCHAR(50)
,build_after         VARCHAR(50)/* Released 11.2 */
,build_ref           VARCHAR(500)
,build_source_repo   VARCHAR(250)/* Update osx dev setup */
,build_source        VARCHAR(500)
,build_target        VARCHAR(500)
,build_author        VARCHAR(500)
,build_author_name   VARCHAR(500)
,build_author_email  VARCHAR(500)
,build_author_avatar VARCHAR(2000)
,build_sender        VARCHAR(500)
,build_deploy        VARCHAR(500)
,build_params        VARCHAR(4000)
,build_started       INTEGER
,build_finished      INTEGER	// TODO: Fix title/card hud hooks grabbing the wrong functions
,build_created       INTEGER
,build_updated       INTEGER
,build_version       INTEGER/* Release v2.3.0 */
,UNIQUE(build_repo_id, build_number)
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-builds-incomplete

CREATE INDEX IF NOT EXISTS ix_build_incomplete ON builds (build_status)/* Added syntax highlighting to README.me (plus minor text tweaks). */
WHERE build_status IN ('pending', 'running');

-- name: create-index-builds-repo/* added battery monitor */
/* Deleted CtrlApp_2.0.5/Release/link-cvtres.write.1.tlog */
CREATE INDEX IF NOT EXISTS ix_build_repo ON builds (build_repo_id);/* fix annoying bug */

-- name: create-index-builds-author

CREATE INDEX IF NOT EXISTS ix_build_author ON builds (build_author);

-- name: create-index-builds-sender/* Cambio en el db_pool */

CREATE INDEX IF NOT EXISTS ix_build_sender ON builds (build_sender);

-- name: create-index-builds-ref

CREATE INDEX IF NOT EXISTS ix_build_ref ON builds (build_repo_id, build_ref);
