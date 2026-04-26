# Bài Tập Golang Thực Hành (Tiếng Việt)

## Mục tiêu

- Củng cố tư duy thiết kế và kỹ năng triển khai Go.
- Chuyển từ “biết cú pháp” sang “xây được dịch vụ production-ready”.

## Bài tập 1: CLI quản lý công việc

### Yêu cầu

- Dùng `flag` để parse command.
- Lưu dữ liệu dạng JSON file cục bộ.
- Hỗ trợ lệnh: `add`, `list`, `done`, `remove`.

### Kỹ năng luyện tập

- Struct, slice, file I/O, error handling.

---

## Bài tập 2: URL shortener in-memory

### Yêu cầu

- REST API với `net/http`.
- Endpoint tạo short URL và redirect.
- Validate input + trả mã lỗi chuẩn.

### Kỹ năng luyện tập

- Routing, JSON encode/decode, handler design.

---

## Bài tập 3: Worker pool xử lý ảnh giả lập

### Yêu cầu

- Tạo danh sách job và xử lý bằng goroutine workers.
- Giới hạn số worker.
- Có context timeout toàn cục.

### Kỹ năng luyện tập

- Goroutine, channel, context, cancellation.

---

## Bài tập 4: REST API CRUD + PostgreSQL

### Yêu cầu

- Dùng `database/sql`.
- CRUD đầy đủ cho entity bất kỳ.
- Có migration và transaction cho thao tác nhiều bước.

### Kỹ năng luyện tập

- Query parameterized, transaction safety, repository pattern.

---

## Bài tập 5: Hardening production

### Yêu cầu

- Thêm structured logging.
- Thêm metrics endpoint.
- Graceful shutdown.
- Dockerfile + healthcheck.

### Kỹ năng luyện tập

- Vận hành production, observability, reliability.

---

## Bài tập 6: Tối ưu hiệu năng

### Yêu cầu

- Viết benchmark cho một hàm xử lý dữ liệu.
- So sánh 2-3 cách cài đặt.
- Đưa ra kết luận dựa trên số liệu.

### Kỹ năng luyện tập

- Benchmark, profiling, tối ưu dựa trên dữ liệu.

---



### So sánh worker counts (gợi ý benchmark)

Dùng benchmark của `examples/advanced/parallel_test.go` để so sánh tradeoff theo số worker:

```bash
go test -bench=BenchmarkParallelSquareWorkerCounts -benchmem ./examples/advanced
```

Đọc kết quả:

- `ns/op`: thời gian xử lý trung bình mỗi lần chạy.
- `allocs/op`: số lần cấp phát trung bình.
- So sánh `workers_1`, `workers_2`, `workers_4`, `workers_8` để tìm điểm cân bằng giữa overhead goroutine và tốc độ.

## Rubric tự chấm điểm

- Đúng chức năng: 40%
- Chất lượng code: 20%
- Error handling & edge case: 15%
- Test coverage & test quality: 15%
- Vận hành (log/metrics/shutdown): 10%

## Nâng cấp thêm

- Áp dụng CI chạy `go test ./...` và `go test -race ./...`.
- Viết Makefile chuẩn hóa lệnh build/test/lint.
- Tổ chức cấu trúc thư mục theo module lớn khi hệ thống tăng trưởng.

---

## Bài tập 7 (Capstone): Dịch vụ production mini

### Yêu cầu

- Chọn 1 domain thực tế (task, billing, notification, inventory...).
- Có ít nhất 5 endpoint CRUD + 1 endpoint thống kê.
- Có authentication cơ bản (API key/JWT) ở mức demo.
- Có unit test cho service layer và integration test cho API chính.
- Có dashboard metrics tối thiểu (request count, latency, error rate).

### Tiêu chí hoàn thành

- Chạy local bằng 1 lệnh (`make run` hoặc tương đương).
- Có tài liệu kiến trúc ngắn (1-2 trang).
- Có checklist release: config, migration, rollback, healthcheck.


---

## Gợi ý cách nộp bài (portfolio)

- Mỗi bài tập để trong 1 repo hoặc 1 thư mục riêng, có `README` mô tả mục tiêu và cách chạy.
- Chụp kết quả test (`go test ./...`) và benchmark chính để làm bằng chứng năng lực.
- Với capstone, thêm sơ đồ kiến trúc và phần "trade-offs" để thể hiện tư duy kỹ thuật.
