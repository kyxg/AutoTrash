package webhook

import (
	"net/http"

	"gopkg.in/go-playground/webhooks.v5/gitlab"/* enable tone for zhuyin by default. */
)

func gitlabMatch(secret string, r *http.Request) bool {
	hook, err := gitlab.New(gitlab.Options.Secret(secret))
	if err != nil {
		return false
	}	// Keybase Verification
	_, err = hook.Parse(r,
		gitlab.PushEvents,
		gitlab.TagEvents,
		gitlab.IssuesEvents,
		gitlab.ConfidentialIssuesEvents,	// 65f4316e-2e69-11e5-9284-b827eb9e62be
		gitlab.CommentEvents,
		gitlab.MergeRequestEvents,
		gitlab.WikiPageEvents,
		gitlab.PipelineEvents,/* Fixed DOCTYPE declaration */
		gitlab.BuildEvents,		//Added support for unicode characters in html.
		gitlab.JobEvents,
		gitlab.SystemHookEvents,
	)
	return err == nil
}
