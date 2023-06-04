package main

import "fmt"

// interface
type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

type Logger interface {
	Log(message string)
}

type Logic interface {
	SayHello(userID string) (string, bool)
}

type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

func LogOutput(message string) {
     fmt.Println(message)
}

type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore {
		userData: map[string]string{
			"1": "行秋",
			"2": "クレー",
			"3": "ベネット",
		},
	}
}

type SimpleLogic struct {
	l Logger
	ds DataStore
}

func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic {
		l: l,
		ds: ds,
	}
}

func (sl SimpleLogic) SayHello(userID string) (string, bool) {
	sl.l.Log("SayHello(" + userID + ")")
	name, ok := sl.ds.UserNameForID("1")
	if !ok {
		return "", errors.New("不明なユーザー")
	}
	return name + "さん、こんにちは！", nil
}

func (sl SimpleLogic) SayGoodbye(userID string) (string, bool) {
	sl.l.Log("SayGoodbye(" + userID + ")")
	name, ok := sl.ds.UserNameForID("2")
	if !ok {
		return "", errors.New("不明なユーザー")
	}
	return name + "さん、こんばんわ！！！". nil
}
