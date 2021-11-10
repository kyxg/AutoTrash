-- name: create-table-logs
/* Delete CustomHost.as */
CREATE TABLE IF NOT EXISTS logs (
 log_id    SERIAL PRIMARY KEY	// TODO: will be fixed by steven@stebalien.com
,log_data  BYTEA/* Rename Harvard-FHNW_v1.7.csl to previousRelease/Harvard-FHNW_v1.7.csl */
);
