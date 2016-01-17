package jImport

import (
	"encoding/csv"
	"fmt"
	"github.com/Sirupsen/logrus"
	"os"
)

var log = logrus.New()

// RunImport executes the jeopardy question import
//   - downloads from such and such
//   - sets up es settings and mappings
//   - uploads to es
func RunImport(args []string) error {
	log.Out = os.Stderr
	log.WithFields(logrus.Fields{
		"args": args,
	}).Info("Running import")

	var filepath = args[0]

	csvFile, err := os.Open(filepath)
	if err != nil {
		return err
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	record, err := reader.Read()
	if err != nil {
		return err
	}
	fmt.Print(record)

	anotherRecord, err := reader.Read()
	if err != nil {
		return err
	}
	fmt.Print(anotherRecord)

	return nil
}
