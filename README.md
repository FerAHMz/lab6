# Lab 6: Go API with Docker - La Liga Tracker

<div align="center">
  <h1>La Liga Tracker</h1>
  <p>
    A RESTful API for tracking football matches using Go and Docker
  </p>
</div>

## Project Description

This project implements a RESTful API for tracking football matches. Built with Go and containerized with Docker, it provides endpoints for managing match data including creating, reading, updating, and deleting matches, as well as tracking match events like goals, cards, and extra time.

## Features

- **CRUD Operations:** Complete set of endpoints for match management
- **Match Events:** Track goals, yellow cards, red cards, and extra time
- **Docker Support:** Containerized application for easy deployment
- **SQLite Storage:** Persistent data storage
- **CORS Enabled:** Frontend-ready with CORS configuration
- **API Documentation:** Interactive Swagger UI documentation

## API Testing

You can find the Postman collection for testing the API endpoints here:
[Postman Collection](https://fernandohernandez-4170971.postman.co/workspace/Fernando-Hernandez's-Workspace~5ebfdd38-b11b-4507-82dd-b9324e34cfa7/collection/43568958-988e5f50-eee8-427a-b2ed-751faae460cc?action=share&creator=43568958)

### POST Request Test
![POST Request Test](https://github.com/FerAHMz/lab6/blob/main/Images/Prueba%20post.png)

## API Endpoints

- `GET /api/matches` - List all matches
- `GET /api/matches/{id}` - Get a specific match
- `POST /api/matches` - Create a new match
- `PUT /api/matches/{id}` - Update a match
- `DELETE /api/matches/{id}` - Delete a match
- `PATCH /api/matches/{id}/goals` - Update match goals
- `PATCH /api/matches/{id}/yellowcards` - Add yellow card
- `PATCH /api/matches/{id}/redcards` - Add red card
- `PATCH /api/matches/{id}/extratime` - Update extra time

## API Documentation
Access the interactive API documentation at:
```bash
http://localhost:8081/swagger-ui.html
```

## Running the Application

### With Docker
```bash
docker compose up --build
```