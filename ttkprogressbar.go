// Copyright 2022 luis
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ttkprogressbar

import (
	"fmt"

	"github.com/srenrique/ttkscreen"
)

type ProgressBar struct {
	Total          int    // The total number of steps to completion.
	Total_complet  int    // The number of steps completed
	Width          int    // The number of terminal columns for displaying a bar excluding other tokens. Defaults to total steps.
	RealWidth      int    // The real width including tokens
	Token_message  string // Token message
	Token_progress string // //
	Complet        string // character for complete steps [========]
	Uncomplet      string // character for uncplete steps [===     ]
}

// Bar Print initial bar
func (this *ProgressBar) Bar() {

	InitializeBar(this)
	SaveCursorPosition()
	PrintBar(this)

}

// private function
func SaveCursorPosition() {
	fmt.Print("\033[s")
}

// private function

func RestoreCursorPosition() {
	fmt.Print("\033[u\033[K")
}

// Advance one step in progress bar
func (this *ProgressBar) Advance() {
	this.Total_complet += 1

	PrintBar(this)
}

// SetPercent printbar in specyfic percent
// get percent in int type
func (this *ProgressBar) SetPercent(percent int) {
	advance := percent * this.Width / 100
	fmt.Println(advance)
	this.Total_complet = advance

	PrintBar(this)
}

// PrintBar
func PrintBar(this *ProgressBar) {
	// Makebar
	RestoreCursorPosition()
	var p_bar string

	p_bar = "["

	if this.Total_complet != 0 {
		for i := 0; i <= this.Total_complet-1; i++ {
			p_bar = p_bar + this.Complet
		}
	}

	for i := this.Total_complet; i < this.Width; i++ {
		p_bar = p_bar + this.Uncomplet
	}
	p_bar = p_bar + "]"

	fmt.Println(this.Token_message + p_bar)
}

func InitializeBar(this *ProgressBar) {
	// Calculate Total if total == 0 defautl 10
	if this.Total == 0 {
		this.Total = 10
	}

	//Default token message "Progress "
	if this.Token_message == "" {
		this.Token_message = "Progress "
	}

	// Width = total if width > consolesize cols then width = consolesize - tocken_message.leg else
	if this.Width == 0 {
		this.Width = this.Total
	}
	this.RealWidth = this.Width + len(this.Token_message) + 2
	cols, _ := ttkscreen.GetConsoleSize()
	if this.RealWidth > cols {
		this.Width = cols - len(this.Token_message) - 2
		this.RealWidth = cols
	}

	if this.Complet == "" {
		this.Complet = "="
	}
	if this.Uncomplet == "" {
		this.Uncomplet = " "
	}
}
