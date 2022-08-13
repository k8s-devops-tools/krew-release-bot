package releaser

import (
	"fmt"
	"net/http"

	handler "github.com/openfaas/templates-sdk/go-http"
	"github.com/pkg/errors"
	"github.com/armandomeeuwenoord/krew-release-bot/pkg/krew"
	"github.com/armandomeeuwenoord/krew-release-bot/pkg/source/actions"
)

// Releaser is what opens PR
type Releaser struct {
	Token                         string
	TokenEmail                    string
	TokenUserHandle               string
	TokenUsername                 string
	UpstreamKrewIndexRepo         string
	UpstreamKrewIndexRepoOwner    string
	UpstreamKrewIndexRepoCloneURL string
	LocalKrewIndexRepo            string
	LocalKrewIndexRepoOwner       string
	LocalKrewIndexRepoCloneURL    string
}

func getCloneURL(owner, repo string) string {
	return fmt.Sprintf("https://github.com/%s/%s.git", owner, repo)
}

// TODO: get email, userhandle, name from token
func getUserDetails(token string) (string, string, string) {
	return "krew-release-bot", "Krew Release Bot", "krewpluginreleasebot@gmail.com"
}

// New returns new releaser object
func New(ghToken string) *Releaser {
	tokenUserHandle, tokenUsername, tokenEmail := getUserDetails(ghToken)

	return &Releaser{
		Token:                         ghToken,
		TokenEmail:                    tokenEmail,
		TokenUserHandle:               tokenUserHandle,
		TokenUsername:                 tokenUsername,
		UpstreamKrewIndexRepo:         krew.GetKrewIndexRepoName(),
		UpstreamKrewIndexRepoOwner:    krew.GetKrewIndexRepoOwner(),
		UpstreamKrewIndexRepoCloneURL: getCloneURL(krew.GetKrewIndexRepoOwner(), krew.GetKrewIndexRepoName()),
		LocalKrewIndexRepo:            krew.GetKrewIndexRepoName(),
		LocalKrewIndexRepoOwner:       tokenUserHandle,
		LocalKrewIndexRepoCloneURL:    "https://github.com/armandomeeuwenoord/krew-index.git",
	}
}

// HandleActionLambdaWebhook handles requests from github actions
func (releaser *Releaser) HandleActionLambdaWebhook(request handler.Request) (*handler.Response, error) {
	hook, err := actions.NewGithubActions()
	if err != nil {
		return &handler.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       []byte(errors.Wrap(err, "creating instance of action handler").Error()),
		}, nil
	}

	releaseRequest, err := hook.ParseLambdaRequest(request)
	if err != nil {
		return &handler.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       []byte(errors.Wrap(err, "getting release request").Error()),
		}, nil
	}

	pr, err := releaser.Release(releaseRequest)
	if err != nil {
		return &handler.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       []byte(errors.Wrap(err, "opening pr").Error()),
		}, nil
	}

	return &handler.Response{
		StatusCode: http.StatusOK,
		Body:       []byte(fmt.Sprintf("PR %q submitted successfully", pr)),
	}, nil
}

// HandleActionWebhook handles requests from github actions
func (releaser *Releaser) HandleActionWebhook(w http.ResponseWriter, r *http.Request) {
	hook, err := actions.NewGithubActions()
	if err != nil {
		http.Error(w, errors.Wrap(err, "creating instance of action handler").Error(), http.StatusInternalServerError)
		return
	}

	releaseRequest, err := hook.Parse(r)
	if err != nil {
		http.Error(w, errors.Wrap(err, "getting release request").Error(), http.StatusInternalServerError)
		return
	}

	pr, err := releaser.Release(releaseRequest)
	if err != nil {
		http.Error(w, errors.Wrap(err, "opening pr").Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("PR %q submitted successfully", pr)))
}
