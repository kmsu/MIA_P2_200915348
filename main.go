package main

import (
	Comandos "MIA_P2_200915348/Comandos"
	DFPM "MIA_P2_200915348/Comandos/AdminPermisosPaths" //DFPM -> Directory, File, Permision Management (Administrador de carpetas, archivos y permisos)
	DM "MIA_P2_200915348/Comandos/AdministradorDiscos"  //DM -> DiskManagement (Administrador de discos)
	FS "MIA_P2_200915348/Comandos/SistemaDeArchivos"    //FS -> FileSystem (sistema de archivos)
	US "MIA_P2_200915348/Comandos/Users"                //US -> UserS
	"MIA_P2_200915348/Herramientas"
	"MIA_P2_200915348/HerramientasInodos"
	"MIA_P2_200915348/Structs"
	"bufio"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/cors"
)

type Entrada struct {
	Text string `json:"text"`
}

type Login struct {
	User string `json:"usuario"`
	Pass string `json:"password"`
	Id   string `json:"id"`
}

type StatusResponse struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func main() {
	//metodos de uso
	http.HandleFunc("/analizar", getCadenaAnalizar)
	http.HandleFunc("/discos", getDiscos)
	http.HandleFunc("/particiones", getParticiones)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/explorer", getContenido)
	http.HandleFunc("/contenido", getContenidoR)
	http.HandleFunc("/file", getFile)
	http.HandleFunc("/back", getBack)
	http.HandleFunc("/reportes", getReportes)
	http.HandleFunc("/descargar", getDescarga)
	http.HandleFunc("/limpiar", getLimpieza)

	// Configurar CORS con opciones predeterminadas
	c := cors.Default()

	// Configurar el manejador HTTP con CORS
	handler := c.Handler(http.DefaultServeMux)

	// Iniciar el servidor en el puerto 8080
	fmt.Println("Servidor escuchando en http://localhost:8080")
	http.ListenAndServe(":8080", handler)
}

// Ejecutar comandos
func getCadenaAnalizar(w http.ResponseWriter, r *http.Request) {
	// Configurar la cabecera de respuesta
	w.Header().Set("Content-Type", "application/json")

	var status StatusResponse
	//verificar que sea un post
	if r.Method == http.MethodPost {
		var entrada Entrada
		if err := json.NewDecoder(r.Body).Decode(&entrada); err != nil {
			http.Error(w, "Error al decodificar JSON", http.StatusBadRequest)
			status = StatusResponse{Message: "Error al decodificar JSON", Type: "unsucces"}
			json.NewEncoder(w).Encode(status)
			return
		}

		//creo un lector de bufer para el archivo
		lector := bufio.NewScanner(strings.NewReader(entrada.Text))
		//leer el archivo linea por linea
		for lector.Scan() {
			//Divido por # para ignorar todo lo que este a la derecha del mismo
			linea := strings.Split(lector.Text(), "#") //lector.Text() retorna la linea leida
			if len(linea[0]) != 0 {
				fmt.Println("\n*********************************************************************************************")
				fmt.Println("Linea en ejecucion: ", linea[0])
				analizar(linea[0])
			}
		}

		//fmt.Println("Cadena recibida ", entrada.Text)
		w.WriteHeader(http.StatusOK)

		status = StatusResponse{Message: "recibido correctamente", Type: "succes"}
		json.NewEncoder(w).Encode(status)

	} else {
		//http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		status = StatusResponse{Message: "Metodo no permitido", Type: "unsucces"}
		json.NewEncoder(w).Encode(status)
	}
}

// Funcion de analisis de comandos
func analizar(entrada string) {
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

	} else if strings.ToLower(parametros[0]) == "exit" {
		fmt.Println("Salida exitosa")
		os.Exit(0)

	} else if strings.ToLower(parametros[0]) == "" {
		//para agregar lineas con cada enter sin tomarlo como error
	} else {
		fmt.Println("Comando no reconocible")
	}
}

func getDiscos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("discos")
	// Configurar la cabecera de respuesta
	w.Header().Set("Content-Type", "application/json")

	directorio := "./MIA/P1"
	//lista de discos encontrados
	var discos []string

	//recorrer el directorio y buscar discos
	err := filepath.Walk(directorio, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			discos = append(discos, info.Name())
		}
		return nil
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al buscar archivos: %s", err), http.StatusInternalServerError)
	}

	respuestaJSON, err := json.Marshal(discos)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al serializar datos a JSON: %s", err), http.StatusInternalServerError)
		return
	}
	w.Write(respuestaJSON)
}

