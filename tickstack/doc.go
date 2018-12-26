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

use localhost:

    $ docker run -d --name=kapacitor -p 9092:9092 \
    --net=container:influxdb
    kapacitor:tag

use configuration on node:

    $ docker run -d --name=kapacitor -p 9092:9092 \
    -v kapacitor-storage/kapacitor.conf:/etc/kapacitor/kapacitor.conf:ro \
    kapacitor:tag

use docker network:

    [[influxdb]]
    urls = ["http://influxdb:8086"]

    # set env
    $ docker run -d --name=kapacitor -p 9092:9092 \
    --net=influxdb \
    -h kapacitor \
    # This option and kapacitor config are the same.
    -e KAPACITOR_INFLUXDB_0_URLS_0=http://influxdb:8086 \
    kapacitor:tag

# Telegraf

telegraf default get cpu/mem from docker not host.

launch with volume:

    $ docker run -d --name=telegraf \
    -v telegraf-storage:/var/lib/telegraf
    telegraf:tag

use localhost:

    $ docker run -d --name=telegraf \
    --net=container:influxdb
    telegraf:tag

use configuration on node:

    $ docker run -d --name=telegraf \
    -v telegraf-storage/telegraf.conf:/etc/telegraf/telegraf.conf:ro \
    telegraf:tag

use docker network:

    [[outputs.influxdb]]
    urls = ["http://influxdb:8086"]

    $ docker run -d --name=telegraf \
    --net=influxdb \
    -v telegraf-storage/telegraf.conf:/etc/telegraf/telegraf.conf:ro \
    telegraf:tag

monitoring host filesystem mount /proc:

    $ docker run -d --name=telegraf --restart=always \
    --net=influxdb --add-host="influxdb:10.103.0.1" \
    -e HOST_PROC=/proc -e HOST_SYS=/sys -e HOST_ETC=/etc -e HOST_VAR=/var \
    -v /proc:/proc:ro -v /sys:/sys -v /etc:/etc -v /var:/var \
    -v telegraf-storage/telegraf.conf:/etc/telegraf/telegraf.conf:ro \
    telegraf:tag

monitoring other container with docker plugin:

    [[inputs.docker]]
    endpoint = "unix:///var/run/docker.sock"

    $ docker run -d --name=telegraf \
    --net=influxdb \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -v telegraf-storage/telegraf.conf:/etc/telegraf/telegraf.conf:ro \
    telegraf:tag

*/
package tickstack
