package utils

import (
	"log"
	"riddles-server/database"
	"riddles-server/models"
)

func SeedDatabase() {
	// Create categories
	categories := []models.Category{
		{Name: "Математика"},
		{Name: "Что? Где? Когда?"},
		{Name: "Логика"},
		{Name: "Шутки"},
		{Name: "Загадки мира"},
	}

	for i := range categories {
		err := database.DB.FirstOrCreate(&categories[i], models.Category{Name: categories[i].Name}).Error
		if err != nil {
			log.Printf("Error creating category %s: %v", categories[i].Name, err)
		}
	}

	// Get category IDs
	var mathCategory, whqCategory, logicCategory, jokesCategory, worldRiddlesCategory models.Category
	database.DB.Where("name = ?", "Математика").First(&mathCategory)
	database.DB.Where("name = ?", "Что? Где? Когда?").First(&whqCategory)
	database.DB.Where("name = ?", "Логика").First(&logicCategory)
	database.DB.Where("name = ?", "Шутки").First(&jokesCategory)
	database.DB.Where("name = ?", "Загадки мира").First(&worldRiddlesCategory)

	// Create riddles for each category (20 per category)
	createMathRiddles(mathCategory.ID)
	createWhqRiddles(whqCategory.ID)
	createLogicRiddles(logicCategory.ID)
	createJokesRiddles(jokesCategory.ID)
	createWorldRiddles(worldRiddlesCategory.ID)

	log.Println("Database seeding completed")
}

func createMathRiddles(categoryID uint) {
	riddles := []models.Riddle{
		{
			Title:       "Простая арифметика",
			Description: "Сколько будет 2+2*2?",
			Answer:      "6",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Возраст",
			Description: "Если тройка больше двойки, то почему двойка больше тройки?",
			Answer:      "На замке",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Числа",
			Description: "Какое число делится на все числа без остатка?",
			Answer:      "0",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Последовательность",
			Description: "Продолжите последовательность: 1, 1, 2, 3, 5, 8, 13, ?",
			Answer:      "21",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Геометрия",
			Description: "Сколько граней у нового шестигранного карандаша?",
			Answer:      "8",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Время",
			Description: "Сколько минут в сутках?",
			Answer:      "1440",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Дроби",
			Description: "Какая дробь больше: 1/3 или 1/4?",
			Answer:      "1/3",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Проценты",
			Description: "Сколько процентов составляет четверть?",
			Answer:      "25",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Уравнение",
			Description: "Решите уравнение: x + 5 = 12",
			Answer:      "7",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Площадь",
			Description: "Чему равна площадь квадрата со стороной 5 см?",
			Answer:      "25",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Объем",
			Description: "Чему равен объем куба с ребром 3 см?",
			Answer:      "27",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Теорема",
			Description: "Квадрат гипотенузы равен...",
			Answer:      "Сумме квадратов катетов",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Пи",
			Description: "Сколько примерно равно число Пи?",
			Answer:      "3.14",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Простые числа",
			Description: "Какое наименьшее простое число?",
			Answer:      "2",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Алгебра",
			Description: "Чему равно (a+b)²?",
			Answer:      "a²+2ab+b²",
			CategoryID:  categoryID,
			Difficulty:  "hard",
		},
		{
			Title:       "Тригонометрия",
			Description: "Чему равен sin(90°)?",
			Answer:      "1",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Логарифмы",
			Description: "Чему равен log₁₀(100)?",
			Answer:      "2",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Производная",
			Description: "Чему равна производная x²?",
			Answer:      "2x",
			CategoryID:  categoryID,
			Difficulty:  "hard",
		},
		{
			Title:       "Интеграл",
			Description: "Чему равен интеграл от 0 до 1 функции f(x)=x?",
			Answer:      "1/2",
			CategoryID:  categoryID,
			Difficulty:  "hard",
		},
		{
			Title:       "Комбинаторика",
			Description: "Сколькими способами можно выбрать 2 предмета из 5?",
			Answer:      "10",
			CategoryID:  categoryID,
			Difficulty:  "hard",
		},
	}

	for i := range riddles {
		err := database.DB.FirstOrCreate(&riddles[i], models.Riddle{Title: riddles[i].Title}).Error
		if err != nil {
			log.Printf("Error creating riddle %s: %v", riddles[i].Title, err)
		}
	}
}

