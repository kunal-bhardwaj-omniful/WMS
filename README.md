# 📦 Warehouse Management System (WMS)

A backend system built with Golang to manage Hubs, SKUs, and Inventory across warehouse locations.

---

## 🧱 Architecture

This project follows a layered **MVC-inspired architecture** with clear separation of concerns:

### ➤ Model (`/domain`)
- Plain Go structs representing database entities: `Hub`, `SKU`, `Inventory`.
- No business logic; just data representation.

### ➤ Repository Layer (`/repo`)
- Responsible for all **database interactions** using GORM.
- Contains methods like `CreateHub`, `GetAllSkus`, `DecreaseAvailableQty`, etc.
- Encapsulates SQL logic and data access patterns.

### ➤ Service Layer (`/service`)
- Implements business logic and validations (e.g. ID checks, name required).
- Serves as a bridge between controllers and repositories.

### ➤ Controller/Handler Layer (`/controller`)
- Exposes HTTP APIs using **Gin** web framework.
- Handles JSON binding, request validation, and response formatting.

---

## 🚀 Features

- Manage **Hubs** (Warehouses)
- Manage **SKUs** (Stock Keeping Units)
- Track inventory quantities per hub
- Standardized JSON API responses
- UUID-based resource identification
- Structured error handling

---

📦 Warehouse Management System (WMS) API

This API manages **Hubs**, **SKUs**, and **Inventory**. It supports creating hubs and SKUs, retrieving data, and decreasing inventory quantities across different types (available, allocated, damaged).

## 🌐 Base URL

/api/v1

---

## 🏬 Hubs

### 🔹 Create Hub

**POST** `/api/v1/hubs`

**Request Body**:
```json
{
  "name": "Main Hub",
  "location": "New York"
}
```
Success Response (201):
```json
{
  "message": "hub created successfully"
}
```
Error Response (400):
```json

{
  "error": "hub name cannot be empty"
}
```
🔹 Get All Hubs
GET /api/v1/hubs

Success Response (200):
```json
[
  {
    "id": "8db7a31f-03fa-4c3b-a2d5-07b6a53de7f1",
    "name": "Main Hub",
    "location": "New York"
  }
]
```
🔹 Get Hub by ID
GET /api/v1/hubs/{id}

Success Response (200):
```json
{
  "id": "8db7a31f-03fa-4c3b-a2d5-07b6a53de7f1",
  "name": "Main Hub",
  "location": "New York"
}
```
Error Response (404):
```json
{
  "error": "hub not found"
}
```
📦 SKUs
🔹 Create SKU
POST /api/v1/skus

Request Body:
```json
{
  "name": "Product A",
  "description": "Blue T-shirt"
}
```

Success Response (201):
```json
{
  "message": "sku created successfully"
}
```

Error Response (400):
```json
{
  "error": "SKU name cannot be empty"
}
```
🔹 Get All SKUs
GET /api/v1/skus

Success Response (200):
```json
[
  {
    "id": "45f7a31e-12ad-46b1-91d4-05c7c6e539ee",
    "name": "Product A",
    "description": "Blue T-shirt"
  }
]
```

🔹 Get SKU by ID
GET /api/v1/skus/{id}

Success Response (200):
```json

{
  "id": "45f7a31e-12ad-46b1-91d4-05c7c6e539ee",
  "name": "Product A",
  "description": "Blue T-shirt"
}
```
Error Response (404):
```json

{
  "error": "SKU not found"
}
```
📊 Inventory
🔹 Get Inventory by SKU and Hub
GET /api/v1/inventory?sku_id={sku_id}&hub_id={hub_id}

Success Response (200):
```json
{
  "sku_id": "45f7a31e-12ad-46b1-91d4-05c7c6e539ee",
  "hub_id": "8db7a31f-03fa-4c3b-a2d5-07b6a53de7f1",
  "available_qty": 50,
  "allocated_qty": 10,
  "damaged_qty": 2
}
```
Error Response (404):
```json
{
  "error": "failed to fetch inventory"
}
```
🔹 Decrease Inventory Quantities
POST /api/v1/inventory/decrease

Request Body:
```json
{
  "sku_id": "45f7a31e-12ad-46b1-91d4-05c7c6e539ee",
  "hub_id": "8db7a31f-03fa-4c3b-a2d5-07b6a53de7f1",
  "available_qty": 5,
  "allocated_qty": 2,
  "damaged_qty": 1
}
```
Success Response (200):
```json
{
  "message": "inventory decreased successfully"
}
```
Error Response (400):
```json

{
  "error": "quantities must be non-negative"
}
```

Error Response (422):
```json

{
  "error": "not enough available/allocated/damaged quantity"
}

```
