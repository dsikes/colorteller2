# ColorTeller2

Inspired by the AWS AppMesh demo, I've re-created the colorteller micro-services and extended them for further simulation purposes.

The updates were designed to illustrate the following:

- Failover
- Circuit Breaking
- Canary Deployments
- Blue/Green Deployments
- Weighted Routing Policies


## Routes

### `GET /` 

returns the defined color of the service as a string w/ http status of 200.
You can set the color by setting the environment variable `COLOR` before starting the service.
If the environment variable is not set, the service will default to `black`.

### `GET /ping`
returns PONG w/ http status of 200. 
This route can be used to serve as a dummy health check endpoint.

### `GET /bad`

returns w/ http status of 400. 
This route simulates a service that has encountered a "bad request".

### `GET /slow` 
returns the defined color of the service as a string w/ http status of 200.
the route determines a random number between 1 and 10 seconds to delay the request.
This route simulates a request that can be slow to respond.

### `GET /slow/2`
the same as the `slow` route, but increases the possible delay duration to 60 seconds.
This route simulates a request that can be REALLY slow to respond.

### `GET /fail`

returns w/ http status of 500. 
This route simulates a service that has encountered an internal server error.

### `GET /random`
returns a random status code between 200, 400 and 500, after a random delay between 1-10 seconds.
If a 200 is returned, the defined or default color is also returned.
This route simulates an real endpoint.
