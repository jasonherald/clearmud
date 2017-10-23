package ticker

import (
	"time"

	objects "../objects"
	sing "../singletons"
)

//UpdateTime sends time related messages at certain times
func UpdateTime() {
	h := time.Now().Hour()
	for _, m := range sing.GetMessagesInstance().MessageList {
		if m.When == h {
			sing.GetUsersInstance().SendWhereAndUpdate(m.Body, func(u *objects.User) bool {
				return u.SentTimeUpdate != m.When
			}, func(u *objects.User) {
				u.SentTimeUpdate = m.When
			})
		}
	}
}
