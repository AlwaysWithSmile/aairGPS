package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	//	"math/rand"
	"strconv"

	"github.com/gorilla/websocket"
)


type gbrNowActiveWorkers struct{
	Id_workings int `json:"id_workings"`
	ObjectNumberPult string `json:"f_object_number_pult"`
	ObjectAdress string `json:"f_object_adress"`
	ObjectName string `json:"f_object_name"`
	Region string `json:"f_region"`
	GbrNumber string `json:"f_gbr_number"`
	GbrNumberRezerv string `json:"f_gbr_number_rezerv"`
	IdGBR string `json:"id_gbr"`
}

type sendStatusOfAlarm struct{
	Status string `json:"status"`
	Param string `json:"param"`
	Id string `json:"id"`
}

type CardBase struct {
	ID                       int    `json:"idd"`
	CARD_TYPE                int    `json:"type_object_cart"`
	CARD_AFFILATION          int    `json:"affiliation"`
	CARD_INSTALLER           int    `json:"installer"`
	CARD_CLIENT              string `json:"field_client"`
	CARD_PULTNUM             string `json:"pult"`
	CARD_RADIO_CHANEL        string `json:"field_radio_chanel"`
	CARD_RADIO_CHANEL_RESERV string `json:"field_radio_chanel_rezerv"`
	CARD_REGION              string `json:"field_region"`
	CARD_PEREZVON            string `json:"perezvon"`
	CARD_GBR_ACTION          int    `json:"gbr_action"`
	CARD_CALL                string `json:"field_call"`
	CARD_CALL_RESERV         string `json:"field_call_rezerv"`
	CARD_CALL_RESERV2        string `json:"field_call_rezerv2"`
	CARD_TIME_RESPONSE       string `json:"field_time_response"`
	CARD_CONTROL_PANEL       string `json:"field_contol_panel"`
	CARD_COMPANY             int    `json:"field_company"`
	CARD_ALIENT_PULT         int    `json:"field_alien_pult"`
	CARD_NAME                string `json:"obname"`
	CARD_ADRES               string `json:"obadr"`
	CARD_MODE                string `json:"field_mode"`
	CARD_TYPE_OBJECT         string `json:"field_type_object"`
	CARD_EXTRACT_ADDRESS     string `json:"exact_address"`
	CARD_STOREYS             string `json:"storeys"`
	CARD_FLOOR               string `json:"floor"`
	CARD_KEY_PRESENT         string `json:"key_availability"`
	CARD_HAVE_DOG            string `json:"having_dog"`
	CARD_OUT_INTO            string `json:"build_out_or_into"`
	CARD_WINDOW_DOOR         string `json:"window_and_dor"`
	CARD_SECURITY            string `json:"security_in_object"`
	CARD_WAYMARK             string `json:"waymark"`
	CARD_PORCH               string `json:"field_porch"`
	CARD_VULNER              string `json:"field_vulnerabilities"`
	CARD_INFO2               string `json:"field_description_2"`
	CARD_EQUIP               string `json:"field_equipment"`
	CARD_WHOSE_EQUIP         string `json:"field_whose_equipment"`
	CARD_AUTHOR              string `json:"field_author"`
	CARD_MANAGER             string `json:"field_manager"`
	CARD_DOGOVOR             string `json:"field_dogovor"`
	CARD_SUM_MONTH           string `json:"field_summ_in_month"`
	CARD_PEOPLE              string `json:"field_people"`
	CARD_SHLEYF              string `json:"field_shleif"`

	CARD_DATE_ENTER string `json:"field_date_enter_object"`
	CARD_START_SEC  string `json:"field_date_start_security"`
	CARD_WARNING    string `json:"field_warning"`
	CARD_LAT        string `json:"lat"`
	CARD_LON        string `json:"lon"`


	CARD_FILES []string `json:"files"`
}

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

