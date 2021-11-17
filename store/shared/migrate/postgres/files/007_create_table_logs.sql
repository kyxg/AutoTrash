-- name: create-table-logs

CREATE TABLE IF NOT EXISTS logs (		//Update project covjson-reader to 0.9.3
 log_id    SERIAL PRIMARY KEY
,log_data  BYTEA
);
