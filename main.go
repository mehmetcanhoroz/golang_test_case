package main

import (
	"YS_TestCase/common"
	"YS_TestCase/domain"
	"YS_TestCase/handler"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func envConfigurationCheck() {
	envConfigs := []string{
		common.SaveInterval,
	}

	for _, config := range envConfigs {
		// It returns the value, which will be empty if the variable is not present.
		// Go explanation, so no need to check null state
		if os.Getenv(config) == "" {
			log.Fatalf("[ERROR] Env config '%s' is missing. Need to set.", config)
		}
	}

	log.Println("[PRE_CHECK] Env check is done!")
}

func main() {

	// We have to check in use env configs to prevent app crash while serving
	envConfigurationCheck()
	prepareDatabaseClient()
	prepareInMemoryDatabaseClient()

	httpRouter := http.NewServeMux()

	gh := &handler.GameHandler{}
	ih := &handler.InMemoryHandler{}

	// net/http doesn't support url parameter without manual work, so, i decided to use query param to make it easier
	httpRouter.HandleFunc("/games/flush", gh.FlushGames)
	httpRouter.HandleFunc("/games/add", gh.AddGame)
	httpRouter.HandleFunc("/games/delete", gh.DeleteGame)
	httpRouter.HandleFunc("/games", gh.GetAllGames)


	httpRouter.HandleFunc("/fake-redis/add", ih.AddValue)
	httpRouter.HandleFunc("/fake-redis/flush", ih.FlushValues)
	httpRouter.HandleFunc("/fake-redis/delete", ih.DeleteValue)
	httpRouter.HandleFunc("/fake-redis", ih.GetAllValues)

	log.Println("[STARTING] Server is starting to serve...")

	go saveDatabase()
	go saveInMemoryDatabase()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, logIncomingRequests(httpRouter)))
}

func prepareDatabaseClient() {
	content, err := ioutil.ReadFile("database.txt")

	if err != nil {
		log.Println("[ERROR] Not able to connect database, trying to create new DB")

		game := domain.Game{
			Id:    1,
			Name:  "New World (Default Item)",
			Price: 109,
		}

		domain.GameDatabase = append(domain.GameDatabase, game)

		dbJsonText, _ := json.Marshal(domain.GameDatabase)

		err2 := ioutil.WriteFile("database.txt", []byte(dbJsonText), 0755)
		if err2 != nil {
			log.Fatalf("[ERROR] Unable to write db file: %v", err2)
		}
		log.Println("[PRE_CHECK] New Database file was created!")
	} else {
		var gameDbHolder []domain.Game
		if err := json.Unmarshal(content, &gameDbHolder); err != nil {
			log.Fatalf("[ERROR] Data in Database file was broken! Have to be deleted!")
		} else {
			domain.GameDatabase = gameDbHolder
			log.Println("[PRE_CHECK] Database file was read!")
		}
	}
}

func prepareInMemoryDatabaseClient() {
	domain.InMemoryDatabase = make(map[string]string)
	content, err := ioutil.ReadFile("fake_redis_database.txt")

	if err != nil {
		log.Println("[ERROR] Not able to connect fake redis database, trying to create new Fake Redis DB")

		dbJsonText, _ := json.Marshal(domain.InMemoryDatabase)

		err2 := ioutil.WriteFile("fake_redis_database.txt", []byte(dbJsonText), 0755)
		if err2 != nil {
			log.Fatalf("[ERROR] Unable to write fake redis db file: %v", err2)
		}
		log.Println("[PRE_CHECK] New Fake Redis Database file was created!")
	} else {
		var valueMemoryHolder map[string]string
		if err := json.Unmarshal(content, &valueMemoryHolder); err != nil {
			log.Fatalf("[ERROR] Fake Redis Data in Database file was broken! Have to be deleted!")
		} else {
			domain.InMemoryDatabase = valueMemoryHolder
			log.Println("[PRE_CHECK] Fake Redis Database file was read!")
		}
	}
}
func logIncomingRequests(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[INCOMING REQUEST] %s %s %s \n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func saveDatabase() {
	for {
		var sec, _ = strconv.ParseInt(os.Getenv(common.SaveInterval), 10, 32)

		time.Sleep(time.Duration(int(sec)) * time.Second)

		dbJsonText, _ := json.Marshal(domain.GameDatabase)

		err := ioutil.WriteFile("database.txt", []byte(dbJsonText), 0755)
		if err != nil {
			log.Fatalf("[ERROR] Unable to write db file: %v", err)
		}
		log.Println("[PRE_CHECK] Database file was updated!")
	}
}

func saveInMemoryDatabase() {
	for {
		var sec, _ = strconv.ParseInt(os.Getenv(common.SaveInterval), 10, 32)

		time.Sleep(time.Duration(int(sec)) * time.Second)

		dbJsonText, _ := json.Marshal(domain.InMemoryDatabase)

		err := ioutil.WriteFile("fake_redis_database.txt", []byte(dbJsonText), 0755)
		if err != nil {
			log.Fatalf("[ERROR] Unable to write fake redis db file: %v", err)
		}
		log.Println("[PRE_CHECK] Fake Redis Database file was updated!")
	}
}