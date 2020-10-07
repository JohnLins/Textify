package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os/exec"
	"runtime"
	"strings"

	"github.com/gorilla/mux"
	"github.com/nexmo-community/nexmo-go"
)

var acronyms = make(map[string]string)

func main() {
	acronyms["idk"] = "I don't know"
	acronyms["ikr"] = "I know right"
	acronyms["omg"] = "oh my gosh"
	acronyms["wtf"] = "what the f***"
	acronyms["wdym"] = "what do you mean"
	acronyms["btw"] = "by the way"

	fmt.Println("Server started...")
	fmt.Println(" * Running on http://127.0.0.1:8080/")
	fmt.Println(" * IP: localhost")
	fmt.Println(" * Port: 8080")

	r := mux.NewRouter()

	r.HandleFunc("/interpret", interpretMessageResponse).Queries("msg", "{msg}").Methods("GET")
	//r.HandleFunc("/interpretDisplay", interpretMessageDisplayResponse).Queries("msg", "{msg}").Methods("GET")

	r.HandleFunc("/add", addAcronymResponse).Queries("acronym", "{acronym}", "def", "{def}").Methods("GET")
	r.HandleFunc("/del", deleteAcronymResponse).Queries("acronym", "{acronym}").Methods("GET")
	r.HandleFunc("/search", searchAcronymResponse).Queries("acronym", "{acronym}").Methods("GET")
	r.HandleFunc("/load", loadPresetResponse).Queries("preset", "{preset}").Methods("GET")

	r.HandleFunc("/sendText", sendText).Queries("phoneNumber", "{phoneNumber}", "msg", "{msg}").Methods("GET")
	r.HandleFunc("/sendEmail", sendEmail).Queries("email", "{email}", "msg", "{msg}").Methods("GET")

	r.HandleFunc("/remove", removeResponse).Queries("remove", "{bool}").Methods("GET")

	url := ""
	switch runtime.GOOS {
	case "windows":
		url = "../Client/"
	case "darwin":
		url = "../Client/"
	default:
		url = "../Client/"
	}

	r.PathPrefix("/").Handler(http.FileServer(http.Dir(url)))

	open("http://localhost:8080/index.html")

	http.ListenAndServe(":8080", r)

	fmt.Println(acronyms)
}

func interpretMessageResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	message := params["msg"]

	for index, element := range acronyms {
		message = strings.Replace(message, index, element, 1)
	}

	blockLetters := []rune{'🅰', '🅱', '🅲', '🅳', '🅴', '🅵', '🅶', '🅷', '🅸', '🅹', '🅺', '🅻', '🅼', '🅽', '🅾', '🅿', '🆀', '🆁', '🆂', '🆃', '🆄', '🆅', '🆆', '🆇', '🆈', '🆉'}

	output := []rune(message)

	memeifi := false
	emojifi := false
	for i := 0; i < len(message); i++ {

		if string(message[i]) == "<" {
			memeifi = true
			output[i] = rune(32)
		}

		if string(message[i]) == ">" {
			memeifi = false
			output[i] = rune(32)
		}

		if memeifi == true && output[i] != ' ' {
			if i%2 == 0 {
				output[i] = rune(byte(output[i]) - 32)
			}
		}

		if string(message[i]) == "[" {
			emojifi = true
			output[i] = 32
		}

		if string(message[i]) == "]" {
			emojifi = false
			output[i] = 32
		}

		if emojifi == true && output[i] != rune(' ') {
			if byte(message[i]) >= 97 && byte(message[i]) <= 122 {
				output[i] = rune(blockLetters[output[i]-97])
			} else if byte(message[i]) >= 65 && byte(message[i]) <= 90 {
				output[i] = rune(blockLetters[output[i]-65])
			}

		}

	}

	fmt.Println(r.Method + " recieved with param " + "'" + params["msg"] + "'" + " Returned: " + message)

	json.NewEncoder(w).Encode(string(output))
}

func removeResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)
	if params["bool"] == "true" {
		for key, element := range acronyms {
			delete(acronyms, key)
			fmt.Println(element + " Removed")
		}
	}

	fmt.Println("remove")

	json.NewEncoder(w).Encode("All Acronyms removed")
}

func addAcronymResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)
	acronyms[params["acronym"]] = params["def"]

	fmt.Println(r.Method + " recieved with params " + "'" + params["acronym"] + "'" + ", " + "'" + params["def"] + "'")

	json.NewEncoder(w).Encode("Acronym added!")
}

func deleteAcronymResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)
	delete(acronyms, params["acronym"])

	fmt.Println(r.Method + " recieved with param " + "'" + params["acronym"] + "'")

	json.NewEncoder(w).Encode("Acronym removed!")
}

func searchAcronymResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	fmt.Println(r.Method + " recieved with param " + "'" + params["acronym"] + "'")

	if params["acronym"] == "ALL" {
		for index, element := range acronyms {
			json.NewEncoder(w).Encode(index + ": " + element)
		}
	} else {
		for index, element := range acronyms {
			if params["acronym"] == index {
				json.NewEncoder(w).Encode(index + ": " + element)
				return
			}
		}

		json.NewEncoder(w).Encode("None")
	}
}

func loadPresetResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	fmt.Println(r.Method + " recieved with param " + "'" + params["preset"] + "'")

	if params["preset"] != "medical" && params["preset"] != "piloting" && params["preset"] != "everyday" {
		return
	}

	path := "./Data/" + params["preset"] + ".json"

	for index, element := range loadPreset(path) {
		data := loadPreset(path)
		acronyms[data[index].Acronym] = element.Def
	}
}

func sendEmail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	from := "customacronyms@gmail.com"
	pass := "ca#gmail1"
	to := params["email"]

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Textify Sent You a Message\n\n" +
		params["msg"]

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
}

func sendText(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := "1b4c064b"
	secret := "MG18mzvuWLVMbszK"
	auth := nexmo.NewAuthSet()
	auth.SetAPISecret(key, secret)

	client := nexmo.NewClient(http.DefaultClient, auth)

	smsContent := nexmo.SendSMSRequest{
		From: "13077351345",
		To:   params["phoneNumber"],
		Text: params["msg"] + "\n",
	}

	smsResponse, _, err := client.SMS.SendSMS(smsContent)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Status:", smsResponse.Messages[0].Status)
}

func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default:
		cmd = "cmd"
	}

	args = append(args, url)

	return exec.Command(cmd, args...).Start()
}
