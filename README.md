# Food Ordering System

> This project was collaboratively developed by a team of students from  
> **National Taiwan University of Science and Technology (NTUST)** – CSIE Department,  
> as part of the **Software Engineering** course project.

## 目錄結構
```
/food-ordering-system

├── /controller        # Handles business logic and route handling
│    ├── user.go           # User-related operations
│    ├── restaurant.go     # Restaurant management
│    ├── rates.go          # Restaurant rating system
│    ├── order.go          # Order placement and tracking
│    └── menu.go           # Menu management
│
├── /models           # Data models and Firebase integration
│    ├── order.go          # Order data model
│    ├── food.go           # Food item data model
│    ├── menu.go           # Menu data model
│    ├── rate.go           # Rating data model
│    ├── restaurant.go     # Restaurant data model
│    └── user_basic.go     # User profile model
│
├── /router           # Route definitions and bindings
│    └── router.go         # Centralized routing
│
├── /utils            # Utility functions and helpers
│    ├── firebase.go       # Firebase initialization
│    └── jwt.go            # JWT token handling
│
├── /config           # Configuration files
│
├── go.mod            # Go module dependency file
└── main.go           # Main entry point (Gin + Firebase)

```

## 如何構建與運行專案

### 1. Install Golang

Make sure Go is installed on your system. You can download the latest version from the [official website](https://golang.org/dl/).

### 2. 設置環境變數

你需要設置 `GOPATH` 來指向 Go 工作空間。如果你的專案位於 `D:\ProgramFile\Go\src\orderfood`，你需要將 `GOPATH` 設置為 `D:\ProgramFile\Go`。

#### Windows 上設置環境變數的步驟：

1. 右鍵 "此電腦"，選擇 "屬性"。
2. 點擊 "進階系統設定"，然後選擇 "環境變數"。
3. 在 "使用者環境變數" 下，新增或編輯 `GOPATH` 變數，設置其值為 `D:\ProgramFile\Go`。
4. 點擊 "確定" 並重新啟動命令行窗口。

### 3. 下載專案依賴項

專案使用 Go Modules 來管理依賴庫，因此你需要下載專案的所有依賴。首先，進入專案的根目錄，然後運行以下命令：

```bash
go mod download
```

## 執行畫面展示
### 主畫面 & 搜尋畫面
![主畫面](screenshots/01.png)

### 登入畫面 & 註冊畫面
![登入畫面](screenshots/02.png)

### 訂單相關畫面
![訂單相關畫面](screenshots/03.png)

### 餐廳資訊畫面
![餐廳資訊畫面](screenshots/04.png)

## 專案報告
- [report.pdf](report.pdf)
<!--
## 我的貢獻
本專案為團隊合作開發，我主要負責後端功能，包括使用者登入註冊、送出訂單、餐廳評價與 API 測試等模組。
-->
