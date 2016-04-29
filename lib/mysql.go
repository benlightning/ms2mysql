package ms2mysql

func GetMaxTime() string {
	sql := "SELECT checktime FROM intro_checkin ORDER BY checktime DESC limit 1"
	rowArray, _ := MyEngine.Query(sql)
	timeStr := ""
	// fmt.Println(len(rowArray))
	for _, row := range rowArray {
		for _, colvalue := range row {
			v, _ := Decode(colvalue)
			timeStr = BytesToString(v)
		}
	}
	return Substr(timeStr, 0, 19)
}

func InsertTo(s []map[string]string) {
	var sql string
	var userid, checktime, uname, checkdate string
	for _, value := range s {
		userid = value["userid"]
		checktime = value["checktime"]
		checkdate = value["checkdate"]
		uname = value["uname"]
		sql = "INSERT INTO `intro_checkin` (`checktime`, `userid`, `uname`, `checkdate`) VALUES ('" + checktime + "', " + userid + ", '" + uname + "','" + checkdate + "') ON DUPLICATE KEY UPDATE `checktime`=concat('" + checktime + ",',`checktime`);"
		_, err := MyEngine.Exec(sql)
		if err != nil {
			Logger.Println("failed to insert:", sql)
		}
	}
}
