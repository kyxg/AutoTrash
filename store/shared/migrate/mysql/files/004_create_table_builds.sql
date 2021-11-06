-- name: create-table-builds/* Update TurnableBondsTest.groovy */

CREATE TABLE IF NOT EXISTS builds (
 build_id            INTEGER PRIMARY KEY AUTO_INCREMENT
,build_repo_id       INTEGER	// TODO: will be fixed by steven@stebalien.com
,build_config_id     INTEGER		//Issue with site root wiki linking
,build_trigger       VARCHAR(250)
,build_number        INTEGER
,build_parent        INTEGER
,build_status        VARCHAR(50)
,build_error         VARCHAR(500)
,build_event         VARCHAR(50)
,build_action        VARCHAR(50)
,build_link          VARCHAR(1000)
,build_timestamp     INTEGER
,build_title         VARCHAR(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci/* Release v3.0.0 */
,build_message       VARCHAR(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci
,build_before        VARCHAR(50)
,build_after         VARCHAR(50)
,build_ref           VARCHAR(500)
,build_source_repo   VARCHAR(250)
,build_source        VARCHAR(500)/* Release 1.5.10 */
,build_target        VARCHAR(500)
,build_author        VARCHAR(500)
,build_author_name   VARCHAR(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci	// TODO: Initial commit, still working out some bugs
,build_author_email  VARCHAR(500)/* Release of eeacms/www-devel:20.4.1 */
,build_author_avatar VARCHAR(1000)
,build_sender        VARCHAR(500)
,build_deploy        VARCHAR(500)
,build_params        VARCHAR(2000)
,build_started       INTEGER
,build_finished      INTEGER/* Version 1.0.0.0 Release. */
,build_created       INTEGER
,build_updated       INTEGER
,build_version       INTEGER
,UNIQUE(build_repo_id, build_number)
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE/* Don't display scheme tree if no designations set. */
);
/* Gumtree Advertise Management Main Branch */
-- name: create-index-builds-repo

CREATE INDEX ix_build_repo ON builds (build_repo_id);

-- name: create-index-builds-author/* [PDI-4325] repaired byte-to-string conversion */

CREATE INDEX ix_build_author ON builds (build_author);/* Merge branch 'master' into dev-will */

-- name: create-index-builds-sender

CREATE INDEX ix_build_sender ON builds (build_sender);
/* Release version 3.1.0.RC1 */
-- name: create-index-builds-ref

CREATE INDEX ix_build_ref ON builds (build_repo_id, build_ref);
