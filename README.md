# CLY TodoApp Backend

API Backend cho ứng dụng quản lý công việc (Todo App) được xây dựng bằng Golang với Gin framework.

## 🚀 Demo

- **Live API**: [34.171.223.47:8000/api/v1/tasks](34.171.223.47:8000/api/v1/tasks) hoặc [cty-todo-app-be.duckdns.org/api/v1/tasks](cty-todo-app-be.duckdns.org/api/v1/tasks)
- **Frontend Demo**: [https://cly-todo-app-fe.vercel.app/](https://cly-todo-app-fe.vercel.app/)

## ✨ Tính năng

- ✅ RESTful API cho quản lý tasks
- ✅ CORS configuration
- ✅ Environment configuration
- ✅ Docker containerization
- ✅ Nginx reverse proxy với SSL
- ✅ Infrastructure as Code với Terraform

## 🛠️ Công nghệ sử dụng

- **Golang 1.25** - Programming language
- **Gin** - HTTP web framework
- **Docker** - Containerization
- **Nginx** - Reverse proxy & SSL termination
- **Terraform** - Infrastructure as Code
- **Google Cloud Platform** - Cloud hosting

## 📦 Cài đặt và chạy locally

### Yêu cầu hệ thống
- Go 1.25 trở lên
- Docker & Docker Compose
- Git

### Bước 1: Clone repository
```bash
git clone <repository-url>
cd CLY_TodoApp_BE
```

### Bước 2: Cài đặt dependencies
```bash
go mod download
```

### Bước 3: Cấu hình environment variables
Tạo file `.env` trong thư mục root từ template:

```bash
# Copy từ template
cp .env.template .env

# Edit file .env
nano .env
```

Cấu hình `.env`:
```bash
PORT=8000
HOST=localhost
ENVIRONMENT=development

CORS_ALLOWED_ORIGINS=http://localhost:3000
# Production: CORS_ALLOWED_ORIGINS=http://localhost:3000,https://cly-todo-app-fe.vercel.app

LOG_LEVEL=info
```

### Bước 4: Chạy ứng dụng

#### Development mode
```bash
go run cmd/server/main.go
```

#### Sử dụng Docker
```bash
# Build image
docker build -t cly-todoapp-backend .

# Run container
docker run -p 8000:8000 --env-file .env cly-todoapp-backend
```

#### Sử dụng Docker Compose (với Nginx)
```bash
# Cấu hình nginx environment
cp _nginx/.env.template _nginx/.env
# Edit _nginx/.env với domain và email của bạn

# Chạy services
docker-compose up -d
```

API sẽ chạy tại `http://localhost:8000`.

## 🌐 Triển khai lên Google Cloud Platform

### Option 1: Triển khai thủ công

#### Bước 1: Tạo VM trên GCP
1. Truy cập [Google Cloud Console](https://console.cloud.google.com/)
2. Tạo project mới hoặc chọn project hiện có
3. Tạo Compute Engine instance:
   - Machine type: e2-medium
   - OS: Ubuntu 22.04 LTS
   - Boot disk: 10GB SSD
   - Firewall: Allow HTTP và HTTPS traffic

#### Bước 2: Cấu hình firewall
```bash
# Tạo firewall rules cho port 8000
gcloud compute firewall-rules create allow-app-port-8000 \
    --allow tcp:8000 \
    --source-ranges 0.0.0.0/0 \
    --description "Allow port 8000 for TodoApp backend"

# Đồng thời cấu hình cho port 80 và 443
```

#### Bước 3: Connect SSH và setup server
```bash
# SSH vào VM
gcloud compute ssh your-instance-name

# Update system
sudo apt update && sudo apt upgrade -y

# Install Git
sudo apt install git -y

# Install Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
sudo usermod -aG docker $USER

# Install Docker Compose
sudo curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

# Logout và login lại để apply docker group
exit
```

#### Bước 4: Deploy application
```bash
# SSH lại vào VM
gcloud compute ssh your-instance-name

# Clone repository
git clone <your-repository-url>
cd CLY_TodoApp_BE

# Cấu hình environment variables
cp .env.template .env
nano .env  # Edit với cấu hình production

# Production .env example:
# PORT=8000
# HOST=0.0.0.0
# ENVIRONMENT=production
# CORS_ALLOWED_ORIGINS=https://cly-todo-app-fe.vercel.app,https://your-domain.com
# LOG_LEVEL=warn

# Cấu hình nginx
cp _nginx/.env.template _nginx/.env
nano _nginx/.env  # Edit với domain và email của bạn

# Deploy với Docker Compose
docker-compose up -d

# Kiểm tra logs
docker-compose logs -f
```

### Option 2: Triển khai tự động với Terraform

#### Bước 1: Chuẩn bị Service Account
1. Truy cập [Google Cloud Console](https://console.cloud.google.com/)
2. Tạo Service Account:
   - IAM & Admin → Service Accounts
   - Create Service Account
   - Assign roles: Compute Admin, Service Account User
3. Tạo JSON key và download về
4. Rename file thành `iam.json` và copy vào thư mục `_terraform/`

#### Bước 2: Cấu hình Terraform
```bash
cd _terraform

# Copy template và cấu hình
cp terraform.tfvars.template terraform.tfvars
cp iam.json.template iam.json  # (Đã có từ bước 1)

# Edit terraform.tfvars
nano terraform.tfvars
```

Cấu hình `terraform.tfvars`:
```hcl
gcp_svc_key    = "./iam.json"
gcp_project    = "your-gcp-project-id"
gcp_region     = "asia-southeast1"  # Hoặc region gần bạn

ssh_public_key = "username:ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAI... your-email@example.com"
```

#### Bước 3: Generate SSH Key (nếu chưa có)
```bash
# Linux/Mac
ssh-keygen -t ed25519 -C "your-email@example.com"
cat ~/.ssh/id_ed25519.pub

# Windows (PowerShell)
ssh-keygen -t ed25519 -C "your-email@example.com"
Get-Content "$env:USERPROFILE\.ssh\id_ed25519.pub"
```

#### Bước 4: Deploy infrastructure
```bash
terraform init

terraform plan

terraform apply -auto-approve

# Lấy IP của VM được tạo
terraform output
```

#### Bước 5: Deploy application trên VM được tạo
```bash
ssh -i ~/.ssh/id_ed25519 username@YOUR_EXTERNAL_IP
```

## 🔧 Cấu hình Domain và SSL

### Bước 1: Gắn domain vào IP
1. Truy cập domain registrar (Namecheap, GoDaddy, etc.)
2. Tạo A record trỏ về external IP của VM:
   ```
   Type: A
   Name: api (hoặc subdomain bạn muốn)
   Value: YOUR_EXTERNAL_IP
   TTL: 300
   ```

### Bước 2: Cấu hình SSL với Let's Encrypt
```bash
# SSH vào VM
ssh username@your-domain.com

# Edit nginx environment
cd CLY_TodoApp_BE
nano _nginx/.env

# Cập nhật:
DOMAIN=api.your-domain.com
LETSENCRYPT_EMAIL=your-email@example.com

# Restart để apply SSL
docker-compose down
docker-compose up -d

# Kiểm tra logs
docker-compose logs nginx
```

## 📝 API Endpoints

```
GET    /api/tasks       # Lấy danh sách tasks
POST   /api/tasks       # Tạo task mới
PUT    /api/tasks/:id   # Cập nhật task
DELETE /api/tasks/:id   # Xóa task
GET    /health          # Health check
```

### Example Request/Response

#### GET /api/tasks
```json
{
  "status": "success",
  "data": [
    {
      "id": 1,
      "title": "Sample Task",
      "description": "Task description",
      "completed": false,
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ]
}
```

#### POST /api/tasks
```json
// Request
{
  "title": "New Task",
  "description": "Task description"
}

// Response
{
  "status": "success",
  "data": {
    "id": 2,
    "title": "New Task",
    "description": "Task description",
    "completed": false,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

## 📁 Cấu trúc thư mục

```
CLY_TodoApp_BE/
├── cmd/
│   └── server/
│       └── main.go          # Entry point
├── internal/
│   ├── config/
│   │   └── config.go        # Configuration
│   ├── handlers/
│   │   └── task_handlers.go # Request handlers
│   ├── middleware/          # Custom middleware
│   ├── models/
│   │   └── task.go         # Data models
│   └── routes/
│       ├── routes.go       # Route setup
│       └── task_routes.go  # Task routes
├── _nginx/
│   ├── Dockerfile          # Nginx Docker config
│   ├── nginx.conf.template # Nginx configuration
│   ├── init-ssl.sh        # SSL initialization
│   └── start-nginx.sh     # Startup script
├── _terraform/
│   ├── main.tf            # Infrastructure definition
│   ├── provider.tf        # Provider configuration
│   ├── variable.tf        # Variable definitions
│   ├── terraform.tfvars.template
│   └── iam.json.template  # Service account template
├── docker-compose.yml     # Multi-container setup
├── Dockerfile            # Backend container
├── go.mod               # Go dependencies
└── .env.template        # Environment template
```

## 🐛 Troubleshooting

### Lỗi kết nối từ Frontend
- Kiểm tra CORS configuration trong `.env`
- Đảm bảo firewall rules cho port 8000
- Kiểm tra domain đã trỏ đúng IP

### SSL Certificate errors
- Kiểm tra domain đã trỏ đúng IP
- Verify email và domain trong `_nginx/.env`
- Check nginx logs: `docker-compose logs nginx`

### Terraform deployment fails
- Verify service account permissions
- Check `iam.json` file exists và đúng format
- Ensure SSH public key format đúng

### Docker issues
- Check disk space: `df -h`
- Restart Docker: `sudo systemctl restart docker`
- Clean up: `docker system prune -a`

## 🔧 Environment Variables

### Backend (.env)
| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | `8000` |
| `HOST` | Server host | `localhost` |
| `ENVIRONMENT` | Environment mode | `development` |
| `CORS_ALLOWED_ORIGINS` | Allowed origins for CORS | `http://localhost:3000` |
| `LOG_LEVEL` | Logging level | `info` |

### Nginx (_nginx/.env)
| Variable | Description | Required |
|----------|-------------|----------|
| `DOMAIN` | Backend domain | Yes |
| `LETSENCRYPT_EMAIL` | Email for SSL cert | Yes |

### Terraform (terraform.tfvars)
| Variable | Description | Required |
|----------|-------------|----------|
| `gcp_project` | GCP Project ID | Yes |
| `gcp_region` | GCP Region | Yes |
| `ssh_public_key` | SSH public key | Yes |

## 👨‍💻 Author

- **Zennisch** - [GitHub Profile](https://github.com/Zennisch)

## 🔗 Links

- **Frontend Repository**: [https://github.com/Zennisch/CLY_TodoApp_FE](https://github.com/Zennisch/CLY_TodoApp_FE)
- **Frontend Demo**: [https://cly-todo-app-fe.vercel.app/](https://cly-todo-app-fe.vercel.app/)
- **Postman Collection**: `CLY_TodoApp.postman_collection.json`
