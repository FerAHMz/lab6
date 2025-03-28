# Lab 6: Go API with Docker - La Liga Tracker

<div align="center">
  <h1>La Liga Tracker</h1>
  <p>
    A RESTful API for tracking football matches using Go and Docker
  </p>
</div>

## Project Description

This project implements a RESTful API for tracking football matches. Built with Go and containerized with Docker, it provides endpoints for managing match data including creating, reading, updating, and deleting matches.

## Features

- **CRUD Operations:** Complete set of endpoints for match management
- **Docker Support:** Containerized application for easy deployment
- **In-Memory Storage:** Lightweight data storage solution
- **CORS Enabled:** Frontend-ready with CORS configuration

## API Testing

### POST Request Test
![POST Request Test](https://github.com/FerAHMz/lab6/blob/main/Images/Prueba%20post.png)

## API Endpoints

- `GET /api/matches` - List all matches
- `GET /api/matches/{id}` - Get a specific match
- `POST /api/matches` - Create a new match
- `PUT /api/matches/{id}` - Update a match
- `DELETE /api/matches/{id}` - Delete a match

## Running the Application

### With Docker
```bash
docker compose up --build
```