package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/abadojack/whatlanggo"
)

// Fungsi untuk menghitung frekuensi kata
func countWords(text string) map[string]int {
	// Mengubah teks menjadi huruf kecil
	text = strings.ToLower(text)

	// Menggunakan regex untuk memisahkan kata-kata
	re := regexp.MustCompile(`[a-zA-Z]+`)
	words := re.FindAllString(text, -1)

	// Menghitung frekuensi kata
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

// Fungsi untuk mendeteksi bahasa
func detectLanguage(text string) whatlanggo.Info {
	return whatlanggo.Detect(text)
}

// Fungsi untuk memisahkan teks berdasarkan kalimat
func splitTextByLanguage(text string) map[string]string {
	// Pisahkan teks berdasarkan titik atau tanda tanya
	sentences := regexp.MustCompile(`[.!?]+`).Split(text, -1)

	languageSegments := make(map[string]string)
	for _, sentence := range sentences {
		if sentence = strings.TrimSpace(sentence); len(sentence) > 0 {
			langInfo := detectLanguage(sentence)
			lang := langInfo.Lang.String()

			// Gabungkan kalimat ke dalam map per bahasa
			languageSegments[lang] += sentence + " "
		}
	}

	return languageSegments
}

func main() {
	// Membaca input teks dari pengguna
	fmt.Println("Masukkan teks:")
	reader := bufio.NewReader(os.Stdin)
	inputText, _ := reader.ReadString('\n')
	inputText = strings.TrimSpace(inputText)

	// Pisahkan teks berdasarkan bahasa
	languageTexts := splitTextByLanguage(inputText)

	// Cek berapa bahasa yang terdeteksi
	fmt.Printf("\nAda %d Bahasa yang terpakai, ini Frekuensinya\n", len(languageTexts))

	// Proses per bahasa
	for lang, text := range languageTexts {
		fmt.Printf("\nBahasa: %s\n", lang)

		// Mendapatkan frekuensi kata untuk bahasa ini
		wordCount := countWords(text)

		// Mengurutkan kata berdasarkan frekuensi kemunculan
		type wordFreq struct {
			Word  string
			Count int
		}

		var wordList []wordFreq
		for word, count := range wordCount {
			wordList = append(wordList, wordFreq{Word: word, Count: count})
		}

		sort.Slice(wordList, func(i, j int) bool {
			return wordList[i].Count > wordList[j].Count
		})

		// Tampilkan daftar kata dan frekuensinya
		for _, wf := range wordList {
			fmt.Printf("- %s: %d\n", wf.Word, wf.Count)
		}
	}
}
