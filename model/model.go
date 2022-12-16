package model

import (
	"time"

	"github.com/google/uuid"
)

type OrderDispatch struct {
	DispatchID           uuid.UUID `json:"dispatch_id"`
	OrderDetailId        uuid.UUID `json:"order_detail_id"`
	PoolCodeExecutor     string    `json:"pool_code_executor"`
	OrderNo              string    `json:"order_no"`
	ItemNo               int32     `json:"item_no"`
	OrderArea            string    `json:"order_area"`
	OrderType            string    `json:"order_type"`
	GroupDestinationCode string    `json:"group_destination_code"`
	StartTime            time.Time `json:"start_time,omitempty"`
	EstimatedHours       time.Time `json:"estimated_hours,omitempty"`
	Currency             string    `json:"currency,omitempty"`
	TarifUsd             float64   `json:"tarif_usd,omitempty"`
	TarifIdr             float64   `json:"tarif_idr,omitempty"`
	VehicleType          string    `json:"vehicle_type,omitempty"`
	VehicleNo            string    `json:"vehicle_no,omitempty"`
	DriverId             string    `json:"driver_id,omitempty"`
	CustomerName         string    `json:"customer_name,omitempty"`
	CompanyName          string    `json:"company_name,omitempty"`
	Telephone            string    `json:"telephone,omitempty"`
	PickupPoint          string    `json:"pickup_point,omitempty"`
	DropPoint            string    `json:"drop_point,omitempty"`
	PickupLatlong        string    `json:"pickup_latlong"`
	CustomerSharing      bool      `json:"customer_sharing"`
	ListCustomerSharing  []string  `json:"list_customer_sharing,omitempty"`
	CreatedBy            string    `json:"created_by"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedBy            string    `json:"updated_by"`
	UpdatedAt            time.Time `json:"updated_at"`
	PaymentType          string    `json:"payment_type,omitempty"`
	PaymentMethod        string    `json:"payment_method,omitempty"`
	ExecutionType        int32     `json:"execution_type"`
	Channel              string    `json:"channel,omitempty"`
}
