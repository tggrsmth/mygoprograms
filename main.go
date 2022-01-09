program main

import (
    "fmt"
    "log"
    "net/http"
    "database/sql"
    "strconv"
    
    _ "github.com/mattn/go-sqlite3"
)

func getHostData(){

}

func getContainerData(){

}

func handleRequests(){
    http.HandleFunc("/host", getHostData)
    http.HandleFunc("/container", getContainerData)
    log.Fatal(http.ListenandServe(":10000", nil))
}

func main(){
    database, _ := sql.Open("sqlite3","./aqua.db")	
    statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS containers(id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,host_id INTEGER NOT NULL,name TEXT NOT NULL,image_name INTEGER NOT NULL,CONSTRAINT containers_FK FOREIGN KEY (host_id) REFERENCES hosts(id))")
    statement.Exec()
    statement.
    handleRequests()

}