type People struct {
	MAN_NUM    string `json:"number_people_line"`
	MAN_NAME   string `json:"field_people_name"`
	MAN_ADDR   string `json:"field_adress"`
	MAN_PHONE  string `json:"field_phone"`
	MAN_ACCESS string `json:"field_access"`
	MAN_NOTE   string `json:"field_note"`

	Peoples []string `json:"users"`
}

type Zone struct {
	ZONE_NUM string `json:"number_shleif_line"`
	ZONE_NAME string `json:"field_shleif_name"`
	ZONE_PLACE string `json:"field_shleif_place"`

	Zoness []string `json:"zones"`
}
//========================= get JSON func =======================================
type StringTable []string

func(st StringTable) Get(i int) string{
	if i < 0 || i > len(st){
		return ""
	}
	return st[i]
}

func(st StringTable) GetIndex(i int) int{
	return i
}

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
func sleepinGoopher(jsonData []byte, a gbrNowActiveWorkers, b gbrNowActiveWorkers, conn *websocket.Conn){
	cardBase := new(CardBase)
	people := new(People)
	zone := new(Zone)



	getJSON(`{
   "number_people_line":"1",
   "field_people_name":"\\u0422\\u0435\\u043b\\u0435\\u0444\\u043e\\u043d \\u0434\\u0438\\u0441\\u043f\\u0435\\u0442\\u0447\\u0435\\u0440\\u0430 ",
   "field_adress":"",
   "field_phone":"207-69-02, 296-79-50, 296-79-90",
   "field_access":"",
   "field_note":""
}`, &people)

	getJSON(`{
	"number_shleif_line":"1",
   "field_shleif_name":"\\u0420\\u0423-0,4 (\\u0421\\u0420\\u041f-600)",
   "field_shleif_place":"\\u0420\\u0423-0,4 (\\u0421\\u0420\\u041f-600)"
}`, &zone)
	getJSON(`
{
	"idd":"7761",
	"lat":"50.4558307",
	"lon":"30.6366454",
	"obadr":"г.Киев, улица Красноткацкая, 40",
	"obname":"Трансформаторная подстанция ",
	"obtel":"",
	"pult":"123",
	"status":"",
	"waymark":"\u043e\u0431\u044a\u0435\u043a\u0442 \u043d\u0430\u0445\u043e\u0434\u0438\u0442\u0441\u044f \u043f\u043e \u0430\u0434\u0440\u0435\u0441\u0443 \u0443\u043b\u0438\u0446\u0430 \u041a\u0440\u0430\u0441\u043d\u043e\u0442\u043a\u0430\u0446\u043a\u0430\u044f, 40 \u043d\u0430 \u0442\u0435\u0440\u0440\u0438\u0442\u043e\u0440\u0438\u0438 \u041f\u041f \u041c\u0438\u043a\u0441\u0435\u0440. \u041f\u0435\u0440\u0435\u0434 \u0444\u0438\u0440\u043c\u0435\u043d\u043d\u044b\u043c \u043c\u0430\u0433\u0430\u0437\u0438\u043d\u043e\u043c \u041c\u0438\u043a\u0443\u043b\u0438\u043d\u0435\u0446\u043a\u0438\u0439 \u0431\u0440\u043e\u0432\u0430\u0440 \u043f\u043e\u0432\u0435\u0440\u043d\u0443\u0442\u044c \u043d\u0430 \u043f\u0440\u0430\u0432\u043e \u0438 \u043f\u0440\u043e\u0435\u0445\u0430\u0442\u044c 20\u043c\u0435\u0442\u0440\u043e\u0432. \u041e\u0434\u043d\u043e\u044d\u0442\u0430\u0436\u043d\u043e\u0435 \u0442\u0435\u0445\u043d\u0438\u0447\u0435\u0441\u043a\u043e\u0435 \u043a\u0430\u043f\u0438\u0442\u0430\u043b\u044c\u043d\u043e\u0435 \u0437\u0434\u0430\u043d\u0438\u0435"
}
`, &cardBase)



	time.Sleep(3 * time.Second)
	fmt.Println("Snore....")
	for  {
		if err := json.Unmarshal(jsonData, &a); err != nil{
			panic(err)
		}
		fmt.Println("gbr id before IF circle: ", b)
		if a != b {
			b = a
			fmt.Println("in IF case B: ", b)

			convertedCardBaseID := strconv.Itoa(cardBase.ID)
			s2 := string(34)
			s_json := "{" + s2 + "obinfo" + s2 + ":" + "[" + "{" + s2 + "id" + s2 + ":" + s2 + convertedCardBaseID + s2 + "," + s2 +"lat" + s2 +":"+ s2 + cardBase.CARD_LAT + s2 +"," + s2 + "lon"+ s2 +":" + s2 + cardBase.CARD_LON + s2 +"," + s2 +"obadr"+ s2 +":" + s2 +cardBase.CARD_ADRES+ s2 +","+ s2 +"obname"+ s2 +":" + s2 +cardBase.CARD_NAME+ s2 +"," + s2 + "pult" + s2 + ":"+ s2 + cardBase.CARD_PULTNUM+ s2 + "," + s2 + "details" + s2 + ":" + s2 + cardBase.CARD_WAYMARK + s2 +"}" +"]"
			s_json += ","
			s_json += s2 + "userlist" + s2 + ":" + "[" + "{" + s2 + "name" + s2 + ":"+ s2 + people.MAN_NAME + s2 + "," + "num" + ":" + s2 + people.MAN_NUM + s2 + "," + s2 + "tel" + s2 + ":" + s2 + people.MAN_PHONE + s2 + "}"+"]"
			s_json += ","
			s_json += s2 + "zonelist" + s2 + ":"+"["+"{" + s2 + "name" + s2 + ":" + s2 + zone.ZONE_NAME + s2 + "," + s2 + "num" + s2 + ":" + s2 + zone.ZONE_NUM + s2 + "," + s2 + "tel" + s2 + ":" + s2 + zone.ZONE_PLACE + s2 + "}"+"]"
			s_json +=","
			s_json += s2 + "eventlist" + s2 + ":" + "[" +"{"+"}"+ "]"
			s_json += ","
			s_json += s2 + "imagelist" + s2 + ":" + "[" + "{" + s2 + "url" + s2 + ":"+ s2 + "https://cs.ohholding.com.ua/view/object_cart/uploads/7761/Screenshot_4.jpg" + s2 + "}" + "]"
			s_json += "}"
			fmt.Println(s_json)
			s_jsonbyte := []byte(s_json)
			conn.WriteMessage(websocket.TextMessage, s_jsonbyte)
		}else{
			fmt.Println("There are not any new alerts....")
		/*	message :=[]byte( "There are not new alerts" )
			conn.WriteMessage(websocket.TextMessage, message)*/
		}
		fmt.Println("After IF circle B: ",b)
		fmt.Println("//==============================================")
		time.Sleep( 1 * time.Second)
	}
}

