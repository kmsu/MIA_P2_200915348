package main

import (
	Comandos "MIA_P2_200915348/Comandos"
	DFPM "MIA_P2_200915348/Comandos/AdminPermisosPaths" //DFPM -> Directory, File, Permision Management (Administrador de carpetas, archivos y permisos)
	DM "MIA_P2_200915348/Comandos/AdministradorDiscos"  //DM -> DiskManagement (Administrador de discos)
	FS "MIA_P2_200915348/Comandos/SistemaDeArchivos"    //FS -> FileSystem (sistema de archivos)
	US "MIA_P2_200915348/Comandos/Users"                //US -> UserS
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Analizar2(entrada string) {
	parametros := strings.Split(entrada, " -")

	//--------------------------------- ADMINISTRADOR DE DISCOS ------------------------------------------------
	if strings.ToLower(parametros[0]) == "mkdisk" {
		//MKDISK
		//crea un archivo binario que simula un disco con su respectivo MBR
		if len(parametros) > 1 {
			DM.Mkdisk(parametros)
		} else {
			fmt.Println("MKDISK ERROR: parametros no encontrados")
		}

	} else if strings.ToLower(parametros[0]) == "rmdisk" {
		//RMDISK
		if len(parametros) > 1 {
			DM.Rmdisk(parametros)
		} else {
			fmt.Println("RMDISK ERROR: parametros no encontrados")
		}

	} else if strings.ToLower(parametros[0]) == "fdisk" {
		//FDISK
		if len(parametros) > 1 {
			DM.Fdisk(parametros)
		} else {
			fmt.Println("FDISK ERROR: parametros no encontrados")
		}

	} else if strings.ToLower(parametros[0]) == "mount" {
		//MOUNT
		if len(parametros) > 1 {
			DM.Mount(parametros)
		} else {
			fmt.Println("MOUNT ERROR: parametros no encontrados")
		}

	} else if strings.ToLower(parametros[0]) == "unmount" {
		//UNMOUNT
		if len(parametros) > 1 {
			DM.Unmount(parametros)
		} else {
			fmt.Println("UNMOUNT ERROR: parametros no encontrados")
		}

		//--------------------------------- SISTEMA DE ARCHIVOS ----------------------------------------------------
	} else if strings.ToLower(parametros[0]) == "mkfs" {
		//MKFS
		if len(parametros) > 1 {
			FS.Mkfs(parametros)
		} else {
			fmt.Println("MKFS ERROR: parametros no encontrados")
		}

		//--------------------------------------- USERS ------------------------------------------------------------
	} else if strings.ToLower(parametros[0]) == "login" {
		//LOGIN
		if len(parametros) > 1 {
			US.Login(parametros)
		} else {
			fmt.Println("LOGIN ERROR: parametros no encontrados")
		}

	} else if strings.ToLower(parametros[0]) == "logout" {
		//LOGOUT
		if len(parametros) == 1 {
			US.Logout()
		} else {
			fmt.Println("LOGOUT ERROR: Este comando no requiere parametros")
		}

	} else if strings.ToLower(parametros[0]) == "mkgrp" {
		//MKGRP
		if len(parametros) > 1 {
			US.Mkgrp(parametros)
		} else {
			fmt.Println("MKGRP ERROR: parametros no encontrados")
		}

	} else if strings.ToLower(parametros[0]) == "rmgrp" {
		//RMGRP
		if len(parametros) > 1 {
			US.Rmgrp(parametros)
		} else {
			fmt.Println("RMGRP ERROR: parametros no encontrados")
		}

	} else if strings.ToLower(parametros[0]) == "mkusr" {
		//MKUSR
		if len(parametros) > 1 {
			US.Mkusr(parametros)
		} else {
			fmt.Println("MKUSR ERROR: parametros no encontrados")
		}

	} else if strings.ToLower(parametros[0]) == "rmusr" {
		//RMUSR
		if len(parametros) > 1 {
			US.Rmusr(parametros)
		} else {
			fmt.Println("RMUSR ERROR: parametros no encontrados")
		}

		// ------------------ ADMINISTRACION DE CARPETAS, ARCHIVOS Y PERMISOS --------------------------------------
	} else if strings.ToLower(parametros[0]) == "cat" {
		//CAT
		if len(parametros) > 1 {
			DFPM.Cat(parametros)
		} else {
			fmt.Println("CAT ERROR: parametros no encontrados")
		}

	} else if strings.ToLower(parametros[0]) == "mkdir" {
		//MKDIR
		if len(parametros) > 1 {
			DFPM.Mkdir(parametros)
		} else {
			fmt.Println("MKDIR ERROR: parametros no encontrados")
		}

	} else if strings.ToLower(parametros[0]) == "mkfile" {
		//MKDIR
		if len(parametros) > 1 {
			DFPM.Mkfile(parametros)
		} else {
			fmt.Println("MKDIR ERROR: parametros no encontrados")
		}

		//--------------------------------------- OTROS ------------------------------------------------------------
	} else if strings.ToLower(parametros[0]) == "rep" {
		//REP
		if len(parametros) > 1 {
			Comandos.Rep(parametros)
		} else {
			fmt.Println("REP ERROR: parametros no encontrados")
		}

	} else if strings.ToLower(parametros[0]) == "pause" {
		fmt.Println("Presione enter para continuar...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')

	} else if strings.ToLower(parametros[0]) == "exit" {
		fmt.Println("Salida exitosa")
		os.Exit(0)

	} else if strings.ToLower(parametros[0]) == "" {
		//para agregar lineas con cada enter sin tomarlo como error
	} else {
		fmt.Println("Comando no reconocible")
	}

}
