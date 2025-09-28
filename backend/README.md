# Task Management Backend

Backend API untuk aplikasi task management menggunakan Go dan Gin framework.

## Features

- JWT Authentication
- Task Management
- Role-based Access Control (Pelaksana, Leader, Manager)
- Task Status Management (Submitted, Revision, Approved, In Progress, Completed)
- Progress Tracking
- Task History
- MySQL Database
- Auto Migration

## Tech Stack

- **Language**: Go 1.21.4
- **Framework**: Gin
- **Database**: MySQL
- **Authentication**: JWT
- **ORM**: Native SQL
- **Goose** Migration

## Setup

### Prerequisites
- Go 1.21.4+
- MySQL 8.0+

### Installation

1. Clone repository
```bash
git clone <repository-url>
cd backend
```

2. Install dependencies
```bash
go mod tidy
```

3. Configure environment
```bash
cp .env.example .env
# Edit .env file with your database credentials
```

4. Run application
```bash
go run main.go
```

## API Endpoints

### Authentication
- `POST /api/auth/register` - Register user baru
- `POST /api/auth/login` - Login user

### Tasks
- `GET /api/tasks/` - Get semua tasks
- `POST /api/tasks/` - Create task baru (Pelaksana only)
- `PUT /api/tasks/:id` - Update task (Pelaksana only)
- `PATCH /api/tasks/:id/progress` - Update progress task
- `PATCH /api/tasks/:id/progress/overide` - Override progress (Leader only)
- `PATCH /api/tasks/:id/revise` - Revise task (Leader only)
- `PATCH /api/tasks/:id/approve` - Approve task (Leader only)
- `PATCH /api/tasks/:id/complete` - Complete task (Leader & Pelaksana)

### Health Check
- `GET /api/health-check` - Check server status

## User Roles

1. **Pelaksana**: Dapat membuat, update, dan melakukan progress task
2. **Leader**: Dapat approve, revise, dan override progress task
3. **Manager**: Dapat melihat semua task untuk monitoring

## Task Status Flow

```
Submitted → Approved → In Progress → Completed
     ↓
  Revision → (back to Submitted)
```

## Development

```bash
# Run with auto-reload
go run main.go
```