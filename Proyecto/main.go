package main

import(
  "database/sql"
  "log"
  "net/http"
  "text/template"

  _ "github.com/go-sql-driver/mysql"
)

//Estructuras
type Evento struct{
  Id int
  Nombre string
  Fecha string
  Lugar string
  Venue int
  Administrador int
}

type Boleto struct{
  IdBoleto int
  Precio float64
  Comprado int
  Asiento string
  Venue string
  Categoria string
}

//

//Función para conexión a la base
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

//Funciones para Eventos
func Index(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM Evento ORDER BY idevento DESC")
    if err != nil {
        panic(err.Error())
    }
    eve := Evento{}
    res := []Evento{}
    for selDB.Next() {
        var id, venue, administrador int
        var fecha, lugar, nombre string
        err = selDB.Scan(&id, &nombre, &lugar, &fecha, &venue, &administrador)
        if err != nil {
            panic(err.Error())
        }
        eve.Id = id
        eve.Nombre = nombre
        eve.Lugar = lugar
        eve.Fecha = fecha
        res = append(res, eve)
    }
    tmpl.ExecuteTemplate(w, "Index", res)
    defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")
    selDB, err := db.Query("SELECT * FROM Evento WHERE idevento=?", nId)
    if err != nil {
        panic(err.Error())
    }
    eve := Evento{}
    for selDB.Next() {
        var id, venue, administrador int
        var fecha, lugar, nombre string
        err = selDB.Scan(&id, &nombre, &lugar, &fecha, &venue, &administrador)
        if err != nil {
            panic(err.Error())
        }
        eve.Id = id
        eve.Nombre = nombre
        eve.Lugar = lugar
        eve.Fecha = fecha
    }
    tmpl.ExecuteTemplate(w, "Show", eve)
    defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
    tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")
    selDB, err := db.Query("SELECT * FROM Evento WHERE idevento=?", nId)
    if err != nil {
        panic(err.Error())
    }
    eve := Evento{}
    for selDB.Next() {
      var id, venue, administrador int
      var fecha, lugar, nombre string
      err = selDB.Scan(&id, &nombre, &lugar, &fecha, &venue, &administrador)
        if err != nil {
            panic(err.Error())
        }
        eve.Id = id
        eve.Nombre = nombre
        eve.Lugar = lugar
        eve.Fecha = fecha
    }
    tmpl.ExecuteTemplate(w, "Edit", eve)
    defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        nombre := r.FormValue("nombre")
        lugar := r.FormValue("lugar")
        fecha := r.FormValue("fecha")
        insForm, err := db.Prepare("INSERT INTO Evento(nombre_evento, lugar_evento, fecha_evento, Venue_idVenue, Administrador_idAdministrador ) VALUES(?,?,?,100,10)")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(nombre, lugar, fecha)
        log.Println("INSERT: Nombre: " + nombre + " | lugar: " + lugar)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        nombre := r.FormValue("nombre")
        lugar := r.FormValue("lugar")
        fecha := r.FormValue("fecha")
        id := r.FormValue("uid")
        insForm, err := db.Prepare("UPDATE Evento SET nombre_evento=?, lugar_evento=?, fecha_evento=?  WHERE idevento=?")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(nombre, lugar, fecha,id)
        log.Println("UPDATE: Nombre: " + nombre + " | Id: " + id)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}
//

func Delete(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    eve := r.URL.Query().Get("id")
    delForm, err := db.Prepare("DELETE FROM Evento WHERE idevento=?")
    if err != nil {
        panic(err.Error())
    }
    delForm.Exec(eve)
    log.Println("DELETE")
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

//Funciones para Boletos
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
        }
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
//

func main() {
    log.Println("Server started on: http://localhost:8080")
    http.HandleFunc("/", Index)
    http.HandleFunc("/show", Show)
    http.HandleFunc("/new", New)
    http.HandleFunc("/edit", Edit)
    http.HandleFunc("/insert", Insert)
    http.HandleFunc("/update", Update)
    http.HandleFunc("/delete", Delete)
    http.HandleFunc("/ticket", Ticket)
    http.HandleFunc("/mostrar", Mostrar)
    http.HandleFunc("/comprar", Comprar)
    http.ListenAndServe(":8080", nil)
}