func createWhqRiddles(categoryID uint) {
	riddles := []models.Riddle{
		{
			Title:       "Классическая",
			Description: "Что больше: 1% от 1 рубля или 1 рубль?",
			Answer:      "1 рубль",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Логика",
			Description: "Можно ли зажечь спичку под водой?",
			Answer:      "Можно, в подводной лодке",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Смекалка",
			Description: "Какой рукой лучше размешивать чай?",
			Answer:      "Ложкой",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Практичность",
			Description: "На какой вопрос нельзя ответить 'да'?",
			Answer:      "Ты спишь?",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Внимание",
			Description: "Когда мы смотрим на цифру 2, а говорим 10?",
			Answer:      "Когда смотрим на часы",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Наблюдательность",
			Description: "Что можно приготовить, но нельзя съесть?",
			Answer:      "Уроки",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Сообразительность",
			Description: "Какой болезнью никто не болеет на суше?",
			Answer:      "Морской",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Эрудиция",
			Description: "Что находится между городом и селом?",
			Answer:      "Союз 'и'",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Игра слов",
			Description: "Когда человек бывает деревом?",
			Answer:      "Когда он со-сна",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Юмор",
			Description: "Что у цапли впереди, а у зайца сзади?",
			Answer:      "Буква 'ц'",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Фантазия",
			Description: "Можно ли с помощью вентилятора увеличить пламя свечи?",
			Answer:      "Можно, но только уменьшить",
			CategoryID:  categoryID,
			Difficulty:  "hard",
		},
		{
			Title:       "Необычность",
			Description: "Каких камней не бывает в море?",
			Answer:      "Сухих",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Логика",
			Description: "Что можно увидеть с закрытыми глазами?",
			Answer:      "Сон",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Смекалка",
			Description: "Под каким деревом сидит заяц, когда идет дождь?",
			Answer:      "Под мокрым",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Остроумие",
			Description: "Какой месяц короче всех?",
			Answer:      "Май, три буквы",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Интуиция",
			Description: "Что будет с вороной, когда ей исполнится 7 лет?",
			Answer:      "Пойдет восьмой",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Сообразительность",
			Description: "Какой остров летает?",
			Answer:      "Птичий",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Эрудиция",
			Description: "Какая птица носит название государства?",
			Answer:      "Турка",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Наблюдательность",
			Description: "Что у коровы впереди, а у быка позади?",
			Answer:      "Буква 'к'",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Юмор",
			Description: "Когда лошадь бывает хищным зверем?",
			Answer:      "Когда бежит рысью",
			CategoryID:  categoryID,
			Difficulty:  "hard",
		},
	}

	for i := range riddles {
		err := database.DB.FirstOrCreate(&riddles[i], models.Riddle{Title: riddles[i].Title}).Error
		if err != nil {
			log.Printf("Error creating riddle %s: %v", riddles[i].Title, err)
		}
	}
}

