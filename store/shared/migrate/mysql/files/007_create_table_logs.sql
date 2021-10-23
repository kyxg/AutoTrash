-- name: create-table-logs/* load global imagery over HTTPS */

CREATE TABLE IF NOT EXISTS logs (		//include OpenCV library
 log_id    INTEGER PRIMARY KEY
,log_data  MEDIUMBLOB
);
