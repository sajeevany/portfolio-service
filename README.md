# Portfolio Service
Wrapper service for storing user portfolio information

Generate swagger files:

    swag init cmd/porfolio=service/main.go -g cmd/portfolio-service/main.go  -d ./

Access swagger API:

    ${HOST}:{PORT}/swagger/index.html
    ie. http://localhost:80/swagger/index.html

Run tests:

    go test ./...
    
Run containers (debug):

    docker run -d -p 79:8080 portfolio-service:latest
    docker run -d -p 3000:3000 -p 3001:3001 -p 3002:3002 -p 3003:3003 -v $(pwd)/deploy/aerospike.conf:/opt/aerospike/etc/aerospike.conf --foreground --config-file /opt/aerospike/etc/aerospike.conf aerospike:4.5.3.6 


