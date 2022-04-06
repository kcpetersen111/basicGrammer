package main

import (
	"math/rand"
	"time"
	"strings"
	"fmt"
)

// Grammar: All counting numbers
// 1. terminals: 0, 1, 2, 3, 4, 5, 6, 7, 8, 9
// 2. nonterminals: S, P, N, D
// 3. start symbol: S
// 4. production rules:
//  S → P | PN
//  P → 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9
//  N → ND | D
//  D → 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9

var (
	S = []string{"P", "PN"}
	P = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	N = []string{"ND", "D"}
	D = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
)

func main() {
	master := []string{"S"}
	rand.Seed(time.Now().UnixNano())
	running := true
	skew := 2
	
	for running {
		pass := false
		master = strings.Split(strings.Join(master, ""), "")
		fmt.Printf("%v \n", master)
		for i, m := range master {
			if m == "S" {
				if rand.Int() % skew == 0 {	// mod by skew to adjust the chances of it being true 
					master[i] = "P"
				} else {
					master[i] = "PN"
				}
				pass = true
			} else if m == "P" {
				master[i] = P[rand.Int() % len(P)]
				pass = true
			} else if m == "N" {
				if rand.Int() % skew == 0 {
					master[i] = "ND"
				} else {
					master[i] = "D"
				}
				pass = true
			} else if m == "D" {
				master[i] = D[rand.Int() % len(D)]
				pass = true
			}
		}
		if !pass{
			running = false
		}
	}
	fmt.Printf("%v \n", master)
}
