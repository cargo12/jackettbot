package main

import (
	"flag"
	"log"
	"os"
)

var (
	telegramTokenBot 1203236369:AAERcUpBp3SWqNEj90WJBbPGhPYInfCpkAU
	jackettAddress   http://cupid.whatbox.ca:11741
	jackettAPIKey    he1iek255fq1lw85a20vokrgfumsvt1l
	dev              bool
)

func readEnvVar() {
	var ok bool

	telegramTokenBot, ok = os.LookupEnv("TELEGRAM_TOKEN")
	if !ok {
		log.Fatal("TELEGRAM_TOKEN env not found.")
	}

	jackettAPIKey, ok = os.LookupEnv("JACKETT_API_KEY")
	if !ok {
		log.Fatal("JACKETT_API_KEY env not found.")
	}

	jackettAddress, ok = os.LookupEnv("JACKETT_ADDRESS")
	if !ok {
		log.Fatal("JACKETT_ADDRESS env not found.")
	}
}

func setCLIParams() {
	flag.BoolVar(&dev, "dev", false, "Dev options")
	flag.Parse()
}
