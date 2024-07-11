package main

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"

	"github.com/niels-hb/consilium-admin/handlers"
	"github.com/thatisuday/commando"
	"google.golang.org/api/option"
)

func main() {
	initializeApp()
	initializeCommando()
}

func initializeApp() (*firebase.App, error) {
	ctx := context.Background()

	opt := option.WithCredentialsFile("service_account.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	firestoreClient, err := app.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("error initializing firestore: %v", err)
	}

	authClient, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("error initializing auth: %v", err)
	}

	handlers.Context = ctx
	handlers.FirestoreClient = firestoreClient
	handlers.AuthClient = authClient

	return app, err
}

func initializeCommando() {
	commando.
		SetExecutableName("consilium-admin").
		SetVersion("1.0.0").
		SetDescription("This tool contains administrative commands for the Consilium backend.")

	commando.
		Register("delete").
		SetShortDescription("delete a user and his documents").
		SetDescription("Deletes a user including his documents (transactions & schedules). The user will have the create a new account after this command has been executed.").
		AddFlag("uid,u", "UID of the user", commando.String, nil).
		AddFlag("dry-run", "don't send actual requests to the server", commando.Bool, nil).
		SetAction(handlers.Delete)

	commando.
		Register("export").
		SetShortDescription("export a users documents to a JSON file").
		SetDescription("Export all transactions and schedules owned by the user identified by <uid>. <target> specifies a JSON file in which the documents should be saved.").
		AddFlag("uid,u", "UID of the user", commando.String, nil).
		AddFlag("target,t", "target file", commando.String, nil).
		SetAction(handlers.Export)

	commando.
		Register("generate").
		SetShortDescription("create random users and data").
		SetDescription("Create <count> random users. Each user will have <transactions-min> <= x <= <transactions-max> transactions and <schedules-min> <= y <= <schedules-max> schedules generated as well.").
		AddFlag("count,c", "number of users to create", commando.Int, 1).
		AddFlag("transactions-min", "minimum number of transactions to generate per created user", commando.Int, 0).
		AddFlag("transactions-max", "maximum number of transactions to generate per created user", commando.Int, 0).
		AddFlag("schedules-min", "minimum number of schedules to generate per created user", commando.Int, 0).
		AddFlag("schedules-max", "maximum number of schedules to generate per created user", commando.Int, 0).
		AddFlag("dry-run", "don't send actual requests to the server", commando.Bool, nil).
		SetAction(handlers.Generate)

	commando.
		Register("import").
		SetShortDescription("import data from a JSON file").
		SetDescription("Import a previously created export JSON file. The owner of the created documents will be the user identified by <uid>. Combining the export and import commands essentially allows the duplication of documents to another user account without modifying the original user. <source> specifies the JSON file created using the export command.").
		AddFlag("uid,u", "UID of the user", commando.String, nil).
		AddFlag("source,s", "source file", commando.String, nil).
		AddFlag("dry-run", "don't send actual requests to the server", commando.Bool, nil).
		SetAction(handlers.Import)

	commando.
		Register("migrate").
		SetShortDescription("migrate documents to a different user").
		SetDescription("Set the UID for all transactions and schedules owned by the user identified by <from> to the UID <to>. This will effectively transfer ownership of all transactions and schedules the user <from> has created.").
		AddFlag("from,f", "UID of the current owner", commando.String, nil).
		AddFlag("to,t", "UID of the user to which ownership should be transferred", commando.String, nil).
		AddFlag("dry-run", "don't send actual requests to the server", commando.Bool, nil).
		SetAction(handlers.Migrate)

	commando.Parse(nil)
}
