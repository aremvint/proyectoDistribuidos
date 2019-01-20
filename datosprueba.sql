INSERT INTO `ventas`.`administrador`
(`idAdministrador`,
`cedula_administrador`,
`nombres`,
`apellidos`,
`edad`,
`correo`,
`usuario`,
`contraseña`,
`departamento`,
`cargo`,
`genero`)
VALUES
(10,
0939558607,
"VINICIO",
"ESPINEL",
"35",
"vespin@gmail.com",
"vespin",
"1234",
"Ventas",
"Admin",
"M");

INSERT INTO `ventas`.`venue`
(`idVenue`,
`tipo_venue`,
`descripcion`,
`Administrador_idAdministrador`)
VALUES
(100,
"Teatro",
"Teatro Perez Peruzzi",
10);

INSERT INTO `ventas`.`evento`
(`idevento`,
`nombre_evento`,
`lugar_evento`,
`fecha_evento`,
`Venue_idVenue`,
`Administrador_idAdministrador`)
VALUES
(3000,
"CONCIERTO SALSA",
"TEATRO PEREZ PERUZZI",
'2018-12-05 12:30:00',
100,
10);
INSERT INTO `ventas`.`categoria`
(`idcategoria`,
`nombre`,
`cantidad_asientos`,
`Venue_idVenue`)
VALUES
(15000,
"PALCO",
500,
100);
INSERT INTO `ventas`.`asiento`
(`idasiento`,
`numero_asiento`,
`categoria_idcategoria`,
`categoria_Venue_idVenue`)
VALUES
(15001,
5,
15000,
100);
INSERT INTO `ventas`.`boleto`
(`idboleto`,
`precio`,
`comprado`,
`asiento_idasiento`,
`asiento_categoria_idcategoria`,
`asiento_categoria_Venue_idVenue`)
VALUES
(4578,
10.50,
0,
15001,
15000,
100);
