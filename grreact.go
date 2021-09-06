package main

import (
	"encoding/json"
	"fmt"
	"time"

	//	"math/rand"
	"strconv"

	"github.com/gorilla/websocket"
)

type AirQuery struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Cmnd  string `json:"cmnd"`
	Param string `json:"param"`
}

//========================= MAIN LOGIC =======================================
func decodeGpsJson(jsonIncoming string, conn *websocket.Conn) string {
	var (
		airDecoding AirQuery
		strJSON     []byte
		//		i_con       int
		js_result string
		js_iden   string
		js_cmnd   string
		js_param  string
		js_name   string
	)
	js_result = "{" + string(0x0D) + string(0x0A)
	js_result += getQuatedJSON("param", "Status error", 1) + string(0x0D) + string(0x0A)
	js_result += "}" + string(0x0D) + string(0x0A)

	//Error json format
	if checkValidJson(jsonIncoming) == true { //Check valid data
		js_result = "Status 8" //Error data
		strJSON = []byte(jsonIncoming)
		err := json.Unmarshal(strJSON, &airDecoding)
		if err != nil {
			defer recoveryAppFunction()
			fmt.Println(getDT(), "Error decoding Json:"+jsonIncoming)
			panic(err)
		}
		js_iden = airDecoding.ID
		js_name = airDecoding.Name
		js_cmnd = airDecoding.Cmnd
		js_param = airDecoding.Param

		switch js_cmnd {
		case "start": //First start
			js_result = startGBR(js_iden, js_name, js_param, getSocketIndex(conn))
		case "login": //Loging for user
			js_result = logGBR(js_iden, js_name, js_param, getSocketIndex(conn))
		case "alarmlist": //Get alarm list
			js_result = getAlarms(js_iden, js_name, js_param)
		case "alarmget": //Receive alarm
			js_result = recAlarms(js_iden, js_cmnd, js_name, js_param)
		case "alarmstart": //GBR starts trip
			js_result = procAlarm(js_iden, js_cmnd, js_name, js_param)
		case "alarmpoint": //GBR at point
			js_result = procAlarm(js_iden, js_cmnd, js_name, js_param)
		case "alarmbreak": //Problem with GBR
			js_result = procAlarm(js_iden, js_cmnd, js_name, js_param)
		case "alarmstop": //Set reaction
			js_result = procAlarm(js_iden, js_cmnd, js_name, js_param)
		case "alarminfo": //Read updates
			js_result = getAlarmInfo(js_iden, js_cmnd, js_name, js_param)
		default:
			js_result = setUnknown(js_iden, js_name, js_cmnd)
		}

	}

	return js_result
}

//------------------------------------------------------------------------------
func sendUpdator(userid int) string {
	s_json := "{" + string(0x0D) + string(0x0A)
	s_json = s_json + getQuatedJSON("id", strconv.Itoa(websock_uin_users[userid]), 1) + "," + string(0x0D) + string(0x0A)
	s_json = s_json + getQuatedJSON("cmnd", "update", 1) + string(0x0D) + string(0x0A)
	s_json = s_json + "}" + string(0x0D) + string(0x0A)
	return s_json
}

//------------------------------------------------------------------------------
func setUnknown(userid, js_name, js_command string) string {
	fmt.Println(getDT(), "Command unknown"+js_command)
	s_json := "{" + string(0x0D) + string(0x0A)
	s_json = s_json + getQuatedJSON("id", userid, 1) + "," + string(0x0D) + string(0x0A)
	s_json = s_json + getQuatedJSON("name", js_name, 1) + "," + string(0x0D) + string(0x0A)
	s_json = s_json + getQuatedJSON("cmnd", js_command, 1) + "," + string(0x0D) + string(0x0A)
	s_json = s_json + getQuatedJSON("param", "STATUS_ERROR", 1) + string(0x0D) + string(0x0A)
	s_json = s_json + "}" + string(0x0D) + string(0x0A)

	return s_json
}

