package elo

import "testing"

func Test_EloRating(t *testing.T) {
	 EloRating(Elo{
		 A:  1500,
		 B:  1600,
		 Sa: 0,
	 })
}
