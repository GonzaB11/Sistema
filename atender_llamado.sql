create or replace function atender_llamado() returns boolean as $$
declare
	id_cliente_en_espera int;
	id_cola_atencion_en_espera int;
	id_operadore_disponible int;

begin
	if not exists (
		select 1 from cola_atencion 
		where estado = 'en espera') then
		insert into error ( 
			operacion, 
			f_error,
			motivo)
		values (
			'atencion llamado', 
			current_timestamp,
			'no existe ningune llamado en espera'
			);
		return false;
	end if;

	select id_cliente,id_cola_atencion 
		into id_cliente_en_espera, id_cola_atencion_en_espera 
		from cola_atencion 
		where estado = 'en espera' 
		order by f_inicio_llamado limit 1;

	if not exists (
		select 1 from operadore 
		where disponible = true) then
		insert into error ( 
			operacion, 
			id_cliente, 
			id_cola_atencion, 
			f_error,
			motivo
			)
		values ( 
			'atencion llamado',
			id_cliente_en_espera,
			id_cola_atencion_en_espera,
			current_timestamp, 
			'no existe ningun operadore disponible'
			);
		return false;
	end if;

	select id_operadore
		into id_operadore_disponible
		from operadore
		where disponible = true
		limit 1;

	update operadore
		set disponible = false
		where id_operadore = id_operadore_disponible;

	update cola_atencion
		set id_operadore = id_operadore_disponible, 
			f_inicio_atencion = current_timestamp, 
			estado = 'en linea'
		where id_cola_atencion = id_cola_atencion_en_espera;

	return true;
end;
$$ language plpgsql;
	
