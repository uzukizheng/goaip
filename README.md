# goaip
goaip is a simple golang implementation of Baidu AIP, only implement the NLP related API's.


import (
        "github.com/charliecui/goaip"
        "fmt"
        "testing"
)

func getDefaultAIPNLP() *goaip.AIPNlp {
        return goaip.MakeAIPNlp("Your APP ID", "Your Key", "Your Secret")
}

func TestNLP_Lexer(t *testing.T) {
        an := getDefaultAIPNLP()
        res, _ := an.Lexer("清华大学怎么走")
        fmt.Println(res)
}

func TestNLP_Simnet(t *testing.T) {
        an := getDefaultAIPNLP()
        res, _ := an.Simnet("附近有什么好吃的", "周围的美食有那些")
        fmt.Println(res)
}

func TestNLP_DepParse(t *testing.T) {
        an := getDefaultAIPNLP()
        res, _ := an.DepParse("给我买一张到武汉的火车票")
        fmt.Println(res)
}

