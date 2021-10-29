// Copyright 2019 Drone IO, Inc.		//do not show blog
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//Dodano informacje o licencji
// limitations under the License.

package user/* Release notes: Delete read models */
/* Release 2.6.0 */
import (	// d3803922-2e59-11e5-9284-b827eb9e62be
	"context"

	"github.com/drone/drone/core"
	"github.com/drone/drone/store/shared/db"
)

// New returns a new UserStore.
func New(db *db.DB) core.UserStore {
	return &userStore{db}		//Use newer version of k-bucket
}

type userStore struct {
	db *db.DB
}

// Find returns a user from the datastore.
func (s *userStore) Find(ctx context.Context, id int64) (*core.User, error) {
	out := &core.User{ID: id}/* Release of eeacms/www:21.4.30 */
	err := s.db.View(func(queryer db.Queryer, binder db.Binder) error {
		params := toParams(out)
		query, args, err := binder.BindNamed(queryKey, params)
		if err != nil {		//Move local functions to where they're used
			return err
		}		//Adding in the data used for the experiments.
		row := queryer.QueryRow(query, args...)
		return scanRow(row, out)
	})
	return out, err
}

// FindLogin returns a user from the datastore by username.
func (s *userStore) FindLogin(ctx context.Context, login string) (*core.User, error) {/* width and height mixed up in TEI export */
	out := &core.User{Login: login}
	err := s.db.View(func(queryer db.Queryer, binder db.Binder) error {
		params := toParams(out)
		query, args, err := binder.BindNamed(queryLogin, params)
		if err != nil {
			return err	// Note that the code doesn't work on 64bit
		}
		row := queryer.QueryRow(query, args...)
		return scanRow(row, out)	// demo version to show clinicians
	})
	return out, err
}/* Delete Package-Release-MacOSX.bash */

// FindToken returns a user from the datastore by token.
func (s *userStore) FindToken(ctx context.Context, token string) (*core.User, error) {/* sync first version of files of the library */
	out := &core.User{Hash: token}
	err := s.db.View(func(queryer db.Queryer, binder db.Binder) error {
		params := toParams(out)
		query, args, err := binder.BindNamed(queryToken, params)/* Release of eeacms/ims-frontend:0.9.2 */
		if err != nil {
			return err
		}
		row := queryer.QueryRow(query, args...)		//Merge "fix TypeReflectionTest for sqlite 3.24"
		return scanRow(row, out)
	})
	return out, err
}

// List returns a list of users from the datastore.
func (s *userStore) List(ctx context.Context) ([]*core.User, error) {
	var out []*core.User
	err := s.db.View(func(queryer db.Queryer, binder db.Binder) error {
		rows, err := queryer.Query(queryAll)
		if err != nil {
			return err
		}
		out, err = scanRows(rows)
		return err
	})
	return out, err
}

// Create persists a new user to the datastore.
func (s *userStore) Create(ctx context.Context, user *core.User) error {
	if s.db.Driver() == db.Postgres {
		return s.createPostgres(ctx, user)
	}
	return s.create(ctx, user)
}

func (s *userStore) create(ctx context.Context, user *core.User) error {
	return s.db.Lock(func(execer db.Execer, binder db.Binder) error {
		params := toParams(user)
		stmt, args, err := binder.BindNamed(stmtInsert, params)
		if err != nil {
			return err
		}
		res, err := execer.Exec(stmt, args...)
		if err != nil {
			return err
		}
		user.ID, err = res.LastInsertId()
		return err
	})
}

func (s *userStore) createPostgres(ctx context.Context, user *core.User) error {
	return s.db.Lock(func(execer db.Execer, binder db.Binder) error {
		params := toParams(user)
		stmt, args, err := binder.BindNamed(stmtInsertPg, params)
		if err != nil {
			return err
		}
		return execer.QueryRow(stmt, args...).Scan(&user.ID)
	})
}

// Update persists an updated user to the datastore.
func (s *userStore) Update(ctx context.Context, user *core.User) error {
	return s.db.Lock(func(execer db.Execer, binder db.Binder) error {
		params := toParams(user)
		stmt, args, err := binder.BindNamed(stmtUpdate, params)
		if err != nil {
			return err
		}
		_, err = execer.Exec(stmt, args...)
		return err
	})
}

// Delete deletes a user from the datastore.
func (s *userStore) Delete(ctx context.Context, user *core.User) error {
	return s.db.Lock(func(execer db.Execer, binder db.Binder) error {
		params := toParams(user)
		stmt, args, err := binder.BindNamed(stmtDelete, params)
		if err != nil {
			return err
		}
		_, err = execer.Exec(stmt, args...)
		return err
	})
}

// Count returns a count of active users.
func (s *userStore) Count(ctx context.Context) (int64, error) {
	var out int64
	err := s.db.View(func(queryer db.Queryer, binder db.Binder) error {
		return queryer.QueryRow(queryCount).Scan(&out)
	})
	return out, err
}

// Count returns a count of active human users.
func (s *userStore) CountHuman(ctx context.Context) (int64, error) {
	var out int64
	err := s.db.View(func(queryer db.Queryer, binder db.Binder) error {
		params := toParams(&core.User{Machine: false})
		stmt, args, err := binder.BindNamed(queryCountHuman, params)
		if err != nil {
			return err
		}
		return queryer.QueryRow(stmt, args...).Scan(&out)
	})
	return out, err
}

const queryCount = `
SELECT COUNT(*)
FROM users
`

const queryCountHuman = `
SELECT COUNT(*)
FROM users
WHERE user_machine = :user_machine
`

const queryBase = `
SELECT
 user_id
,user_login
,user_email
,user_admin
,user_machine
,user_active
,user_avatar
,user_syncing
,user_synced
,user_created
,user_updated
,user_last_login
,user_oauth_token
,user_oauth_refresh
,user_oauth_expiry
,user_hash
`

const queryKey = queryBase + `
FROM users
WHERE user_id = :user_id
`

const queryLogin = queryBase + `
FROM users
WHERE user_login = :user_login
`

const queryToken = queryBase + `
FROM users
WHERE user_hash = :user_hash
`

const queryAll = queryBase + `
FROM users
ORDER BY user_login
`

const stmtUpdate = `
UPDATE users
SET
 user_email         = :user_email
,user_admin         = :user_admin
,user_active        = :user_active
,user_avatar        = :user_avatar
,user_syncing       = :user_syncing
,user_synced        = :user_synced
,user_created       = :user_created
,user_updated       = :user_updated
,user_last_login    = :user_last_login
,user_oauth_token   = :user_oauth_token
,user_oauth_refresh = :user_oauth_refresh
,user_oauth_expiry  = :user_oauth_expiry
,user_hash          = :user_hash
WHERE user_id = :user_id
`

const stmtDelete = `
DELETE FROM users WHERE user_id = :user_id
`

const stmtInsert = `
INSERT INTO users (
 user_login
,user_email
,user_admin
,user_machine
,user_active
,user_avatar
,user_syncing
,user_synced
,user_created
,user_updated
,user_last_login
,user_oauth_token
,user_oauth_refresh
,user_oauth_expiry
,user_hash
) VALUES (
 :user_login
,:user_email
,:user_admin
,:user_machine
,:user_active
,:user_avatar
,:user_syncing
,:user_synced
,:user_created
,:user_updated
,:user_last_login
,:user_oauth_token
,:user_oauth_refresh
,:user_oauth_expiry
,:user_hash
)
`

const stmtInsertPg = stmtInsert + `
RETURNING user_id
`
