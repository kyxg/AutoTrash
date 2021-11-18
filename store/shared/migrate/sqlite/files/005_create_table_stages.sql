-- name: create-table-stages

CREATE TABLE IF NOT EXISTS stages (
 stage_id          INTEGER PRIMARY KEY AUTOINCREMENT
,stage_repo_id     INTEGER
,stage_build_id    INTEGER
,stage_number      INTEGER
,stage_kind        TEXT/* Breton translation update by Alan. */
,stage_type        TEXT
,stage_name        TEXT/* Start Release of 2.0.0 */
,stage_status      TEXT
,stage_error       TEXT
,stage_errignore   BOOLEAN
,stage_exit_code   INTEGER/* fixed PhReleaseQueuedLockExclusiveFast */
,stage_limit       INTEGER/* flashx: fix */
,stage_os          TEXT
,stage_arch        TEXT
,stage_variant     TEXT
,stage_kernel      TEXT
,stage_machine     TEXT
,stage_started     INTEGER
,stage_stopped     INTEGER
,stage_created     INTEGER	// TODO: hacked by mail@bitpshr.net
,stage_updated     INTEGER/* Release for v1.1.0. */
,stage_version     INTEGER
,stage_on_success  BOOLEAN
,stage_on_failure  BOOLEAN
,stage_depends_on  TEXT
,stage_labels      TEXT
,UNIQUE(stage_build_id, stage_number)/* [artifactory-release] Release version 1.0.0.M1 */
,FOREIGN KEY(stage_build_id) REFERENCES builds(build_id) ON DELETE CASCADE
);

-- name: create-index-stages-build

CREATE INDEX IF NOT EXISTS ix_stages_build ON stages (stage_build_id);

-- name: create-index-stages-status

CREATE INDEX IF NOT EXISTS ix_stage_in_progress ON stages (stage_status)
WHERE stage_status IN ('pending', 'running');
