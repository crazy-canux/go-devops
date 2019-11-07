/*
# httpd/apach2

copy html in  Dockerfile:

    FROM httpd:tag
    COPY ./web/ /usr/local/apache2/htdocs/

launch your apache with specify port:

    $ docker run -dit --name admin-gui -p 80:80 -p 443:443 myweb:tag

# haproxy

*/
package http
