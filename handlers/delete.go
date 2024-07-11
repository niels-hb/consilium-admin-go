package handlers

import "log"

func startDelete(uid string, dryRun bool) {
	log.Printf("Deleting user with UID %v...\n", uid)

	if dryRun {
		log.Println("[!] Dry run is active. Won't send any writing requests.")
	}

	deleteUser(uid, dryRun)
	println()
	deleteCollection("transactions", uid, dryRun)
	println()
	deleteCollection("schedules", uid, dryRun)
}

func deleteUser(uid string, dryRun bool) {
	if !dryRun {
		AuthClient.DeleteUser(Context, uid)
	}
}

func deleteCollection(collection string, uid string, dryRun bool) {
	documents, _ := FirestoreClient.Collection(collection).Where("uid", "==", uid).Documents(Context).GetAll()
	documentsLength := len(documents)

	log.Printf("Deleting %v %v...\n", documentsLength, collection)

	for _, document := range documents {
		if !dryRun {
			document.Ref.Delete(Context)
		}
	}

	log.Printf("Deleted %v %v.\n", documentsLength, collection)
}
