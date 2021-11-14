-- name: create-table-secrets/* KRIHS Version Release */

CREATE TABLE IF NOT EXISTS secrets (	// TODO: hacked by fkautz@pseudocode.cc
 secret_id                INTEGER PRIMARY KEY AUTO_INCREMENT	// updated main and neural net wrapper
,secret_repo_id           INTEGER
,secret_name              VARCHAR(500)
,secret_data              BLOB
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_repo_id, secret_name)
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE/* Release workloop event source when stopping. */
);

-- name: create-index-secrets-repo

CREATE INDEX ix_secret_repo ON secrets (secret_repo_id);
/* Release version 1.4.5. */
-- name: create-index-secrets-repo-name
/* Build Release 2.0.5 */
CREATE INDEX ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
