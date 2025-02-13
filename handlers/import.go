package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/niels-hb/consilium-admin/models"
)

func startImport(uid string, source string, dryRun bool) {
	log.Printf("Importing transactions and schedules from file %v to user %v.\n", source, uid)

	println()

	log.Printf("Reading %v...\n", source)
	export := readFile(source)
	log.Printf("Read %v. Found %v transactions and %v schedules.\n", source, len(export.Transactions), len(export.Schedules))

	println()

	log.Printf("Importing %v transactions...\n", len(export.Transactions))
	writeTransactions(uid, export.Transactions, dryRun)
	log.Printf("Imported %v transactions.\n", len(export.Transactions))

	println()

	log.Printf("Importing %v schedules...\n", len(export.Schedules))
	writeSchedules(uid, export.Schedules, dryRun)
	log.Printf("Imported %v schedules.\n", len(export.Schedules))
}

func readFile(source string) models.FileExport {
	jsonFile, err := os.Open(source)
	if err != nil {
		log.Fatal(err.Error())
	}

	var export models.FileExport

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &export)

	return export
}

func writeTransactions(uid string, documents []models.TransactionExport, dryRun bool) {
	for _, document := range documents {
		document.UID = uid

		if !dryRun {
			ref, _, err := FirestoreClient.Collection("transactions").Add(Context, document.ToMap())

			if err != nil {
				log.Fatal(err.Error())
			} else {
				log.Printf("Created transaction with ID: %v", ref.ID)
			}
		}
	}
}

func writeSchedules(uid string, documents []models.ScheduleExport, dryRun bool) {
	for _, document := range documents {
		document.UID = uid

		if !dryRun {
			ref, _, err := FirestoreClient.Collection("schedules").Add(Context, document.ToMap())

			if err != nil {
				log.Fatal(err.Error())
			} else {
				log.Printf("Created schedule with ID: %v", ref.ID)
			}
		}
	}
}
