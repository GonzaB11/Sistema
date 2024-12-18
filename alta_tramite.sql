create or replace function alta_tramite(ingreso_id_cola_atencion int, 
		ingreso_tipo_tramite char(10), ingreso_descripcion text) returns int as $$

declare
	id_tramite_alta int;
	id_cliente_alta int;
begin
	if ingreso_tipo_tramite not in ('consulta', 'reclamo') then
		insert into error (
			operacion,  
			id_cola_atencion, 
			tipo_tramite, 
			id_tramite, 
			estado_cierre_tramite, 
			f_error, 
			motivo
			)
		values ( 
			'Alta tramite',
			 ingreso_id_cola_atencion, 
			 ingreso_tipo_tramite,   
			 current_timestamp, 
			 'Tipo de trámite no válido'
			 );
		return -1;
	end if;

	select id_cliente into id_cliente_alta
	from cola_atencion
	where id_cola_atencion = ingreso_id_cola_atencion
	and estado != 'espera';
	if not found then
		insert into error ( 
			operacion,  
			id_cola_atencion, 
			tipo_tramite, 
			f_error, 
			motivo
			)
		values ( 
			'Alta tramite',  
			ingreso_id_cola_atencion, 
			ingreso_tipo_tramite, 
			current_timestamp, 
			'El id de cola de atención no es válido'
			);
		return -1;
	end if;

	insert into tramite( 
		id_cliente, 
		id_cola_atencion,
		tipo_tramite, 
		f_inicio_gestion, 
		descripcion,  
		estado)
	values( 
		id_cliente_alta, 
		ingreso_id_cola_atencion, 
		ingreso_tipo_tramite, 
		current_timestamp, 
		ingreso_descripcion,  
		'iniciado'
		)
		returning id_tramite into id_tramite_alta;
		
	return id_tramite_alta;

end;
$$ language plpgsql;
