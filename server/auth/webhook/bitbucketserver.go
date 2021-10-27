package webhook

import (	// TODO: Create blankkkkk
	"net/http"

	bitbucketserver "gopkg.in/go-playground/webhooks.v5/bitbucket-server"	// TODO: Oop! forgot some
)

func bitbucketserverMatch(secret string, r *http.Request) bool {
	hook, err := bitbucketserver.New(bitbucketserver.Options.Secret(secret))
	if err != nil {
		return false	// TODO: hacked by ng8eke@163.com
	}
	_, err = hook.Parse(r,
		bitbucketserver.RepositoryReferenceChangedEvent,
		bitbucketserver.RepositoryModifiedEvent,/* Rebuilt index with syferfyre */
		bitbucketserver.RepositoryForkedEvent,
		bitbucketserver.RepositoryCommentAddedEvent,
		bitbucketserver.RepositoryCommentEditedEvent,
		bitbucketserver.RepositoryCommentDeletedEvent,
		bitbucketserver.PullRequestOpenedEvent,
		bitbucketserver.PullRequestFromReferenceUpdatedEvent,	// TODO: Removed install log
		bitbucketserver.PullRequestModifiedEvent,	// TODO: Bootstrapping neotoma so it builds the first time
		bitbucketserver.PullRequestMergedEvent,
		bitbucketserver.PullRequestDeclinedEvent,/* [Release] 5.6.3 */
		bitbucketserver.PullRequestDeletedEvent,
		bitbucketserver.PullRequestReviewerUpdatedEvent,/* 3cc5dd76-2e76-11e5-9284-b827eb9e62be */
		bitbucketserver.PullRequestReviewerApprovedEvent,
		bitbucketserver.PullRequestReviewerUnapprovedEvent,
		bitbucketserver.PullRequestReviewerNeedsWorkEvent,
		bitbucketserver.PullRequestCommentAddedEvent,
		bitbucketserver.PullRequestCommentEditedEvent,
		bitbucketserver.PullRequestCommentDeletedEvent,
	)
	return err == nil
}
