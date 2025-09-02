# Web Application for Riddles - Design Document

## 1. Overview

This document outlines the design for a full-stack web application for riddles with user authentication, progress tracking, categorized riddle catalog with filtering capabilities, and rating system. The application will be built using modern technologies for both frontend and backend.

### 1.1 Purpose
To create a web application that allows users to browse, solve, and interact with riddles while tracking their progress and preferences.

### 1.2 Key Features
- User registration and authentication system
- Riddle catalog with filtering by category, difficulty, and status
- Progress tracking for solved riddles
- Favorites system for bookmarking preferred riddles
- Like/Dislike system for riddles with counters
- Daily riddle feature
- Responsive design with Russian language UI

## 2. Technology Stack

### 2.1 Frontend (client/)
- **Framework**: Next.js 15 with App Router
- **Language**: TypeScript
- **State Management**: TanStack Query (React Query)
- **UI Components**: Shadcn/ui with Tailwind CSS, designed to closely resemble riddles.com with a clean, modern interface
- **Directory Structure**: src/ directory structure with components organized for easy maintenance
- **Build Tool**: Turbopack

### 2.2 Backend (server/)
- **Language**: Go
- **Framework**: Echo
- **ORM**: GORM
- **Database**: PostgreSQL

### 2.3 Database
- **System**: PostgreSQL
- **Connection Details**:
  - Host: localhost
  - User: postgres
  - Password: admin
  - Database Name: riddles
  - Port: 5432

## 3. Architecture

### 3.1 High-Level Architecture
The application follows a client-server architecture with a clear separation of concerns:
- **Frontend**: Handles user interface, user interactions, and presentation logic
- **Backend**: Manages business logic, data processing, and database interactions
- **Database**: Stores all application data including users, riddles, progress, favorites, and ratings

### 3.2 Backend Architecture
The backend follows a layered architecture pattern with the following components:

#### 3.2.1 Handler Layer
Responsible for handling HTTP requests and responses:
- Authentication Handlers (registration, login, token refresh)
- User Handlers (profile, progress)
- Riddle Handlers (list, details, answer validation)
- Favorites Handlers (add/remove favorites)
- Rating Handlers (like/dislike riddles)
- Daily Riddle Handler (get daily riddle)

#### 3.2.2 Service Layer
Contains business logic and acts as an intermediary between handlers and repositories:
- Authentication Service (user registration, login, token generation)
- User Service (user profile management)
- Riddle Service (riddle retrieval, answer validation, like/dislike management)
- Favorites Service (managing user favorites)
- Rating Service (managing riddle likes/dislikes)
- Daily Riddle Service (selecting and retrieving daily riddles)

#### 3.2.3 Repository Layer
Handles data access and persistence:
- User Repository
- Riddle Repository
- Progress Repository
- Favorite Repository
- Rating Repository
- Daily Riddle Repository

### 3.3 Frontend Architecture
The frontend follows a component-based architecture with clear separation of concerns. The design will use Shadcn/ui components with Tailwind CSS styling to create a UI that closely resembles riddles.com:

#### 3.3.1 Component Structure
- Layout Components (header, footer, navigation)
- Page Components (home, riddles list, profile, auth)
- Feature Components (riddle cards, daily riddle, filters)
- UI Components (buttons, forms, modals from Shadcn/ui)

#### 3.3.2 State Management
- TanStack Query for server state management
- React Context API for client state when needed
- localStorage for persisting authentication tokens

## 4. Database Design

### 4.1 Entity Relationship Diagram
Entity relationship diagram showing relationships between Users, Riddles, Categories, Progress, Favorites, Ratings, and Daily Riddles tables. The Ratings table specifically supports the like/dislike functionality where each user can rate a riddle once with either a like (+1) or dislike (-1).

### 4.2 Database Schema Details

#### 4.2.1 Users Table
| Column | Type | Description |
|--------|------|-------------|
| id | SERIAL (PK) | Unique user identifier |
| username | VARCHAR(50) | User's display name |
| email | VARCHAR(100) | User's email (unique) |
| password_hash | VARCHAR(255) | Hashed password |
| created_at | TIMESTAMP | Account creation timestamp |
| updated_at | TIMESTAMP | Last update timestamp |

