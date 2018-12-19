/*
# Chronograf

launch with volume:

    $ docker run -d -p 8888:8888 \
    -v chronograf-storage:/var/lib/chronograf \
	chronograf:tag

# Influxdb

launch with volume:

    $ docker run -d -name=influxdb -p 8086:8086 \
    -v influxdb-storage:/var/lib/influxdb \
	influxdb:tag

use docker network:

    $ docker network create influxdb

    $ docker run -d --name=influxdb -p 8086:8086 \
    --net=influxdb
    influxdb:tag

# Kapacitor

launch with volume:

    $ docker run -d -name=kapacitor -p 9092:9092 \
    -v kapacitor:/var/lib/kapacitor \
	kapacitor:tag

use configuration on node:

    $ docker run -d -name=kapacitor -p 9092:9092 \
    -v kapacitor-storage/kapacitor.conf:/etc/kapacitor/kapacitor.conf:ro \
    kapacitor:tag

use docker network:

    $ docker run -d -name=kapacitor -p 9092:9092 \
    --net=influxdb \
    -h kapacitor \
	-e KAPACITOR_INFLUXDB_0_URLS_0=http://influxdb:8086 \
	kapacitor:tag

    $ docker run -d -name=kapacitor -p 9092:9092 \
    --net=container:influxdb
    kapacitor:tag

# Telegraf

*/
package tickstack
