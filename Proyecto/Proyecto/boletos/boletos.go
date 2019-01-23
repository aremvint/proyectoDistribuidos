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
  Precio float64
  Comprado int
  Asiento string
  Venue string
  Categoria string
}

func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "1234"
    dbName := "ventas"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}

var tmpl = template.Must(template.ParseGlob("form-boletos/*"))

func Ticket(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM Boleto WHERE comprado=0")
    if err != nil {
        panic(err.Error())
    }
    bol := Boleto{}
    res := []Boleto{}
    for selDB.Next() {
        var idboleto, idcategoria, idVenue, idasiento, comprado int
        var precio float64
        err = selDB.Scan(&idboleto, &precio, &comprado, &idasiento, &idcategoria, &idVenue)
        if err != nil {
            panic(err.Error())
        }
        //
        sqelDB, err := db.Query("SELECT * FROM Categoria WHERE idcategoria=?",idcategoria)
        if err != nil {
            panic(err.Error())
        }
        for sqelDB.Next() {
            var idcategoria, cantidad_asientos, Venue_idVenue int
            var nombre string
            err = sqelDB.Scan(&idcategoria, &nombre, &cantidad_asientos, &Venue_idVenue)
            bol.Categoria = nombre
        }
        srelDB, err := db.Query("SELECT * FROM Venue WHERE idVenue=?",idVenue)
        if err != nil {
            panic(err.Error())
        }
        for srelDB.Next() {
            var idVenue, Administrador_idAdministrador int
            var descripcion, tipo_venue string
            err = srelDB.Scan(&idVenue, &tipo_venue, &descripcion, &Administrador_idAdministrador)
            bol.Venue = descripcion
        }
        sselDB, err := db.Query("SELECT * FROM Asiento WHERE idasiento=?",idasiento)
        if err != nil {
            panic(err.Error())
        }
        for sselDB.Next() {
            var idasiento, categoria_idcategoria, categoria_Venue_idVenue int
            var numero_asiento string
            err = sselDB.Scan(&idasiento, &numero_asiento, &categoria_idcategoria, &categoria_Venue_idVenue)
            bol.Asiento = numero_asiento
        }//
        bol.IdBoleto = idboleto
        bol.Comprado = comprado
        bol.Precio = precio
        res = append(res, bol)
    }
    tmpl.ExecuteTemplate(w, "Ticket", res)
    log.Println(bol)
    defer db.Close()
}

func Mostrar(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")
    selDB, err := db.Query("SELECT * FROM Boleto WHERE idboleto=?", nId)
    if err != nil {
        panic(err.Error())
    }
    bol := Boleto{}
    for selDB.Next() {
        var idboleto, idcategoria, idVenue, idasiento, comprado int
        var precio float64
        err = selDB.Scan(&idboleto, &precio, &comprado, &idasiento, &idcategoria, &idVenue)
        if err != nil {
            panic(err.Error())
        }
        bol.IdBoleto = idboleto
        bol.Precio = precio
    }
    tmpl.ExecuteTemplate(w, "Mostrar", bol)
    defer db.Close()
}

func Comprar(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        id := r.FormValue("tid")
        precio := r.FormValue("precio")
        insForm, err := db.Prepare("UPDATE Boleto SET comprado=?  WHERE idboleto=?")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(1, id)
        log.Println("Boleto Comprado: Id: " + id + " | Precio: " + precio)
    }
    defer db.Close()
    http.Redirect(w, r, "/ticket", 301)
}


func main() {
    log.Println("Server started on: http://localhost:8080")
    http.HandleFunc("/ticket", Ticket)
    http.HandleFunc("/mostrar", Mostrar)
    http.HandleFunc("/comprar", Comprar)
    http.ListenAndServe(":8080", nil)
}