func getParticiones(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Particiones")
	// Configurar la cabecera de respuesta
	w.Header().Set("Content-Type", "application/json")

	var entrada string
	if err := json.NewDecoder(r.Body).Decode(&entrada); err != nil {
		http.Error(w, "Error al decodificar JSON", http.StatusBadRequest)
		return
	}

	filepath := "./MIA/P1/" + entrada

	disco, err := Herramientas.OpenFile(filepath)
	if err != nil {
		fmt.Println("MOUNT Error: No se pudo leer el disco")
		return
	}

	//Se crea un mbr para cargar el mbr del disco
	var mbr Structs.MBR
	//Guardo el mbr leido
	if err := Herramientas.ReadObject(disco, &mbr, 0); err != nil {
		return
	}

	// cerrar el archivo del disco
	defer disco.Close()

	//lista de discos encontrados
	var particiones []string

	for i := 0; i < 4; i++ {
		estado := string(mbr.Partitions[i].Status[:])
		if estado == "A" {
			particiones = append(particiones, string(mbr.Partitions[i].Id[:]))
		}
	}

	fmt.Println("Montadas ", particiones)
	respuestaJSON, err := json.Marshal(particiones)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al serializar datos a JSON: %s", err), http.StatusInternalServerError)
		return
	}
	w.Write(respuestaJSON)
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login")
	// Configurar la cabecera de respuesta
	w.Header().Set("Content-Type", "application/json")

	var entrada Login
	if err := json.NewDecoder(r.Body).Decode(&entrada); err != nil {
		http.Error(w, "Error al decodificar JSON", http.StatusBadRequest)
		return
	}

	fmt.Println("User ", entrada.User)
	fmt.Println("Pass ", entrada.Pass)
	fmt.Println("Id ", entrada.Id)
	//Construir cadena para ejecutar el comando login
	//login -user=root -pass=123 -id=A148

	logear := [4]string{"login", "user=" + entrada.User, "pass=" + entrada.Pass, "id=" + entrada.Id}
	fmt.Println("logear ", logear)

	respuestaJSON, err := json.Marshal(US.Login(logear[:]))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al serializar datos a JSON: %s", err), http.StatusInternalServerError)
		return
	}
	w.Write(respuestaJSON)
}

func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("logout")
	// Configurar la cabecera de respuesta
	w.Header().Set("Content-Type", "application/json")

	respuestaJSON, err := json.Marshal(US.Logout())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al serializar datos a JSON: %s", err), http.StatusInternalServerError)
		return
	}
	w.Write(respuestaJSON)
}

// variables para manejar los inodos (carpetas y archivos)
var idActual int32
var initSuperBloque int64

func getContenido(w http.ResponseWriter, r *http.Request) {
	fmt.Println("contenido")
	// Configurar la cabecera de respuesta
	w.Header().Set("Content-Type", "application/json")

	id := Structs.UsuarioActual.Id
	disk := id[0:1]
	//abrir disco a reportar
	carpeta := "./MIA/P1/" //Ruta (carpeta donde se leera el disco)
	extension := ".dsk"
	rutaDisco := carpeta + disk + extension

	disco, err := Herramientas.OpenFile(rutaDisco)
	if err != nil {
		return
	}

	//Se crea un mbr para cargar el mbr del disco
	var mbr Structs.MBR
	//Guardo el mbr leido
	if err := Herramientas.ReadObject(disco, &mbr, 0); err != nil {
		return
	}

	// cerrar el archivo del disco
	defer disco.Close()

	for i := 0; i < 4; i++ {
		identificador := Structs.GetId(string(mbr.Partitions[i].Id[:]))
		if identificador == id {
			initSuperBloque = int64(mbr.Partitions[i].Start)
			break //para que ya no siga recorriendo si ya encontro la particion independientemente si se pudo o no reducir
		}
	}

	var superBloque Structs.Superblock
	Herramientas.ReadObject(disco, &superBloque, initSuperBloque)

	var Inode0 Structs.Inode
	Herramientas.ReadObject(disco, &Inode0, int64(superBloque.S_inode_start))

	//establezco valores de id (como es raiz ambos seran 0)
	idActual = 0

	//lista de discos encontrados
	var contenido []string

	var folderBlock Structs.Folderblock
	for i := 0; i < 12; i++ {
		idBloque := Inode0.I_block[i]
		if idBloque != -1 {
			Herramientas.ReadObject(disco, &folderBlock, int64(superBloque.S_block_start+(idBloque*int32(binary.Size(Structs.Folderblock{})))))
			//Recorrer el bloque actual buscando la carpeta/archivo en la raiz
			for j := 2; j < 4; j++ {
				//apuntador es el apuntador del bloque al inodo (carpeta/archivo), si existe es distinto a -1
				apuntador := folderBlock.B_content[j].B_inodo
				if apuntador != -1 {
					pathActual := Structs.GetB_name(string(folderBlock.B_content[j].B_name[:]))
					contenido = append(contenido, pathActual)
				}
			}
		}
	}

	respuestaJSON, err := json.Marshal(contenido)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al serializar datos a JSON: %s", err), http.StatusInternalServerError)
		return
	}
	w.Write(respuestaJSON)
}

