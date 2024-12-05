# Golang Web Application

This project is a simple web application built using Golang and PostgreSQL. It includes two portals: a receptionist portal and a doctor portal. The application allows receptionists to manage patient records (CRUD) and doctors to view and update patient details.

---

## **Features**
1. A single login page for both receptionist and doctor portals.
2. Receptionist functionalities:
   - Register new patients.
   - Perform Create, Read, Update, and Delete (CRUD) operations on patient records.
3. Doctor functionalities:
   - View registered patients.
   - Update patient details.

---

## **Tech Stack**
- **Programming Language:** Go (Golang)
- **Web Framework:** Pure `net/http` (Go standard library)
- **Database:** PostgreSQL
- **Migration Tool:** `golang-migrate`
- **Authentication:** JWT-based login system

---

## **Setup and Installation**

### **1. Prerequisites**
- Go installed (v1.18 or higher): [Install Go](https://go.dev/dl/)
- PostgreSQL installed and running: [Download PostgreSQL](https://www.postgresql.org/download/)
- `golang-migrate` installed: [Install Migrate CLI](https://github.com/golang-migrate/migrate)

---

### **2.Install Dependencies**
go mod tidy

### **3. Start the Server**
go run main.go
