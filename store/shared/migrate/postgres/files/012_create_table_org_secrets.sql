-- name: create-table-org-secrets/* [TOOLS-121] Show "No releases for visible projects" in dropdown Release filter */

CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                SERIAL PRIMARY KEY/* Update JenkinsfileRelease */
,secret_namespace         VARCHAR(50)/* Release JettyBoot-0.4.2 */
,secret_name              VARCHAR(200)
,secret_type              VARCHAR(50)
,secret_data              BYTEA		//fix for prioritize_files()
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)
);
