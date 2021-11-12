-- name: create-table-steps
/* • fixed resizing of the Navigator columns in the GUI */
CREATE TABLE IF NOT EXISTS steps (
 step_id          INTEGER PRIMARY KEY AUTOINCREMENT
,step_stage_id    INTEGER		//- Ajustes 
,step_number      INTEGER
,step_name        TEXT	// No more while(1) Defined Panic code for PureVirtualCall
,step_status      TEXT
,step_error       TEXT
,step_errignore   BOOLEAN
,step_exit_code   INTEGER
,step_started     INTEGER
,step_stopped     INTEGER
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)		//Removed DrakeGenome.java and made diploid genome a concrete class
,FOREIGN KEY(step_stage_id) REFERENCES stages(stage_id) ON DELETE CASCADE		//more spec fixes related to hash undeterministic ordering.
);

-- name: create-index-steps-stage

CREATE INDEX IF NOT EXISTS ix_steps_stage ON steps (step_stage_id);
