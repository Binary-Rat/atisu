package atisu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	mock_atisu "github.com/Binary-Rat/atisu/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_GetCarsWithFilter(t *testing.T) {

	ctrl := gomock.NewController(t)
	HTTPClient := mock_atisu.NewMockHTTPClient(ctrl)

	cl := &Client{
		isDemo: false,
		client: HTTPClient,
		token:  "token",
	}
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(testBodyCars)),
	}
	HTTPClient.EXPECT().Do(gomock.Any()).Return(resp, nil)
	cars, err := cl.GetCarsWithFilter(1, 1, Filter{})
	assert.NoError(t, err)
	_ = cars // TODO

}

func Test_GetCityID(t *testing.T) {

	ctrl := gomock.NewController(t)
	HTTPClient := mock_atisu.NewMockHTTPClient(ctrl)

	cl := &Client{
		isDemo: false,
		client: HTTPClient,
		token:  "token",
	}
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer([]byte(testBodyCity))),
	}
	HTTPClient.EXPECT().Do(gomock.Any()).Return(resp, nil)
	cities, err := cl.GetCityID([]string{"Москва"})
	assert.NoError(t, err)
	fmt.Println(cities)
	assert.Equal(t, 3611, (*cities)["Москва"].CityID)
}

func Test_GetCityIDHTTP(t *testing.T) {
	cl := &Client{
		isDemo: true,
		client: &http.Client{},
		token:  "-",
	}
	cities, err := cl.GetCityID([]string{"Москва"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cities)
}

func Test_GetCarsWithFilterHTTP(t *testing.T) {
	cl := &Client{
		isDemo: true,
		client: &http.Client{},
		token:  "token",
	}
	filter := Filter{}
	filter.Dates.DateOption = "today"
	filter.From.ID = 2
	filter.From.Type = 0
	filter.To.ID = 2
	filter.To.Type = 0
	filter.Weight.Min = 0.5
	filter.Weight.Max = 0.5
	filter.Volume.Min = 0.5
	filter.Volume.Max = 0.5
	cars, err := cl.GetCarsWithFilter(1, 10, filter)
	if err != nil {
		t.Fatal(err)
	}
	var data any
	json.Unmarshal(cars, &data)
	t.Log(data)
}

func Test_ParseJSON(t *testing.T) {

	type args struct {
		body []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "1", args: args{body: []byte(testBodyCars)}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r Response
			err := json.Unmarshal(tt.args.body, &r)
			assert.NoError(t, err)
			fmt.Println(r.Trucks[0].AtiID)
		})
	}
}

