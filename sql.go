package main                                                                                                                                                                                                                        
import "database/sql"                                                                                                                                                                                                               
import "fmt"                                                                                                                                                                                                                        
import _ "github.com/go-sql-driver/mysql"                                                                                                                                                                                           

func main() {                                                                                                                                                                                                                       
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/x_ucenter?charset=utf8")                                                                                                                                            
        if err != nil {                                                                                                                                                                                                             
            panic(err.Error())                                                                                                                                                                                                  
        }                                                                                                                                                                                                                           
    defer db.Close()                                                                                                                                                                                                            
        stmtOut, err := db.Prepare("Select uid, username from uc_members where uid = ? limit 1;")                                                                                                                                   
        if err != nil {                                                                                                                                                                                                             
            panic(err.Error())                                                                                                                                                                                                  
        }                                                                                                                                                                                                                           
    var uid int                                                                                                                                                                                                                 
        var username string                                                                                                                                                                                                         
        err = stmtOut.QueryRow(196824).Scan(&uid, &username)                                                                                                                                                                        
        if err != nil {                                                                                                                                                                                                             
            panic(err.Error())                                                                                                                                                                                                  
        }                                                                                                                                                                                                                           
    fmt.Printf(" %d => %s \n", uid, username)                                                                                                                                                                                   

}                                                                                                                                                                                                                                   



