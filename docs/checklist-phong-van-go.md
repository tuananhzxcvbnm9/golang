# Checklist Ôn Phỏng Vấn Golang (Tiếng Việt)

## 1) Ngôn ngữ nền tảng

- [ ] Phân biệt `array` vs `slice` (len/cap, copy, append, reallocation).
- [ ] Hiểu `map` là reference type, chú ý zero-value và kiểm tra key tồn tại.
- [ ] Nắm pointer trong Go và vì sao không có pointer arithmetic.
- [ ] Nắm method set của value receiver vs pointer receiver.
- [ ] Hiểu interface là implicit, cách tránh interface quá lớn.

## 2) Concurrency

- [ ] Viết được worker pool với goroutine + channel.
- [ ] Hiểu khi nào dùng buffered channel.
- [ ] Biết xử lý cancellation với `context.Context`.
- [ ] Nhận diện deadlock, goroutine leak, data race.
- [ ] Biết dùng `go test -race ./...` để bắt race condition.

## 3) Error handling và thiết kế API

- [ ] Luôn kiểm tra lỗi trả về; bọc lỗi bằng `%w`.
- [ ] Dùng `errors.Is/As` đúng ngữ cảnh.
- [ ] Thiết kế HTTP status code nhất quán.
- [ ] Validate input và trả lỗi có cấu trúc.
- [ ] Có timeout cho server/client/database.

## 4) Testing

- [ ] Viết table-driven tests cho logic chính.
- [ ] Viết benchmark cơ bản và đọc `ns/op`, `B/op`, `allocs/op`.
- [ ] Biết phân biệt unit test và integration test.
- [ ] Có test cho edge cases và error paths.

## 5) System design mini (Go backend)

- [ ] Trình bày được kiến trúc handler/service/repository.
- [ ] Nêu được chiến lược transaction và idempotency.
- [ ] Mô tả logging/metrics/tracing căn bản.
- [ ] Mô tả graceful shutdown và healthcheck.

## 6) Bài tập phỏng vấn gợi ý (60-90 phút)

1. **Implement rate limiter in-memory** (token bucket đơn giản).
2. **Build concurrent URL checker** với context timeout và worker pool.
3. **Design tiny order API** với idempotency key cho tạo đơn.
4. **Refactor code có race/deadlock** và giải thích root cause.

## 7) Cách tự đánh giá trước phỏng vấn

- Nếu giải trọn vẹn 3/4 bài tập ở mục 6 trong thời gian giới hạn, bạn ở mức sẵn sàng khá tốt.
- Nếu còn chậm ở concurrency, ưu tiên luyện mục 2 + chạy race detector thường xuyên.
- Nếu yếu phần API design, luyện thêm capstone trong `docs/bai-tap-go.md`.
