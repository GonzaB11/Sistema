package main

import(
		"encoding/json"
		"fmt"
		"log"
		"os"
		"database/sql"
		"time"
		_"github.com/lib/pq" 
		"io/ioutil"
		"strconv"
	    bolt "go.etcd.io/bbolt"
)

type Cliente struct{
	IdCliente int 					`json:"id_cliente"`
	Nombre string 					`json:"nombre"`
	Apellido string 				`json:"apellido"`
	Dni int 						`json:"dni"`
	FechaNacimiento string 		 	`json:"fecha_nacimiento"`
	Telefono string 				`json:"telefono"`
	Email string 					`json:"email"`
}

type Operadore struct{
	IdOperadore int 				`json:"id_operadore"`
	Nombre string 					`json:"nombre"`
	Apellido string 				`json:"apellido"`
	Dni int 						`json:"dni"`
	FechaIngreso string 			`json:"fecha_ingreso"`
	Disponible bool 				`json:"disponible"`
}

type ColaAtencion struct{
	IdColaAtencion int 				`json:"id_cola_atencion"`
	IdCliente int 					`json:"id_cliente"`
	FInicioLlamado string		`json:"f_inicio_llamado"`
	IdOperadore int 				`json:"id_operadore"`
	FInicioAtencion string 		`json:"f_inicio_atencion"`
	FFinAtencion string 		`json:"f_fin_atencion"`
	Estado string 				`json:"estado"`
}



type Tramite struct {
	IDTramite int 					`json:"id_tramite"`
	IDCliente int 					`json:"id_cliente"`
	IDColaAtencion int 				`json:"id_cola_atencion"`
	TipoTramite string 				`json:"tipo_tramite"` 
	FInicioGestion string		 	`json:"f_inicio_gestion"`
	Descripcion string 				`json:"descripcion"` 
	FFinGestion string	 			`json:"f_fin_gestion"`
	Respuesta string 				`json:"respuesta"` 
	Estado string 					`json:"estado"`
}


type RendimientoOperadore struct {
	IDOperadore int 								`json:"id_operadore"`
	FechaAtencion string						`json:"fecha_atencion"`
	DuracionTotalAtenciones string 					`json:"duracion_total_atenciones"` 
	CantidadTotalAtenciones int 					`json:"cantidad_total_atenciones"`
	DuracionPromedioTotalAtenciones string 			`json:"duracion_promedio_total_atenciones"`
	DuracionAtencionesFinalizadas string 			`json:"duracion_atenciones_finalizadas"`
	CantidadAtencionesFinalizadas int 				`json:"cantidad_atenciones_finalizadas"`
	DuracionPromedioAtencionesFinalizadas string 	`json:"duracion_promedio_atenciones_finalizadas"`
	DuracionAtencionesDesistidas string 			`json:"duracion_atenciones_desistidas"`
	CantidadAtencionesDesistidas int 				`json:"cantidad_atenciones_desistidas"`
	DuracionPromedioAtencionesDesistidas string 	`json:"duracion_promedio_atenciones_desistidas"`
}


type Error struct {
	IDError int 				`json:"id_error"`
	Operacion string 			`json:"operacion"`
	IDCliente int				`json:"id_cliente"`
	IDColaAtencion int 			`json:"id_cola_atencion"`
	TipoTramite string 			`json:"tipo_tramite"`
	IDTramite int 				`json:"id_tramite"`
	EstadoCierreTramite string  `json:"estado_cierre_tramite"`
	FError string			 	`json:"f_error"`
	Motivo string 				`json:"motivo"` 
}

type EnvioEmail struct {
	IDEmail int 			`json:"id_email"`
	FGeneracion string	`json:"f_generacion"`
	EmailCliente string 	`json:"email_cliente"`
	Asunto string 			`json:"asunto"`
	Cuerpo string 			`json:"cuerpo"`
	FEnvio string 		`json:"f_envio"`
	Estado string 			`json:"estado"` 
}


