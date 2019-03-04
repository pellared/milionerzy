package main

type stage struct {
	money        int
	isGuaranteed bool
	quiz
}

func getStages() []stage {
	awards := []struct {
		money        int
		isGuaranteed bool
	}{
		{money: 500, isGuaranteed: false},
		{money: 1000, isGuaranteed: true},
		{money: 2000, isGuaranteed: false},
		{money: 5000, isGuaranteed: false},
		{money: 10000, isGuaranteed: false},
		{money: 20000, isGuaranteed: false},
		{money: 40000, isGuaranteed: true},
		{money: 75000, isGuaranteed: false},
		{money: 125000, isGuaranteed: false},
		{money: 250000, isGuaranteed: false},
		{money: 500000, isGuaranteed: false},
		{money: 1000000, isGuaranteed: false},
	}

	quizes := getQuizes("quizes.csv")
	quizesIds, err := kPerm(len(awards), len(quizes))
	check(err)
	stages := make([]stage, len(awards))
	for i, award := range awards {
		stages[i] = stage{
			money:        award.money,
			isGuaranteed: award.isGuaranteed,
			quiz:         quizes[quizesIds[i]],
		}
	}

	return stages
}
