/*
# Grafana

launch:

    docker run -d --name=grafana -p 3000:3000 grafana/grafana:tag

launch with volume:

    $ docker volume create grafana-storage

    $ docker run -d --name=grafana -p 3000:3000 \
	-v grafana-storage:/var/lib/grafana \
	grafana/grafana:tag
*/
package grafana