// para manejar el anterior
var listaAnterior []int32

func getContenidoR(w http.ResponseWriter, r *http.Request) {
	fmt.Println("contenido 2")
	// Configurar la cabecera de respuesta
	w.Header().Set("Content-Type", "application/json")

	var entrada string
	if err := json.NewDecoder(r.Body).Decode(&entrada); err != nil {
		http.Error(w, "Error al decodificar JSON", http.StatusBadRequest)
		return
	}

	fmt.Println("Buscar ", entrada)

	id := Structs.UsuarioActual.Id
	disk := id[0:1]
	//abrir disco a reportar
	carpeta := "./MIA/P1/" //Ruta (carpeta donde se leera el disco)
	extension := ".dsk"
	rutaDisco := carpeta + disk + extension

	disco, err := Herramientas.OpenFile(rutaDisco)
	if err != nil {
		return
	}
	// Close bin file
	defer disco.Close()

	var superBloque Structs.Superblock
	Herramientas.ReadObject(disco, &superBloque, initSuperBloque)

	//agrego el actual a la pila de anteriores (este sera el anterior)
	listaAnterior = append(listaAnterior, idActual)
	//busco en el actual
	idActual = HerramientasInodos.BuscarInodo(idActual, "/"+entrada, superBloque, disco)

	//cargo el inodo actual
	var Inode Structs.Inode
	Herramientas.ReadObject(disco, &Inode, int64(superBloque.S_inode_start+(idActual*int32(binary.Size(Structs.Inode{})))))

	//lista de discos encontrados
	var contenido []string

	var folderBlock Structs.Folderblock
	for i := 0; i < 12; i++ {
		idBloque := Inode.I_block[i]
		if idBloque != -1 {
			Herramientas.ReadObject(disco, &folderBlock, int64(superBloque.S_block_start+(idBloque*int32(binary.Size(Structs.Folderblock{})))))
			//Recorrer el bloque actual buscando la carpeta/archivo en la raiz
			for j := 2; j < 4; j++ {
				//apuntador es el apuntador del bloque al inodo (carpeta/archivo), si existe es distinto a -1
				apuntador := folderBlock.B_content[j].B_inodo
				if apuntador != -1 {
					pathActual := Structs.GetB_name(string(folderBlock.B_content[j].B_name[:]))
					contenido = append(contenido, pathActual)
				}
			}
		}
	}

	respuestaJSON, err := json.Marshal(contenido)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al serializar datos a JSON: %s", err), http.StatusInternalServerError)
		return
	}
	w.Write(respuestaJSON)
}

func getFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("file")
	// Configurar la cabecera de respuesta
	w.Header().Set("Content-Type", "application/json")

	var entrada string
	if err := json.NewDecoder(r.Body).Decode(&entrada); err != nil {
		http.Error(w, "Error al decodificar JSON", http.StatusBadRequest)
		return
	}

	fmt.Println("Buscar ", entrada)

	id := Structs.UsuarioActual.Id
	disk := id[0:1]
	//abrir disco a reportar
	carpeta := "./MIA/P1/" //Ruta (carpeta donde se leera el disco)
	extension := ".dsk"
	rutaDisco := carpeta + disk + extension

	disco, err := Herramientas.OpenFile(rutaDisco)
	if err != nil {
		return
	}
	// Close bin file
	defer disco.Close()

	var superBloque Structs.Superblock
	Herramientas.ReadObject(disco, &superBloque, initSuperBloque)

	//agrego el actual a la pila de anteriores (este sera el anterior)
	listaAnterior = append(listaAnterior, idActual)
	//busco en el actual
	idActual = HerramientasInodos.BuscarInodo(idActual, "/"+entrada, superBloque, disco)

	//cargo el inodo actual
	var Inode Structs.Inode
	Herramientas.ReadObject(disco, &Inode, int64(superBloque.S_inode_start+(idActual*int32(binary.Size(Structs.Inode{})))))

	//lista de discos encontrados
	var contenido []string
	var textFile string
	var fileBlock Structs.Fileblock

	for _, idBlock := range Inode.I_block {
		if idBlock != -1 {
			Herramientas.ReadObject(disco, &fileBlock, int64(superBloque.S_block_start+(idBlock*int32(binary.Size(Structs.Fileblock{})))))
			textFile += string(fileBlock.B_content[:])
		}
	}

	contenido = append(contenido, textFile)
	fmt.Println("Contenido archivo ", contenido[0])
	respuestaJSON, err := json.Marshal(contenido)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al serializar datos a JSON: %s", err), http.StatusInternalServerError)
		return
	}
	w.Write(respuestaJSON)
}

