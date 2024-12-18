alter table cliente add constraint cliente_pk primary key (id_cliente);
alter table operadore add constraint operadore_pk primary key (id_operadore);
alter table cola_atencion add constraint cola_atencion_pk primary key (id_cola_atencion);
alter table tramite add constraint tramite_pk primary key (id_tramite);
alter table rendimiento_operadore add constraint rendimiento_operadore_pk primary key (id_operadore, fecha_atencion);
alter table error add constraint error_pk primary key (id_error);
alter table envio_email add constraint envio_email_pk primary key (id_email);
alter table datos_de_prueba add constraint datos_de_prueba_pk primary key (id_orden);

alter table cola_atencion add constraint cola_atencion_cliente_fk foreign key (id_cliente) references cliente (id_cliente);
alter table cola_atencion add constraint cola_atencion_operadore_fk foreign key (id_operadore) references operadore (id_operadore);
alter table tramite add constraint tramite_cliente_fk foreign key (id_cliente) references cliente (id_cliente);
alter table tramite add constraint tramite_cola_atencion_fk foreign key (id_cola_atencion) references cola_atencion (id_cola_atencion);
alter table rendimiento_operadore add constraint rendimiento_operadore_id_fk foreign key (id_operadore) references operadore (id_operadore);

