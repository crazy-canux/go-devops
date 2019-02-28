/*
# Redis

expose 6379 and storage in /data.

launch with specify volume:

    $ docker run --name myredis -v redis-storage:/data -d redis redis-server --appendonly yes

# Mongo

 */

package nosql
