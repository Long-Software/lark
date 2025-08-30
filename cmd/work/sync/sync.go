package sync

import (
	"flag"
	"fmt"
	"log"
	"os"
	"github.com/Long-Software/Sonality/internal/consts"
	"github.com/Long-Software/Sonality/internal/utils"
	"strings"
)

var (
	file = flag.String("file", "work.json", "Config file for syncing with the workspace")
)

type SyncCommand struct{}

func (s *SyncCommand) Run() {
	flag.Usage = s.Help
	flag.CommandLine.Parse(os.Args[2:])
	if *file == "" {
		s.Help()
	}
	var config consts.WorkConfig
	err := utils.LoadJsonConfig(&config, *file)
	if err != nil {
		log.Fatalf("Failed to load config file: %v", err)
		return
	}
	work_file, err := os.Create(consts.WORK_FILE)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer work_file.Close()

	goVersion := utils.GetGOVersion()
	actives := make([]string, len(config.Active))
	inactives := make([]string, len(config.InActive))

	for i, a := range config.Active {
		actives[i] = fmt.Sprintf("\t%s", a)
	}
	for i, a := range config.InActive {
		inactives[i] = fmt.Sprintf("\t//\t%s", a)
	}

	output := fmt.Sprintf("go %s\n\nuse (\n%s\n%s\n)",
		goVersion,
		strings.Join(actives, "\n"),
		strings.Join(inactives, "\n"),
	)
	_, err = work_file.WriteString(output)
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}
	fmt.Println("Data successfully written to ", consts.WORK_FILE)
}

func (s *SyncCommand) Help() {
	fmt.Println("usage: work sync [options]")
	flag.PrintDefaults()
}
