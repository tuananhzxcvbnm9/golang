# Hướng Dẫn Học Golang Toàn Diện (Tiếng Việt)

## 1) Go là gì và khi nào nên dùng?

Go (Golang) là ngôn ngữ lập trình do Google phát triển, tối ưu cho:

- Dịch vụ backend hiệu năng cao.
- Hệ thống microservices.
- Ứng dụng hạ tầng: CLI, DevOps tools, networking services.
- Hệ thống cần concurrency đơn giản, an toàn.

### Điểm mạnh

- Cú pháp gọn, dễ đọc.
- Biên dịch nhanh, binary độc lập.
- Concurrency mạnh với goroutine và channel.
- Thư viện chuẩn tốt, phù hợp cho hệ thống production.

### Khi không phải lựa chọn tốt nhất

- Ứng dụng UI frontend phức tạp.
- Kịch bản cần hệ sinh thái machine learning lớn như Python.

---

## 2) Thiết lập môi trường

### Cài đặt

1. Cài Go từ trang chính thức: `https://go.dev/dl/`
2. Kiểm tra:

```bash
go version
go env
```

### Khởi tạo dự án

```bash
mkdir hello-go && cd hello-go
go mod init example.com/hello-go
```

### Cấu trúc dự án gợi ý

```text
.
├── cmd/                # entry points (app main)
├── internal/           # logic nội bộ
├── pkg/                # thư viện có thể tái sử dụng
├── api/                # contract/API specs
├── configs/            # cấu hình
├── scripts/            # script hỗ trợ
├── test/               # integration test data
├── go.mod
└── go.sum
```

---

## 3) Nền tảng ngôn ngữ Go

## 3.1 Biến, hằng, kiểu dữ liệu

- Biến: `var`, `:=`
- Hằng: `const`
- Kiểu cơ bản: `int`, `float64`, `string`, `bool`, `rune`, `byte`
- Composite: `array`, `slice`, `map`, `struct`

## 3.2 Hàm và package

- Hàm có thể trả nhiều giá trị.
- Hàm có thể trả `error`.
- Quy tắc export: tên viết hoa đầu là public.

## 3.3 Control flow

- `if`, `switch`, `for` (Go chỉ có `for`, nhưng dùng linh hoạt như while).
- `defer` để dọn tài nguyên.

## 3.4 Pointer

Go có pointer nhưng không có pointer arithmetic.

## 3.5 Interface

- Interface trong Go là implicit implementation.
- Nên định nghĩa interface nhỏ, gần nơi sử dụng.

---

## 4) Error handling đúng chuẩn

### Nguyên tắc

- Không bỏ qua lỗi.
- Bọc lỗi bằng `fmt.Errorf("...: %w", err)`.
- Dùng sentinel error / custom error khi cần phân loại.

### Ví dụ

```go
func ReadConfig(path string) ([]byte, error) {
    b, err := os.ReadFile(path)
    if err != nil {
        return nil, fmt.Errorf("read config %s: %w", path, err)
    }
    return b, nil
}
```

---

## 5) Concurrency thực chiến

## 5.1 Goroutine

- Nhẹ hơn thread hệ điều hành.
- Tạo bằng `go fn()`.

## 5.2 Channel

- Đồng bộ giao tiếp giữa goroutine.
- Pattern phổ biến: fan-in, fan-out, worker pool.

## 5.3 Context

- Dùng `context.Context` cho timeout, cancellation, request-scope values.
- Luôn truyền `ctx` là tham số đầu tiên trong hàm xử lý nghiệp vụ/request.

### Ví dụ worker pool đơn giản

```go
type Job struct{ ID int }
type Result struct{ ID int; Err error }

func worker(ctx context.Context, jobs <-chan Job, results chan<- Result) {
    for {
        select {
        case <-ctx.Done():
            return
        case j, ok := <-jobs:
            if !ok {
                return
            }
            // xử lý công việc
            results <- Result{ID: j.ID, Err: nil}
        }
    }
}
```

---

## 6) Testing trong Go

### 6.1 Unit test

- File test tên `*_test.go`
- Dùng package `testing`
- Chạy: `go test ./...`

### 6.2 Table-driven tests

Thực hành tốt nhất để kiểm tra nhiều case gọn, dễ đọc.

### 6.3 Benchmark

- Hàm benchmark có dạng `BenchmarkXxx(b *testing.B)`
- Chạy: `go test -bench=. ./...`

### 6.4 Race detector

- Chạy: `go test -race ./...`
- Bắt data race trong code concurrent.

---

## 7) Thiết kế API/Service với Go

### 7.1 HTTP service cơ bản

- Dùng `net/http` hoặc framework nhẹ.
- Áp dụng middleware cho logging, auth, recover, tracing.

### 7.2 Clean architecture (gợi ý)

- `handler`: nhận request/response.
- `service/usecase`: nghiệp vụ.
- `repository`: truy cập dữ liệu.
- `domain`: model và rule lõi.

### 7.3 Cấu hình

- Đọc từ env (12-factor app).
- Có default hợp lý, validate lúc startup.

