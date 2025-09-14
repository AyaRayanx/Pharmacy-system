package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "time"
)

func main() {
    db := OpenDB()
    defer db.Close()

    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Println("\n--- Pharmacy Tracker ---")
        fmt.Println("1. Add Medicine")
        fmt.Println("2. List Medicines")
        fmt.Println("3. Sell Medicine")
        fmt.Println("4. List Sales")
        fmt.Println("5. Show Alerts")
        fmt.Println("6. Daily Sales Summary")
        fmt.Println("7. Top Selling Medicines")
        fmt.Println("8. Expiring Soon Report")
        fmt.Println("0. Exit")
        fmt.Print("Choose: ")
        choice, _ := reader.ReadString('\n')
        choice = strings.TrimSpace(choice)

        switch choice {
        case "1":
            fmt.Print("Medicine Name: ")
            name, _ := reader.ReadString('\n')
            name = strings.TrimSpace(name)

            fmt.Print("Quantity: ")
            qtyStr, _ := reader.ReadString('\n')
            qty, _ := strconv.Atoi(strings.TrimSpace(qtyStr))

            fmt.Print("Price: ")
            priceStr, _ := reader.ReadString('\n')
            price, _ := strconv.ParseFloat(strings.TrimSpace(priceStr), 64)

            fmt.Print("Expiry (YYYY-MM-DD): ")
            expStr, _ := reader.ReadString('\n')
            expiry, _ := time.Parse("2006-01-02", strings.TrimSpace(expStr))

            fmt.Print("Supplier: ")
            supp, _ := reader.ReadString('\n')
            supp = strings.TrimSpace(supp)

            AddMedicine(db, Medicine{Name: name, Quantity: qty, Price: price, Expiry: expiry, Supplier: supp})
            fmt.Println("✅ Medicine added!")

        case "2":
            ListMedicines(db)

        case "3":
            fmt.Print("Enter Medicine ID: ")
            idStr, _ := reader.ReadString('\n')
            id, _ := strconv.Atoi(strings.TrimSpace(idStr))

            fmt.Print("Enter Quantity: ")
            qtyStr, _ := reader.ReadString('\n')
            qty, _ := strconv.Atoi(strings.TrimSpace(qtyStr))

            err := SellMedicine(db, id, qty)
            if err != nil {
                fmt.Println("❌ Error:", err)
            } else {
                fmt.Println("✅ Sale recorded!")
            }

        case "4":
            ListSales(db)

        case "5":
            ListAlerts(db)

        case "6":
            DailySalesSummary(db)

        case "7":
            TopSellingMedicines(db)

        case "8":
            ExpiringSoon(db)

        case "0":
            fmt.Println("Exiting the system. Have a nice day.")
            return
        }
    }
}
