Passos a serem executados:

1. Subir containers por meio do **_docker-compose up -d_** na pasta cmd/ordersystem
2. Executar servidores por meio do comando **_go run main.go wire_gen.go_** na pasta cmd/ordersystem  
3. Executar chamadas no arquivo **_api/api.http_**

Portas dos serviços:

Web server na porta 8000  
gRPC server na porta 50051  
GraphQL server na porta 8080  
RabbitMQ na porta 5672  
