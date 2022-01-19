package main

import (
	"log"
)

/*
	The list object has these
*/

type listT struct {
	displayName  string
	listName     string
	mailHost     string
	listId       string
	description  string
	selfLink     string
	fqdnListname string
}

var lists []listT

func findList(list string) listT {
	loadLists()
	for _, l := range lists {
		if l.listName == list {
			return l
		}
	}
	log.Fatalf("list %s not found", list)
	panic("not reached")
}

func loadLists() {
	res, ok := get(domains[domainId].url + "/lists")
	if !ok {
		log.Fatal("Cannot load lists collection")
	}

	resMap, ok := res.(map[string]interface{})
	if !ok {
		log.Fatal("badly formed json in response to get lists")
	}

	entriesRaw, ok := resMap["entries"]
	if !ok {
		log.Fatal("missing entries in list response")
	}

	entries, ok := entriesRaw.([]interface{})
	if !ok {
		log.Fatal("badly formed json in response to get lists/entries")
	}

	lists = make([]listT, len(entries))

	for i, e := range entries {
		entry, ok := e.(map[string]interface{})
		if !ok {
			log.Fatal("badly formed json in response to get lists/entries/entry")
		}

		lists[i].displayName = jsonDecode(entry, "lists", "display_name")
		lists[i].listName = jsonDecode(entry, "lists", "list_name")
		lists[i].mailHost = jsonDecode(entry, "lists", "mail_host")
		lists[i].description = jsonDecode(entry, "lists", "description")
		lists[i].listId = jsonDecode(entry, "lists", "list_id")
		lists[i].selfLink = jsonDecode(entry, "lists", "self_link")
		lists[i].fqdnListname = jsonDecode(entry, "lists", "fqdn_listname")
	}
}
