FROM redis:7-alpine

ARG DBIVTOOL_REDIS_PASSWORD
ENV DBIVTOOL_REDIS_PASSWORD $DBIVTOOL_REDIS_PASSWORD

COPY conf/redis.conf /usr/local/etc/redis/redis.conf
RUN sed -i "s/{DBIVTOOL_REDIS_PASSWORD}/$DBIVTOOL_REDIS_PASSWORD/g" /usr/local/etc/redis/redis.conf
CMD ["redis-server", "/usr/local/etc/redis/redis.conf"]
