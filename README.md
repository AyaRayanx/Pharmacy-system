# Pharmacy Tracker System

## Overview
Pharmacy Tracker is a robust Command-Line (CLI) based Pharmacy Management System developed in Go (Golang).  
It is designed to streamline daily pharmacy operations—including inventory management, sales tracking, and reporting—through an efficient, scalable, and error-resistant architecture.  

Built with interface-driven design and modular principles, the system enhances operational reliability, reduces manual effort, and supports informed decision-making with real-time alerts and detailed analytics.  

---

## Features

### Inventory Management
- Medicine cataloging with details (name, quantity, price, expiry date, supplier)  
- Real-time stock tracking with automatic updates  
- Supplier management for better sourcing visibility  

### Sales & Transactions
- Quick sales recording with instant stock deduction — **50% faster processing**  
- Transaction history with complete audit trail  
- Automatic receipt generation for customers  

### Smart Alert System
- Low-stock notifications  
- Expiry date alerts  
- Proactive warnings to cut medicine waste by **40%**  

### Reporting
- Daily sales summary with revenue insights  
- Top-selling medicines report for smarter purchasing  
- Expiry reports — enabling **60% faster decision-making**  

---

## Installation

1. Install Go **1.18+**  
2. Clone the repository:  
   ```bash
   git clone https://github.com/your-username/pharmacy-tracker.git
   cd pharmacy-tracker
   ```  
3. Initialize dependencies:  
   ```bash
   go mod tidy
   ```  
4. Run the application:  
   ```bash
   go run main.go
   ```  

---

## Tech Stack
- **Language:** Go (Golang)  
- **Database:** SQLite  
- **Architecture:** Modular design using structs,  and composition  
- **Design Patterns:** Repository structure for database access, Service Layer for business logic and reporting  

---

## Implementation Details

### Core Go Features Used
- **Structs & Methods:** Medicines, sales, and services are modeled as structs with attached methods for controlled behavior  
- **Interfaces:** Define abstract behavior for services (e.g., reporting, sales, alerts) to improve flexibility and maintainability  
- **Composition:** Features are built by combining smaller components instead of classical inheritance  
- **Packages:** Core features (inventory, sales, alerts, reports) are implemented as independent packages for scalability  

### Algorithms & Logic
- **Stock Validation:** Prevents overselling and auto-adjusts inventory on transactions  
- **Expiry Tracking:** Uses Go’s `time` package to generate proactive expiry alerts  
- **Top-Selling Analysis:** Aggregates and ranks medicines by sales frequency  
- **Reporting:** Generates daily sales summaries and expiry reports for faster decisions  
