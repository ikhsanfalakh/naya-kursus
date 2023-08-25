package tugas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/olekukonko/tablewriter"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func GetUsers(url string) ([]User, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var users []User
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func ShowData(url string) {
	users, err := GetUsers(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Username", "Email"})

	for _, user := range users {
		table.Append([]string{
			fmt.Sprintf("%d", user.ID),
			user.Name,
			user.Username,
			user.Email,
		})
	}
	table.Render()
}
