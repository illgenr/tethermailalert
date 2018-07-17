package main

import (
"fmt"
"io"
"log"
"os/exec"
"strconv"
)

func sendAlert(amount float64) {
				cmd := exec.Command("sendmail", "illgenr@gmail.com")
				stdin, err := cmd.StdinPipe()
				if err != nil {
								log.Fatal(err)
				}

				go func() {
								defer stdin.Close()
								io.WriteString(stdin, "From: tetheralerts@raleighillgen.com\n")
								io.WriteString(stdin,	"To: illgenr@gmail.com\n")
								io.WriteString(stdin,	"Subject: Tether Alert\n")
								io.WriteString(stdin, "Alert! " + strconv.FormatFloat(amount, 'f', 8, 64) + " Tethers printed!\n")
								io.WriteString(stdin,	".\n")
				}()

				out, err := cmd.CombinedOutput()
				if err != nil {
								log.Fatal(err)
				}

				fmt.Printf("%s\n", out)
}
