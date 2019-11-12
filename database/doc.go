/*
# Mysql

expose 3306.

launch with specify volume:

    $ docker run --name mymysql -v mysql-storage:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=password -d mysql:tag

create database and user:

    $ docker run --name mymysql -v mysql-storage:/var/lib/mysql -e MYSQL_DATABASE=sandbox -e MYSQL_USER=sandbox MYSQL_PASSWORD=password -d mysql:tag

files in /docker-entrypoint-initdb.d will be executed when container started(.sh,.sql,.sql.gz).

    $ docker run --name mymysql -v mysql-storage:/var/lib/mysql -v dbpatch:/docker-entrypoint-initdb.d -e MYSQL_DATABASE=sandbox -e MYSQL_USER=sandbox -e MYSQL_PASSWORD=password -d mysql:tag

# SqlServer

*/
package database
