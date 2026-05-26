# Application

- ~Frontend: Nginx serving a static page~
- ~API: A small Go service~
- Database: PostgreSQL
  - Create POST endpoint
    - Authenticate with postgres
    - INSERT
    - return "Saved!"
- Cache: Redis
- Message queue: RabbitMQ
- Worker: A small consumer service that reads from RabbitMQ and writes analytics to the DB