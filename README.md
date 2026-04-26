# Golang Study Guide (Tiếng Việt)

Kho tài liệu này cung cấp bộ hướng dẫn học Go bằng tiếng Việt, theo lộ trình từ cơ bản đến nâng cao.

## Nội dung chính

- [Hướng dẫn học Go toàn diện](docs/huong-dan-go-viet-nam.md)
- [Bài tập thực hành Go](docs/bai-tap-go.md)
- [Checklist ôn phỏng vấn Go](docs/checklist-phong-van-go.md)
- [Từ điển thuật ngữ Go](docs/thuat-ngu-go.md)

## Mục tiêu

- Giúp người mới bắt đầu học Go có lộ trình rõ ràng.
- Cung cấp tài liệu tham khảo nhanh cho người đã có kinh nghiệm.
- Tập trung vào thực hành, tư duy thiết kế, và các chuẩn production.

## Cách sử dụng

1. Bắt đầu với phần **Lộ trình 12 tuần** trong tài liệu chính.
2. Sau mỗi chương, làm bài tập tương ứng trong `docs/bai-tap-go.md`.
3. Dùng checklist cuối tài liệu để tự đánh giá năng lực.

## Đối tượng phù hợp

- Sinh viên CNTT.
- Lập trình viên backend muốn chuyển sang Go.
- Kỹ sư cần xây dựng dịch vụ hiệu năng cao bằng Go.

## Gợi ý học tiếp

- Hoàn thành bài tập 1-3 trước để nắm chắc nền tảng.
- Chuyển sang bài tập 4-6 khi đã quen testing và concurrency.
- Cuối cùng, xây 1 dự án capstone (dịch vụ thật có log/metrics/test/deploy).
- Nếu học theo nhóm, áp dụng code review theo checklist trong tài liệu chính.


## Lộ trình đọc nhanh theo mục tiêu

- **Mới bắt đầu**: đọc mục 1-6 trong tài liệu chính, sau đó làm Bài tập 1-2.
- **Backend đã có kinh nghiệm**: ưu tiên mục 5-11, rồi làm Bài tập 3-7.
- **Chuẩn bị phỏng vấn Go**: ôn checklist + tự triển khai lại Capstone từ đầu trong 1 tuần.


## Chạy thử nhanh

```bash
go test ./...
```

Repo hiện có ví dụ Go tối thiểu ở `examples/basic` để bạn thực hành test ngay.

- `examples/advanced`: ví dụ concurrency nâng cao với worker pool + context cancellation.
