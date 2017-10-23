package singletons

import (
	"io/ioutil"
	"sync"

	objects "../objects"
	"github.com/chrislusf/glow/flow"
)

type singleton struct {
	UserList []*objects.User // array of user objects
}

var instance *singleton // single singleton instance
var once sync.Once      // threadsafe runner

//GetInstance returns the singleton instance
func GetUsersInstance() *singleton {
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

//SendWhere sends a message to users matching the criteria function f
func (s *singleton) SendWhere(message string, f func(u *objects.User) bool) {
	flow.New().Slice(GetUsersInstance().UserList).Filter(f).Map(func(u *objects.User) {
		u.C <- message
	}).Run()

}

//SendWhereAndUpdate sends a message to users matching the criteria function f and running an after function up
func (s *singleton) SendWhereAndUpdate(message string, f func(u *objects.User) bool, up func(u *objects.User)) {
	flow.New().Slice(GetUsersInstance().UserList).Filter(f).Map(func(u *objects.User) {
		u.C <- message
		up(u)
	}).Run()

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