#### 4.2.2 Riddles Table
| Column | Type | Description |
|--------|------|-------------|
| id | SERIAL (PK) | Unique riddle identifier |
| title | VARCHAR(255) | Riddle title |
| description | TEXT | Full riddle text |
| answer | TEXT | Correct answer (case-insensitive comparison) |
| category_id | INTEGER (FK) | Reference to category |
| difficulty | VARCHAR(20) | Difficulty level (easy, medium, hard) |
| created_at | TIMESTAMP | Creation timestamp |
| updated_at | TIMESTAMP | Last update timestamp |

#### 4.2.3 Categories Table
| Column | Type | Description |
|--------|------|-------------|
| id | SERIAL (PK) | Unique category identifier |
| name | VARCHAR(50) | Category name |

#### 4.2.4 User Riddle Progress Table
| Column | Type | Description |
|--------|------|-------------|
| id | SERIAL (PK) | Unique progress identifier |
| user_id | INTEGER (FK) | Reference to user |
| riddle_id | INTEGER (FK) | Reference to riddle |
| solved | BOOLEAN | Whether riddle is solved |
| solved_at | TIMESTAMP | When riddle was solved |
| created_at | TIMESTAMP | Progress creation timestamp |
| updated_at | TIMESTAMP | Last update timestamp |

#### 4.2.5 Favorites Table
| Column | Type | Description |
|--------|------|-------------|
| id | SERIAL (PK) | Unique favorite identifier |
| user_id | INTEGER (FK) | Reference to user |
| riddle_id | INTEGER (FK) | Reference to riddle |
| created_at | TIMESTAMP | When riddle was favorited |

#### 4.2.6 Riddle Ratings Table
| Column | Type | Description |
|--------|------|-------------|
| id | SERIAL (PK) | Unique rating identifier |
| user_id | INTEGER (FK) | Reference to user |
| riddle_id | INTEGER (FK) | Reference to riddle |
| rating | INTEGER | Rating value (1 for like, -1 for dislike) |
| created_at | TIMESTAMP | When rating was given |

#### 4.2.7 Daily Riddle Table
| Column | Type | Description |
|--------|------|-------------|
| id | SERIAL (PK) | Unique daily riddle identifier |
| riddle_id | INTEGER (FK) | Reference to riddle |
| featured_date | DATE | Date when riddle is featured |
| created_at | TIMESTAMP | Creation timestamp |

## 5. API Design

### 5.1 Authentication API
#### POST /api/auth/register
- **Description**: Register a new user
- **Request Body**: Registration credentials (username, email, password)
- **Response**: Access token, refresh token, and user information

#### POST /api/auth/login
- **Description**: Authenticate user
- **Request Body**: Login credentials (email, password)
- **Response**: Access token, refresh token, and user information

#### POST /api/auth/refresh
- **Description**: Refresh access token
- **Request Body**: Refresh token
- **Response**: New access token and refresh token

### 5.2 Users API
#### GET /api/users/profile
- **Description**: Get authenticated user's profile
- **Headers**: Authorization: Bearer <access_token>
- **Response**: User profile information (id, username, email, created_at)

#### GET /api/users/progress
- **Description**: Get user's riddle solving progress
- **Headers**: Authorization: Bearer <access_token>
- **Response**: Progress statistics (total riddles, solved riddles, success rate)

### 5.3 Riddles API
#### GET /api/riddles
- **Description**: Get list of riddles with filtering and pagination
- **Query Parameters**:
  - category: string (optional)
  - difficulty: string (optional)
  - status: string (solved/unsolved/all) (optional)
  - favorite: boolean (optional)
  - page: number (default: 1)
  - limit: number (default: 10)
- **Response**: List of riddles with pagination information, including like/dislike counts for each riddle

#### GET /api/riddles/{id}
- **Description**: Get specific riddle details
- **Response**: Detailed riddle information including metadata, user interaction status, and like/dislike counts

#### POST /api/riddles/{id}/answer
- **Description**: Submit answer for a riddle
- **Headers**: Authorization: Bearer <access_token>
- **Request Body**: Submitted answer
- **Response**: Answer validation result (correctness and feedback message)

### 5.4 Favorites API
#### POST /api/favorites/{riddle_id}
- **Description**: Add riddle to favorites
- **Headers**: Authorization: Bearer <access_token>
- **Response**: 204 No Content

#### DELETE /api/favorites/{riddle_id}
- **Description**: Remove riddle from favorites
- **Headers**: Authorization: Bearer <access_token>
- **Response**: 204 No Content

### 5.5 Ratings API
#### POST /api/ratings/{riddle_id}
- **Description**: Rate a riddle (like/dislike)
- **Headers**: Authorization: Bearer <access_token>
- **Request Body**: Rating value (1 for like, -1 for dislike)
- **Response**: 204 No Content

