package singletons

import (
	"io/ioutil"
	"sync"

	objects "../objects"
)

type singleton struct {
	UserList []*objects.User
}

var instance *singleton
var once sync.Once

//GetInstance returns the singleton instance
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func (s *singleton) Add(u *objects.User) {
	s.UserList = append(s.UserList, u)
	u.C <- loadMotd()
}

func (s *singleton) SendAll(message string) {
	for _, element := range s.UserList {
		element.C <- message
	}
}

func (s *singleton) SendOne(message string, index int) {
	s.UserList[index].C <- message
}

func (s *singleton) Count() int {
	return len(s.UserList)
}

func (s *singleton) Remove(index int) {
	s.UserList[index] = s.UserList[len(s.UserList)-1]
	s.UserList = s.UserList[:len(s.UserList)-1]
}

func loadMotd() string {
	content, err := ioutil.ReadFile("MOTD") // load the file contents for the MOTD
	if err != nil {                         // if reading the file wasn't successulf
		panic(1) // oh dear god!!!!
	}
	return string(content) // return the string representation of the file's contents
}
