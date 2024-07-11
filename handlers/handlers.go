package handlers

import (
	"context"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
	"github.com/thatisuday/commando"
)

var Context context.Context
var FirestoreClient *firestore.Client
var AuthClient *auth.Client

func Delete(_ map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
	uid, _ := flags["uid"].GetString()
	dryRun, _ := flags["dry-run"].GetBool()

	startDelete(uid, dryRun)
}

func Export(_ map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
	uid, _ := flags["uid"].GetString()
	target, _ := flags["target"].GetString()

	startExport(uid, target)
}

func Generate(_ map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
	userCount, _ := flags["count"].GetInt()
	transactionsMin, _ := flags["transactions-min"].GetInt()
	transactionsMax, _ := flags["transactions-max"].GetInt()
	schedulesMin, _ := flags["schedules-min"].GetInt()
	schedulesMax, _ := flags["schedules-max"].GetInt()
	dryRun, _ := flags["dry-run"].GetBool()

	startGeneration(userCount, transactionsMin, transactionsMax, schedulesMin, schedulesMax, dryRun)
}

func Import(_ map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
	uid, _ := flags["uid"].GetString()
	source, _ := flags["source"].GetString()
	dryRun, _ := flags["dry-run"].GetBool()

	startImport(uid, source, dryRun)
}

func Migrate(_ map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
	from, _ := flags["from"].GetString()
	to, _ := flags["to"].GetString()
	dryRun, _ := flags["dry-run"].GetBool()

	startMigration(from, to, dryRun)
}
