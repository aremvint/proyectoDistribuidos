package main

import(
  "database/sql"
  "log"
  "net/http"
  "text/template"

  _ "github.com/go-sql-driver/mysql"
)

type Boleto struct{
  IdBoleto int
  Precio float
  Comprado int
  idAsiento int
  idVenue int
  idCategoria int
}


type Asiento struct{
  idAsiento int
  NumeroAsiento string
  idVenue int
  idCategoria int
}

func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "password"
    dbName := "ventas"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM Boleto WHERE comprado=0")
    if err != nil {
        panic(err.Error())
    }
    bol := Boleto{}
    res := []Boleto{}
    for selDB.Next() {
        var idboleto, idcategoria, idvenue, idasiento, comprado int
        var precio float
        err = selDB.Scan(&idboleto, &idcategoria, &idvenue, &idasiento, &comprado, &precio)
        if err != nil {
            panic(err.Error())
        }
        bol.IdBoleto = idboleto
        bol.idCategoria = idcategoria
        bol.IdVenue = idvenue
        bol.IdAsiento = idasiento
        bol.Comprado = comprado
        bol.Precio = precio
        res = append(res, bol)
    }
    tmpl.ExecuteTemplate(w, "Index", res)
    defer db.Close()
}
