package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	//"log"

	"math"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//===================== Delphi convert date functions ==========================
// 4x faster than dateutils version
func UNIXTimeToDateTimeFAST(intUnixTime int64) float64 {
	//2021-06-08 16:34:19.3509522 +0300 EEST
	s := time.Now().UTC().Local().String()
	i_pos := strings.Index(s, "+")
	sec_local := 3600 * convertInt64Val(s[i_pos+1:i_pos+3])
	return float64(intUnixTime+sec_local)/86400 + 25569
}

//------------------------------------------------------------------------------
// 10x faster than dateutils version
func DateTimeToUNIXTimeFAST(DelphiTime float64) int64 {
	return int64(math.Round((DelphiTime - 25569) * 86400))
}

//------------------------------------------------------------------------------
func floatToStr64(f_data float64) string {
	return strconv.FormatFloat(f_data, 'f', -1, 64)
}

//------------------------------------------------------------------------------
func delphiDateToSQL(unixConvDT int64) string {
	f_date := UNIXTimeToDateTimeFAST(unixConvDT)
	s_date := floatToStr64(f_date)
	return doReplaceStr(s_date, ",", ".")
}

//======================== General utils ==================================
func recoveryAppFunction() {
	if recoveryMessage := recover(); recoveryMessage != nil {
		fmt.Println(recoveryMessage)
	}
	fmt.Println(getDT(), "Application restored after Error...")
}

//-----------------------------------------------------------------------------
func getDT() string {
	currentTime := time.Now()
	return currentTime.Format("2006.01.02 15:04:05") + " > "
	//return currentTime.String() + " > "
}

//-------------------------------------------------------------------------
func getQuatedJSON(datJSON string, valJSON string, tpJSON int) string {
	var rez_json string
	rez_json = ""
	rez_json = string(34) + datJSON + string(34) + ": "
	if tpJSON == 1 {
		rez_json = rez_json + string(34) + valJSON + string(34)
	} else {
		rez_json = rez_json + valJSON
	}
	return rez_json
}

//------------------------------------------------------------------------
func convertIntVal(strColumn string) int {
	var i_convert int
	i_convert = 0
	j, err := strconv.Atoi(strColumn)
	if err != nil {
		defer recoveryAppFunction()
		fmt.Println(getDT(), err)
	} else {
		i_convert = j
	}
	return i_convert
}

//--------------------------------------------------------------------------
func convertInt64Val(strColumn string) int64 {
	var i_convert int64
	i_convert = 0
	j, err := strconv.ParseInt(strColumn, 10, 64)
	if err != nil {
		fmt.Println(getDT(), err)
		defer recoveryAppFunction()
	}
	i_convert = j
	return i_convert
}

//--------------------------------------------------------------------------
func dec2hex(dec int) string {
	//hexStr := dec * 255 / 100
	//return fmt.Sprintf("%02x", hexStr)
	hexStr := fmt.Sprintf("%02x", dec)
	return strings.ToUpper(hexStr)

}

//--------------------------------------------------------------------------
func hex2dec(hex string) int64 {
	dec, err := strconv.ParseInt("0x"+hex, 0, 16)
	if err != nil {
		fmt.Println(getDT(), err)
		defer recoveryAppFunction()
	}
	dec = dec * 100 / 255
	return dec

}

//---------------- convert -----------------------
func hex2int(hexStr string) int {
	result, err := strconv.ParseUint(hexStr, 16, 32)
	if err != nil {
		return -1
	}
	return int(result)
}

//--------------------------------------------------------------------------
func string_int32(snumeric string, defnimeric int64) int64 {
	var i_result int64
	i_result = defnimeric
	i, err := strconv.ParseInt(snumeric, 10, 64)
	if err != nil {
		defer recoveryAppFunction()
		fmt.Println(getDT(), "Error Convertion type:"+snumeric, err)
		//panic(err)
	}
	i_result = i
	return i_result
}

//------------------------------------------------------------------------------
func strToFloat64(strFloat string) float64 {
	value, err := strconv.ParseFloat(strFloat, 64)
	if err != nil {
		defer recoveryAppFunction()
		return 0
	}
	return float64(value)
}

//=================== String functions =====================================//
func doReplaceStr(whole_String, old_Str, new_Str string) string {
	var s_result string
	s_result = whole_String
	for ok := true; ok; ok = (strings.Index(s_result, old_Str) > 0) {
		s_result = strings.Replace(s_result, old_Str, new_Str, -1)
	}
	return s_result
}

//---------------------------------------------------------------------------
func dbQuatedString(queryString string) string {
	return string(39) + queryString + string(39)
}

//---------------------------------------------------------------------------
func jsonQuatedString(queryString string) string {
	return string(34) + queryString + string(34)
}

