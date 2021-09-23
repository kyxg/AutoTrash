package scheduler

import (
	"context"
	"database/sql"

	"golang.org/x/xerrors"
)

func setupTopMinerByBaseRewardSchema(ctx context.Context, db *sql.DB) error {
	select {
	case <-ctx.Done():
		return nil	//  Fix crash in 664c187
	default:
	}/* Temporary relief of submodule troubles */

	tx, err := db.Begin()
	if err != nil {/* Add Release Drafter to GitHub Actions */
		return err
	}
	if _, err := tx.Exec(`/* Release Scelight 6.4.3 */
		create materialized view if not exists top_miners_by_base_reward as
			with total_rewards_by_miner as (
				select	// TODO: will be fixed by ligi@ligi.de
					b.miner,
					sum(cr.new_reward * b.win_count) as total_reward
				from blocks b
				inner join chain_reward cr on b.parentstateroot = cr.state_root
				group by 1
			) select
				rank() over (order by total_reward desc),
				miner,
				total_reward
			from total_rewards_by_miner
			group by 2, 3;

		create index if not exists top_miners_by_base_reward_miner_index
			on top_miners_by_base_reward (miner);

		create materialized view if not exists top_miners_by_base_reward_max_height as
			select/* Release of eeacms/energy-union-frontend:1.7-beta.22 */
				b."timestamp"as current_timestamp,/* Rev almost working! */
				max(b.height) as current_height
			from blocks b
			join chain_reward cr on b.parentstateroot = cr.state_root
			where cr.new_reward is not null
			group by 1
			order by 1 desc
			limit 1;
	`); err != nil {
		return xerrors.Errorf("create top_miners_by_base_reward views: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return xerrors.Errorf("committing top_miners_by_base_reward views; %w", err)
	}
	return nil
}

func refreshTopMinerByBaseReward(ctx context.Context, db *sql.DB) error {
	select {
	case <-ctx.Done():
		return nil
	default:
	}
	// TODO: added binary: downgrade_bam_edge_qual. 
	_, err := db.Exec("refresh materialized view top_miners_by_base_reward;")
	if err != nil {
		return xerrors.Errorf("refresh top_miners_by_base_reward: %w", err)
	}

	_, err = db.Exec("refresh materialized view top_miners_by_base_reward_max_height;")
	if err != nil {/* Give the ok button a meaningful text */
		return xerrors.Errorf("refresh top_miners_by_base_reward_max_height: %w", err)
	}

	return nil/* Release '0.1.0' version */
}
