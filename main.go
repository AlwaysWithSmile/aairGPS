package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	//"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

//const connect_db string = "root:BSDcloud2019@tcp(localhost:3306)/skmaircloud"

var (
	skm_dbcmd       bool
	start_pass      string
	skm_dbname      string
	skm_host        string
	skm_port        string
	sg_host         string
	sg_port         string
	sg_time         string
	skm_user        string
	skm_password    string
	connect_db      string
	gps_socket_port string
	starting_log    string

	uptAlarm  bool
	uptList   bool
	lastAlarm float64
	lastUIN   int64
)

//=============== Reads info from config file===============================
func ReadConfig() {
	var (
		air_param_int int
		air_param_str string
	)
	skm_dbcmd = false
	uptAlarm = false
	uptList = false
	lastAlarm = 0
	lastUIN = 0
	start_pass = "admin"
	starting_log := ""
	connect_db = "root:BSDcloud2019@tcp(localhost:3306)/skmaircloud"
	skm_dbname = "skmaircloud"
	skm_host = "localhost"
	skm_port = ":3306"

	sg_host = "192.168.88.178"
	sg_port = "5020"
	sg_time = "650"

	gps_socket_port = ":9090"

	skm_user = "root"
	skm_password = "BSDcloud2019"
	file, err := os.Open("bsdbroker.cfg")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		air_param_str = scanner.Text()
		air_param_int = strings.Index(air_param_str, "=")
		if air_param_int > 0 {
			if air_param_int < len(air_param_str) {
				if strings.Index(air_param_str, "startpass=") == 0 {
					start_pass = air_param_str[air_param_int+1 : len(air_param_str)]
					starting_log += getDT() + "startpass:" + start_pass + string(13) + string(10)
				} else if strings.Index(air_param_str, "dbname=") == 0 {
					skm_dbname = air_param_str[air_param_int+1 : len(air_param_str)]
					starting_log += getDT() + "dbname:" + skm_dbname + string(13) + string(10)
				} else if strings.Index(air_param_str, "dbhost=") == 0 {
					skm_host = air_param_str[air_param_int+1 : len(air_param_str)]
					starting_log += getDT() + "dbhost:" + skm_host + string(13) + string(10)
				} else if strings.Index(air_param_str, "dbport=") == 0 {
					skm_port = ":" + air_param_str[air_param_int+1:len(air_param_str)]
					starting_log += getDT() + "dbport" + skm_port + string(13) + string(10)
				} else if strings.Index(air_param_str, "dbuser=") == 0 {
					skm_user = air_param_str[air_param_int+1 : len(air_param_str)]
					starting_log += getDT() + "dbuser:" + skm_user + string(13) + string(10)
				} else if strings.Index(air_param_str, "dbpassword=") == 0 {
					skm_password = air_param_str[air_param_int+1 : len(air_param_str)]
					starting_log += getDT() + "dbpassword:" + skm_password + string(13) + string(10)
				} else if strings.Index(air_param_str, "gpsport=") == 0 {
					gps_socket_port = ":" + air_param_str[air_param_int+1:len(air_param_str)]
					starting_log += getDT() + "gpsport" + gps_socket_port + string(13) + string(10)
				}

			}
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	connect_db = skm_user + ":" + skm_password + "@tcp(" + skm_host + skm_port + ")/" + skm_dbname
	fmt.Println("Connection String:" + connect_db)

}

//============================ TIMER FUNCTIONS ================================

func gpsTiming() {

	if len(starting_log) > 10 {
		fmt.Println(starting_log)
		starting_log = ""
	}

	tickerSG := time.NewTicker(time.Second) //20 *
	b_time := false

	for now := range tickerSG.C {
		if b_time {
			fmt.Println("Check timer picker", now)
		}

		checkUpdateAlarms()
		searchMyAlarms()
		broker_upgrade()
	}

}

//============================ SOCKET FUNCTIONS ================================

func webProcesses() {
	fmt.Println("Start websocket", gps_socket_port)
	websock_addr_user = make([]*websocket.Conn, 0)
	websock_last_connect = make([]int64, 0)
	websock_uin_users = make([]int, 0)
	websock_send_repeat = make([]int, 0)
	websock_send_device = make([]string, 0)
	// Configure websocket route
	http.HandleFunc("/ws", wsHandler)
	//http.HandleFunc("/", rootHandler)
	if err := http.ListenAndServe(gps_socket_port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

//*****************************************************************************
func main() {
	fmt.Println("START GPS SERVICE")
	ReadConfig()
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		webProcesses()
	}()

	go func() {
		defer wg.Done()
		gpsTiming()
	}()
	starting_log += getDT() + "Waiting To Finish" + string(13) + string(10)
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
