package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type SearchEnhanced struct {
	Packages        HotelEnhancedPackages `json:"packages,omitempty"`
	ComparisonRates ComparisonRates       `json:"comparison_rates,omitempty"`
	Detailed        bool                  `json:"detailed"`
	CreatedAt       *time.Time            `json:"-"`
}

type HotelEnhancedPackages map[string][]EnhancedPackage

type EnhancedPackage struct {
	*BasicPackage
	HotelID     HotelID     `json:"hotel_id"`
	RoomDetails RoomDetails `json:"room_details"`
}

type BasicPackage struct {
	Partner  string `json:"partner"`
	Supplier string `json:"supplier"`

	SupplierSellRate CurrencyValue `json:"supplier_sell_rate"`
	//	SupplierCommission *CurrencyValue `json:"supplier_commission"`
	ZUMATACommission *CurrencyValue `json:"zumata_commission,omitempty"` // Only applicable for commissionable rates

	SupplierHotelID       SupplierHotelID          `json:"supplier_hotel_id"`
	SupplierRoomDetails   SupplierRoomDetails      `json:"supplier_room_details"`
	SupplierDetails       map[string]string        `json:"supplier_details"`
	SupplierRoomIDs       []SupplierRoomID         `json:"supplier_room_ids"`
	TaxesAndFees          map[string]CurrencyValue `json:"taxes_and_fees,omitempty"`
	IsBundledRate         bool                     `json:"is_bundled_rate"`
	SupplierLoyaltyPoints *LoyaltyPoints           `json:"supplier_loyalty_points,omitempty"`
}

type RoomDetails struct {
	Description       string         `json:"description" valid:"MinSize(1)"`
	Food              FoodCode       `json:"food"`
	NonRefundable     *bool          `json:"non_refundable,omitempty"`
	RoomType          string         `json:"room_type"`
	RoomView          string         `json:"room_view"`
	SupplierBedChoice bool           `json:"supplier_bed_choice,omitempty"`
	Beds              map[string]int `json:"beds"`
	Extras            []string       `json:"extras,omitempty"`
	Restrictions      []string       `json:"restrictions,omitempty"`
	Score             float32        `json:"score"`

	// These fields carry un-normalized values from the supplier
	SupplierDescription    string `json:"supplier_description,omitempty"`
	SupplierBedDescription string `json:"supplier_bed_description,omitempty"`
}

type ComparisonRates map[string][]ComparisonRate

type ComparisonRate struct {
	Type       string        `json:"type"`
	Rate       CurrencyValue `json:"rate"`
	Provider   string        `json:"provider"`
	DeepLinkID string        `json:"deeplink_id"`
}

type HotelID string

type SupplierHotelID string

type SupplierRoomID struct {
	ID string `json:"id,omitempty"`
}

type SupplierRoomDetails struct {
	Description          string                `json:"description"`
	Food                 FoodCode              `json:"food"`
	NonRefundable        *bool                 `json:"non_refundable,omitempty"`
	RoomType             string                `json:"room_type" valid:"MinSize(1)"`
	RoomView             string                `json:"room_view"`
	RoomDescriptionHints *RoomDescriptionHints `json:"-"`
}

type RoomDescriptionHints struct {
	BedHint  string // TODO:: @kevin enable other fields
	RoomHint string
	// ViewHint         string
	// ExtrasHint       string
	// RestrictionsHint string
}

type CurrencyValue struct {
	Currency string  `json:"currency"`
	Value    float64 `json:"value"`
}

type LoyaltyPoints struct {
	Amount float64 `json:"amount"`
	Units  string  `json:"units"`
}

type FoodCode int

var (
	hotelPackages *SearchEnhanced
)

func init() {
	file, err := os.Open("./data.json")
	if err != nil {
		log.Fatalf("unable to open data file due to %v", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalf("unable to get file info due to %v", err)
	}

	log.Printf("Read data file of %v bytes", fileInfo.Size())

	hotelPackages = &SearchEnhanced{}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(hotelPackages)
	if err != nil {
		log.Fatalf("unable to unmarshal data file due to %v", err)
	}
}
