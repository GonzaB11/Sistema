create or replace function procesar_datos_prueba()
returns void as $$
declare
    dato record;
begin
    for dato in select * from datos_de_prueba loop
        if dato.operacion = 'nuevo llamado' then
            perform ingreso_de_llamado(dato.id_cliente);
        elsif dato.operacion = 'atencion llamado' then
            perform atender_llamado();
        elsif dato.operacion = 'baja llamado' then
            perform desistir_llamado(dato.id_cola_atencion);
        elsif dato.operacion = 'alta tramite' then
            perform alta_tramite(dato.id_cola_atencion, dato.tipo_tramite, dato.descripcion_tramite);
        elsif dato.operacion = 'fin llamado' then
            perform finalizar_llamado(dato.id_cola_atencion);
        end if;
    end loop;
end;
$$ language plpgsql;
