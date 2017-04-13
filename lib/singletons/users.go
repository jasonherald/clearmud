package singletons

import (
	"io/ioutil"
	"sync"

	objects "../objects"
)

type singleton struct {
	UserList []*objects.User // array of user objects
}

var instance *singleton // single singleton instance
var once sync.Once      // threadsafe runner

//GetInstance returns the singleton instance
func GetInstance() *singleton {
	once.Do(func() { // run once
		instance = &singleton{} // instantiate the singleton
	})
	return instance // return the instance we created (if we created one)
}

//Add add a user to the list
func (s *singleton) Add(u *objects.User) {
	s.UserList = append(s.UserList, u) // append the user
	u.C <- loadMotd()                  // send the MOTD
}

//SendAll send a message to all users
func (s *singleton) SendAll(message string) {
	for _, element := range s.UserList { // loop over all the users
		element.C <- message // send the message through the channel
	}
}

//SendOne send a message to one user
func (s *singleton) SendOne(message string, index int) {
	s.UserList[index].C <- message // send the message
}

//Count get a count of logged in users
func (s *singleton) Count() int {
	return len(s.UserList)
}

//Remove drop a user who either disconnects or logs out
func (s *singleton) Remove(index int) {
	s.UserList[index] = s.UserList[len(s.UserList)-1] // all of this resizes the array
	s.UserList = s.UserList[:len(s.UserList)-1]       // drop the element
}

func loadMotd() string {
	content, err := ioutil.ReadFile("MOTD") // load the file contents for the MOTD
	if err != nil {                         // if reading the file wasn't successful
		panic(1) // oh dear god!!!!
	}
	return string(content) // return the string representation of the file's contents
}
