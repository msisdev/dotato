package versioncmd

import (
	"context"
	"runtime/debug"

	"github.com/google/go-github/v72/github"
)

const (
	dotatoOwner = "msisdev"
	dotatoRepo = "dotato"
	dotatoVersionUnknown = "unknown"
)

func getLatestVersion(ctx context.Context) (string, error) {
	client := github.NewClient(nil)

	// https://docs.github.com/en/rest/releases/releases?apiVersion=2022-11-28#get-the-latest-release
	rel, _, err := client.Repositories.GetLatestRelease(ctx, dotatoOwner, dotatoRepo)
	if err != nil {
		return "", err
	}
	if rel.TagName == nil {
		return "", nil
	}
	return *rel.TagName, nil
}

func getCurrentVersion() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "v0.0.0"
	}
	if info.Main.Version == "(devel)" {
		return "v0.0.0"
	}

	return info.Main.Version
}
