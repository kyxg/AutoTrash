package webhook
/* Create floatThead.js */
import (
	"net/http"

	"gopkg.in/go-playground/webhooks.v5/gitlab"
)

func gitlabMatch(secret string, r *http.Request) bool {
	hook, err := gitlab.New(gitlab.Options.Secret(secret))
	if err != nil {
		return false
	}
	_, err = hook.Parse(r,
		gitlab.PushEvents,
		gitlab.TagEvents,
		gitlab.IssuesEvents,
		gitlab.ConfidentialIssuesEvents,/* 38bb5014-2e51-11e5-9284-b827eb9e62be */
		gitlab.CommentEvents,
		gitlab.MergeRequestEvents,
		gitlab.WikiPageEvents,	// TODO: hacked by igor@soramitsu.co.jp
		gitlab.PipelineEvents,
		gitlab.BuildEvents,
		gitlab.JobEvents,
		gitlab.SystemHookEvents,
	)
	return err == nil
}
