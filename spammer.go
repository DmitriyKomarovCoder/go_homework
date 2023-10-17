package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
)

func RunPipeline(cmds ...cmd) {
	var wg sync.WaitGroup
	in := make(chan interface{})
	out := in

	for _, c := range cmds {
		in = out
		out = make(chan interface{})

		wg.Add(1)
		go func(c cmd, in, out chan interface{}) {
			defer wg.Done()
			defer close(out)

			c(in, out)
		}(c, in, out)
	}

	wg.Wait()
}

func SelectUsers(in, out chan interface{}) {
	var uniqUser = &sync.Map{}
	var wg sync.WaitGroup

	for emailInterface := range in {
		email := emailInterface.(string)
		wg.Add(1)
		go func(email string) {
			defer wg.Done()

			user := GetUser(email)
			_, ucheck := uniqUser.LoadOrStore(user.ID, user)
			if !ucheck {
				out <- user
			}
		}(email)
	}

	wg.Wait()
}

func SelectMessages(in, out chan interface{}) {
	var wg sync.WaitGroup
	usrSlice := []User{}
	for userInterface := range in {
		usrSlice = append(usrSlice, userInterface.(User))
	}
	for i := 0; i < len(usrSlice); i += GetMessagesMaxUsersBatch {
		wg.Add(1)
		go func(curI int) {
			defer wg.Done()

			end := curI + GetMessagesMaxUsersBatch

			if end > len(usrSlice) {
				end = len(usrSlice)
			}

			msg, err := GetMessages(usrSlice[curI:end]...)

			if err != nil {
				fmt.Println("Error")
				return
			}

			for _, msgID := range msg {
				out <- msgID
			}
		}(i)
	}

	wg.Wait()
}

func worker(in <-chan MsgData, out chan<- interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	var err error
	for msg := range in {
		msg.HasSpam, err = HasSpam(msg.ID)

		if err != nil {
			fmt.Print("some err")
			return
		}

		out <- msg
	}
}

func CheckSpam(in, out chan interface{}) {
	var wg sync.WaitGroup

	numW := 5
	wChannels := make(chan MsgData, numW)

	for i := 1; i <= numW; i++ {
		wg.Add(1)
		go worker(wChannels, out, &wg)
	}

	for msgID := range in {
		msgData := MsgData{msgID.(MsgID), true}
		wChannels <- msgData
	}
	close(wChannels)
	wg.Wait()
}

func CombineResults(in, out chan interface{}) {
	var msgData []MsgData
	for res := range in {
		msgData = append(msgData, res.(MsgData))
	}

	sort.Slice(msgData, func(i, j int) bool {
		if msgData[i].HasSpam && !msgData[j].HasSpam {
			return true

		} else if !msgData[i].HasSpam && msgData[j].HasSpam {
			return false
		}

		return msgData[i].ID < msgData[j].ID
	})

	for _, m := range msgData {
		out <- strconv.FormatBool(m.HasSpam) + " " + strconv.FormatUint(uint64(m.ID), 10)
	}
}
