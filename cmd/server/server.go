package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"wtk/biz"
	"wtk/biz/credential"
	"wtk/ext/secure"
	"wtk/ext/store/ram"
	"wtk/ext/store/sqlite"
	"wtk/ui/api"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	host := flag.String("host", "0.0.0.0", "host name to listen on")
	port := flag.Int("port", 8000, "port to listen on")
	static := flag.String("static", "./ui/web/", "static files folder")
	help := flag.Bool("help", false, "show command usage")
	flag.Parse()

	if *help {
		flag.Usage()
		return nil
	}

	secure, err := secure.New()
	if err != nil {
		return fmt.Errorf("impossible d'initialiser le module de sécurité: %w", err)
	}

	initSchema := false

	// storage, initSchema, err := useRamDB()
	storage, initSchema, err := useSqliteDB()
	if err != nil {
		return fmt.Errorf("impossible d'obtenir une base: %w", err)
	}
	defer func() {
		err = storage.Close()
		if err != nil {
			log.Fatalf("closing database: %s", err)
		}
	}()

	// configure le métier avec le storage
	// mais on peut imaginer d'avoir un système de stockage différencié
	biz.Initialize(storage, secure)

	if initSchema {
		err = storage.InitSchema()
		if err != nil {
			return fmt.Errorf("impossible de créer le schéma de la base: %w", err)
		}
		// ajoute un user de test
		err = credential.Add("test", "test", "test", credential.Administrator)
		if err != nil {
			return fmt.Errorf("impossible de créer un utilisateur de test: %w", err)
		}

		// ajoute 5 fake user dans realm test
		for i := 1; i <= 5; i++ {
			err = credential.Add("test", "test"+strconv.Itoa(i), "test", credential.Customer)
		}
		// ajoute 5 fake user dans realm fakerealm
		for i := 1; i <= 5; i++ {
			err = credential.Add("fakerealm", "test"+strconv.Itoa(i), "test", credential.Customer)
		}
	}

	// initialise le serveur api
	server := api.New(*static, *host, *port)

	// lance le server
	return server.Run()
}

// composant de stockage en RAM
func useRamDB() (store biz.Storage, needSchema bool, err error) {
	st, err := ram.New()
	if err != nil {
		return store, false, err
	}
	return &st, true, nil
}

// composant de stockage sqlite
func useSqliteDB() (store biz.Storage, needSchema bool, err error) {
	const pragma = "_pragma=foreign_keys(1)&_pragma=journal_mode(WAL)&_pragma=synchronous(NORMAL)&_pragma=busy_timeout(8000)&_pragma=journal_size_limit(100000000)"
	const databasePath = "database.db"
	_, err = os.Stat(databasePath)
	if errors.Is(err, os.ErrNotExist) {
		_, err = os.Create(databasePath)
		if err != nil {
			return store, false, err
		}
		needSchema = true
	} else if err != nil {
		return store, false, err
	}

	store, err = sqlite.New(databasePath, pragma)
	if err != nil {
		return store, false, err
	}

	return store, needSchema, err
}
