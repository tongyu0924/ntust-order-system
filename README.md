## 目錄結構

/food-ordering-system

├── /controllers  # 控制器：處理業務邏輯與路由的映射  
│    ├── orderController.go  # 點餐相關的業務邏輯  
│    └── userController.go  # 用戶相關的業務邏輯  
│  
├── /models  # 模型：資料結構和 Firebase 交互定義  
│    ├── order.go  # 點餐的數據模型  
│    └── user.go  # 用戶的數據模型  
│  
├── /router  # 路由：定義 HTTP 路徑與控制器的映射  
│    └── router.go  # 所有路由的統一管理  
│  
├── /services  # 服務：處理具體功能邏輯，如 Firebase 操作  
│    ├── firebaseService.go  # Firebase 的初始化和操作  
│    └── orderService.go  # 點餐邏輯（如創建訂單、查詢訂單）  
│  
├── /utils  # 公用工具：放置一些工具函數或通用邏輯  
│    └── response.go  # 標準化的 API 回應結構  
│  
├── /config  # 配置：存放專案相關的配置文件  
│    └── firebaseConfig.go  # Firebase 的配置初始化  
│  
├── /views  # 前端視圖：HTML 或其他模板文件（如果需要網頁前端）  
│    └── index.html  # 首頁或其他視圖  
│  
├── go.mod  # Go module 文件，管理依賴項  
└── main.go  # 主文件，啟動 Gin 框架和 Firebase


## 如何構建與運行專案

### 1. 安裝 Golang

請先確保已安裝 Golang，可以從 [Golang 官網](https://golang.org/dl/) 下載並安裝最新版本。

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
