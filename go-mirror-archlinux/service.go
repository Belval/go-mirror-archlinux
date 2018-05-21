package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/jasonlvhit/gocron"
)

func rsyncArgsBuilder(useBackupServer bool) []string {
	args := make([]string, 1)
	args = append(args, "-rtlH")
	args = append(args, "--safe-links")
	args = append(args, "--delete-after")
	args = append(args, "--timeout=600")
	if config.BandwidthLimit > 0 {
		args = append(args, "--bwlimit="+string(config.BandwidthLimit))
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
	fmt.Println(config.RepoDirectory)
	args = append(args, config.RepoDirectory)
	args = append(args, "--quiet")
	return args
}

func synchronize() {
	fmt.Println("Syncing...")
	fmt.Println(rsyncArgsBuilder(false))
	if err := exec.Command("rsync", rsyncArgsBuilder(false)...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		if err := exec.Command("rsync", rsyncArgsBuilder(true)...).Run(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	fmt.Println("Done syncing!")
}

func serve() {
	gocron.Every(uint64(config.SyncInterval)).Hours().Do(synchronize)
	http.Handle("/", http.FileServer(http.Dir(config.RepoDirectory)))
	http.ListenAndServe(":"+string(config.Port), nil)
}
