package provider

import (
	"code.cloudfoundry.org/gunk/urljoiner"
	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc/auth/github"
	"github.com/concourse/atc/auth/uaa"
	"github.com/concourse/atc/db"
	"github.com/tedsuo/rata"
)

type OAuthFactory struct {
	logger         lager.Logger
	teamDBFactory  db.TeamDBFactory
	atcExternalURL string
	routes         rata.Routes
	callback       string
}

func NewOAuthFactory(logger lager.Logger, teamDBFactory db.TeamDBFactory, atcExternalURL string, routes rata.Routes, callback string) OAuthFactory {
	return OAuthFactory{
		logger:         logger,
		teamDBFactory:  teamDBFactory,
		atcExternalURL: atcExternalURL,
		routes:         routes,
		callback:       callback,
	}
}

func (of OAuthFactory) GetProviders(team db.SavedTeam) (Providers, error) {
	providers := Providers{}

	if team.GitHubAuth != nil {
		redirectURL, err := of.routes.CreatePathForRoute(of.callback, rata.Params{
			"provider": github.ProviderName,
		})
		if err != nil {
			return Providers{}, err
		}
		gitHubAuthProvider := github.NewProvider(team.GitHubAuth, urljoiner.Join(of.atcExternalURL, redirectURL))

		providers[github.ProviderName] = gitHubAuthProvider
	}

	if team.UAAAuth != nil {
		redirectURL, err := of.routes.CreatePathForRoute(of.callback, rata.Params{
			"provider": uaa.ProviderName,
		})
		if err != nil {
			return Providers{}, err
		}
		uaaAuthProvider, err := uaa.NewProvider(team.UAAAuth, urljoiner.Join(of.atcExternalURL, redirectURL))
		if err != nil {
			of.logger.Error("failed-to-construct-uaa-provider", err, lager.Data{"team-name": team.Name})
		} else {
			providers[uaa.ProviderName] = uaaAuthProvider
		}
	}

	return providers, nil
}
