create or replace function desistir_llamado (ingreso_id_cola_atencion int) returns boolean as $$  
declare
	id  int;
begin
	if not exists (
		select 1 from cola_atencion 
		where id_cola_atencion= ingreso_id_cola_atencion) then
		insert into error( 
			operacion, 
			f_error,
			motivo
			)
		values (
			'baja llamado',  
			current_timestamp,
			'id cola de atencion no valido'
			);
		return false;
	end if;

	select id_cliente into id 
	from cola_atencion 
	where id_cola_atencion = ingreso_id_cola_atencion;

	if not exists (
		select 1 from cola_atencion 
		where id_cola_atencion= ingreso_id_cola_atencion 
		and estado= 'en espera'
		or estado= 'en linea') then
		insert into error(
			operacion, 
			id_cliente, 
			id_cola_atencion, 
		 	f_error,
		 	motivo
		 	)
		values ( 
			'baja llamado',
			id,
			ingreso_id_cola_atencion, 
			current_timestamp,
			'el llamado no esta en espera ni en linea'
			);
		return false;
	end if; 				

	if exists (
		select 1 from cola_atencion 
		where id_cola_atencion= ingreso_id_cola_atencion 
		and estado = 'en linea') then
		update cola_atencion 
		set f_fin_atencion= current_timestamp 
		where id_cola_atencion = ingreso_id_cola_atencion;
	end if;

	update cola_atencion 
	set estado= 'desistido' 
	where id_cola_atencion = ingreso_id_cola_atencion;
	return true;
end;
$$ language plpgsql;	
