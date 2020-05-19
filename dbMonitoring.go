package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
	"sync"
	"time"
)


type logEntity struct {
	//ID			primitive.ObjectID `bson:"_id"`
	logType		string	`bson:"logType"`
	logValue	string	`bson:"logValue"`
	dbPrefix	string	`bson:"dbPrefix"`
	time		string	`bson:"time"`
}

type queryInfo struct {
	queryName		string
	querySchedule	int
	query			string
}

type dbInfo struct {
	alias			string
	user			string
	pwd				string
	url				string
	engine			string
	database		string
	queryInfos		map[string]queryInfo
}

type dbList =[]dbInfo
//var schedule1 = queryPeriod{1,1,1,1,1,1,1}
var db1 = dbInfo{"mysql-test-localhost","root","!dlatl00","localhost:3306","mysql","test",queries}
var curDbList = dbList{db1}


var queries = map[string]queryInfo{
	"totSess"	:queryInfo{"totSess",2000,"SELECT COUNT(*) FROM Information_schema.processlist;"},
	"connUtil"	:queryInfo{"connUtil",5000,"select round((select Count(*) from  information_schema.processlist  )*100/variable_value,0) from performance_schema.global_variables where variable_name=\"MAX_CONNECTIONS\""},
	"blkThrCnt"	:queryInfo{"blkThrCnt",10000,"SELECT COUNT(*) FROM INFORMATION_SCHEMA.INNODB_LOCK_WAITS"},
	"runThrCnt"	:queryInfo{"runThrCnt",3000,"SELECT COUNT(*) FROM information_schema.PROCESSLIST WHERE COMMAND = 'Query'"},
	"mtlCnt"	:queryInfo{"mtlCnt",2000,"select count(*)  from information_schema.processlist p where state like '%meta%'"},
	"txnCnt"	:queryInfo{"txnCnt",4000,"SELECT count(c.trx_mysql_thread_id) AS lock_count FROM information_schema.innodb_trx AS a JOIN information_schema.innodb_lock_waits b ON (a.trx_id = b.requesting_trx_id) JOIN information_schema.innodb_trx AS c ON (c.trx_id = b.blocking_trx_id) where timestampdiff(SECOND, a.trx_wait_started, now()) > 1800;"},
	"lngQuery"	:queryInfo{"lngQuery",10000,"SELECT COUNT(*) FROM information_schema.PROCESSLIST WHERE COMMAND = 'Query' AND TIME > 1800"},
}


func dbConnect(db dbInfo) (dbConn *sql.DB){
	dataSource := db.user+":"+db.pwd+"@tcp("+db.url+")/"+db.database
	conn,err := sql.Open(db.engine,dataSource)
	if err != nil {
		log.Fatal(err)
	}
	//defer conn.Close()
	fmt.Println("MySQL Connection")
	return conn
}

func dbQuery(db *sql.DB,writeDb *mongo.Collection, dbInfo dbInfo, queryInfo queryInfo) {
	for {
		fmt.Println("Inside For")
		var count int
		err := db.QueryRow(queryInfo.query).Scan(&count)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(count)
		logWrite(queryInfo.queryName, count, dbInfo, writeDb)
		//return count
		time.Sleep(time.Millisecond * time.Duration(queryInfo.querySchedule))
	}
}

func mongoConn()(client *mongo.Client){
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection")
	return client
}


func logWrite(name string, value int, db dbInfo, writeDb *mongo.Collection){
	currentTime:= time.Now().Format(time.RFC3339)
	currentLog := logEntity{name,strconv.Itoa(value),db.alias,currentTime}
	consoleLog := "["+name+":"+strconv.Itoa(value)+"]"+"["+db.alias+"]"+"["+currentTime+"]"
	fmt.Println(currentLog)
	//insertResult, err := writeDb.InsertOne(context.TODO(), &logEntity{
	//	logType: name,
	//	logValue: strconv.Itoa(value),
	//	dbPrefix: db.alias,
	//	time:currentTime,
	//})
	insertResult,err := writeDb.InsertOne(context.TODO(),bson.D{
		{Key:"logType",Value:name},
		{Key:"logValue",Value:strconv.Itoa(value)},
		{Key:"dbPrefix",Value:db.alias},
		{Key:"time",Value:currentTime},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Insert Result: ",insertResult.InsertedID)
	fmt.Println("Inserted Log: ",consoleLog)
}

func main() {
	var wait sync.WaitGroup
	wait.Add(len(curDbList))
	//MongoDB Connection
	conn := mongoConn()
	mongo := conn.Database("log_data").Collection("logs")


	//Query to MySQL -> DB List 순환 + Queries 순환 작
	for _,thisDb := range curDbList {
		dbConn := dbConnect(thisDb)
		go dbQuery(dbConn, mongo, thisDb, thisDb.queryInfos["totSess"])
		go dbQuery(dbConn, mongo, thisDb, thisDb.queryInfos["connUtil"])
		//go dbQuery(dbConn, mongo, thisDb, thisDb.queryInfos["blkThrCnt"])
		go dbQuery(dbConn, mongo, thisDb, thisDb.queryInfos["runThrCnt"])
		//go dbQuery(dbConn, mongo, thisDb, thisDb.queryInfos["mtlCnt"])
		//go dbQuery(dbConn, mongo, thisDb, thisDb.queryInfos["txnCnt"])
		//go dbQuery(dbConn, mongo, thisDb, thisDb.queryInfos["lngQuery"])
	}
	//defer dbConn.Close()

	//err := conn.Disconnect(context.TODO())
	//if err != nil {
	//	log.Fatal(err)
	//}

	//result := dbConnect(db1,totSess)
	//print("Query Result: ",result)
	wait.Wait()
}
