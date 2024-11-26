package converter

import (
	"github.com/vitaliysev/mts_go_project/internal/booking/model"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

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
		PeriodUse: info.Period_use,
		HotelId:   info.Hotel_id,
	}
}

func ToBookInfoFromDesc(info *desc.BookInfo) *model.BookInfo {
	return &model.BookInfo{
		Period_use: info.PeriodUse,
		Hotel_id:   info.HotelId,
	}
}

func ToUpdateBookInfoFromDesc(info *desc.UpdateBookInfo) *model.BookInfo {
	return &model.BookInfo{
		Period_use: unwrapStringValue(info.GetPeriodUse()),
		Hotel_id:   info.GetHotelId().Value,
	}
}

func unwrapStringValue(value *wrapperspb.StringValue) string {
	if value != nil {
		return value.Value
	}
	return ""
}
