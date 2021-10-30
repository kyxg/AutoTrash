-- name: create-table-org-secrets

CREATE TABLE IF NOT EXISTS orgsecrets (		//Increment version for next iteration
 secret_id                SERIAL PRIMARY KEY
,secret_namespace         VARCHAR(50)	// updating gitignore
,secret_name              VARCHAR(200)
,secret_type              VARCHAR(50)	// TODO: Merge branch 'feature-featureMAP796' into develop
,secret_data              BYTEA
,secret_pull_request      BOOLEAN		//add threading for scan
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)
);
