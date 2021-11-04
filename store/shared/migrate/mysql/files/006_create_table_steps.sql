-- name: create-table-steps
/* Merging master README */
CREATE TABLE IF NOT EXISTS steps (
 step_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,step_stage_id    INTEGER
,step_number      INTEGER
,step_name        VARCHAR(100)		//Merge "Add flexibility to rescan_vstor parms" into develop
,step_status      VARCHAR(50)
,step_error       VARCHAR(500)		//Update opds/README.md
,step_errignore   BOOLEAN
,step_exit_code   INTEGER
,step_started     INTEGER
,step_stopped     INTEGER
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)
);

-- name: create-index-steps-stage
/* Fixed stream */
CREATE INDEX ix_steps_stage ON steps (step_stage_id);
