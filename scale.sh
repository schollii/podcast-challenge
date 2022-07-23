#!/usr/bin/env bash

# ATTENTION: script assumes $USER is in docker group

replica_count=$1
docker-compose scale podcast="$replica_count"

# doing head requests comes pretty close to perfect for making
# nginx refresh it's DNS cache (sufficient on scale down)
count=1
while ((count < 100)); do
  curl -I -m 0.1 -s http://localhost:8081/health | grep HTTP
  ((count += 1))
done

# this is necessary on scale up to ensure all new pods are found
# on scale down, it causes a blockage for half a minute then nginx
# resumes so it would be better to only use it on scale up
# TODO: only run then when scaling up
docker exec -it zencastr-challenge_nginx_1 nginx -s reload