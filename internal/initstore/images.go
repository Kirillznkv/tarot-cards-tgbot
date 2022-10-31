package initstore

import (
	"tarot-cards-tgbot/internal/model"
	"tarot-cards-tgbot/internal/store"
)

var (
	imges = []*model.Image{
		{
			URL:         "https://i0.wp.com/all-tarot.ru/wp-content/uploads/2020/01/lightseers-00-fool-tarot-meaning.png?resize=406%2C700&ssl=1",
			Name:        "Шут",
			Description: "Светлое пророчество: Новые начинания, потенциал, приключения, энтузиазм, пробуждение, невинность, оптимизм.\nТеневое пророчество: Самонадеянность, ошибочное мнение, что ответ у вас уже есть, импульсивность и скоропалительность решений, недостаток опыта, глупость, аналитический паралич (чрезмерное анализирование ситуации).",
		},
		{
			URL:         "https://i0.wp.com/all-tarot.ru/wp-content/uploads/2020/01/lightseers-01-magician-tarot-meaning.png?resize=406%2C700&ssl=1",
			Name:        "Маг",
			Description: "Светлое пророчество: Мастерство, природные таланты, проявление сил. Творчество, находчивость.\nТеневое пророчество: Неиспользованный потенциал, скрытые таланты. Сомнительные намерения. Манипулирование, эгоизм. Невнимательность или заблокированная творческая энергия.",
		},
		{
			URL:         "https://i0.wp.com/all-tarot.ru/wp-content/uploads/2020/01/lightseers-02-high-priestess-tarot-meaning.png?resize=406%2C700&ssl=1",
			Name:        "Верховная Жрица",
			Description: "Светлое пророчество: Мечты, мощные видения, экстрасенсорные способности, случайность, медитация. Самоанализ, интуиция, духовный опыт\nТеневое пророчество: Секреты, сплетни или неправды, заговоры, боязнь интуитивных способностей, невнимание к внутреннему голосу.",
		},
		//{
		//	URL:         "",
		//	Name:        "",
		//	Description: "",
		//},
		//{
		//	URL:         "",
		//	Name:        "",
		//	Description: "",
		//},
		//{
		//	URL:         "",
		//	Name:        "",
		//	Description: "",
		//},
		//{
		//	URL:         "",
		//	Name:        "",
		//	Description: "",
		//},
		//{
		//	URL:         "",
		//	Name:        "",
		//	Description: "",
		//},
		//{
		//	URL:         "",
		//	Name:        "",
		//	Description: "",
		//},
		//{
		//	URL:         "",
		//	Name:        "",
		//	Description: "",
		//},
		//{
		//	URL:         "",
		//	Name:        "",
		//	Description: "",
		//},
		//{
		//	URL:         "",
		//	Name:        "",
		//	Description: "",
		//},
		//{
		//	URL:         "",
		//	Name:        "",
		//	Description: "",
		//},
		//{
		//	URL:         "",
		//	Name:        "",
		//	Description: "",
		//},
		//{
		//	URL:         "",
		//	Name:        "",
		//	Description: "",
		//},
		//{
		//	URL:         "",
		//	Name:        "",
		//	Description: "",
		//},
		//{
		//	URL:         "",
		//	Name:        "",
		//	Description: "",
		//},
		//{
		//	URL:         "",
		//	Name:        "",
		//	Description: "",
		//},
		//{
		//	URL:         "",
		//	Name:        "",
		//	Description: "",
		//},
		//{
		//	URL:         "",
		//	Name:        "",
		//	Description: "",
		//},
		//{
		//	URL:         "",
		//	Name:        "",
		//	Description: "",
		//},
		//{
		//	URL:         "",
		//	Name:        "",
		//	Description: "",
		//},
	}
)

func Images(store store.Store) {
	for _, img := range imges {
		img.BeforeUse()
		store.Images().AddImage(img)
	}
}
