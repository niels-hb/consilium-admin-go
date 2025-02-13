package handlers

import (
	"encoding/json"
	"log"
	"os"

	"github.com/niels-hb/consilium-admin/models"
)

func startExport(uid string, target string) {
	log.Printf("Writing documents for user %v into %v.\n", uid, target)
	println()

	var data models.FileExport

	log.Println("Exporting transactions...")
	data.Transactions = exportTransactions(uid)
	log.Printf("Exported %v transactions.\n", len(data.Transactions))

	println()

	log.Println("Exporting schedules...")
	data.Schedules = exportSchedules(uid)
	log.Printf("Exported %v schedules.\n", len(data.Schedules))

	println()

	log.Println("Writing export to file...")
	writeFile(target, data)
	log.Println("Wrote export to file.")
}

func exportTransactions(uid string) []models.TransactionExport {
	documents, _ := FirestoreClient.Collection("transactions").Where("uid", "==", uid).Documents(Context).GetAll()
	var results []models.TransactionExport

	for _, document := range documents {
		var documentExport models.TransactionExport
		documentExport.FromJSON(document.Data())

		results = append(results, documentExport)
	}

	return results
}

func exportSchedules(uid string) []models.ScheduleExport {
	documents, _ := FirestoreClient.Collection("schedules").Where("uid", "==", uid).Documents(Context).GetAll()
	var results []models.ScheduleExport

	for _, document := range documents {
		var documentExport models.ScheduleExport
		documentExport.FromJSON(document.Data())

		results = append(results, documentExport)
	}

	return results
}

func writeFile(filename string, data models.FileExport) {
	content, _ := json.MarshalIndent(data, "", " ")
	err := os.WriteFile(filename, content, 0644)

	if err != nil {
		log.Fatal(err.Error())
	}
}
