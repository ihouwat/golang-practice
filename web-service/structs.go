package main

type widget struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type location struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Widgets []widget `json:"widgets"`
}

var locations = []location{
	{ID: 1, Name: "Igor's Warehouse", Widgets: []widget{
		{ID: 1, Name: "Widget A", Quantity: 100},
		{ID: 2, Name: "Widget B", Quantity: 200},
	}},
	{ID: 2, Name: "Gordon's Warehouse", Widgets: []widget{
		{ID: 3, Name: "Widget C", Quantity: 150},
		{ID: 4, Name: "Widget D", Quantity: 300},
	}},
	{ID: 3, Name: "Tin's Warehouse", Widgets: []widget{
		{ID: 5, Name: "Widget E", Quantity: 50},
		{ID: 6, Name: "Widget F", Quantity: 75},
	}},
	{ID: 4, Name: "Sang's Warehouse", Widgets: []widget{
		{ID: 7, Name: "Widget G", Quantity: 120},
		{ID: 8, Name: "Widget H", Quantity: 80},
	}},
	{ID: 5, Name: "Mike's Warehouse", Widgets: []widget{
		{ID: 9, Name: "Widget I", Quantity: 90},
		{ID: 10, Name: "Widget J", Quantity: 110}},
	},
}
