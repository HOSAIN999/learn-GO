package logger

import "log"

func LisenForLog(ch chan string) {
	for {
		mgs := <-ch
		log.Print(mgs)
	}
}
