create or replace function finalizar_llamado (id_cola_atencion_a_finalizar int) returns boolean as $$  
declare
	id  int;
	_id_operadore int;
begin
	if not exists (
		select 1 from cola_atencion 
		where id_cola_atencion = id_cola_atencion_a_finalizar) then
		insert into error( 
			operacion, 
			f_error,
			motivo)
		values ( 
			'fin llamado',  
			current_timestamp,
			'id cola de atencion no valido'
			);
		return false;
	end if;

	select id_cliente into id 
	from cola_atencion 
	where id_cola_atencion = id_cola_atencion_a_finalizar;
	
	if not exists (
		select 1 from cola_atencion 
		where id_cola_atencion= id_cola_atencion_a_finalizar 
		and estado= 'en linea') then
		insert into error( 
			operacion, 
			id_cliente, 
			id_cola_atencion, 
			f_error,
			motivo
			)
		values (
			'fin llamado',
			id,
			id_cola_atencion_a_finalizar, 
			current_timestamp,
			'el llamado no esta en linea'
			);
		return false;
	end if; 

	select id_operadore into _id_operadore
	from cola_atencion
	where id_cola_atencion = id_cola_atencion_a_finalizar;

	update operadore
	set disponible = true
	where id_operadore = _id_operadore; 

	update cola_atencion 
	set estado= 'finalizado', f_fin_atencion= current_timestamp  
	where id_cola_atencion = id_cola_atencion_a_finalizar;
	return true;
end;
$$ language plpgsql;
