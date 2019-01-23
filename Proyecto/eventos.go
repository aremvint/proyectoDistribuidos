package main

import(
  "database/sql"
  "log"
  "net/http"
  "text/template"

  _ "github.com/go-sql-driver/mysql"
)

type Evento struct{
  Id int
  Nombre string
  Fecha string
  Lugar string
  Venue int
  Administrador int
}

type Administrador struct{
  ID int
  Cedula int
  Nombres string
  Apellidos string
  Edad int
  Correo string
  Usuario string
  Contrasena string
  Departamento string
  Cargo string
  Genero string
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

var tmpl = template.Must(template.ParseGlob("form-eventos/*"))

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
        id := r.FormValue("uid")
        nombre := r.FormValue("nombre")
        lugar := r.FormValue("lugar")
        fecha := r.FormValue("fecha")
        insForm, err := db.Prepare("UPDATE Evento SET nombre_evento=?, lugar_evento=?, fecha_evento=?  WHERE idevento=?")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(id, nombre, lugar, fecha)
        log.Println("UPDATE: Nombre: " + nombre + " | Id: " + id)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

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

func main() {
    log.Println("Server started on: http://localhost:8080")
    http.HandleFunc("/", Index)
    http.HandleFunc("/show", Show)
    http.HandleFunc("/new", New)
    http.HandleFunc("/edit", Edit)
    http.HandleFunc("/insert", Insert)
    http.HandleFunc("/update", Update)
    http.HandleFunc("/delete", Delete)
    http.ListenAndServe(":8080", nil)
}
