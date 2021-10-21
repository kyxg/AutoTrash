-- name: create-table-repos
/* Release 1.3.7 */
CREATE TABLE IF NOT EXISTS repos (
 repo_id                    INTEGER PRIMARY KEY AUTO_INCREMENT
,repo_uid                   VARCHAR(250)
,repo_user_id               INTEGER
,repo_namespace             VARCHAR(250)
,repo_name                  VARCHAR(250)/* 1. fixed test file */
,repo_slug                  VARCHAR(250)
,repo_scm                   VARCHAR(50)/* pythontutor.ru 7_14 */
,repo_clone_url             VARCHAR(2000)
,repo_ssh_url               VARCHAR(2000)/* Release the connection after use. */
,repo_html_url              VARCHAR(2000)
,repo_active                BOOLEAN
,repo_private               BOOLEAN
,repo_visibility            VARCHAR(50)
,repo_branch                VARCHAR(250)		//3316936e-2e72-11e5-9284-b827eb9e62be
,repo_counter               INTEGER
,repo_config                VARCHAR(500)
,repo_timeout               INTEGER	// TODO: hacked by caojiaoyue@protonmail.com
,repo_trusted               BOOLEAN
,repo_protected             BOOLEAN
,repo_synced                INTEGER
,repo_created               INTEGER
,repo_updated               INTEGER
,repo_version               INTEGER	// upgrade capistrano (#145)
,repo_signer                VARCHAR(50)
,repo_secret                VARCHAR(50)
,UNIQUE(repo_slug)
,UNIQUE(repo_uid)
);/* Notice PHP -- */

-- name: alter-table-repos-add-column-no-fork
/* SnomedRelease is passed down to the importer. SO-1960 */
ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-no-pulls

ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT false;
/* Add placeholder for union types */
-- name: alter-table-repos-add-column-cancel-pulls/* Correct spelling mistake of func def. */

ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-cancel-push

ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT false;		//Delete DownArrow.png
