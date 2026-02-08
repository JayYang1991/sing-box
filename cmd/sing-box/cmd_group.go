package main

import (
	"fmt"

	"github.com/sagernet/sing-box/adapter"
	"github.com/sagernet/sing-box/experimental/cachefile"
	"github.com/sagernet/sing-box/log"
	"github.com/sagernet/sing-box/option"

	"github.com/spf13/cobra"
)

var commandGroup = &cobra.Command{
	Use:   "group",
	Short: "Manage proxy groups",
}

var commandGroupList = &cobra.Command{
	Use:   "list",
	Short: "List proxy groups and their selected outbounds",
	Run: func(cmd *cobra.Command, args []string) {
		err := runGroupList()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	mainCommand.AddCommand(commandGroup)
	commandGroup.AddCommand(commandGroupList)
}

func runGroupList() error {
	options, err := readConfigAndMerge()
	if err != nil {
		return err
	}

	var cachePath string
	if options.Experimental != nil && options.Experimental.CacheFile != nil {
		cachePath = options.Experimental.CacheFile.Path
	}
	if cachePath == "" {
		return fmt.Errorf("cache file not configured")
	}

	cf := cachefile.New(globalCtx, option.CacheFileOptions{
		Path: cachePath,
	})
	err = cf.Start(adapter.StartStateInitialize)
	if err != nil {
		return fmt.Errorf("open cache file: %w", err)
	}
	defer cf.Close()

	fmt.Printf("%-20s %-20s\n", "GROUP", "SELECTED")
	fmt.Printf("%-20s %-20s\n", "-----", "--------")

	for _, outbound := range options.Outbounds {
		if outbound.Type == "selector" || outbound.Type == "urltest" {
			selected := cf.LoadSelected(outbound.Tag)
			if selected == "" {
				selected = "(none)"
			}
			fmt.Printf("%-20s %-20s\n", outbound.Tag, selected)
		}
	}

	return nil
}