type DatosDePrueba struct {
	IDOrden int 				`json:"id_orden"`
	Operacion string 			`json:"operacion"` 
	IDCliente int 				`json:"id_cliente"`
	IDColaAtencion int 			`json:"id_cola_atencion"`
	TipoTramite string 			`json:"tipo_tramite"`
	DescripcionTramite string 	`json:"descripcion"`
	IDTramite int 				`json:"id_tramite"`
	EstadoCierreTramite string 	`json:"estado_cierre_tramite"`
	RespuestaTramite string 	`json:"respuesta_tramite"`
}

func crearBaseDatos(){
	db, err:= sql.Open("postgres","user=postgres host=localhost dbname=postgres sslmode=disable")
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	_, err= db.Exec("drop database if exists aguayo_aquino_benitez_ojeda_db1")
	if err != nil{
		log.Fatal(err)
	}

	_, err= db.Exec("create database aguayo_aquino_benitez_ojeda_db1")
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Base de datos 'aguayo_aquino_benitez_ojeda_db1' creada correctamente")
}

func cargarDatos(ruta string, lista interface{}) {
	err := cargarDatosJson(ruta, lista)
	if err != nil {
		log.Fatalf("Error al cargar datos desde el archivo %s: %v", ruta, err)
	}
}

func cargarDatosJson(archivo string, destino interface{}) error {
	file, err := os.Open(archivo)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(destino); err != nil {
		return err
	}
	return nil
}

func insertarClientes(lista []Cliente, db *sql.DB) {
        for _, cliente := range lista {
                fechaNacimiento, err := time.Parse("2006-01-02", cliente.FechaNacimiento)
                if err != nil {
                        log.Fatal(err)
                }
                _, err = db.Exec("insert into cliente (id_cliente, nombre, apellido, dni, fecha_nacimiento, telefono, email) values ($1, $2, $3, $4, $5, $6, $7)",
                        cliente.IdCliente, cliente.Nombre, cliente.Apellido, cliente.Dni, fechaNacimiento, cliente.Telefono, cliente.Email)
                if err != nil {
                        log.Fatal(err)
                }
        }
}

func insertarOperadores(lista [] Operadore, db*sql.DB){
	for _, operadore := range lista {
			fechaIngreso, err := time.Parse("2006-01-02", operadore.FechaIngreso)
			if err != nil {
				log.Fatal(err)
			}
			_, err= db.Exec("insert into operadore (id_operadore, nombre, apellido, dni, fecha_ingreso, disponible) values ($1, $2, $3, $4, $5, $6)",
					operadore.IdOperadore, operadore.Nombre, operadore.Apellido, operadore.Dni, fechaIngreso, operadore.Disponible)
			if err != nil {
				log.Fatal(err)
			}
	 }
}

func insertarDatosDePruebas(lista []DatosDePrueba, db *sql.DB) {
    for _, dato := range lista {
        _, err := db.Exec("insert into  datos_de_prueba (id_orden, operacion, id_cliente, id_cola_atencion, tipo_tramite, descripcion_tramite, id_tramite, estado_cierre_tramite, respuesta_tramite) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
            dato.IDOrden, dato.Operacion, dato.IDCliente, dato.IDColaAtencion, dato.TipoTramite, dato.DescripcionTramite, dato.IDTramite, dato.EstadoCierreTramite, dato.RespuestaTramite)
        if err != nil {
            log.Fatal(err)
        }
    }
}

func leerArchivo(direccion string) string {
	resultado, err:= ioutil.ReadFile(direccion)
	if err != nil {
		log.Fatal(err)
	}
	return string(resultado)
}

func mostrarMenu() {
    fmt.Println(`
    	#### Menu ####
	1- Crear base de datos
	2- Crear tablas
	3- Agregar PKs y FKs
	4- Eliminar PKs y FKs
	5- Cargar Datos
	6- Crear stored procedures y triggers
	7- Iniciar pruebas
	8- Cargar datos en BoltDB
	9- Salir
    	##############
	`)
}


func iniciarPruebas(lista []DatosDePrueba, db *sql.DB) {
	_, err := db.Exec("select procesar_datos_prueba();")
	if err != nil {
	log.Fatalf("error al ejectura la funcion: %d",err)
	}

	fmt.Println("Pruebas procesadas correctamente.")
}

