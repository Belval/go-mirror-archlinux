package main

import (
	"fmt"
	"strconv"
	"net/http"
	"os"
	"os/exec"

	"github.com/jasonlvhit/gocron"
)

func rsyncArgsBuilder(useBackupServer bool) []string {
	args := make([]string, 0)
	args = append(args, "-rtlH")
	args = append(args, "--safe-links")
	args = append(args, "--delete-after")
	args = append(args, "--timeout=600")
	args = append(args, "--quiet")
	if config.BandwidthLimit > 0 {
		args = append(args, "--bwlimit="+strconv.Itoa(config.BandwidthLimit))
	}
	args = append(args, "--exclude='*.links.tar.gz*'")
	if !config.SyncISO {
		args = append(args, "--exclude='/iso'")
	}
	if !config.SyncOther {
		args = append(args, "--exclude='/other'")
	}
	if !config.SyncSources {
		args = append(args, "--exclude='/sources'")
	}
	if useBackupServer {
		args = append(args, config.BackupServer)
		fmt.Println(config.BackupServer)
	} else {
		args = append(args, config.PrimaryServer)
		fmt.Println(config.PrimaryServer)
	}
	args = append(args, config.RepoDirectory)
	return args
}

func synchronize() {
	fmt.Println("Syncing...")
	if err := exec.Command("rsync", rsyncArgsBuilder(false)...).Run(); err != nil {
		fmt.Println(err)
		if err := exec.Command("rsync", rsyncArgsBuilder(true)...).Run(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	fmt.Println("Done syncing!")
}

func serve() {
	fmt.Println("Initial sync")
	synchronize()
	fmt.Println("Mirror should be up and running")
	gocron.Every(uint64(config.SyncInterval)).Hours().Do(synchronize)
	http.Handle("/", http.FileServer(http.Dir(config.RepoDirectory)))
	fmt.Println(http.ListenAndServe(":"+strconv.Itoa(config.Port), nil))
}
