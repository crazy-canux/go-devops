/*
# Chronograf

launch with volume:

    $ docker run -d -p 8888:8888 \
    -v chronograf-storage:/var/lib/chronograf \
    chronograf:tag

# Influxdb

launch with volume:

    $ docker run -d --name=influxdb -p 8086:8086 \
    -v influxdb-storage:/var/lib/influxdb \
    influxdb:tag

use docker network:

    $ docker network create influxdb

    $ docker run -d --name=influxdb -p 8086:8086 \
    --net=influxdb
    influxdb:tag

# Kapacitor

launch with volume:

    $ docker run -d --name=kapacitor -p 9092:9092 \
    -v kapacitor:/var/lib/kapacitor \
    kapacitor:tag

use configuration on node:

    $ docker run -d --name=kapacitor -p 9092:9092 \
    -v kapacitor-storage/kapacitor.conf:/etc/kapacitor/kapacitor.conf:ro \
    kapacitor:tag

use docker network:

    $ docker run -d --name=kapacitor -p 9092:9092 \
    --net=container:influxdb
    kapacitor:tag

    # set env
    $ docker run -d --name=kapacitor -p 9092:9092 \
    --net=influxdb \
    -h kapacitor \
    -e KAPACITOR_INFLUXDB_0_URLS_0=http://influxdb:8086 \
    kapacitor:tag

# Telegraf

launch with volume:

    $ docker run -d --name=telegraf -p 8083:8083 \
    -v telegraf-storage:/var/lib/telegraf
    telegraf:tag

use configuration on node:

    $ docker run -d --name=telegraf -p 8083:8083 \
    -v telegraf-storage/telegraf.conf:/etc/telegraf/telegraf.conf:ro \
    telegraf:tag

use docker network:

    $ docker run -d --name=telegraf -p 8083:8083 \
    --net=container:influxdb
    telegraf:tag

    # set env: urls = ["http://influxdb:8086"]
    $ docker run -d --name=telegraf -p 8083:8083 \
    --net=influxdb \
    -v telegraf-storage/telegraf.conf:/etc/telegraf/telegraf.conf:ro \
    telegraf:tag

*/
package tickstack