//======================= JSON FUNCTIONS =====================================
func checkValidJson(strChecked string) bool {
	var is_valid bool
	is_valid = true
	if len(strChecked) < 16 {
		return false
	}
	if strings.Index(strChecked, "{") < 0 {
		return false
	}
	if strings.Index(strChecked, "id") < 1 {
		return false
	}

	if strings.Index(strChecked, "cmnd") < 1 {
		return false
	}
	if strings.Index(strChecked, "name") < 1 {
		return false
	}
	if strings.Index(strChecked, "param") < 1 {
		return false
	}
	if strings.Index(strChecked, "}") < 1 {
		return false
	}
	return is_valid
}

//===================== GET DATABASE TYPES =====================================
func dbUpdateData(uptSQL string) {
	db, err := sqlx.Connect("mysql", connect_db)

	if err != nil {
		defer recoveryAppFunction()
		fmt.Println(getDT(), "Error connect:"+connect_db, err)
		panic(err)

	}
	defer db.Close()
	result, err := db.Exec(uptSQL)
	if err != nil {
		defer recoveryAppFunction()
		fmt.Println(getDT(), "Error update:"+uptSQL, err)
		panic(err)
	}
	if result == nil {
		fmt.Println(getDT(), "Update result:", uptSQL, result)
		//	fmt.Println(result.LastInsertId()) // id последнего удаленого объекта
		//	fmt.Println(result.RowsAffected()) // количество затронутых строк
	}

}

