# Medioxe - AI-Powered Medical Appointment Platform

## Description

Medioxe is a RESTful API application built using Golang, Echo Framework, GORM, and PostgreSQL.

This application helps patients connect with doctors more efficiently through AI-assisted symptom analysis, appointment scheduling, appointment tracking, and email notifications.

The system supports three user roles:

* Admin
* Doctor
* Patient

Patients can analyze symptoms using AI recommendations, discover suitable doctors, book appointments, and track appointment history. Doctors can manage schedules and appointments, while admins can manage doctors and monitor appointment histories.

---

## Features

### Authentication

* Register User
* Login User
* JWT Authentication

### Patient

* Get User Profile
* AI Symptom Analysis using Gemini AI
* Get Doctors
* Get Doctor Detail
* View Doctor Schedules
* Create Appointment
* View My Appointments
* View Appointment History

### Doctor

* Get Doctor Profile
* Update Doctor Profile
* Create Schedule
* Update Schedule
* Delete Schedule
* Get My Schedules
* Get Doctor Appointments
* Confirm Appointment
* Complete Appointment
* Cancel Appointment

### Admin

* Create Doctor
* Update Doctor
* Activate Doctor
* Deactivate Doctor
* View All Appointment Histories
* View Appointment History by Appointment ID

### Appointment History

* Automatic Appointment History Tracking
* Pending Status
* Confirmed Status
* Completed Status
* Cancelled Status

### Notification

* Email Notification using Mailjet
* Appointment Created Notification
* Appointment Confirmed Notification
* Appointment Completed Notification
* Appointment Cancelled Notification

---

## Tech Stack

* Golang
* PostgreSQL
* Echo Framework
* GORM
* Supabase
* JWT Authentication
* Gemini AI
* Mailjet
* Postman
* GoMock
* Unit Testing
* Railway

---

## Installation

### 1. Clone Repository

```bash
git clone <repository-url>
cd medioxe
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Configure Environment Variables

Create `.env` file and configure:

```env
APP_PORT=8080

DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_SSLMODE=

JWT_SECRET=

GEMINI_API_KEY=

MAILJET_BASE_URL=https://api.mailjet.com
MAILJET_API_KEY=
MAILJET_SECRET_KEY=
MAILJET_SENDER_EMAIL=
MAILJET_SENDER_NAME=
```

### 4. Create Database

Run SQL scripts located in:

```text
sql/
```

Make sure all required tables are created:

* users
* specialties
* hospitals
* doctor_profiles
* schedules
* symptom_analyses
* appointments
* appointment_histories

### 5. Run Application

```bash
go run ./cmd
```

Server will run on:

```text
http://localhost:8080
```

---

## API Documentation

Complete API documentation and example requests are available in:


[Postman](https://documenter.getpostman.com/view/27897753/2sBXwvJ8Xy)


---

## Third Party Services

### Gemini AI

Used for symptom analysis and doctor specialty recommendations.

Required environment variable:

```env
GEMINI_API_KEY
```

### Mailjet

Used for email notifications.

Required environment variables:

```env
MAILJET_API_KEY
MAILJET_SECRET_KEY
MAILJET_SENDER_EMAIL
MAILJET_SENDER_NAME
```

---

## Testing

Unit testing is implemented using:

* Go Testing Package
* GoMock
* Testify

Example:

```bash
go test ./internal/service -v
```

---

## Deployment

API Base URL:


[railway](https://medioxe-production.up.railway.app)


---

## Future Improvements

* Payment Gateway Integration
* Appointment Reminder System
* Medical Records Management
* Doctor Rating & Review System
* Admin Dashboard
* Doctor Dashboard
* Microservices Architecture

```
```
