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

// CreateInitialFolder creates the root folder for logs and its subfolders
//
// Parameters:
// folder - a list of subfolder names to create under the root folder
//
// Returns:
// none
func CreateInitialFolder(folder []string) {
	// Create first the root folder for logs
	initLogFolder := "./logs/"
	CreateDirectory(initLogFolder)

	// Create the subfolders
	for _, v := range folder {
		CreateDirectory(fmt.Sprintf("%s/%s", initLogFolder, v))
	}
}

// CreateDirectory creates a directory if it does not exist
//
// Parameters:
// path - the path to the directory to be created
//
// Returns:
// error - an error if the directory already exists or if there is an error
//
//	creating the directory
func CreateDirectory(path string) error {
	// Check if the directory already exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Create the directory with permissions 0755
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
		return nil
	} else {
		return err
	}
}

// SystemLogger writes a log entry to a file with the given class, folder, filename, process name, request and response.
//
// Parameters:
// class - the name of the class that is writing the log entry
// folder - the folder path where the log file will be written
// filename - the filename of the log file that will be written
// process - the name of the process that is writing the log entry
// request - the request object that will be logged
// response - the response object that will be logged
//
// Returns:
// none
//
// Example:
// folder := "test"
// filename := "test"
// process := "test"
// request := map[string]string{"key": "value"}
// response := map[string]string{"key": "value"}
// SystemLogger("test", folder, filename, process, request, response)
//
// Output:
// 2023/10/24 14:30:00 INFO: TEST: - - - - : TEST : - - - -
// 2023/10/24 14:30:00 INFO: TEST: PROCESS TIME: 2023-10-24 14:30:00
// 2023/10/24 14:30:00 INFO: TEST: REQUEST: {"key":"value"}
// 2023/10/24 14:30:00 INFO: TEST: RESPONSE: {"key":"value"}
func SystemLogger(class, folder, filename, process string, request, response interface{}) {
	// Checking folder name if exists
	currentTime := time.Now()
	folderName := "./logs/" + folder + "/" + currentTime.Format("01-January")
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
	file.Close()
}
func ApplicationLogger(class, folder, filename, process, status string, request, response any) {
	// Checking folder name if exists
	currentTime := time.Now()
	folderName := "./logs/" + folder + "/" + currentTime.Format("01-January")
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
	file.Close()
}

// ApplicationErrorLogger logs the application logs in the specified folder and filename.
// The file is created if it does not exist and the logs are appended to the file.
// The logs are in the format of:
// 2023/10/24 14:30:00 ERROR: CLASS: - - - - : PROCESS : - - - -
// 2023/10/24 14:30:00 ERROR: CLASS: PROCESS TIME: 2023-10-24 14:30:00
// 2023/10/24 14:30:00 ERROR: CLASS: REQUEST: {"key":"value"}
// 2023/10/24 14:30:00 ERROR: CLASS: CODE: 200
// 2023/10/24 14:30:00 ERROR: CLASS: ERROR: {"key":"value"}
//
// Parameters:
// class - the class name of the application
// folder - the folder name where the log file is located
// filename - the filename of the log file
// process - the name of the process
// code - the code of the request
// request - the request object
// err - the error object
//
// Example:
// folder := "test"
// filename := "test"
// process := "test"
// request := map[string]string{"key": "value"}
// err := map[string]string{"key": "value"}
// code := "200"
// ApplicationErrorLogger("test", folder, filename, process, code, request, err)
//
// Output:
// 2023/10/24 14:30:00 ERROR: TEST: - - - - : TEST : - - - -
// 2023/10/24 14:30:00 ERROR: TEST: PROCESS TIME: 2023-10-24 14:30:00
// 2023/10/24 14:30:00 ERROR: TEST: REQUEST: {"key":"value"}
// 2023/10/24 14:30:00 ERROR: TEST: CODE: 200
// 2023/10/24 14:30:00 ERROR: TEST:
func ApplicationErrorLogger(class, folder, filename, process, code string, request, err any) {
	// Checking folder name if exists
	currentTime := time.Now()
	folderName := "./logs/" + folder + "/" + currentTime.Format("01-January")
	CreateDirectory(folderName)
	file, filErr := os.OpenFile(folderName+"/"+strings.ToLower(filename)+"-"+currentTime.Format("01022006")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if filErr != nil {
		fmt.Println("error wrting the logs:", filErr.Error())
	}

	strRequest, _ := json.Marshal(request)
	strResponse, _ := json.Marshal(err)

	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime)
	Separator = log.New(file, "", log.Ldate|log.Ltime)

	Separator.Println("")
	ErrorLogger.Println(class + ": - - - - : " + strings.ToUpper(process) + " : - - - -")
	ErrorLogger.Println(class + ": PROCESS TIME: " + currentTime.Format(time.DateTime))
	ErrorLogger.Println(class + ": REQUEST: " + string(strRequest))
	ErrorLogger.Println(class + ": ERROR CODE: " + code)
	ErrorLogger.Println(class + ": ERROR MESSAGE: " + string(strResponse))

	fmt.Printf("New entry for %s: %v\n", strings.ToUpper(process), currentTime.Format(time.DateTime))
	file.Close()
}