//---------------------- Read int value from column 1 --------------------------
func dbGetIntData(queryString string, rownum int) int {
	var uin_result int
	uin_result = -1

	db, err := sqlx.Connect("mysql", connect_db)

	if err != nil {
		defer recoveryAppFunction()
		fmt.Println(getDT(), "Error read value :"+queryString, err)
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query(queryString)
	if err != nil {
		defer recoveryAppFunction()
		fmt.Println("Error read value :" + queryString)
	}
	cols, err := rows.Columns()
	if err != nil {
		defer recoveryAppFunction()
		fmt.Println("Error read value :" + queryString)
	}
	data := make(map[string]string)

	if rows.Next() {
		columns := make([]string, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}
		rows.Scan(columnPointers...)
		for i, colName := range cols {
			data[colName] = columns[i]
			if i == rownum {
				j, err := strconv.Atoi(columns[i])
				if err != nil {
					defer recoveryAppFunction()
					fmt.Println(getDT(), "Error convert value :"+columns[i], err)
					panic(err)
				} else {
					uin_result = j
				}
			}
		}
	}
	if rows != nil {
		defer rows.Close()
	}
	return (uin_result)
}

//---------------------- Read string value from column 1 -----------------------
func dbGetStringData(queryString string, col_poz int) string {
	var str_result string
	str_result = ""

	db, err := sqlx.Connect("mysql", connect_db)

	if err != nil {
		defer recoveryAppFunction()
		fmt.Println("Error read value :"+queryString, err)
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query(queryString)
	if err != nil {
		defer recoveryAppFunction()
		fmt.Println("Error read value :" + queryString)
	}
	cols, err := rows.Columns()
	if err != nil {
		defer recoveryAppFunction()
		fmt.Println("Error read value :" + queryString)
	}
	data := make(map[string]string)

	if rows.Next() {
		columns := make([]string, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}
		rows.Scan(columnPointers...)
		for i, colName := range cols {
			data[colName] = columns[i]
			if i == col_poz {
				str_result = columns[i]
			}
		}
	}
	if rows != nil {
		defer rows.Close()
	}
	return (str_result)
}

//===================== DATABASE FUNCTIONS ===================================
func getGBRlist(isGBRlist int) string {
	s_query := ""

	if isGBRlist == 0 {
		s_query = "SELECT IDCONST, CONSTVALUE FROM consttable "
		s_query += "WHERE (CONSTVISIB = 0) AND (CONSTKIND = 5) ORDER BY CONSTVALUE"
	} else if isGBRlist == 1 {
		s_query = "SELECT IDCONST, CONSTVALUE FROM consttable "
		s_query += "WHERE (CONSTVISIB = 0) AND (CONSTKIND = 4) ORDER BY CONSTVALUE"
	} else {
		s_query = "SELECT IDPERS, FIOPERS, PERSKIND FROM personality "
		s_query += "WHERE PERSSTATUS=5 ORDER BY FIOPERS"
	}

	db, err := sqlx.Connect("mysql", connect_db)
	s_json := ""
	if err != nil {
		defer recoveryAppFunction()
		fmt.Println(getDT(), "Error read value :"+s_query, err)
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query(s_query)
	if err != nil {
		defer recoveryAppFunction()
		fmt.Println("Error read value :" + s_query)
	}
	cols, err := rows.Columns()
	if err != nil {
		defer recoveryAppFunction()
		fmt.Println("Error read value :" + s_query)
	}
	data := make(map[string]string)
	for rows.Next() {
		columns := make([]string, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}
		rows.Scan(columnPointers...)

		for i, colName := range cols {
			data[colName] = columns[i]
			if i == 0 { // 0 - Get Id
				if len(s_json) > 5 {
					s_json += string(0x0D) + string(0x0A)
				}
				s_json += "{" + string(0x0D) + string(0x0A)
				s_json += getQuatedJSON("id", columns[i], 1) + "," + string(0x0D) + string(0x0A)
			}
			if i == 1 { // 1 - Get Name
				s_json += getQuatedJSON("name", strings.ReplaceAll(columns[i], string(34), " "), 1) + string(0x0D) + string(0x0A)
				s_json += "},"
			}
		}
	}
	if rows != nil {
		defer rows.Close()
	}
	return s_json
}

//------------------------------------------------------------------------------
func checkGBRneed(uinOBJ, uinGBR string) bool {
	resultat := false
	s_sql := "SELECT IDOB,EKIPAJ FROM objectlist WHERE (IDOB="
	s_sql += uinOBJ + ") AND (EKIPAJ=" + uinGBR + ")"
	if dbGetIntData(s_sql, 0) > 0 {
		resultat = true
	} else {
		s_sql = "SELECT IDOB,NUM FROM objectlist WHERE IDOB=" + uinOBJ
		uinEQUIP := dbGetIntData(s_sql, 1)
		if uinEQUIP > 0 {
			s_sql = "SELECT IDGBR,GBREQID,GBRPERSID FROM gbrlist WHERE (GBREQID="
			s_sql += strconv.Itoa(uinEQUIP) + ") AND (GBRPERSID=" + uinGBR + ") LIMIT 1"
			if dbGetIntData(s_sql, 0) > 0 {
				resultat = true
			}
		}
	}
	return resultat
}


//------------------------------------------------------------------------------
func getALARMlist(objectIden, gbrUIN string) string {
	s_sql := ""
	is_Single := (len(objectIden) > 0)
	if is_Single == true {
		s_sql = "SELECT eventlist.IDEV,eventlist.EVDATA,eventlist.CODEID,eventlist.ISNEW,"
		s_sql += "eventlist.OBJECTID,objectlist.IDOB,objectlist.OBNAME,objectlist.OBADR,"
		s_sql += "objectlist.OBTEL,objectlist.DOLGOTA,objectlist.SHIROTA,objectlist.OHRANA,objectlist.PULTADR,"
		s_sql += "objectlist.AVATAR,codelist.IDCODE,codelist.EVENTS,eventlist.ZONENUM,eventlist.ISGBR "
		s_sql += "FROM objectlist INNER JOIN eventlist on (objectlist.IDOB = eventlist.OBJECTID) "
		s_sql += "INNER JOIN codelist on (eventlist.CODEID = codelist.IDCODE) WHERE (ISNEW>0) AND (ISFINISH=0) AND "
		s_sql += "(objectlist.IDOB=" + objectIden + ") ORDER BY objectlist.IDOB, eventlist.IDEV DESC"

		/**TODO Testing Layout*/
		/*alarmlist - разбить obtel на obowner & obphone*/

		//s_sql = "SELECT IDMAPA,MAPAOBJ,MAPAOPYS,MAPAKIND,MAPAWAY FROM mapalist"
		//s_sql +="WHERE MAPAOBJ=\" + objUIN + \" ORDER BY MAPAOPYS"
	} else {
		s_sql = "SELECT eventlist.IDEV,eventlist.EVDATA,eventlist.CODEID,eventlist.ISNEW,"
		s_sql += "eventlist.OBJECTID,objectlist.IDOB,objectlist.OBNAME,objectlist.OBADR,"
		s_sql += "objectlist.OBTEL,objectlist.DOLGOTA,objectlist.SHIROTA,objectlist.OHRANA,objectlist.PULTADR,"
		s_sql += "objectlist.AVATAR,codelist.IDCODE,codelist.EVENTS,eventlist.ZONENUM,eventlist.ISGBR "
		s_sql += "FROM objectlist INNER JOIN eventlist on (objectlist.IDOB = eventlist.OBJECTID) "
		s_sql += "INNER JOIN codelist on (eventlist.CODEID = codelist.IDCODE) WHERE (ISNEW>0) AND (ISFINISH=0) "
		s_sql += "ORDER BY objectlist.IDOB, eventlist.IDEV DESC"
	}

	db, err := sqlx.Connect("mysql", connect_db)
	s_json := ""
	if err != nil {
		defer recoveryAppFunction()
		fmt.Println(getDT(), "Error read value :"+s_sql, err)
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query(s_sql)
	if err != nil {
		defer recoveryAppFunction()
		fmt.Println("Error read value :" + s_sql)
	}
	cols, err := rows.Columns()
	if err != nil {
		defer recoveryAppFunction()
		fmt.Println("Error read value :" + s_sql)
	}
	data := make(map[string]string)
	for rows.Next() {
		columns := make([]string, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}
		rows.Scan(columnPointers...)
		s_pult := ""
		s_json_temp := ""
		b_json_need := false
		for i, colName := range cols {
			data[colName] = columns[i]
			if i == 0 { // 0 - Get Id

				s_json_temp += "{" + string(0x0D) + string(0x0A)
				s_json_temp += getQuatedJSON("id", columns[i], 1) + "," + string(0x0D) + string(0x0A)
			}
			if i == 1 { // 1 - Date
				f_date := strToFloat64(columns[i])
				i_date := DateTimeToUNIXTimeFAST(f_date)
				s_json_temp += getQuatedJSON("evdata", strconv.FormatInt(i_date, 10), 1) + "," + string(0x0D) + string(0x0A)
			}
			if i == 5 { // Object Id
				s_json_temp += getQuatedJSON("idob", columns[i], 1) + "," + string(0x0D) + string(0x0A)
				b_json_need = checkGBRneed(columns[i], gbrUIN)
				if len(s_json) > 5 && b_json_need {
					s_json += string(0x0D) + string(0x0A)
				}

			}

			if i == 6 && is_Single == false { // Object name
				s_json_temp += getQuatedJSON("obname", columns[i], 1) + "," + string(0x0D) + string(0x0A)
			}
			if i == 7 && is_Single == false { // Object addr
				s_json_temp += getQuatedJSON("obadr", columns[i], 1) + "," + string(0x0D) + string(0x0A)
			}
			if i == 8 && is_Single == false { // Object tel
				s_json_temp += getQuatedJSON("obowner", columns[i], 1) + "," + string(0x0D) + string(0x0A)
			}
			if i == 9 && is_Single == false { // Object longitude
				s_json_temp += getQuatedJSON("lon", columns[i], 1) + "," + string(0x0D) + string(0x0A)
			}
			if i == 10 && is_Single == false { // Object latitude
				s_json_temp += getQuatedJSON("lat", columns[i], 1) + "," + string(0x0D) + string(0x0A)
			}
			if i == 11 && is_Single == false { // Object status 0-unknown; 1 - open, 2-part closed; 3 - closed
				s_json_temp += getQuatedJSON("status", columns[i], 1) + "," + string(0x0D) + string(0x0A)
			}
			if i == 12 && is_Single == false { // Pult Addr
				s_pult = columns[i] + ".0"
			}
			if i == 13 && is_Single == false { // Pult Group
				s_json_temp += getQuatedJSON("num", s_pult+columns[i], 1) + "," + string(0x0D) + string(0x0A)
			}
			if i == 15 { // Event
				s_json_temp += getQuatedJSON("event", columns[i], 1) + "," + string(0x0D) + string(0x0A)
			}

			if i == 16 { // Zone
				if is_Single == false {
					s_json_temp += getQuatedJSON("zone", columns[i], 1) + "," + string(0x0D) + string(0x0A)
				} else {
					s_json_temp += getQuatedJSON("zone", columns[i], 1) + "},"
					if b_json_need {
						s_json += s_json_temp

					}
				}

			}

			if i == 17 && is_Single == false { // Is GBR comes, Unix time
				f_date := strToFloat64(columns[i])
				i_date := DateTimeToUNIXTimeFAST(f_date)
				s_json_temp += getQuatedJSON("gbr", strconv.FormatInt(i_date, 10), 1) + string(0x0D) + string(0x0A)
				s_json_temp += "},"
				if b_json_need {
					s_json += s_json_temp

				}
			}
		}
	}
	if rows != nil {
		defer rows.Close()
	}
	return s_json
}

//-------------------------------GET CARD INFO ---------------------------------
func getObjectStatus(objUIN string) (string, string) {
	s_sql := "SELECT IDOB,OHRANA,SVIZI FROM objectlist WHERE IDOB=" + objUIN + " LIMIT 1"
	db, err := sqlx.Connect("mysql", connect_db)
	s_con := "0"
	s_arm := "0"
	if err != nil {
		defer recoveryAppFunction()
		fmt.Println(getDT(), "Error read value :"+s_sql, err)
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(s_sql)
	if err != nil {
		defer recoveryAppFunction()
		fmt.Println("Error read value :" + s_sql)
	}
	cols, err := rows.Columns()
	if err != nil {
		defer recoveryAppFunction()
		fmt.Println("Error read value :" + s_sql)
	}
	data := make(map[string]string)
	if rows.Next() {
		columns := make([]string, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}
		rows.Scan(columnPointers...)

		for i, colName := range cols {
			data[colName] = columns[i]

			if i == 1 { // OHRANA
				s_arm = columns[i]
			}
			if i == 2 { // SVIZI
				s_con = columns[i]
			}
		}
	}

	if rows != nil {
		defer rows.Close()
	}
	return s_arm, s_con

}

//------------------------------------------------------------------------------
func getObjGeneral(objUIN string, isKey bool) string {
	s_sql := "SELECT IDOB,OBNAME,OBADR,OBTEL,OHRANA,SVIZI,MOREINFO,INFODETAIL,"
	s_sql += "concat(PULTADR,'.0',AVATAR) AS PULTNUMBER,DOLGOTA,SHIROTA "
	s_sql += "FROM objectlist WHERE IDOB=" + objUIN + " LIMIT 1"

	db, err := sqlx.Connect("mysql", connect_db)
	s_json := ""
	if isKey == false {
		s_json += getQuatedJSON("obinfo", "[", 0) + string(0x0D) + string(0x0A)
	}

	if err != nil {
		defer recoveryAppFunction()
		fmt.Println(getDT(), "Error read value :"+s_sql, err)
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query(s_sql)
	if err != nil {
		defer recoveryAppFunction()
		fmt.Println("Error read value :" + s_sql)
	}
	cols, err := rows.Columns()
	if err != nil {
		defer recoveryAppFunction()
		fmt.Println("Error read value :" + s_sql)
	}
	data := make(map[string]string)
	if rows.Next() {
		columns := make([]string, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}
		rows.Scan(columnPointers...)

		for i, colName := range cols {
			data[colName] = columns[i]
			if i == 0 { // 0 - Get IDOB
				s_json += "{" + string(0x0D) + string(0x0A)
				s_json += getQuatedJSON("id", columns[i], 1) + "," + string(0x0D) + string(0x0A)
			}
			if i == 1 { // 1 - OBNAME
				s_json += getQuatedJSON("obname", columns[i], 1) + "," + string(0x0D) + string(0x0A)
			}
			if i == 2 { // OBADR
				s_json += getQuatedJSON("obadr", columns[i], 1) + "," + string(0x0D) + string(0x0A)
			}
			if i == 4 { // OHRANA
				s_json += getQuatedJSON("status", columns[i], 1) + "," + string(0x0D) + string(0x0A)
			}
			if i == 5 { // SVIZI
				s_json += getQuatedJSON("con", columns[i], 1) + "," + string(0x0D) + string(0x0A)
			}
			if i == 6 { // MOREINFO
				s_json += getQuatedJSON("more", columns[i], 1) + "," + string(0x0D) + string(0x0A)
			}
			if i == 7 { // INFODETAIL
				s_json += getQuatedJSON("detail", columns[i], 1) + "," + string(0x0D) + string(0x0A)
			}
			if i == 8 { // PULTNUMBER
				s_json += getQuatedJSON("pult", columns[i], 1) + "," + string(0x0D) + string(0x0A)
			}
			if i == 9 { // LONGITUDE
				s_json += getQuatedJSON("lon", columns[i], 1) + "," + string(0x0D) + string(0x0A)
			}
			if i == 10 { // LATITUDE
				s_json += getQuatedJSON("lat", columns[i], 1) + string(0x0D) + string(0x0A)
				s_json += "}" + string(0x0D) + string(0x0A)
			}
		}
	}
	if isKey == false {
		s_json += "]"
	}
	if rows != nil {
		defer rows.Close()
	}
	return s_json

}

//------------------------------------------------------------------------------
func getZoneUserList(objUIN string, infoListType int) string {
	s_sql := "SELECT ZONENUM,OBID,OPYSZONE,ZONEKIND,LINKZONA FROM zonelist "
	if infoListType == 0 { // Zones
		s_sql += "WHERE (ZONEKIND<50) AND (OBID="
		s_sql += objUIN + ") ORDER BY ZONENUM"
	} else if infoListType == 1 { // Users
		s_sql += "WHERE  ((ZONEKIND=100) OR (ZONEKIND=101)) AND (OBID="
		s_sql += objUIN + ") ORDER BY ZONENUM"
	} else { //Maps and Images
		s_sql = "SELECT IDMAPA,MAPAOBJ,MAPAOPYS,MAPAKIND,MAPAWAY FROM mapalist "
		s_sql += "WHERE MAPAOBJ=" + objUIN + " ORDER BY MAPAOPYS"
	}
	db, err := sqlx.Connect("mysql", connect_db)
	s_json := ""
	if err != nil {
		defer recoveryAppFunction()
		fmt.Println(getDT(), "Error read value :"+s_sql, err)
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query(s_sql)
	if err != nil {
		defer recoveryAppFunction()
		fmt.Println("Error read value :" + s_sql)
	}
	cols, err := rows.Columns()
	if err != nil {
		defer recoveryAppFunction()
		fmt.Println("Error read value :" + s_sql)
	}
	data := make(map[string]string)
	for rows.Next() {
		columns := make([]string, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}
		rows.Scan(columnPointers...)
		for i, colName := range cols {
			data[colName] = columns[i]
			if i == 0 { // 0 - Get ZONENUM / IDMAPA
				if len(s_json) > 5 {
					s_json += string(0x0D) + string(0x0A)
				}
				s_json += "{" + string(0x0D) + string(0x0A)
				s_json += getQuatedJSON("num", columns[i], 1) + "," + string(0x0D) + string(0x0A)
			}
			if i == 2 { // 1 - OPYSZONE
				s_json += getQuatedJSON("name", columns[i], 1) + "," + string(0x0D) + string(0x0A)
			}
			if i == 4 { // LINKZONA
				if infoListType == 2 {
					s_json += getQuatedJSON("lnk", columns[i], 1) + string(0x0D) + string(0x0A)
				} else {
					s_json += getQuatedJSON("tel", columns[i], 1) + string(0x0D) + string(0x0A)
				}
				s_json += "},"
			}
		}
	}
	if rows != nil {
		defer rows.Close()
	}
	return s_json

}

//------------------------------------------------------------------------------
func getGBRuser(update_GBR_id string) string {
	s_sql := "SELECT IDCONST,CONSTDOP FROM  consttable WHERE IDCONST=" + update_GBR_id + " LIMIT 1"
	s_uin := dbGetStringData(s_sql, 1)
	if len(s_uin) < 1 {
		s_uin = "0"
	}
	return s_uin
}

//------------------------------------------------------------------------------
func updateGBRstatus(update_GBR_id, update_USER_id, update_GEO, update_REPORT string, update_GBR_status int) {
	/*
		update_GBR_status
		0 - Login USER
		1 - Start Alarm
		2 - Point Alarm
		3 - Break Alarm
		4 - Stop Alarm
		5 - Geo Alarm
	*/
	if update_GBR_status == 5 {

		data := []byte(`{
		"status":"alarmstart",
		"param":"X0Y0",
		"id":"123456"
	}`)
		var testJsonUnmarshal sendStatusOfAlarm
		if err := json.Unmarshal(data, &testJsonUnmarshal); err != nil {
			panic(err)
		}
		r := bytes.NewReader(data)
		resp, err := http.Post("http://api-cs.ohholding.com.ua/api/set-status?status=" + testJsonUnmarshal.Status+"&param="+testJsonUnmarshal.Param+"&id="+testJsonUnmarshal.Id, "application/json", r)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v %v", err, resp)
	}

	if update_GBR_status == 4 {

		data := []byte(`{
		"status":"alarmstop",
		"param":"X0Y0",
		"id":"123456"
	}`)
		var testJsonUnmarshal sendStatusOfAlarm
		if err := json.Unmarshal(data, &testJsonUnmarshal); err != nil {
			panic(err)
		}
		r := bytes.NewReader(data)
		resp, err := http.Post("http://api-cs.ohholding.com.ua/api/set-status?status=" + testJsonUnmarshal.Status+"&param="+testJsonUnmarshal.Param+"&id="+testJsonUnmarshal.Id, "application/json", r)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v %v", err, resp)
	}


	if update_GBR_status == 3 {

		data := []byte(`{
		"status":"alarmbreak",
		"param":"X0Y0",
		"id":"123456"
	}`)
		var testJsonUnmarshal sendStatusOfAlarm
		if err := json.Unmarshal(data, &testJsonUnmarshal); err != nil {
			panic(err)
		}
		r := bytes.NewReader(data)
		resp, err := http.Post("http://api-cs.ohholding.com.ua/api/set-status?status=" + testJsonUnmarshal.Status+"&param="+testJsonUnmarshal.Param+"&id="+testJsonUnmarshal.Id, "application/json", r)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v %v", err, resp)
	}

	if update_GBR_status == 2 {
		data := []byte(`{
		"status":"alarmpoint",
		"param":"X0Y0",
		"id":"123456"
	}`)
		var testJsonUnmarshal sendStatusOfAlarm
		if err := json.Unmarshal(data, &testJsonUnmarshal); err != nil {
			panic(err)
		}
		r := bytes.NewReader(data)
		resp, err := http.Post("http://api-cs.ohholding.com.ua/api/set-status?status=" + testJsonUnmarshal.Status+"&param="+testJsonUnmarshal.Param+"&id="+testJsonUnmarshal.Id, "application/json", r)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v %v", err, resp)
	}

	s_sql := "INSERT INTO gbrhistory (GBRIDH,GBRUSERH,GBRSTATUSH,GBRGEOH,GBRTIMEH,GBRREPORT) VALUES ("
	s_sql += update_GBR_id + "," + update_USER_id + "," + strconv.Itoa(update_GBR_status) + ","
	s_sql += dbQuatedString(update_GEO) + "," + delphiDateToSQL(time.Now().Unix()) + ","
	s_sql += dbQuatedString(update_REPORT) + ")"
	dbUpdateData(s_sql)

	if update_GBR_status == 0 {
		s_sql = "UPDATE consttable SET CONSTDATE=" + delphiDateToSQL(time.Now().Unix())
		s_sql += ",GBRUSERH=" + update_USER_id
		if len(update_REPORT) > 10 {
			s_sql += ",GBRTOCKEN=" + dbQuatedString(update_REPORT)
		}

		s_sql += " WHERE (CONSTKIND = 4) AND (IDCONST = " + update_GBR_id + ") LIMIT 1"
		fmt.Println("Find tocken", s_sql)
		dbUpdateData(s_sql)
	}

}

//=====================AUTO UPDATE MODULES =====================================
func checkUpdateView() {
	if uptList {
		updateSockList()
	}
	uptList = false
}

//------------------------------------------------------------------------------
func checkUpdateAlarms() {
	//f_date := float64(0)
	if uptAlarm == false { //Wasn't update event
		s_date := dbGetStringData("SELECT IDEQ,SHVYDKIST FROM equiplist WHERE IDEQ=0 LIMIT 1", 1)
		f_date := strToFloat64(s_date)
		if lastAlarm != f_date {
			lastAlarm = f_date
			uptAlarm = true

		}
	}
	if uptAlarm {
		lastUIN = getLastUin()
	}
	checkUpdateView()
	uptAlarm = false
}

//------------------------------------------------------------------------------
func getLastUin() int64 {
	i_uin := lastUIN
	s_sql := "SELECT IDEV,GBRID,OBJECTID,CODEID,RANG,ISNEW FROM eventlist WHERE "
	s_sql += "(ISNEW>0) AND (GBRID=0) AND (IDEV>" + strconv.FormatInt(lastUIN, 10) + ") ORDER BY IDEV"
	db, err := sqlx.Connect("mysql", connect_db)
	if err != nil {
		defer recoveryAppFunction()
		fmt.Println(getDT(), "Error read value :"+s_sql, err)
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query(s_sql)
	if err != nil {
		defer recoveryAppFunction()
		fmt.Println("Error read value :" + s_sql)
	}
	cols, err := rows.Columns()
	if err != nil {
		defer recoveryAppFunction()
		fmt.Println("Error read value :" + s_sql)
	}
	data := make(map[string]string)
	for rows.Next() {
		i_gbr := 0
		i_obj := 0
		i_code := 0
		columns := make([]string, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}
		rows.Scan(columnPointers...)
		for i, colName := range cols {
			data[colName] = columns[i]
			if i == 0 { // 0 - IDEV
				i_temp_uin := convertInt64Val(columns[i])
				if i_temp_uin > 0 {
					i_uin = i_temp_uin
				}

			}
			if i == 1 { // 1 - GBRID
				i_gbr = convertIntVal(columns[i])
			}
			if i == 2 { // OBJECTID
				i_obj = convertIntVal(columns[i])
			}
			if i == 3 { // CODEID
				i_code = convertIntVal(columns[i])
			}
			if i == 4 { //RANG
				if convertIntVal(columns[i]) == 1 && i_obj > 0 {
					uptList = true
					s_msg := dbGetStringData("SELECT IDOB,OBNAME FROM objectlist WHERE IDOB="+strconv.Itoa(i_obj), 1) + "; "
					if i_code > 0 {
						s_msg += dbGetStringData("SELECT IDCODE,EVENTS FROM codelist WHERE IDCODE="+strconv.Itoa(i_code), 1)
					}
					sendFcmToGbr(i_obj, i_gbr, 0, s_msg)
				}
			}
		}
	}
	if rows != nil {
		defer rows.Close()
	}

	return i_uin
}

//SELECT objectlist.IDOB,objectlist.NUM,gbrlist.IDGBR,gbrlist.GBREQID,gbrlist.GBRPERSID,consttable.CONSTVALUE,consttable
//------------------------------------------------------------------------------
func sendFcmToGbr(uinOBJECT, uinGROUP, sendValueType int, fcmMESSAGE string) {
	s_sql := "SELECT IDOB,NUM FROM objectlist WHERE IDOB=" + strconv.Itoa(uinOBJECT) + " LIMIT 1"
	i_eqid := dbGetIntData(s_sql, 1)
	if i_eqid > 0 {
		s_sql = "SELECT objectlist.IDOB,objectlist.NUM,gbrlist.IDGBR,gbrlist.GBREQID,gbrlist.GBRPERSID,"
		s_sql += "consttable.CONSTVALUE,consttable.IDCONST,IFNULL(consttable.GBRTOCKEN,'') "
		s_sql += "FROM consttable "
		s_sql += "inner join gbrlist on (consttable.IDCONST = gbrlist.GBRPERSID) "
		s_sql += "inner join objectlist on (gbrlist.GBREQID = objectlist.NUM) "
		s_sql += "WHERE GBREQID=" + strconv.Itoa(i_eqid) // + " LIMIT 1"
		//--------------------------------------------------------------------
		db, err := sqlx.Connect("mysql", connect_db)
		if err != nil {
			defer recoveryAppFunction()
			fmt.Println(getDT(), "Error read value :"+s_sql, err)
			panic(err)
		}
		defer db.Close()
		rows, err := db.Query(s_sql)
		if err != nil {
			defer recoveryAppFunction()
			fmt.Println("Error read value :" + s_sql)
		}
		cols, err := rows.Columns()
		if err != nil {
			defer recoveryAppFunction()
			fmt.Println("Error read value :" + s_sql)
		}
		data := make(map[string]string)
		for rows.Next() {
			s_gbr := ""
			columns := make([]string, len(cols))
			columnPointers := make([]interface{}, len(cols))
			for i, _ := range columns {
				columnPointers[i] = &columns[i]
			}
			rows.Scan(columnPointers...)
			for i, colName := range cols {
				data[colName] = columns[i]
				if i == 0 { //3
					s_gbr = columns[i]
					fmt.Println("Find GBR", s_gbr, fcmMESSAGE)
				}
				if i == 7 { //TOCKEN
					s_TOCKEN := columns[i]
					if len(s_TOCKEN) > 5 {
						//		checkDroid()
						getTokenList(s_TOCKEN, s_gbr, fcmMESSAGE, sendValueType)
					}
				}
			}
		}
		if rows != nil {
			defer rows.Close()
		}

		//--------------------------------------------------------------------
	}

}

//------------------------------------------------------------------------------
func sendAlarmToGbr() {

	s_sql := "SELECT regtemp.IDREG,regtemp.IDOBJ,regtemp.REGRESULT,objectlist.IDOB,objectlist.OBNAME,"
	s_sql += "consttable.CONSTVALUE,consttable.IDCONST,IFNULL(consttable.GBRTOCKEN,'') "
	s_sql += "FROM consttable "
	s_sql += "inner join regtemp on (consttable.IDCONST = regtemp.IDREG) "
	s_sql += "inner join objectlist on (regtemp.IDOBJ = objectlist.IDOB) "
	s_sql += "WHERE REGRESULT<4" // LIMIT 1
	//--------------------------------------------------------------------
	db, err := sqlx.Connect("mysql", connect_db)
	if err != nil {
		defer recoveryAppFunction()
		fmt.Println(getDT(), "Error read value :"+s_sql, err)
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query(s_sql)
	if err != nil {
		defer recoveryAppFunction()
		fmt.Println("Error read value :" + s_sql)
	}
	cols, err := rows.Columns()
	if err != nil {
		defer recoveryAppFunction()
		fmt.Println("Error read value :" + s_sql)
	}
	data := make(map[string]string)
	for rows.Next() {
		s_object := ""
		s_uin := ""
		//s_gbr := ""
		i_status := 0
		columns := make([]string, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}
		rows.Scan(columnPointers...)
		for i, colName := range cols {
			data[colName] = columns[i]
			if i == 0 {
				//s_gbr = columns[i]
			}

			if i == 2 {
				i_status = convertIntVal(columns[i])
			}

			if i == 3 {
				s_uin = columns[i]
			}

			if i == 4 {
				if i_status < 2 {
					s_object = columns[i] + "; Выезд на тревогу"
				} else if i_status == 2 {
					s_object = columns[i] + "; Отмена выезда на тревогу"
				} else {
					s_object = columns[i] + "; Завершение обработки тревоги"
				}
			}
			if i == 7 { //TOCKEN
				s_TOCKEN := columns[i]
				if len(s_TOCKEN) > 5 {
					//		checkDroid()

					getTokenList(s_TOCKEN, s_uin, s_object, i_status)
				}
			}
		}
	}
	if rows != nil {
		defer rows.Close()
	}

}

//--------------------------------------------------------------------