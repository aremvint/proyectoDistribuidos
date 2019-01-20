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
  Fecha string
  Lugar string
  Nombre string
  IdAdministrador *Administrador
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

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM Evento ORDER BY idevento DESC")
    if err != nil {
        panic(err.Error())
    }
    eve := Evento{}
    res := []Evento{}
    for selDB.Next() {
        var id int
        var fecha, lugar, nombre string
        err = selDB.Scan(&id, &fecha, &lugar, &nombre)
        if err != nil {
            panic(err.Error())
        }
        eve.Id = id
        eve.Fecha = fecha
        eve.Lugar = lugar
        eve.Nombre = nombre
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
        var id int
        var fecha, lugar, nombre string
        err = selDB.Scan(&id, &fecha, &lugar, &nombre)
        if err != nil {
            panic(err.Error())
        }
        eve.Id = id
        eve.Fecha = fecha
        eve.Lugar = lugar
        eve.Nombre = nombre
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
        var id int
        var fecha, lugar, nombre string
        err = selDB.Scan(&id, &fecha, &lugar, &nombre)
        if err != nil {
            panic(err.Error())
        }
        eve.Id = id
        eve.Fecha = fecha
        eve.Lugar = lugar
        eve.Nombre = nombre
    }
    tmpl.ExecuteTemplate(w, "Edit", eve)
    defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        id := r.FormValue("id")
        fecha := r.FormValue("fecha")
        lugar := r.FormValue("lugar")
        nombre := r.FormValue("nombre")
        insForm, err := db.Prepare("INSERT INTO Evento(idevento, fecha_evento, lugar_evento, nombre_evento) VALUES(?,?,?,?)")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(id, fecha, lugar, nombre)
        log.Println("INSERT: Nombre: " + nombre + " | Id: " + id)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        fecha := r.FormValue("fecha")
        lugar := r.FormValue("lugar")
        nombre := r.FormValue("nombre")
        id := r.FormValue("uid")
        insForm, err := db.Prepare("UPDATE Evento SET fecha_evento=?, lugar_evento=?, nombre_evento=? WHERE idevento=?")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(fecha, lugar, nombre, id)
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
