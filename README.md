# network-monitoring
A home network monitoring project with a frontend dashboard

Requirements:
    - Docker Desktop installed
    - Go installed

1.) cp .env.examples to .env and set values

2.) From root directory run:
    - "docker compose up --build"

3.) From scanner directory run:
    - "go run ."

The Go scanner will run and persist any devices it has found on the network to the docker postgres db via the monitor-api services running in the container. 

To test and see what devices are in the postgres db run the following from a terminal:
    - "curl http://localhost:8080/monitor-api/devices"

To stop the docker container, run the following command from root directory:
    - "docker compose down" 

To stop and destroy the volume, run:
    - "docker compose down -v"

