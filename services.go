package main

import (
    "database/sql"
    "errors"
    "fmt"
    "time"
)

// Add medicine
func AddMedicine(db *sql.DB, m Medicine) error {
    _, err := db.Exec(
        "INSERT INTO medicines (name, quantity, price, expiry, supplier) VALUES (?, ?, ?, ?, ?)",
        m.Name, m.Quantity, m.Price, m.Expiry.Format("2006-01-02"), m.Supplier,
    )
    return err
}

// Sell medicine 
func SellMedicine(db *sql.DB, id int, qty int) error {
    var price float64
    err := db.QueryRow("SELECT price FROM medicines WHERE id = ?", id).Scan(&price)
    if err != nil {
        return errors.New("❌ medicine not found")
    }

    total := float64(qty) * price
    _, err = db.Exec(
        "INSERT INTO sales (medicine_id, quantity, total, date) VALUES (?, ?, ?, ?)",
        id, qty, total, time.Now().Format("2006-01-02 15:04:05"),
    )
    return err
}

// List medicines 
func ListMedicines(db *sql.DB) {
    rows, err := db.Query("SELECT id, name, quantity, price, supplier, expiry, status FROM medicine_summary")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer rows.Close()

    fmt.Println("\n--- Medicines Summary ---")
    for rows.Next() {
        var id, qty int
        var name, supplier, expiry, status string
        var price float64
        rows.Scan(&id, &name, &qty, &price, &supplier, &expiry, &status)

        fmt.Printf("%d. %s | %d pcs | $%.2f | Exp: %s | %s | %s\n",
            id, name, qty, price, expiry, supplier, status)
    }
}

// List sales
func ListSales(db *sql.DB) {
    rows, err := db.Query(`
        SELECT sales.id, medicines.name, sales.quantity, sales.total, sales.date
        FROM sales
        JOIN medicines ON sales.medicine_id = medicines.id`)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer rows.Close()

    fmt.Println("\n--- Sales ---")
    for rows.Next() {
        var s Sale
        var date string
        rows.Scan(&s.ID, &s.Medicine, &s.Quantity, &s.Total, &date)
        s.Date, _ = time.Parse("2006-01-02 15:04:05", date)
        fmt.Printf("%d. %s x%d = $%.2f (%s)\n", s.ID, s.Medicine, s.Quantity, s.Total, s.Date.Format("2006-01-02"))
    }
}

// Show alerts
func ListAlerts(db *sql.DB) {
    rows, _ := db.Query("SELECT id, message, date FROM alerts ORDER BY id DESC")
    defer rows.Close()

    fmt.Println("\n--- Alerts ---")
    for rows.Next() {
        var a Alert
        var date string
        rows.Scan(&a.ID, &a.Message, &date)
        a.Date, _ = time.Parse("2006-01-02 15:04:05", date)
        fmt.Printf("%d. %s (%s)\n", a.ID, a.Message, a.Date.Format("2006-01-02 15:04"))
    }
}

// Daily sales summary
func DailySalesSummary(db *sql.DB) {
    rows, _ := db.Query("SELECT sale_date, total_items, total_revenue FROM daily_sales_summary")
    defer rows.Close()

    fmt.Println("\n--- Daily Sales Summary ---")
    for rows.Next() {
        var date string
        var items int
        var revenue float64
        rows.Scan(&date, &items, &revenue)
        fmt.Printf("%s | %d items | $%.2f\n", date, items, revenue)
    }
}

// Top selling medicines
func TopSellingMedicines(db *sql.DB) {
    rows, _ := db.Query("SELECT name, total_sold, revenue FROM top_selling_medicines")
    defer rows.Close()

    fmt.Println("\n--- Top Selling Medicines ---")
    for rows.Next() {
        var name string
        var sold int
        var revenue float64
        rows.Scan(&name, &sold, &revenue)
        fmt.Printf("%s | %d sold | $%.2f revenue\n", name, sold, revenue)
    }
}

// Expiring soon
func ExpiringSoon(db *sql.DB) {
    rows, _ := db.Query("SELECT id, name, expiry FROM expiring_soon")
    defer rows.Close()

    fmt.Println("\n--- Expiring Soon ---")
    for rows.Next() {
        var id int
        var name, expiry string
        rows.Scan(&id, &name, &expiry)
        fmt.Printf("%d. %s | Exp: %s\n", id, name, expiry)
    }
}
