package main

import (
	"flag"
	"fmt"
	"log"
)

type subscribeT struct {
	ListID      string `json:"list_id"`
	Subscriber  string `json:"subscriber"`
	PreAproved  string `json:"pre_approved"`
	PreVerified string `json:"pre_verified"`
}

type subscribeTbool struct {
	ListID      string `json:"list_id"`
	Subscriber  string `json:"subscriber"`
	PreAproved  bool   `json:"pre_approved"`
	PreVerified bool   `json:"pre_verified"`
}

var (
	address string
	list    string
)

func sSetup() (listId, userId string) {
	setDomain()

	if flag.NArg() < 2 {
		fmt.Println("subscribe/unsubscribe requires a list and an email address")
		usage()
	}

	list = flag.Arg(0)
	address = flag.Arg(1)

	if flag.NArg() > 2 {
		fmt.Println("too many arguments to subscribe")
		usage()
	}

	listEntry := findList(list)
	listId = listEntry.listId

	resRaw, ok := get(configuration.Url + "/users/" + address)
	if !ok {
		log.Fatalf("User %s not found, not subscribed", address)
	}

	res, ok := resRaw.(map[string]interface{})
	if !ok {
		log.Fatal("badly formed json in response to get users/address (subscribe)")
	}

	_, ok = res["self_link"]
	if !ok {
		log.Fatalf("User %s not found (subscribe)", address)
	}

	// Get the user ID
	userIdraw, ok := res["user_id"]
	if !ok {
		log.Fatalf("User %s does not have a user_id (subscribe)", address)
	}
	userId, ok = userIdraw.(string)
	if !ok {
		log.Fatalf("User %s has bad json (subscribe)", address)
	}

	return
}

func subscribeCmd() {
	listId, userId := sSetup()

	if flagb {
		var subscribeStruct subscribeTbool
		subscribeStruct.ListID = listId
		subscribeStruct.Subscriber = userId
		subscribeStruct.PreAproved = true
		subscribeStruct.PreVerified = true

		post(configuration.Url+"/members", subscribeStruct)
	} else {
		var subscribeStruct subscribeT
		subscribeStruct.ListID = listId
		subscribeStruct.Subscriber = userId
		subscribeStruct.PreAproved = "true"
		subscribeStruct.PreVerified = "true"

		post(configuration.Url+"/members", subscribeStruct)
	}
	log.Printf("user %s subscribed to list %s", address, list)
}

func jsonDecode(entry map[string]interface{}, collection string, field string) string {
	raw, ok := entry[field]
	if !ok {
		log.Fatalf("no field \"%s\" in in collection %s", field, collection)
	}
	r, ok := raw.(string)
	if !ok {
		log.Fatalf("badly formed json in response to get %s / %s", collection, field)
	}
	return r
}
