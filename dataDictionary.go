package mssql

/*

Description of the object type:
AGGREGATE_FUNCTION
CHECK_CONSTRAINT
CLR_SCALAR_FUNCTION
CLR_STORED_PROCEDURE
CLR_TABLE_VALUED_FUNCTION
CLR_TRIGGER
DEFAULT_CONSTRAINT
EXTENDED_STORED_PROCEDURE
FOREIGN_KEY_CONSTRAINT
INTERNAL_TABLE
PLAN_GUIDE
PRIMARY_KEY_CONSTRAINT
REPLICATION_FILTER_PROCEDURE
RULE
SEQUENCE_OBJECT
SERVICE_QUEUE
SQL_INLINE_TABLE_VALUED_FUNCTION
SQL_SCALAR_FUNCTION
SQL_STORED_PROCEDURE
SQL_TABLE_VALUED_FUNCTION
SQL_TRIGGER
SYNONYM
SYSTEM_TABLE
TABLE_TYPE
UNIQUE_CONSTRAINT
USER_TABLE
VIEW

*/

//Query to get all columns of user tables
var SqlGetColumn = `
SELECT
OBJECT_NAME(c.OBJECT_ID) TableName
,c.name AS ColumnName
,c.column_id
,SCHEMA_NAME(o.schema_id) AS SchemaName
,t.name AS TypeName
,t.user_type_id
,t.is_user_defined
,t.is_assembly_type
,c.max_length
,c.PRECISION
,c.scale
--into tmp_ntext_fields
FROM sys.columns AS c
JOIN sys.types AS t ON c.user_type_id=t.user_type_id
join sys.objects o on c.object_id = o.object_id 
where
--t.name in ('text','ntext')
--and 
o.type = 'U'  -- user table only
ORDER BY TableName;
`