func CreateUpdate(db *bolt.DB, bucketName string, key []byte, val []byte) error {
	// abre transacción de escritura
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	b, _ := tx.CreateBucketIfNotExists([]byte(bucketName))

	err = b.Put(key, val)
	if err != nil {
		return err
	}

	// cierra transacción
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func cargarClientesEnBolt(db *bolt.DB){	
	clientes :=[]Cliente{
		{1,"Ken","Thompson",5153057,"1995-05-05","15-2889-7948","ken@thompson.org"},
		{2,"Dennis","Ritchie",25610126,"1955-04-11","15-7811-5045","dennis@ritchie.org"},
		{3,"Donald","Knuth",9168297,"1984-04-05","15-2780-6005","don@knuth.org"},
	}
		
	for _, Cliente := range clientes {
		data, err := json.Marshal(Cliente)
		if err != nil {
			log.Fatal(err)
		}
		CreateUpdate(db, "cliente", []byte(strconv.Itoa(Cliente.IdCliente)), data)
	}
}
	
	
func cargarOperadoresEnBolt(db *bolt.DB) {
	operadores :=[]Operadore{
		{1,"Wilhelm","Steinitz",5053058,"2018-05-14",true},
		{2,"Emanuel","Lasker",24610127,"2018-12-24",true},
		{3,"Jose Raul","Capablanca",9068298,"2019-11-19",true},
	}
	
	for _, Operadore := range operadores {
		data, err := json.Marshal(Operadore)
		if err != nil {
			log.Fatal(err)
		}
		CreateUpdate(db, "operadore", []byte(strconv.Itoa(Operadore.IdOperadore)), data)
	}
}	
	
func cargarColaAtencionEnBolt(db *bolt.DB) {
	colaAtenciones :=[]ColaAtencion{
		{1,1,"2024-11-25 00:12:49.596638",2,"2024-11-25 00:13:57.249012","","en linea"},
		{2,2,"2024-11-25 00:12:49.59663",1,"2024-11-25 00:13:57.249012","2024-11-25 00:16:35.4719","desistido"},
		{3,3,"2024-11-25 00:07:16.863089",3,"2024-11-25 00:07:16.863089","2024-11-25 00:07:16.863089","finalizado"},
	}
	
	for _, ColaAtencion := range colaAtenciones {
		data, err := json.Marshal(ColaAtencion)
		if err != nil {
			log.Fatal(err)
		}
		CreateUpdate(db, "colaAtencion", []byte(strconv.Itoa(ColaAtencion.IdColaAtencion)), data)
	}
}
	
func cargarTramitesEnBolt(db *bolt.DB) {
	tramites :=[]Tramite{
		{1,1,1,"consulta", "","","","","iniciado"},
		{2,2,2,"reclamo", "","","","","solucionado"},
		{3,3,3,"reclamo", "","","","","solucionado"},
	}
					
	for _, Tramite := range tramites {
		data, err := json.Marshal(Tramite)
		if err != nil {
			log.Fatal(err)
		}
		CreateUpdate(db, "Tramite", []byte(strconv.Itoa(Tramite.IDTramite)), data)
	}
}

func cargarDatosEnBolt(db *bolt.DB) {
	cargarClientesEnBolt(db)
	cargarOperadoresEnBolt(db)
	cargarColaAtencionEnBolt(db)
	cargarTramitesEnBolt(db)
	
	fmt.Println("Los datos fueron cargados en la base de datos NoSQL correctamente.")	
}




func main() {
	var db *sql.DB
	var clientes []Cliente
	var operadores []Operadore
	var datosPrueba []DatosDePrueba

	for {
		mostrarMenu()
					
		var opcion int
		fmt.Print("Elegir una opcion: ")
		_, err := fmt.Scanf("%d", &opcion)
					
		if err != nil {
			fmt.Println (err)
		}
								
		switch opcion {
			case 1: 
				crearBaseDatos()
				db, err = sql.Open("postgres", "user=postgres host=localhost dbname=aguayo_aquino_benitez_ojeda_db1 sslmode=disable")
				if err != nil {
					log.Fatal(err)
				}
				defer db.Close()
			case 2:
				db, err := sql.Open("postgres", "user=postgres host=localhost dbname=aguayo_aquino_benitez_ojeda_db1 sslmode=disable")
				if err != nil {
			 		log.Fatal(err)
			 	}
				defer db.Close()
				_, err = db.Exec(leerArchivo("tablas.sql"))
				if err!= nil {
					log.Fatal(err)
				}
				fmt.Println ("Tablas creadas")	
			case 3:
				db, err := sql.Open("postgres", "user=postgres host=localhost dbname=aguayo_aquino_benitez_ojeda_db1 sslmode=disable")
				if err != nil {
					log.Fatal(err)
				}
				defer db.Close()
				_, err = db.Exec(leerArchivo("pk_fk.sql"))
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Pk y Fk creadas")
			case 4:
				db, err := sql.Open("postgres", "user=postgres host=localhost dbname=aguayo_aquino_benitez_ojeda_db1 sslmode=disable")
				if err != nil {
					log.Fatal(err)
				}
				defer db.Close()
				_, err = db.Exec(leerArchivo("borrar_pk_fk.sql"))
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Se borraron Pk y Fk")
			case 5:
				db, err := sql.Open("postgres", "user=postgres host=localhost dbname=aguayo_aquino_benitez_ojeda_db1 sslmode=disable")
				if err != nil {cargarDatos("datos_de_prueba.json", &datosPrueba)
								insertarDatosDePruebas(datosPrueba,db)
					log.Fatal(err)
				}
				defer db.Close()
						        
				cargarDatos("clientes.json", &clientes)
				insertarClientes(clientes, db)
					
				cargarDatos("operadores.json", &operadores)
				insertarOperadores(operadores, db)

				cargarDatos("datos_de_prueba.json", &datosPrueba)
				insertarDatosDePruebas(datosPrueba,db)
	
				fmt.Println("Datos cargados en las tablas")	
			case 6:
				db, err := sql.Open("postgres", "user=postgres host=localhost dbname=aguayo_aquino_benitez_ojeda_db1 sslmode=disable")
				if err != nil {
					log.Fatal(err)
				}
				defer db.Close()

				_, err = db.Exec(leerArchivo("ingreso_de_llamado.sql"))
				if err != nil {
					log.Fatalf("ingreso llamado %d",err)
				}

				_, err = db.Exec(leerArchivo("atender_llamado.sql"))
				if err != nil {
					log.Fatalf("atender llamado %d",err)
				}

				_, err = db.Exec(leerArchivo("alta_tramite.sql"))
				if err != nil {
					log.Fatalf("alta tramite %d",err)
				}
					
				_, err = db.Exec(leerArchivo("desistir_llamado.sql"))
				if err != nil {
					log.Fatalf("desistir llamado %d",err)
				}

				_, err = db.Exec(leerArchivo("finalizar_llamado.sql"))
				if err != nil {
					log.Fatalf("finalizar llamado %d", err)
				}

				_, err = db.Exec(leerArchivo("rendimiento_operadore_desistido.sql"))
				if err != nil {
					log.Fatalf("rendimiento operadore desistido %d",err)
				}

				_, err = db.Exec(leerArchivo("rendimiento_operadore_finalizado.sql"))
				if err != nil {
					log.Fatalf("rendimiento operadore finalizado %d",err)
				}
				
				_, err = db.Exec(leerArchivo("procesar_datos_prueba.sql"))
				if err != nil {
					log.Fatalf("procesar datos finalizado %d",err)
				}				

				fmt.Println("Stored procedures y triggers creados")
			case 7:
				db, err := sql.Open("postgres", "user=postgres host=localhost dbname=aguayo_aquino_benitez_ojeda_db1 sslmode=disable")
				if err !=  nil {
					log.Fatal(err)
				} 
				defer db.Close()

				iniciarPruebas(datosPrueba,db)
			case 8:
				db, err := bolt.Open("aguayo_aquino_benitez_ojeda_NoSQL.db", 0600, nil)
				if err != nil{
					log.Fatal(err)
				}
				defer db.Close()
				
				cargarDatosEnBolt(db)											
			case 9:
				fmt.Println("Saliendo")
				os.Exit(0)				
		    }
        }    
}



