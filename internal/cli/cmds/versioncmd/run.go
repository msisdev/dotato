package versioncmd

import (
	"context"
	"time"

	"github.com/charmbracelet/log"
	"github.com/msisdev/dotato/internal/cli/args"
	"github.com/msisdev/dotato/internal/component/mxspinner"
	"github.com/msisdev/dotato/internal/lib/store"
	"golang.org/x/mod/semver"
)

const (
	timeout = time.Second * 7
)

func Run(logger *log.Logger, args *args.VersionArgs) {
	current := getCurrentVersion()
	if !semver.IsValid(current) {
		logger.Error("Failed to get current version")
		return
	}

	var latest string
	title := "Fetching verseion info ..."
	err := mxspinner.Run(title, func(store *store.Store[string], quit <-chan bool) error {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		var err error
		latest, err = getLatestVersion(ctx)
		if err != nil {
			return err
		}

		store.Set("Fetched version info")

		return nil
	})
	if err != nil {
		logger.Error("Failed to fetch latest release", "error", err)
		return
	}
	if !semver.IsValid(latest) {
		logger.Error("Failed to get latest version")
		println(renderCurrent(current))
		return
	}

	if semver.Compare(current, latest) == 0 {
		println(renderCurrentIsLatest(current))
		return
	}
	
	println(renderLatestAvailable(current, latest))

	if semver.Compare(semver.Major(current), semver.Major(latest)) == -1 {
		println("This is a major release, please check the changelog before upgrading.")
	}
}
