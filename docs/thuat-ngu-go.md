# Từ Điển Thuật Ngữ Golang (Tiếng Việt)

## Mục tiêu

Tài liệu ngắn giúp quy đổi thuật ngữ Go từ tiếng Anh sang ngữ cảnh dễ hiểu bằng tiếng Việt.

## Thuật ngữ cốt lõi

- **Goroutine**: luồng thực thi nhẹ do runtime Go quản lý.
- **Channel**: kênh giao tiếp đồng bộ/bất đồng bộ giữa goroutine.
- **Context**: cơ chế truyền timeout/cancel/deadline xuyên suốt call chain.
- **Interface**: tập hành vi (method set) mà type có thể triển khai ngầm định.
- **Receiver**: tham số đặc biệt gắn method với type.
- **Zero value**: giá trị mặc định của một kiểu khi chưa khởi tạo.
- **Escape analysis**: phân tích của compiler để quyết định biến nằm stack hay heap.
- **Race condition**: lỗi truy cập dữ liệu đồng thời gây kết quả không xác định.
- **Deadlock**: các goroutine chờ nhau vô hạn, chương trình không tiến triển.
- **Idempotency**: gọi lặp lại cùng request vẫn cho cùng trạng thái kết quả.

## Thuật ngữ testing/performance

- **Table-driven test**: kỹ thuật gom nhiều case test vào bảng dữ liệu.
- **Benchmark**: đo hiệu năng hàm bằng `testing.B`.
- **pprof**: bộ công cụ profiling CPU/memory/goroutine.
- **allocs/op**: số lần cấp phát bộ nhớ trung bình trên mỗi operation benchmark.
- **Race detector**: công cụ runtime phát hiện data race (`go test -race`).

## Thuật ngữ vận hành hệ thống

- **Graceful shutdown**: dừng service có kiểm soát, xử lý nốt request đang chạy.
- **Liveness probe**: kiểm tra process còn sống.
- **Readiness probe**: kiểm tra service đã sẵn sàng nhận traffic.
- **SLI/SLO**: chỉ số đo độ tin cậy và mục tiêu cam kết dịch vụ.
- **Runbook**: tài liệu quy trình xử lý sự cố vận hành.