### 5.6 Daily Riddle API
#### GET /api/daily-riddle
- **Description**: Get today's featured riddle
- **Response**: Daily riddle information (id, title, description, category, difficulty)

## 6. Frontend Component Architecture

### 6.1 Page Components
#### Home Page (/)
- Daily Riddle Component
- Random Riddles Preview Component
- Navigation Header
- Footer

#### Riddles List Page (/riddles)
- Riddle Filters Component
- Riddle List Component
- Pagination Component

#### Authentication Pages (/login, /register)
- Login Form Component
- Registration Form Component

#### Profile Page (/profile)
- User Profile Component
- Progress Statistics Component

### 6.2 Feature Components
#### Daily Riddle Component
- Displays the riddle of the day
- Shows title, description, and category
- Provides answer submission form (for authenticated users)
- Shows like/dislike counters

#### Riddle Card Component
- Displays riddle preview (title, category, difficulty)
- Shows solution status for authenticated users
- Shows like and dislike counters
- Expandable to show full details and answer form

#### Riddle Detail Component
- Displays full riddle details
- Provides answer submission form for authenticated users
- Shows like/dislike counters and buttons
- Favorite toggle button

#### Filters Component
- Category filter dropdown
- Difficulty filter
- Status filter (solved/unsolved/all)
- Favorites filter (for authenticated users)
- Rating filter (most liked, most disliked)

### 6.3 UI Components
Using Shadcn/ui components with Tailwind CSS styling to match riddles.com design:
- Buttons (with like/dislike icons and counters)
- Forms
- Input fields
- Select dropdowns
- Cards (with consistent styling and shadow effects)
- Modals
- Toast notifications
- Badges for difficulty levels
- Progress indicators for user statistics

## 7. Authentication Flow

### 7.1 User Registration
1. User fills registration form
2. Frontend sends request to /api/auth/register
3. Backend validates input and creates user
4. Backend returns access and refresh tokens
5. Frontend stores tokens and redirects to home page

### 7.2 User Login
1. User fills login form
2. Frontend sends request to /api/auth/login
3. Backend validates credentials
4. Backend returns access and refresh tokens
5. Frontend stores tokens and redirects to home page

### 7.3 Token Management
- Access tokens stored in memory
- Refresh tokens stored in secure HTTP-only cookies
- Token refresh handled automatically when access token expires

### 7.4 Protected Routes
- Profile page requires authentication
- Answer submission requires authentication
- Favorites management requires authentication
- Progress tracking requires authentication

## 8. Data Flow

### 8.1 Backend Data Flow
Data flow diagram showing the request processing path from HTTP Request through Middleware, Handler, Service, Repository, and Database.

### 8.2 Frontend Data Flow
Data flow diagram showing the user interaction path from User Interaction through Component, TanStack Query, API Client, and Backend API.

## 9. Deployment Configuration

### 9.1 Environment Variables
#### Backend (.env)
Environment variables for database connection and JWT secret key

#### Frontend (.env.local)
Environment variable for API URL

### 9.2 Database Seeding
- Initial seeding: 20 riddles per category
- Future expansion: 50 riddles per category
- Total target: 500 riddles across categories

### 9.3 Daily Riddle Selection
- Algorithm to select 6 riddles daily based on date
- Ensures consistent selection for all users on the same day

## 10. Testing Strategy

### 10.1 Backend Testing
- Unit tests for services and repositories
- Integration tests for API endpoints
- Database migration tests

### 10.2 Frontend Testing
- Component unit tests
- Integration tests for API client
- End-to-end tests for critical user flows

## 11. Security Considerations

### 11.1 Authentication Security
- Password hashing with bcrypt
- JWT token expiration and refresh
- HTTP-only cookies for refresh tokens
- CORS configuration

### 11.2 Data Security
- Input validation and sanitization
- SQL injection prevention through GORM
- Rate limiting for API endpoints

## 12. Performance Considerations

### 12.1 Backend Performance
- Database indexing for frequently queried fields
- Connection pooling
- Caching for non-sensitive data

### 12.2 Frontend Performance
- TanStack Query caching
- Code splitting for pages
- Image optimization (if applicable)

## 13. Future Enhancements

### 13.1 Planned Features
- Social sharing of riddles
- User leaderboards
- Hint system for difficult riddles
- Multi-language support beyond Russian

### 13.2 Scalability Improvements
- Database read replicas
- CDN for static assets
- Load balancing for multiple server instances