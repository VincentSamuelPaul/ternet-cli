package cmd_test

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

var (
	binName = "ternet"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")
	
	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}
	fmt.Println("Running tests...")
	result := m.Run()
	fmt.Println("cleaning up...")
	os.Remove(binName)
	os.Exit(result)
}

func TestTernetCLI(t *testing.T) {

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	cmdPath := filepath.Join(dir, binName)

	// 1st test - Login User
	t.Run("LoginUser", func(t *testing.T) {
		username := "vincent\n"
		password := "vince123\n"

		cmd := exec.Command(cmdPath, "-login")
		stdin, err := cmd.StdinPipe()
		if err != nil {
			t.Fatal(err)
		}
		go func() {
			defer stdin.Close()
			io.WriteString(stdin, username)
			io.WriteString(stdin, password)
		}()
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(string(out))
	})

	// 2nd test - SignUp User
	t.Run("SignUp", func(t *testing.T) {
		username := "nigga\n"
		password := "nigga123\n"

		cmd := exec.Command(cmdPath, "-signup")
		stdin, err := cmd.StdinPipe()
		if err != nil {
			t.Fatal(err)
		}
		go func() {
			defer stdin.Close()
			io.WriteString(stdin, username)
			io.WriteString(stdin, password)
		}()
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(string(out))
	})

	// 3rd test - Create New Post
	t.Run("NewPost", func(t *testing.T) {
		username := "samuel\n"
		data := "IOT is the next big thing.\n"

		cmd := exec.Command(cmdPath, "-newpost")
		stdin, err := cmd.StdinPipe()
		if err != nil {
			t.Fatal(err)
		}
		go func() {
			defer stdin.Close()
			io.WriteString(stdin, username)
			io.WriteString(stdin, data)
		}()
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(string(out))
	})

	// 4th test - Browse Posts
	// t.Run("BrowsePosts", func(t *testing.T) {
	// 	cmd := exec.Command(cmdPath, "-browse")
	// 	for i := 0; i < 3; i++ {
	// 		out, err := cmd.CombinedOutput()
	// 		if err != nil {
	// 			t.Fatal(err)
	// 		}
	// 		fmt.Println(string(out))
	// 		stdin, err := cmd.StdinPipe()
	// 		if err != nil {
	// 			t.Fatal(err)
	// 		}
	// 		go func() {
	// 			defer stdin.Close()
	// 			io.WriteString(stdin, "\t")
	// 		}()
	// 	}
	// })
}