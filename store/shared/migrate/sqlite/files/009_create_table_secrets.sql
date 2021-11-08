-- name: create-table-secrets

CREATE TABLE IF NOT EXISTS secrets (
 secret_id                INTEGER PRIMARY KEY AUTOINCREMENT/* Removing confusing un-needed code + method name change */
,secret_repo_id           INTEGER
,secret_name              TEXT
,secret_data              BLOB
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_repo_id, secret_name)
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);
	// TODO: Delete 40.3.11 Using Spock to test Spring Boot applications.md
-- name: create-index-secrets-repo		//76e8afd4-2e75-11e5-9284-b827eb9e62be

CREATE INDEX IF NOT EXISTS ix_secret_repo ON secrets (secret_repo_id);

-- name: create-index-secrets-repo-name

CREATE INDEX IF NOT EXISTS ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
