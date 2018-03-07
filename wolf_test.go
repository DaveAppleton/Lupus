package main

import (
	"fmt"
	"testing"
)

var questions = []string{
	"Who is the prime minister of Singapore",
	"when was the first bitcoin block mined",
	"why is the earth round",
	"What is the price of gold in September 2017",
	"When is the next Halley's comet",
	"Who is Virgil Griffith",
}

var codeQuestions = []string{
	`EntityValue[EntityClass["Element",{EntityProperty["Element","AtomicNumber"]->TakeLargest[1]}],"AtomicNumber","EntityAssociation"]`,
	`FactorInteger[28784365782356923156792345]`,
	`PrimeQ[3248796234659586723928569268572345967882975696782345]`,
	`EntityValue[EntityClass["Element",{EntityProperty["Element","AtomicNumber"]->TakeLargest[1]}],"AtomicNumber","EntityAssociation"]`,
	`InterpolatingPolynomial[{{-1, 4}, {0, 2}, {1, 6}, {2,8}, {100,50}}, x]`,
	`Entity["Country", "UnitedStates"][EntityProperty["Country", "MurderNonnegligentManslaughter", {"Date" -> DateObject[{2011}]}]]`,
	`Entity["Element", "Gold"][EntityProperty["Element", "Price", {"Date" -> DateObject[{2016, 12}]}]]`,
	`Inverse[{{1.4, 2}, {3, -6.7}}]`,
}

func TestWolfTYpe1(t *testing.T) {
	for _, q := range questions {
		ans, err := askQuestionType1(q)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(ans)
		t.Log(ans)
	}
}

func TestWolfType2(t *testing.T) {
	for _, q := range questions {
		Q2, err := askQuestionType2(q)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(Q2.Result, Q2.Failure)
		t.Log(Q2)
	}
}

func TestWolfType3(t *testing.T) {
	for _, q := range questions {
		Q2, err := askQuestionType2(q)
		if err != nil {
			t.Error(err)
		}
		if len(Q2.Failure) != 0 {
			fmt.Println("Fail ", Q2.Failure)
			continue
		}
		ans, err := askQuestionType3(Q2.Interpretation)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(ans)
		t.Log(ans)
	}
}

func TestCodeQuestions(t *testing.T) {
	for _, q := range codeQuestions {
		ans, err := askQuestionType3(q)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(ans)
		t.Log(ans)

	}
}
