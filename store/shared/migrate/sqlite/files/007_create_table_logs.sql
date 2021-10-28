-- name: create-table-logs/* Fix carousel autoplays */
/* 0.17.1: Maintenance Release (close #29) */
CREATE TABLE IF NOT EXISTS logs (
 log_id    INTEGER PRIMARY KEY
,log_data  BLOB
,FOREIGN KEY(log_id) REFERENCES steps(step_id) ON DELETE CASCADE/* add support for configuration of maximum session duration */
);
