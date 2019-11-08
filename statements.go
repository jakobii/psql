package psql

var sqlGetColumns = `
DECLARE @Schema nvarchar(128) = N'%s' 
DECLARE @Table nvarchar(128) = N'%s' 

SELECT
	t.COLUMN_NAME AS [Name],
	t.ORDINAL_POSITION AS [Position],
	t.DATA_TYPE [Type],
	iif(t.IS_NULLABLE = 'YES',cast(1 as bit), cast(0 as bit)) AS [Nullable],
	iif(c.CONSTRAINT_TYPE = 'PRIMARY KEY',cast(1 as bit), cast(0 as bit)) AS [Key],
	iif(c.CONSTRAINT_TYPE = 'UNIQUE',cast(1 as bit), cast(0 as bit)) AS [Unique],
	t.COLUMN_DEFAULT AS [Default],
	t.CHARACTER_SET_NAME AS [CharacterSet],
	COALESCE(t.NUMERIC_PRECISION,t.DATETIME_PRECISION,0) AS [Percision],
	COALESCE(t.NUMERIC_SCALE,0) AS [Scale],
	COALESCE(t.CHARACTER_MAXIMUM_LENGTH,0) AS [Length]

FROM [INFORMATION_SCHEMA].[COLUMNS] AS t

LEFT JOIN [INFORMATION_SCHEMA].[KEY_COLUMN_USAGE] AS u
    on  u.COLUMN_NAME = t.COLUMN_NAME
    and u.TABLE_NAME = t.TABLE_NAME
	and u.TABLE_SCHEMA = t.TABLE_SCHEMA

LEFT JOIN [INFORMATION_SCHEMA].[TABLE_CONSTRAINTS] AS c
	on c.CONSTRAINT_NAME = u.CONSTRAINT_NAME

WHERE
    t.[TABLE_NAME] = @Table
    AND t.[TABLE_SCHEMA] = @Schema

ORDER BY t.ORDINAL_POSITION
`
