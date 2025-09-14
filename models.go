package main

import "time"

type Medicine struct {
    ID       int
    Name     string
    Quantity int
    Price    float64
    Expiry   time.Time
    Supplier string
}
type Sale struct {
    ID        int
    Medicine  string
    Quantity  int
    Total     float64
    Date      time.Time
}
type Alert struct {
    ID      int
    Message string
    Date    time.Time
}