const (
	testBodyCity = `{
    "Москва": {
        "is_success": true,
        "city_id": 3611,
        "street": null
    }
}`
	testBodyCars = `{
  "total_count": 0,
  "hidden_count": 0,
  "accounts": {
    "any-key": {
    "ati_id": "string",
    "firm_name": "string",
    "full_firm_name": "string",
    "brand": "string",
    "ownership": "string",
    "ownership_id": 0,
    "firm_type_id": 0,
    "firm_type": "string",
    "claims_sum": 0.5,
    "claims_count": 0,
    "bad_partner_mentions_count": 0,
    "bad_partner_firms_count": 0,
    "recommendation_count": 0,
    "recommendations_count": 0,
    "is_ati_partner": false,
    "is_odks_member": false,
    "city_id": 0,
    "passport": {
      "ati_data_match_point": 0.5,
      "account_lifetime_point": 0.5,
      "business_activity_point": 0.5,
      "round_table_point": 0.5,
      "claim_point": 0.5,
      "prof_activity_point": 0.5,
      "ati_administration_point": 0.5,
      "clones_point": 0.5,
      "egr_point": 0.5,
      "mass_registration_point": 0.5,
      "mass_founder_point": 0.5,
      "firm_lifetime_point": 0.5,
      "negative_points_sum": 0.5,
      "total_score": 0.5,
      "status": 0,
      "status_description": "string"
    },
    "contacts": [
      {
      "id": 0,
      "name": "string",
      "telephone": "string",
      "telephone_brand": "string",
      "telephone_region": "string",
      "email": false,
      "icq": "string",
      "mobile": "string",
      "mobile_brand": "string",
      "mobile_region": "string",
      "skype_name": "string",
      "fax": "string",
      "mobile_operator": "string",
      "city_id": 0
      }
    ],
    "user_lists": {
      "white": false,
      "neutral": false,
      "black": false,
      "emoji_firm_lists": [
      {
        "emoji": "string",
        "total_lists_count": 0,
        "head_lists": [
        {
          "list_name": "string",
          "list_id": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
          "emoji": "string",
          "type": 0,
          "fixed": false
        }
        ]
      }
      ]
    },
    "comment": "string"
    }
  },
  "trucks": [
    {
    "ati_id": "string",
    "truck_id": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
    "added_at": "1970-01-01T00:00:00.000Z",
    "updated_at": "1970-01-01T00:00:00.000Z",
    "contact_id1": 0,
    "contact_id2": 0,
    "department_id": 0,
    "note": "string",
    "first_date": "1970-01-01T00:00:00.000Z",
    "last_date": "1970-01-01T00:00:00.000Z",
    "periodicity_id": 0,
    "date_type": 0,
    "true_price": 0.5,
    "true_currency_id": 0,
    "transport": {
      "car_type": 0,
      "body_type": 0,
      "loading_type": 0,
      "weight": 0.5,
      "volume": 0.5,
      "truck_length": 0.5,
      "truck_width": 0.5,
      "truck_height": 0.5,
      "trailer_length": 0.5,
      "trailer_width": 0.5,
      "trailer_height": 0.5,
      "hydrolift": false,
      "partial_load": false,
      "koniki": false,
      "tir": false,
      "ekmt": false,
      "adr_types": 0
    },
    "loading": {
      "city_id": 0,
      "radius": 0,
      "distance": 0
    },
    "unloading_list": [
      {
      "main": false,
      "visible": false,
      "cash_sum": 0.5,
      "sum_with_nds": 0.5,
      "sum_without_nds": 0.5,
      "currency_id": 0,
      "point_type": 0,
      "point_id": 0,
      "radius": 0,
      "distance": 0
      }
    ],
    "payment_options": {
      "card": false,
      "torg": false
    },
    "show_contacts": false,
    "comments": [
      {
      "comment_id": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
      "message": "string",
      "contact_id": 0,
      "contact_name": "string",
      "comment_date_time": "1970-01-01T00:00:00.000Z"
      }
    ]
    }
  ],
  "responses": {
    "any-key": {
    "offers": [
      {
      "truck_id": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
      "addition_date": "1970-01-01T00:00:00.000Z",
      "change_date": "1970-01-01T00:00:00.000Z",
      "cash_sum": 0.5,
      "currency_id": 0,
      "sum_with_nds": 0.5,
      "nds_currency_id": 0,
      "sum_without_nds": 0.5,
      "not_nds_currency_id": 0,
      "note": "string",
      "prepay_percent": 0,
      "delay_payment_days": 0,
      "loading_city_id": 0,
      "unloading_city_id": 0,
      "weight": 0.5,
      "volume": 0.5,
      "car_delivery_date": "1970-01-01T00:00:00.000Z",
      "unload_payment": false,
      "coloading": false,
      "sender_firm_info": {
        "ati_id": "string",
        "contact_id": 0,
        "contact": {
        "name": "string",
        "phone": "string",
        "mobile_phone": "string",
        "city": "string",
        "icq": "string",
        "skype_name": "string",
        "fax": "string",
        "email": "string"
        },
        "name": "string",
        "score": 0.5,
        "status": 0
      }
      }
    ],
    "absents": [
      {
      "truck_id": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
      "add_date": "1970-01-01T00:00:00.000Z",
      "sender_firm_info": {
        "ati_id": "string",
        "contact_id": 0,
        "contact": {
        "name": "string",
        "phone": "string",
        "mobile_phone": "string",
        "city": "string",
        "icq": "string",
        "skype_name": "string",
        "fax": "string",
        "email": "string"
        },
        "name": "string",
        "score": 0.5,
        "status": 0
      }
      }
    ],
    "complaints": [
      {
      "truck_id": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
      "add_date": "1970-01-01T00:00:00.000Z",
      "change_date": "1970-01-01T00:00:00.000Z",
      "note": "string",
      "sender_firm_info": {
        "ati_id": "string",
        "contact_id": 0,
        "contact": {
        "name": "string",
        "phone": "string",
        "mobile_phone": "string",
        "city": "string",
        "icq": "string",
        "skype_name": "string",
        "fax": "string",
        "email": "string"
        },
        "name": "string",
        "score": 0.5,
        "status": 0
      }
      }
    ]
    }
  }
  }`
)
