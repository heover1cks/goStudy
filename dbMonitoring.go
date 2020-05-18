package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	//"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo"
	"log"
)


type Log struct {
	logType		string
	logValue	int
	dbPrefix	string
	time		string
}

type dbInfo struct {
	user			string
	pwd				string
	url				string
	engine			string
	database		string
	timeTotSess		int
	timeConnUtil	int
	timeBlkThrCn	int
	timeRunThrCn	int
	timeMtlCnt		int
	timeTxnCnt		int
	timeLngQuery	int
}

type dbList =[]dbInfo
var db1 = dbInfo{"root","!dlatl00","localhost:3306","mysql","test",1,1,1,1,1,1,1}
var curDbList = dbList{db1}

const totSess	="SELECT COUNT(*) FROM Information_schema.processlist;"
const connUtil	="select round((select Count(*) from  information_schema.processlist  )*100/variable_value,0) from performance_schema.global_variables where variable_name=\"MAX_CONNECTIONS\""
const blkThrCnt	="SELECT COUNT(*) FROM INFORMATION_SCHEMA.INNODB_LOCK_WAITS"
const runThrCnt	="SELECT COUNT(*) FROM information_schema.PROCESSLIST WHERE COMMAND = 'Query'"
const mtlCnt	="select count(*)  from information_schema.processlist p where state like '%meta%'"
const txnCnt	="SELECT count(c.trx_mysql_thread_id) AS lock_count FROM information_schema.innodb_trx AS a JOIN information_schema.innodb_lock_waits b ON (a.trx_id = b.requesting_trx_id) JOIN information_schema.innodb_trx AS c ON (c.trx_id = b.blocking_trx_id) where timestampdiff(SECOND, a.trx_wait_started, now()) > 1800;"
const lngQuery	="SELECT COUNT(*) FROM information_schema.PROCESSLIST WHERE COMMAND = 'Query' AND TIME > 1800"

func dbConnect(db dbInfo,query string) (count int) {
	dataSource := db.user+":"+db.pwd+"@tcp("+db.url+")/"+db.database
	conn,err := sql.Open(db.engine,dataSource)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	err = conn.QueryRow(query).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
	return count
}

func dbQuery(){

}

func main() {
	result := dbConnect(db1,totSess)
	print("Query Result: ",result)
}

 go run dbMonitoring.go
3
Query Result: 3