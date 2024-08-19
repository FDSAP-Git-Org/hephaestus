package apilogs

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	Separator     *log.Logger
)

// Used to create initial folders for logs
func CreateInitialFolder(folder []string) {
	// Create first the root folder for logs
	initLogFolder := "./logs/"
	CreateDirectory(initLogFolder)

	for _, v := range folder {
		CreateDirectory(fmt.Sprintf("%s/%s", initLogFolder, v))
	}
}

func CreateDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
		return nil
	} else {
		return err
	}
}

// Default system log
func SystemLog(class, folder, filename, process string, request, response interface{}) {
	// Checking folder name if exists
	currentTime := time.Now()
	folderName := "./logs/" + strings.ToUpper(folder) + "/" + currentTime.Format("01-January")
	CreateDirectory(folderName)
	file, filErr := os.OpenFile(folderName+"/"+strings.ToLower(filename)+"-"+currentTime.Format("01022006")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if filErr != nil {
		fmt.Println("error wrting the logs:", filErr.Error())
	}

	strRequest, _ := json.Marshal(request)
	strResponse, _ := json.Marshal(response)

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	Separator = log.New(file, "", log.Ldate|log.Ltime)

	Separator.Println("")
	InfoLogger.Println(class + ": - - - - : " + strings.ToUpper(process) + " : - - - -")
	InfoLogger.Println(class + ": PROCESS TIME: " + currentTime.Format(time.DateTime))
	InfoLogger.Println(class + ": REQUEST: " + string(strRequest))
	InfoLogger.Println(class + ": RESPONSE: " + string(strResponse))

	fmt.Printf("New entry for %s: %v\n", strings.ToUpper(process), currentTime.Format(time.DateTime))
}

func APIlog(class, folder, filename, process, status string, request, response interface{}) {
	// Checking folder name if exists
	currentTime := time.Now()
	folderName := "./logs/" + strings.ToUpper(folder) + "/" + currentTime.Format("01-January")
	CreateDirectory(folderName)
	file, filErr := os.OpenFile(folderName+"/"+strings.ToLower(filename)+"-"+currentTime.Format("01022006")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if filErr != nil {
		fmt.Println("error wrting the logs:", filErr.Error())
	}

	strRequest, _ := json.Marshal(request)
	strResponse, _ := json.Marshal(response)

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	Separator = log.New(file, "", log.Ldate|log.Ltime)

	Separator.Println("")
	InfoLogger.Println(class + ": - - - - : " + strings.ToUpper(process) + " : - - - -")
	InfoLogger.Println(class + ": PROCESS TIME: " + currentTime.Format(time.DateTime))
	InfoLogger.Println(class + ": REQUEST: " + string(strRequest))
	InfoLogger.Println(class + ": RESPONSE: " + string(strResponse))
	InfoLogger.Println(class + ": STATUS: " + status)

	fmt.Printf("New entry for %s: %v\n", strings.ToUpper(process), currentTime.Format(time.DateTime))
}
