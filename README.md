# Crypto Aggregator

A Kubernetes-based microservices for aggregating cryptocurrency data, featuring:

- **Fetcher**: a CronJob that runs every 5 minutes to fetch and save data to a database.
- **GraphQL API**: reads from the database, caches in Redis, and exposes data to clients.
- **Redis**: caching layer to speed up responses and reduce database load.

---

## Architecture

```mermaid
flowchart LR
    A[Fetcher CronJob] -->|writes data| B[(Database)]
    B -->|queries| C[GraphQL API]
    C -->|caches| D[(Redis)]
    C -->|responds| E[Client]

    style A fill:#f9f,stroke:#333,stroke-width:2px
    style B fill:#bbf,stroke:#333,stroke-width:2px
    style C fill:#fbf,stroke:#333,stroke-width:2px
    style D fill:#bfb,stroke:#333,stroke-width:2px
    style E fill:#ffb,stroke:#333,stroke-width:2px
````

---

## What’s Done

* CronJob configured to run every 5 minutes.
* Redis service running inside Kubernetes with appropriate Service and Pods.
* GraphQL API running and connected to Redis cache.
* Basic tests implemented for the fetcher.
* Kubernetes manifests configured for Deployments, Services, and CronJob.

---

## What’s Left to Do

* Configure cleanup for old CronJob runs (`successfulJobsHistoryLimit`, `failedJobsHistoryLimit`).
* Add health checks (`readinessProbe`, `livenessProbe`) for API and Redis pods.
* Set up persistence for Redis (volumes for data durability).
* Improve GraphQL API test coverage.
* Document GraphQL queries and schema.
* Implement monitoring and alerting (e.g., Prometheus + Grafana).
* Set up CI/CD pipeline for automated deployments.

---

## Running Locally (via Port-Forwarding)

1. Deploy Redis to your cluster:

   ```bash
   kubectl apply -f redis-deployment.yaml
   ```

2. Deploy the GraphQL API:

   ```bash
   kubectl apply -f graphql-api-deployment.yaml
   ```

3. Forward the GraphQL API port to localhost:

   ```bash
   kubectl port-forward pod/<graphql-pod-name> 4000:4000
   ```

4. Access the GraphQL API at `http://localhost:4000`
