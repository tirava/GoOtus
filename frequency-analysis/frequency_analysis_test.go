/*
 * HomeWork-3: Frequency Analysis tests
 * Created on 13.09.19 19:20
 * Copyright (c) 2019 - Eugene Klimov
 */

package frequency

import (
	"reflect"
	"testing"
)

const MaxWords = 10

var testCases = []struct {
	description string
	input       string
	output      Frequency
}{
	{
		"empty string",
		"",
		Frequency{},
	},
	{
		"count one word",
		"word",
		Frequency{"word": 1},
	},
	{
		"count one of each word",
		"one of each - каждого по одному",
		Frequency{"each": 1, "of": 1, "one": 1, "каждого": 1, "по": 1, "одному": 1},
	},
	{
		"multiple occurrences of a word",
		"One fish two fish red fish blue fish Отус Отус Отус",
		Frequency{"blue": 1, "fish": 4, "one": 1, "red": 1, "two": 1, "отус": 3},
	},
	{
		"ignore punctuation and apostrophes",
		"Car: 'carpet' as java:\n javascript!!&@$%^&",
		Frequency{"as": 1, "car": 1, "carpet": 1, "java": 1, "javascript": 1},
	},
	{
		"include numbers",
		"Testing, 1, 2, 3, 3 testing",
		Frequency{"1": 1, "2": 1, "testing": 2, "3": 2},
	},
	{
		"multiple spaces not detected as a word",
		"    multiple   whitespaces       ",
		Frequency{"multiple": 1, "whitespaces": 1},
	},
	{
		"10 - simple cyrillic text",
		"От запаха краски полковник плохо соображал, однако полковник плохо знал: стоит выпить большую чашку кофе, и он снова придет в норму. Вот только плохо этот доклад после краски...",
		Frequency{"плохо": 3, "полковник": 2, "краски": 2, "большую": 1, "в": 1, "вот": 1, "выпить": 1, "доклад": 1, "запаха": 1, "знал": 1},
	},
	{
		"10 - cyrillic refrain text",
		`
В те времена, когда роились грезы
В сердцах людей, прозрачны и ясны,
Как хороши, как свежи были розы
Моей любви, и славы, и весны!

Прошли лета, и всюду льются слезы…
Нет ни страны, ни тех, кто жил в стране…
Как хороши, как свежи ныне розы
Воспоминаний о минувшем дне!

Но дни идут - уже стихают грозы.
Вернуться в дом Россия ищет троп…
Как хороши, как свежи будут розы,
Моей страной мне брошенные в гроб!
		`,
		Frequency{"брошенные": 1, "будут": 1, "в": 5, "и": 4, "как": 6, "моей": 2, "ни": 2, "розы": 3, "свежи": 3, "хороши": 3},
	},
	{
		"10 - cyrillic lexical repetition Simonov",
		`
Жди меня, и я вернусь.
Только очень жди,
Жди, когда наводят грусть
Желтые дожди,
Жди, когда снега метут,
Жди, когда жара,
Жди, когда других не ждут,
Позабыв вчера.
Жди, когда из дальних мест
Писем не придет,
Жди, когда уж надоест
Всем, кто вместе ждет.
		`,
		Frequency{"вернусь": 1, "вместе": 1, "всем": 1, "вчера": 1, "грусть": 1, "дальних": 1, "дожди": 1, "жди": 8, "когда": 6, "не": 2},
	},
	{
		"10 - big english text",
		`
It is a Sunday afternoon in the countryside. Is it hot? A small boy is in a rice field. He is ten years old. When is his birthday? He walks up a hill. Is it a big hill? He goes into a graveyard. Is it an old graveyard? He sees his father’s gravestone. What colour is it? He sits down in front of the gravestone. Is he sad?
Suddenly the small boy hears a sound. Is it loud? The boy turns around. He sees a man. Is he tall? The man has very short hair. He has handcuffs on his wrists. The man grabs the boy. Is the boy afraid? The man turns the boy upside down. He shakes him.
There is a rice ball in the boy’s pocket. Is it round or triangular? Is it very big? The rice ball falls out. The man picks up the rice ball. He eats it. Is it tasty? He eats very quickly because he is very very hungry. The boy feels sorry for the hungry man. The man asks for food. He asks for some tools too. He wants to take the handcuffs off.
The boy goes to get food. The man sleeps on the hill behind a tree near the graveyard. Is he very tired? Is he afraid of ghosts?
Two hours later the boy comes back. He has some bread. Is it fresh? He also has some cooked rice. Is it hot? Is it in a bag? He has some fruit too. What kind of fruit is it? He has a bottle of sake too. Is it expensive sake? He has some tools too. What tools are they?
The man is very grateful. He eats the food. He takes the handcuffs off. He runs away.
Two days later the boy is in the rice fields again. Is it a warm day? The boy hears a sound. It is a police siren. A police car comes. What colour is it? Four policemen get out. Are they young or old? They run up the hill. They catch the man. They put handcuffs on him. They put him in the car? The car drives away. Is it fast? The small boy watches the car. He thinks about the man. He feels sorry for the man.
The man is an escaped prisoner. The police take him back to prison. Are the police nice to him? The man is from a poor family. He lives in prison now. Where is the prison? Is it a big prison?
The boy lives with his aunt and his uncle. His father is dead and his mother is dead too. Is he lonely? He lives in an old house in a village. Is it a big house? Is it a traditional Japanese farmhouse?
The boy likes school. Is his school far from his house? What is the name of his favourite teacher? The boy wants to go to university. Is university expensive? The boy wants to go to university but he cannot go to university because his family is poor. Is his uncle a farmer?
At night the boy thinks about his father and his mother and he thinks about school and he thinks about the escaped prisoner and he thinks about his dream. He wants to go to university. Is he happy?
`,
		Frequency{"a": 22, "boy": 18, "he": 36, "his": 15, "in": 10, "is": 45, "it": 22, "man": 14, "the": 49, "to": 11},
	},
}

func TestWordCount(t *testing.T) {
	for _, tt := range testCases {
		expected := tt.output
		actual := CountFrequency(tt.input, MaxWords)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%s\n\tExpected: %v\n\tGot: %v", tt.description, expected, actual)
			continue
		}
		t.Logf("PASS: %s", tt.description)
	}
}

func BenchmarkWordCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range testCases {
			CountFrequency(tt.input, MaxWords)
		}
	}
}