func alarmbreak(conn *websocket.Conn) {
	postTestData := []byte(`{
		"status":"alarmbreak",
		"param":"X0Y0",
		"id":"123456"
	}`)
	var testjsonUnmarshal sendStatusOfAlarm
	if err := json.Unmarshal(postTestData, testjsonUnmarshal); err != nil {
		panic(err)
	}
	r := bytes.NewReader(postTestData)
	resp, err := http.Post("http://api-cs.ohholding.com.ua/api/set-status?status=" + testjsonUnmarshal.Status+"&param="+testjsonUnmarshal.Param+"&id="+testjsonUnmarshal.Id, "application/json", r);
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v %v", err, resp)
	message := []byte( "{\"cmnd\":\"connect\",\"id\":\"8\",\"name\":\"-1\",\"param\":\"Alarmbreak\"}" )
	conn.WriteMessage(websocket.TextMessage, message)
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

		fmt.Println(js_name)

		getJSON("http://api-cs.ohholding.com.ua/gbr_list/get", &gbrlistSout)
		b := []byte(`{"gbrlist":[{"id_gbr":1,"region":"\u041a\u0438\u0435\u0432","number":"71"},{"id_gbr":2,"region":"\u041a\u0438\u0435\u0432","number":"72"},{"id_gbr":3,"region":"\u041a\u0438\u0435\u0432","number":"73"},{"id_gbr":4,"region":"\u041a\u0438\u0435\u0432","number":"74"},{"id_gbr":5,"region":"\u041a\u0438\u0435\u0432","number":"75"},{"id_gbr":6,"region":"\u041a\u0438\u0435\u0432","number":"78"},{"id_gbr":7,"region":"\u041a\u0438\u0435\u0432","number":"79"},{"id_gbr":8,"region":"\u041a\u0438\u0435\u0432","number":"80"},{"id_gbr":9,"region":"\u041a\u0438\u0435\u0432","number":"81"},{"id_gbr":10,"region":"\u041a\u0438\u0435\u0432","number":"82"},{"id_gbr":11,"region":"\u041a\u0438\u0435\u0432","number":"83"},{"id_gbr":12,"region":"\u041a\u0438\u0435\u0432","number":"84"},{"id_gbr":13,"region":"\u041a\u0438\u0435\u0432","number":"85"},{"id_gbr":14,"region":"\u041a\u0438\u0435\u0432","number":"86"},{"id_gbr":15,"region":"\u041a\u0438\u0435\u0432","number":"88"},{"id_gbr":16,"region":"\u041a\u0438\u0435\u0432","number":"89"},{"id_gbr":17,"region":"\u041a\u0438\u0435\u0432","number":"92"},{"id_gbr":18,"region":"\u0417\u0430\u043f\u043e\u0440\u043e\u0436\u044c\u0435","number":"\u0411\u0430\u0439\u043a\u0430\u043b 1"},{"id_gbr":19,"region":"\u0417\u0430\u043f\u043e\u0440\u043e\u0436\u044c\u0435","number":"\u0411\u0430\u0439\u043a\u0430\u043b 2"},{"id_gbr":20,"region":"\u0417\u0430\u043f\u043e\u0440\u043e\u0436\u044c\u0435","number":"\u0411\u0430\u0439\u043a\u0430\u043b 3"},{"id_gbr":21,"region":"\u0417\u0430\u043f\u043e\u0440\u043e\u0436\u044c\u0435","number":"\u0411\u0430\u0439\u043a\u0430\u043b 4"},{"id_gbr":22,"region":"\u0417\u0430\u043f\u043e\u0440\u043e\u0436\u044c\u0435","number":"\u0411\u0430\u0439\u043a\u0430\u043b 5"},{"id_gbr":23,"region":"\u0417\u0430\u043f\u043e\u0440\u043e\u0436\u044c\u0435","number":"\u0411\u0430\u0439\u043a\u0430\u043b 6"},{"id_gbr":24,"region":"\u0417\u0430\u043f\u043e\u0440\u043e\u0436\u044c\u0435","number":"\u0411\u0430\u0439\u043a\u0430\u043b 7"},{"id_gbr":25,"region":"\u0414\u043d\u0435\u043f\u0440","number":"\u0414\u043d\u0435\u043f\u0440 1"},{"id_gbr":26,"region":"\u0414\u043d\u0435\u043f\u0440","number":"\u0414\u043d\u0435\u043f\u0440 2"},{"id_gbr":27,"region":"\u0414\u043d\u0435\u043f\u0440","number":"\u0414\u043d\u0435\u043f\u0440 3"},{"id_gbr":28,"region":"\u0414\u043d\u0435\u043f\u0440","number":"\u0414\u043d\u0435\u043f\u0440 4"},{"id_gbr":29,"region":"\u0414\u043d\u0435\u043f\u0440","number":"\u0414\u043d\u0435\u043f\u0440 5"},{"id_gbr":30,"region":"\u0414\u043d\u0435\u043f\u0440","number":"\u0414\u043d\u0435\u043f\u0440 6"},{"id_gbr":31,"region":"\u041a\u0430\u043c\u0435\u043d\u0441\u043a","number":"\u041a\u0430\u043c\u0435\u043d\u0441\u043a 1"},{"id_gbr":32,"region":"\u041a\u0430\u043c\u0435\u043d\u0441\u043a","number":"\u041a\u0430\u043c\u0435\u043d\u0441\u043a 2"},{"id_gbr":35,"region":"\u041a\u0440\u0438\u0432\u043e\u0439 \u0420\u043e\u0433","number":"\u041a\u0440\u0438\u0432\u0431\u0430\u0441 1"},{"id_gbr":36,"region":"\u041a\u0440\u0438\u0432\u043e\u0439 \u0420\u043e\u0433","number":"\u041a\u0440\u0438\u0432\u0431\u0430\u0441 2\r\n"},{"id_gbr":37,"region":"\u041a\u0440\u0438\u0432\u043e\u0439 \u0420\u043e\u0433","number":"\u041a\u0440\u0438\u0432\u0431\u0430\u0441 3"},{"id_gbr":38,"region":"\u041a\u0440\u0438\u0432\u043e\u0439 \u0420\u043e\u0433","number":"\u041a\u0440\u0438\u0432\u0431\u0430\u0441 4"},{"id_gbr":39,"region":"\u041a\u0440\u0438\u0432\u043e\u0439 \u0420\u043e\u0433","number":"\u041a\u0440\u0438\u0432\u0431\u0430\u0441 7"},{"id_gbr":40,"region":"\u041a\u0440\u0438\u0432\u043e\u0439 \u0420\u043e\u0433","number":"\u041a\u0440\u0438\u0432\u0431\u0430\u0441 6"},{"id_gbr":43,"region":"\u0414\u043e\u0431\u0440\u043e\u043f\u043e\u043b\u044c\u0435","number":"\u0421\u043e\u043a\u043e\u043b"},{"id_gbr":49,"region":"\u041b\u044c\u0432\u043e\u0432","number":"\u041b\u044c\u0432\u043e\u0432 1"},{"id_gbr":50,"region":"\u041b\u044c\u0432\u043e\u0432","number":"\u041b\u044c\u0432\u043e\u0432 2"},{"id_gbr":51,"region":"\u041b\u044c\u0432\u043e\u0432","number":"\u041b\u044c\u0432\u043e\u0432 3"},{"id_gbr":52,"region":"\u041b\u044c\u0432\u043e\u0432","number":"\u041b\u044c\u0432\u043e\u0432 4"},{"id_gbr":60,"region":"\u041a\u0438\u0435\u0432","number":"76"},{"id_gbr":65,"region":"\u041c\u0430\u0440\u0438\u0443\u043f\u043e\u043b\u044c","number":"\u041c\u0430\u0440\u0438\u0443\u043f\u043e\u043b\u044c-1"},{"id_gbr":66,"region":"\u041c\u0430\u0440\u0438\u0443\u043f\u043e\u043b\u044c","number":"\u041c\u0430\u0440\u0438\u0443\u043f\u043e\u043b\u044c-2"},{"id_gbr":67,"region":"\u041c\u0430\u0440\u0438\u0443\u043f\u043e\u043b\u044c","number":"\u041c\u0430\u0440\u0438\u0443\u043f\u043e\u043b\u044c-4"},{"id_gbr":72,"region":"\u041a\u0438\u0435\u0432","number":"87"},{"id_gbr":74,"region":"\u041f\u043e\u043a\u0440\u043e\u0432\u0441\u043a\u043e\u0435","number":"\u041f\u043e\u043a\u0440\u043e\u0432\u0441\u043a"},{"id_gbr":76,"region":"\u042d\u043d\u0435\u0440\u0433\u043e\u0434\u0430\u0440","number":"\u042d\u043d\u0435\u0440\u0433\u043e\u0434\u0430\u0440"},{"id_gbr":77,"region":"\u041a\u0438\u0435\u0432","number":"77"},{"id_gbr":80,"region":"\u041f\u0430\u0432\u043b\u043e\u0433\u0440\u0430\u0434","number":"\u041f\u0430\u0432\u043b\u043e\u0433\u0440\u0430\u0434 1"},{"id_gbr":81,"region":"\u041f\u0430\u0432\u043b\u043e\u0433\u0440\u0430\u0434","number":"\u041f\u0430\u0432\u043b\u043e\u0433\u0440\u0430\u0434 2"},{"id_gbr":82,"region":"\u041a\u0430\u043c\u0435\u043d\u0441\u043a","number":"\u041a\u0430\u043c\u0435\u043d\u0441\u043a 3"},{"id_gbr":83,"region":"\u041a\u0430\u043c\u0435\u043d\u0441\u043a","number":"\u041a\u0430\u043c\u0435\u043d\u0441\u043a 4"},{"id_gbr":84,"region":"\u0414\u043d\u0435\u043f\u0440","number":"\u0414\u043d\u0435\u043f\u0440 7"},{"id_gbr":85,"region":"\u0417\u0430\u043f\u043e\u0440\u043e\u0436\u044c\u0435","number":"\u0411\u0430\u0439\u043a\u0430\u043b 8"},{"id_gbr":86,"region":"\u041b\u044c\u0432\u043e\u0432","number":"\u041b\u044c\u0432\u043e\u0432 5"},{"id_gbr":88,"region":"\u041a\u0438\u0435\u0432","number":"91"},{"id_gbr":89,"region":"\u041a\u0430\u043c\u0435\u043d\u0441\u043a","number":"\u041a\u0430\u043c\u0435\u043d\u0441\u043a 5"},{"id_gbr":90,"region":"\u041c\u0430\u0440\u0438\u0443\u043f\u043e\u043b\u044c","number":"\u041c\u0430\u0440\u0438\u0443\u043f\u043e\u043b\u044c-3"},{"id_gbr":91,"region":"\u041a\u0438\u0435\u0432","number":"\u041c\u041e\u0422\u041e - 1"}]}`)
		switch js_cmnd {
		case "start": //First start
			//	js_result = startGBR(js_iden, js_name, js_param, getSocketIndex(conn))
			json.Unmarshal(b, &gbrlistSout)
			err = conn.WriteMessage(websocket.TextMessage, b)
		case "login": //Loging for user
			js_result = logGBR(js_iden, js_name, js_param, getSocketIndex(conn))
		case "connect":
			fmt.Println("In connect case")
			message := []byte( "{\"cmnd\":\"connect\",\"id\":\"8\",\"name\":\"-1\",\"param\":\"Connect_OK\"}" )
			err = conn.WriteMessage(websocket.TextMessage, message)
			fmt.Println("Successfully connected....")
			//TODO comparation json files
			//	case "alarmlist": //Get alarm list
	//		js_result = getAlarms(js_iden, js_name, js_param)
		case "alarmget": //Receive alarm
			jsonData := []byte(`
{
	"id_workings":245115,
	"f_object_number_pult":"89",
	"f_object_adress":"\u0433. \u041a\u0438\u0435\u0432, \u0443\u043b. \u041c\u0438\u0440\u043e\u043f\u043e\u043b\u044c\u0441\u043a\u0430\u044f, 1",
	"f_object_name":"\u0422\u041f 2594",
	"f_region":"\u041a\u0438\u0435\u0432",
	"f_gbr_number":"80",
	"f_gbr_number_rezerv":"",
	"id_gbr":"8"
}
`)
			var newGbrActiveWorker gbrNowActiveWorkers
			var activeGbrWorker gbrNowActiveWorkers
			go sleepinGoopher(jsonData, newGbrActiveWorker, activeGbrWorker, conn)

			fmt.Println("/==================================================")
			testCheck := []byte (jsonIncoming)
			if err := json.Unmarshal(testCheck, &airDecoding); err != nil {
				panic(err)
			}
			fmt.Println("AirDecoding: ", airDecoding.ID + airDecoding.Param)
			fmt.Println("/==================================================")
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
var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}
//func handler(w http.ResponseWriter, r *http.Request){
// if  conn, err := upgrader.Upgrade(w,r,nil); err != nil {
//	 panic(err)
//	 }
//}
func logGBR(userid, js_name, js_param string, conPosition int) string {

	s_sql := "-2"
	s_sql += "(CONSTVISIB = 0) AND (CONSTKIND = 4) AND (IDCONST = " + userid
	s_sql += ") LIMIT 1"
//TODO remake valid method
	gbrvalid := bla(userid)

	s_sql = "SELECT IDPERS,FIOPERS,PAROL FROM personality WHERE IDPERS=" + dbQuatedString(js_name)
	s_json := ""

	if gbrvalid==false || js_name!="-2" || js_param!="-111"{
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

//TODO testing getAlarmListMethod

/*func getAlarmListTest(objectid string, userid string) string{

/*	myTable := StringTable{
		8:"78",
	}
	convertedobjectId, _ := strconv.Atoi(objectid)
	converteduserId, _ := strconv.Atoi(userid)
	if convertedobjectId != myTable.GetIndex(converteduserId) && converteduserId != convertedobjectId {
		return ""
	}else{
	sout_json += "id_gbr" + nowActiveWorkers.IdGBR + "," + "f_gbr_number" + nowActiveWorkers.GbrNumber
	}
	return sout_json
}*/
//------------------------------------------------------------------------------

func getAlarms(userid, js_name, js_param string) string {
	s_json := "{" + string(0x0D) + string(0x0A)
	s_alarms := getALARMlist("8", userid)
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
func recAlarms(userid, js_cmnd, js_name, js_param string)string {
/*	//procPosition(userid, js_cmnd, js_name)
	//updateGBRstatus(userid, getGBRuser(userid), "", js_param, 1)

	jsonData := []byte(`{
	"id_workings":245115,
	"f_object_number_pult":"89",
	"f_object_adress":"\u0433. \u041a\u0438\u0435\u0432, \u0443\u043b. \u041c\u0438\u0440\u043e\u043f\u043e\u043b\u044c\u0441\u043a\u0430\u044f, 1",
	"f_object_name":"\u0422\u041f 2594",
	"f_region":"\u041a\u0438\u0435\u0432",
	"f_gbr_number":"80",
	"f_gbr_number_rezerv":"",
	"id_gbr":"8"
}`)
	var nowActiveWorkers gbrNowActiveWorkers
	if err := json.Unmarshal(jsonData, &nowActiveWorkers); err != nil{
		panic(err)
	}

	var n map[string]gbrNowActiveWorkers
	n = make(map[string]gbrNowActiveWorkers)
	n[nowActiveWorkers.IdGBR] = gbrNowActiveWorkers{nowActiveWorkers.Id_workings,
		nowActiveWorkers.IdGBR, nowActiveWorkers.GbrNumber,
		nowActiveWorkers.GbrNumberRezerv, nowActiveWorkers.ObjectAdress,
		nowActiveWorkers.ObjectName,
		nowActiveWorkers.GbrNumberRezerv, nowActiveWorkers.ObjectNumberPult}

	jsonStr, err := json.Marshal(n[userid])
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
		return(string(jsonStr))*/
	return ""
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
