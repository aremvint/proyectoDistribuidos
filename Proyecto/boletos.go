package boletos


type Boleto struct{
  IdBoleto int
  Precio float64
  Comprado int
  Asiento string
  Venue string
  Categoria string
}

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
