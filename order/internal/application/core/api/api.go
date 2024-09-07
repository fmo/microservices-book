package api

import (
	"context"
	"github.com/fmo/microservices-book/order/internal/application/core/domain"
	"github.com/fmo/microservices-book/order/internal/ports"
)

type Application struct {
	db ports.DBPort
	//payment ports.PaymentPort
}

//func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
//	return &Application{
//		db: db,
//		payment: payment,
//	}
//}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (a Application) PlaceOrder(ctx context.Context, order domain.Order) (domain.Order, error) {
	err := a.db.Save(ctx, &order)
	if err != nil {
		return domain.Order{}, err
	}
	//paymentErr := a.payment.Charge(ctx, &order)
	//if paymentErr != nil {
	//	st, _ := status.FromError(paymentErr)
	//	fieldErr := &errdetails.BadRequest_FieldViolation{
	//		Field:       "payment",
	//		Description: st.Message(),
	//	}
	//	badReq := &errdetails.BadRequest{}
	//	badReq.FieldViolations = append(badReq.FieldViolations, fieldErr)
	//	orderStatus := status.New(codes.InvalidArgument, "order creation failed")
	//	statusWithDetails, _ := orderStatus.WithDetails(badReq)
	//	return domain.Order{}, statusWithDetails.Err()
	//}
	return order, nil
}

func (a Application) GetOrder(ctx context.Context, id int64) (domain.Order, error) {
	return a.db.Get(ctx, id)
}
