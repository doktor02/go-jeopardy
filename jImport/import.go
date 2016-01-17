package jImport

import (
	"encoding/csv"
	"errors"
	"github.com/Sirupsen/logrus"
	"gopkg.in/olivere/elastic.v3"
	"os"
)

var log = logrus.New()

const jeopardyIndex = "jeopardy"

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

	_ = csv.NewReader(csvFile)

	client, err := elastic.NewClient(
		elastic.SetURL("http://elasticsearch:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		return err
	}

	exists, err := client.IndexExists(jeopardyIndex).Do()
	if err != nil {
		return err
	}
	if !exists {
		createIndex, err := client.CreateIndex(jeopardyIndex).Do()
		if err != nil {
			return err
		}
		if !createIndex.Acknowledged {
			return errors.New("Unacknowledged on index creation")
		}
	}

	log.Info("Csv and Elasticsearch ready")

	return nil
}
