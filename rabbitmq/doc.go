/*

# Rabbitmq

rabbitmq:tag default expose 5672.

rabbitmq:tag-management default expost 5672 and 15672.

storage: /var/lib/rabbitmq

launch with management:

    # http://ip:15672 guest/guest
    $ docker run --hostname mq-node1 --name rabbitmq-node1 -d rabbitmq:tag-management

launch with specify auth for management:

    # http://ip:15672 sandbox/password
    $ docker run -v rabbitmq-storage:/var/lib/rabbitmq --hostname mq-node1 --name rabbitmq-node1 -e RABBITMQ_DEFAULT_USER=sandbox -e RABBITMQ_DEFAULT_PASS=password -d rabbitmq:tag-management

# Rabbitmq-HAProxy

specify erlang cookie:

    $ docker run --hostname mq-master -e RABBITMQ_ERLANG_COOKIE='secret cookie here' -d rabbitmq:tag-management

*/
package rabbitmq
