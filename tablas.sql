create table cliente (
    id_cliente int,
    nombre text, apellido text,
    dni int,
    fecha_nacimiento date,
    telefono char(12),
    email text
);

create table operadore (
	id_operadore int,
	nombre text, apellido text, 
	dni int, fecha_ingreso date, 
	disponible boolean
);

create table cola_atencion (
	id_cola_atencion serial, 
	id_cliente int, 
	f_inicio_llamado timestamp, 
	id_operadore int,
	f_inicio_atencion timestamp,
	f_fin_atencion timestamp, 
	estado char(15)
);

create table tramite (
	id_tramite serial, 
	id_cliente int, 
	id_cola_atencion int, 
	tipo_tramite char(10), 
	f_inicio_gestion timestamp,
	descripcion text, 
	f_fin_gestion timestamp, 
	respuesta text, 
	estado char(15)
);

create table rendimiento_operadore (
	id_operadore int, 
	fecha_atencion date, 
	duracion_total_atenciones interval, 
	cantidad_total_atenciones int,
	duracion_promedio_total_atenciones interval, 
	duracion_atenciones_finalizadas interval, 
	cantidad_atenciones_finalizadas int,
	duracion_promedio_atenciones_finalizadas interval, 
	duracion_atenciones_desistidas interval, 
	cantidad_atenciones_desistidas int,
	duracion_promedio_atenciones_desistidas interval
);			

create table error (
	id_error serial, 
	operacion char(16), 
	id_cliente int, 
	id_cola_atencion int, 
	tipo_tramite char(10), 
	id_tramite int, 
	estado_cierre_tramite char (15),
	f_error timestamp, 
	motivo varchar(80)
);

create table envio_email (
	id_email int, 
	f_generacion timestamp, 
	email_cliente text, 
	asunto text, 
	cuerpo text, 
	f_envio timestamp, 
	estado char(10)
);

create table datos_de_prueba( 
	id_orden int,
	operacion char (16), 
	id_cliente int, 
	id_cola_atencion int,
	tipo_tramite char (10), 
	descripcion_tramite text, 
	id_tramite int, 
	estado_cierre_tramite char(15),
	respuesta_tramite text
);


