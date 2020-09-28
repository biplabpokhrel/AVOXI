
/* We need table for maping purpose
	network: field should contain network address and the subnet mask
  autonomous_system_organization: Should contain name of the country

*/
CREATE TABLE network_map_tbl (
    id SERIAL primary key,
    network character varying(250) not null,
    autonomous_system_number character varying(250),
    autonomous_system_organization character varying(250) not null
);

/* Sample data */
insert into network_map_tbl
	(network, autonomous_system_organization)
 values
 ( '36.0.4.0/22','Nepal' ),
 ( '42.104.0.0/18','India' ),
 ('43.225.116.0/22','India')
 