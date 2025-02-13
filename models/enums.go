package models

func GetCategories() []string {
	return []string{
		"fitness",
		"food",
		"housing",
		"insurance",
		"medical",
		"miscellaneous",
		"hygiene",
		"recreation",
		"savings",
		"subscriptions",
		"transportation",
		"utilities",
	}
}

func GetScheduleTypes() []string {
	return []string{
		"incoming",
		"outgoing",
	}
}
