package main

import (
	"fmt"
)

const (
	NMAX int = 2000  // Adjust NMAX to accommodate both user types
	DMAX int = 100   // Adjust DMAX for the maximum number of doctors
	chat int = 10000 // Maximum number of forum posts
)

type User struct {
	Name     string
	Email    string
	Password string
	Role     string // Role can be "Patient" or "Doctor"
}

type ConsultationData struct {
	UserDetails       User
	Question          string
	Responses         [NMAX]ResponseData
	NumberOfResponses int
	Tags              string
}

type ResponseData struct {
	UserDetails User
	Response    string
}

type pesan struct {
	nama, post, tag1, tag2, tag3 string
	noPesan                      int
}

var forumapk [chat]pesan

type daftarpasien [NMAX]User // Array of patients
type daftardokter [DMAX]User // Array of doctors

var Users [NMAX]User // Array of all users (patients and doctors)
var Consultations [NMAX]ConsultationData

var NumberOfUsers int         // Total number of users
var NumberOfConsultations int // Total number of consultations

var LoggedInUser User // Current logged-in user

func main() {
	var (
		choice int
	)

	for {
		fmt.Println("               Main Menu               ")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Lihat Forum Konsultasi")
		fmt.Println("4. Exit")

		fmt.Print("Pilih: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			Register()
		case 2:
			Login()
		case 3:
			ViewForum()
		case 4:
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func Register() {
	var newUser User

	fmt.Println("                Registration                  ")
	fmt.Print("Nama: ")
	fmt.Scan(&newUser.Name)
	fmt.Print("Password: ")
	fmt.Scan(&newUser.Password)
	fmt.Print("Kamu pasien atau dokter? : ")
	fmt.Scan(&newUser.Role)

	if isAlreadyRegistered(newUser.Password) {
		fmt.Println("Password Terdaftar")
	} else {
		Users[NumberOfUsers] = newUser
		NumberOfUsers++
		fmt.Println("Pendaftaran sukses")
	}
}

func isAlreadyRegistered(password string) bool {
	for i := 0; i < NumberOfUsers; i++ {
		if Users[i].Password == password {
			return true
		}
	}
	return false
}

func Login() {
	var status, password string

	fmt.Println("                 Login                  ")
	fmt.Print("Login dokter atau pasien? ")
	fmt.Scan(&status)
	fmt.Print("Password : ")
	fmt.Scan(&password)
	fmt.Println("------------------------------------------")

	if isAlreadyRegistered(password) {
		fmt.Println("Login successful")

		// Find the logged-in user
		for i := 0; i < NumberOfUsers; i++ {
			if Users[i].Password == password {
				LoggedInUser = Users[i] // Set LoggedInUser to the logged-in user
				break
			}
		}

		if status == "Pasien" || status == "pasien" {
			PatientConsultationMenu()
		} else if status == "Dokter" || status == "dokter" {
			DoctorConsultationMenu()
		}
	} else {
		fmt.Println("Login failed")
	}
}

func ViewForum() {
	var choice int

	fmt.Println("          Consultation Forum          ")
	fmt.Println("1. Lihat pertanyaan secara terurut menurun")
	fmt.Println("2. Lihat pertanyaan secara terurut menaik")
	fmt.Println("3. Exit")
	fmt.Print("Pilih : ")
	fmt.Scan(&choice)

	if choice == 1 {
		ViewQuestionsDescending()
	} else if choice == 2 {
		ViewQuestionsAscending()
	} else if choice == 3 {
		return
	} else {
		fmt.Println("Invalid choice")
	}
}

func ViewQuestionsDescending() {
	// Display all questions and responses in descending order based on tags
	for i := 1; i < NumberOfConsultations; i++ {
		key := Consultations[i]
		j := i - 1

		for j >= 0 && Consultations[j].Tags < key.Tags {
			Consultations[j+1] = Consultations[j]
			j = j - 1
		}
		Consultations[j+1] = key
	}

	for i := 0; i < NumberOfConsultations; i++ {
		DisplayQuestionAndResponses(i)
	}
}

func ViewQuestionsAscending() {
	// Display all questions and responses in ascending order based on tags
	for i := 0; i < NumberOfConsultations-1; i++ {
		minIndex := i
		for j := i + 1; j < NumberOfConsultations; j++ {
			if Consultations[j].Tags < Consultations[minIndex].Tags {
				minIndex = j
			}
		}
		Consultations[i], Consultations[minIndex] = Consultations[minIndex], Consultations[i]
	}

	for i := 0; i < NumberOfConsultations; i++ {
		DisplayQuestionAndResponses(i)
	}
}

func DisplayQuestionAndResponses(index int) {
	// Display the question and responses
	fmt.Printf("%d. Pertanyaan: %s ; Penanya: %s (%s)\n", (index + 1), Consultations[index].Question, Consultations[index].UserDetails.Name, Consultations[index].UserDetails.Role)
	for j := 0; j < Consultations[index].NumberOfResponses; j++ {
		fmt.Printf("\tResponse %d: %s ; Responded by: %s (%s)\n", (j + 1), Consultations[index].Responses[j].Response, Consultations[index].Responses[j].UserDetails.Name, Consultations[index].Responses[j].UserDetails.Role)
	}
}

// Add other functions like PatientConsultationMenu, DoctorConsultationMenu, PostQuestion, RespondToQuestion, etc.
func menuforum() {
	fmt.Println("Selamat datang di forum! Silakan pilih opsi di bawah ini:")
	fmt.Println("1. Cari posting berdasarkan tag atau kata kunci")
	fmt.Println("2. Lihat semua posting")
	fmt.Println("3. Tambahkan posting baru")
	fmt.Println("4. Kembali ke menu utama")
	fmt.Println()
}

func addPost() {
	var question, tag string

	fmt.Println("                Post Question              ")
	fmt.Print("Kasih Pertanyaan: ")
	fmt.Scan(&question)
	fmt.Println("Berikan tag pertanyaan dengan suatu kata agar mudah untuk mencari pertanyaan mu")
	fmt.Print("tag pertanyaan : ")
	fmt.Scan(&tag)

	Consultations[NumberOfConsultations].UserDetails = LoggedInUser
	Consultations[NumberOfConsultations].Question = question
	Consultations[NumberOfConsultations].Tags = tag
	Consultations[NumberOfConsultations].NumberOfResponses = 0
	NumberOfConsultations++

	fmt.Println("Question posted successfully")
}

func search() {
	var choice int
	var tag string

	fmt.Println("             Search Question Menu            ")
	fmt.Println("1. Search pertanyaan menurun")
	fmt.Println("2. Search pertanyaan menaik")
	fmt.Println("3. Back to Consultation Menu")
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Print("tag pertanyaan yang ingin dicari ")
		fmt.Scan(&tag)
		SearchQuestionsDescending(tag)
	} else if choice == 2 {
		fmt.Print("tag pertanyaan yang ingin dicari ")
		fmt.Scan(&tag)
		SearchQuestionsAscending(tag)
	} else if choice == 3 {
		return
	} else {
		fmt.Println("Invalid choice")
	}
}

func SearchQuestionsDescending(tag string) {
	// Implement searching questions in descending order based on tags
	// This function would need to integrate the binary search logic from your previous implementation
	// and display questions and responses accordingly.
}

func SearchQuestionsAscending(tag string) {
	// Implement searching questions in ascending order based on tags
	// This function would need
}
