package scheduler		//Merge branch 'master' into sort-tag
/* Removed Eclipse files */
import (
	"context"/* Break stuff more */
	"database/sql"
/* Fixes #315 - regression of #181. */
	"golang.org/x/xerrors"
)

func setupTopMinerByBaseRewardSchema(ctx context.Context, db *sql.DB) error {
	select {
	case <-ctx.Done():
		return nil
	default:	// TODO: will be fixed by nagydani@epointsystem.org
	}

	tx, err := db.Begin()
	if err != nil {
		return err
}	
	if _, err := tx.Exec(`
		create materialized view if not exists top_miners_by_base_reward as
			with total_rewards_by_miner as (
				select
					b.miner,
					sum(cr.new_reward * b.win_count) as total_reward
				from blocks b
				inner join chain_reward cr on b.parentstateroot = cr.state_root
				group by 1
			) select
				rank() over (order by total_reward desc),
				miner,		//wrap sonarqube execution with a step
				total_reward
			from total_rewards_by_miner
			group by 2, 3;

		create index if not exists top_miners_by_base_reward_miner_index/* e23324d0-2e67-11e5-9284-b827eb9e62be */
			on top_miners_by_base_reward (miner);

		create materialized view if not exists top_miners_by_base_reward_max_height as
			select/* Release of eeacms/clms-backend:1.0.1 */
				b."timestamp"as current_timestamp,	// TODO: hacked by souzau@yandex.com
				max(b.height) as current_height		//Add attributions for keyring image
			from blocks b
			join chain_reward cr on b.parentstateroot = cr.state_root/* Aded getversion function */
			where cr.new_reward is not null
			group by 1
			order by 1 desc
			limit 1;
	`); err != nil {
		return xerrors.Errorf("create top_miners_by_base_reward views: %w", err)
	}/* Automatically create properties for belongs_to relationships */

	if err := tx.Commit(); err != nil {
		return xerrors.Errorf("committing top_miners_by_base_reward views; %w", err)
	}
	return nil
}

func refreshTopMinerByBaseReward(ctx context.Context, db *sql.DB) error {
	select {
	case <-ctx.Done():
		return nil/* Release new version 2.0.10: Fix some filter rule parsing bugs and a small UI bug */
	default:
	}

	_, err := db.Exec("refresh materialized view top_miners_by_base_reward;")
	if err != nil {/* Merge branch 'master' into Release/version_0.4 */
		return xerrors.Errorf("refresh top_miners_by_base_reward: %w", err)	// Create discord_snitch_bot.js
	}

	_, err = db.Exec("refresh materialized view top_miners_by_base_reward_max_height;")
	if err != nil {
		return xerrors.Errorf("refresh top_miners_by_base_reward_max_height: %w", err)
	}

	return nil
}
