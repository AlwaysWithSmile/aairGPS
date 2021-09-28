package main

import (
	"bytes"
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

type gbrList []struct{

	Region string `json:"region"`
	Number string `json:"number"`
	GbrlistArray []string `json:"gbrlist"`
}

type AppSend struct{
	Command string `json:"cmnd"`
	ID string `json:"id"`
}

//========================= get JSON func =======================================
func getJSON(url string, target interface{})error{
	return json.NewDecoder(bytes.NewBufferString(url)).Decode(target)
}

func bla(userid string) bool {
	mytable := StringTable{
		1:"71",
		2:"72",
		3:"73",
		4:"74",
		5:"75",
		6:"78",
		7:"79",
		8:"80",
		9:"81",
		10:"82",
		11:"83",
		12:"84",
		13:"85",
		14:"86",
		15:"88",
		16:"89",
		17:"92",
		18:"1",
		19:"2",
		20:"3",
		21:"4",
		22:"5",
		23:"6",
		24:"7",
		25:"1",
		26:"2",
		27:"3",
		28:"4",
		29:"5",
		30:"6",
		31:"1",
		32:"2",
		35:"1",
	}
	intvar,_ := strconv.Atoi(userid)
	if(mytable.Get(intvar) == "") {
		return false
	}
	return true
}
var gbrJsonRawEscaped json.RawMessage
var gbrJsonRawUnescaped json.RawMessage
//========================= MAIN LOGIC =======================================
func decodeGpsJson(jsonIncoming string, conn *websocket.Conn) string {
	var gbrlistSout gbrList
	getJSON("http://api-cs.ohholding.com.ua/gbr_list/get", &gbrlistSout)
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


		getJSON("http://api-cs.ohholding.com.ua/gbr_list/get", &gbrlistSout)
		b := []byte(`{"gbrlist":[{"id_gbr":1,"region":"\u041a\u0438\u0435\u0432","number":"71"},{"id_gbr":2,"region":"\u041a\u0438\u0435\u0432","number":"72"},{"id_gbr":3,"region":"\u041a\u0438\u0435\u0432","number":"73"},{"id_gbr":4,"region":"\u041a\u0438\u0435\u0432","number":"74"},{"id_gbr":5,"region":"\u041a\u0438\u0435\u0432","number":"75"},{"id_gbr":6,"region":"\u041a\u0438\u0435\u0432","number":"78"},{"id_gbr":7,"region":"\u041a\u0438\u0435\u0432","number":"79"},{"id_gbr":8,"region":"\u041a\u0438\u0435\u0432","number":"80"},{"id_gbr":9,"region":"\u041a\u0438\u0435\u0432","number":"81"},{"id_gbr":10,"region":"\u041a\u0438\u0435\u0432","number":"82"},{"id_gbr":11,"region":"\u041a\u0438\u0435\u0432","number":"83"},{"id_gbr":12,"region":"\u041a\u0438\u0435\u0432","number":"84"},{"id_gbr":13,"region":"\u041a\u0438\u0435\u0432","number":"85"},{"id_gbr":14,"region":"\u041a\u0438\u0435\u0432","number":"86"},{"id_gbr":15,"region":"\u041a\u0438\u0435\u0432","number":"88"},{"id_gbr":16,"region":"\u041a\u0438\u0435\u0432","number":"89"},{"id_gbr":17,"region":"\u041a\u0438\u0435\u0432","number":"92"},{"id_gbr":18,"region":"\u0417\u0430\u043f\u043e\u0440\u043e\u0436\u044c\u0435","number":"\u0411\u0430\u0439\u043a\u0430\u043b 1"},{"id_gbr":19,"region":"\u0417\u0430\u043f\u043e\u0440\u043e\u0436\u044c\u0435","number":"\u0411\u0430\u0439\u043a\u0430\u043b 2"},{"id_gbr":20,"region":"\u0417\u0430\u043f\u043e\u0440\u043e\u0436\u044c\u0435","number":"\u0411\u0430\u0439\u043a\u0430\u043b 3"},{"id_gbr":21,"region":"\u0417\u0430\u043f\u043e\u0440\u043e\u0436\u044c\u0435","number":"\u0411\u0430\u0439\u043a\u0430\u043b 4"},{"id_gbr":22,"region":"\u0417\u0430\u043f\u043e\u0440\u043e\u0436\u044c\u0435","number":"\u0411\u0430\u0439\u043a\u0430\u043b 5"},{"id_gbr":23,"region":"\u0417\u0430\u043f\u043e\u0440\u043e\u0436\u044c\u0435","number":"\u0411\u0430\u0439\u043a\u0430\u043b 6"},{"id_gbr":24,"region":"\u0417\u0430\u043f\u043e\u0440\u043e\u0436\u044c\u0435","number":"\u0411\u0430\u0439\u043a\u0430\u043b 7"},{"id_gbr":25,"region":"\u0414\u043d\u0435\u043f\u0440","number":"\u0414\u043d\u0435\u043f\u0440 1"},{"id_gbr":26,"region":"\u0414\u043d\u0435\u043f\u0440","number":"\u0414\u043d\u0435\u043f\u0440 2"},{"id_gbr":27,"region":"\u0414\u043d\u0435\u043f\u0440","number":"\u0414\u043d\u0435\u043f\u0440 3"},{"id_gbr":28,"region":"\u0414\u043d\u0435\u043f\u0440","number":"\u0414\u043d\u0435\u043f\u0440 4"},{"id_gbr":29,"region":"\u0414\u043d\u0435\u043f\u0440","number":"\u0414\u043d\u0435\u043f\u0440 5"},{"id_gbr":30,"region":"\u0414\u043d\u0435\u043f\u0440","number":"\u0414\u043d\u0435\u043f\u0440 6"},{"id_gbr":31,"region":"\u041a\u0430\u043c\u0435\u043d\u0441\u043a","number":"\u041a\u0430\u043c\u0435\u043d\u0441\u043a 1"},{"id_gbr":32,"region":"\u041a\u0430\u043c\u0435\u043d\u0441\u043a","number":"\u041a\u0430\u043c\u0435\u043d\u0441\u043a 2"},{"id_gbr":35,"region":"\u041a\u0440\u0438\u0432\u043e\u0439 \u0420\u043e\u0433","number":"\u041a\u0440\u0438\u0432\u0431\u0430\u0441 1"},{"id_gbr":36,"region":"\u041a\u0440\u0438\u0432\u043e\u0439 \u0420\u043e\u0433","number":"\u041a\u0440\u0438\u0432\u0431\u0430\u0441 2\r\n"},{"id_gbr":37,"region":"\u041a\u0440\u0438\u0432\u043e\u0439 \u0420\u043e\u0433","number":"\u041a\u0440\u0438\u0432\u0431\u0430\u0441 3"},{"id_gbr":38,"region":"\u041a\u0440\u0438\u0432\u043e\u0439 \u0420\u043e\u0433","number":"\u041a\u0440\u0438\u0432\u0431\u0430\u0441 4"},{"id_gbr":39,"region":"\u041a\u0440\u0438\u0432\u043e\u0439 \u0420\u043e\u0433","number":"\u041a\u0440\u0438\u0432\u0431\u0430\u0441 7"},{"id_gbr":40,"region":"\u041a\u0440\u0438\u0432\u043e\u0439 \u0420\u043e\u0433","number":"\u041a\u0440\u0438\u0432\u0431\u0430\u0441 6"},{"id_gbr":43,"region":"\u0414\u043e\u0431\u0440\u043e\u043f\u043e\u043b\u044c\u0435","number":"\u0421\u043e\u043a\u043e\u043b"},{"id_gbr":49,"region":"\u041b\u044c\u0432\u043e\u0432","number":"\u041b\u044c\u0432\u043e\u0432 1"},{"id_gbr":50,"region":"\u041b\u044c\u0432\u043e\u0432","number":"\u041b\u044c\u0432\u043e\u0432 2"},{"id_gbr":51,"region":"\u041b\u044c\u0432\u043e\u0432","number":"\u041b\u044c\u0432\u043e\u0432 3"},{"id_gbr":52,"region":"\u041b\u044c\u0432\u043e\u0432","number":"\u041b\u044c\u0432\u043e\u0432 4"},{"id_gbr":60,"region":"\u041a\u0438\u0435\u0432","number":"76"},{"id_gbr":65,"region":"\u041c\u0430\u0440\u0438\u0443\u043f\u043e\u043b\u044c","number":"\u041c\u0430\u0440\u0438\u0443\u043f\u043e\u043b\u044c-1"},{"id_gbr":66,"region":"\u041c\u0430\u0440\u0438\u0443\u043f\u043e\u043b\u044c","number":"\u041c\u0430\u0440\u0438\u0443\u043f\u043e\u043b\u044c-2"},{"id_gbr":67,"region":"\u041c\u0430\u0440\u0438\u0443\u043f\u043e\u043b\u044c","number":"\u041c\u0430\u0440\u0438\u0443\u043f\u043e\u043b\u044c-4"},{"id_gbr":72,"region":"\u041a\u0438\u0435\u0432","number":"87"},{"id_gbr":74,"region":"\u041f\u043e\u043a\u0440\u043e\u0432\u0441\u043a\u043e\u0435","number":"\u041f\u043e\u043a\u0440\u043e\u0432\u0441\u043a"},{"id_gbr":76,"region":"\u042d\u043d\u0435\u0440\u0433\u043e\u0434\u0430\u0440","number":"\u042d\u043d\u0435\u0440\u0433\u043e\u0434\u0430\u0440"},{"id_gbr":77,"region":"\u041a\u0438\u0435\u0432","number":"77"},{"id_gbr":80,"region":"\u041f\u0430\u0432\u043b\u043e\u0433\u0440\u0430\u0434","number":"\u041f\u0430\u0432\u043b\u043e\u0433\u0440\u0430\u0434 1"},{"id_gbr":81,"region":"\u041f\u0430\u0432\u043b\u043e\u0433\u0440\u0430\u0434","number":"\u041f\u0430\u0432\u043b\u043e\u0433\u0440\u0430\u0434 2"},{"id_gbr":82,"region":"\u041a\u0430\u043c\u0435\u043d\u0441\u043a","number":"\u041a\u0430\u043c\u0435\u043d\u0441\u043a 3"},{"id_gbr":83,"region":"\u041a\u0430\u043c\u0435\u043d\u0441\u043a","number":"\u041a\u0430\u043c\u0435\u043d\u0441\u043a 4"},{"id_gbr":84,"region":"\u0414\u043d\u0435\u043f\u0440","number":"\u0414\u043d\u0435\u043f\u0440 7"},{"id_gbr":85,"region":"\u0417\u0430\u043f\u043e\u0440\u043e\u0436\u044c\u0435","number":"\u0411\u0430\u0439\u043a\u0430\u043b 8"},{"id_gbr":86,"region":"\u041b\u044c\u0432\u043e\u0432","number":"\u041b\u044c\u0432\u043e\u0432 5"},{"id_gbr":88,"region":"\u041a\u0438\u0435\u0432","number":"91"},{"id_gbr":89,"region":"\u041a\u0430\u043c\u0435\u043d\u0441\u043a","number":"\u041a\u0430\u043c\u0435\u043d\u0441\u043a 5"},{"id_gbr":90,"region":"\u041c\u0430\u0440\u0438\u0443\u043f\u043e\u043b\u044c","number":"\u041c\u0430\u0440\u0438\u0443\u043f\u043e\u043b\u044c-3"},{"id_gbr":91,"region":"\u041a\u0438\u0435\u0432","number":"\u041c\u041e\u0422\u041e - 1"}]}`)
		switch js_cmnd {
		case "start": //First start
			//	js_result = startGBR(js_iden, js_name, js_param, getSocketIndex(conn))
			json.Unmarshal(b, &gbrlistSout)
			//msg := []byte("Hello on Serverside!")
			err = conn.WriteMessage(websocket.TextMessage, b)
		case "login": //Loging for user
				fmt.Println("In login case")
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
		//userid - Token of device; js_name - token; js_param - password
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
func loginChecker(userid, js_name, js_param string, conPosition int) string {
	return ""
}

type StringTable []string

func (st StringTable) Get(i int) string {
	if i < 0 || i >= len(st) {
		return ""
	}
	return st[i]
}

//------------------------------------------------------------------------------
func logGBR(userid, js_name, js_param string, conPosition int) string {
	//myTable := StringTable{
	//	1: "71",
	//	2: "72",
	//	3: "73",
	//}
	s_sql := "-1"
	s_sql += "(CONSTVISIB = 0) AND (CONSTKIND = 4) AND (IDCONST = " + userid
	s_sql += ") LIMIT 1"
//TODO remake valid method
	gbrvalid := bla(userid)

	s_sql = "SELECT IDPERS,FIOPERS,PAROL FROM personality WHERE IDPERS=" + dbQuatedString(js_name)
	s_json := ""

	if gbrvalid==false && js_name!="-1"&& js_param!="-111"{
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
		s_psw := "-111"
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
			s_json = s_json + getQuatedJSON("param", "LOG_FALSE_x", 1) + string(0x0D) + string(0x0A)
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
