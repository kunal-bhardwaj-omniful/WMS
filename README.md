# 📦 Warehouse Management System (WMS)

A backend service built using **Golang** and **Gin** for managing hubs, SKUs, and their inventories. The system supports standard CRUD operations, along with inventory management between different warehouses (hubs).

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

yaml
Copy
Edit

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
Success Response (201):

json
Copy
Edit
{
  "message": "hub created successfully"
}
Error Response (400):

json
Copy
Edit
{
  "error": "hub name cannot be empty"
}
🔹 Get All Hubs
GET /api/v1/hubs

Success Response (200):

json
Copy
Edit
[
  {
    "id": "8db7a31f-03fa-4c3b-a2d5-07b6a53de7f1",
    "name": "Main Hub",
    "location": "New York"
  }
]
🔹 Get Hub by ID
GET /api/v1/hubs/{id}

Success Response (200):

json
Copy
Edit
{
  "id": "8db7a31f-03fa-4c3b-a2d5-07b6a53de7f1",
  "name": "Main Hub",
  "location": "New York"
}
Error Response (404):

json
Copy
Edit
{
  "error": "hub not found"
}
📦 SKUs
🔹 Create SKU
POST /api/v1/skus

Request Body:

json
Copy
Edit
{
  "name": "Product A",
  "description": "Blue T-shirt"
}
Success Response (201):

json
Copy
Edit
{
  "message": "sku created successfully"
}
Error Response (400):

json
Copy
Edit
{
  "error": "SKU name cannot be empty"
}
🔹 Get All SKUs
GET /api/v1/skus

Success Response (200):

json
Copy
Edit
[
  {
    "id": "45f7a31e-12ad-46b1-91d4-05c7c6e539ee",
    "name": "Product A",
    "description": "Blue T-shirt"
  }
]
🔹 Get SKU by ID
GET /api/v1/skus/{id}

Success Response (200):

json
Copy
Edit
{
  "id": "45f7a31e-12ad-46b1-91d4-05c7c6e539ee",
  "name": "Product A",
  "description": "Blue T-shirt"
}
Error Response (404):

json
Copy
Edit
{
  "error": "SKU not found"
}
📊 Inventory
🔹 Get Inventory by SKU and Hub
GET /api/v1/inventory?sku_id={sku_id}&hub_id={hub_id}

Success Response (200):

json
Copy
Edit
{
  "sku_id": "45f7a31e-12ad-46b1-91d4-05c7c6e539ee",
  "hub_id": "8db7a31f-03fa-4c3b-a2d5-07b6a53de7f1",
  "available_qty": 50,
  "allocated_qty": 10,
  "damaged_qty": 2
}
Error Response (404):

json
Copy
Edit
{
  "error": "failed to fetch inventory"
}
🔹 Decrease Inventory Quantities
POST /api/v1/inventory/decrease

Request Body:

json
Copy
Edit
{
  "sku_id": "45f7a31e-12ad-46b1-91d4-05c7c6e539ee",
  "hub_id": "8db7a31f-03fa-4c3b-a2d5-07b6a53de7f1",
  "available_qty": 5,
  "allocated_qty": 2,
  "damaged_qty": 1
}
Success Response (200):

json
Copy
Edit
{
  "message": "inventory decreased successfully"
}
Error Response (400):

json
Copy
Edit
{
  "error": "quantities must be non-negative"
}
Error Response (422):

json
Copy
Edit
{
  "error": "not enough available/allocated/damaged quantity"
}

