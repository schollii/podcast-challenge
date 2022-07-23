# Podcast Challenge

## Description 

Build and deploy a very simple REST API:

1. Build a standalone very simple REST app
2. Modify or retrieve a podcast. 3 routes max.
3. Create a docker container for this service
4. Create script / config for deploying and autoscaling this container in the method 
   that you best see fit
5. Create an integration test for one route on the REST API
6. Run integration test as a pre commit hook


## Assumptions

I have assumed that:

- It's ok to use Go for the micro-service.
- Using `docker-compose scale N` is acceptable form of auto-scaling even though it 
  is only partial automation: docker-compose takes care of starting and stopping 
  the containers to meet N, but it's up to me to run the command. Fully automated 
  scaling would be possible in kubernetes, but then several other aspects of deployment 
  and testing as significantly more complicated. 
- I don't need to actually store or create the podcast, in the interest of simplicity.
  This results in some limitations, discussed below.
- It is ok for the micro-service to be stateful. This is for simplicity, I would 
  normally favor stateless. 


## Limitations 

I have kept the ReST API very simple: when the service starts, the podcast is assumed to 
exist in the container, so all requests for the "known" podcast will work from the start. 
Once deleted, PUT and GET and further DELETE will fail (404). 

Since testing is setup to use docker-compose as deployment/runtime environment, scaling is 
not fully automated. And since the free nginx is used, DNS refresh is clunky. This means 
that unit tests for N>1 replica will fail for current architecture, and scale down can 
lead to a few seconds of downtime (scaling up seems to be fine).

When using run-test.sh, any existing docker-compose up will be restarted using a fresh
build of the podcast service. 

The podcast micro-service is stateful, as it stores podcast existence flag in the container. 
This means that for N>1 replicas, integration test will fail: round-robin load-balancing by 
nginx, without session affinity in free version, will cause requests to be handled by 
different replicas. Shared state could be achieved via redis (which would actually make the 
pods stateless, far better) or by a replica notifying its siblings when state changes
(harder than using redis). 

Scaling down can cause downtime for a minute sometimes, because nginx free load balancer 
does not refresh its dns cache without a reload (and refreshing on every request just 
for this situation is almost certain a serious performance hit).

TODO: pretend podcast is stored in a DB: use a redis db to share file existence flag


## Requirements

For deploying and testing: 

- Ubuntu 20.x
- docker (>= 20.10)
- docker-compose
- bash
- connection to internet
- $USER is member of `docker` group

For further development: 

- (optional) Go 1.18 (to get code completion etc)
- make (to install the hooks)


# Development 

Run `make init-clone` to setup the git hooks. Thereafter, any commit will trigger a 
rebuild-deploy-test.

Run `./run-test.sh` to build the image, deploy it, and test the service, and tear down.

Run `./scale N` to scale # of containers of podcast service up and down.

If you play around with different HTTP requests to the service, ensure replicas = 1, OR 
do not use DELETE (since this will delete the podcast only in one replica).