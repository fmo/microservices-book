package grpc

import (
	"context"
	"fmt"
	"github.com/fmo/microservices-book/payment/internal/application/core/domain"
	"github.com/fmo/microservices-proto/golang/payment"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a Adapter) Create(ctx context.Context, request *payment.CreatePaymentRequest) (*payment.CretePaymentResponse, error) {
	logrus.WithContext(ctx).Info("Creating payment...")
	newPayment := domain.NewPayment(request.UserId, request.OrderId, request.TotalPrice)
	result, err := a.api.Charge(ctx, newPayment)
	if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("failed to charge. %v ", err)).Err()
	}
	return &payment.CretePaymentResponse{PaymentId: result.ID}, nil
}