func createLogicRiddles(categoryID uint) {
	riddles := []models.Riddle{
		{
			Title:       "Классическая логика",
			Description: "Если в 12 часов ночи идет дождь, то можно ли ожидать, что через 72 часа будет солнечная погода?",
			Answer:      "Нет, будет ночь",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Парадокс",
			Description: "Что тяжелее: килограмм пуха или килограмм железа?",
			Answer:      "Одинаково",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Смекалка",
			Description: "У отца Мэри пять дочерей: 1. Чача 2. Чече 3. Чичи 4. Чочо. Вопрос: Как зовут пятую дочь?",
			Answer:      "Мэри",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Анализ",
			Description: "Шли два отца и два сына, нашли три апельсина. Решили делить - всем по одному досталось. Как это могло быть?",
			Answer:      "Это были дед, отец и сын",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Внимание",
			Description: "Петух, стоя на одной ноге, весит 5 кг. Сколько он будет весить, стоя на двух ногах?",
			Answer:      "5 кг",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Логика",
			Description: "Один поезд едет из Москвы в Санкт-Петербург с опозданием 10 минут, а другой - из Санкт-Петербурга в Москву с опозданием 20 минут. Какой из этих поездов будет ближе к Москве в момент их встречи?",
			Answer:      "Оба будут на одинаковом расстоянии",
			CategoryID:  categoryID,
			Difficulty:  "hard",
		},
		{
			Title:       "Сообразительность",
			Description: "Сколько яиц можно съесть натощак?",
			Answer:      "Одно",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Размышление",
			Description: "В комнате горело 50 свечей, 20 из них задули. Сколько останется?",
			Answer:      "20",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Анализ",
			Description: "Человек живет на 17-м этаже. На лифт он никогда не заходит. Когда идет дождь, он едет на лифте до 17-го этажа. А когда дождя нет, он доезжает до 10-го этажа, а дальше идет пешком. Почему?",
			Answer:      "Он карлик и не достает до кнопки 17",
			CategoryID:  categoryID,
			Difficulty:  "hard",
		},
		{
			Title:       "Смекалка",
			Description: "У меня две монеты на общую сумму 15 копеек. Одна из них не пятак. Что это за монеты?",
			Answer:      "10 копеек и 5 копеек",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Логика",
			Description: "Собака привязана к 10-метровой веревке, а пройти она может 300 метров. Как ей это удается?",
			Answer:      "Веревка ни к чему не привязана",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Анализ",
			Description: "3 курицы за 3 дня снесли 3 яйца. Сколько яиц снесут 12 кур за 12 дней?",
			Answer:      "48",
			CategoryID:  categoryID,
			Difficulty:  "hard",
		},
		{
			Title:       "Сообразительность",
			Description: "Вы заходите в тёмную комнату. В ней есть свеча, керосиновая лампа и газовая плита. Что вы зажжёте в первую очередь?",
			Answer:      "Спичку",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Внимание",
			Description: "Как можно поместить два литра молока в литровую бутылку?",
			Answer:      "Налить полбутылки, закупорить, перевернуть, залить оставшееся",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Логика",
			Description: "Стоит стена из бетона высотой 3 метра, по другую сторону стены - смертельная опасность. Прыгать с бетонной стены нельзя, падать с нее нельзя. Как человеку попасть на другую сторону?",
			Answer:      "Обойти вокруг",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Смекалка",
			Description: "Без рук, без ног, а двери и окна открывает.",
			Answer:      "Ветер",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Анализ",
			Description: "Что принадлежит вам, но другие используют его чаще, чем вы?",
			Answer:      "Ваше имя",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Сообразительность",
			Description: "Что становится больше, если его поставить вверх ногами?",
			Answer:      "Число 6",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Логика",
			Description: "Какой рукой лучше размешивать чай?",
			Answer:      "Той, в которой ложка",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Интеллект",
			Description: "Можно ли предсказать счет любого матча до его начала?",
			Answer:      "Да, 0:0",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
	}

	for i := range riddles {
		err := database.DB.FirstOrCreate(&riddles[i], models.Riddle{Title: riddles[i].Title}).Error
		if err != nil {
			log.Printf("Error creating riddle %s: %v", riddles[i].Title, err)
		}
	}
}

func createJokesRiddles(categoryID uint) {
	riddles := []models.Riddle{
		{
			Title:       "С юмором",
			Description: "Как написать 'сухая трава' четырьмя буквами?",
			Answer:      "Сено",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Веселая",
			Description: "Из какой посуды нельзя ничего поесть?",
			Answer:      "Из пустой",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Остроумная",
			Description: "Сколько месяцев в году имеют 28 дней?",
			Answer:      "Все",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Забавная",
			Description: "Что можно видеть с закрытыми глазами?",
			Answer:      "Сон",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Интересная",
			Description: "Что в огне не горит и в воде не тонет?",
			Answer:      "Лёд",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Смешная",
			Description: "Что нужно делать, когда видишь зелёного человечка?",
			Answer:      "Переходить дорогу",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Веселая",
			Description: "Что у коровы впереди, а у быка позади?",
			Answer:      "Буква К",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Юмористическая",
			Description: "Каких камней в море нет?",
			Answer:      "Сухих",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Забавная",
			Description: "Под каким кустом сидит заяц во время дождя?",
			Answer:      "Под мокрым",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Остроумная",
			Description: "Какой месяц короче всех?",
			Answer:      "Май",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "С юмором",
			Description: "Что будет с вороной, когда ей исполнится 7 лет?",
			Answer:      "Пойдёт восьмой",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Веселая",
			Description: "Когда лошадь бывает хищным зверем?",
			Answer:      "Когда бежит рысью",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Интересная",
			Description: "Какой болезнью никто на земле не болеет?",
			Answer:      "Морской",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Забавная",
			Description: "Можно ли зажечь спичку под водой?",
			Answer:      "В подводной лодке можно",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Смешная",
			Description: "Какой город летает?",
			Answer:      "Орёл",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Юмористическая",
			Description: "Какая река самая страшная?",
			Answer:      "Тигр",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Остроумная",
			Description: "Что с земли легко поднимешь, но далеко не закинешь?",
			Answer:      "Пух",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Веселая",
			Description: "Какой конь не ест овса?",
			Answer:      "Шахматный",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "С юмором",
			Description: "Кто под проливным дождём не намочит волосы?",
			Answer:      "Лысый",
			CategoryID:  categoryID,
			Difficulty:  "easy",
		},
		{
			Title:       "Забавная",
			Description: "Что у цапли впереди, а у зайца сзади?",
			Answer:      "Буква Ц",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
	}

	for i := range riddles {
		err := database.DB.FirstOrCreate(&riddles[i], models.Riddle{Title: riddles[i].Title}).Error
		if err != nil {
			log.Printf("Error creating riddle %s: %v", riddles[i].Title, err)
		}
	}
}

func createWorldRiddles(categoryID uint) {
	riddles := []models.Riddle{
		{
			Title:       "Американская",
			Description: "Что в Америке делают открытым, а в России - закрытым?",
			Answer:      "Холодильник",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Японская",
			Description: "Что у японцев перед нами, а у нас за нами?",
			Answer:      "Нос",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Французская",
			Description: "Что во Франции делают в лифте, а в Америке в самолёте?",
			Answer:      "Произносят 'пти-пти'",
			CategoryID:  categoryID,
			Difficulty:  "hard",
		},
		{
			Title:       "Английская",
			Description: "Что англичане делают стоя, а французы - лёжа?",
			Answer:      "Едят спаржу",
			CategoryID:  categoryID,
			Difficulty:  "hard",
		},
		{
			Title:       "Итальянская",
			Description: "Что итальянцы делают в четыре руки?",
			Answer:      "Играют в карты",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Немецкая",
			Description: "Что немцы делают в четыре ноги?",
			Answer:      "Ползут",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Русская",
			Description: "Что русские делают в три руки?",
			Answer:      "Тройной прыжок",
			CategoryID:  categoryID,
			Difficulty:  "medium",
		},
		{
			Title:       "Китайская",
			Description: "Что китайцы делают в пять рук?",
			Answer:      "Играют в маджонг",
			CategoryID:  categoryID,
			Difficulty:  "hard",
		},
		{
			Title:       "Индийская",
			Description: "Что индусы делают в шесть рук?",
			Answer:      "Молятся",
			CategoryID:  categoryID,
			Difficulty:  "hard",
		},
		{
			Title:       "Бразильская",
			Description: "Что бразильцы делают в семь ног?",
			Answer:      "Танцуют самбу",
			CategoryID:  categoryID,
			Difficulty:  "hard",
		},
		{
			Title:       "Австралийская",
			Description: "Что австралийцы делают в восемь ног?",
			Answer:      "Ловят кенгуру",
			CategoryID:  categoryID,
			Difficulty:  "hard",
		},
		{
			Title:       "Африканская",
			Description: "Что африканцы делают в девять рук?",
			Answer:      "Управляют баобабом",
			CategoryID:  categoryID,
			Difficulty:  "hard",
		},
		{
			Title:       "Европейская",
			Description: "Что европейцы делают в десять рук?",
			Answer:      "Играют в оркестре",
			CategoryID:  categoryID,
			Difficulty:  "hard",
		},
		{
			Title:       "Скандинавская",
			Description: "Что скандинавы делают в одиннадцать рук?",
			Answer:      "Строят викингов",
			CategoryID:  categoryID,
			Difficulty:  "hard",
		},
		{
			Title:       "Полярная",
			Description: "Что полярники делают в двенадцать рук?",
			Answer:      "Строят снеговика",
			CategoryID:  categoryID,
			Difficulty:  "hard",
		},
		{
			Title:       "Горная",
			Description: "Что альпинисты делают в тринадцать рук?",
			Answer:      "Лазают по скалам",
			CategoryID:  categoryID,
			Difficulty:  "hard",
		},
		{
			Title:       "Морская",
			Description: "Что моряки делают в четырнадцать рук?",
			Answer:      "Завязывают морские узлы",
			CategoryID:  categoryID,
			Difficulty:  "hard",
		},
		{
			Title:       "Космическая",
			Description: "Что космонавты делают в пятнадцать рук?",
			Answer:      "Управляют ракетой",
			CategoryID:  categoryID,
			Difficulty:  "hard",
		},
		{
			Title:       "Цирковая",
			Description: "Что цирковые артисты делают в шестнадцать рук?",
			Answer:      "Выступают в цирке",
			CategoryID:  categoryID,
			Difficulty:  "hard",
		},
		{
			Title:       "Мировая",
			Description: "Что все народы мира делают в бесконечность рук?",
			Answer:      "Живут",
			CategoryID:  categoryID,
			Difficulty:  "hard",
		},
	}

	for i := range riddles {
		err := database.DB.FirstOrCreate(&riddles[i], models.Riddle{Title: riddles[i].Title}).Error
		if err != nil {
			log.Printf("Error creating riddle %s: %v", riddles[i].Title, err)
		}
	}
}