//------------------------------------------------------------------------------
func getAlarmInfo(userid, js_command, js_name, js_param string) string {
	updateGBRstatus(userid, getGBRuser(userid), js_param, "", 5)
	s_json := "{" + string(0x0D) + string(0x0A)
	s_temp := getALARMlist(js_name, userid)
	i_arm, i_con := getObjectStatus(js_name)
	if len(s_temp) > 10 {
		s_json += getQuatedJSON("alarminfo", "[", 0) + string(0x0D) + string(0x0A)
		s_json += "{" + string(0x0D) + string(0x0A)
		s_json += getQuatedJSON("id", userid, 1) + "," + string(0x0D) + string(0x0A)
		s_json += getQuatedJSON("name", js_name, 1) + "," + string(0x0D) + string(0x0A)
		s_json += getQuatedJSON("status", i_arm, 1) + "," + string(0x0D) + string(0x0A)
		s_json += getQuatedJSON("con", i_con, 1) + "," + string(0x0D) + string(0x0A)
		s_json += getQuatedJSON("param", "ALARM_VALID", 1) + string(0x0D) + string(0x0A)
		s_json += "}]," + string(0x0D) + string(0x0A)
		s_json += getQuatedJSON("eventlist", "[", 0) + string(0x0D) + string(0x0A)
		s_json += s_temp
		s_json += "]" + string(0x0D) + string(0x0A)
	} else {
		s_json += getQuatedJSON("id", userid, 1) + "," + string(0x0D) + string(0x0A)
		s_json += getQuatedJSON("name", js_name, 1) + "," + string(0x0D) + string(0x0A)
		s_json += getQuatedJSON("cmnd", js_command, 1) + "," + string(0x0D) + string(0x0A)
		s_json += getQuatedJSON("param", "ALARM_EMPTY", 1) + string(0x0D) + string(0x0A)
	}
	s_json = s_json + "}" + string(0x0D) + string(0x0A)
	s_json = doReplaceStr(s_json, "},]", "}]")
	return s_json
}

//------------------------------------------------------------------------------
func startGBR(userid, js_name, js_param string, conPosition int) string {
	//{"cmnd":"start","id":"0","name":"token","param":"semen2021"}
	s_json := ""
	if start_pass == js_param { //Device loging
		//userid - IMEI of device; js_name - google accound addr; js_param - password
		s_json = "{" + string(0x0D) + string(0x0A)
		s_json += getQuatedJSON("reglist", "[", 0) + string(0x0D) + string(0x0A)
		s_json += getGBRlist(0) + "],"
		s_json += getQuatedJSON("gbrlist", "[", 0) + string(0x0D) + string(0x0A)
		s_json += getGBRlist(1) + "],"
		//s_json = doReplaceStr(s_json, "},]", "}]")
		s_json += getQuatedJSON("usrlist", "[", 0) + string(0x0D) + string(0x0A)
		s_json += getGBRlist(2)
		s_json += "]" + string(0x0D) + string(0x0A)
		s_json += "}" + string(0x0D) + string(0x0A)

		s_json = doReplaceStr(s_json, "},]", "}]")

		if conPosition >= 0 && conPosition < websock_addr_counter {
			websock_send_device[conPosition] = js_name
		}

	} else {
		s_json = "{" + string(0x0D) + string(0x0A)
		s_json += getQuatedJSON("id", userid, 1) + "," + string(0x0D) + string(0x0A)
		s_json += getQuatedJSON("name", js_name, 1) + "," + string(0x0D) + string(0x0A)
		s_json += getQuatedJSON("cmnd", "start", 1) + "," + string(0x0D) + string(0x0A)
		s_json += getQuatedJSON("param", "START_ERR", 1) + string(0x0D) + string(0x0A)
		s_json += "}" + string(0x0D) + string(0x0A)
	}
	return s_json
}

