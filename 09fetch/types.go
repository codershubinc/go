package main

type Name struct {
	FirstName   string `json:"first_name"`
	MiddleName  string `json:"middle_name"`
	LastName    string `json:"last_name"`
	FullName    string `json:"full_name"`
	Prefix      string `json:"prefix"`
	Title       string `json:"title"`
	Gender      string `json:"gender"`
	CountryCode string `json:"contryCode"` // Note: keeping the typo from API
}

type Devices struct {
	Browser  string `json:"browser"`
	OS       string `json:"os"`
	Location string `json:"location"`
}

type Info struct {
	Hobby []string `json:"hobby"`
}

type Prefs struct {
	Avatar  string  `json:"avatar"`
	Secret  string  `json:"secret"`
	Mode    string  `json:"mode"`
	Devices Devices `json:"devices"`
	Info    Info    `json:"info"`
}

type Avatar struct {
	Avatar1      string `json:"avatar1"`
	Avatar2      string `json:"avatar2"`
	Avatar3      string `json:"avatar3"`
	Banner       string `json:"banner"`
	Thumbnail    string `json:"thumbnail"`
	AvatarSecret string `json:"avatarSecret"`
}

type User struct {
	Name     Name   `json:"name"`
	ID       string `json:"id"`
	Password string `json:"password"`
	Prefs    Prefs  `json:"prefs"`
	Avatar   Avatar `json:"avatar"`
}

type Coordinates struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type TimeZone struct {
	Name   string `json:"name"`
	Offset string `json:"offset"`
	Zone   string `json:"zone"`
}

type Street struct {
	Name              string `json:"name"`
	Number            int    `json:"number"`
	EncryptedHomeName string `json:"encrypted_home_name"`
}

type CountryName struct {
	Country string `json:"country"`
	Alpha2  string `json:"alpha2"`
	Alpha3  string `json:"alpha3"`
	Numeric string `json:"numeric"`
}

type Country struct {
	Code string      `json:"code"`
	Name CountryName `json:"name"`
}

type Address struct {
	City        string      `json:"city"`
	Coordinates Coordinates `json:"co_ordinates"`
	State       string      `json:"state"`
	PostCode    int         `json:"post_code"`
	TimeZone    TimeZone    `json:"time_zone"`
	Street      Street      `json:"street"`
	Country     Country     `json:"country"`
}

type Data struct {
	User       User    `json:"user"`
	Address    Address `json:"address"`
	DocumentID string  `json:"document_id"`
}

type APIResponse struct {
	StatusCode int    `json:"statusCode"`
	Data       Data   `json:"data"`
	Message    string `json:"message"`
	Success    bool   `json:"success"`
}
