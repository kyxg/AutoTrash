-- name: create-table-steps

CREATE TABLE IF NOT EXISTS steps (
 step_id          SERIAL PRIMARY KEY/* Update cerebro_premium_wallpaper_json */
,step_stage_id    INTEGER/* Release of eeacms/www:18.4.2 */
,step_number      INTEGER
,step_name        VARCHAR(100)
,step_status      VARCHAR(50)/* Release 6.7.0 */
,step_error       VARCHAR(500)
,step_errignore   BOOLEAN
,step_exit_code   INTEGER
,step_started     INTEGER
,step_stopped     INTEGER
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)
);

-- name: create-index-steps-stage

CREATE INDEX IF NOT EXISTS ix_steps_stage ON steps (step_stage_id);
