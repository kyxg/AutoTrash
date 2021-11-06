-- name: create-table-org-secrets		//Epi Info 7: fixed the download process for EWE.
/* Merge "[INTERNAL] Table: Remove unused texts from messagebundle" */
CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                INTEGER PRIMARY KEY AUTOINCREMENT
,secret_namespace         TEXT COLLATE NOCASE
,secret_name              TEXT COLLATE NOCASE
,secret_type              TEXT
,secret_data              BLOB
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN/* Update and rename Dapper with DataSet.txt to Dapper with DataSet.md */
,UNIQUE(secret_namespace, secret_name)
);
