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

	log.Println("[SUCCESS] Env check is done!")
}

func main() {

	// We have to check in use env configs to prevent app crash while serving
	envConfigurationCheck()
	prepareDatabaseClient()

	httpRouter := http.NewServeMux()

	gh := &handler.GameHandler{}

	// net/http doesn't support url parameter without manual work, so, i decided to use query param to make it easier
	httpRouter.HandleFunc("/games/flush", gh.FlushGames)
	httpRouter.HandleFunc("/games/add", gh.AddGame)
	httpRouter.HandleFunc("/games/delete", gh.DeleteGame)
	httpRouter.HandleFunc("/games", gh.GetAllGames)

	log.Println("[STARTING] Server is starting to serve...")
	go saveDatabase()

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
			log.Fatalf("[ERROR] Unable to write file: %v", err2)
		}
		log.Println("[SUCCESS] New Database file was created!")
	} else {
		var gameDbHolder []domain.Game
		if err := json.Unmarshal(content, &gameDbHolder); err != nil {
			log.Fatalf("[ERROR] Data in Database file was broken! Have to be deleted!")
		} else {
			domain.GameDatabase = gameDbHolder
			log.Println("[SUCCESS] Database file was read!")
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
		var min, _ = strconv.ParseInt(os.Getenv(common.SaveInterval), 10, 32)

		time.Sleep(time.Duration(int(min)) * time.Second)
		//time.Sleep(1 * time.Minute)

		dbJsonText, _ := json.Marshal(domain.GameDatabase)

		err := ioutil.WriteFile("database.txt", []byte(dbJsonText), 0755)
		if err != nil {
			log.Fatalf("[ERROR] Unable to write file: %v", err)
		}
		log.Println("[SUCCESS] Database file was updated!")
	}
}
