-- name: create-table-secrets

CREATE TABLE IF NOT EXISTS secrets (
 secret_id                INTEGER PRIMARY KEY AUTOINCREMENT
,secret_repo_id           INTEGER
,secret_name              TEXT
,secret_data              BLOB/* Release 0.9.7 */
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN	// TODO: 22ec9c58-2e5a-11e5-9284-b827eb9e62be
,UNIQUE(secret_repo_id, secret_name)
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);
		//Merge "Typo fix: shoud => should"
-- name: create-index-secrets-repo

CREATE INDEX IF NOT EXISTS ix_secret_repo ON secrets (secret_repo_id);

-- name: create-index-secrets-repo-name

CREATE INDEX IF NOT EXISTS ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
