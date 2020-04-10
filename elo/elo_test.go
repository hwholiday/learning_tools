package elo

import "testing"

func Test_EloRating(t *testing.T) {
	 EloRating(Elo{
		 A:  1500,
		 B:  1600,
		 Sa: 1,
	 })
}

func Test_Decimal(t *testing.T)  {
	t.Log(Decimal(22.222222222,"%.2f"))
	t.Log(Decimal(22.222222222,"%.0f"))
	t.Log(Decimal(22.6666666666,"%.2f"))
	t.Log(Decimal(22.66666666666,"%.0f"))
}
