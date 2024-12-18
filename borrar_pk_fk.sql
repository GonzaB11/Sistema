alter table cola_atencion drop constraint if exists cola_atencion_cliente_fk;
alter table cola_atencion drop constraint if exists cola_atencion_operadore_fk;
alter table tramite drop constraint if exists tramite_cliente_fk;
alter table tramite drop constraint if exists tramite_cola_atencion_fk;
alter table rendimiento_operadore drop constraint if exists rendimiento_operadore_id_fk;




alter table cliente drop constraint if exists cliente_pk;
alter table operadore drop constraint if exists operadore_pk;
alter table cola_atencion drop constraint if exists cola_atencion_pk;
alter table tramite drop constraint if exists tramite_pk;
alter table rendimiento_operadore drop constraint if exists rendimiento_operadore_pk;
alter table error drop constraint if exists error_pk;
alter table envio_email drop constraint if exists envio_email_pk;
alter table datos_de_prueba drop constraint if exists datos_de_prueba_pk;
