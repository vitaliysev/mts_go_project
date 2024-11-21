package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/vitaliysev/mts_go_project/internal/model"
	desc "github.com/vitaliysev/mts_go_project/pkg/booking_v1"
)

func ToBookFromService(book *model.Book) *desc.Book {
	var updatedAt *timestamppb.Timestamp
	if book.UpdatedAt.Valid {
		updatedAt = timestamppb.New(book.UpdatedAt.Time)
	}

	return &desc.Book{
		Id:        book.ID,
		Info:      ToBookInfoFromService(book.Info),
		CreatedAt: timestamppb.New(book.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToBookInfoFromService(info model.BookInfo) *desc.BookInfo {
	return &desc.BookInfo{
		Title:     info.Title,
		PeriodUse: info.Period_use,
	}
}

func ToBookInfoFromDesc(info *desc.BookInfo) *model.BookInfo {
	return &model.BookInfo{
		Title:      info.Title,
		Period_use: info.PeriodUse,
	}
}

func ToUpdateBookInfoFromDesc(info *desc.UpdateBookInfo) *model.BookInfo {
	return &model.BookInfo{
		Title:      unwrapStringValue(info.GetTitle()),
		Period_use: unwrapStringValue(info.GetPeriodUse()),
	}
}

func unwrapStringValue(value *wrapperspb.StringValue) string {
	if value != nil {
		return value.Value
	}
	return ""
}
