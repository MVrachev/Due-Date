package data

import "github.com/end-date/user"

// This object will be used to handle and encrypt all the data
type DataHandler struct{}

// Add a user as a record
func (d *DataHandler) addIntoDatabase(user *user.User) {}

// Checks if a user is registered
func (d *DataHandler) doesRecordExists() bool {
	return false
}

func (d *DataHandler) deleteFromDatabase() {}

// Not sure what arugments will have
func (d *DataHandler) editRecord() {

}
