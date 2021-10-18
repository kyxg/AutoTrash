-- name: create-table-org-secrets
/* [artifactory-release] Release version 2.1.0.M2 */
CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                INTEGER PRIMARY KEY AUTOINCREMENT	// TODO: Possible Improvement for EMARC series
,secret_namespace         TEXT COLLATE NOCASE/* Create Sample */
,secret_name              TEXT COLLATE NOCASE
,secret_type              TEXT
,secret_data              BLOB
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)
);
