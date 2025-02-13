package handlers

import (
	"fmt"
	"log"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/bxcodec/faker/v3"
	"github.com/niels-hb/consilium-admin/models"
	"github.com/niels-hb/consilium-admin/random"
)

func startGeneration(userCount int, transactionsMin int, transactionsMax int, schedulesMin int, schedulesMax int, dryRun bool) {
	log.Printf("Generating %v users with %v <= transactions <= %v and %v <= schedules <= %v.\n", userCount, transactionsMin, transactionsMax, schedulesMin, schedulesMax)

	if dryRun {
		log.Println("[!] Dry run is active. Won't send any writing requests.")
	}

	transactionCount, err := random.GetRandomIntInRange(transactionsMin, transactionsMax)
	if err != nil {
		log.Fatalf("transactions: %v", err.Error())
	}

	scheduleCount, err := random.GetRandomIntInRange(schedulesMin, schedulesMax)
	if err != nil {
		log.Fatalf("schedules: %v", err.Error())
	}

	log.Printf("Creating %v users...\n", userCount)
	for i := 0; i < userCount; i++ {
		uid, email, password := createUser(dryRun)

		log.Printf("Generated user with UID %v (%v / %v).\n", uid, email, password)

		println()

		log.Printf("Generating %v transactions...\n", transactionCount)
		createTransactions(uid, transactionCount, dryRun)
		log.Printf("Generated %v transactions.\n", transactionCount)

		println()

		log.Printf("Generating %v schedules...\n", scheduleCount)
		createSchedules(uid, scheduleCount, dryRun)
		log.Printf("Generated %v schedules.\n", scheduleCount)
	}
}

func createUser(dryRun bool) (string, string, string) {
	firstName := faker.FirstName()
	lastName := faker.LastName()
	email := strings.ToLower(fmt.Sprintf("%v.%v@generated.com", firstName, lastName))
	password := faker.Password()

	params := (&auth.UserToCreate{}).
		Email(email).
		EmailVerified(true).
		Password(password).
		DisplayName(fmt.Sprintf("%v %v", firstName, lastName))

	if dryRun {
		return "nil", email, password
	}

	userRecord, err := AuthClient.CreateUser(Context, params)
	if err != nil {
		log.Fatalf("User couldn't be created. %v", err.Error())
	}

	return userRecord.UID, email, password
}

func createTransactions(uid string, count int, dryRun bool) {
	for i := 0; i < count; i++ {
		amountCents, _ := random.GetRandomIntInRange(0.01*100, 1000*100)

		document := models.TransactionExport{
			UID:         uid,
			AmountCents: amountCents,
			Category:    random.GetRandomCategory(),
			CreatedOn:   random.GetRandomTime(),
			Name:        faker.Word(),
			Note:        faker.Sentence(),
		}

		if !dryRun {
			FirestoreClient.Collection("transactions").Add(Context, document.ToMap())
		}
	}
}

func createSchedules(uid string, count int, dryRun bool) {
	for i := 0; i < count; i++ {
		amountCents, _ := random.GetRandomIntInRange(0.01*100, 1000*100)
		frequencyMonths, _ := random.GetRandomIntInRange(1, 12)

		document := models.ScheduleExport{
			UID:             uid,
			AmountCents:     amountCents,
			Category:        random.GetRandomCategory(),
			CreatedOn:       random.GetRandomTime(),
			Name:            faker.Word(),
			Note:            faker.Sentence(),
			CanceledOn:      nil,
			FrequencyMonths: frequencyMonths,
			StartedOn:       random.GetRandomTime(),
			ScheduleType:    random.GetRandomScheduleType(),
		}

		if !dryRun {
			FirestoreClient.Collection("schedules").Add(Context, document.ToMap())
		}
	}
}
