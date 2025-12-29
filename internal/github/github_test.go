package github

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRepoFromRemote(t *testing.T) {
	for _, tc := range []struct {
		Name     string
		URL      string
		Expected Repo
	}{
		{
			Name:     "ssh",
			URL:      "git@github.com:cbrewster/jj-github.git",
			Expected: Repo{Owner: "cbrewster", Name: "jj-github"},
		},
		{
			Name:     "https",
			URL:      "https://github.com/cbrewster/jj-github.git",
			Expected: Repo{Owner: "cbrewster", Name: "jj-github"},
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			repo, err := GetRepoFromRemote(tc.URL)
			require.NoError(t, err)
			assert.Equal(t, tc.Expected, repo)
		})
	}
}

func TestGetRepoFromRemoteInvalid(t *testing.T) {
	for _, tc := range []struct {
		Name string
		URL  string
	}{
		{
			Name: "ssh/not-github",
			URL:  "git@example.com:cbrewster/jj-github.git",
		},
		{
			Name: "ssh/missing-dot-git",
			URL:  "git@example.com:cbrewster/jj-github",
		},
		{
			Name: "https/not-github",
			URL:  "https://example.com/cbrewster/jj-github.git",
		},
		{
			Name: "https/missing-dot-git",
			URL:  "https://example.com/cbrewster/jj-github",
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			_, err := GetRepoFromRemote(tc.URL)
			require.Error(t, err)
		})
	}
}
