# go_lang_microservices
Learning go lang

Consists of following
1. UI
2. Broker service to talk to other service
3. Auth service
4. Log Service
5. Mailer Service
6. RabbitMq support
7. Mongo Support
8. Postgres Support
9. RPC support
10. GRPC support
11. REST support
12. Docker support
13. Kubernetes support

# TO Start Services
goto /project folder
run the command = "make down up_build"
To create the rabbitmq exchange -
- goto "localhost:15672"
- goto exchanges tab
- create exchange logs_topic (durable)
