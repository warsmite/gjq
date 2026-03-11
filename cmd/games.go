package cmd

import (
	"fmt"
	"strings"

	"github.com/0xkowalskidev/gsq"
	"github.com/spf13/cobra"
)

func NewGamesCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "games",
		Short: "List supported games",
		RunE: func(cmd *cobra.Command, args []string) error {
			games := gsq.SupportedGames()

			// Find max widths for alignment
			nameWidth := len("NAME")
			slugWidth := len("SLUG")
			for _, g := range games {
				if len(g.Name) > nameWidth {
					nameWidth = len(g.Name)
				}
				if len(g.Slug) > slugWidth {
					slugWidth = len(g.Slug)
				}
			}
			nameWidth += 2
			slugWidth += 2

			fmtStr := fmt.Sprintf("%%-%ds %%-%ds %%-10s %%-12s %%-15s %%s\n", nameWidth, slugWidth)

			fmt.Printf(fmtStr, "NAME", "SLUG", "GAME PORT", "QUERY PORT", "ALIASES", "PROTOCOL")
			fmt.Printf(fmtStr, strings.Repeat("-", nameWidth-2), strings.Repeat("-", slugWidth-2), "---------", "----------", "-------", "--------")

			for _, g := range games {
				aliases := "-"
				if len(g.Aliases) > 0 {
					aliases = strings.Join(g.Aliases, ", ")
				}
				fmt.Printf(fmt.Sprintf("%%-%ds %%-%ds %%-10d %%-12d %%-15s %%s\n", nameWidth, slugWidth), g.Name, g.Slug, g.DefaultGamePort, g.DefaultQueryPort, aliases, g.Protocol)
			}

			return nil
		},
	}
}
