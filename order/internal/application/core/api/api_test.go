package api

import (
	"context"
	"github.com/fmo/microservices-book/order/internal/application/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockedPayment struct {
	mock.Mock
}

func (p *mockedPayment) Charge(ctx context.Context, order *domain.Order) error {
	args := p.Called(ctx, order)
	return args.Error(0)
}

type mockedDb struct {
	mock.Mock
}

func (d *mockedDb) Save(ctx context.Context, order *domain.Order) error {
	args := d.Called(ctx, order)
	return args.Error(0)
}

func (d *mockedDb) Get(ctx context.Context, id int64) (domain.Order, error) {
	args := d.Called(ctx, id)
	return args.Get(0).(domain.Order), args.Error(1)
}

func Test_Should_Place_Order(t *testing.T) {
	payment := new(mockedPayment)
	db := new(mockedDb)
	payment.On("Charge", mock.Anything, mock.Anything).Return(nil)
	db.On("Save", mock.Anything, mock.Anything).Return(nil)

	application := NewApplication(db, payment)
	_, err := application.PlaceOrder(context.Background(), domain.Order{
		CustomerID: 123,
		OrderItems: []domain.OrderItem{
			{
				ProductCode: "camera",
				UnitPrice:   12.3,
				Quantity:    3,
			},
		},
		CreatedAt: 0,
	})
	assert.Nil(t, err)
}
