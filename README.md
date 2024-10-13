/food-ordering-system
│
├── /controllers           # 控制器：處理業務邏輯與路由的映射
│   ├── orderController.go  # 點餐相關的業務邏輯
│   └── userController.go   # 用戶相關的業務邏輯
│
├── /models                # 模型：資料結構和 Firebase 交互定義
│   ├── order.go           # 點餐的數據模型
│   └── user.go            # 用戶的數據模型
│
├── /router               # 路由：定義 HTTP 路徑與控制器的映射
│   └── router.go          # 所有路由的統一管理
│
├── /services              # 服務：處理具體功能邏輯，如 Firebase 操作
│   ├── firebaseService.go  # Firebase 的初始化和操作
│   └── orderService.go     # 點餐邏輯（如創建訂單、查詢訂單）
│
├── /utils                 # 公用工具：放置一些工具函數或通用邏輯
│   └── response.go        # 標準化的 API 回應結構
│
├── /config                # 配置：存放專案相關的配置文件
│   └── firebaseConfig.go  # Firebase 的配置初始化
│
├── /views                 # 前端視圖：HTML 或其他模板文件（如果需要網頁前端）
│   └── index.html         # 首頁或其他視圖
│
├── go.mod                 # Go module 文件，管理依賴項
└── main.go                # 主文件，啟動 Gin 框架和 Firebase

如何Build 這個GO專案?
先裝好Golang，然後接下來要配置環境變數，將GO_PATH配置到此專案資料夾位置，假設此檔案(Readme.md)在D:\ProgramFile\Go\src\orderfood，則GO_PATH就設定成D:\ProgramFile\Go