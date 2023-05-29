package cli

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strings"
	"time"

	"github.com/hako/durafmt"
	"github.com/spf13/cobra"
	"github.com/uztadh/ledger"
)

// statsCmd represents the stats command
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "A small report of transaction stats",
	Run: func(cmd *cobra.Command, args []string) {
		transactions, terr := cliTransactions()
		if terr != nil {
			log.Fatalln(terr)
		}
		printStats(transactions)
	},
}

func printStats(generalLedger []*ledger.Transaction) {
	if len(generalLedger) < 1 {
		fmt.Println("Empty ledger.")
		return
	}
	sort.Slice(generalLedger, func(i, j int) bool {
		return generalLedger[i].Date.Before(generalLedger[j].Date)
	})

	startDate := generalLedger[0].Date
	endDate := generalLedger[len(generalLedger)-1].Date

	payees := make(map[string]struct{})
	accounts := make(map[string]struct{})

	var postings int64
	for _, trans := range generalLedger {
		payees[strings.ToLower(strings.TrimSpace(trans.Payee))] = struct{}{}
		for _, account := range trans.AccountChanges {
			postings++
			accounts[account.Name] = struct{}{}
		}
	}

	days := math.Floor(endDate.Sub(startDate).Hours() / 24)

	fmt.Printf("%-25s : %s to %s (%s)\n", "Time period", startDate.Format(time.DateOnly), endDate.Format(time.DateOnly), durafmt.Parse(endDate.Sub(startDate)).String())
	fmt.Printf("%-25s : %d\n", "Unique payees", len(payees))
	fmt.Printf("%-25s : %d\n", "Unique accounts", len(accounts))
	fmt.Printf("%-25s : %d (%.1f per day)\n", "Number of transactions", len(generalLedger), float64(len(generalLedger))/days)
	fmt.Printf("%-25s : %d (%.1f per day)\n", "Number of postings", postings, float64(postings)/days)
	fmt.Printf("%-25s : %s\n", "Time since last post", durafmt.ParseShort(time.Since(endDate)).String())
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
