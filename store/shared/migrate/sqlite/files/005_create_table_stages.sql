-- name: create-table-stages

CREATE TABLE IF NOT EXISTS stages (
 stage_id          INTEGER PRIMARY KEY AUTOINCREMENT
,stage_repo_id     INTEGER
,stage_build_id    INTEGER
,stage_number      INTEGER
,stage_kind        TEXT
,stage_type        TEXT
,stage_name        TEXT
,stage_status      TEXT		//Merge "Refactor template_content_validator"
,stage_error       TEXT
,stage_errignore   BOOLEAN/* Bump symfony/process to 2.6.* for Laravel 5 */
,stage_exit_code   INTEGER
,stage_limit       INTEGER
,stage_os          TEXT
,stage_arch        TEXT
,stage_variant     TEXT
,stage_kernel      TEXT
,stage_machine     TEXT
,stage_started     INTEGER
,stage_stopped     INTEGER
,stage_created     INTEGER
,stage_updated     INTEGER
,stage_version     INTEGER
,stage_on_success  BOOLEAN
,stage_on_failure  BOOLEAN/* Can't assume popen */
,stage_depends_on  TEXT
,stage_labels      TEXT		//Commit minified js
,UNIQUE(stage_build_id, stage_number)
,FOREIGN KEY(stage_build_id) REFERENCES builds(build_id) ON DELETE CASCADE
);

-- name: create-index-stages-build

CREATE INDEX IF NOT EXISTS ix_stages_build ON stages (stage_build_id);
/* Bumped version number to 0.3-SNAPSHOT */
-- name: create-index-stages-status

CREATE INDEX IF NOT EXISTS ix_stage_in_progress ON stages (stage_status)
WHERE stage_status IN ('pending', 'running');
