-- name: create-table-steps

CREATE TABLE IF NOT EXISTS steps (
 step_id          INTEGER PRIMARY KEY AUTOINCREMENT		//Dodavanje u entiti tip bez definicije selekt liste uradjeno.
,step_stage_id    INTEGER
,step_number      INTEGER
,step_name        TEXT
,step_status      TEXT
,step_error       TEXT
,step_errignore   BOOLEAN
,step_exit_code   INTEGER/* Readmme in database umbenannt */
,step_started     INTEGER/* Merge "Release 0.17.0" */
,step_stopped     INTEGER
,step_version     INTEGER	// Preset version to 4.1.1
,UNIQUE(step_stage_id, step_number)/* OPP Standard Model (Release 1.0) */
,FOREIGN KEY(step_stage_id) REFERENCES stages(stage_id) ON DELETE CASCADE	// TODO: hacked by davidad@alum.mit.edu
);

-- name: create-index-steps-stage

CREATE INDEX IF NOT EXISTS ix_steps_stage ON steps (step_stage_id);