//------------------------------------------------------------------------------
func logGBR(userid, js_name, js_param string, conPosition int) string {

	s_sql := "SELECT IDCONST, CONSTVALUE FROM consttable WHERE "
	s_sql += "(CONSTVISIB = 0) AND (CONSTKIND = 4) AND (IDCONST = " + userid
	s_sql += ") LIMIT 1"

	gbr_valid := (dbGetIntData(s_sql, 0) > 0)

	s_sql = "SELECT IDPERS,FIOPERS,PAROL FROM personality WHERE IDPERS=" + dbQuatedString(js_name)
	s_json := ""

	if gbr_valid == false {
		s_json = "{" + string(0x0D) + string(0x0A)
		s_json = s_json + getQuatedJSON("id", userid, 1) + "," + string(0x0D) + string(0x0A)
		s_json = s_json + getQuatedJSON("name", js_name, 1) + "," + string(0x0D) + string(0x0A)
		s_json = s_json + getQuatedJSON("cmnd", "login", 1) + "," + string(0x0D) + string(0x0A)
		s_json = s_json + getQuatedJSON("param", "GBR_ERR", 1) + string(0x0D) + string(0x0A)
		s_json = s_json + "}" //+ string(0x0D) + string(0x0A)
	} else if len(js_name) < 1 || len(js_param) < 1 { //Input data is empty
		s_json = "{" + string(0x0D) + string(0x0A)
		s_json = s_json + getQuatedJSON("id", userid, 1) + "," + string(0x0D) + string(0x0A)
		s_json = s_json + getQuatedJSON("name", js_name, 1) + "," + string(0x0D) + string(0x0A)
		s_json = s_json + getQuatedJSON("cmnd", "login", 1) + "," + string(0x0D) + string(0x0A)
		s_json = s_json + getQuatedJSON("param", "LOG_EMPTY", 1) + string(0x0D) + string(0x0A)
		s_json = s_json + "}" //+ string(0x0D) + string(0x0A)
	} else { //Not EMpty data
		s_psw := dbGetStringData(s_sql, 2)
		if len(s_psw) < 1 { //NOT LOGIN ENABLE
			s_json = "{" + string(0x0D) + string(0x0A)
			s_json = s_json + getQuatedJSON("id", userid, 1) + "," + string(0x0D) + string(0x0A)
			s_json = s_json + getQuatedJSON("name", js_name, 1) + "," + string(0x0D) + string(0x0A)
			s_json = s_json + getQuatedJSON("cmnd", "login", 1) + "," + string(0x0D) + string(0x0A)
			s_json = s_json + getQuatedJSON("param", "LOG_FALSE_L", 1) + string(0x0D) + string(0x0A)
			s_json = s_json + "}" //+ string(0x0D) + string(0x0A)
		} else if s_psw == js_param { //ALL OK
			s_json = "{" + string(0x0D) + string(0x0A)
			s_json = s_json + getQuatedJSON("id", userid, 1) + "," + string(0x0D) + string(0x0A)
			s_json = s_json + getQuatedJSON("name", js_name, 1) + "," + string(0x0D) + string(0x0A)
			s_json = s_json + getQuatedJSON("cmnd", "login", 1) + "," + string(0x0D) + string(0x0A)
			s_json = s_json + getQuatedJSON("param", "LOG_OK", 1) + string(0x0D) + string(0x0A)
			s_json = s_json + "}" //+ string(0x0D) + string(0x0A)
			s_tocken := ""
			if conPosition >= 0 && conPosition < websock_addr_counter {
				s_tocken = websock_send_device[conPosition]
				websock_uin_users[conPosition] = convertIntVal(userid)
			}
			fmt.Println("Try update tocken", userid, js_name, s_tocken)
			updateGBRstatus(userid, js_name, "", s_tocken, 0)

		} else { //NOT PASSWORD ENABLE
			s_json = "{" + string(0x0D) + string(0x0A)
			s_json = s_json + getQuatedJSON("id", userid, 1) + "," + string(0x0D) + string(0x0A)
			s_json = s_json + getQuatedJSON("name", js_name, 1) + "," + string(0x0D) + string(0x0A)
			s_json = s_json + getQuatedJSON("cmnd", "login", 1) + "," + string(0x0D) + string(0x0A)
			s_json = s_json + getQuatedJSON("param", "LOG_FALSE_P", 1) + string(0x0D) + string(0x0A)
			s_json = s_json + "}" //+ string(0x0D) + string(0x0A)
		}
	}
	return s_json
}

//------------------------------------------------------------------------------
func getAlarms(userid, js_name, js_param string) string {
	s_json := "{" + string(0x0D) + string(0x0A)
	s_alarms := getALARMlist("", userid)
	if len(s_alarms) > 10 {
		s_json += getQuatedJSON("alarmlist", "[", 0) + string(0x0D) + string(0x0A)
		s_json += s_alarms
		s_json += "]" + string(0x0D) + string(0x0A)
		s_json += "}" + string(0x0D) + string(0x0A)
		s_json = doReplaceStr(s_json, "},]", "}]")
	} else {
		s_json = "{" + string(0x0D) + string(0x0A)
		s_json = s_json + getQuatedJSON("id", userid, 1) + "," + string(0x0D) + string(0x0A)
		s_json = s_json + getQuatedJSON("name", js_name, 1) + "," + string(0x0D) + string(0x0A)
		s_json = s_json + getQuatedJSON("cmnd", "alarmlist", 1) + "," + string(0x0D) + string(0x0A)
		s_json = s_json + getQuatedJSON("param", "Status empty", 1) + string(0x0D) + string(0x0A)
		s_json = s_json + "}" + string(0x0D) + string(0x0A)
	}

	return s_json
}

