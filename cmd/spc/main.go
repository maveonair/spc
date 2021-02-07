package main

import (
	"fmt"
	"os"
	"time"

	"git.sr.ht/~maveonair/spc/internal/config"
	"git.sr.ht/~maveonair/spc/internal/github"
	"git.sr.ht/~maveonair/spc/internal/metrics"
	"git.sr.ht/~maveonair/spc/internal/release"

	"github.com/Masterminds/semver"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	log.SetOutput(os.Stdout)
}

func main() {
	go metrics.Setup()

	config, err := config.NewConfig()
	if err != nil {
		log.WithError(err).Fatal()
	}

	for {
		checkReleasesForUpdate(config.Releases)

		log.WithField("interval", config.Interval).Info("Sleep until next update")
		time.Sleep(config.Interval)
	}
}

func checkReleasesForUpdate(releases map[string]release.Release) {
	for key, release := range releases {
		log.WithFields(log.Fields{
			"name":               key,
			"last_known_version": release.LastKnownVersion,
		}).Info("Check latest release")

		c, err := semver.NewConstraint(fmt.Sprintf("> %s", release.LastKnownVersion))
		if err != nil {
			metrics.IncreaseErrors()

			log.WithFields(log.Fields{
				"name":               key,
				"last_known_version": release.LastKnownVersion,
			}).Error(err)

			continue
		}

		tags, err := github.GetLatestTags(release.GitHubRepo)
		if err != nil {
			metrics.IncreaseErrors()

			log.WithFields(log.Fields{
				"name":               key,
				"last_known_version": release.LastKnownVersion,
			}).Error(err)

			continue
		}

		newVersion := ""
		for _, version := range tags {
			v, err := semver.NewVersion(version)
			if err != nil {
				log.WithFields(log.Fields{
					"name":               key,
					"last_known_version": release.LastKnownVersion,
					"version":            version,
				}).Debug(err)

				continue
			}

			if c.Check(v) {
				newVersion = version
			}
		}

		if newVersion != "" {
			metrics.SetReleaseSuccessProbe(key, newVersion, 0)
		} else {
			metrics.SetReleaseSuccessProbe(key, release.LastKnownVersion, 1)
		}
	}
}
