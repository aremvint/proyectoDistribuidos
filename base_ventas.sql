-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema ventas
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema ventas
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `ventas` DEFAULT CHARACTER SET utf8 ;
USE `ventas` ;

-- -----------------------------------------------------
-- Table `ventas`.`Usuario`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `ventas`.`Usuario` (
  `idUsuario` INT NOT NULL AUTO_INCREMENT,
  `cedula_usuario` INT NOT NULL,
  `nombres` VARCHAR(100) NULL,
  `apellidos` VARCHAR(100) NULL,
  `telefono` VARCHAR(45) NULL,
  `correo` VARCHAR(45) NULL,
  `boletos_comprados` INT NULL,
  `nombre_usuario` VARCHAR(45) NOT NULL,
  `contrasena` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`idUsuario`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `ventas`.`Administrador`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `ventas`.`Administrador` (
  `idAdministrador` INT NOT NULL AUTO_INCREMENT,
  `cedula_administrador` INT NOT NULL,
  `nombres` VARCHAR(100) NULL,
  `apellidos` VARCHAR(100) NULL,
  `edad` VARCHAR(45) NULL,
  `correo` VARCHAR(45) NULL,
  `usuario` VARCHAR(45) NOT NULL,
  `contraseña` VARCHAR(45) NOT NULL,
  `departamento` VARCHAR(45) NULL,
  `cargo` VARCHAR(45) NULL,
  `genero` VARCHAR(45) NULL,
  PRIMARY KEY (`idAdministrador`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `ventas`.`Venue`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `ventas`.`Venue` (
  `idVenue` INT NOT NULL,
  `tipo_venue` VARCHAR(45) NULL,
  `descripcion` VARCHAR(45) NULL,
  `Administrador_idAdministrador` INT NOT NULL,
  PRIMARY KEY (`idVenue`, `Administrador_idAdministrador`),
  INDEX `fk_Venue_Administrador1_idx` (`Administrador_idAdministrador` ASC) VISIBLE,
  CONSTRAINT `fk_Venue_Administrador1`
    FOREIGN KEY (`Administrador_idAdministrador`)
    REFERENCES `ventas`.`Administrador` (`idAdministrador`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `ventas`.`categoria`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `ventas`.`categoria` (
  `idcategoria` INT NOT NULL,
  `nombre` VARCHAR(45) NOT NULL,
  `cantidad_asientos` INT NULL,
  `Venue_idVenue` INT NOT NULL,
  PRIMARY KEY (`idcategoria`, `Venue_idVenue`),
  INDEX `fk_categoria_Venue1_idx` (`Venue_idVenue` ASC) VISIBLE,
  CONSTRAINT `fk_categoria_Venue1`
    FOREIGN KEY (`Venue_idVenue`)
    REFERENCES `ventas`.`Venue` (`idVenue`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `ventas`.`asiento`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `ventas`.`asiento` (
  `idasiento` INT NOT NULL,
  `numero_asiento` VARCHAR(45) NOT NULL,
  `categoria_idcategoria` INT NOT NULL,
  `categoria_Venue_idVenue` INT NOT NULL,
  PRIMARY KEY (`idasiento`, `categoria_idcategoria`, `categoria_Venue_idVenue`),
  INDEX `fk_asiento_categoria1_idx` (`categoria_idcategoria` ASC, `categoria_Venue_idVenue` ASC) VISIBLE,
  CONSTRAINT `fk_asiento_categoria1`
    FOREIGN KEY (`categoria_idcategoria` , `categoria_Venue_idVenue`)
    REFERENCES `ventas`.`categoria` (`idcategoria` , `Venue_idVenue`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `ventas`.`boleto`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `ventas`.`boleto` (
  `idboleto` INT NOT NULL,
  `precio` FLOAT NULL,
  `comprado` TINYINT NOT NULL,
  `asiento_idasiento` INT NOT NULL,
  `asiento_categoria_idcategoria` INT NOT NULL,
  `asiento_categoria_Venue_idVenue` INT NOT NULL,
  PRIMARY KEY (`idboleto`, `asiento_idasiento`, `asiento_categoria_idcategoria`, `asiento_categoria_Venue_idVenue`),
  INDEX `fk_boleto_asiento1_idx` (`asiento_idasiento` ASC, `asiento_categoria_idcategoria` ASC, `asiento_categoria_Venue_idVenue` ASC) VISIBLE,
  CONSTRAINT `fk_boleto_asiento1`
    FOREIGN KEY (`asiento_idasiento` , `asiento_categoria_idcategoria` , `asiento_categoria_Venue_idVenue`)
    REFERENCES `ventas`.`asiento` (`idasiento` , `categoria_idcategoria` , `categoria_Venue_idVenue`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `ventas`.`Reporte`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `ventas`.`Reporte` (
  `idReporte` INT NOT NULL,
  `titulo` VARCHAR(45) NULL,
  `fecha` DATETIME NULL,
  `tipo` VARCHAR(45) NULL,
  `descripcion` VARCHAR(500) NULL,
  `Administrador_idAdministrador` INT NOT NULL,
  PRIMARY KEY (`idReporte`, `Administrador_idAdministrador`),
  INDEX `fk_Reporte_Administrador1_idx` (`Administrador_idAdministrador` ASC) VISIBLE,
  CONSTRAINT `fk_Reporte_Administrador1`
    FOREIGN KEY (`Administrador_idAdministrador`)
    REFERENCES `ventas`.`Administrador` (`idAdministrador`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `ventas`.`boleto_has_Usuario`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `ventas`.`boleto_has_Usuario` (
  `boleto_idboleto` INT NOT NULL,
  `Usuario_idUsuario` INT NOT NULL,
  `Reporte_idReporte` INT NOT NULL,
  PRIMARY KEY (`boleto_idboleto`, `Usuario_idUsuario`, `Reporte_idReporte`),
  INDEX `fk_boleto_has_Usuario_Usuario1_idx` (`Usuario_idUsuario` ASC) VISIBLE,
  INDEX `fk_boleto_has_Usuario_boleto1_idx` (`boleto_idboleto` ASC) VISIBLE,
  INDEX `fk_boleto_has_Usuario_Reporte1_idx` (`Reporte_idReporte` ASC) VISIBLE,
  CONSTRAINT `fk_boleto_has_Usuario_boleto1`
    FOREIGN KEY (`boleto_idboleto`)
    REFERENCES `ventas`.`boleto` (`idboleto`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_boleto_has_Usuario_Usuario1`
    FOREIGN KEY (`Usuario_idUsuario`)
    REFERENCES `ventas`.`Usuario` (`idUsuario`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_boleto_has_Usuario_Reporte1`
    FOREIGN KEY (`Reporte_idReporte`)
    REFERENCES `ventas`.`Reporte` (`idReporte`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `ventas`.`evento`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `ventas`.`evento` (
  `idevento` INT NOT NULL,
  `nombre_evento` VARCHAR(45) NULL,
  `lugar_evento` VARCHAR(45) NULL,
  `fecha_evento` DATETIME NULL,
  `Venue_idVenue` INT NOT NULL,
  `Administrador_idAdministrador` INT NOT NULL,
  PRIMARY KEY (`idevento`, `Venue_idVenue`, `Administrador_idAdministrador`),
  INDEX `fk_evento_Venue1_idx` (`Venue_idVenue` ASC) VISIBLE,
  INDEX `fk_evento_Administrador1_idx` (`Administrador_idAdministrador` ASC) VISIBLE,
  CONSTRAINT `fk_evento_Venue1`
    FOREIGN KEY (`Venue_idVenue`)
    REFERENCES `ventas`.`Venue` (`idVenue`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_evento_Administrador1`
    FOREIGN KEY (`Administrador_idAdministrador`)
    REFERENCES `ventas`.`Administrador` (`idAdministrador`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `ventas`.`evento_has_Usuario`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `ventas`.`evento_has_Usuario` (
  `evento_idevento` INT NOT NULL,
  `Usuario_idUsuario` INT NOT NULL,
  PRIMARY KEY (`evento_idevento`, `Usuario_idUsuario`),
  INDEX `fk_evento_has_Usuario_Usuario1_idx` (`Usuario_idUsuario` ASC) VISIBLE,
  INDEX `fk_evento_has_Usuario_evento1_idx` (`evento_idevento` ASC) VISIBLE,
  CONSTRAINT `fk_evento_has_Usuario_evento1`
    FOREIGN KEY (`evento_idevento`)
    REFERENCES `ventas`.`evento` (`idevento`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_evento_has_Usuario_Usuario1`
    FOREIGN KEY (`Usuario_idUsuario`)
    REFERENCES `ventas`.`Usuario` (`idUsuario`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `ventas`.`evento_has_boleto`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `ventas`.`evento_has_boleto` (
  `evento_idevento` INT NOT NULL,
  `boleto_idboleto` INT NOT NULL,
  PRIMARY KEY (`evento_idevento`, `boleto_idboleto`),
  INDEX `fk_evento_has_boleto_boleto1_idx` (`boleto_idboleto` ASC) VISIBLE,
  INDEX `fk_evento_has_boleto_evento1_idx` (`evento_idevento` ASC) VISIBLE,
  CONSTRAINT `fk_evento_has_boleto_evento1`
    FOREIGN KEY (`evento_idevento`)
    REFERENCES `ventas`.`evento` (`idevento`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_evento_has_boleto_boleto1`
    FOREIGN KEY (`boleto_idboleto`)
    REFERENCES `ventas`.`boleto` (`idboleto`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
