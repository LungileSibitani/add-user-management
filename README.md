# User Management System

## Project Overview

This repository contains a **full-stack User Management System** with separate **backend** and **frontend** components.

### Backend

- **Tech:** Go (Golang)  
- **Purpose:** Provides a RESTful API for user management (CRUD operations).  
- **Files & Features:**
  - `main.go`: Starts the server and registers API routes.
  - `db.go`: Handles database connection (PostgreSQL).
  - `handlers.go`: Contains API request handlers (`GetUsers`, `CreateUser`, `UpdateUser`, `DeleteUser`).
  - `models.go`: Defines data structures (e.g., `User`).

- **Flow:**
  1. `main.go` initializes the database connection.
  2. API routes are defined (e.g., `/api/users`).
  3. Handlers interact with the database and return JSON responses.
  4. Backend listens on a port (default: `8080`) for frontend requests.

### Frontend

- **Tech:** Vue.js (JavaScript framework)  
- **Purpose:** Provides a Single Page Application (SPA) to interact with the backend.  
- **Files & Features:**
  - `App.vue`: Main component handling UI for forms, user list, and modals.
  - `services/`: Contains Axios API calls to backend.
  - **Styling:** Modern glassmorphism design with gradients and responsive buttons.
  - **Features:**
    - Add new users
    - Edit existing users via modal
    - Delete users
    - Toggle and search user list

- **Flow:**
  1. User fills a form and submits.
  2. Axios sends HTTP requests to the Go backend (`POST /api/users` etc.).
  3. Backend responds with data, frontend updates the UI dynamically.

