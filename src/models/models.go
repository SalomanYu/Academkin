package models


type Vuz struct {
	VuzId		string	`bson:"vuz_id"`
	ShortName 	string	`bson:"short_name"`
	FullName	string	`bson:"full_name"`
	Logo		string	`bson:"logo"`
	City		string	`bson:"city"`
	Locality	string	`bson:"locality"`
}


type Specialization struct {
	Id					int 	`bson:"spec_id"`
	VuzId				string	`bson:"vuz_id"`
	VuzFullName 		string	`bson:"vuz_name"` 
	Name 				string	`bson:"name"`
	FormEducation 		string	`bson:"form_educations"`
	Duration 			string	`bson:"duration"`
	PreparationLevel 	string	`bson:"preparation_level"`
	Qualification		string	`bson:"qualification"`
}