package main

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"time"
)

func doWork(db *sql.DB, work int64) {
	// work here
}

func getWork(db *sql.DB) {
	for {
		// get work from the database here
		var work sql.NullInt64
		err := db.QueryRow("SELECT get_work()").Scan(&work)
		if err != nil {
			fmt.Println("call to get_work() failed: ", err)
			time.Sleep(10 * time.Second)
			continue
		}
		if !work.Valid {
			// no more work to do
			fmt.Println("ran out of work")
			return
		}

		fmt.Println("starting work on ", work.Int64)
		go doWork(db, work.Int64)
	}
}

func waitForNotification(l *pq.Listener) {
	select {
	case <-l.Notify:
		fmt.Println("received notification, new work available")
	case <-time.After(90 * time.Second):
		go l.Ping()
		// Check if there's more work available, just in case it takes
		// a while for the Listener to notice connection loss and
		// reconnect.
		fmt.Println("received no work for 90 seconds, checking for new work")
	}
}

func main() {
	var conninfo string = "postgres://postgres:postgresql2016@192.168.199.216:5432/pinto?sslmode=disable"

	db, err := sql.Open("postgres", conninfo)
	if err != nil {
		panic(err)
	}

	reportProblem := func(ev pq.ListenerEventType, err error) {
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	listener := pq.NewListener(conninfo, 10*time.Second, time.Minute, reportProblem)
	err = listener.Listen("getwork")
	if err != nil {
		panic(err)
	}

	fmt.Println("entering main loop")
	for {
		// process all available work before waiting for notifications
		getWork(db)
		waitForNotification(listener)
	}
}
