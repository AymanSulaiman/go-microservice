# go-microservice


Building a gin microservice that is the backend of a blog. Just learning how to use Go better

````
┌──────────────────┐
│     main.go      │
├──────────────────┤
│  - import Gin    │
│  - define routes │
│  - start server  │
└──────────────────┘
         │
         ▼
┌──────────────────┐
│    handlers/     │
├──────────────────┤
│  - handle routes │
└──────────────────┘
         │
         ▼
┌──────────────────┐
│   controllers/   │
├──────────────────┤
│  - handle logic  │
│  - interact with │
│    database      │
└──────────────────┘
         │
         ▼
┌──────────────────┐
│    models/       │
├──────────────────┤
│  - define        │
│    data models   │
└──────────────────┘
         │
         ▼
┌──────────────────┐
│    database/     │
├──────────────────┤
│  - interact with │
│    database      │
└──────────────────┘
         │
         ▼
┌──────────────────┐
│    config/       │
├──────────────────┤
│  - store app     │
│    configuration │
└──────────────────┘
````