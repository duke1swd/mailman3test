package main

import (
	"log"
)

type domain struct {
	mailHost    string
	description string
	url         string
}

var (
	domainId int
	domains  []domain
)

func setDomain() {
	loadDomains()
	domainId = 0
}

func loadDomains() {
	res, ok := get(configuration.Url + "/domains")
	if !ok {
		log.Fatal("Cannot load domains collection")
	}

	resMap, ok := res.(map[string]interface{})
	if !ok {
		log.Fatal("badly formed json in response to get domains")
	}

	entriesRaw, ok := resMap["entries"]
	if !ok {
		log.Fatal("missing entries in domain response")
	}

	entries, ok := entriesRaw.([]interface{})
	if !ok {
		log.Fatal("badly formed json in response to get domains/entries")
	}

	domains = make([]domain, len(entries))

	for i, e := range entries {
		//fmt.Printf("Entry %d is %v\n\n", k, e)
		entry, ok := e.(map[string]interface{})
		if !ok {
			log.Fatal("badly formed json in response to get domains/entries/entry")
		}
		mhRaw, ok := entry["mail_host"]
		if !ok {
			log.Fatal("missing mail_host")
		}
		domains[i].mailHost, ok = mhRaw.(string)
		if !ok {
			log.Fatal("mail_host not a string")
		}

		desRaw, ok := entry["description"]
		if !ok {
			log.Fatal("missing description")
		}
		domains[i].description, ok = desRaw.(string)
		if !ok {
			log.Fatal("description not a string")
		}

		urlRaw, ok := entry["self_link"]
		if !ok {
			log.Fatal("missing self_link")
		}
		domains[i].url, ok = urlRaw.(string)
		if !ok {
			log.Fatal("self_link not a string")
		}
	}
}
