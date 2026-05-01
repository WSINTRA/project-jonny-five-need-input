# ResearchBot

A general-purpose AI research platform that takes a location + topic, browses the web for real-world data, generates summaries, and stores everything as a knowledge graph.

**Work in progress.**

## Goal

Build a knowledge acquirer agent that can research specific queries and store structured findings into a graph database. Example: *"What are good venues for musicians to reach out to in Middlesbrough?"* → list of venues with contact details, live music ratings, and recent gig data.

## Stack

- **Agent** — TypeScript browser app (Vite)
- **Server** — Go REST API, tool execution, graph writes
- **Database** — Neo4j knowledge graph
- **LLM** — llama-server (local inference)
- **Diagrams** — PlantUML server (Docker, port 7070)

## Diagrams

PlantUML diagrams live in `docs/`. Render with:

```bash
curl -s http://localhost:7070/png -H 'Content-Type: text/plain' \
  --data-binary @docs/your-diagram.puml \
  --output docs/your-diagram.png
```

## Status

Foundation phase complete. Knowledge graph schema, research pipeline, and tool integration in progress.
