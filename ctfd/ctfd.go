package ctfd

import (
	"os"
	"path"

	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var categories = []string{"crypto", "web", "misc", "pwn", "rev"}

var log = zlog.Output(zerolog.ConsoleWriter{Out: os.Stderr}).Level(zerolog.InfoLevel)

type ChallengeInfo struct {
	Name        string   `yaml:"name"`
	Author      string   `yaml:"author"`
	Category    string   `yaml:"category"`
	Description string   `yaml:"description"`
	Value       int      `yaml:"value"`
	Type        string   `yaml:"type"`
	ConnInfo    string   `yaml:"connection_info"`
	Flags       []string `yaml:"flags"`
	Tags        []string `yaml:"tags"`
	Files       []string `yaml:"files"`
	Hints       []string `yaml:"hints"`
	PreReqs     []string `yaml:"requirements"`
	State       string   `yaml:"state"`
}

func readChallengeYaml(chalPath string) (ChallengeInfo, error) {
	info := ChallengeInfo{}
	challengeYaml, err := os.ReadFile(path.Join(chalPath, "challenge.yaml"))
	if err != nil {
		return info, err
	}
	if err = yaml.Unmarshal(challengeYaml, &info); err != nil {
		return info, err
	}
	return info, nil
}

func uploadDist(chalPath string) error {
	distPath := path.Join(chalPath, "dist")
	files, err := os.ReadDir(distPath)
	if err != nil {
		log.Info().Str("path", distPath).Msg("no dist found")
		return err
	}
	for _, file := range files {
		log.Info().Str("file", file.Name()).Msg("unimplemented: uploading file")
	}
	return nil
}

// syncChallenge synchronizes a challenge to CTFd given a challenge folder
func syncChallenge(chalPath string) error {
	info, err := readChallengeYaml(chalPath)
	if err != nil {
		return err
	}
	log.Info().Any("object", info).Msg("parsed challenge info")
	err = uploadDist(chalPath)
	if err != nil {
		return err
	}
	return nil
}

func Sync(c *cobra.Command, args []string) {
	// Synchronize challenges
	for _, cat := range categories {
		dir := path.Join("challenges", cat)
		log.Info().Str("path", dir).Msg("reading directory")
		challenges, err := os.ReadDir(dir)
		if err != nil {
			log.Error().Err(err).Msg("could not read challenge directory")
			continue
		}
		for _, challenge := range challenges {
			chalFolder := path.Join(dir, challenge.Name())
			log.Debug().Str("path", chalFolder).Msg("reading directory")
			err := syncChallenge(chalFolder)
			if err != nil {
				log.Warn().Err(err).Str("path", chalFolder).Msg("could not sync challenge")
			}
		}
	}
}

func Query(c *cobra.Command, args []string) {
	client := NewAPIClient("ctfd_b565c8f736af89490959917d5b932f85e379e09c5dd2b33fced654d774767f99", "https://ctf.maplebacon.org")
	data, err := client.getChallenges()
	if err != nil {
		log.Error().Err(err).Msg("borked")
	}
	log.Info().Any("challenges", data).Msg("Listing Challenges")
}
