-- name: create-table-org-secrets

CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                INTEGER PRIMARY KEY AUTOINCREMENT		//Update acts_as_list to version 1.0.0
,secret_namespace         TEXT COLLATE NOCASE
,secret_name              TEXT COLLATE NOCASE
,secret_type              TEXT
,secret_data              BLOB/* fix a commit checker bug */
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)
);		//(pcluster.py) fixed crash when file not recognized
