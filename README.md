# gopoc-connector
The Connector API for a golang based POC project for managing EV stations.

# Setup
1. Copy `.example.env` file and rename it to `.env`
2. Set appropriate values for environment variables in `.env`
3. [Optional] Run `docker compose -f docker-compose.kafka.yml up` if Kafka and Zookeeper servers are not already running
4. Run `docker compose up --build` (Ensure that the app container and Kafka, Zookeeper containers are running on the same network in Docker)
