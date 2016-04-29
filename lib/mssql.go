package ms2mysql

func GetData(time string) []map[string]string {
	var sql string
	if len(time) > 0 {
		sql = "SELECT CK.USERID,CONVERT(varchar, CK.CHECKTIME, 120) as CHECKTIME,UI.USERID,UI.NAME FROM CHECKINOUT CK LEFT JOIN USERINFO UI ON CK.USERID=UI.USERID WHERE CHECKTIME>'" + time + "' ORDER BY CK.USERID ASC, CHECKTIME DESC;"
	} else {
		sql = "SELECT CK.USERID,CONVERT(varchar, CK.CHECKTIME, 120) as CHECKTIME,UI.USERID,UI.NAME FROM CHECKINOUT CK LEFT JOIN USERINFO UI ON CK.USERID=UI.USERID ORDER BY CK.USERID ASC, CHECKTIME DESC;"
	}
	rowArray, _ := MsEngine.Query(sql)
	var userid, checktime, uname, checkdate, oldid, olddate string
	var uid, ctime, name []byte
	//var m map[string]string = {"userid":"","checktime":"","checkdate":"","uname":""}
	m := make(map[string]string)
	var s []map[string]string
	len := len(rowArray)
	i := 1
	for _, row := range rowArray {
		uid, _ = Decode(row["USERID"])
		ctime, _ = Decode(row["CHECKTIME"])
		name, _ = Decode(row["NAME"])

		userid = BytesToString(uid)
		checktime = BytesToString(ctime)
		checkdate = Substr(checktime, 0, 10)
		uname = BytesToString(name)

		if oldid == "" {
			oldid = userid
			olddate = checkdate
			m["userid"] = userid
			m["checktime"] = checktime
			m["checkdate"] = checkdate
			m["uname"] = uname
		} else if oldid == userid && olddate == checkdate {
			m["checktime"] = m["checktime"] + "," + checktime
		} else {
			m1 := make(map[string]string)
			for key, value := range m {
				m1[key] = value
			}
			s = append(s, m1)
			oldid = userid
			olddate = checkdate
			m["userid"] = userid
			m["checktime"] = checktime
			m["checkdate"] = checkdate
			m["uname"] = uname
		}
		if len == i {
			m2 := make(map[string]string)
			for key, value := range m {
				m2[key] = value
			}
			s = append(s, m2)
		}
		i = i + 1
		//sql = "INSERT INTO `intro_checkin` (`checktime`, `userid`, `uname`) VALUES ('" + checktime + "', " + userid + ", '" + uname + "');"
		//_, err := MyEngine.Exec(sql)
		//if err != nil {
		//	Logger.Println("failed to insert:", sql)
		//}
	}
	return s
}
