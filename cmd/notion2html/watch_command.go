package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/jomei/notionapi"
	"github.com/kapitanov/notion2html/internal/convert"
	"github.com/kapitanov/notion2html/internal/emit"
	"github.com/kapitanov/notion2html/internal/tree"
	"github.com/spf13/cobra"
)

func init() {
	command := &cobra.Command{
		Use:   "watch",
		Short: "generate a static website from Notion and watch for changes",
	}
	rootCommand.AddCommand(command)
	flags := command.Flags()

	flagOutputDirectory := flags.StringP("output", "o", "", "path to output directory")
	flagAccessToken := flags.StringP("token", "t", "", "access token for Notion")
	flagForce := flags.BoolP("force", "f", false, "rebuild all pages")
	flagPeriod := flags.DurationP("period", "p", 10*time.Minute, "refresh timer period")

	command.RunE = func(cmd *cobra.Command, args []string) error {
		outputDirectory := *flagOutputDirectory
		log.Printf("output directory %s", outputDirectory)
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		defer func() {
			signal.Stop(c)
			cancel()
		}()
		go func() {
			select {
			case <-c:
				cancel()
			case <-ctx.Done():
			}
		}()

		notion := notionapi.NewClient(notionapi.Token(*flagAccessToken))

		isFirstIteration := true
		for {
			pageSet, err := tree.Load(ctx, notion)
			if err != nil {
				return err
			}

			builder := convert.NewASTBuilder(notion)
			emitter, err := emit.NewEmitter(outputDirectory, builder, *flagForce)
			if err != nil {
				return err
			}

			n, err := emitter.Generate(ctx, pageSet)
			if err != nil {
				return err
			}

			if n > 0 {
				log.Printf("%d page(s) generated", n)
			}

			if isFirstIteration {
				log.Printf("%d page(s) generated", n)
				log.Printf("will check for updates every %s", *flagPeriod)
				isFirstIteration = false
			} else {
				if n > 0 {
					log.Printf("%d page(s) generated", n)
				}
			}

			timer := time.NewTimer(*flagPeriod)
			select {
			case <-ctx.Done():
				if !timer.Stop() {
					<-timer.C
				}
			case <-timer.C:
			}
		}
	}
}
