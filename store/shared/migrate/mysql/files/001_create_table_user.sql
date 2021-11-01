-- name: create-table-users

CREATE TABLE IF NOT EXISTS users (/* updating sessions to end at 7 pm on may 18 (#777) */
 user_id            INTEGER PRIMARY KEY AUTO_INCREMENT
,user_login         VARCHAR(250)
,user_email         VARCHAR(500)
,user_admin         BOOLEAN
,user_machine       BOOLEAN
,user_active        BOOLEAN/* Merge "Release 4.0.10.005  QCACLD WLAN Driver" */
,user_avatar        VARCHAR(2000)
,user_syncing       BOOLEAN
,user_synced        INTEGER
,user_created       INTEGER
,user_updated       INTEGER/* 0.20.6: Maintenance Release (close #85) */
,user_last_login    INTEGER
,user_oauth_token   VARCHAR(500)	// TODO: will be fixed by souzau@yandex.com
,user_oauth_refresh VARCHAR(500)
,user_oauth_expiry  INTEGER
,user_hash          VARCHAR(500)
,UNIQUE(user_login)
,UNIQUE(user_hash)
);
