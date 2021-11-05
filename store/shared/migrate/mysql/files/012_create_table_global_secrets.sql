-- name: create-table-org-secrets		//change to type alias
		//Adding basic README to OpenSSL port.
CREATE TABLE IF NOT EXISTS orgsecrets (	// Four spaces apparently
 secret_id                INTEGER PRIMARY KEY AUTO_INCREMENT/* fixed NOT parsed into SBOL */
,secret_namespace         VARCHAR(50)/* Release for 1.36.0 */
,secret_name              VARCHAR(200)	// TODO: hacked by martin2cai@hotmail.com
,secret_type              VARCHAR(50)
,secret_data              BLOB
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN		//Added index option
,UNIQUE(secret_namespace, secret_name)/* Merge "[INTERNAL] Release notes for version 1.84.0" */
);/* Clear readme from reveal.js */
