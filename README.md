# Order Management Service

This service handles fetching customer and product details from external APIs, creating orders, and storing them in the database. 

## Overview
The Order Management Service is a backend application built using Go, with functionalities to:
- Fetch customer data from an external service.
- Fetch product data from an external service.
- Create orders based on fetched customer and product data.
- Store orders in a MySQL database.


## Features
- **Product Fetch**: Retrieve product details from an external API.
- ![Screenshot 2025-01-21 104030](https://github.com/user-attachments/assets/ce19b382-8238-48eb-b095-b2fb233c2550)
![Screenshot 2025-01-21 104118](https://github.com/user-attachments/assets/c49b5e83-cf13-4304-8558-dabda0400ec9)

- **Customer Fetch**: Retrieve customer details from an external API.
- ![Screenshot 2025-01-21 104045](https://github.com/user-attachments/assets/d12d6c5a-eb32-435b-9c15-a02f19101aa5)
- ![Screenshot 2025-01-21 104211](https://github.com/user-attachments/assets/1acdf184-2bdd-40b9-8061-7110685cb7c2)


- **Order Creation**: Create orders by linking customers and products.
- **Database Storage**: Save orders in a MySQL database.

## Prerequisites
- **Go**: Ensure you have Go installed.
- **MySQL Database**: A MySQL database named `orderdb` is required.
- **External APIs**: APIs for fetching customer and product data.
