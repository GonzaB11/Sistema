create or replace function ingreso_de_llamado (ingreso_id_cliente int) returns int as $$
declare
	id_cola int;
begin
	if not exists (
		select 1 from cliente 
		where id_cliente = ingreso_id_cliente) then
		insert into error(
			operacion, 
			f_error,
			motivo
			)
		values( 
			'nuevo llamado',
			current_timestamp, 
			'cliente no valido'
			);	
		return -1;
	end if;

	insert into cola_atencion(
		id_cliente, 
		f_inicio_llamado, 
		estado
		)
	values (
		ingreso_id_cliente, 
		current_timestamp,  
		'en espera'
		)
	returning id_cola_atencion into id_cola;
	return id_cola;
end;
$$ language plpgsql;