func getBack(w http.ResponseWriter, r *http.Request) {
	fmt.Println("back")
	// Configurar la cabecera de respuesta
	w.Header().Set("Content-Type", "application/json")

	id := Structs.UsuarioActual.Id
	disk := id[0:1]
	//abrir disco a reportar
	carpeta := "./MIA/P1/" //Ruta (carpeta donde se leera el disco)
	extension := ".dsk"
	rutaDisco := carpeta + disk + extension

	disco, err := Herramientas.OpenFile(rutaDisco)
	if err != nil {
		return
	}

	var superBloque Structs.Superblock
	Herramientas.ReadObject(disco, &superBloque, initSuperBloque)

	//obtengo el ultimo elemento de la lista en el idActual global
	idActual = listaAnterior[len(listaAnterior)-1]
	//elimino el elemento de la lista
	listaAnterior = listaAnterior[:len(listaAnterior)-1]
	//cargo el inodo actual
	var Inode Structs.Inode
	Herramientas.ReadObject(disco, &Inode, int64(superBloque.S_inode_start+(idActual*int32(binary.Size(Structs.Inode{})))))

	//lista de discos encontrados
	var contenido []string

	var folderBlock Structs.Folderblock
	for i := 0; i < 12; i++ {
		idBloque := Inode.I_block[i]
		if idBloque != -1 {
			Herramientas.ReadObject(disco, &folderBlock, int64(superBloque.S_block_start+(idBloque*int32(binary.Size(Structs.Folderblock{})))))
			//Recorrer el bloque actual buscando la carpeta/archivo en la raiz
			for j := 2; j < 4; j++ {
				//apuntador es el apuntador del bloque al inodo (carpeta/archivo), si existe es distinto a -1
				apuntador := folderBlock.B_content[j].B_inodo
				if apuntador != -1 {
					pathActual := Structs.GetB_name(string(folderBlock.B_content[j].B_name[:]))
					contenido = append(contenido, pathActual)
				}
			}
		}
	}

	respuestaJSON, err := json.Marshal(contenido)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al serializar datos a JSON: %s", err), http.StatusInternalServerError)
		return
	}
	w.Write(respuestaJSON)
}

func getReportes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("reportes")
	// Configurar la cabecera de respuesta
	w.Header().Set("Content-Type", "application/json")

	directorio := "./MIA/Reportes"
	//lista de discos encontrados
	var discos []string

	//recorrer el directorio y buscar discos
	err := filepath.Walk(directorio, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			if strings.HasSuffix(info.Name(), ".png") {
				discos = append(discos, info.Name())
			}
		}
		return nil
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al buscar archivos: %s", err), http.StatusInternalServerError)
	}

	respuestaJSON, err := json.Marshal(discos)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al serializar datos a JSON: %s", err), http.StatusInternalServerError)
		return
	}
	w.Write(respuestaJSON)
}

func getDescarga(w http.ResponseWriter, r *http.Request) {
	fmt.Println("file")
	// Configurar la cabecera de respuesta
	w.Header().Set("Content-Type", "application/json")

	var entrada string
	if err := json.NewDecoder(r.Body).Decode(&entrada); err != nil {
		http.Error(w, "Error al decodificar JSON", http.StatusBadRequest)
		return
	}

	directorio := "./MIA/Reportes/"
	fmt.Println("Buscar ", directorio)

	//abrir el archivo
	file, err := http.Dir(directorio).Open(entrada)
	if err != nil {
		http.Error(w, fmt.Sprintf("No se puede abrir el archivo %s para leer: %v", entrada, err), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	//configurar cabeceras
	// Establecer la cabecera HTTP
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", entrada))
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	w.Header().Set("Cache-Control", "must-revalidate")
	w.Header().Set("Pragma", "public")

	// Copiar el archivo al cuerpo de la respuesta
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al copiar el archivo al cuerpo de la respuesta: %v", err), http.StatusInternalServerError)
		return
	}
}

func getLimpieza(w http.ResponseWriter, r *http.Request) {
	fmt.Println("limpieza")
	// Configurar la cabecera de respuesta
	w.Header().Set("Content-Type", "application/json")
}