---

## 8) Cơ sở dữ liệu và transaction

- Dùng `database/sql` + driver.
- Luôn có timeout context khi query.
- Bọc transaction trong hàm nghiệp vụ.
- Không giữ transaction lâu hơn cần thiết.

---

## 9) Quan sát hệ thống (Observability)

- Logging có cấu trúc (JSON).
- Metrics (Prometheus).
- Tracing (OpenTelemetry).
- Health checks: liveness/readiness.

---

## 10) Bảo mật căn bản

- Validate input đầu vào.
- Không hard-code secrets.
- Dùng TLS trong môi trường production.
- Thiết lập timeout cho server/client để tránh treo tài nguyên.
- Chặn panic lan rộng bằng middleware recover.

---

## 11) Tối ưu hiệu năng

- Dùng benchmark trước khi tối ưu.
- Tránh cấp phát không cần thiết.
- Tận dụng profiler:

```bash
go test -bench=. -benchmem ./...
go tool pprof
```

- Kiểm tra memory/CPU hotspot trước khi refactor.

---

## 12) Lộ trình học Go 12 tuần

### Giai đoạn 1: Cơ bản (Tuần 1-3)

- Cú pháp, kiểu dữ liệu, hàm, struct, interface.
- Bài tập: CRUD in-memory bằng map/slice.

### Giai đoạn 2: Trung cấp (Tuần 4-6)

- Error handling, package design, testing, modules.
- Bài tập: xây REST API nhỏ có unit test.

### Giai đoạn 3: Concurrency & I/O (Tuần 7-9)

- Goroutine, channel, context, worker pool.
- Bài tập: xử lý batch jobs có timeout/retry.

### Giai đoạn 4: Production mindset (Tuần 10-12)

- Logging, metrics, tracing, Docker, CI/CD, security.
- Bài tập: triển khai dịch vụ hoàn chỉnh với monitoring.

---

## 13) Checklist tự đánh giá

- [ ] Hiểu rõ slice vs array vs map.
- [ ] Viết được interface nhỏ, đúng ngữ cảnh.
- [ ] Biết bọc/làm giàu lỗi và kiểm tra bằng `errors.Is/As`.
- [ ] Viết unit test + table test + benchmark.
- [ ] Sử dụng context chuẩn trong API/database.
- [ ] Tạo được HTTP service với middleware và graceful shutdown.
- [ ] Biết dùng race detector và pprof để phân tích lỗi/hiệu năng.

---

## 14) Tài nguyên khuyến nghị

- Tài liệu Go chính thức: https://go.dev/doc/
- Effective Go: https://go.dev/doc/effective_go
- Go by Example: https://gobyexample.com/
- Tour of Go: https://go.dev/tour/
- Checklist phỏng vấn Go (repo này): docs/checklist-phong-van-go.md
- Từ điển thuật ngữ (repo này): docs/thuat-ngu-go.md


---

## 15) Gợi ý học tiếp theo (sau 12 tuần)

### Nhánh A: Backend API chuyên sâu

- Tập trung tối ưu truy vấn database, transaction, và cache.
- Áp dụng distributed tracing cho luồng request đa service.
- Thực hành thiết kế idempotency cho API ghi dữ liệu.

### Nhánh B: Platform/Infra với Go

- Viết CLI nội bộ (deploy, migration, log tools).
- Xây custom exporter/collector cho metrics.
- Tìm hiểu sâu về networking trong `net`, `http`, `context`.

### Nhánh C: Reliability/SRE mindset

- Thiết kế cơ chế retry/backoff/circuit-breaker.
- Đặt SLI/SLO đơn giản cho service.
- Viết runbook xử lý sự cố và diễn tập incident nhỏ.

### Mục tiêu 30-60-90 ngày

- 30 ngày: ship 1 service nhỏ có test + CI.
- 60 ngày: thêm observability đầy đủ + benchmark.
- 90 ngày: tách module hợp lý, cải thiện reliability theo dữ liệu thực tế.


---

## 16) Lỗi thường gặp khi học Go và cách tránh

- **Lạm dụng goroutine**: tạo quá nhiều goroutine không có cơ chế dừng -> luôn dùng context hoặc channel đóng để shutdown.
- **Quên `defer rows.Close()`**: dễ rò rỉ kết nối DB -> đóng tài nguyên ngay sau khi kiểm tra lỗi.
- **Bỏ qua lỗi trả về**: khiến bug âm thầm -> xử lý hoặc bọc lỗi ngay tại chỗ.
- **Interface quá lớn**: khó mock/test -> tách interface nhỏ theo use case.
- **Tối ưu sớm**: làm code phức tạp -> đo benchmark/profile trước khi tối ưu.

## 17) Mẫu lịch học hàng tuần (8 giờ/tuần)

- **2 giờ đọc**: học 1-2 mục lý thuyết trong tài liệu chính.
- **4 giờ code**: triển khai bài tập tương ứng.
- **1 giờ test/benchmark**: thêm test case và chạy benchmark/race detector.
- **1 giờ review**: tự review hoặc peer review theo checklist.
