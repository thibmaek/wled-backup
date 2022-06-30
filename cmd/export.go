package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	util "thibmaek/wled-export/util"
)

const (
	ExportCmdName = "export"
)

type ExportCmd struct {
	name      string
	hosts     string
	outputDir string
}

func NewExportCmd(flags []string) *ExportCmd {
	cmd := &ExportCmd{name: ExportCmdName}

	dir := util.GetRootDir()

	fs := flag.NewFlagSet(ExportCmdName, flag.ExitOnError)
	fs.StringVar(&cmd.hosts, "hosts", "", "A host or list of hosts to export configs from")
	fs.StringVar(&cmd.outputDir, "outputDir", dir, "Directory to save the exported JSON files to.")

	err := fs.Parse(flags)
	if err != nil {
		fs.PrintDefaults()
		os.Exit(1)
	}

	return cmd
}

func (cmd *ExportCmd) Name() string {
	return cmd.name
}

func (cmd *ExportCmd) Run() {
	if len(cmd.hosts) < 1 {
		fmt.Println("Please supply at least 1 host:")
		fmt.Println("  $ ./wled-backup --hosts=wled-tv.local,192.168.1.66")
		os.Exit(1)
	}

	hosts := []string{}
	if strings.Contains(cmd.hosts, ",") {
		cmd.hosts = strings.TrimSuffix(cmd.hosts, ",")
		hosts = strings.Split(cmd.hosts, ",")
	} else {
		hosts = append(hosts, cmd.hosts)
	}

	var wg sync.WaitGroup

	for _, h := range hosts {
		hostname, err := util.GetHostname(h)
		if err == nil {
			wg.Add(1)

			go func(hn string) {
				defer wg.Done()
				downloadConfig(h, filepath.Join(cmd.outputDir, hn))
				downloadPresets(h, filepath.Join(cmd.outputDir, hn))
			}(hostname)
		} else {
			fmt.Printf("Failed to get the hostname for host '%s': %v", h, err)
		}
	}

	wg.Wait()
}

func downloadConfig(host string, destDir string) {
	fmt.Printf("Downloading config for host '%s'\n", host)
	url := fmt.Sprintf("http://%s/edit?download=/cfg.json", host)
	err := util.DownloadFile(fmt.Sprintf("%s.config.json", destDir), url)
	if err != nil {
		log.Fatal(err)
	}
}

func downloadPresets(host string, destDir string) {
	fmt.Printf("Downloading presets for host '%s'\n", host)
	url := fmt.Sprintf("http://%s/edit?download=/presets.json", host)
	err := util.DownloadFile(fmt.Sprintf("%s.presets.json", destDir), url)
	if err != nil {
		log.Fatal(err)
	}
}
