package main

import (
	"errors"
	"fmt"
	"net/http"
)

func listeningController() {
	l := LoggerAdapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	http.HandleFunc("/hello", c.HandleGreeting)
	http.ListenAndServe(":8080", nil)
}

func LogOutput(message string) {
	fmt.Println(message)
}

type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForId(userId string) (string, bool) {
	name, ok := sds.userData[userId]
	return name, ok
}

func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "joy",
			"2": "world",
			"3": "warm",
			"4": "oil",
		},
	}
}

type DataStore interface {
	UserNameForId(userId string) (string, bool)
}

type Logger interface {
	Log(message string)
}

type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

type SimpleLogic struct {
	l  Logger
	ds DataStore
}

func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

func (sl SimpleLogic) SayHello(userId string) (string, error) {
	sl.l.Log("in sayHello for " + userId)
	name, ok := sl.ds.UserNameForId(userId)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "hello " + name, nil
}

func (sl SimpleLogic) SayGoodBye(userId string) (string, error) {
	sl.l.Log("in say goodBye for " + userId)
	name, ok := sl.ds.UserNameForId(userId)
	if !ok {
		return "", errors.New("unknown user ")
	}
	return "Goodbye ," + name, nil
}

type Logic interface {
	SayHello(userId string) (string, error)
}

type Controller struct {
	l     Logger
	logic Logic
}

func (c Controller) HandleGreeting(w http.ResponseWriter, r *http.Request) {
	c.l.Log("In say hello")
	userId := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

func NewController(l Logger, logic Logic) Controller {
	return Controller{
		l:     l,
		logic: logic,
	}
}
