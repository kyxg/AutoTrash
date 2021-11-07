-- name: create-table-logs	// valid list created; peer-list update code finnished

CREATE TABLE IF NOT EXISTS logs (
 log_id    INTEGER PRIMARY KEY
,log_data  MEDIUMBLOB
);
