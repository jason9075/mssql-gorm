# Dokcer

## Pull
```
docker pull mcr.microsoft.com/mssql/server
```

## Run:

```
docker run -e "ACCEPT_EULA=Y" -e "SA_PASSWORD=zaq12345" -e "MSSQL_PID=Express" -p 1433:1433 -d --name mssql001 mcr.microsoft.com/mssql/server
```

## Azure Data Studio:
https://learn.microsoft.com/zh-tw/sql/azure-data-studio/download-azure-data-studio?view=sql-server-ver15

# Golang

## Package
```
go get github.com/denisenkom/go-mssqldb
```
