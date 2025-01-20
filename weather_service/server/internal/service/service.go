package service

/*
message Response {
    string city_name = 1;    // Название города
    float temperature = 2;   // Температура воздуха в градусах цельсия
    float precipitation = 3; // Осадки в процентах
    float wind_speed = 4;    // Скорость ветра м/с
    string condition = 5;    // Состояние (ясно, облачно или дождь)
}
*/
type weather struct {
	cityName string
	temperature float32
	precipitation float32
	windSpeed float32
	condition string
}

func (w *weather) GetSityName() string {
	return w.cityName
}

func (w *weather) GetTemperature() float32 {
	return w.temperature
}

func (w *weather) GetPercipitation() float32 {
	return w.precipitation
}

func (w *weather) GetWindSpeed() float32 {
	return w.windSpeed
}

func (w *weather) GetCondition() string {
	return w.condition
}

/* Интерфейс для самой погоды */
type Weather interface {
	GetSityName() string
	GetTemperature() float32
	GetPercipitation() float32
	GetWindSpeed() float32
	GetCondition() string
}

type GetWeather struct {}

func NewGetWeather() (*GetWeather, error) {
	return &GetWeather{}, nil
}

func (gw *GetWeather)GetWeatherInfo(sity string) (Weather, error) {
	return &weather{
		cityName: "New York",
		temperature: 23.1,
		precipitation: 40.0,
		windSpeed: 0.5,
		condition: "Sunny",
	}, nil
}
