create or replace function rendimiento_operadore_desistido() returns trigger as $$
declare
    duracion interval;
    r_id_operadore int;
    fecha_fin date;
begin
    duracion := new.f_fin_atencion - new.f_inicio_atencion;
    r_id_operadore := new.id_operadore;
    fecha_fin := new.f_fin_atencion;

	if not exists (
		select 1 from  rendimiento_operadore 
		where id_operadore = r_id_operadore)then
		insert into rendimiento_operadore (
			id_operadore, 
			fecha_atencion,
			duracion_total_atenciones,
			cantidad_total_atenciones,
			duracion_promedio_total_atenciones,
			duracion_atenciones_desistidas, 
			cantidad_atenciones_desistidas,
			duracion_promedio_atenciones_desistidas
			)
		values(
			r_id_operadore,
			fecha_fin,
			duracion,
		 	1,
			duracion,
			duracion,
			1,
			duracion
	  		);	
	else
		update rendimiento_operadore 
    	set
        	duracion_total_atenciones = duracion_total_atenciones + duracion,
        	cantidad_total_atenciones = cantidad_total_atenciones + 1,
        	duracion_promedio_total_atenciones = duracion_total_atenciones / cantidad_total_atenciones,
        	duracion_atenciones_desistidas = duracion_atenciones_desistidas + duracion,
        	cantidad_atenciones_desistidas = cantidad_atenciones_desistidas + 1,
        	duracion_promedio_atenciones_desistidas = duracion_atenciones_desistidas / cantidad_atenciones_desistidas
    	where
        	id_operadore = r_id_operadore
       		and fecha_atencion = fecha_fin;
    end if;
    return new;
end;
$$ language plpgsql;

create trigger rendimiento_operadore_desistido_trg
after update on cola_atencion
for each row
when (new.estado = 'desistido' and new.id_operadore is not null)
execute function rendimiento_operadore_desistido();