//------------------------------------------------------------------------------
func recAlarms(userid, js_cmnd, js_name, js_param string) string {
	//procPosition(userid, js_cmnd, js_name)
	//updateGBRstatus(userid, getGBRuser(userid), "", js_param, 1)

	s_json := "{" + string(0x0D) + string(0x0A) + getObjGeneral(js_name, false)
	s_temp := getZoneUserList(js_name, 0)
	//userid - GBR ID
	//js_name - OBJECT ID
	if len(s_temp) > 10 {
		s_json += "," + string(0x0D) + string(0x0A) + getQuatedJSON("zonelist", "[", 0) + string(0x0D) + string(0x0A)
		s_json += s_temp
		s_json += "]"
	}
	s_temp = getZoneUserList(js_name, 1)
	if len(s_temp) > 10 {
		s_json += "," + string(0x0D) + string(0x0A) + getQuatedJSON("userlist", "[", 0) + string(0x0D) + string(0x0A)
		s_json += s_temp
		s_json += "]"
	}

	s_temp = getZoneUserList(js_name, 2)
	if len(s_temp) > 10 {
		s_json += "," + string(0x0D) + string(0x0A) + getQuatedJSON("imagelist", "[", 0) + string(0x0D) + string(0x0A)
		s_json += s_temp
		s_json += "]" + string(0x0D) + string(0x0A)
	}

	s_temp = getALARMlist(js_name, userid)
	if len(s_temp) > 10 {
		s_json += "," + string(0x0D) + string(0x0A) + getQuatedJSON("eventlist", "[", 0) + string(0x0D) + string(0x0A)
		s_json += s_temp
		s_json += "]" + string(0x0D) + string(0x0A)
	}
	/*	*/
	s_json += "}" + string(0x0D) + string(0x0A)
	s_json = doReplaceStr(s_json, "},]", "}]")
	return s_json
}

//------------------------------------------------------------------------------
func procAlarm(userid, js_cmnd, js_name, js_param string) string {
	s_status := "ALARM_ERR"
	switch js_cmnd {
	case "alarmstart": //GBR starts trip
		s_status = "START_OK"
	case "alarmpoint": //GBR at point
		s_status = "POINT_OK"
	case "alarmbreak": //Problem with GBR
		s_status = "BREAK_OK"
	case "alarmstop": //Set reaction
		s_status = "STOP_OK"
	}

	switch js_cmnd {
	case "alarmstart": //GBR starts trip
		updateGBRstatus(userid, getGBRuser(userid), "", js_param, 1)
	case "alarmpoint": //GBR at point
		updateGBRstatus(userid, getGBRuser(userid), "", js_param, 2)
	case "alarmbreak": //Problem with GBR
		updateGBRstatus(userid, getGBRuser(userid), "", js_param, 3)
	case "alarmstop": //Set reaction
		updateGBRstatus(userid, getGBRuser(userid), "", js_param, 4)
	}

	procPosition(userid, js_cmnd, js_name)
	s_json := "{" + string(0x0D) + string(0x0A)
	s_json = s_json + getQuatedJSON("id", userid, 1) + "," + string(0x0D) + string(0x0A)
	s_json = s_json + getQuatedJSON("name", js_name, 1) + "," + string(0x0D) + string(0x0A)
	s_json = s_json + getQuatedJSON("cmnd", js_cmnd, 1) + "," + string(0x0D) + string(0x0A)
	s_json = s_json + getQuatedJSON("param", s_status, 1) + string(0x0D) + string(0x0A)
	s_json = s_json + "}" + string(0x0D) + string(0x0A)

	return s_json
}

//------------------------------------------------------------------------------
func procPosition(userid, js_cmnd, js_name string) {
	s_sql := ""
	s_time := delphiDateToSQL(time.Now().Unix())
	switch js_cmnd {
	case "alarmstart": //Receive alarm
		s_sql = "UPDATE eventlist SET GBRID=" + userid + ", ISSEND=" + s_time
	case "alarmpoint": //GBR at point
		s_sql = "UPDATE eventlist SET ISPRIB=" + s_time
	case "alarmbreak": //Problem with GBR
		s_sql = "UPDATE eventlist SET GBRID=0,ISGBR=0,ISSEND=0,ISPRIB=0,ISOTBOY=" + s_time
	case "alarmstop": //Stop alarm
		s_sql = "UPDATE eventlist SET ISFINISH=" + s_time
	}
	if len(s_sql) > 5 {
		s_sql += " WHERE (ISNEW>0) AND (OBJECTID=" + js_name + ")"
		dbUpdateData(s_sql)
		i_time := time.Now().Unix()
		dbUpdateData("UPDATE equiplist SET SHVYDKIST=" + delphiDateToSQL(i_time))
		lastAlarm = UNIXTimeToDateTimeFAST(i_time)
		uptAlarm = true
	}
	/*
		if js_cmnd == "alarmbreak" {

		}
	*/
	return
}

//------------------------------------------------------------------------------
func searchMyAlarms() { //New info Presents
	if dbGetIntData("SELECT COUNT(IDREG) AS SEARCHFLD FROM regtemp WHERE REGRESULT<4", 0) > 0 {
		sendAlarmToGbr()
		dbUpdateData("DELETE FROM regtemp WHERE REGRESULT<4")
		updateSockList()
	}
}
